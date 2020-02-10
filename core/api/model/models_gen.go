// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AcceptGroupInvitationInput struct {
	ID string `json:"id"`
}

type AddPaymentMethodInput struct {
	StripeID string  `json:"stripeId"`
	GroupID  *string `json:"groupId"`
}

type BillingPlan struct {
	ID             string          `json:"id"`
	Price          float64         `json:"price"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	IsActive       bool            `json:"isActive"`
	Tier           BillingPlanTier `json:"tier"`
	AllowedStorage Int64           `json:"allowedStorage"`
}

type BillingPlanInput struct {
	ID          *string         `json:"id"`
	Name        string          `json:"name"`
	Price       float64         `json:"price"`
	Tier        BillingPlanTier `json:"tier"`
	StripeID    string          `json:"stripeId"`
	Description string          `json:"description"`
	IsActive    *bool           `json:"isActive"`
}

type CancelGroupInvitationInput struct {
	ID string `json:"id"`
}

type ChangeBillingPlanInput struct {
	ID      string  `json:"id"`
	UserID  *string `json:"userId"`
	GroupID *string `json:"groupId"`
}

type ChangeDefaultPaymentMethodInput struct {
	ID string `json:"id"`
}

type CompleteRegistrationInput struct {
	ID       string              `json:"id"`
	Username string              `json:"username"`
	AuthKey  Bytes               `json:"authKey"`
	Device   *SessionDeviceInput `json:"device"`
}

type CreateGroupInput struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	UsersToInvite []string `json:"usersToInvite"`
}

type DeclineGroupInvitationInput struct {
	ID string `json:"id"`
}

type DeleteBillingPlanInput struct {
	ID string `json:"id"`
}

type DeleteGroupInput struct {
	ID string `json:"id"`
}

type GroupInput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GroupInvitation struct {
	ID      string `json:"id"`
	Group   *Group `json:"group"`
	Inviter *User  `json:"inviter"`
}

type GroupMember struct {
	User *User           `json:"user"`
	Role GroupMemberRole `json:"role"`
}

type InviteUsersInGroupInput struct {
	ID        string   `json:"id"`
	Usernames []string `json:"usernames"`
}

type Invoice struct {
	ID string `json:"id"`
}

type PaymentMethod struct {
	ID                  string    `json:"id"`
	CreatedAt           time.Time `json:"createdAt"`
	CardLast4           string    `json:"cardLast4"`
	CardExpirationMonth int       `json:"cardExpirationMonth"`
	CardExpirationYear  int       `json:"cardExpirationYear"`
}

type QuitGroupInput struct {
	ID string `json:"id"`
}

type RegisterInput struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type RegistrationStarted struct {
	ID string `json:"id"`
}

type RemoveGroupMembersInput struct {
	ID        string   `json:"id"`
	Usernames []string `json:"usernames"`
}

type RemovePaymentMethodInput struct {
	ID string `json:"id"`
}

type RevokeSessionInput struct {
	ID string `json:"id"`
}

type SendNewRegistrationCodeInput struct {
	ID string `json:"id"`
}

type ServerVersion struct {
	Os        string `json:"os"`
	Arch      string `json:"arch"`
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
}

type Session struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	Token     *string        `json:"token"`
	Device    *SessionDevice `json:"device"`
}

type SessionDevice struct {
	Os   SessionDeviceOs   `json:"os"`
	Type SessionDeviceType `json:"type"`
}

type SessionDeviceInput struct {
	Os   SessionDeviceOs   `json:"os"`
	Type SessionDeviceType `json:"type"`
}

type SignInInput struct {
	Username string              `json:"username"`
	AuthKey  Bytes               `json:"authKey"`
	Device   *SessionDeviceInput `json:"device"`
}

type SignedIn struct {
	Session *Session `json:"session"`
	Me      *User    `json:"me"`
}

type UpdateUserProfileInput struct {
	ID          *string `json:"id"`
	DisplayName *string `json:"displayName"`
	Bio         *string `json:"bio"`
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
}

type VerifyRegistrationInput struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

type BillingPlanTier string

const (
	BillingPlanTierFree  BillingPlanTier = "FREE"
	BillingPlanTierBasic BillingPlanTier = "BASIC"
	BillingPlanTierPro   BillingPlanTier = "PRO"
	BillingPlanTierUltra BillingPlanTier = "ULTRA"
)

var AllBillingPlanTier = []BillingPlanTier{
	BillingPlanTierFree,
	BillingPlanTierBasic,
	BillingPlanTierPro,
	BillingPlanTierUltra,
}

func (e BillingPlanTier) IsValid() bool {
	switch e {
	case BillingPlanTierFree, BillingPlanTierBasic, BillingPlanTierPro, BillingPlanTierUltra:
		return true
	}
	return false
}

func (e BillingPlanTier) String() string {
	return string(e)
}

func (e *BillingPlanTier) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BillingPlanTier(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BillingPlanTier", str)
	}
	return nil
}

func (e BillingPlanTier) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GroupMemberRole string

const (
	GroupMemberRoleAdmin  GroupMemberRole = "ADMIN"
	GroupMemberRoleMember GroupMemberRole = "MEMBER"
)

var AllGroupMemberRole = []GroupMemberRole{
	GroupMemberRoleAdmin,
	GroupMemberRoleMember,
}

func (e GroupMemberRole) IsValid() bool {
	switch e {
	case GroupMemberRoleAdmin, GroupMemberRoleMember:
		return true
	}
	return false
}

func (e GroupMemberRole) String() string {
	return string(e)
}

func (e *GroupMemberRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GroupMemberRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GroupMemberRole", str)
	}
	return nil
}

func (e GroupMemberRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionDeviceOs string

const (
	SessionDeviceOsLinux   SessionDeviceOs = "LINUX"
	SessionDeviceOsMacos   SessionDeviceOs = "MACOS"
	SessionDeviceOsWindows SessionDeviceOs = "WINDOWS"
	SessionDeviceOsAndroid SessionDeviceOs = "ANDROID"
	SessionDeviceOsIos     SessionDeviceOs = "IOS"
)

var AllSessionDeviceOs = []SessionDeviceOs{
	SessionDeviceOsLinux,
	SessionDeviceOsMacos,
	SessionDeviceOsWindows,
	SessionDeviceOsAndroid,
	SessionDeviceOsIos,
}

func (e SessionDeviceOs) IsValid() bool {
	switch e {
	case SessionDeviceOsLinux, SessionDeviceOsMacos, SessionDeviceOsWindows, SessionDeviceOsAndroid, SessionDeviceOsIos:
		return true
	}
	return false
}

func (e SessionDeviceOs) String() string {
	return string(e)
}

func (e *SessionDeviceOs) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionDeviceOs(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionDeviceOS", str)
	}
	return nil
}

func (e SessionDeviceOs) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionDeviceType string

const (
	SessionDeviceTypeTv       SessionDeviceType = "TV"
	SessionDeviceTypeConsole  SessionDeviceType = "CONSOLE"
	SessionDeviceTypeMobile   SessionDeviceType = "MOBILE"
	SessionDeviceTypeTablet   SessionDeviceType = "TABLET"
	SessionDeviceTypeWatch    SessionDeviceType = "WATCH"
	SessionDeviceTypeComputer SessionDeviceType = "COMPUTER"
	SessionDeviceTypeCar      SessionDeviceType = "CAR"
)

var AllSessionDeviceType = []SessionDeviceType{
	SessionDeviceTypeTv,
	SessionDeviceTypeConsole,
	SessionDeviceTypeMobile,
	SessionDeviceTypeTablet,
	SessionDeviceTypeWatch,
	SessionDeviceTypeComputer,
	SessionDeviceTypeCar,
}

func (e SessionDeviceType) IsValid() bool {
	switch e {
	case SessionDeviceTypeTv, SessionDeviceTypeConsole, SessionDeviceTypeMobile, SessionDeviceTypeTablet, SessionDeviceTypeWatch, SessionDeviceTypeComputer, SessionDeviceTypeCar:
		return true
	}
	return false
}

func (e SessionDeviceType) String() string {
	return string(e)
}

func (e *SessionDeviceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionDeviceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionDeviceType", str)
	}
	return nil
}

func (e SessionDeviceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
