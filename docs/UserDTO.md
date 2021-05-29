# UserDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountDisabled** | Pointer to **bool** |  | [optional] 
**AccountExpirationDate** | Pointer to **time.Time** |  | [optional] 
**AccountExpires** | Pointer to **bool** |  | [optional] 
**AllowUserToChangePassword** | Pointer to **bool** |  | [optional] 
**AutomaticallyGeneratePassword** | Pointer to **bool** |  | [optional] 
**BusinessCategory** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**CountryName** | Pointer to **string** |  | [optional] 
**DaysBeforeExpiration** | Pointer to **int32** |  | [optional] 
**DaysBetweenChanges** | Pointer to **int32** |  | [optional] 
**DistinguishedName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailNewPasword** | Pointer to **bool** |  | [optional] 
**ExtraAttributes** | Pointer to [**[]AttributeValueDTO**](AttributeValueDTO.md) |  | [optional] 
**FacsimilTelephoneNumber** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**ForcePeriodicPasswordChanges** | Pointer to **bool** |  | [optional] 
**GenerationQualifier** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]GroupDTO**](GroupDTO.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Initials** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**LimitSimultaneousLogin** | Pointer to **bool** |  | [optional] 
**LocalityName** | Pointer to **string** |  | [optional] 
**MaximunLogins** | Pointer to **int32** |  | [optional] 
**NotifyPasswordExpiration** | Pointer to **bool** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**OrganizationUnitName** | Pointer to **string** |  | [optional] 
**PasswordExpirationDate** | Pointer to **time.Time** |  | [optional] 
**PersonalTitle** | Pointer to **string** |  | [optional] 
**PostOfficeBox** | Pointer to **string** |  | [optional] 
**PostalAddress** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**PreventNewSession** | Pointer to **bool** |  | [optional] 
**StateOrProvinceName** | Pointer to **string** |  | [optional] 
**StreetAddress** | Pointer to **string** |  | [optional] 
**Surename** | Pointer to **string** |  | [optional] 
**TelephoneNumber** | Pointer to **string** |  | [optional] 
**TerminatePreviousSession** | Pointer to **bool** |  | [optional] 
**UserCertificate** | Pointer to **[]string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewUserDTO

`func NewUserDTO() *UserDTO`

NewUserDTO instantiates a new UserDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDTOWithDefaults

`func NewUserDTOWithDefaults() *UserDTO`

NewUserDTOWithDefaults instantiates a new UserDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountDisabled

`func (o *UserDTO) GetAccountDisabled() bool`

GetAccountDisabled returns the AccountDisabled field if non-nil, zero value otherwise.

### GetAccountDisabledOk

`func (o *UserDTO) GetAccountDisabledOk() (*bool, bool)`

GetAccountDisabledOk returns a tuple with the AccountDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDisabled

`func (o *UserDTO) SetAccountDisabled(v bool)`

SetAccountDisabled sets AccountDisabled field to given value.

### HasAccountDisabled

`func (o *UserDTO) HasAccountDisabled() bool`

HasAccountDisabled returns a boolean if a field has been set.

### GetAccountExpirationDate

`func (o *UserDTO) GetAccountExpirationDate() time.Time`

GetAccountExpirationDate returns the AccountExpirationDate field if non-nil, zero value otherwise.

### GetAccountExpirationDateOk

`func (o *UserDTO) GetAccountExpirationDateOk() (*time.Time, bool)`

GetAccountExpirationDateOk returns a tuple with the AccountExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationDate

`func (o *UserDTO) SetAccountExpirationDate(v time.Time)`

SetAccountExpirationDate sets AccountExpirationDate field to given value.

### HasAccountExpirationDate

`func (o *UserDTO) HasAccountExpirationDate() bool`

HasAccountExpirationDate returns a boolean if a field has been set.

### GetAccountExpires

`func (o *UserDTO) GetAccountExpires() bool`

GetAccountExpires returns the AccountExpires field if non-nil, zero value otherwise.

### GetAccountExpiresOk

`func (o *UserDTO) GetAccountExpiresOk() (*bool, bool)`

GetAccountExpiresOk returns a tuple with the AccountExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpires

`func (o *UserDTO) SetAccountExpires(v bool)`

SetAccountExpires sets AccountExpires field to given value.

### HasAccountExpires

`func (o *UserDTO) HasAccountExpires() bool`

HasAccountExpires returns a boolean if a field has been set.

### GetAllowUserToChangePassword

`func (o *UserDTO) GetAllowUserToChangePassword() bool`

GetAllowUserToChangePassword returns the AllowUserToChangePassword field if non-nil, zero value otherwise.

### GetAllowUserToChangePasswordOk

`func (o *UserDTO) GetAllowUserToChangePasswordOk() (*bool, bool)`

GetAllowUserToChangePasswordOk returns a tuple with the AllowUserToChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserToChangePassword

`func (o *UserDTO) SetAllowUserToChangePassword(v bool)`

SetAllowUserToChangePassword sets AllowUserToChangePassword field to given value.

### HasAllowUserToChangePassword

`func (o *UserDTO) HasAllowUserToChangePassword() bool`

HasAllowUserToChangePassword returns a boolean if a field has been set.

### GetAutomaticallyGeneratePassword

`func (o *UserDTO) GetAutomaticallyGeneratePassword() bool`

GetAutomaticallyGeneratePassword returns the AutomaticallyGeneratePassword field if non-nil, zero value otherwise.

### GetAutomaticallyGeneratePasswordOk

`func (o *UserDTO) GetAutomaticallyGeneratePasswordOk() (*bool, bool)`

GetAutomaticallyGeneratePasswordOk returns a tuple with the AutomaticallyGeneratePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyGeneratePassword

`func (o *UserDTO) SetAutomaticallyGeneratePassword(v bool)`

SetAutomaticallyGeneratePassword sets AutomaticallyGeneratePassword field to given value.

### HasAutomaticallyGeneratePassword

`func (o *UserDTO) HasAutomaticallyGeneratePassword() bool`

HasAutomaticallyGeneratePassword returns a boolean if a field has been set.

### GetBusinessCategory

`func (o *UserDTO) GetBusinessCategory() string`

GetBusinessCategory returns the BusinessCategory field if non-nil, zero value otherwise.

### GetBusinessCategoryOk

`func (o *UserDTO) GetBusinessCategoryOk() (*string, bool)`

GetBusinessCategoryOk returns a tuple with the BusinessCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCategory

`func (o *UserDTO) SetBusinessCategory(v string)`

SetBusinessCategory sets BusinessCategory field to given value.

### HasBusinessCategory

`func (o *UserDTO) HasBusinessCategory() bool`

HasBusinessCategory returns a boolean if a field has been set.

### GetCommonName

`func (o *UserDTO) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *UserDTO) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *UserDTO) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *UserDTO) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountryName

`func (o *UserDTO) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *UserDTO) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *UserDTO) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *UserDTO) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetDaysBeforeExpiration

`func (o *UserDTO) GetDaysBeforeExpiration() int32`

GetDaysBeforeExpiration returns the DaysBeforeExpiration field if non-nil, zero value otherwise.

### GetDaysBeforeExpirationOk

`func (o *UserDTO) GetDaysBeforeExpirationOk() (*int32, bool)`

GetDaysBeforeExpirationOk returns a tuple with the DaysBeforeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBeforeExpiration

`func (o *UserDTO) SetDaysBeforeExpiration(v int32)`

SetDaysBeforeExpiration sets DaysBeforeExpiration field to given value.

### HasDaysBeforeExpiration

`func (o *UserDTO) HasDaysBeforeExpiration() bool`

HasDaysBeforeExpiration returns a boolean if a field has been set.

### GetDaysBetweenChanges

`func (o *UserDTO) GetDaysBetweenChanges() int32`

GetDaysBetweenChanges returns the DaysBetweenChanges field if non-nil, zero value otherwise.

### GetDaysBetweenChangesOk

`func (o *UserDTO) GetDaysBetweenChangesOk() (*int32, bool)`

GetDaysBetweenChangesOk returns a tuple with the DaysBetweenChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBetweenChanges

`func (o *UserDTO) SetDaysBetweenChanges(v int32)`

SetDaysBetweenChanges sets DaysBetweenChanges field to given value.

### HasDaysBetweenChanges

`func (o *UserDTO) HasDaysBetweenChanges() bool`

HasDaysBetweenChanges returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *UserDTO) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *UserDTO) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *UserDTO) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *UserDTO) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetEmail

`func (o *UserDTO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDTO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDTO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserDTO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailNewPasword

`func (o *UserDTO) GetEmailNewPasword() bool`

GetEmailNewPasword returns the EmailNewPasword field if non-nil, zero value otherwise.

### GetEmailNewPaswordOk

`func (o *UserDTO) GetEmailNewPaswordOk() (*bool, bool)`

GetEmailNewPaswordOk returns a tuple with the EmailNewPasword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNewPasword

`func (o *UserDTO) SetEmailNewPasword(v bool)`

SetEmailNewPasword sets EmailNewPasword field to given value.

### HasEmailNewPasword

`func (o *UserDTO) HasEmailNewPasword() bool`

HasEmailNewPasword returns a boolean if a field has been set.

### GetExtraAttributes

`func (o *UserDTO) GetExtraAttributes() []AttributeValueDTO`

GetExtraAttributes returns the ExtraAttributes field if non-nil, zero value otherwise.

### GetExtraAttributesOk

`func (o *UserDTO) GetExtraAttributesOk() (*[]AttributeValueDTO, bool)`

GetExtraAttributesOk returns a tuple with the ExtraAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraAttributes

`func (o *UserDTO) SetExtraAttributes(v []AttributeValueDTO)`

SetExtraAttributes sets ExtraAttributes field to given value.

### HasExtraAttributes

`func (o *UserDTO) HasExtraAttributes() bool`

HasExtraAttributes returns a boolean if a field has been set.

### GetFacsimilTelephoneNumber

`func (o *UserDTO) GetFacsimilTelephoneNumber() string`

GetFacsimilTelephoneNumber returns the FacsimilTelephoneNumber field if non-nil, zero value otherwise.

### GetFacsimilTelephoneNumberOk

`func (o *UserDTO) GetFacsimilTelephoneNumberOk() (*string, bool)`

GetFacsimilTelephoneNumberOk returns a tuple with the FacsimilTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacsimilTelephoneNumber

`func (o *UserDTO) SetFacsimilTelephoneNumber(v string)`

SetFacsimilTelephoneNumber sets FacsimilTelephoneNumber field to given value.

### HasFacsimilTelephoneNumber

`func (o *UserDTO) HasFacsimilTelephoneNumber() bool`

HasFacsimilTelephoneNumber returns a boolean if a field has been set.

### GetFirstName

`func (o *UserDTO) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserDTO) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserDTO) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserDTO) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetForcePeriodicPasswordChanges

`func (o *UserDTO) GetForcePeriodicPasswordChanges() bool`

GetForcePeriodicPasswordChanges returns the ForcePeriodicPasswordChanges field if non-nil, zero value otherwise.

### GetForcePeriodicPasswordChangesOk

`func (o *UserDTO) GetForcePeriodicPasswordChangesOk() (*bool, bool)`

GetForcePeriodicPasswordChangesOk returns a tuple with the ForcePeriodicPasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePeriodicPasswordChanges

`func (o *UserDTO) SetForcePeriodicPasswordChanges(v bool)`

SetForcePeriodicPasswordChanges sets ForcePeriodicPasswordChanges field to given value.

### HasForcePeriodicPasswordChanges

`func (o *UserDTO) HasForcePeriodicPasswordChanges() bool`

HasForcePeriodicPasswordChanges returns a boolean if a field has been set.

### GetGenerationQualifier

`func (o *UserDTO) GetGenerationQualifier() string`

GetGenerationQualifier returns the GenerationQualifier field if non-nil, zero value otherwise.

### GetGenerationQualifierOk

`func (o *UserDTO) GetGenerationQualifierOk() (*string, bool)`

GetGenerationQualifierOk returns a tuple with the GenerationQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationQualifier

`func (o *UserDTO) SetGenerationQualifier(v string)`

SetGenerationQualifier sets GenerationQualifier field to given value.

### HasGenerationQualifier

`func (o *UserDTO) HasGenerationQualifier() bool`

HasGenerationQualifier returns a boolean if a field has been set.

### GetGivenName

`func (o *UserDTO) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *UserDTO) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *UserDTO) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *UserDTO) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetGroups

`func (o *UserDTO) GetGroups() []GroupDTO`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserDTO) GetGroupsOk() (*[]GroupDTO, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserDTO) SetGroups(v []GroupDTO)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserDTO) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *UserDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitials

`func (o *UserDTO) GetInitials() string`

GetInitials returns the Initials field if non-nil, zero value otherwise.

### GetInitialsOk

`func (o *UserDTO) GetInitialsOk() (*string, bool)`

GetInitialsOk returns a tuple with the Initials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitials

`func (o *UserDTO) SetInitials(v string)`

SetInitials sets Initials field to given value.

### HasInitials

`func (o *UserDTO) HasInitials() bool`

HasInitials returns a boolean if a field has been set.

### GetLanguage

`func (o *UserDTO) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserDTO) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserDTO) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserDTO) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLimitSimultaneousLogin

`func (o *UserDTO) GetLimitSimultaneousLogin() bool`

GetLimitSimultaneousLogin returns the LimitSimultaneousLogin field if non-nil, zero value otherwise.

### GetLimitSimultaneousLoginOk

`func (o *UserDTO) GetLimitSimultaneousLoginOk() (*bool, bool)`

GetLimitSimultaneousLoginOk returns a tuple with the LimitSimultaneousLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSimultaneousLogin

`func (o *UserDTO) SetLimitSimultaneousLogin(v bool)`

SetLimitSimultaneousLogin sets LimitSimultaneousLogin field to given value.

### HasLimitSimultaneousLogin

`func (o *UserDTO) HasLimitSimultaneousLogin() bool`

HasLimitSimultaneousLogin returns a boolean if a field has been set.

### GetLocalityName

`func (o *UserDTO) GetLocalityName() string`

GetLocalityName returns the LocalityName field if non-nil, zero value otherwise.

### GetLocalityNameOk

`func (o *UserDTO) GetLocalityNameOk() (*string, bool)`

GetLocalityNameOk returns a tuple with the LocalityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalityName

`func (o *UserDTO) SetLocalityName(v string)`

SetLocalityName sets LocalityName field to given value.

### HasLocalityName

`func (o *UserDTO) HasLocalityName() bool`

HasLocalityName returns a boolean if a field has been set.

### GetMaximunLogins

`func (o *UserDTO) GetMaximunLogins() int32`

GetMaximunLogins returns the MaximunLogins field if non-nil, zero value otherwise.

### GetMaximunLoginsOk

`func (o *UserDTO) GetMaximunLoginsOk() (*int32, bool)`

GetMaximunLoginsOk returns a tuple with the MaximunLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximunLogins

`func (o *UserDTO) SetMaximunLogins(v int32)`

SetMaximunLogins sets MaximunLogins field to given value.

### HasMaximunLogins

`func (o *UserDTO) HasMaximunLogins() bool`

HasMaximunLogins returns a boolean if a field has been set.

### GetNotifyPasswordExpiration

`func (o *UserDTO) GetNotifyPasswordExpiration() bool`

GetNotifyPasswordExpiration returns the NotifyPasswordExpiration field if non-nil, zero value otherwise.

### GetNotifyPasswordExpirationOk

`func (o *UserDTO) GetNotifyPasswordExpirationOk() (*bool, bool)`

GetNotifyPasswordExpirationOk returns a tuple with the NotifyPasswordExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPasswordExpiration

`func (o *UserDTO) SetNotifyPasswordExpiration(v bool)`

SetNotifyPasswordExpiration sets NotifyPasswordExpiration field to given value.

### HasNotifyPasswordExpiration

`func (o *UserDTO) HasNotifyPasswordExpiration() bool`

HasNotifyPasswordExpiration returns a boolean if a field has been set.

### GetOrganizationName

`func (o *UserDTO) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *UserDTO) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *UserDTO) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *UserDTO) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationUnitName

`func (o *UserDTO) GetOrganizationUnitName() string`

GetOrganizationUnitName returns the OrganizationUnitName field if non-nil, zero value otherwise.

### GetOrganizationUnitNameOk

`func (o *UserDTO) GetOrganizationUnitNameOk() (*string, bool)`

GetOrganizationUnitNameOk returns a tuple with the OrganizationUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnitName

`func (o *UserDTO) SetOrganizationUnitName(v string)`

SetOrganizationUnitName sets OrganizationUnitName field to given value.

### HasOrganizationUnitName

`func (o *UserDTO) HasOrganizationUnitName() bool`

HasOrganizationUnitName returns a boolean if a field has been set.

### GetPasswordExpirationDate

`func (o *UserDTO) GetPasswordExpirationDate() time.Time`

GetPasswordExpirationDate returns the PasswordExpirationDate field if non-nil, zero value otherwise.

### GetPasswordExpirationDateOk

`func (o *UserDTO) GetPasswordExpirationDateOk() (*time.Time, bool)`

GetPasswordExpirationDateOk returns a tuple with the PasswordExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationDate

`func (o *UserDTO) SetPasswordExpirationDate(v time.Time)`

SetPasswordExpirationDate sets PasswordExpirationDate field to given value.

### HasPasswordExpirationDate

`func (o *UserDTO) HasPasswordExpirationDate() bool`

HasPasswordExpirationDate returns a boolean if a field has been set.

### GetPersonalTitle

`func (o *UserDTO) GetPersonalTitle() string`

GetPersonalTitle returns the PersonalTitle field if non-nil, zero value otherwise.

### GetPersonalTitleOk

`func (o *UserDTO) GetPersonalTitleOk() (*string, bool)`

GetPersonalTitleOk returns a tuple with the PersonalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalTitle

`func (o *UserDTO) SetPersonalTitle(v string)`

SetPersonalTitle sets PersonalTitle field to given value.

### HasPersonalTitle

`func (o *UserDTO) HasPersonalTitle() bool`

HasPersonalTitle returns a boolean if a field has been set.

### GetPostOfficeBox

`func (o *UserDTO) GetPostOfficeBox() string`

GetPostOfficeBox returns the PostOfficeBox field if non-nil, zero value otherwise.

### GetPostOfficeBoxOk

`func (o *UserDTO) GetPostOfficeBoxOk() (*string, bool)`

GetPostOfficeBoxOk returns a tuple with the PostOfficeBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOfficeBox

`func (o *UserDTO) SetPostOfficeBox(v string)`

SetPostOfficeBox sets PostOfficeBox field to given value.

### HasPostOfficeBox

`func (o *UserDTO) HasPostOfficeBox() bool`

HasPostOfficeBox returns a boolean if a field has been set.

### GetPostalAddress

`func (o *UserDTO) GetPostalAddress() string`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *UserDTO) GetPostalAddressOk() (*string, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *UserDTO) SetPostalAddress(v string)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *UserDTO) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### GetPostalCode

`func (o *UserDTO) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UserDTO) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UserDTO) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UserDTO) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPreventNewSession

`func (o *UserDTO) GetPreventNewSession() bool`

GetPreventNewSession returns the PreventNewSession field if non-nil, zero value otherwise.

### GetPreventNewSessionOk

`func (o *UserDTO) GetPreventNewSessionOk() (*bool, bool)`

GetPreventNewSessionOk returns a tuple with the PreventNewSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventNewSession

`func (o *UserDTO) SetPreventNewSession(v bool)`

SetPreventNewSession sets PreventNewSession field to given value.

### HasPreventNewSession

`func (o *UserDTO) HasPreventNewSession() bool`

HasPreventNewSession returns a boolean if a field has been set.

### GetStateOrProvinceName

`func (o *UserDTO) GetStateOrProvinceName() string`

GetStateOrProvinceName returns the StateOrProvinceName field if non-nil, zero value otherwise.

### GetStateOrProvinceNameOk

`func (o *UserDTO) GetStateOrProvinceNameOk() (*string, bool)`

GetStateOrProvinceNameOk returns a tuple with the StateOrProvinceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvinceName

`func (o *UserDTO) SetStateOrProvinceName(v string)`

SetStateOrProvinceName sets StateOrProvinceName field to given value.

### HasStateOrProvinceName

`func (o *UserDTO) HasStateOrProvinceName() bool`

HasStateOrProvinceName returns a boolean if a field has been set.

### GetStreetAddress

`func (o *UserDTO) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UserDTO) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UserDTO) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *UserDTO) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetSurename

`func (o *UserDTO) GetSurename() string`

GetSurename returns the Surename field if non-nil, zero value otherwise.

### GetSurenameOk

`func (o *UserDTO) GetSurenameOk() (*string, bool)`

GetSurenameOk returns a tuple with the Surename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurename

`func (o *UserDTO) SetSurename(v string)`

SetSurename sets Surename field to given value.

### HasSurename

`func (o *UserDTO) HasSurename() bool`

HasSurename returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *UserDTO) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *UserDTO) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *UserDTO) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *UserDTO) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetTerminatePreviousSession

`func (o *UserDTO) GetTerminatePreviousSession() bool`

GetTerminatePreviousSession returns the TerminatePreviousSession field if non-nil, zero value otherwise.

### GetTerminatePreviousSessionOk

`func (o *UserDTO) GetTerminatePreviousSessionOk() (*bool, bool)`

GetTerminatePreviousSessionOk returns a tuple with the TerminatePreviousSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatePreviousSession

`func (o *UserDTO) SetTerminatePreviousSession(v bool)`

SetTerminatePreviousSession sets TerminatePreviousSession field to given value.

### HasTerminatePreviousSession

`func (o *UserDTO) HasTerminatePreviousSession() bool`

HasTerminatePreviousSession returns a boolean if a field has been set.

### GetUserCertificate

`func (o *UserDTO) GetUserCertificate() []string`

GetUserCertificate returns the UserCertificate field if non-nil, zero value otherwise.

### GetUserCertificateOk

`func (o *UserDTO) GetUserCertificateOk() (*[]string, bool)`

GetUserCertificateOk returns a tuple with the UserCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCertificate

`func (o *UserDTO) SetUserCertificate(v []string)`

SetUserCertificate sets UserCertificate field to given value.

### HasUserCertificate

`func (o *UserDTO) HasUserCertificate() bool`

HasUserCertificate returns a boolean if a field has been set.

### GetUserName

`func (o *UserDTO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserDTO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserDTO) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserDTO) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPassword

`func (o *UserDTO) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *UserDTO) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *UserDTO) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *UserDTO) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


