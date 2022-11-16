# ActivateExecEnvReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivateSamples** | Pointer to **bool** |  | [optional] 
**ApplianceId** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**ExecEnvFolder** | Pointer to **string** |  | [optional] 
**ExecEnvName** | Pointer to **string** |  | [optional] 
**IdOrName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Reactivate** | Pointer to **bool** |  | [optional] 
**Replace** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewActivateExecEnvReq

`func NewActivateExecEnvReq() *ActivateExecEnvReq`

NewActivateExecEnvReq instantiates a new ActivateExecEnvReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateExecEnvReqWithDefaults

`func NewActivateExecEnvReqWithDefaults() *ActivateExecEnvReq`

NewActivateExecEnvReqWithDefaults instantiates a new ActivateExecEnvReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivateSamples

`func (o *ActivateExecEnvReq) GetActivateSamples() bool`

GetActivateSamples returns the ActivateSamples field if non-nil, zero value otherwise.

### GetActivateSamplesOk

`func (o *ActivateExecEnvReq) GetActivateSamplesOk() (*bool, bool)`

GetActivateSamplesOk returns a tuple with the ActivateSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateSamples

`func (o *ActivateExecEnvReq) SetActivateSamples(v bool)`

SetActivateSamples sets ActivateSamples field to given value.

### HasActivateSamples

`func (o *ActivateExecEnvReq) HasActivateSamples() bool`

HasActivateSamples returns a boolean if a field has been set.

### GetApplianceId

`func (o *ActivateExecEnvReq) GetApplianceId() string`

GetApplianceId returns the ApplianceId field if non-nil, zero value otherwise.

### GetApplianceIdOk

`func (o *ActivateExecEnvReq) GetApplianceIdOk() (*string, bool)`

GetApplianceIdOk returns a tuple with the ApplianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceId

`func (o *ActivateExecEnvReq) SetApplianceId(v string)`

SetApplianceId sets ApplianceId field to given value.

### HasApplianceId

`func (o *ActivateExecEnvReq) HasApplianceId() bool`

HasApplianceId returns a boolean if a field has been set.

### GetDestination

`func (o *ActivateExecEnvReq) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ActivateExecEnvReq) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ActivateExecEnvReq) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ActivateExecEnvReq) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetExecEnvFolder

`func (o *ActivateExecEnvReq) GetExecEnvFolder() string`

GetExecEnvFolder returns the ExecEnvFolder field if non-nil, zero value otherwise.

### GetExecEnvFolderOk

`func (o *ActivateExecEnvReq) GetExecEnvFolderOk() (*string, bool)`

GetExecEnvFolderOk returns a tuple with the ExecEnvFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnvFolder

`func (o *ActivateExecEnvReq) SetExecEnvFolder(v string)`

SetExecEnvFolder sets ExecEnvFolder field to given value.

### HasExecEnvFolder

`func (o *ActivateExecEnvReq) HasExecEnvFolder() bool`

HasExecEnvFolder returns a boolean if a field has been set.

### GetExecEnvName

`func (o *ActivateExecEnvReq) GetExecEnvName() string`

GetExecEnvName returns the ExecEnvName field if non-nil, zero value otherwise.

### GetExecEnvNameOk

`func (o *ActivateExecEnvReq) GetExecEnvNameOk() (*string, bool)`

GetExecEnvNameOk returns a tuple with the ExecEnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnvName

`func (o *ActivateExecEnvReq) SetExecEnvName(v string)`

SetExecEnvName sets ExecEnvName field to given value.

### HasExecEnvName

`func (o *ActivateExecEnvReq) HasExecEnvName() bool`

HasExecEnvName returns a boolean if a field has been set.

### GetIdOrName

`func (o *ActivateExecEnvReq) GetIdOrName() string`

GetIdOrName returns the IdOrName field if non-nil, zero value otherwise.

### GetIdOrNameOk

`func (o *ActivateExecEnvReq) GetIdOrNameOk() (*string, bool)`

GetIdOrNameOk returns a tuple with the IdOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdOrName

`func (o *ActivateExecEnvReq) SetIdOrName(v string)`

SetIdOrName sets IdOrName field to given value.

### HasIdOrName

`func (o *ActivateExecEnvReq) HasIdOrName() bool`

HasIdOrName returns a boolean if a field has been set.

### GetPassword

`func (o *ActivateExecEnvReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ActivateExecEnvReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ActivateExecEnvReq) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ActivateExecEnvReq) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetReactivate

`func (o *ActivateExecEnvReq) GetReactivate() bool`

GetReactivate returns the Reactivate field if non-nil, zero value otherwise.

### GetReactivateOk

`func (o *ActivateExecEnvReq) GetReactivateOk() (*bool, bool)`

GetReactivateOk returns a tuple with the Reactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivate

`func (o *ActivateExecEnvReq) SetReactivate(v bool)`

SetReactivate sets Reactivate field to given value.

### HasReactivate

`func (o *ActivateExecEnvReq) HasReactivate() bool`

HasReactivate returns a boolean if a field has been set.

### GetReplace

`func (o *ActivateExecEnvReq) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *ActivateExecEnvReq) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *ActivateExecEnvReq) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *ActivateExecEnvReq) HasReplace() bool`

HasReplace returns a boolean if a field has been set.

### GetUsername

`func (o *ActivateExecEnvReq) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActivateExecEnvReq) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActivateExecEnvReq) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ActivateExecEnvReq) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


