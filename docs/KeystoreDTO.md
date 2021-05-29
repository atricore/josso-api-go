# KeystoreDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateAlias** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**KeystorePassOnly** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PrivateKeyName** | Pointer to **string** |  | [optional] 
**PrivateKeyPassword** | Pointer to **string** |  | [optional] 
**Store** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewKeystoreDTO

`func NewKeystoreDTO() *KeystoreDTO`

NewKeystoreDTO instantiates a new KeystoreDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoreDTOWithDefaults

`func NewKeystoreDTOWithDefaults() *KeystoreDTO`

NewKeystoreDTOWithDefaults instantiates a new KeystoreDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAlias

`func (o *KeystoreDTO) GetCertificateAlias() string`

GetCertificateAlias returns the CertificateAlias field if non-nil, zero value otherwise.

### GetCertificateAliasOk

`func (o *KeystoreDTO) GetCertificateAliasOk() (*string, bool)`

GetCertificateAliasOk returns a tuple with the CertificateAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAlias

`func (o *KeystoreDTO) SetCertificateAlias(v string)`

SetCertificateAlias sets CertificateAlias field to given value.

### HasCertificateAlias

`func (o *KeystoreDTO) HasCertificateAlias() bool`

HasCertificateAlias returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeystoreDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeystoreDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeystoreDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeystoreDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetElementId

`func (o *KeystoreDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *KeystoreDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *KeystoreDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *KeystoreDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *KeystoreDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeystoreDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeystoreDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeystoreDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeystorePassOnly

`func (o *KeystoreDTO) GetKeystorePassOnly() bool`

GetKeystorePassOnly returns the KeystorePassOnly field if non-nil, zero value otherwise.

### GetKeystorePassOnlyOk

`func (o *KeystoreDTO) GetKeystorePassOnlyOk() (*bool, bool)`

GetKeystorePassOnlyOk returns a tuple with the KeystorePassOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystorePassOnly

`func (o *KeystoreDTO) SetKeystorePassOnly(v bool)`

SetKeystorePassOnly sets KeystorePassOnly field to given value.

### HasKeystorePassOnly

`func (o *KeystoreDTO) HasKeystorePassOnly() bool`

HasKeystorePassOnly returns a boolean if a field has been set.

### GetName

`func (o *KeystoreDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeystoreDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeystoreDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeystoreDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *KeystoreDTO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeystoreDTO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeystoreDTO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KeystoreDTO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivateKeyName

`func (o *KeystoreDTO) GetPrivateKeyName() string`

GetPrivateKeyName returns the PrivateKeyName field if non-nil, zero value otherwise.

### GetPrivateKeyNameOk

`func (o *KeystoreDTO) GetPrivateKeyNameOk() (*string, bool)`

GetPrivateKeyNameOk returns a tuple with the PrivateKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyName

`func (o *KeystoreDTO) SetPrivateKeyName(v string)`

SetPrivateKeyName sets PrivateKeyName field to given value.

### HasPrivateKeyName

`func (o *KeystoreDTO) HasPrivateKeyName() bool`

HasPrivateKeyName returns a boolean if a field has been set.

### GetPrivateKeyPassword

`func (o *KeystoreDTO) GetPrivateKeyPassword() string`

GetPrivateKeyPassword returns the PrivateKeyPassword field if non-nil, zero value otherwise.

### GetPrivateKeyPasswordOk

`func (o *KeystoreDTO) GetPrivateKeyPasswordOk() (*string, bool)`

GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassword

`func (o *KeystoreDTO) SetPrivateKeyPassword(v string)`

SetPrivateKeyPassword sets PrivateKeyPassword field to given value.

### HasPrivateKeyPassword

`func (o *KeystoreDTO) HasPrivateKeyPassword() bool`

HasPrivateKeyPassword returns a boolean if a field has been set.

### GetStore

`func (o *KeystoreDTO) GetStore() ResourceDTO`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *KeystoreDTO) GetStoreOk() (*ResourceDTO, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *KeystoreDTO) SetStore(v ResourceDTO)`

SetStore sets Store field to given value.

### HasStore

`func (o *KeystoreDTO) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetType

`func (o *KeystoreDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeystoreDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeystoreDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeystoreDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


