# GetApplianceStateRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetApplianceStateRes

`func NewGetApplianceStateRes() *GetApplianceStateRes`

NewGetApplianceStateRes instantiates a new GetApplianceStateRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplianceStateResWithDefaults

`func NewGetApplianceStateResWithDefaults() *GetApplianceStateRes`

NewGetApplianceStateResWithDefaults instantiates a new GetApplianceStateRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetApplianceStateRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetApplianceStateRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetApplianceStateRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetApplianceStateRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetState

`func (o *GetApplianceStateRes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetApplianceStateRes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetApplianceStateRes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetApplianceStateRes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetApplianceStateRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetApplianceStateRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetApplianceStateRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetApplianceStateRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


