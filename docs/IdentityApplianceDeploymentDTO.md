# IdentityApplianceDeploymentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployedRevision** | Pointer to **int32** |  | [optional] 
**DeploymentTime** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**FeatureName** | Pointer to **string** |  | [optional] 
**FeatureUri** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Idaus** | Pointer to [**[]IdentityApplianceUnitDTO**](IdentityApplianceUnitDTO.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityApplianceDeploymentDTO

`func NewIdentityApplianceDeploymentDTO() *IdentityApplianceDeploymentDTO`

NewIdentityApplianceDeploymentDTO instantiates a new IdentityApplianceDeploymentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityApplianceDeploymentDTOWithDefaults

`func NewIdentityApplianceDeploymentDTOWithDefaults() *IdentityApplianceDeploymentDTO`

NewIdentityApplianceDeploymentDTOWithDefaults instantiates a new IdentityApplianceDeploymentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployedRevision

`func (o *IdentityApplianceDeploymentDTO) GetDeployedRevision() int32`

GetDeployedRevision returns the DeployedRevision field if non-nil, zero value otherwise.

### GetDeployedRevisionOk

`func (o *IdentityApplianceDeploymentDTO) GetDeployedRevisionOk() (*int32, bool)`

GetDeployedRevisionOk returns a tuple with the DeployedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedRevision

`func (o *IdentityApplianceDeploymentDTO) SetDeployedRevision(v int32)`

SetDeployedRevision sets DeployedRevision field to given value.

### HasDeployedRevision

`func (o *IdentityApplianceDeploymentDTO) HasDeployedRevision() bool`

HasDeployedRevision returns a boolean if a field has been set.

### GetDeploymentTime

`func (o *IdentityApplianceDeploymentDTO) GetDeploymentTime() time.Time`

GetDeploymentTime returns the DeploymentTime field if non-nil, zero value otherwise.

### GetDeploymentTimeOk

`func (o *IdentityApplianceDeploymentDTO) GetDeploymentTimeOk() (*time.Time, bool)`

GetDeploymentTimeOk returns a tuple with the DeploymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTime

`func (o *IdentityApplianceDeploymentDTO) SetDeploymentTime(v time.Time)`

SetDeploymentTime sets DeploymentTime field to given value.

### HasDeploymentTime

`func (o *IdentityApplianceDeploymentDTO) HasDeploymentTime() bool`

HasDeploymentTime returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityApplianceDeploymentDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityApplianceDeploymentDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityApplianceDeploymentDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityApplianceDeploymentDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *IdentityApplianceDeploymentDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *IdentityApplianceDeploymentDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *IdentityApplianceDeploymentDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *IdentityApplianceDeploymentDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetFeatureName

`func (o *IdentityApplianceDeploymentDTO) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *IdentityApplianceDeploymentDTO) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *IdentityApplianceDeploymentDTO) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *IdentityApplianceDeploymentDTO) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetFeatureUri

`func (o *IdentityApplianceDeploymentDTO) GetFeatureUri() string`

GetFeatureUri returns the FeatureUri field if non-nil, zero value otherwise.

### GetFeatureUriOk

`func (o *IdentityApplianceDeploymentDTO) GetFeatureUriOk() (*string, bool)`

GetFeatureUriOk returns a tuple with the FeatureUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUri

`func (o *IdentityApplianceDeploymentDTO) SetFeatureUri(v string)`

SetFeatureUri sets FeatureUri field to given value.

### HasFeatureUri

`func (o *IdentityApplianceDeploymentDTO) HasFeatureUri() bool`

HasFeatureUri returns a boolean if a field has been set.

### GetId

`func (o *IdentityApplianceDeploymentDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityApplianceDeploymentDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityApplianceDeploymentDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityApplianceDeploymentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdaus

`func (o *IdentityApplianceDeploymentDTO) GetIdaus() []IdentityApplianceUnitDTO`

GetIdaus returns the Idaus field if non-nil, zero value otherwise.

### GetIdausOk

`func (o *IdentityApplianceDeploymentDTO) GetIdausOk() (*[]IdentityApplianceUnitDTO, bool)`

GetIdausOk returns a tuple with the Idaus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdaus

`func (o *IdentityApplianceDeploymentDTO) SetIdaus(v []IdentityApplianceUnitDTO)`

SetIdaus sets Idaus field to given value.

### HasIdaus

`func (o *IdentityApplianceDeploymentDTO) HasIdaus() bool`

HasIdaus returns a boolean if a field has been set.

### GetState

`func (o *IdentityApplianceDeploymentDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IdentityApplianceDeploymentDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IdentityApplianceDeploymentDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IdentityApplianceDeploymentDTO) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


