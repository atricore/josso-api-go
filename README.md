# Go API client for jossoappi

# Atricore Console API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.4.3-SNAPSHOT
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.atricore.com](https://www.atricore.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./jossoappi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8081/atricore-res/services*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateAppliance**](docs/DefaultApi.md#createappliance) | **Post** /iam-deploy/appliance | 
*DefaultApi* | [**CreateExtSaml2Sp**](docs/DefaultApi.md#createextsaml2sp) | **Post** /iam-deploy/extsaml2sp | 
*DefaultApi* | [**CreateIdP**](docs/DefaultApi.md#createidp) | **Post** /iam-deploy/idp | 
*DefaultApi* | [**CreateIdSourceDb**](docs/DefaultApi.md#createidsourcedb) | **Post** /iam-deploy/idsourcedb | 
*DefaultApi* | [**CreateIdSourceLdap**](docs/DefaultApi.md#createidsourceldap) | **Post** /iam-deploy/idsourceldap | 
*DefaultApi* | [**CreateIdVault**](docs/DefaultApi.md#createidvault) | **Post** /iam-deploy/idvault | 
*DefaultApi* | [**CreateIntSaml2Sp**](docs/DefaultApi.md#createintsaml2sp) | **Post** /iam-deploy/intsaml2sp | 
*DefaultApi* | [**CreateOidcRp**](docs/DefaultApi.md#createoidcrp) | **Post** /iam-deploy/oidcrp | 
*DefaultApi* | [**CreateVirtSaml2Sp**](docs/DefaultApi.md#createvirtsaml2sp) | **Post** /iam-deploy/virtsaml2sp | 
*DefaultApi* | [**DeleteAppliance**](docs/DefaultApi.md#deleteappliance) | **Delete** /iam-deploy/appliance | 
*DefaultApi* | [**DeleteExtSaml2Sp**](docs/DefaultApi.md#deleteextsaml2sp) | **Delete** /iam-deploy/extsaml2sp | 
*DefaultApi* | [**DeleteIdP**](docs/DefaultApi.md#deleteidp) | **Delete** /iam-deploy/idp | 
*DefaultApi* | [**DeleteIdSourceDb**](docs/DefaultApi.md#deleteidsourcedb) | **Delete** /iam-deploy/idsourcedb | 
*DefaultApi* | [**DeleteIdSourceLdap**](docs/DefaultApi.md#deleteidsourceldap) | **Delete** /iam-deploy/idsourceldap | 
*DefaultApi* | [**DeleteIdVault**](docs/DefaultApi.md#deleteidvault) | **Delete** /iam-deploy/idvault | 
*DefaultApi* | [**DeleteIntSaml2Sp**](docs/DefaultApi.md#deleteintsaml2sp) | **Delete** /iam-deploy/intsaml2sp | 
*DefaultApi* | [**DeleteOidcRp**](docs/DefaultApi.md#deleteoidcrp) | **Delete** /iam-deploy/oidcrp | 
*DefaultApi* | [**DeleteVirtSaml2Sp**](docs/DefaultApi.md#deletevirtsaml2sp) | **Delete** /iam-deploy/virtsaml2sp | 
*DefaultApi* | [**GetAppliance**](docs/DefaultApi.md#getappliance) | **Get** /iam-deploy/appliance | 
*DefaultApi* | [**GetAppliances**](docs/DefaultApi.md#getappliances) | **Get** /iam-deploy/appliances | 
*DefaultApi* | [**GetExtSaml2Sp**](docs/DefaultApi.md#getextsaml2sp) | **Get** /iam-deploy/extsaml2sp | 
*DefaultApi* | [**GetExtSaml2Sps**](docs/DefaultApi.md#getextsaml2sps) | **Get** /iam-deploy/extsaml2sps | 
*DefaultApi* | [**GetIdP**](docs/DefaultApi.md#getidp) | **Get** /iam-deploy/idp | 
*DefaultApi* | [**GetIdPs**](docs/DefaultApi.md#getidps) | **Get** /iam-deploy/idps | 
*DefaultApi* | [**GetIdSourceDb**](docs/DefaultApi.md#getidsourcedb) | **Get** /iam-deploy/idsourcedb | 
*DefaultApi* | [**GetIdSourceDbs**](docs/DefaultApi.md#getidsourcedbs) | **Get** /iam-deploy/idsourcedbs | 
*DefaultApi* | [**GetIdSourceLdap**](docs/DefaultApi.md#getidsourceldap) | **Get** /iam-deploy/idsourceldap | 
*DefaultApi* | [**GetIdSourceLdaps**](docs/DefaultApi.md#getidsourceldaps) | **Get** /iam-deploy/idsourceldaps | 
*DefaultApi* | [**GetIdVault**](docs/DefaultApi.md#getidvault) | **Get** /iam-deploy/idvault | 
*DefaultApi* | [**GetIdVaults**](docs/DefaultApi.md#getidvaults) | **Get** /iam-deploy/idvaults | 
*DefaultApi* | [**GetIntSaml2Sp**](docs/DefaultApi.md#getintsaml2sp) | **Get** /iam-deploy/intsaml2sp | 
*DefaultApi* | [**GetIntSaml2Sps**](docs/DefaultApi.md#getintsaml2sps) | **Get** /iam-deploy/intsaml2sps | 
*DefaultApi* | [**GetOidcRp**](docs/DefaultApi.md#getoidcrp) | **Get** /iam-deploy/oidcrp | 
*DefaultApi* | [**GetOidcRps**](docs/DefaultApi.md#getoidcrps) | **Get** /iam-deploy/oidcrps | 
*DefaultApi* | [**GetVirtSaml2Sp**](docs/DefaultApi.md#getvirtsaml2sp) | **Get** /iam-deploy/virtsaml2sp | 
*DefaultApi* | [**GetVirtSaml2Sps**](docs/DefaultApi.md#getvirtsaml2sps) | **Get** /iam-deploy/virtsaml2sps | 
*DefaultApi* | [**ImportAppliance**](docs/DefaultApi.md#importappliance) | **Post** /iam-deploy/appliance/import | 
*DefaultApi* | [**SignOn**](docs/DefaultApi.md#signon) | **Post** /iam-authn/sign-on | 
*DefaultApi* | [**UdpateExtSaml2Sp**](docs/DefaultApi.md#udpateextsaml2sp) | **Put** /iam-deploy/extsaml2sp | 
*DefaultApi* | [**UdpateIntSaml2Sp**](docs/DefaultApi.md#udpateintsaml2sp) | **Put** /iam-deploy/intsaml2sp | 
*DefaultApi* | [**UdpateVirtSaml2Sp**](docs/DefaultApi.md#udpatevirtsaml2sp) | **Put** /iam-deploy/virtsaml2sp | 
*DefaultApi* | [**UpdateAppliance**](docs/DefaultApi.md#updateappliance) | **Put** /iam-deploy/appliance | 
*DefaultApi* | [**UpdateIdP**](docs/DefaultApi.md#updateidp) | **Put** /iam-deploy/idp | 
*DefaultApi* | [**UpdateIdSourceDb**](docs/DefaultApi.md#updateidsourcedb) | **Put** /iam-deploy/idsourcedb | 
*DefaultApi* | [**UpdateIdSourceLdap**](docs/DefaultApi.md#updateidsourceldap) | **Put** /iam-deploy/idsourceldap | 
*DefaultApi* | [**UpdateIdVault**](docs/DefaultApi.md#updateidvault) | **Put** /iam-deploy/idvault | 
*DefaultApi* | [**UpdateOidcRp**](docs/DefaultApi.md#updateoidcrp) | **Put** /iam-deploy/oidcrp | 


## Documentation For Models

 - [AccountLinkagePolicyDTO](docs/AccountLinkagePolicyDTO.md)
 - [ActivationDTO](docs/ActivationDTO.md)
 - [AttributeProfileDTO](docs/AttributeProfileDTO.md)
 - [AttributeValueDTO](docs/AttributeValueDTO.md)
 - [AuthenticationAssertionEmissionPolicyDTO](docs/AuthenticationAssertionEmissionPolicyDTO.md)
 - [AuthenticationContractDTO](docs/AuthenticationContractDTO.md)
 - [AuthenticationMechanismDTO](docs/AuthenticationMechanismDTO.md)
 - [AuthenticationServiceDTO](docs/AuthenticationServiceDTO.md)
 - [CustomClassDTO](docs/CustomClassDTO.md)
 - [CustomClassPropertyDTO](docs/CustomClassPropertyDTO.md)
 - [DbIdentitySourceDTO](docs/DbIdentitySourceDTO.md)
 - [DelegatedAuthenticationDTO](docs/DelegatedAuthenticationDTO.md)
 - [DeleteReq](docs/DeleteReq.md)
 - [DeleteRes](docs/DeleteRes.md)
 - [EmbeddedIdentityVaultDTO](docs/EmbeddedIdentityVaultDTO.md)
 - [EntitySelectionStrategyDTO](docs/EntitySelectionStrategyDTO.md)
 - [ExecutionEnvironmentDTO](docs/ExecutionEnvironmentDTO.md)
 - [ExtensionDTO](docs/ExtensionDTO.md)
 - [ExternalOpenIDConnectRelayingPartyDTO](docs/ExternalOpenIDConnectRelayingPartyDTO.md)
 - [ExternalSaml2ServiceProviderDTO](docs/ExternalSaml2ServiceProviderDTO.md)
 - [FederatedChannelDTO](docs/FederatedChannelDTO.md)
 - [FederatedConnectionDTO](docs/FederatedConnectionDTO.md)
 - [FederatedProviderDTO](docs/FederatedProviderDTO.md)
 - [GetApplianceReq](docs/GetApplianceReq.md)
 - [GetApplianceRes](docs/GetApplianceRes.md)
 - [GetAppliancesRes](docs/GetAppliancesRes.md)
 - [GetExtSaml2SpReq](docs/GetExtSaml2SpReq.md)
 - [GetExtSaml2SpRes](docs/GetExtSaml2SpRes.md)
 - [GetExtSaml2SpsRes](docs/GetExtSaml2SpsRes.md)
 - [GetIdPReq](docs/GetIdPReq.md)
 - [GetIdPRes](docs/GetIdPRes.md)
 - [GetIdPsRes](docs/GetIdPsRes.md)
 - [GetIdSourceDbReq](docs/GetIdSourceDbReq.md)
 - [GetIdSourceDbRes](docs/GetIdSourceDbRes.md)
 - [GetIdSourceDbsRes](docs/GetIdSourceDbsRes.md)
 - [GetIdSourceLdapReq](docs/GetIdSourceLdapReq.md)
 - [GetIdSourceLdapRes](docs/GetIdSourceLdapRes.md)
 - [GetIdSourceLdapsRes](docs/GetIdSourceLdapsRes.md)
 - [GetIdVaultReq](docs/GetIdVaultReq.md)
 - [GetIdVaultRes](docs/GetIdVaultRes.md)
 - [GetIdVaultsRes](docs/GetIdVaultsRes.md)
 - [GetIntSaml2SpReq](docs/GetIntSaml2SpReq.md)
 - [GetIntSaml2SpRes](docs/GetIntSaml2SpRes.md)
 - [GetIntSaml2SpsRes](docs/GetIntSaml2SpsRes.md)
 - [GetOidcRpReq](docs/GetOidcRpReq.md)
 - [GetOidcRpRes](docs/GetOidcRpRes.md)
 - [GetOidcRpsRes](docs/GetOidcRpsRes.md)
 - [GetVirtSaml2SpReq](docs/GetVirtSaml2SpReq.md)
 - [GetVirtSaml2SpRes](docs/GetVirtSaml2SpRes.md)
 - [GetVirtSaml2SpsRes](docs/GetVirtSaml2SpsRes.md)
 - [GroupDTO](docs/GroupDTO.md)
 - [IdentityApplianceDefinitionDTO](docs/IdentityApplianceDefinitionDTO.md)
 - [IdentityApplianceSecurityConfigDTO](docs/IdentityApplianceSecurityConfigDTO.md)
 - [IdentityLookupDTO](docs/IdentityLookupDTO.md)
 - [IdentityMappingPolicyDTO](docs/IdentityMappingPolicyDTO.md)
 - [IdentityProviderDTO](docs/IdentityProviderDTO.md)
 - [IdentitySourceDTO](docs/IdentitySourceDTO.md)
 - [ImportApplianceReq](docs/ImportApplianceReq.md)
 - [ImportApplianceRes](docs/ImportApplianceRes.md)
 - [InternalSaml2ServiceProviderDTO](docs/InternalSaml2ServiceProviderDTO.md)
 - [KeystoreDTO](docs/KeystoreDTO.md)
 - [LdapIdentitySourceDTO](docs/LdapIdentitySourceDTO.md)
 - [LocationDTO](docs/LocationDTO.md)
 - [OAuth2ClientDTO](docs/OAuth2ClientDTO.md)
 - [OIDCSignOnRequest](docs/OIDCSignOnRequest.md)
 - [OIDCSignOnResponse](docs/OIDCSignOnResponse.md)
 - [PointDTO](docs/PointDTO.md)
 - [ProviderConfigDTO](docs/ProviderConfigDTO.md)
 - [ProviderDTO](docs/ProviderDTO.md)
 - [ResourceDTO](docs/ResourceDTO.md)
 - [SamlR2IDPConfigDTO](docs/SamlR2IDPConfigDTO.md)
 - [SamlR2SPConfigDTO](docs/SamlR2SPConfigDTO.md)
 - [ServerContext](docs/ServerContext.md)
 - [ServiceConnectionDTO](docs/ServiceConnectionDTO.md)
 - [ServiceResourceDTO](docs/ServiceResourceDTO.md)
 - [SessionManagerFactoryDTO](docs/SessionManagerFactoryDTO.md)
 - [StoreApplianceReq](docs/StoreApplianceReq.md)
 - [StoreApplianceRes](docs/StoreApplianceRes.md)
 - [StoreExtSaml2SpReq](docs/StoreExtSaml2SpReq.md)
 - [StoreExtSaml2SpRes](docs/StoreExtSaml2SpRes.md)
 - [StoreIdPReq](docs/StoreIdPReq.md)
 - [StoreIdPRes](docs/StoreIdPRes.md)
 - [StoreIdSourceDbReq](docs/StoreIdSourceDbReq.md)
 - [StoreIdSourceDbRes](docs/StoreIdSourceDbRes.md)
 - [StoreIdSourceLdapReq](docs/StoreIdSourceLdapReq.md)
 - [StoreIdSourceLdapRes](docs/StoreIdSourceLdapRes.md)
 - [StoreIdVaultReq](docs/StoreIdVaultReq.md)
 - [StoreIdVaultRes](docs/StoreIdVaultRes.md)
 - [StoreIntSaml2SpReq](docs/StoreIntSaml2SpReq.md)
 - [StoreIntSaml2SpRes](docs/StoreIntSaml2SpRes.md)
 - [StoreOidcRpReq](docs/StoreOidcRpReq.md)
 - [StoreOidcRpRes](docs/StoreOidcRpRes.md)
 - [StoreVirtSaml2SpReq](docs/StoreVirtSaml2SpReq.md)
 - [StoreVirtSaml2SpRes](docs/StoreVirtSaml2SpRes.md)
 - [SubjectAuthenticationPolicyDTO](docs/SubjectAuthenticationPolicyDTO.md)
 - [SubjectNameIdentifierPolicyDTO](docs/SubjectNameIdentifierPolicyDTO.md)
 - [UserDTO](docs/UserDTO.md)
 - [UserDashboardBrandingDTO](docs/UserDashboardBrandingDTO.md)
 - [VirtualSaml2ServiceProviderDTO](docs/VirtualSaml2ServiceProviderDTO.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

sgonzalez@atricore.com

