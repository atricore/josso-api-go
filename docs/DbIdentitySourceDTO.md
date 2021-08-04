# DbIdentitySourceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquireIncrement** | Pointer to **int32** |  | [optional] 
**Admin** | Pointer to **string** |  | [optional] 
**ConnectionUrl** | Pointer to **string** |  | [optional] 
**CredentialsQueryString** | Pointer to **string** |  | [optional] 
**CustomClass** | Pointer to [**CustomClassDTO**](CustomClassDTO.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Driver** | Pointer to [**ResourceDTO**](ResourceDTO.md) |  | [optional] 
**DriverName** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IdleConnectionTestPeriod** | Pointer to **int32** |  | [optional] 
**InitialPoolSize** | Pointer to **int32** |  | [optional] 
**MaxIdleTime** | Pointer to **int32** |  | [optional] 
**MaxPoolSize** | Pointer to **int32** |  | [optional] 
**MinPoolSize** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PooledDatasource** | Pointer to **bool** |  | [optional] 
**RelayCredentialQueryString** | Pointer to **string** |  | [optional] 
**ResetCredentialDml** | Pointer to **string** |  | [optional] 
**RolesQueryString** | Pointer to **string** |  | [optional] 
**UseColumnNamesAsPropertyNames** | Pointer to **bool** |  | [optional] 
**UserPropertiesQueryString** | Pointer to **string** |  | [optional] 
**UserQueryString** | Pointer to **string** |  | [optional] 
**X** | Pointer to **float64** |  | [optional] 
**Y** | Pointer to **float64** |  | [optional] 

## Methods

### NewDbIdentitySourceDTO

`func NewDbIdentitySourceDTO() *DbIdentitySourceDTO`

NewDbIdentitySourceDTO instantiates a new DbIdentitySourceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbIdentitySourceDTOWithDefaults

`func NewDbIdentitySourceDTOWithDefaults() *DbIdentitySourceDTO`

NewDbIdentitySourceDTOWithDefaults instantiates a new DbIdentitySourceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquireIncrement

`func (o *DbIdentitySourceDTO) GetAcquireIncrement() int32`

GetAcquireIncrement returns the AcquireIncrement field if non-nil, zero value otherwise.

### GetAcquireIncrementOk

`func (o *DbIdentitySourceDTO) GetAcquireIncrementOk() (*int32, bool)`

GetAcquireIncrementOk returns a tuple with the AcquireIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquireIncrement

`func (o *DbIdentitySourceDTO) SetAcquireIncrement(v int32)`

SetAcquireIncrement sets AcquireIncrement field to given value.

### HasAcquireIncrement

`func (o *DbIdentitySourceDTO) HasAcquireIncrement() bool`

HasAcquireIncrement returns a boolean if a field has been set.

### GetAdmin

`func (o *DbIdentitySourceDTO) GetAdmin() string`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *DbIdentitySourceDTO) GetAdminOk() (*string, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *DbIdentitySourceDTO) SetAdmin(v string)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *DbIdentitySourceDTO) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetConnectionUrl

`func (o *DbIdentitySourceDTO) GetConnectionUrl() string`

GetConnectionUrl returns the ConnectionUrl field if non-nil, zero value otherwise.

### GetConnectionUrlOk

`func (o *DbIdentitySourceDTO) GetConnectionUrlOk() (*string, bool)`

GetConnectionUrlOk returns a tuple with the ConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrl

`func (o *DbIdentitySourceDTO) SetConnectionUrl(v string)`

SetConnectionUrl sets ConnectionUrl field to given value.

### HasConnectionUrl

`func (o *DbIdentitySourceDTO) HasConnectionUrl() bool`

HasConnectionUrl returns a boolean if a field has been set.

### GetCredentialsQueryString

`func (o *DbIdentitySourceDTO) GetCredentialsQueryString() string`

GetCredentialsQueryString returns the CredentialsQueryString field if non-nil, zero value otherwise.

### GetCredentialsQueryStringOk

`func (o *DbIdentitySourceDTO) GetCredentialsQueryStringOk() (*string, bool)`

GetCredentialsQueryStringOk returns a tuple with the CredentialsQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsQueryString

`func (o *DbIdentitySourceDTO) SetCredentialsQueryString(v string)`

SetCredentialsQueryString sets CredentialsQueryString field to given value.

### HasCredentialsQueryString

`func (o *DbIdentitySourceDTO) HasCredentialsQueryString() bool`

HasCredentialsQueryString returns a boolean if a field has been set.

### GetCustomClass

`func (o *DbIdentitySourceDTO) GetCustomClass() CustomClassDTO`

GetCustomClass returns the CustomClass field if non-nil, zero value otherwise.

### GetCustomClassOk

`func (o *DbIdentitySourceDTO) GetCustomClassOk() (*CustomClassDTO, bool)`

GetCustomClassOk returns a tuple with the CustomClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClass

`func (o *DbIdentitySourceDTO) SetCustomClass(v CustomClassDTO)`

SetCustomClass sets CustomClass field to given value.

### HasCustomClass

`func (o *DbIdentitySourceDTO) HasCustomClass() bool`

HasCustomClass returns a boolean if a field has been set.

### GetDescription

`func (o *DbIdentitySourceDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DbIdentitySourceDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DbIdentitySourceDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DbIdentitySourceDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDriver

`func (o *DbIdentitySourceDTO) GetDriver() ResourceDTO`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *DbIdentitySourceDTO) GetDriverOk() (*ResourceDTO, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *DbIdentitySourceDTO) SetDriver(v ResourceDTO)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *DbIdentitySourceDTO) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetDriverName

`func (o *DbIdentitySourceDTO) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *DbIdentitySourceDTO) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *DbIdentitySourceDTO) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *DbIdentitySourceDTO) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetElementId

`func (o *DbIdentitySourceDTO) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *DbIdentitySourceDTO) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *DbIdentitySourceDTO) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *DbIdentitySourceDTO) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *DbIdentitySourceDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbIdentitySourceDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbIdentitySourceDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DbIdentitySourceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdleConnectionTestPeriod

`func (o *DbIdentitySourceDTO) GetIdleConnectionTestPeriod() int32`

GetIdleConnectionTestPeriod returns the IdleConnectionTestPeriod field if non-nil, zero value otherwise.

### GetIdleConnectionTestPeriodOk

`func (o *DbIdentitySourceDTO) GetIdleConnectionTestPeriodOk() (*int32, bool)`

GetIdleConnectionTestPeriodOk returns a tuple with the IdleConnectionTestPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleConnectionTestPeriod

`func (o *DbIdentitySourceDTO) SetIdleConnectionTestPeriod(v int32)`

SetIdleConnectionTestPeriod sets IdleConnectionTestPeriod field to given value.

### HasIdleConnectionTestPeriod

`func (o *DbIdentitySourceDTO) HasIdleConnectionTestPeriod() bool`

HasIdleConnectionTestPeriod returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *DbIdentitySourceDTO) GetInitialPoolSize() int32`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *DbIdentitySourceDTO) GetInitialPoolSizeOk() (*int32, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *DbIdentitySourceDTO) SetInitialPoolSize(v int32)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *DbIdentitySourceDTO) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxIdleTime

`func (o *DbIdentitySourceDTO) GetMaxIdleTime() int32`

GetMaxIdleTime returns the MaxIdleTime field if non-nil, zero value otherwise.

### GetMaxIdleTimeOk

`func (o *DbIdentitySourceDTO) GetMaxIdleTimeOk() (*int32, bool)`

GetMaxIdleTimeOk returns a tuple with the MaxIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdleTime

`func (o *DbIdentitySourceDTO) SetMaxIdleTime(v int32)`

SetMaxIdleTime sets MaxIdleTime field to given value.

### HasMaxIdleTime

`func (o *DbIdentitySourceDTO) HasMaxIdleTime() bool`

HasMaxIdleTime returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *DbIdentitySourceDTO) GetMaxPoolSize() int32`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *DbIdentitySourceDTO) GetMaxPoolSizeOk() (*int32, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *DbIdentitySourceDTO) SetMaxPoolSize(v int32)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *DbIdentitySourceDTO) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetMinPoolSize

`func (o *DbIdentitySourceDTO) GetMinPoolSize() int32`

GetMinPoolSize returns the MinPoolSize field if non-nil, zero value otherwise.

### GetMinPoolSizeOk

`func (o *DbIdentitySourceDTO) GetMinPoolSizeOk() (*int32, bool)`

GetMinPoolSizeOk returns a tuple with the MinPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoolSize

`func (o *DbIdentitySourceDTO) SetMinPoolSize(v int32)`

SetMinPoolSize sets MinPoolSize field to given value.

### HasMinPoolSize

`func (o *DbIdentitySourceDTO) HasMinPoolSize() bool`

HasMinPoolSize returns a boolean if a field has been set.

### GetName

`func (o *DbIdentitySourceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbIdentitySourceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbIdentitySourceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbIdentitySourceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *DbIdentitySourceDTO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DbIdentitySourceDTO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DbIdentitySourceDTO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DbIdentitySourceDTO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPooledDatasource

`func (o *DbIdentitySourceDTO) GetPooledDatasource() bool`

GetPooledDatasource returns the PooledDatasource field if non-nil, zero value otherwise.

### GetPooledDatasourceOk

`func (o *DbIdentitySourceDTO) GetPooledDatasourceOk() (*bool, bool)`

GetPooledDatasourceOk returns a tuple with the PooledDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooledDatasource

`func (o *DbIdentitySourceDTO) SetPooledDatasource(v bool)`

SetPooledDatasource sets PooledDatasource field to given value.

### HasPooledDatasource

`func (o *DbIdentitySourceDTO) HasPooledDatasource() bool`

HasPooledDatasource returns a boolean if a field has been set.

### GetRelayCredentialQueryString

`func (o *DbIdentitySourceDTO) GetRelayCredentialQueryString() string`

GetRelayCredentialQueryString returns the RelayCredentialQueryString field if non-nil, zero value otherwise.

### GetRelayCredentialQueryStringOk

`func (o *DbIdentitySourceDTO) GetRelayCredentialQueryStringOk() (*string, bool)`

GetRelayCredentialQueryStringOk returns a tuple with the RelayCredentialQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayCredentialQueryString

`func (o *DbIdentitySourceDTO) SetRelayCredentialQueryString(v string)`

SetRelayCredentialQueryString sets RelayCredentialQueryString field to given value.

### HasRelayCredentialQueryString

`func (o *DbIdentitySourceDTO) HasRelayCredentialQueryString() bool`

HasRelayCredentialQueryString returns a boolean if a field has been set.

### GetResetCredentialDml

`func (o *DbIdentitySourceDTO) GetResetCredentialDml() string`

GetResetCredentialDml returns the ResetCredentialDml field if non-nil, zero value otherwise.

### GetResetCredentialDmlOk

`func (o *DbIdentitySourceDTO) GetResetCredentialDmlOk() (*string, bool)`

GetResetCredentialDmlOk returns a tuple with the ResetCredentialDml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCredentialDml

`func (o *DbIdentitySourceDTO) SetResetCredentialDml(v string)`

SetResetCredentialDml sets ResetCredentialDml field to given value.

### HasResetCredentialDml

`func (o *DbIdentitySourceDTO) HasResetCredentialDml() bool`

HasResetCredentialDml returns a boolean if a field has been set.

### GetRolesQueryString

`func (o *DbIdentitySourceDTO) GetRolesQueryString() string`

GetRolesQueryString returns the RolesQueryString field if non-nil, zero value otherwise.

### GetRolesQueryStringOk

`func (o *DbIdentitySourceDTO) GetRolesQueryStringOk() (*string, bool)`

GetRolesQueryStringOk returns a tuple with the RolesQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesQueryString

`func (o *DbIdentitySourceDTO) SetRolesQueryString(v string)`

SetRolesQueryString sets RolesQueryString field to given value.

### HasRolesQueryString

`func (o *DbIdentitySourceDTO) HasRolesQueryString() bool`

HasRolesQueryString returns a boolean if a field has been set.

### GetUseColumnNamesAsPropertyNames

`func (o *DbIdentitySourceDTO) GetUseColumnNamesAsPropertyNames() bool`

GetUseColumnNamesAsPropertyNames returns the UseColumnNamesAsPropertyNames field if non-nil, zero value otherwise.

### GetUseColumnNamesAsPropertyNamesOk

`func (o *DbIdentitySourceDTO) GetUseColumnNamesAsPropertyNamesOk() (*bool, bool)`

GetUseColumnNamesAsPropertyNamesOk returns a tuple with the UseColumnNamesAsPropertyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseColumnNamesAsPropertyNames

`func (o *DbIdentitySourceDTO) SetUseColumnNamesAsPropertyNames(v bool)`

SetUseColumnNamesAsPropertyNames sets UseColumnNamesAsPropertyNames field to given value.

### HasUseColumnNamesAsPropertyNames

`func (o *DbIdentitySourceDTO) HasUseColumnNamesAsPropertyNames() bool`

HasUseColumnNamesAsPropertyNames returns a boolean if a field has been set.

### GetUserPropertiesQueryString

`func (o *DbIdentitySourceDTO) GetUserPropertiesQueryString() string`

GetUserPropertiesQueryString returns the UserPropertiesQueryString field if non-nil, zero value otherwise.

### GetUserPropertiesQueryStringOk

`func (o *DbIdentitySourceDTO) GetUserPropertiesQueryStringOk() (*string, bool)`

GetUserPropertiesQueryStringOk returns a tuple with the UserPropertiesQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertiesQueryString

`func (o *DbIdentitySourceDTO) SetUserPropertiesQueryString(v string)`

SetUserPropertiesQueryString sets UserPropertiesQueryString field to given value.

### HasUserPropertiesQueryString

`func (o *DbIdentitySourceDTO) HasUserPropertiesQueryString() bool`

HasUserPropertiesQueryString returns a boolean if a field has been set.

### GetUserQueryString

`func (o *DbIdentitySourceDTO) GetUserQueryString() string`

GetUserQueryString returns the UserQueryString field if non-nil, zero value otherwise.

### GetUserQueryStringOk

`func (o *DbIdentitySourceDTO) GetUserQueryStringOk() (*string, bool)`

GetUserQueryStringOk returns a tuple with the UserQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQueryString

`func (o *DbIdentitySourceDTO) SetUserQueryString(v string)`

SetUserQueryString sets UserQueryString field to given value.

### HasUserQueryString

`func (o *DbIdentitySourceDTO) HasUserQueryString() bool`

HasUserQueryString returns a boolean if a field has been set.

### GetX

`func (o *DbIdentitySourceDTO) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *DbIdentitySourceDTO) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *DbIdentitySourceDTO) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *DbIdentitySourceDTO) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *DbIdentitySourceDTO) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *DbIdentitySourceDTO) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *DbIdentitySourceDTO) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *DbIdentitySourceDTO) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


