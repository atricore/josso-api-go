# ExecEnvContainerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Captive** | Pointer to **bool** |  | [optional] 
**ExecEnv** | Pointer to [**ExecutionEnvironmentDTO**](ExecutionEnvironmentDTO.md) |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewExecEnvContainerDTO

`func NewExecEnvContainerDTO() *ExecEnvContainerDTO`

NewExecEnvContainerDTO instantiates a new ExecEnvContainerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecEnvContainerDTOWithDefaults

`func NewExecEnvContainerDTOWithDefaults() *ExecEnvContainerDTO`

NewExecEnvContainerDTOWithDefaults instantiates a new ExecEnvContainerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptive

`func (o *ExecEnvContainerDTO) GetCaptive() bool`

GetCaptive returns the Captive field if non-nil, zero value otherwise.

### GetCaptiveOk

`func (o *ExecEnvContainerDTO) GetCaptiveOk() (*bool, bool)`

GetCaptiveOk returns a tuple with the Captive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptive

`func (o *ExecEnvContainerDTO) SetCaptive(v bool)`

SetCaptive sets Captive field to given value.

### HasCaptive

`func (o *ExecEnvContainerDTO) HasCaptive() bool`

HasCaptive returns a boolean if a field has been set.

### GetExecEnv

`func (o *ExecEnvContainerDTO) GetExecEnv() ExecutionEnvironmentDTO`

GetExecEnv returns the ExecEnv field if non-nil, zero value otherwise.

### GetExecEnvOk

`func (o *ExecEnvContainerDTO) GetExecEnvOk() (*ExecutionEnvironmentDTO, bool)`

GetExecEnvOk returns a tuple with the ExecEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnv

`func (o *ExecEnvContainerDTO) SetExecEnv(v ExecutionEnvironmentDTO)`

SetExecEnv sets ExecEnv field to given value.

### HasExecEnv

`func (o *ExecEnvContainerDTO) HasExecEnv() bool`

HasExecEnv returns a boolean if a field has been set.

### GetLocation

`func (o *ExecEnvContainerDTO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExecEnvContainerDTO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExecEnvContainerDTO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExecEnvContainerDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ExecEnvContainerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecEnvContainerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecEnvContainerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecEnvContainerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ExecEnvContainerDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExecEnvContainerDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExecEnvContainerDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExecEnvContainerDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


