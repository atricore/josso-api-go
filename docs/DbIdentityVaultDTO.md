# DbIdentityVaultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquireIncrement** | Pointer to **int32** |  | [optional] 
**ConnectionUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DriverName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ExternalDB** | Pointer to **bool** |  | [optional] 
**HashAlgorithm** | Pointer to **string** |  | [optional] 
**HashEncoding** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdleConnectionTestPeriod** | Pointer to **int32** |  | [optional] 
**InitialPoolSize** | Pointer to **int32** |  | [optional] 
**MaxIdleTime** | Pointer to **int32** |  | [optional] 
**MaxPoolSize** | Pointer to **int32** |  | [optional] 
**MinPoolSize** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PooledDatasource** | Pointer to **bool** |  | [optional] 
**SaltLength** | Pointer to **int32** |  | [optional] 
**SaltValue** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewDbIdentityVaultDTO

`func NewDbIdentityVaultDTO() *DbIdentityVaultDTO`

NewDbIdentityVaultDTO instantiates a new DbIdentityVaultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbIdentityVaultDTOWithDefaults

`func NewDbIdentityVaultDTOWithDefaults() *DbIdentityVaultDTO`

NewDbIdentityVaultDTOWithDefaults instantiates a new DbIdentityVaultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquireIncrement

`func (o *DbIdentityVaultDTO) GetAcquireIncrement() int32`

GetAcquireIncrement returns the AcquireIncrement field if non-nil, zero value otherwise.

### GetAcquireIncrementOk

`func (o *DbIdentityVaultDTO) GetAcquireIncrementOk() (*int32, bool)`

GetAcquireIncrementOk returns a tuple with the AcquireIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquireIncrement

`func (o *DbIdentityVaultDTO) SetAcquireIncrement(v int32)`

SetAcquireIncrement sets AcquireIncrement field to given value.

### HasAcquireIncrement

`func (o *DbIdentityVaultDTO) HasAcquireIncrement() bool`

HasAcquireIncrement returns a boolean if a field has been set.

### GetConnectionUrl

`func (o *DbIdentityVaultDTO) GetConnectionUrl() string`

GetConnectionUrl returns the ConnectionUrl field if non-nil, zero value otherwise.

### GetConnectionUrlOk

`func (o *DbIdentityVaultDTO) GetConnectionUrlOk() (*string, bool)`

GetConnectionUrlOk returns a tuple with the ConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrl

`func (o *DbIdentityVaultDTO) SetConnectionUrl(v string)`

SetConnectionUrl sets ConnectionUrl field to given value.

### HasConnectionUrl

`func (o *DbIdentityVaultDTO) HasConnectionUrl() bool`

HasConnectionUrl returns a boolean if a field has been set.

### GetDescription

`func (o *DbIdentityVaultDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DbIdentityVaultDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DbIdentityVaultDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DbIdentityVaultDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDriverName

`func (o *DbIdentityVaultDTO) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *DbIdentityVaultDTO) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *DbIdentityVaultDTO) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *DbIdentityVaultDTO) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetElementId

`func (o *DbIdentityVaultDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *DbIdentityVaultDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *DbIdentityVaultDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *DbIdentityVaultDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetExternalDB

`func (o *DbIdentityVaultDTO) GetExternalDB() bool`

GetExternalDB returns the ExternalDB field if non-nil, zero value otherwise.

### GetExternalDBOk

`func (o *DbIdentityVaultDTO) GetExternalDBOk() (*bool, bool)`

GetExternalDBOk returns a tuple with the ExternalDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDB

`func (o *DbIdentityVaultDTO) SetExternalDB(v bool)`

SetExternalDB sets ExternalDB field to given value.

### HasExternalDB

`func (o *DbIdentityVaultDTO) HasExternalDB() bool`

HasExternalDB returns a boolean if a field has been set.

### GetHashAlgorithm

`func (o *DbIdentityVaultDTO) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *DbIdentityVaultDTO) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *DbIdentityVaultDTO) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *DbIdentityVaultDTO) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### GetHashEncoding

`func (o *DbIdentityVaultDTO) GetHashEncoding() string`

GetHashEncoding returns the HashEncoding field if non-nil, zero value otherwise.

### GetHashEncodingOk

`func (o *DbIdentityVaultDTO) GetHashEncodingOk() (*string, bool)`

GetHashEncodingOk returns a tuple with the HashEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashEncoding

`func (o *DbIdentityVaultDTO) SetHashEncoding(v string)`

SetHashEncoding sets HashEncoding field to given value.

### HasHashEncoding

`func (o *DbIdentityVaultDTO) HasHashEncoding() bool`

HasHashEncoding returns a boolean if a field has been set.

### GetId

`func (o *DbIdentityVaultDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbIdentityVaultDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbIdentityVaultDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DbIdentityVaultDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdleConnectionTestPeriod

`func (o *DbIdentityVaultDTO) GetIdleConnectionTestPeriod() int32`

GetIdleConnectionTestPeriod returns the IdleConnectionTestPeriod field if non-nil, zero value otherwise.

### GetIdleConnectionTestPeriodOk

`func (o *DbIdentityVaultDTO) GetIdleConnectionTestPeriodOk() (*int32, bool)`

GetIdleConnectionTestPeriodOk returns a tuple with the IdleConnectionTestPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleConnectionTestPeriod

`func (o *DbIdentityVaultDTO) SetIdleConnectionTestPeriod(v int32)`

SetIdleConnectionTestPeriod sets IdleConnectionTestPeriod field to given value.

### HasIdleConnectionTestPeriod

`func (o *DbIdentityVaultDTO) HasIdleConnectionTestPeriod() bool`

HasIdleConnectionTestPeriod returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *DbIdentityVaultDTO) GetInitialPoolSize() int32`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *DbIdentityVaultDTO) GetInitialPoolSizeOk() (*int32, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *DbIdentityVaultDTO) SetInitialPoolSize(v int32)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *DbIdentityVaultDTO) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxIdleTime

`func (o *DbIdentityVaultDTO) GetMaxIdleTime() int32`

GetMaxIdleTime returns the MaxIdleTime field if non-nil, zero value otherwise.

### GetMaxIdleTimeOk

`func (o *DbIdentityVaultDTO) GetMaxIdleTimeOk() (*int32, bool)`

GetMaxIdleTimeOk returns a tuple with the MaxIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdleTime

`func (o *DbIdentityVaultDTO) SetMaxIdleTime(v int32)`

SetMaxIdleTime sets MaxIdleTime field to given value.

### HasMaxIdleTime

`func (o *DbIdentityVaultDTO) HasMaxIdleTime() bool`

HasMaxIdleTime returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *DbIdentityVaultDTO) GetMaxPoolSize() int32`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *DbIdentityVaultDTO) GetMaxPoolSizeOk() (*int32, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *DbIdentityVaultDTO) SetMaxPoolSize(v int32)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *DbIdentityVaultDTO) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetMinPoolSize

`func (o *DbIdentityVaultDTO) GetMinPoolSize() int32`

GetMinPoolSize returns the MinPoolSize field if non-nil, zero value otherwise.

### GetMinPoolSizeOk

`func (o *DbIdentityVaultDTO) GetMinPoolSizeOk() (*int32, bool)`

GetMinPoolSizeOk returns a tuple with the MinPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoolSize

`func (o *DbIdentityVaultDTO) SetMinPoolSize(v int32)`

SetMinPoolSize sets MinPoolSize field to given value.

### HasMinPoolSize

`func (o *DbIdentityVaultDTO) HasMinPoolSize() bool`

HasMinPoolSize returns a boolean if a field has been set.

### GetName

`func (o *DbIdentityVaultDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbIdentityVaultDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbIdentityVaultDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbIdentityVaultDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *DbIdentityVaultDTO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DbIdentityVaultDTO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DbIdentityVaultDTO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DbIdentityVaultDTO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPooledDatasource

`func (o *DbIdentityVaultDTO) GetPooledDatasource() bool`

GetPooledDatasource returns the PooledDatasource field if non-nil, zero value otherwise.

### GetPooledDatasourceOk

`func (o *DbIdentityVaultDTO) GetPooledDatasourceOk() (*bool, bool)`

GetPooledDatasourceOk returns a tuple with the PooledDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooledDatasource

`func (o *DbIdentityVaultDTO) SetPooledDatasource(v bool)`

SetPooledDatasource sets PooledDatasource field to given value.

### HasPooledDatasource

`func (o *DbIdentityVaultDTO) HasPooledDatasource() bool`

HasPooledDatasource returns a boolean if a field has been set.

### GetSaltLength

`func (o *DbIdentityVaultDTO) GetSaltLength() int32`

GetSaltLength returns the SaltLength field if non-nil, zero value otherwise.

### GetSaltLengthOk

`func (o *DbIdentityVaultDTO) GetSaltLengthOk() (*int32, bool)`

GetSaltLengthOk returns a tuple with the SaltLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLength

`func (o *DbIdentityVaultDTO) SetSaltLength(v int32)`

SetSaltLength sets SaltLength field to given value.

### HasSaltLength

`func (o *DbIdentityVaultDTO) HasSaltLength() bool`

HasSaltLength returns a boolean if a field has been set.

### GetSaltValue

`func (o *DbIdentityVaultDTO) GetSaltValue() string`

GetSaltValue returns the SaltValue field if non-nil, zero value otherwise.

### GetSaltValueOk

`func (o *DbIdentityVaultDTO) GetSaltValueOk() (*string, bool)`

GetSaltValueOk returns a tuple with the SaltValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltValue

`func (o *DbIdentityVaultDTO) SetSaltValue(v string)`

SetSaltValue sets SaltValue field to given value.

### HasSaltValue

`func (o *DbIdentityVaultDTO) HasSaltValue() bool`

HasSaltValue returns a boolean if a field has been set.

### GetUsername

`func (o *DbIdentityVaultDTO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DbIdentityVaultDTO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DbIdentityVaultDTO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DbIdentityVaultDTO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetX

`func (o *DbIdentityVaultDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *DbIdentityVaultDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *DbIdentityVaultDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *DbIdentityVaultDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *DbIdentityVaultDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *DbIdentityVaultDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *DbIdentityVaultDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *DbIdentityVaultDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


