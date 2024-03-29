/*
Atricore Console :: Remote : API

# Atricore Console API

API version: 1.5.1-SNAPSHOT
Contact: sgonzalez@atricore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// DbIdentityVaultDTO struct for DbIdentityVaultDTO
type DbIdentityVaultDTO struct {
	AcquireIncrement *int32 `json:"acquireIncrement,omitempty"`
	ConnectionUrl *string `json:"connectionUrl,omitempty"`
	Description *string `json:"description,omitempty"`
	DriverName *string `json:"driverName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	ExternalDB *bool `json:"externalDB,omitempty"`
	HashAlgorithm *string `json:"hashAlgorithm,omitempty"`
	HashEncoding *string `json:"hashEncoding,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IdleConnectionTestPeriod *int32 `json:"idleConnectionTestPeriod,omitempty"`
	InitialPoolSize *int32 `json:"initialPoolSize,omitempty"`
	MaxIdleTime *int32 `json:"maxIdleTime,omitempty"`
	MaxPoolSize *int32 `json:"maxPoolSize,omitempty"`
	MinPoolSize *int32 `json:"minPoolSize,omitempty"`
	Name *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	PooledDatasource *bool `json:"pooledDatasource,omitempty"`
	SaltLength *int32 `json:"saltLength,omitempty"`
	SaltValue *string `json:"saltValue,omitempty"`
	Username *string `json:"username,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DbIdentityVaultDTO DbIdentityVaultDTO

// NewDbIdentityVaultDTO instantiates a new DbIdentityVaultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbIdentityVaultDTO() *DbIdentityVaultDTO {
	this := DbIdentityVaultDTO{}
	return &this
}

// NewDbIdentityVaultDTOWithDefaults instantiates a new DbIdentityVaultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbIdentityVaultDTOWithDefaults() *DbIdentityVaultDTO {
	this := DbIdentityVaultDTO{}
	return &this
}

// GetAcquireIncrement returns the AcquireIncrement field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetAcquireIncrement() int32 {
	if o == nil || isNil(o.AcquireIncrement) {
		var ret int32
		return ret
	}
	return *o.AcquireIncrement
}

// GetAcquireIncrementOk returns a tuple with the AcquireIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetAcquireIncrementOk() (*int32, bool) {
	if o == nil || isNil(o.AcquireIncrement) {
    return nil, false
	}
	return o.AcquireIncrement, true
}

// HasAcquireIncrement returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasAcquireIncrement() bool {
	if o != nil && !isNil(o.AcquireIncrement) {
		return true
	}

	return false
}

// SetAcquireIncrement gets a reference to the given int32 and assigns it to the AcquireIncrement field.
func (o *DbIdentityVaultDTO) SetAcquireIncrement(v int32) {
	o.AcquireIncrement = &v
}

// GetConnectionUrl returns the ConnectionUrl field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetConnectionUrl() string {
	if o == nil || isNil(o.ConnectionUrl) {
		var ret string
		return ret
	}
	return *o.ConnectionUrl
}

// GetConnectionUrlOk returns a tuple with the ConnectionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetConnectionUrlOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionUrl) {
    return nil, false
	}
	return o.ConnectionUrl, true
}

// HasConnectionUrl returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasConnectionUrl() bool {
	if o != nil && !isNil(o.ConnectionUrl) {
		return true
	}

	return false
}

// SetConnectionUrl gets a reference to the given string and assigns it to the ConnectionUrl field.
func (o *DbIdentityVaultDTO) SetConnectionUrl(v string) {
	o.ConnectionUrl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DbIdentityVaultDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDriverName returns the DriverName field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetDriverName() string {
	if o == nil || isNil(o.DriverName) {
		var ret string
		return ret
	}
	return *o.DriverName
}

// GetDriverNameOk returns a tuple with the DriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetDriverNameOk() (*string, bool) {
	if o == nil || isNil(o.DriverName) {
    return nil, false
	}
	return o.DriverName, true
}

// HasDriverName returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasDriverName() bool {
	if o != nil && !isNil(o.DriverName) {
		return true
	}

	return false
}

// SetDriverName gets a reference to the given string and assigns it to the DriverName field.
func (o *DbIdentityVaultDTO) SetDriverName(v string) {
	o.DriverName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetElementId() string {
	if o == nil || isNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetElementIdOk() (*string, bool) {
	if o == nil || isNil(o.ElementId) {
    return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasElementId() bool {
	if o != nil && !isNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *DbIdentityVaultDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetExternalDB returns the ExternalDB field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetExternalDB() bool {
	if o == nil || isNil(o.ExternalDB) {
		var ret bool
		return ret
	}
	return *o.ExternalDB
}

// GetExternalDBOk returns a tuple with the ExternalDB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetExternalDBOk() (*bool, bool) {
	if o == nil || isNil(o.ExternalDB) {
    return nil, false
	}
	return o.ExternalDB, true
}

// HasExternalDB returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasExternalDB() bool {
	if o != nil && !isNil(o.ExternalDB) {
		return true
	}

	return false
}

// SetExternalDB gets a reference to the given bool and assigns it to the ExternalDB field.
func (o *DbIdentityVaultDTO) SetExternalDB(v bool) {
	o.ExternalDB = &v
}

// GetHashAlgorithm returns the HashAlgorithm field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetHashAlgorithm() string {
	if o == nil || isNil(o.HashAlgorithm) {
		var ret string
		return ret
	}
	return *o.HashAlgorithm
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetHashAlgorithmOk() (*string, bool) {
	if o == nil || isNil(o.HashAlgorithm) {
    return nil, false
	}
	return o.HashAlgorithm, true
}

// HasHashAlgorithm returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasHashAlgorithm() bool {
	if o != nil && !isNil(o.HashAlgorithm) {
		return true
	}

	return false
}

// SetHashAlgorithm gets a reference to the given string and assigns it to the HashAlgorithm field.
func (o *DbIdentityVaultDTO) SetHashAlgorithm(v string) {
	o.HashAlgorithm = &v
}

// GetHashEncoding returns the HashEncoding field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetHashEncoding() string {
	if o == nil || isNil(o.HashEncoding) {
		var ret string
		return ret
	}
	return *o.HashEncoding
}

// GetHashEncodingOk returns a tuple with the HashEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetHashEncodingOk() (*string, bool) {
	if o == nil || isNil(o.HashEncoding) {
    return nil, false
	}
	return o.HashEncoding, true
}

// HasHashEncoding returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasHashEncoding() bool {
	if o != nil && !isNil(o.HashEncoding) {
		return true
	}

	return false
}

// SetHashEncoding gets a reference to the given string and assigns it to the HashEncoding field.
func (o *DbIdentityVaultDTO) SetHashEncoding(v string) {
	o.HashEncoding = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DbIdentityVaultDTO) SetId(v int64) {
	o.Id = &v
}

// GetIdleConnectionTestPeriod returns the IdleConnectionTestPeriod field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetIdleConnectionTestPeriod() int32 {
	if o == nil || isNil(o.IdleConnectionTestPeriod) {
		var ret int32
		return ret
	}
	return *o.IdleConnectionTestPeriod
}

// GetIdleConnectionTestPeriodOk returns a tuple with the IdleConnectionTestPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetIdleConnectionTestPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.IdleConnectionTestPeriod) {
    return nil, false
	}
	return o.IdleConnectionTestPeriod, true
}

// HasIdleConnectionTestPeriod returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasIdleConnectionTestPeriod() bool {
	if o != nil && !isNil(o.IdleConnectionTestPeriod) {
		return true
	}

	return false
}

// SetIdleConnectionTestPeriod gets a reference to the given int32 and assigns it to the IdleConnectionTestPeriod field.
func (o *DbIdentityVaultDTO) SetIdleConnectionTestPeriod(v int32) {
	o.IdleConnectionTestPeriod = &v
}

// GetInitialPoolSize returns the InitialPoolSize field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetInitialPoolSize() int32 {
	if o == nil || isNil(o.InitialPoolSize) {
		var ret int32
		return ret
	}
	return *o.InitialPoolSize
}

// GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetInitialPoolSizeOk() (*int32, bool) {
	if o == nil || isNil(o.InitialPoolSize) {
    return nil, false
	}
	return o.InitialPoolSize, true
}

// HasInitialPoolSize returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasInitialPoolSize() bool {
	if o != nil && !isNil(o.InitialPoolSize) {
		return true
	}

	return false
}

// SetInitialPoolSize gets a reference to the given int32 and assigns it to the InitialPoolSize field.
func (o *DbIdentityVaultDTO) SetInitialPoolSize(v int32) {
	o.InitialPoolSize = &v
}

// GetMaxIdleTime returns the MaxIdleTime field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetMaxIdleTime() int32 {
	if o == nil || isNil(o.MaxIdleTime) {
		var ret int32
		return ret
	}
	return *o.MaxIdleTime
}

// GetMaxIdleTimeOk returns a tuple with the MaxIdleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetMaxIdleTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxIdleTime) {
    return nil, false
	}
	return o.MaxIdleTime, true
}

// HasMaxIdleTime returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasMaxIdleTime() bool {
	if o != nil && !isNil(o.MaxIdleTime) {
		return true
	}

	return false
}

// SetMaxIdleTime gets a reference to the given int32 and assigns it to the MaxIdleTime field.
func (o *DbIdentityVaultDTO) SetMaxIdleTime(v int32) {
	o.MaxIdleTime = &v
}

// GetMaxPoolSize returns the MaxPoolSize field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetMaxPoolSize() int32 {
	if o == nil || isNil(o.MaxPoolSize) {
		var ret int32
		return ret
	}
	return *o.MaxPoolSize
}

// GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetMaxPoolSizeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPoolSize) {
    return nil, false
	}
	return o.MaxPoolSize, true
}

// HasMaxPoolSize returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasMaxPoolSize() bool {
	if o != nil && !isNil(o.MaxPoolSize) {
		return true
	}

	return false
}

// SetMaxPoolSize gets a reference to the given int32 and assigns it to the MaxPoolSize field.
func (o *DbIdentityVaultDTO) SetMaxPoolSize(v int32) {
	o.MaxPoolSize = &v
}

// GetMinPoolSize returns the MinPoolSize field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetMinPoolSize() int32 {
	if o == nil || isNil(o.MinPoolSize) {
		var ret int32
		return ret
	}
	return *o.MinPoolSize
}

// GetMinPoolSizeOk returns a tuple with the MinPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetMinPoolSizeOk() (*int32, bool) {
	if o == nil || isNil(o.MinPoolSize) {
    return nil, false
	}
	return o.MinPoolSize, true
}

// HasMinPoolSize returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasMinPoolSize() bool {
	if o != nil && !isNil(o.MinPoolSize) {
		return true
	}

	return false
}

// SetMinPoolSize gets a reference to the given int32 and assigns it to the MinPoolSize field.
func (o *DbIdentityVaultDTO) SetMinPoolSize(v int32) {
	o.MinPoolSize = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DbIdentityVaultDTO) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DbIdentityVaultDTO) SetPassword(v string) {
	o.Password = &v
}

// GetPooledDatasource returns the PooledDatasource field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetPooledDatasource() bool {
	if o == nil || isNil(o.PooledDatasource) {
		var ret bool
		return ret
	}
	return *o.PooledDatasource
}

// GetPooledDatasourceOk returns a tuple with the PooledDatasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetPooledDatasourceOk() (*bool, bool) {
	if o == nil || isNil(o.PooledDatasource) {
    return nil, false
	}
	return o.PooledDatasource, true
}

// HasPooledDatasource returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasPooledDatasource() bool {
	if o != nil && !isNil(o.PooledDatasource) {
		return true
	}

	return false
}

// SetPooledDatasource gets a reference to the given bool and assigns it to the PooledDatasource field.
func (o *DbIdentityVaultDTO) SetPooledDatasource(v bool) {
	o.PooledDatasource = &v
}

// GetSaltLength returns the SaltLength field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetSaltLength() int32 {
	if o == nil || isNil(o.SaltLength) {
		var ret int32
		return ret
	}
	return *o.SaltLength
}

// GetSaltLengthOk returns a tuple with the SaltLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetSaltLengthOk() (*int32, bool) {
	if o == nil || isNil(o.SaltLength) {
    return nil, false
	}
	return o.SaltLength, true
}

// HasSaltLength returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasSaltLength() bool {
	if o != nil && !isNil(o.SaltLength) {
		return true
	}

	return false
}

// SetSaltLength gets a reference to the given int32 and assigns it to the SaltLength field.
func (o *DbIdentityVaultDTO) SetSaltLength(v int32) {
	o.SaltLength = &v
}

// GetSaltValue returns the SaltValue field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetSaltValue() string {
	if o == nil || isNil(o.SaltValue) {
		var ret string
		return ret
	}
	return *o.SaltValue
}

// GetSaltValueOk returns a tuple with the SaltValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetSaltValueOk() (*string, bool) {
	if o == nil || isNil(o.SaltValue) {
    return nil, false
	}
	return o.SaltValue, true
}

// HasSaltValue returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasSaltValue() bool {
	if o != nil && !isNil(o.SaltValue) {
		return true
	}

	return false
}

// SetSaltValue gets a reference to the given string and assigns it to the SaltValue field.
func (o *DbIdentityVaultDTO) SetSaltValue(v string) {
	o.SaltValue = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DbIdentityVaultDTO) SetUsername(v string) {
	o.Username = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetX() float64 {
	if o == nil || isNil(o.X) {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetXOk() (*float64, bool) {
	if o == nil || isNil(o.X) {
    return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasX() bool {
	if o != nil && !isNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *DbIdentityVaultDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *DbIdentityVaultDTO) GetY() float64 {
	if o == nil || isNil(o.Y) {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbIdentityVaultDTO) GetYOk() (*float64, bool) {
	if o == nil || isNil(o.Y) {
    return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *DbIdentityVaultDTO) HasY() bool {
	if o != nil && !isNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *DbIdentityVaultDTO) SetY(v float64) {
	o.Y = &v
}

func (o DbIdentityVaultDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AcquireIncrement) {
		toSerialize["acquireIncrement"] = o.AcquireIncrement
	}
	if !isNil(o.ConnectionUrl) {
		toSerialize["connectionUrl"] = o.ConnectionUrl
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DriverName) {
		toSerialize["driverName"] = o.DriverName
	}
	if !isNil(o.ElementId) {
		toSerialize["elementId"] = o.ElementId
	}
	if !isNil(o.ExternalDB) {
		toSerialize["externalDB"] = o.ExternalDB
	}
	if !isNil(o.HashAlgorithm) {
		toSerialize["hashAlgorithm"] = o.HashAlgorithm
	}
	if !isNil(o.HashEncoding) {
		toSerialize["hashEncoding"] = o.HashEncoding
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IdleConnectionTestPeriod) {
		toSerialize["idleConnectionTestPeriod"] = o.IdleConnectionTestPeriod
	}
	if !isNil(o.InitialPoolSize) {
		toSerialize["initialPoolSize"] = o.InitialPoolSize
	}
	if !isNil(o.MaxIdleTime) {
		toSerialize["maxIdleTime"] = o.MaxIdleTime
	}
	if !isNil(o.MaxPoolSize) {
		toSerialize["maxPoolSize"] = o.MaxPoolSize
	}
	if !isNil(o.MinPoolSize) {
		toSerialize["minPoolSize"] = o.MinPoolSize
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.PooledDatasource) {
		toSerialize["pooledDatasource"] = o.PooledDatasource
	}
	if !isNil(o.SaltLength) {
		toSerialize["saltLength"] = o.SaltLength
	}
	if !isNil(o.SaltValue) {
		toSerialize["saltValue"] = o.SaltValue
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !isNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DbIdentityVaultDTO) UnmarshalJSON(bytes []byte) (err error) {
	varDbIdentityVaultDTO := _DbIdentityVaultDTO{}

	if err = json.Unmarshal(bytes, &varDbIdentityVaultDTO); err == nil {
		*o = DbIdentityVaultDTO(varDbIdentityVaultDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "acquireIncrement")
		delete(additionalProperties, "connectionUrl")
		delete(additionalProperties, "description")
		delete(additionalProperties, "driverName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "externalDB")
		delete(additionalProperties, "hashAlgorithm")
		delete(additionalProperties, "hashEncoding")
		delete(additionalProperties, "id")
		delete(additionalProperties, "idleConnectionTestPeriod")
		delete(additionalProperties, "initialPoolSize")
		delete(additionalProperties, "maxIdleTime")
		delete(additionalProperties, "maxPoolSize")
		delete(additionalProperties, "minPoolSize")
		delete(additionalProperties, "name")
		delete(additionalProperties, "password")
		delete(additionalProperties, "pooledDatasource")
		delete(additionalProperties, "saltLength")
		delete(additionalProperties, "saltValue")
		delete(additionalProperties, "username")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDbIdentityVaultDTO struct {
	value *DbIdentityVaultDTO
	isSet bool
}

func (v NullableDbIdentityVaultDTO) Get() *DbIdentityVaultDTO {
	return v.value
}

func (v *NullableDbIdentityVaultDTO) Set(val *DbIdentityVaultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDbIdentityVaultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDbIdentityVaultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbIdentityVaultDTO(val *DbIdentityVaultDTO) *NullableDbIdentityVaultDTO {
	return &NullableDbIdentityVaultDTO{value: val, isSet: true}
}

func (v NullableDbIdentityVaultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbIdentityVaultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


