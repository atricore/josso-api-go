# ClientCertAuthnServiceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClrEnabled** | Pointer to **bool** |  | [optional] 
**CrlRefreshSeconds** | Pointer to **int32** |  | [optional] 
**CrlUrl** | Pointer to **string** |  | [optional] 
**CustomClass** | Pointer to [**CustomClassDTO**](CustomClassDTO.md) |  | [optional] 
**DelegatedAuthentications** | Pointer to [**[]DelegatedAuthenticationDTO**](DelegatedAuthenticationDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OcspEnabled** | Pointer to **bool** |  | [optional] 
**OcspServer** | Pointer to **string** |  | [optional] 
**Ocspserver** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewClientCertAuthnServiceDTO

`func NewClientCertAuthnServiceDTO() *ClientCertAuthnServiceDTO`

NewClientCertAuthnServiceDTO instantiates a new ClientCertAuthnServiceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertAuthnServiceDTOWithDefaults

`func NewClientCertAuthnServiceDTOWithDefaults() *ClientCertAuthnServiceDTO`

NewClientCertAuthnServiceDTOWithDefaults instantiates a new ClientCertAuthnServiceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClrEnabled

`func (o *ClientCertAuthnServiceDTO) GetClrEnabled() bool`

GetClrEnabled returns the ClrEnabled field if non-nil, zero value otherwise.

### GetClrEnabledOk

`func (o *ClientCertAuthnServiceDTO) GetClrEnabledOk() (*bool, bool)`

GetClrEnabledOk returns a tuple with the ClrEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClrEnabled

`func (o *ClientCertAuthnServiceDTO) SetClrEnabled(v bool)`

SetClrEnabled sets ClrEnabled field to given value.

### HasClrEnabled

`func (o *ClientCertAuthnServiceDTO) HasClrEnabled() bool`

HasClrEnabled returns a boolean if a field has been set.

### GetCrlRefreshSeconds

`func (o *ClientCertAuthnServiceDTO) GetCrlRefreshSeconds() int32`

GetCrlRefreshSeconds returns the CrlRefreshSeconds field if non-nil, zero value otherwise.

### GetCrlRefreshSecondsOk

`func (o *ClientCertAuthnServiceDTO) GetCrlRefreshSecondsOk() (*int32, bool)`

GetCrlRefreshSecondsOk returns a tuple with the CrlRefreshSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlRefreshSeconds

`func (o *ClientCertAuthnServiceDTO) SetCrlRefreshSeconds(v int32)`

SetCrlRefreshSeconds sets CrlRefreshSeconds field to given value.

### HasCrlRefreshSeconds

`func (o *ClientCertAuthnServiceDTO) HasCrlRefreshSeconds() bool`

HasCrlRefreshSeconds returns a boolean if a field has been set.

### GetCrlUrl

`func (o *ClientCertAuthnServiceDTO) GetCrlUrl() string`

GetCrlUrl returns the CrlUrl field if non-nil, zero value otherwise.

### GetCrlUrlOk

`func (o *ClientCertAuthnServiceDTO) GetCrlUrlOk() (*string, bool)`

GetCrlUrlOk returns a tuple with the CrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlUrl

`func (o *ClientCertAuthnServiceDTO) SetCrlUrl(v string)`

SetCrlUrl sets CrlUrl field to given value.

### HasCrlUrl

`func (o *ClientCertAuthnServiceDTO) HasCrlUrl() bool`

HasCrlUrl returns a boolean if a field has been set.

### GetCustomClass

`func (o *ClientCertAuthnServiceDTO) GetCustomClass() CustomClassDTO`

GetCustomClass returns the CustomClass field if non-nil, zero value otherwise.

### GetCustomClassOk

`func (o *ClientCertAuthnServiceDTO) GetCustomClassOk() (*CustomClassDTO, bool)`

GetCustomClassOk returns a tuple with the CustomClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClass

`func (o *ClientCertAuthnServiceDTO) SetCustomClass(v CustomClassDTO)`

SetCustomClass sets CustomClass field to given value.

### HasCustomClass

`func (o *ClientCertAuthnServiceDTO) HasCustomClass() bool`

HasCustomClass returns a boolean if a field has been set.

### GetDelegatedAuthentications

`func (o *ClientCertAuthnServiceDTO) GetDelegatedAuthentications() []DelegatedAuthenticationDTO`

GetDelegatedAuthentications returns the DelegatedAuthentications field if non-nil, zero value otherwise.

### GetDelegatedAuthenticationsOk

`func (o *ClientCertAuthnServiceDTO) GetDelegatedAuthenticationsOk() (*[]DelegatedAuthenticationDTO, bool)`

GetDelegatedAuthenticationsOk returns a tuple with the DelegatedAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAuthentications

`func (o *ClientCertAuthnServiceDTO) SetDelegatedAuthentications(v []DelegatedAuthenticationDTO)`

SetDelegatedAuthentications sets DelegatedAuthentications field to given value.

### HasDelegatedAuthentications

`func (o *ClientCertAuthnServiceDTO) HasDelegatedAuthentications() bool`

HasDelegatedAuthentications returns a boolean if a field has been set.

### GetDescription

`func (o *ClientCertAuthnServiceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientCertAuthnServiceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientCertAuthnServiceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientCertAuthnServiceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ClientCertAuthnServiceDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClientCertAuthnServiceDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClientCertAuthnServiceDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ClientCertAuthnServiceDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *ClientCertAuthnServiceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *ClientCertAuthnServiceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *ClientCertAuthnServiceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *ClientCertAuthnServiceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *ClientCertAuthnServiceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientCertAuthnServiceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientCertAuthnServiceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClientCertAuthnServiceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClientCertAuthnServiceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientCertAuthnServiceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientCertAuthnServiceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientCertAuthnServiceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOcspEnabled

`func (o *ClientCertAuthnServiceDTO) GetOcspEnabled() bool`

GetOcspEnabled returns the OcspEnabled field if non-nil, zero value otherwise.

### GetOcspEnabledOk

`func (o *ClientCertAuthnServiceDTO) GetOcspEnabledOk() (*bool, bool)`

GetOcspEnabledOk returns a tuple with the OcspEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspEnabled

`func (o *ClientCertAuthnServiceDTO) SetOcspEnabled(v bool)`

SetOcspEnabled sets OcspEnabled field to given value.

### HasOcspEnabled

`func (o *ClientCertAuthnServiceDTO) HasOcspEnabled() bool`

HasOcspEnabled returns a boolean if a field has been set.

### GetOcspServer

`func (o *ClientCertAuthnServiceDTO) GetOcspServer() string`

GetOcspServer returns the OcspServer field if non-nil, zero value otherwise.

### GetOcspServerOk

`func (o *ClientCertAuthnServiceDTO) GetOcspServerOk() (*string, bool)`

GetOcspServerOk returns a tuple with the OcspServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServer

`func (o *ClientCertAuthnServiceDTO) SetOcspServer(v string)`

SetOcspServer sets OcspServer field to given value.

### HasOcspServer

`func (o *ClientCertAuthnServiceDTO) HasOcspServer() bool`

HasOcspServer returns a boolean if a field has been set.

### GetOcspserver

`func (o *ClientCertAuthnServiceDTO) GetOcspserver() string`

GetOcspserver returns the Ocspserver field if non-nil, zero value otherwise.

### GetOcspserverOk

`func (o *ClientCertAuthnServiceDTO) GetOcspserverOk() (*string, bool)`

GetOcspserverOk returns a tuple with the Ocspserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspserver

`func (o *ClientCertAuthnServiceDTO) SetOcspserver(v string)`

SetOcspserver sets Ocspserver field to given value.

### HasOcspserver

`func (o *ClientCertAuthnServiceDTO) HasOcspserver() bool`

HasOcspserver returns a boolean if a field has been set.

### GetUid

`func (o *ClientCertAuthnServiceDTO) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ClientCertAuthnServiceDTO) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ClientCertAuthnServiceDTO) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ClientCertAuthnServiceDTO) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetX

`func (o *ClientCertAuthnServiceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ClientCertAuthnServiceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ClientCertAuthnServiceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *ClientCertAuthnServiceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ClientCertAuthnServiceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ClientCertAuthnServiceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ClientCertAuthnServiceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *ClientCertAuthnServiceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


