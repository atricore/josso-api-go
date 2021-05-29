# IdentityApplianceSecurityConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptSensitiveData** | Pointer to **bool** |  | [optional] 
**Encryption** | Pointer to **string** |  | [optional] 
**EncryptionConfig** | Pointer to **string** |  | [optional] 
**EncryptionConfigFile** | Pointer to **string** |  | [optional] 
**EncryptionPassword** | Pointer to **string** |  | [optional] 
**ExternalConfig** | Pointer to **bool** |  | [optional] 
**ExternalConfigFile** | Pointer to **string** |  | [optional] 
**PasswordProperty** | Pointer to **string** |  | [optional] 
**Salt** | Pointer to **string** |  | [optional] 
**SaltProperty** | Pointer to **string** |  | [optional] 
**SaltValue** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityApplianceSecurityConfigDTO

`func NewIdentityApplianceSecurityConfigDTO() *IdentityApplianceSecurityConfigDTO`

NewIdentityApplianceSecurityConfigDTO instantiates a new IdentityApplianceSecurityConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityApplianceSecurityConfigDTOWithDefaults

`func NewIdentityApplianceSecurityConfigDTOWithDefaults() *IdentityApplianceSecurityConfigDTO`

NewIdentityApplianceSecurityConfigDTOWithDefaults instantiates a new IdentityApplianceSecurityConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptSensitiveData

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptSensitiveData() bool`

GetEncryptSensitiveData returns the EncryptSensitiveData field if non-nil, zero value otherwise.

### GetEncryptSensitiveDataOk

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptSensitiveDataOk() (*bool, bool)`

GetEncryptSensitiveDataOk returns a tuple with the EncryptSensitiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptSensitiveData

`func (o *IdentityApplianceSecurityConfigDTO) SetEncryptSensitiveData(v bool)`

SetEncryptSensitiveData sets EncryptSensitiveData field to given value.

### HasEncryptSensitiveData

`func (o *IdentityApplianceSecurityConfigDTO) HasEncryptSensitiveData() bool`

HasEncryptSensitiveData returns a boolean if a field has been set.

### GetEncryption

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *IdentityApplianceSecurityConfigDTO) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *IdentityApplianceSecurityConfigDTO) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetEncryptionConfig

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptionConfig() string`

GetEncryptionConfig returns the EncryptionConfig field if non-nil, zero value otherwise.

### GetEncryptionConfigOk

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptionConfigOk() (*string, bool)`

GetEncryptionConfigOk returns a tuple with the EncryptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionConfig

`func (o *IdentityApplianceSecurityConfigDTO) SetEncryptionConfig(v string)`

SetEncryptionConfig sets EncryptionConfig field to given value.

### HasEncryptionConfig

`func (o *IdentityApplianceSecurityConfigDTO) HasEncryptionConfig() bool`

HasEncryptionConfig returns a boolean if a field has been set.

### GetEncryptionConfigFile

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptionConfigFile() string`

GetEncryptionConfigFile returns the EncryptionConfigFile field if non-nil, zero value otherwise.

### GetEncryptionConfigFileOk

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptionConfigFileOk() (*string, bool)`

GetEncryptionConfigFileOk returns a tuple with the EncryptionConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionConfigFile

`func (o *IdentityApplianceSecurityConfigDTO) SetEncryptionConfigFile(v string)`

SetEncryptionConfigFile sets EncryptionConfigFile field to given value.

### HasEncryptionConfigFile

`func (o *IdentityApplianceSecurityConfigDTO) HasEncryptionConfigFile() bool`

HasEncryptionConfigFile returns a boolean if a field has been set.

### GetEncryptionPassword

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptionPassword() string`

GetEncryptionPassword returns the EncryptionPassword field if non-nil, zero value otherwise.

### GetEncryptionPasswordOk

`func (o *IdentityApplianceSecurityConfigDTO) GetEncryptionPasswordOk() (*string, bool)`

GetEncryptionPasswordOk returns a tuple with the EncryptionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassword

`func (o *IdentityApplianceSecurityConfigDTO) SetEncryptionPassword(v string)`

SetEncryptionPassword sets EncryptionPassword field to given value.

### HasEncryptionPassword

`func (o *IdentityApplianceSecurityConfigDTO) HasEncryptionPassword() bool`

HasEncryptionPassword returns a boolean if a field has been set.

### GetExternalConfig

`func (o *IdentityApplianceSecurityConfigDTO) GetExternalConfig() bool`

GetExternalConfig returns the ExternalConfig field if non-nil, zero value otherwise.

### GetExternalConfigOk

`func (o *IdentityApplianceSecurityConfigDTO) GetExternalConfigOk() (*bool, bool)`

GetExternalConfigOk returns a tuple with the ExternalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfig

`func (o *IdentityApplianceSecurityConfigDTO) SetExternalConfig(v bool)`

SetExternalConfig sets ExternalConfig field to given value.

### HasExternalConfig

`func (o *IdentityApplianceSecurityConfigDTO) HasExternalConfig() bool`

HasExternalConfig returns a boolean if a field has been set.

### GetExternalConfigFile

`func (o *IdentityApplianceSecurityConfigDTO) GetExternalConfigFile() string`

GetExternalConfigFile returns the ExternalConfigFile field if non-nil, zero value otherwise.

### GetExternalConfigFileOk

`func (o *IdentityApplianceSecurityConfigDTO) GetExternalConfigFileOk() (*string, bool)`

GetExternalConfigFileOk returns a tuple with the ExternalConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfigFile

`func (o *IdentityApplianceSecurityConfigDTO) SetExternalConfigFile(v string)`

SetExternalConfigFile sets ExternalConfigFile field to given value.

### HasExternalConfigFile

`func (o *IdentityApplianceSecurityConfigDTO) HasExternalConfigFile() bool`

HasExternalConfigFile returns a boolean if a field has been set.

### GetPasswordProperty

`func (o *IdentityApplianceSecurityConfigDTO) GetPasswordProperty() string`

GetPasswordProperty returns the PasswordProperty field if non-nil, zero value otherwise.

### GetPasswordPropertyOk

`func (o *IdentityApplianceSecurityConfigDTO) GetPasswordPropertyOk() (*string, bool)`

GetPasswordPropertyOk returns a tuple with the PasswordProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProperty

`func (o *IdentityApplianceSecurityConfigDTO) SetPasswordProperty(v string)`

SetPasswordProperty sets PasswordProperty field to given value.

### HasPasswordProperty

`func (o *IdentityApplianceSecurityConfigDTO) HasPasswordProperty() bool`

HasPasswordProperty returns a boolean if a field has been set.

### GetSalt

`func (o *IdentityApplianceSecurityConfigDTO) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *IdentityApplianceSecurityConfigDTO) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *IdentityApplianceSecurityConfigDTO) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *IdentityApplianceSecurityConfigDTO) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetSaltProperty

`func (o *IdentityApplianceSecurityConfigDTO) GetSaltProperty() string`

GetSaltProperty returns the SaltProperty field if non-nil, zero value otherwise.

### GetSaltPropertyOk

`func (o *IdentityApplianceSecurityConfigDTO) GetSaltPropertyOk() (*string, bool)`

GetSaltPropertyOk returns a tuple with the SaltProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltProperty

`func (o *IdentityApplianceSecurityConfigDTO) SetSaltProperty(v string)`

SetSaltProperty sets SaltProperty field to given value.

### HasSaltProperty

`func (o *IdentityApplianceSecurityConfigDTO) HasSaltProperty() bool`

HasSaltProperty returns a boolean if a field has been set.

### GetSaltValue

`func (o *IdentityApplianceSecurityConfigDTO) GetSaltValue() string`

GetSaltValue returns the SaltValue field if non-nil, zero value otherwise.

### GetSaltValueOk

`func (o *IdentityApplianceSecurityConfigDTO) GetSaltValueOk() (*string, bool)`

GetSaltValueOk returns a tuple with the SaltValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltValue

`func (o *IdentityApplianceSecurityConfigDTO) SetSaltValue(v string)`

SetSaltValue sets SaltValue field to given value.

### HasSaltValue

`func (o *IdentityApplianceSecurityConfigDTO) HasSaltValue() bool`

HasSaltValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


