export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  /**
   * This schema is only used to auto genrate the struct for the communication between the core and
   * the apps.
   */
  BloomMetadata: any;
  UserConnection: any;
  Group: any;
  BillingPlanConnection: any;
  User: any;
  Time: any;
  Bytes: any;
};


export type DashboardData = {
  __typename?: 'DashboardData';
  metadata?: Maybe<Scalars['BloomMetadata']>;
  users?: Maybe<Scalars['UserConnection']>;
};

export type FetchGroupProfileParams = {
  __typename?: 'FetchGroupProfileParams';
  id: Scalars['ID'];
};

export type GroupBillingProfile = {
  __typename?: 'GroupBillingProfile';
  group?: Maybe<Scalars['Group']>;
  billingPlans?: Maybe<Scalars['BillingPlanConnection']>;
  stripePublicKey?: Maybe<Scalars['String']>;
};

export type UserBillingProfile = {
  __typename?: 'UserBillingProfile';
  user?: Maybe<Scalars['User']>;
  billingPlans?: Maybe<Scalars['BillingPlanConnection']>;
};

export type NewStripeCard = {
  __typename?: 'NewStripeCard';
  number: Scalars['String'];
  expMonth: Scalars['String'];
  expYear: Scalars['String'];
  cvc: Scalars['String'];
};

export type AddPaymentMethodParams = {
  __typename?: 'AddPaymentMethodParams';
  stripePublicKey?: Maybe<Scalars['String']>;
  groupID?: Maybe<Scalars['ID']>;
  card: NewStripeCard;
};

export type MyBillingProfile = {
  __typename?: 'MyBillingProfile';
  me?: Maybe<Scalars['User']>;
  billingPlans?: Maybe<Scalars['BillingPlanConnection']>;
  stripePublicKey: Scalars['String'];
};

export type CalcParams = {
  __typename?: 'CalcParams';
  expression: Scalars['String'];
};

export type CalcResult = {
  __typename?: 'CalcResult';
  result: Scalars['String'];
};

export type CalendarFindEventsParams = {
  __typename?: 'CalendarFindEventsParams';
  startAt?: Maybe<Scalars['Time']>;
  endAt?: Maybe<Scalars['Time']>;
  groupID?: Maybe<Scalars['ID']>;
};

export type CalendarCreateEventParams = {
  __typename?: 'CalendarCreateEventParams';
  title: Scalars['String'];
  location: Scalars['String'];
  description: Scalars['String'];
  startAt: Scalars['Time'];
  endAt: Scalars['Time'];
  groupID?: Maybe<Scalars['ID']>;
};

export type CalendarDeleteEventParams = {
  __typename?: 'CalendarDeleteEventParams';
  id: Scalars['Bytes'];
};

/**
 * type Contacts struct {
 *   Contacts []Contact `json:"contacts"`
 * }
 * type CreateContactParams struct {
 *   DeviceID      string        `json:"deviceId"`
 *   FirstName     string        `json:"firstName"`
 *   LastName      string        `json:"lastName"`
 *   Notes         string        `json:"notes"`
 *   Birthday      *time.Time    `json:"birthday"`
 *   BloomUsername string        `json:"bloomUsername" db:"bloom_username"`
 *   Organizations Organizations `json:"organizations"`
 *   Addresses     Addresses     `json:"addresses"`
 *   Emails        Emails        `json:"emails"`
 *   Phones        Phones        `json:"phones"`
 *   Websites      Websites      `json:"websites"`
 * }
 */
export type DeleteContactParams = {
  __typename?: 'DeleteContactParams';
  id: Scalars['Bytes'];
};

export type ContactsFindParams = {
  __typename?: 'ContactsFindParams';
  groupID?: Maybe<Scalars['ID']>;
};

export type GroupsCreateParams = {
  __typename?: 'GroupsCreateParams';
  name: Scalars['String'];
  description: Scalars['String'];
};

export type GroupsDeleteParams = {
  __typename?: 'GroupsDeleteParams';
  groupID: Scalars['ID'];
};

export type GroupsFetchMembersParams = {
  __typename?: 'GroupsFetchMembersParams';
  groupID: Scalars['ID'];
};

export type GroupsFetchDetailsParams = {
  __typename?: 'GroupsFetchDetailsParams';
  groupID: Scalars['ID'];
};

export type GroupsInviteUserParams = {
  __typename?: 'GroupsInviteUserParams';
  groupID: Scalars['ID'];
  username: Scalars['String'];
};

export type GroupsCancelInvitationParams = {
  __typename?: 'GroupsCancelInvitationParams';
  invitationID: Scalars['ID'];
};

export type GroupsDeclineInvitationParams = {
  __typename?: 'GroupsDeclineInvitationParams';
  invitationID: Scalars['ID'];
};

export type GroupsRemoveMembersParams = {
  __typename?: 'GroupsRemoveMembersParams';
  groupID: Scalars['ID'];
  username: Scalars['String'];
};

export type GroupsQuitParams = {
  __typename?: 'GroupsQuitParams';
  groupID: Scalars['ID'];
};

export type Empty = {
  __typename?: 'Empty';
  empty?: Maybe<Scalars['Boolean']>;
};

export type NotesCreateParams = {
  __typename?: 'NotesCreateParams';
  title: Scalars['String'];
  body: Scalars['String'];
  color: Scalars['String'];
  groupID?: Maybe<Scalars['ID']>;
};

export type NotesDeleteParams = {
  __typename?: 'NotesDeleteParams';
  id: Scalars['Bytes'];
};

export type NotesFindParams = {
  __typename?: 'NotesFindParams';
  groupID?: Maybe<Scalars['ID']>;
};

export type PreferencesSetParams = {
  __typename?: 'PreferencesSetParams';
  key: Scalars['String'];
  value: Scalars['String'];
};

export type PreferencesGetParams = {
  __typename?: 'PreferencesGetParams';
  key: Scalars['String'];
};

export type PreferencesDeleteParams = {
  __typename?: 'PreferencesDeleteParams';
  key: Scalars['String'];
};
