# GetApplianceContainersRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appliances** | Pointer to [**[]IdentityApplianceDTO**](IdentityApplianceDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetApplianceContainersRes

`func NewGetApplianceContainersRes() *GetApplianceContainersRes`

NewGetApplianceContainersRes instantiates a new GetApplianceContainersRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplianceContainersResWithDefaults

`func NewGetApplianceContainersResWithDefaults() *GetApplianceContainersRes`

NewGetApplianceContainersResWithDefaults instantiates a new GetApplianceContainersRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliances

`func (o *GetApplianceContainersRes) GetAppliances() []IdentityApplianceDTO`

GetAppliances returns the Appliances field if non-nil, zero value otherwise.

### GetAppliancesOk

`func (o *GetApplianceContainersRes) GetAppliancesOk() (*[]IdentityApplianceDTO, bool)`

GetAppliancesOk returns a tuple with the Appliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliances

`func (o *GetApplianceContainersRes) SetAppliances(v []IdentityApplianceDTO)`

SetAppliances sets Appliances field to given value.

### HasAppliances

`func (o *GetApplianceContainersRes) HasAppliances() bool`

HasAppliances returns a boolean if a field has been set.

### GetError

`func (o *GetApplianceContainersRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetApplianceContainersRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetApplianceContainersRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetApplianceContainersRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetApplianceContainersRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetApplianceContainersRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetApplianceContainersRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetApplianceContainersRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


