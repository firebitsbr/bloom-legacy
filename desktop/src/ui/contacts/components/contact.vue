<template>
  <div>
    <v-toolbar elevation="0" v-if="contact">
      <v-toolbar-title>{{ contact.data.firstName }} {{ contact.data.lastName }}</v-toolbar-title>

      <v-spacer />
      <v-menu>
        <template v-slot:activator="{ on }">
          <v-btn icon v-on="on">
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>
        <v-list class="text-left">
          <v-list-item @click="deleteContact">
            <v-list-item-icon>
              <v-icon>mdi-delete</v-icon>
            </v-list-item-icon>
            <v-list-item-title>Delete forever</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-toolbar>

    <div style="height: calc(100vh - 65px);" class="overflow-y-auto">
      <v-container class="ps-3">
        <v-row>
          <v-col cols="6">
            <v-text-field
              label="First name"
              v-model="contact.data.firstName"
              prepend-icon="mdi-account"
            />
          </v-col>
          <v-col cols="6">
            <v-text-field
              label="Last name"
              v-model="contact.data.lastName"
            />
          </v-col>
          <v-col cols="12">
            <v-text-field
              label="Bloom username"
              v-model="contact.data.bloomUsername"
              prefix="@"
            />
          </v-col>
        </v-row>


        <!-- birthday -->
          <v-row>
            <v-col cols="12">
              <v-menu
                ref="birthdayMenu"
                :close-on-content-click="false"
                v-model="birthdayMenu"
                :nudge-right="40"
                transition="scale-transition"
                offset-y
                min-width="290px"
              >
                <template v-slot:activator="{ on }">
                  <v-text-field
                    v-on="on"
                    v-model="birthdayFormatted"
                    label="Birthday"
                    prepend-icon="mdi-calendar"
                    @blur="birthday = parseDate(birthdayFormatted)"
                    readonly
                    clearable
                    @click:clear="birthdayCleared"
                  />
                </template>
                <v-date-picker
                  ref="birthdayPicker"
                  v-model="birthday"
                  :max="new Date().toISOString().substr(0, 10)"
                  min="1900-01-01"
                  @change="saveBirthday"
                ></v-date-picker>
              </v-menu>
            </v-col>
          </v-row>
          <!-- /birthday -->

          <!-- addresses: TODO(z0mbie42) -->

          <!-- organizations -->
          <!-- <v-row>
            <v-col cols="6">
              <v-text-field
                label="Company"
                v-model="contact.data.organizations[0].name"
                prepend-icon="mdi-domain"
              ></v-text-field>
            </v-col>
            <v-col cols="6">
              <v-text-field
                label="Job title"
                v-model="contact.data.organizations[0].title"
              ></v-text-field>
            </v-col>
          </v-row> -->
          <!-- /organizations -->


          <!-- phones -->
          <v-row>
            <v-col cols="12">
              <v-list class="pt-0 pb-0">
                <v-list-item
                  v-for="(phone, index) in contact.data.phones"
                  :key="index"
                  class="contacts-add-row"
                >
                  <v-row align="center">
                    <v-col cols="1" class="pl-0 pr-0">
                      <v-icon v-if="index === 0">mdi-phone</v-icon>
                    </v-col>
                    <v-col cols="5">
                      <v-text-field
                        label="Phone"
                        v-model="phone.value"
                      ></v-text-field>
                    </v-col>
                    <v-col cols="4">
                      <v-autocomplete
                        :items="phoneLabels"
                        v-model="phone.label"
                        label="Label"
                        single-line
                      >
                      </v-autocomplete>
                    </v-col>
                    <v-col cols="1" class="pr-0 pl-0">
                      <v-tooltip bottom>
                        <template v-slot:activator="{ on }">
                          <v-btn
                            text
                            icon
                            small
                            v-on="on"
                            color="grey"
                            class="action-button"
                            @click="removePhone(index)"
                          >
                            <v-icon>mdi-close</v-icon>
                          </v-btn>
                        </template>
                        <span>Remove</span>
                      </v-tooltip>
                    </v-col>
                    <v-col cols="1" class="pl-0 pr-0"
                      v-if="index === contact.data.phones.length - 1">
                      <v-tooltip bottom>
                        <template v-slot:activator="{ on }">
                          <v-btn
                            text
                            icon
                            small
                            v-on="on"
                            color="blue darken-2"
                            class="ml-0"
                            @click="addPhone"
                          >
                            <v-icon>mdi-plus-circle</v-icon>
                          </v-btn>
                        </template>
                        <span>Add</span>
                      </v-tooltip>
                    </v-col>
                  </v-row>
                </v-list-item>
              </v-list>
            </v-col>
          </v-row>
          <!-- /phones -->


          <!-- emails -->
          <v-row>
            <v-col cols="12">
              <v-list class="pt-0 pb-0">
                <v-list-item
                  v-for="(email, index) in contact.data.emails"
                  :key="index"
                  class="contacts-add-row"
                >
                  <v-row align="center">
                    <v-col
                      cols="1"
                      class="pl-0 pr-0"
                    >
                      <v-icon v-if="index === 0">mdi-email-outline</v-icon>
                    </v-col>
                    <v-col cols="5">
                      <v-text-field
                        label="Email"
                        v-model="email.value"
                      ></v-text-field>
                    </v-col>
                    <v-col cols="4">
                      <v-autocomplete
                        :items="emailLabels"
                        v-model="email.label"
                        label="Label"
                        single-line
                      >
                      </v-autocomplete>
                    </v-col>
                    <v-col
                      cols="1"
                      class="pr-0 pl-0"
                    >
                      <v-tooltip bottom>
                        <template v-slot:activator="{ on }">
                          <v-btn
                            text
                            icon
                            small
                            v-on="on"
                            color="grey"
                            class="action-button"
                            @click="removeEmail(index)"
                          >
                            <v-icon>mdi-close</v-icon>
                          </v-btn>
                        </template>
                        <span>Remove</span>
                      </v-tooltip>
                    </v-col>
                    <v-col
                      cols="1"
                      class="pl-0 pr-0"
                    >
                      <v-tooltip
                        bottom
                        v-if="index === contact.data.emails.length - 1"
                      >
                        <template v-slot:activator="{ on }">
                          <v-btn
                            text
                            icon
                            small
                            v-on="on"
                            color="blue darken-2"
                            class="ml-0"
                            @click="addEmail"
                          >
                            <v-icon>mdi-plus-circle</v-icon>
                          </v-btn>
                        </template>
                        <span>Add</span>
                      </v-tooltip>
                    </v-col>
                  </v-row>
                </v-list-item>
              </v-list>
            </v-col>
          </v-row>
          <!-- /emails -->

          <!-- websites -->
          <v-row>
            <v-col cols="12">
              <v-list class="pt-0 pb-0">
                <v-list-item
                  v-for="(website, index) in contact.data.online"
                  :key="index"
                  class="contacts-add-row"
                >
                  <v-row align="center">
                    <v-col
                      cols="1"
                      class="pl-0 pr-0"
                    >
                      <v-icon v-if="index === 0">mdi-earth</v-icon>
                    </v-col>
                    <v-col cols="5">
                      <v-text-field
                        label="Website"
                        v-model="website.value"
                      ></v-text-field>
                    </v-col>
                    <v-col cols="4">
                      <v-autocomplete
                        :items="websiteLabels"
                        v-model="website.label"
                        label="Label"
                        single-line
                      >
                      </v-autocomplete>
                    </v-col>
                    <v-col
                      cols="1"
                      class="pr-0 pl-0"
                    >
                      <v-tooltip bottom>
                        <template v-slot:activator="{ on }">
                          <v-btn
                            text
                            icon
                            small
                            v-on="on"
                            color="grey"
                            class="action-button"
                            @click="removeWebsite(index)"
                          >
                            <v-icon>mdi-close</v-icon>
                          </v-btn>
                        </template>
                        <span>Remove</span>
                      </v-tooltip>
                    </v-col>
                    <v-col
                      cols="1"
                      class="pl-0 pr-0"
                    >
                      <v-tooltip
                        bottom
                        v-if="index === contact.data.online.length - 1"
                      >
                        <template v-slot:activator="{ on }">
                          <v-btn
                            text
                            icon
                            small
                            v-on="on"
                            color="blue darken-2"
                            class="ml-0"
                            @click="addWebsite"
                          >
                            <v-icon>mdi-plus-circle</v-icon>
                          </v-btn>
                        </template>
                        <span>Add</span>
                      </v-tooltip>
                    </v-col>
                  </v-row>
                </v-list-item>
              </v-list>
            </v-col>
          </v-row>
          <!-- /websites -->

          <!-- addresses -->
          <v-row>
            <v-col cols="12">
              <v-list class="pt-0 pb-0">
                <v-list-item v-for="(address, index) in contact.data.addresses" :key="index">
                  <v-row align="center">
                    <v-col cols="1" class="pl-0 pr-0">
                      <v-icon v-if="index === 0">mdi-home</v-icon>
                    </v-col>
                    <v-col cols="9">
                      <v-expansion-panels flat>
                        <v-expansion-panel>
                          <v-expansion-panel-header>Address</v-expansion-panel-header>
                          <v-expansion-panel-content>
                            <v-row>
                              <v-col cols="12">
                                <v-text-field
                                  label="Street"
                                  v-model="address.street"
                                ></v-text-field>
                              </v-col>
                              <v-col cols="6">
                                <v-text-field
                                  label="City"
                                  v-model="address.city"
                                ></v-text-field>
                              </v-col>
                              <v-col cols="6">
                                <v-text-field
                                  label="Postal code"
                                  v-model="address.postalCode"
                                ></v-text-field>
                              </v-col>
                              <v-col cols="6">
                                <v-text-field
                                  label="State"
                                  v-model="address.state"
                                ></v-text-field>
                              </v-col>
                              <v-col cols="6">
                                <v-text-field
                                  label="Country"
                                  v-model="address.country"
                                ></v-text-field>
                              </v-col>
                              <v-col cols="12">
                                <v-autocomplete
                                  :items="addressLabels"
                                  v-model="address.label"
                                  label="Label"
                                  single-line
                                >
                                </v-autocomplete>
                              </v-col>
                            </v-row>
                          </v-expansion-panel-content>
                        </v-expansion-panel>
                      </v-expansion-panels>
                    </v-col>
                    <v-col cols="1" class="pr-0 pl-0">
                      <v-tooltip bottom>
                        <template v-slot:activator="{ on }">
                          <v-btn
                            text
                            icon
                            small
                            v-on="on"
                            color="grey"
                            class="action-button"
                            @click="removeAddress(index)"
                          >
                            <v-icon>mdi-close</v-icon>
                          </v-btn>
                        </template>
                        <span>Remove</span>
                      </v-tooltip>
                    </v-col>
                    <v-col cols="1" class="pl-0 pr-0"
                      v-if="index === contact.data.addresses.length - 1">
                      <v-tooltip bottom>
                        <template v-slot:activator="{ on }">
                          <v-btn
                            text
                            icon
                            small
                            v-on="on"
                            color="blue darken-2"
                            class="ml-0"
                            @click="addAddress"
                          >
                            <v-icon>mdi-plus-circle</v-icon>
                          </v-btn>
                        </template>
                        <span>Add</span>
                      </v-tooltip>
                    </v-col>
                  </v-row>
                </v-list-item>
              </v-list>
            </v-col>
          </v-row>
          <!-- addresses -->

          <!-- notes -->
          <v-row>
            <v-col cols="12">
              <v-textarea
                label="Notes"
                v-model="contact.data.notes"
                prepend-icon="mdi-calendar-text"
              ></v-textarea>
            </v-col>
          </v-row>


      </v-container>
    </div>
  </div>
</template>


<script lang="ts">
import {
  Component,
  Vue,
  Prop,
  Watch,
} from 'vue-property-decorator';
import {
  Contact,
  // Organization,
  // Website,
  // Email,
  // Phone,
  DeleteContact,
  Method,
  // Address,
} from '@/core/contacts';
import core, { BlmObject } from '@/core';

const DEFAULT_EMAIL = { value: '', label: 'Personal' };
const DEFAULT_WEBSITE = { value: '', label: 'Personal' };
const DEFAULT_PHONE = { value: '', label: 'Mobile' };
// const DEFAULT_ORGANIZATION = { name: '', title: '' };
const DEFAULT_ADDRESS = {
  city: '',
  country: '',
  label: 'Home',
  postalCode: '',
  street: '',
  state: '',
};

@Component
export default class BlmContact extends Vue {
  // props
  @Prop({ type: Object, required: true }) contact!: BlmObject<Contact>;

  // data
  error = '';
  loading = false;
  // firstName: string = '';
  // lastName: string = '';
  birthday: string | null = null;
  birthdayFormatted: string | null = null;
  birthdayMenu = false;
  // notes: string = '';
  // bloomUsername: string = '';
  // organizations: Organization[] = [{ ...DEFAULT_ORGANIZATION }];
  // phones: Phone[] = [{ ...DEFAULT_PHONE }];
  // emails: Email[] = [{ ...DEFAULT_EMAIL }];
  // websites: Website[] = [{ ...DEFAULT_WEBSITE }];
  // addresses: Address[] = [{ ...DEFAULT_ADDRESS }];
  phoneLabels = [
    'Home',
    'Work',
    'Other',
    'Mobile',
    'Main',
    'Home fax',
    'Work fax',
  ];
  addressLabels = [
    'Home',
    'Work',
    'Other',
  ];
  websiteLabels = [
    'Personal',
    'Blog',
    'Home page',
    'Work',
    'Other',
  ];
  emailLabels = [
    'Personal',
    'Work',
    'Other',
  ];

  // computed
  // lifecycle
  // watch
  @Watch('birthday')
  onBirthdayUpdated() {
    this.birthdayFormatted = this.formatDate(this.birthday);
  }

  @Watch('birthdayMenu')
  onBirthdayMenuChanged(birthdayMenu: boolean) {
    if (birthdayMenu) {
      // eslint-disable-next-line
      this.$nextTick(() => ((this.$refs.birthdayPicker) as any).activePicker = 'YEAR');
    }
  }


  @Watch('contact')
  onContactChanged(contact: BlmObject<Contact>) {
    // this.firstName = contact.data.firstName;
    // this.lastName = contact.data.lastName;
    // this.notes = contact.data.notes;
    // this.bloomUsername = contact.data.bloomUsername;
    this.birthday = contact.data.birthday
      ? new Date(contact.data.birthday!).toISOString().substr(0, 10) : null;
    this.birthdayFormatted = this.formatDate(this.birthday);
    this.contact.data.emails = this.contact.data.emails.length > 0
      ? this.contact.data.emails
      : [{ ...DEFAULT_EMAIL }];
    this.contact.data.online = contact.data.online.length > 0
      ? this.contact.data.online
      : [{ ...DEFAULT_WEBSITE }];
    this.contact.data.phones = contact.data.phones.length > 0
      ? this.contact.data.phones
      : [{ ...DEFAULT_PHONE }];
    // this.contact.data.organizations = contact.data.organizations.length > 0
    //   ? this.contact.data.organizations
    //   : [{ ...DEFAULT_ORGANIZATION }];
    this.contact.data.addresses = contact.data.addresses.length > 0
      ? this.contact.data.addresses
      : [{ ...DEFAULT_ADDRESS }];
  }


  // methods
  async deleteContact() {
    this.error = '';
    this.loading = true;
    const params: DeleteContact = {
      id: this.contact!.id,
    };
    try {
      await core.call(Method.DeleteContact, params);
      this.$emit('deleted', this.contact);
      // this.close();
    } catch (err) {
      this.error = err.message;
    } finally {
      this.loading = false;
    }
  }

  addPhone() {
    this.contact.data.phones.push({ ...DEFAULT_PHONE });
  }

  removePhone(index: number) {
    this.contact.data.phones.splice(index, 1);
    if (this.contact.data.phones.length === 0) {
      this.addPhone();
    }
  }

  addEmail() {
    this.contact.data.emails.push({ ...DEFAULT_EMAIL });
  }
  removeEmail(index: number) {
    this.contact.data.emails.splice(index, 1);
    if (this.contact.data.emails.length === 0) {
      this.addEmail();
    }
  }

  addWebsite() {
    this.contact.data.online.push({ ...DEFAULT_WEBSITE });
  }
  removeWebsite(index: number) {
    this.contact.data.online.splice(index, 1);
    if (this.contact.data.online.length === 0) {
      this.addWebsite();
    }
  }

  addAddress() {
    this.contact.data.addresses.push({ ...DEFAULT_ADDRESS });
  }
  removeAddress(index: number) {
    this.contact.data.addresses.splice(index, 1);
    if (this.contact.data.addresses.length === 0) {
      this.addAddress();
    }
  }

  saveBirthday(date: any) {
    (this.$refs.birthdayMenu as any).save(date);
  }

  formatDate(date: any) {
    if (!date) {
      return null;
    }

    const [year, month, day] = date.split('-');
    return `${year}/${month}/${day}`;
  }

  parseDate(date: any) {
    if (!date) {
      return null;
    }

    const [year, month, day] = date.split('/');
    return `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`;
  }

  birthdayCleared() {
    this.birthday = null;
    this.birthdayFormatted = null;
  }
}
</script>


<style lang="scss" scoped>
.v-toolbar {
  border-bottom: 1px solid rgba($color: #000000, $alpha: 0.1) !important;
  left: 0px !important;
}
</style>
