# GetAppliancesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appliances** | Pointer to [**[]IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ValidationErrors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetAppliancesRes

`func NewGetAppliancesRes() *GetAppliancesRes`

NewGetAppliancesRes instantiates a new GetAppliancesRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppliancesResWithDefaults

`func NewGetAppliancesResWithDefaults() *GetAppliancesRes`

NewGetAppliancesResWithDefaults instantiates a new GetAppliancesRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliances

`func (o *GetAppliancesRes) GetAppliances() []IdentityApplianceDefinitionDTO`

GetAppliances returns the Appliances field if non-nil, zero value otherwise.

### GetAppliancesOk

`func (o *GetAppliancesRes) GetAppliancesOk() (*[]IdentityApplianceDefinitionDTO, bool)`

GetAppliancesOk returns a tuple with the Appliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliances

`func (o *GetAppliancesRes) SetAppliances(v []IdentityApplianceDefinitionDTO)`

SetAppliances sets Appliances field to given value.

### HasAppliances

`func (o *GetAppliancesRes) HasAppliances() bool`

HasAppliances returns a boolean if a field has been set.

### GetError

`func (o *GetAppliancesRes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAppliancesRes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAppliancesRes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetAppliancesRes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetValidationErrors

`func (o *GetAppliancesRes) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *GetAppliancesRes) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *GetAppliancesRes) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *GetAppliancesRes) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


