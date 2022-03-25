# IdentityProviderChannelDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLinkagePolicy** | Pointer to [**AccountLinkagePolicyDTO**](AccountLinkagePolicyDTO.md) |  | [optional] 
**ActiveBindings** | Pointer to **[]string** |  | [optional] 
**ActiveProfiles** | Pointer to **[]string** |  | [optional] 
**AuthenticationContract** | Pointer to [**AuthenticationContractDTO**](AuthenticationContractDTO.md) |  | [optional] 
**AuthenticationMechanism** | Pointer to [**AuthenticationMechanismDTO**](AuthenticationMechanismDTO.md) |  | [optional] 
**ConnectionA** | Pointer to [**FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**ConnectionB** | Pointer to [**FederatedConnectionDTO**](FederatedConnectionDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**EnableProxyExtension** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdentityMappingPolicy** | Pointer to [**IdentityMappingPolicyDTO**](IdentityMappingPolicyDTO.md) |  | [optional] 
**Location** | Pointer to [**LocationDTO**](LocationDTO.md) |  | [optional] 
**MessageTtl** | Pointer to **int32** |  | [optional] 
**MessageTtlTolerance** | Pointer to **int32** |  | [optional] 
**MultivaluedAttrGroups** | Pointer to **bool** |  | [optional] 
**MultivaluedAttrInternal** | Pointer to **bool** |  | [optional] 
**MultivaluedAttrUserDefined** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverrideProviderSetup** | Pointer to **bool** |  | [optional] 
**Preferred** | Pointer to **bool** |  | [optional] 
**SignAuthenticationRequests** | Pointer to **bool** |  | [optional] 
**SignatureHash** | Pointer to **string** |  | [optional] 
**WantAssertionSigned** | Pointer to **bool** |  | [optional] 

## Methods

### NewIdentityProviderChannelDTO

`func NewIdentityProviderChannelDTO() *IdentityProviderChannelDTO`

NewIdentityProviderChannelDTO instantiates a new IdentityProviderChannelDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderChannelDTOWithDefaults

`func NewIdentityProviderChannelDTOWithDefaults() *IdentityProviderChannelDTO`

NewIdentityProviderChannelDTOWithDefaults instantiates a new IdentityProviderChannelDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLinkagePolicy

`func (o *IdentityProviderChannelDTO) GetAccountLinkagePolicy() AccountLinkagePolicyDTO`

GetAccountLinkagePolicy returns the AccountLinkagePolicy field if non-nil, zero value otherwise.

### GetAccountLinkagePolicyOk

`func (o *IdentityProviderChannelDTO) GetAccountLinkagePolicyOk() (*AccountLinkagePolicyDTO, bool)`

GetAccountLinkagePolicyOk returns a tuple with the AccountLinkagePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLinkagePolicy

`func (o *IdentityProviderChannelDTO) SetAccountLinkagePolicy(v AccountLinkagePolicyDTO)`

SetAccountLinkagePolicy sets AccountLinkagePolicy field to given value.

### HasAccountLinkagePolicy

`func (o *IdentityProviderChannelDTO) HasAccountLinkagePolicy() bool`

HasAccountLinkagePolicy returns a boolean if a field has been set.

### GetActiveBindings

`func (o *IdentityProviderChannelDTO) GetActiveBindings() []string`

GetActiveBindings returns the ActiveBindings field if non-nil, zero value otherwise.

### GetActiveBindingsOk

`func (o *IdentityProviderChannelDTO) GetActiveBindingsOk() (*[]string, bool)`

GetActiveBindingsOk returns a tuple with the ActiveBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBindings

`func (o *IdentityProviderChannelDTO) SetActiveBindings(v []string)`

SetActiveBindings sets ActiveBindings field to given value.

### HasActiveBindings

`func (o *IdentityProviderChannelDTO) HasActiveBindings() bool`

HasActiveBindings returns a boolean if a field has been set.

### GetActiveProfiles

`func (o *IdentityProviderChannelDTO) GetActiveProfiles() []string`

GetActiveProfiles returns the ActiveProfiles field if non-nil, zero value otherwise.

### GetActiveProfilesOk

`func (o *IdentityProviderChannelDTO) GetActiveProfilesOk() (*[]string, bool)`

GetActiveProfilesOk returns a tuple with the ActiveProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfiles

`func (o *IdentityProviderChannelDTO) SetActiveProfiles(v []string)`

SetActiveProfiles sets ActiveProfiles field to given value.

### HasActiveProfiles

`func (o *IdentityProviderChannelDTO) HasActiveProfiles() bool`

HasActiveProfiles returns a boolean if a field has been set.

### GetAuthenticationContract

`func (o *IdentityProviderChannelDTO) GetAuthenticationContract() AuthenticationContractDTO`

GetAuthenticationContract returns the AuthenticationContract field if non-nil, zero value otherwise.

### GetAuthenticationContractOk

`func (o *IdentityProviderChannelDTO) GetAuthenticationContractOk() (*AuthenticationContractDTO, bool)`

GetAuthenticationContractOk returns a tuple with the AuthenticationContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationContract

`func (o *IdentityProviderChannelDTO) SetAuthenticationContract(v AuthenticationContractDTO)`

SetAuthenticationContract sets AuthenticationContract field to given value.

### HasAuthenticationContract

`func (o *IdentityProviderChannelDTO) HasAuthenticationContract() bool`

HasAuthenticationContract returns a boolean if a field has been set.

### GetAuthenticationMechanism

`func (o *IdentityProviderChannelDTO) GetAuthenticationMechanism() AuthenticationMechanismDTO`

GetAuthenticationMechanism returns the AuthenticationMechanism field if non-nil, zero value otherwise.

### GetAuthenticationMechanismOk

`func (o *IdentityProviderChannelDTO) GetAuthenticationMechanismOk() (*AuthenticationMechanismDTO, bool)`

GetAuthenticationMechanismOk returns a tuple with the AuthenticationMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMechanism

`func (o *IdentityProviderChannelDTO) SetAuthenticationMechanism(v AuthenticationMechanismDTO)`

SetAuthenticationMechanism sets AuthenticationMechanism field to given value.

### HasAuthenticationMechanism

`func (o *IdentityProviderChannelDTO) HasAuthenticationMechanism() bool`

HasAuthenticationMechanism returns a boolean if a field has been set.

### GetConnectionA

`func (o *IdentityProviderChannelDTO) GetConnectionA() FederatedConnectionDTO`

GetConnectionA returns the ConnectionA field if non-nil, zero value otherwise.

### GetConnectionAOk

`func (o *IdentityProviderChannelDTO) GetConnectionAOk() (*FederatedConnectionDTO, bool)`

GetConnectionAOk returns a tuple with the ConnectionA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionA

`func (o *IdentityProviderChannelDTO) SetConnectionA(v FederatedConnectionDTO)`

SetConnectionA sets ConnectionA field to given value.

### HasConnectionA

`func (o *IdentityProviderChannelDTO) HasConnectionA() bool`

HasConnectionA returns a boolean if a field has been set.

### GetConnectionB

`func (o *IdentityProviderChannelDTO) GetConnectionB() FederatedConnectionDTO`

GetConnectionB returns the ConnectionB field if non-nil, zero value otherwise.

### GetConnectionBOk

`func (o *IdentityProviderChannelDTO) GetConnectionBOk() (*FederatedConnectionDTO, bool)`

GetConnectionBOk returns a tuple with the ConnectionB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionB

`func (o *IdentityProviderChannelDTO) SetConnectionB(v FederatedConnectionDTO)`

SetConnectionB sets ConnectionB field to given value.

### HasConnectionB

`func (o *IdentityProviderChannelDTO) HasConnectionB() bool`

HasConnectionB returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityProviderChannelDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProviderChannelDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProviderChannelDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProviderChannelDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityProviderChannelDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderChannelDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderChannelDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProviderChannelDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *IdentityProviderChannelDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *IdentityProviderChannelDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *IdentityProviderChannelDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *IdentityProviderChannelDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEnableProxyExtension

`func (o *IdentityProviderChannelDTO) GetEnableProxyExtension() bool`

GetEnableProxyExtension returns the EnableProxyExtension field if non-nil, zero value otherwise.

### GetEnableProxyExtensionOk

`func (o *IdentityProviderChannelDTO) GetEnableProxyExtensionOk() (*bool, bool)`

GetEnableProxyExtensionOk returns a tuple with the EnableProxyExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxyExtension

`func (o *IdentityProviderChannelDTO) SetEnableProxyExtension(v bool)`

SetEnableProxyExtension sets EnableProxyExtension field to given value.

### HasEnableProxyExtension

`func (o *IdentityProviderChannelDTO) HasEnableProxyExtension() bool`

HasEnableProxyExtension returns a boolean if a field has been set.

### GetId

`func (o *IdentityProviderChannelDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderChannelDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderChannelDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderChannelDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityMappingPolicy

`func (o *IdentityProviderChannelDTO) GetIdentityMappingPolicy() IdentityMappingPolicyDTO`

GetIdentityMappingPolicy returns the IdentityMappingPolicy field if non-nil, zero value otherwise.

### GetIdentityMappingPolicyOk

`func (o *IdentityProviderChannelDTO) GetIdentityMappingPolicyOk() (*IdentityMappingPolicyDTO, bool)`

GetIdentityMappingPolicyOk returns a tuple with the IdentityMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMappingPolicy

`func (o *IdentityProviderChannelDTO) SetIdentityMappingPolicy(v IdentityMappingPolicyDTO)`

SetIdentityMappingPolicy sets IdentityMappingPolicy field to given value.

### HasIdentityMappingPolicy

`func (o *IdentityProviderChannelDTO) HasIdentityMappingPolicy() bool`

HasIdentityMappingPolicy returns a boolean if a field has been set.

### GetLocation

`func (o *IdentityProviderChannelDTO) GetLocation() LocationDTO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IdentityProviderChannelDTO) GetLocationOk() (*LocationDTO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IdentityProviderChannelDTO) SetLocation(v LocationDTO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IdentityProviderChannelDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMessageTtl

`func (o *IdentityProviderChannelDTO) GetMessageTtl() int32`

GetMessageTtl returns the MessageTtl field if non-nil, zero value otherwise.

### GetMessageTtlOk

`func (o *IdentityProviderChannelDTO) GetMessageTtlOk() (*int32, bool)`

GetMessageTtlOk returns a tuple with the MessageTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtl

`func (o *IdentityProviderChannelDTO) SetMessageTtl(v int32)`

SetMessageTtl sets MessageTtl field to given value.

### HasMessageTtl

`func (o *IdentityProviderChannelDTO) HasMessageTtl() bool`

HasMessageTtl returns a boolean if a field has been set.

### GetMessageTtlTolerance

`func (o *IdentityProviderChannelDTO) GetMessageTtlTolerance() int32`

GetMessageTtlTolerance returns the MessageTtlTolerance field if non-nil, zero value otherwise.

### GetMessageTtlToleranceOk

`func (o *IdentityProviderChannelDTO) GetMessageTtlToleranceOk() (*int32, bool)`

GetMessageTtlToleranceOk returns a tuple with the MessageTtlTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtlTolerance

`func (o *IdentityProviderChannelDTO) SetMessageTtlTolerance(v int32)`

SetMessageTtlTolerance sets MessageTtlTolerance field to given value.

### HasMessageTtlTolerance

`func (o *IdentityProviderChannelDTO) HasMessageTtlTolerance() bool`

HasMessageTtlTolerance returns a boolean if a field has been set.

### GetMultivaluedAttrGroups

`func (o *IdentityProviderChannelDTO) GetMultivaluedAttrGroups() bool`

GetMultivaluedAttrGroups returns the MultivaluedAttrGroups field if non-nil, zero value otherwise.

### GetMultivaluedAttrGroupsOk

`func (o *IdentityProviderChannelDTO) GetMultivaluedAttrGroupsOk() (*bool, bool)`

GetMultivaluedAttrGroupsOk returns a tuple with the MultivaluedAttrGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivaluedAttrGroups

`func (o *IdentityProviderChannelDTO) SetMultivaluedAttrGroups(v bool)`

SetMultivaluedAttrGroups sets MultivaluedAttrGroups field to given value.

### HasMultivaluedAttrGroups

`func (o *IdentityProviderChannelDTO) HasMultivaluedAttrGroups() bool`

HasMultivaluedAttrGroups returns a boolean if a field has been set.

### GetMultivaluedAttrInternal

`func (o *IdentityProviderChannelDTO) GetMultivaluedAttrInternal() bool`

GetMultivaluedAttrInternal returns the MultivaluedAttrInternal field if non-nil, zero value otherwise.

### GetMultivaluedAttrInternalOk

`func (o *IdentityProviderChannelDTO) GetMultivaluedAttrInternalOk() (*bool, bool)`

GetMultivaluedAttrInternalOk returns a tuple with the MultivaluedAttrInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivaluedAttrInternal

`func (o *IdentityProviderChannelDTO) SetMultivaluedAttrInternal(v bool)`

SetMultivaluedAttrInternal sets MultivaluedAttrInternal field to given value.

### HasMultivaluedAttrInternal

`func (o *IdentityProviderChannelDTO) HasMultivaluedAttrInternal() bool`

HasMultivaluedAttrInternal returns a boolean if a field has been set.

### GetMultivaluedAttrUserDefined

`func (o *IdentityProviderChannelDTO) GetMultivaluedAttrUserDefined() bool`

GetMultivaluedAttrUserDefined returns the MultivaluedAttrUserDefined field if non-nil, zero value otherwise.

### GetMultivaluedAttrUserDefinedOk

`func (o *IdentityProviderChannelDTO) GetMultivaluedAttrUserDefinedOk() (*bool, bool)`

GetMultivaluedAttrUserDefinedOk returns a tuple with the MultivaluedAttrUserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivaluedAttrUserDefined

`func (o *IdentityProviderChannelDTO) SetMultivaluedAttrUserDefined(v bool)`

SetMultivaluedAttrUserDefined sets MultivaluedAttrUserDefined field to given value.

### HasMultivaluedAttrUserDefined

`func (o *IdentityProviderChannelDTO) HasMultivaluedAttrUserDefined() bool`

HasMultivaluedAttrUserDefined returns a boolean if a field has been set.

### GetName

`func (o *IdentityProviderChannelDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderChannelDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderChannelDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityProviderChannelDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideProviderSetup

`func (o *IdentityProviderChannelDTO) GetOverrideProviderSetup() bool`

GetOverrideProviderSetup returns the OverrideProviderSetup field if non-nil, zero value otherwise.

### GetOverrideProviderSetupOk

`func (o *IdentityProviderChannelDTO) GetOverrideProviderSetupOk() (*bool, bool)`

GetOverrideProviderSetupOk returns a tuple with the OverrideProviderSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideProviderSetup

`func (o *IdentityProviderChannelDTO) SetOverrideProviderSetup(v bool)`

SetOverrideProviderSetup sets OverrideProviderSetup field to given value.

### HasOverrideProviderSetup

`func (o *IdentityProviderChannelDTO) HasOverrideProviderSetup() bool`

HasOverrideProviderSetup returns a boolean if a field has been set.

### GetPreferred

`func (o *IdentityProviderChannelDTO) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *IdentityProviderChannelDTO) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *IdentityProviderChannelDTO) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.

### HasPreferred

`func (o *IdentityProviderChannelDTO) HasPreferred() bool`

HasPreferred returns a boolean if a field has been set.

### GetSignAuthenticationRequests

`func (o *IdentityProviderChannelDTO) GetSignAuthenticationRequests() bool`

GetSignAuthenticationRequests returns the SignAuthenticationRequests field if non-nil, zero value otherwise.

### GetSignAuthenticationRequestsOk

`func (o *IdentityProviderChannelDTO) GetSignAuthenticationRequestsOk() (*bool, bool)`

GetSignAuthenticationRequestsOk returns a tuple with the SignAuthenticationRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthenticationRequests

`func (o *IdentityProviderChannelDTO) SetSignAuthenticationRequests(v bool)`

SetSignAuthenticationRequests sets SignAuthenticationRequests field to given value.

### HasSignAuthenticationRequests

`func (o *IdentityProviderChannelDTO) HasSignAuthenticationRequests() bool`

HasSignAuthenticationRequests returns a boolean if a field has been set.

### GetSignatureHash

`func (o *IdentityProviderChannelDTO) GetSignatureHash() string`

GetSignatureHash returns the SignatureHash field if non-nil, zero value otherwise.

### GetSignatureHashOk

`func (o *IdentityProviderChannelDTO) GetSignatureHashOk() (*string, bool)`

GetSignatureHashOk returns a tuple with the SignatureHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureHash

`func (o *IdentityProviderChannelDTO) SetSignatureHash(v string)`

SetSignatureHash sets SignatureHash field to given value.

### HasSignatureHash

`func (o *IdentityProviderChannelDTO) HasSignatureHash() bool`

HasSignatureHash returns a boolean if a field has been set.

### GetWantAssertionSigned

`func (o *IdentityProviderChannelDTO) GetWantAssertionSigned() bool`

GetWantAssertionSigned returns the WantAssertionSigned field if non-nil, zero value otherwise.

### GetWantAssertionSignedOk

`func (o *IdentityProviderChannelDTO) GetWantAssertionSignedOk() (*bool, bool)`

GetWantAssertionSignedOk returns a tuple with the WantAssertionSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantAssertionSigned

`func (o *IdentityProviderChannelDTO) SetWantAssertionSigned(v bool)`

SetWantAssertionSigned sets WantAssertionSigned field to given value.

### HasWantAssertionSigned

`func (o *IdentityProviderChannelDTO) HasWantAssertionSigned() bool`

HasWantAssertionSigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


