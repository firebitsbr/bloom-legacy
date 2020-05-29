// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package messages

import (
	"time"

	"gitlab.com/bloom42/bloom/core/api/model"
	"gitlab.com/bloom42/gobox/uuid"
)

type AddPaymentMethodParams struct {
	StripePublicKey *string        `json:"stripePublicKey"`
	GroupID         *uuid.UUID     `json:"groupId"`
	Card            *NewStripeCard `json:"card"`
}

type CalcParams struct {
	Expression string `json:"expression"`
}

type CalcResult struct {
	Result string `json:"result"`
}

type CalendarCreateEventParams struct {
	Title       string    `json:"title"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
}

type CalendarDeleteEventParams struct {
	ID []byte `json:"id"`
}

type CalendarListEventsParams struct {
	StartAt *time.Time `json:"startAt"`
	EndAt   *time.Time `json:"endAt"`
}

type CreateNoteParams struct {
	Title   string     `json:"title"`
	Body    string     `json:"body"`
	Color   string     `json:"color"`
	GroupID *uuid.UUID `json:"groupId"`
}

type DashboardData struct {
	Metadata *model.BloomMetadata  `json:"metadata"`
	Users    *model.UserConnection `json:"users"`
}

// type Contacts struct {
//   Contacts []Contact `json:"contacts"`
// }
// type CreateContactParams struct {
//   DeviceID      string        `json:"deviceId"`
//   FirstName     string        `json:"firstName"`
//   LastName      string        `json:"lastName"`
//   Notes         string        `json:"notes"`
//   Birthday      *time.Time    `json:"birthday"`
//   BloomUsername string        `json:"bloomUsername" db:"bloom_username"`
//   Organizations Organizations `json:"organizations"`
//   Addresses     Addresses     `json:"addresses"`
//   Emails        Emails        `json:"emails"`
//   Phones        Phones        `json:"phones"`
//   Websites      Websites      `json:"websites"`
// }
type DeleteContactParams struct {
	ID []byte `json:"id"`
}

type DeleteNoteParams struct {
	ID []byte `json:"id"`
}

type Empty struct {
	Empty *bool `json:"empty"`
}

type FetchGroupProfileParams struct {
	ID uuid.UUID `json:"id"`
}

type GroupBillingProfile struct {
	Group           *model.Group                 `json:"group"`
	BillingPlans    *model.BillingPlanConnection `json:"billingPlans"`
	StripePublicKey *string                      `json:"stripePublicKey"`
}

type GroupsCancelInvitationParams struct {
	InvitationID uuid.UUID `json:"invitationID"`
}

type GroupsCreateParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GroupsFetchDetailsParams struct {
	GroupID uuid.UUID `json:"groupID"`
}

type GroupsFetchMembersParams struct {
	GroupID uuid.UUID `json:"groupID"`
}

type GroupsInviteUserParams struct {
	GroupID  uuid.UUID `json:"groupID"`
	Username string    `json:"username"`
}

type MyBillingProfile struct {
	Me              *model.User                  `json:"me"`
	BillingPlans    *model.BillingPlanConnection `json:"billingPlans"`
	StripePublicKey string                       `json:"stripePublicKey"`
}

type NewStripeCard struct {
	Number   string `json:"number"`
	ExpMonth string `json:"expMonth"`
	ExpYear  string `json:"expYear"`
	Cvc      string `json:"cvc"`
}

type PreferencesDeleteParams struct {
	Key string `json:"key"`
}

type PreferencesGetParams struct {
	Key string `json:"key"`
}

type PreferencesSetParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UserBillingProfile struct {
	User         *model.User                  `json:"user"`
	BillingPlans *model.BillingPlanConnection `json:"billingPlans"`
}
