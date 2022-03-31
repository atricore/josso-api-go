# IdentityApplianceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdApplianceDefinition** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdApplianceDefinitionBin** | Pointer to **string** |  | [optional] 
**IdApplianceDeployment** | Pointer to [**IdentityApplianceDeploymentDTO**](IdentityApplianceDeploymentDTO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityApplianceDTO

`func NewIdentityApplianceDTO() *IdentityApplianceDTO`

NewIdentityApplianceDTO instantiates a new IdentityApplianceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityApplianceDTOWithDefaults

`func NewIdentityApplianceDTOWithDefaults() *IdentityApplianceDTO`

NewIdentityApplianceDTOWithDefaults instantiates a new IdentityApplianceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IdentityApplianceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityApplianceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityApplianceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityApplianceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityApplianceDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityApplianceDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityApplianceDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityApplianceDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *IdentityApplianceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *IdentityApplianceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *IdentityApplianceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *IdentityApplianceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *IdentityApplianceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityApplianceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityApplianceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityApplianceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdApplianceDefinition

`func (o *IdentityApplianceDTO) GetIdApplianceDefinition() IdentityApplianceDefinitionDTO`

GetIdApplianceDefinition returns the IdApplianceDefinition field if non-nil, zero value otherwise.

### GetIdApplianceDefinitionOk

`func (o *IdentityApplianceDTO) GetIdApplianceDefinitionOk() (*IdentityApplianceDefinitionDTO, bool)`

GetIdApplianceDefinitionOk returns a tuple with the IdApplianceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdApplianceDefinition

`func (o *IdentityApplianceDTO) SetIdApplianceDefinition(v IdentityApplianceDefinitionDTO)`

SetIdApplianceDefinition sets IdApplianceDefinition field to given value.

### HasIdApplianceDefinition

`func (o *IdentityApplianceDTO) HasIdApplianceDefinition() bool`

HasIdApplianceDefinition returns a boolean if a field has been set.

### GetIdApplianceDefinitionBin

`func (o *IdentityApplianceDTO) GetIdApplianceDefinitionBin() string`

GetIdApplianceDefinitionBin returns the IdApplianceDefinitionBin field if non-nil, zero value otherwise.

### GetIdApplianceDefinitionBinOk

`func (o *IdentityApplianceDTO) GetIdApplianceDefinitionBinOk() (*string, bool)`

GetIdApplianceDefinitionBinOk returns a tuple with the IdApplianceDefinitionBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdApplianceDefinitionBin

`func (o *IdentityApplianceDTO) SetIdApplianceDefinitionBin(v string)`

SetIdApplianceDefinitionBin sets IdApplianceDefinitionBin field to given value.

### HasIdApplianceDefinitionBin

`func (o *IdentityApplianceDTO) HasIdApplianceDefinitionBin() bool`

HasIdApplianceDefinitionBin returns a boolean if a field has been set.

### GetIdApplianceDeployment

`func (o *IdentityApplianceDTO) GetIdApplianceDeployment() IdentityApplianceDeploymentDTO`

GetIdApplianceDeployment returns the IdApplianceDeployment field if non-nil, zero value otherwise.

### GetIdApplianceDeploymentOk

`func (o *IdentityApplianceDTO) GetIdApplianceDeploymentOk() (*IdentityApplianceDeploymentDTO, bool)`

GetIdApplianceDeploymentOk returns a tuple with the IdApplianceDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdApplianceDeployment

`func (o *IdentityApplianceDTO) SetIdApplianceDeployment(v IdentityApplianceDeploymentDTO)`

SetIdApplianceDeployment sets IdApplianceDeployment field to given value.

### HasIdApplianceDeployment

`func (o *IdentityApplianceDTO) HasIdApplianceDeployment() bool`

HasIdApplianceDeployment returns a boolean if a field has been set.

### GetName

`func (o *IdentityApplianceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityApplianceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityApplianceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityApplianceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *IdentityApplianceDTO) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IdentityApplianceDTO) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IdentityApplianceDTO) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IdentityApplianceDTO) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetState

`func (o *IdentityApplianceDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IdentityApplianceDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IdentityApplianceDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IdentityApplianceDTO) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


