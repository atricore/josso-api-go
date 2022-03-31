# GetApplianceContainerRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appliance** | Pointer to [**IdentityApplianceDTO**](IdentityApplianceDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetApplianceContainerRes

`func NewGetApplianceContainerRes() *GetApplianceContainerRes`

NewGetApplianceContainerRes instantiates a new GetApplianceContainerRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplianceContainerResWithDefaults

`func NewGetApplianceContainerResWithDefaults() *GetApplianceContainerRes`

NewGetApplianceContainerResWithDefaults instantiates a new GetApplianceContainerRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliance

`func (o *GetApplianceContainerRes) GetAppliance() IdentityApplianceDTO`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *GetApplianceContainerRes) GetApplianceOk() (*IdentityApplianceDTO, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *GetApplianceContainerRes) SetAppliance(v IdentityApplianceDTO)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *GetApplianceContainerRes) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetError

`func (o *GetApplianceContainerRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetApplianceContainerRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetApplianceContainerRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetApplianceContainerRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetApplianceContainerRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetApplianceContainerRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetApplianceContainerRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetApplianceContainerRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


