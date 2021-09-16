# InternalSaml2ServiceProviderChannelDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**AttributeProfile** | Pointer to [**AttributeProfileDTO**](AttributeProfileDTO.md) |  | [optional] 
**AuthenticationContract** | Pointer to [**AuthenticationContractDTO**](AuthenticationContractDTO.md) |  | [optional] 
**AuthenticationMechanism** | Pointer to [**AuthenticationMechanismDTO**](AuthenticationMechanismDTO.md) |  | [optional] 
**ConnectionA** | Pointer to [**FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**ConnectionB** | Pointer to [**FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**EmissionPolicy** | Pointer to [**AuthenticationAssertionEmissionPolicyDTO**](AuthenticationAssertionEmissionPolicyDTO.md) |  | [optional] 
**EncryptAssertion** | Pointer to **bool** |  | [optional] 
**EncryptAssertionAlgorithm** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IgnoreRequestedNameIDPolicy** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**MessageTtl** | Pointer to **int32** |  | [optional] 
**MessageTtlTolerance** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverrideProviderSetup** | Pointer to **bool** |  | [optional] 
**RequiredRoles** | Pointer to **[]string** |  | [optional] 
**RequiredRolesMatchMode** | Pointer to **int32** |  | [optional] 
**RestrictedRoles** | Pointer to **[]string** |  | [optional] 
**RestrictedRolesMatchMode** | Pointer to **int32** |  | [optional] 
**SignatureHash** | Pointer to **string** |  | [optional] 
**SubjectNameIDPolicy** | Pointer to [**SubjectNameIdentifierPolicyDTO**](SubjectNameIdentifierPolicyDTO.md) |  | [optional] 
**WantAuthnRequestsSigned** | Pointer to **bool** |  | [optional] 

## Methods

### NewInternalSaml2ServiceProviderChannelDTO

`func NewInternalSaml2ServiceProviderChannelDTO() *InternalSaml2ServiceProviderChannelDTO`

NewInternalSaml2ServiceProviderChannelDTO instantiates a new InternalSaml2ServiceProviderChannelDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalSaml2ServiceProviderChannelDTOWithDefaults

`func NewInternalSaml2ServiceProviderChannelDTOWithDefaults() *InternalSaml2ServiceProviderChannelDTO`

NewInternalSaml2ServiceProviderChannelDTOWithDefaults instantiates a new InternalSaml2ServiceProviderChannelDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveBindings

`func (o *InternalSaml2ServiceProviderChannelDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *InternalSaml2ServiceProviderChannelDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *InternalSaml2ServiceProviderChannelDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *InternalSaml2ServiceProviderChannelDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *InternalSaml2ServiceProviderChannelDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *InternalSaml2ServiceProviderChannelDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAttributeProfile

`func (o *InternalSaml2ServiceProviderChannelDTO) GetAttributeProfile() AttributeProfileDTO`

GetAttributeProfile returns the AttributeProfile field if non-nil, zero value otherwise.

### GetAttributeProfileOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetAttributeProfileOk() (*AttributeProfileDTO, bool)`

GetAttributeProfileOk returns a tuple with the AttributeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeProfile

`func (o *InternalSaml2ServiceProviderChannelDTO) SetAttributeProfile(v AttributeProfileDTO)`

SetAttributeProfile sets AttributeProfile field to given value.

### HasAttributeProfile

`func (o *InternalSaml2ServiceProviderChannelDTO) HasAttributeProfile() bool`

HasAttributeProfile returns a boolean if a field has been set.

### GetAuthenticationContract

`func (o *InternalSaml2ServiceProviderChannelDTO) GetAuthenticationContract() AuthenticationContractDTO`

GetAuthenticationContract returns the AuthenticationContract field if non-nil, zero value otherwise.

### GetAuthenticationContractOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetAuthenticationContractOk() (*AuthenticationContractDTO, bool)`

GetAuthenticationContractOk returns a tuple with the AuthenticationContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationContract

`func (o *InternalSaml2ServiceProviderChannelDTO) SetAuthenticationContract(v AuthenticationContractDTO)`

SetAuthenticationContract sets AuthenticationContract field to given value.

### HasAuthenticationContract

`func (o *InternalSaml2ServiceProviderChannelDTO) HasAuthenticationContract() bool`

HasAuthenticationContract returns a boolean if a field has been set.

### GetAuthenticationMechanism

`func (o *InternalSaml2ServiceProviderChannelDTO) GetAuthenticationMechanism() AuthenticationMechanismDTO`

GetAuthenticationMechanism returns the AuthenticationMechanism field if non-nil, zero value otherwise.

### GetAuthenticationMechanismOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetAuthenticationMechanismOk() (*AuthenticationMechanismDTO, bool)`

GetAuthenticationMechanismOk returns a tuple with the AuthenticationMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMechanism

`func (o *InternalSaml2ServiceProviderChannelDTO) SetAuthenticationMechanism(v AuthenticationMechanismDTO)`

SetAuthenticationMechanism sets AuthenticationMechanism field to given value.

### HasAuthenticationMechanism

`func (o *InternalSaml2ServiceProviderChannelDTO) HasAuthenticationMechanism() bool`

HasAuthenticationMechanism returns a boolean if a field has been set.

### GetConnectionA

`func (o *InternalSaml2ServiceProviderChannelDTO) GetConnectionA() FederatedConnectionDTO`

GetConnectionA returns the ConnectionA field if non-nil, zero value otherwise.

### GetConnectionAOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetConnectionAOk() (*FederatedConnectionDTO, bool)`

GetConnectionAOk returns a tuple with the ConnectionA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionA

`func (o *InternalSaml2ServiceProviderChannelDTO) SetConnectionA(v FederatedConnectionDTO)`

SetConnectionA sets ConnectionA field to given value.

### HasConnectionA

`func (o *InternalSaml2ServiceProviderChannelDTO) HasConnectionA() bool`

HasConnectionA returns a boolean if a field has been set.

### GetConnectionB

`func (o *InternalSaml2ServiceProviderChannelDTO) GetConnectionB() FederatedConnectionDTO`

GetConnectionB returns the ConnectionB field if non-nil, zero value otherwise.

### GetConnectionBOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetConnectionBOk() (*FederatedConnectionDTO, bool)`

GetConnectionBOk returns a tuple with the ConnectionB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionB

`func (o *InternalSaml2ServiceProviderChannelDTO) SetConnectionB(v FederatedConnectionDTO)`

SetConnectionB sets ConnectionB field to given value.

### HasConnectionB

`func (o *InternalSaml2ServiceProviderChannelDTO) HasConnectionB() bool`

HasConnectionB returns a boolean if a field has been set.

### GetDescription

`func (o *InternalSaml2ServiceProviderChannelDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalSaml2ServiceProviderChannelDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalSaml2ServiceProviderChannelDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *InternalSaml2ServiceProviderChannelDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InternalSaml2ServiceProviderChannelDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InternalSaml2ServiceProviderChannelDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *InternalSaml2ServiceProviderChannelDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *InternalSaml2ServiceProviderChannelDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *InternalSaml2ServiceProviderChannelDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEmissionPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) GetEmissionPolicy() AuthenticationAssertionEmissionPolicyDTO`

GetEmissionPolicy returns the EmissionPolicy field if non-nil, zero value otherwise.

### GetEmissionPolicyOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetEmissionPolicyOk() (*AuthenticationAssertionEmissionPolicyDTO, bool)`

GetEmissionPolicyOk returns a tuple with the EmissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmissionPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) SetEmissionPolicy(v AuthenticationAssertionEmissionPolicyDTO)`

SetEmissionPolicy sets EmissionPolicy field to given value.

### HasEmissionPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) HasEmissionPolicy() bool`

HasEmissionPolicy returns a boolean if a field has been set.

### GetEncryptAssertion

`func (o *InternalSaml2ServiceProviderChannelDTO) GetEncryptAssertion() bool`

GetEncryptAssertion returns the EncryptAssertion field if non-nil, zero value otherwise.

### GetEncryptAssertionOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetEncryptAssertionOk() (*bool, bool)`

GetEncryptAssertionOk returns a tuple with the EncryptAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertion

`func (o *InternalSaml2ServiceProviderChannelDTO) SetEncryptAssertion(v bool)`

SetEncryptAssertion sets EncryptAssertion field to given value.

### HasEncryptAssertion

`func (o *InternalSaml2ServiceProviderChannelDTO) HasEncryptAssertion() bool`

HasEncryptAssertion returns a boolean if a field has been set.

### GetEncryptAssertionAlgorithm

`func (o *InternalSaml2ServiceProviderChannelDTO) GetEncryptAssertionAlgorithm() string`

GetEncryptAssertionAlgorithm returns the EncryptAssertionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptAssertionAlgorithmOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetEncryptAssertionAlgorithmOk() (*string, bool)`

GetEncryptAssertionAlgorithmOk returns a tuple with the EncryptAssertionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertionAlgorithm

`func (o *InternalSaml2ServiceProviderChannelDTO) SetEncryptAssertionAlgorithm(v string)`

SetEncryptAssertionAlgorithm sets EncryptAssertionAlgorithm field to given value.

### HasEncryptAssertionAlgorithm

`func (o *InternalSaml2ServiceProviderChannelDTO) HasEncryptAssertionAlgorithm() bool`

HasEncryptAssertionAlgorithm returns a boolean if a field has been set.

### GetId

`func (o *InternalSaml2ServiceProviderChannelDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalSaml2ServiceProviderChannelDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InternalSaml2ServiceProviderChannelDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreRequestedNameIDPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) GetIgnoreRequestedNameIDPolicy() bool`

GetIgnoreRequestedNameIDPolicy returns the IgnoreRequestedNameIDPolicy field if non-nil, zero value otherwise.

### GetIgnoreRequestedNameIDPolicyOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetIgnoreRequestedNameIDPolicyOk() (*bool, bool)`

GetIgnoreRequestedNameIDPolicyOk returns a tuple with the IgnoreRequestedNameIDPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRequestedNameIDPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) SetIgnoreRequestedNameIDPolicy(v bool)`

SetIgnoreRequestedNameIDPolicy sets IgnoreRequestedNameIDPolicy field to given value.

### HasIgnoreRequestedNameIDPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) HasIgnoreRequestedNameIDPolicy() bool`

HasIgnoreRequestedNameIDPolicy returns a boolean if a field has been set.

### GetLocation

`func (o *InternalSaml2ServiceProviderChannelDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InternalSaml2ServiceProviderChannelDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InternalSaml2ServiceProviderChannelDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMessageTtl

`func (o *InternalSaml2ServiceProviderChannelDTO) GetMessageTtl() int32`

GetMessageTtl returns the MessageTtl field if non-nil, zero value otherwise.

### GetMessageTtlOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetMessageTtlOk() (*int32, bool)`

GetMessageTtlOk returns a tuple with the MessageTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtl

`func (o *InternalSaml2ServiceProviderChannelDTO) SetMessageTtl(v int32)`

SetMessageTtl sets MessageTtl field to given value.

### HasMessageTtl

`func (o *InternalSaml2ServiceProviderChannelDTO) HasMessageTtl() bool`

HasMessageTtl returns a boolean if a field has been set.

### GetMessageTtlTolerance

`func (o *InternalSaml2ServiceProviderChannelDTO) GetMessageTtlTolerance() int32`

GetMessageTtlTolerance returns the MessageTtlTolerance field if non-nil, zero value otherwise.

### GetMessageTtlToleranceOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetMessageTtlToleranceOk() (*int32, bool)`

GetMessageTtlToleranceOk returns a tuple with the MessageTtlTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtlTolerance

`func (o *InternalSaml2ServiceProviderChannelDTO) SetMessageTtlTolerance(v int32)`

SetMessageTtlTolerance sets MessageTtlTolerance field to given value.

### HasMessageTtlTolerance

`func (o *InternalSaml2ServiceProviderChannelDTO) HasMessageTtlTolerance() bool`

HasMessageTtlTolerance returns a boolean if a field has been set.

### GetName

`func (o *InternalSaml2ServiceProviderChannelDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalSaml2ServiceProviderChannelDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalSaml2ServiceProviderChannelDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideProviderSetup

`func (o *InternalSaml2ServiceProviderChannelDTO) GetOverrideProviderSetup() bool`

GetOverrideProviderSetup returns the OverrideProviderSetup field if non-nil, zero value otherwise.

### GetOverrideProviderSetupOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetOverrideProviderSetupOk() (*bool, bool)`

GetOverrideProviderSetupOk returns a tuple with the OverrideProviderSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideProviderSetup

`func (o *InternalSaml2ServiceProviderChannelDTO) SetOverrideProviderSetup(v bool)`

SetOverrideProviderSetup sets OverrideProviderSetup field to given value.

### HasOverrideProviderSetup

`func (o *InternalSaml2ServiceProviderChannelDTO) HasOverrideProviderSetup() bool`

HasOverrideProviderSetup returns a boolean if a field has been set.

### GetRequiredRoles

`func (o *InternalSaml2ServiceProviderChannelDTO) GetRequiredRoles() []string`

GetRequiredRoles returns the RequiredRoles field if non-nil, zero value otherwise.

### GetRequiredRolesOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetRequiredRolesOk() (*[]string, bool)`

GetRequiredRolesOk returns a tuple with the RequiredRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRoles

`func (o *InternalSaml2ServiceProviderChannelDTO) SetRequiredRoles(v []string)`

SetRequiredRoles sets RequiredRoles field to given value.

### HasRequiredRoles

`func (o *InternalSaml2ServiceProviderChannelDTO) HasRequiredRoles() bool`

HasRequiredRoles returns a boolean if a field has been set.

### GetRequiredRolesMatchMode

`func (o *InternalSaml2ServiceProviderChannelDTO) GetRequiredRolesMatchMode() int32`

GetRequiredRolesMatchMode returns the RequiredRolesMatchMode field if non-nil, zero value otherwise.

### GetRequiredRolesMatchModeOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetRequiredRolesMatchModeOk() (*int32, bool)`

GetRequiredRolesMatchModeOk returns a tuple with the RequiredRolesMatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRolesMatchMode

`func (o *InternalSaml2ServiceProviderChannelDTO) SetRequiredRolesMatchMode(v int32)`

SetRequiredRolesMatchMode sets RequiredRolesMatchMode field to given value.

### HasRequiredRolesMatchMode

`func (o *InternalSaml2ServiceProviderChannelDTO) HasRequiredRolesMatchMode() bool`

HasRequiredRolesMatchMode returns a boolean if a field has been set.

### GetRestrictedRoles

`func (o *InternalSaml2ServiceProviderChannelDTO) GetRestrictedRoles() []string`

GetRestrictedRoles returns the RestrictedRoles field if non-nil, zero value otherwise.

### GetRestrictedRolesOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetRestrictedRolesOk() (*[]string, bool)`

GetRestrictedRolesOk returns a tuple with the RestrictedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedRoles

`func (o *InternalSaml2ServiceProviderChannelDTO) SetRestrictedRoles(v []string)`

SetRestrictedRoles sets RestrictedRoles field to given value.

### HasRestrictedRoles

`func (o *InternalSaml2ServiceProviderChannelDTO) HasRestrictedRoles() bool`

HasRestrictedRoles returns a boolean if a field has been set.

### GetRestrictedRolesMatchMode

`func (o *InternalSaml2ServiceProviderChannelDTO) GetRestrictedRolesMatchMode() int32`

GetRestrictedRolesMatchMode returns the RestrictedRolesMatchMode field if non-nil, zero value otherwise.

### GetRestrictedRolesMatchModeOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetRestrictedRolesMatchModeOk() (*int32, bool)`

GetRestrictedRolesMatchModeOk returns a tuple with the RestrictedRolesMatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedRolesMatchMode

`func (o *InternalSaml2ServiceProviderChannelDTO) SetRestrictedRolesMatchMode(v int32)`

SetRestrictedRolesMatchMode sets RestrictedRolesMatchMode field to given value.

### HasRestrictedRolesMatchMode

`func (o *InternalSaml2ServiceProviderChannelDTO) HasRestrictedRolesMatchMode() bool`

HasRestrictedRolesMatchMode returns a boolean if a field has been set.

### GetSignatureHash

`func (o *InternalSaml2ServiceProviderChannelDTO) GetSignatureHash() string`

GetSignatureHash returns the SignatureHash field if non-nil, zero value otherwise.

### GetSignatureHashOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetSignatureHashOk() (*string, bool)`

GetSignatureHashOk returns a tuple with the SignatureHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureHash

`func (o *InternalSaml2ServiceProviderChannelDTO) SetSignatureHash(v string)`

SetSignatureHash sets SignatureHash field to given value.

### HasSignatureHash

`func (o *InternalSaml2ServiceProviderChannelDTO) HasSignatureHash() bool`

HasSignatureHash returns a boolean if a field has been set.

### GetSubjectNameIDPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) GetSubjectNameIDPolicy() SubjectNameIdentifierPolicyDTO`

GetSubjectNameIDPolicy returns the SubjectNameIDPolicy field if non-nil, zero value otherwise.

### GetSubjectNameIDPolicyOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetSubjectNameIDPolicyOk() (*SubjectNameIdentifierPolicyDTO, bool)`

GetSubjectNameIDPolicyOk returns a tuple with the SubjectNameIDPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameIDPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) SetSubjectNameIDPolicy(v SubjectNameIdentifierPolicyDTO)`

SetSubjectNameIDPolicy sets SubjectNameIDPolicy field to given value.

### HasSubjectNameIDPolicy

`func (o *InternalSaml2ServiceProviderChannelDTO) HasSubjectNameIDPolicy() bool`

HasSubjectNameIDPolicy returns a boolean if a field has been set.

### GetWantAuthnRequestsSigned

`func (o *InternalSaml2ServiceProviderChannelDTO) GetWantAuthnRequestsSigned() bool`

GetWantAuthnRequestsSigned returns the WantAuthnRequestsSigned field if non-nil, zero value otherwise.

### GetWantAuthnRequestsSignedOk

`func (o *InternalSaml2ServiceProviderChannelDTO) GetWantAuthnRequestsSignedOk() (*bool, bool)`

GetWantAuthnRequestsSignedOk returns a tuple with the WantAuthnRequestsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantAuthnRequestsSigned

`func (o *InternalSaml2ServiceProviderChannelDTO) SetWantAuthnRequestsSigned(v bool)`

SetWantAuthnRequestsSigned sets WantAuthnRequestsSigned field to given value.

### HasWantAuthnRequestsSigned

`func (o *InternalSaml2ServiceProviderChannelDTO) HasWantAuthnRequestsSigned() bool`

HasWantAuthnRequestsSigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


