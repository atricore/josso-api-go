# CreateApplianceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appliance** | Pointer to [**IdentityApplianceDefinitionDTO**](IdentityApplianceDefinitionDTO.md) |  | [optional] 
**IdaName** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateApplianceReq

`func NewCreateApplianceReq() *CreateApplianceReq`

NewCreateApplianceReq instantiates a new CreateApplianceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplianceReqWithDefaults

`func NewCreateApplianceReqWithDefaults() *CreateApplianceReq`

NewCreateApplianceReqWithDefaults instantiates a new CreateApplianceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliance

`func (o *CreateApplianceReq) GetAppliance() IdentityApplianceDefinitionDTO`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *CreateApplianceReq) GetApplianceOk() (*IdentityApplianceDefinitionDTO, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *CreateApplianceReq) SetAppliance(v IdentityApplianceDefinitionDTO)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *CreateApplianceReq) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetIdaName

`func (o *CreateApplianceReq) GetIdaName() string`

GetIdaName returns the IdaName field if non-nil, zero value otherwise.

### GetIdaNameOk

`func (o *CreateApplianceReq) GetIdaNameOk() (*string, bool)`

GetIdaNameOk returns a tuple with the IdaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdaName

`func (o *CreateApplianceReq) SetIdaName(v string)`

SetIdaName sets IdaName field to given value.

### HasIdaName

`func (o *CreateApplianceReq) HasIdaName() bool`

HasIdaName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


