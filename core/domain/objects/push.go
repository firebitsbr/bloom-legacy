package objects

import (
	"context"

	"gitlab.com/bloom42/bloom/core/api"
	"gitlab.com/bloom42/bloom/core/api/model"
	"gitlab.com/bloom42/bloom/core/db"
	"gitlab.com/bloom42/bloom/core/domain/groups"
	"gitlab.com/bloom42/bloom/core/domain/kernel"
	"gitlab.com/bloom42/bloom/core/domain/keys"
	"gitlab.com/bloom42/bloom/core/domain/users"
	"gitlab.com/bloom42/gobox/crypto"
	"gitlab.com/bloom42/gobox/graphql"
	"gitlab.com/bloom42/gobox/rz"
	"gitlab.com/bloom42/gobox/rz/log"
)

func push() error {
	var err error
	client := api.Client()
	var storedObjects []Object
	objectsToPush := map[string][]Object{}
	var masterKey []byte
	ctx := context.Background()

	input := model.PushInput{
		Repositories: []*model.RepositoryPushInput{},
	}

	// find objects that need to be pushed
	storedObjects, err = FindOutOfSyncObjects(ctx, nil)
	if err != nil {
		return err
	}
	// if no objects to push, abort
	if len(storedObjects) == 0 {
		return nil
	}

	for _, object := range storedObjects {
		var groupID string

		if object.GroupID != nil {
			groupID = object.GroupID.String()
		}
		objectsToPush[groupID] = append(objectsToPush[groupID], object)
	}

	masterKey, err = keys.FindUserMasterKey(ctx, nil)
	if err != nil {
		return err
	}
	defer crypto.Zeroize(masterKey) // clear masterKey from memory

	myGroups, err := groups.FindGroups(ctx, nil)
	if err != nil {
		return err
	}

	// format and encrypt objects
	for _, group := range myGroups.Groups {
		repo := &model.RepositoryPushInput{
			CurrentState: group.State,
			GroupID:      &group.ID,
		}
		input.Repositories = append(input.Repositories, repo)
	}
	myRepo := &model.RepositoryPushInput{
		CurrentState: *kernel.Me.State,
		GroupID:      nil,
	}
	input.Repositories = append(input.Repositories, myRepo)

	// for each repo, add the outOfSync objects
	for i, repo := range input.Repositories {
		var groupIDStr string
		objectsPushInput := []*model.ObjectInput{}

		if repo.GroupID != nil {
			groupIDStr = repo.GroupID.String()
		}
		for _, object := range objectsToPush[groupIDStr] {
			var objectToPush *model.ObjectInput

			objectToPush, err = encryptObject(object, masterKey, compressAlgoSnappy)
			if err != nil {
				return err
			}
			objectsPushInput = append(objectsPushInput, objectToPush)
		}
		input.Repositories[i].Objects = objectsPushInput
	}

	var resp struct {
		Push *model.Push `json:"push"`
	}
	req := graphql.NewRequest(`
		mutation ($input: PushInput!) {
			push(input: $input) {
				repositories {
					oldState
					newState
					groupID
				}
			}
		}
	`)
	req.Var("input", input)

	err = client.Do(ctx, req, &resp)
	if err != nil {
		log.Error("error pushing objects", rz.Err(err))
		return err
	}

	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}

	// update groups and me states
	for _, repo := range resp.Push.Repositories {
		if repo.GroupID != nil {
			err = groups.SaveGroupState(ctx, tx, *(repo.GroupID), repo.NewState)
		} else {
			kernel.Me.State = &repo.NewState
			err = users.SaveMe(ctx, tx, kernel.Me)
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// delete empty object
	deleteEmptyObjectQuery := "DELETE FROM objects WHERE length(data) < 3 AND out_of_sync = ?"
	_, err = tx.Exec(deleteEmptyObjectQuery, true)
	if err != nil {
		tx.Rollback()
		return err
	}

	// update outofSync object to not OufOfSync
	updateOutOfSyncObjectsQuery := "UPDATE objects SET out_of_sync = ? WHERE out_of_sync = ?"
	_, err = tx.Exec(updateOutOfSyncObjectsQuery, false, true)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
