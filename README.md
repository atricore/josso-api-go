# Go API client for jossoappi

# Atricore Console API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.5.0-SNAPSHOT
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
import jossoappi "github.com/atricore/josso-api-go"
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
ctx := context.WithValue(context.Background(), jossoappi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), jossoappi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), jossoappi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), jossoappi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8081/atricore-res/services*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ActivateExecEnv**](docs/DefaultApi.md#activateexecenv) | **Get** /iam-deploy/execenv/activate | 
*DefaultApi* | [**BuildAppliance**](docs/DefaultApi.md#buildappliance) | **Get** /iam-deploy/appliance/build | 
*DefaultApi* | [**CreateAppliance**](docs/DefaultApi.md#createappliance) | **Post** /iam-deploy/appliance | 
*DefaultApi* | [**CreateBranding**](docs/DefaultApi.md#createbranding) | **Post** /iam-branding/branding | 
*DefaultApi* | [**CreateDbIdVault**](docs/DefaultApi.md#createdbidvault) | **Post** /iam-deploy/dbidvault | 
*DefaultApi* | [**CreateExtSaml2Sp**](docs/DefaultApi.md#createextsaml2sp) | **Post** /iam-deploy/extsaml2sp | 
*DefaultApi* | [**CreateIdP**](docs/DefaultApi.md#createidp) | **Post** /iam-deploy/idp | 
*DefaultApi* | [**CreateIdSourceDb**](docs/DefaultApi.md#createidsourcedb) | **Post** /iam-deploy/idsourcedb | 
*DefaultApi* | [**CreateIdSourceLdap**](docs/DefaultApi.md#createidsourceldap) | **Post** /iam-deploy/idsourceldap | 
*DefaultApi* | [**CreateIdVault**](docs/DefaultApi.md#createidvault) | **Post** /iam-deploy/idvault | 
*DefaultApi* | [**CreateIdpAzure**](docs/DefaultApi.md#createidpazure) | **Post** /iam-deploy/idp_azure | 
*DefaultApi* | [**CreateIdpFacebook**](docs/DefaultApi.md#createidpfacebook) | **Post** /iam-deploy/idp_fb | 
*DefaultApi* | [**CreateIdpGoogle**](docs/DefaultApi.md#createidpgoogle) | **Post** /iam-deploy/idp_google | 
*DefaultApi* | [**CreateIisExecEnv**](docs/DefaultApi.md#createiisexecenv) | **Post** /iam-deploy/iisexecenv | 
*DefaultApi* | [**CreateIntSaml2Sp**](docs/DefaultApi.md#createintsaml2sp) | **Post** /iam-deploy/intsaml2sp | 
*DefaultApi* | [**CreateJossoRs**](docs/DefaultApi.md#createjossors) | **Post** /iam-deploy/jossors | 
*DefaultApi* | [**CreateOidcRp**](docs/DefaultApi.md#createoidcrp) | **Post** /iam-deploy/oidcrp | 
*DefaultApi* | [**CreatePhpExecEnv**](docs/DefaultApi.md#createphpexecenv) | **Post** /iam-deploy/phpexecenv | 
*DefaultApi* | [**CreateSharepointRs**](docs/DefaultApi.md#createsharepointrs) | **Post** /iam-deploy/sharepointrs | 
*DefaultApi* | [**CreateTomcatExecEnv**](docs/DefaultApi.md#createtomcatexecenv) | **Post** /iam-deploy/tomcatexecenv | 
*DefaultApi* | [**CreateVirtSaml2Sp**](docs/DefaultApi.md#createvirtsaml2sp) | **Post** /iam-deploy/virtsaml2sp | 
*DefaultApi* | [**CreateWeblogicExecEnv**](docs/DefaultApi.md#createweblogicexecenv) | **Post** /iam-deploy/weblogicexecenv | 
*DefaultApi* | [**DeleteAppliance**](docs/DefaultApi.md#deleteappliance) | **Delete** /iam-deploy/appliance | 
*DefaultApi* | [**DeleteBranding**](docs/DefaultApi.md#deletebranding) | **Delete** /iam-branding/branding | 
*DefaultApi* | [**DeleteDbIdVault**](docs/DefaultApi.md#deletedbidvault) | **Delete** /iam-deploy/dbidvault | 
*DefaultApi* | [**DeleteExtSaml2Sp**](docs/DefaultApi.md#deleteextsaml2sp) | **Delete** /iam-deploy/extsaml2sp | 
*DefaultApi* | [**DeleteIdP**](docs/DefaultApi.md#deleteidp) | **Delete** /iam-deploy/idp | 
*DefaultApi* | [**DeleteIdSourceDb**](docs/DefaultApi.md#deleteidsourcedb) | **Delete** /iam-deploy/idsourcedb | 
*DefaultApi* | [**DeleteIdSourceLdap**](docs/DefaultApi.md#deleteidsourceldap) | **Delete** /iam-deploy/idsourceldap | 
*DefaultApi* | [**DeleteIdVault**](docs/DefaultApi.md#deleteidvault) | **Delete** /iam-deploy/idvault | 
*DefaultApi* | [**DeleteIdpAzure**](docs/DefaultApi.md#deleteidpazure) | **Delete** /iam-deploy/idp_azure | 
*DefaultApi* | [**DeleteIdpFacebook**](docs/DefaultApi.md#deleteidpfacebook) | **Delete** /iam-deploy/idp_fb | 
*DefaultApi* | [**DeleteIdpGoogle**](docs/DefaultApi.md#deleteidpgoogle) | **Delete** /iam-deploy/idp_google | 
*DefaultApi* | [**DeleteIisExecEnv**](docs/DefaultApi.md#deleteiisexecenv) | **Delete** /iam-deploy/iisexecenv | 
*DefaultApi* | [**DeleteIntSaml2Sp**](docs/DefaultApi.md#deleteintsaml2sp) | **Delete** /iam-deploy/intsaml2sp | 
*DefaultApi* | [**DeleteJossoRs**](docs/DefaultApi.md#deletejossors) | **Delete** /iam-deploy/jossors | 
*DefaultApi* | [**DeleteOidcRp**](docs/DefaultApi.md#deleteoidcrp) | **Delete** /iam-deploy/oidcrp | 
*DefaultApi* | [**DeletePhpExecEnv**](docs/DefaultApi.md#deletephpexecenv) | **Delete** /iam-deploy/phpexecenv | 
*DefaultApi* | [**DeleteSharepointRs**](docs/DefaultApi.md#deletesharepointrs) | **Delete** /iam-deploy/sharepointrs | 
*DefaultApi* | [**DeleteTomcatExecEnv**](docs/DefaultApi.md#deletetomcatexecenv) | **Delete** /iam-deploy/tomcatexecenv | 
*DefaultApi* | [**DeleteVirtSaml2Sp**](docs/DefaultApi.md#deletevirtsaml2sp) | **Delete** /iam-deploy/virtsaml2sp | 
*DefaultApi* | [**DeleteWeblogicExecEnv**](docs/DefaultApi.md#deleteweblogicexecenv) | **Delete** /iam-deploy/weblogicexecenv | 
*DefaultApi* | [**GetAllBrandings**](docs/DefaultApi.md#getallbrandings) | **Get** /iam-branding/brandings | 
*DefaultApi* | [**GetAppliance**](docs/DefaultApi.md#getappliance) | **Get** /iam-deploy/appliance | 
*DefaultApi* | [**GetApplianceContainer**](docs/DefaultApi.md#getappliancecontainer) | **Get** /iam-deploy/appliance-container | 
*DefaultApi* | [**GetApplianceContainers**](docs/DefaultApi.md#getappliancecontainers) | **Get** /iam-deploy/appliance-containers | 
*DefaultApi* | [**GetApplianceState**](docs/DefaultApi.md#getappliancestate) | **Get** /iam-deploy/appliance/state | 
*DefaultApi* | [**GetAppliances**](docs/DefaultApi.md#getappliances) | **Get** /iam-deploy/appliances | 
*DefaultApi* | [**GetBranding**](docs/DefaultApi.md#getbranding) | **Get** /iam-branding/branding | 
*DefaultApi* | [**GetDbIdVault**](docs/DefaultApi.md#getdbidvault) | **Get** /iam-deploy/dbidvault | 
*DefaultApi* | [**GetDbIdVaults**](docs/DefaultApi.md#getdbidvaults) | **Get** /iam-deploy/dbidvaults | 
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
*DefaultApi* | [**GetIdpAzure**](docs/DefaultApi.md#getidpazure) | **Get** /iam-deploy/idp_azure | 
*DefaultApi* | [**GetIdpAzures**](docs/DefaultApi.md#getidpazures) | **Get** /iam-deploy/idp_azures | 
*DefaultApi* | [**GetIdpFacebook**](docs/DefaultApi.md#getidpfacebook) | **Get** /iam-deploy/idp_fb | 
*DefaultApi* | [**GetIdpFacebooks**](docs/DefaultApi.md#getidpfacebooks) | **Get** /iam-deploy/idp_fbs | 
*DefaultApi* | [**GetIdpGoogle**](docs/DefaultApi.md#getidpgoogle) | **Get** /iam-deploy/idp_google | 
*DefaultApi* | [**GetIdpGoogles**](docs/DefaultApi.md#getidpgoogles) | **Get** /iam-deploy/idp_googles | 
*DefaultApi* | [**GetIisExecEnv**](docs/DefaultApi.md#getiisexecenv) | **Get** /iam-deploy/iisexecenv | 
*DefaultApi* | [**GetIisExecEnvs**](docs/DefaultApi.md#getiisexecenvs) | **Get** /iam-deploy/iisexecenvs | 
*DefaultApi* | [**GetIntSaml2Sp**](docs/DefaultApi.md#getintsaml2sp) | **Get** /iam-deploy/intsaml2sp | 
*DefaultApi* | [**GetIntSaml2Sps**](docs/DefaultApi.md#getintsaml2sps) | **Get** /iam-deploy/intsaml2sps | 
*DefaultApi* | [**GetJossoRs**](docs/DefaultApi.md#getjossors) | **Get** /iam-deploy/jossors | 
*DefaultApi* | [**GetJossoRss**](docs/DefaultApi.md#getjossorss) | **Get** /iam-deploy/jossorss | 
*DefaultApi* | [**GetOidcRp**](docs/DefaultApi.md#getoidcrp) | **Get** /iam-deploy/oidcrp | 
*DefaultApi* | [**GetOidcRps**](docs/DefaultApi.md#getoidcrps) | **Get** /iam-deploy/oidcrps | 
*DefaultApi* | [**GetPhpExecEnv**](docs/DefaultApi.md#getphpexecenv) | **Get** /iam-deploy/phpexecenv | 
*DefaultApi* | [**GetPhpExecEnvs**](docs/DefaultApi.md#getphpexecenvs) | **Get** /iam-deploy/phpexecenvs | 
*DefaultApi* | [**GetProvider**](docs/DefaultApi.md#getprovider) | **Get** /iam-deploy/provider | 
*DefaultApi* | [**GetProviders**](docs/DefaultApi.md#getproviders) | **Get** /iam-deploy/providers | 
*DefaultApi* | [**GetSharepointRs**](docs/DefaultApi.md#getsharepointrs) | **Get** /iam-deploy/sharepointrs | 
*DefaultApi* | [**GetSharepointRss**](docs/DefaultApi.md#getsharepointrss) | **Get** /iam-deploy/sharepointrss | 
*DefaultApi* | [**GetTomcatExecEnv**](docs/DefaultApi.md#gettomcatexecenv) | **Get** /iam-deploy/tomcatexecenv | 
*DefaultApi* | [**GetTomcatExecEnvs**](docs/DefaultApi.md#gettomcatexecenvs) | **Get** /iam-deploy/tomcatexecenvs | 
*DefaultApi* | [**GetTypes**](docs/DefaultApi.md#gettypes) | **Get** /iam-deploy/noop/types | 
*DefaultApi* | [**GetVirtSaml2Sp**](docs/DefaultApi.md#getvirtsaml2sp) | **Get** /iam-deploy/virtsaml2sp | 
*DefaultApi* | [**GetVirtSaml2Sps**](docs/DefaultApi.md#getvirtsaml2sps) | **Get** /iam-deploy/virtsaml2sps | 
*DefaultApi* | [**GetWeblogicExecEnv**](docs/DefaultApi.md#getweblogicexecenv) | **Get** /iam-deploy/weblogicexecenv | 
*DefaultApi* | [**GetWeblogicExecEnvs**](docs/DefaultApi.md#getweblogicexecenvs) | **Get** /iam-deploy/weblogicexecenvs | 
*DefaultApi* | [**ImportAppliance**](docs/DefaultApi.md#importappliance) | **Post** /iam-deploy/appliance/import | 
*DefaultApi* | [**LayoutAppliance**](docs/DefaultApi.md#layoutappliance) | **Get** /iam-deploy/appliance/layout | 
*DefaultApi* | [**RefreshBrandings**](docs/DefaultApi.md#refreshbrandings) | **Get** /iam-branding/brandings/refresh | 
*DefaultApi* | [**SignOn**](docs/DefaultApi.md#signon) | **Post** /iam-authn/sign-on | 
*DefaultApi* | [**StartAppliance**](docs/DefaultApi.md#startappliance) | **Get** /iam-deploy/appliance/start | 
*DefaultApi* | [**StopAppliance**](docs/DefaultApi.md#stopappliance) | **Get** /iam-deploy/appliance/stop | 
*DefaultApi* | [**UpdateAppliance**](docs/DefaultApi.md#updateappliance) | **Put** /iam-deploy/appliance | 
*DefaultApi* | [**UpdateBranding**](docs/DefaultApi.md#updatebranding) | **Put** /iam-branding/branding | 
*DefaultApi* | [**UpdateDbIdVault**](docs/DefaultApi.md#updatedbidvault) | **Put** /iam-deploy/dbidvault | 
*DefaultApi* | [**UpdateExtSaml2Sp**](docs/DefaultApi.md#updateextsaml2sp) | **Put** /iam-deploy/extsaml2sp | 
*DefaultApi* | [**UpdateIdP**](docs/DefaultApi.md#updateidp) | **Put** /iam-deploy/idp | 
*DefaultApi* | [**UpdateIdSourceDb**](docs/DefaultApi.md#updateidsourcedb) | **Put** /iam-deploy/idsourcedb | 
*DefaultApi* | [**UpdateIdSourceLdap**](docs/DefaultApi.md#updateidsourceldap) | **Put** /iam-deploy/idsourceldap | 
*DefaultApi* | [**UpdateIdVault**](docs/DefaultApi.md#updateidvault) | **Put** /iam-deploy/idvault | 
*DefaultApi* | [**UpdateIdpAzure**](docs/DefaultApi.md#updateidpazure) | **Put** /iam-deploy/idp_azure | 
*DefaultApi* | [**UpdateIdpFacebook**](docs/DefaultApi.md#updateidpfacebook) | **Put** /iam-deploy/idp_fb | 
*DefaultApi* | [**UpdateIdpGoogle**](docs/DefaultApi.md#updateidpgoogle) | **Put** /iam-deploy/idp_google | 
*DefaultApi* | [**UpdateIisExecEnv**](docs/DefaultApi.md#updateiisexecenv) | **Put** /iam-deploy/iisexecenv | 
*DefaultApi* | [**UpdateIntSaml2Sp**](docs/DefaultApi.md#updateintsaml2sp) | **Put** /iam-deploy/intsaml2sp | 
*DefaultApi* | [**UpdateJossoRs**](docs/DefaultApi.md#updatejossors) | **Put** /iam-deploy/jossors | 
*DefaultApi* | [**UpdateOidcRp**](docs/DefaultApi.md#updateoidcrp) | **Put** /iam-deploy/oidcrp | 
*DefaultApi* | [**UpdatePhpExecEnv**](docs/DefaultApi.md#updatephpexecenv) | **Put** /iam-deploy/phpexecenv | 
*DefaultApi* | [**UpdateSharepointRs**](docs/DefaultApi.md#updatesharepointrs) | **Put** /iam-deploy/sharepointrs | 
*DefaultApi* | [**UpdateTomcatExecEnv**](docs/DefaultApi.md#updatetomcatexecenv) | **Put** /iam-deploy/tomcatexecenv | 
*DefaultApi* | [**UpdateVirtSaml2Sp**](docs/DefaultApi.md#updatevirtsaml2sp) | **Put** /iam-deploy/virtsaml2sp | 
*DefaultApi* | [**UpdateWeblogicExecEnv**](docs/DefaultApi.md#updateweblogicexecenv) | **Put** /iam-deploy/weblogicexecenv | 
*DefaultApi* | [**ValidateAppliance**](docs/DefaultApi.md#validateappliance) | **Get** /iam-deploy/appliance/validate | 


## Documentation For Models

 - [AccountLinkagePolicyDTO](docs/AccountLinkagePolicyDTO.md)
 - [ActivateExecEnvReq](docs/ActivateExecEnvReq.md)
 - [ActivateExecEnvRes](docs/ActivateExecEnvRes.md)
 - [ActivationDTO](docs/ActivationDTO.md)
 - [AttributeMapperProfileDTO](docs/AttributeMapperProfileDTO.md)
 - [AttributeMappingDTO](docs/AttributeMappingDTO.md)
 - [AttributeProfileDTO](docs/AttributeProfileDTO.md)
 - [AttributeValueDTO](docs/AttributeValueDTO.md)
 - [AuthenticationAssertionEmissionPolicyDTO](docs/AuthenticationAssertionEmissionPolicyDTO.md)
 - [AuthenticationContractDTO](docs/AuthenticationContractDTO.md)
 - [AuthenticationMechanismDTO](docs/AuthenticationMechanismDTO.md)
 - [AuthenticationServiceDTO](docs/AuthenticationServiceDTO.md)
 - [AzureOpenIDConnectIdentityProviderDTO](docs/AzureOpenIDConnectIdentityProviderDTO.md)
 - [BasicAuthenticationDTO](docs/BasicAuthenticationDTO.md)
 - [BindAuthenticationDTO](docs/BindAuthenticationDTO.md)
 - [BrandingDefinitionDTO](docs/BrandingDefinitionDTO.md)
 - [BuiltInAttributeProfileDTO](docs/BuiltInAttributeProfileDTO.md)
 - [CalcLayoutReq](docs/CalcLayoutReq.md)
 - [CalcLayoutRes](docs/CalcLayoutRes.md)
 - [CustomBrandingDefinitionDTO](docs/CustomBrandingDefinitionDTO.md)
 - [CustomClassDTO](docs/CustomClassDTO.md)
 - [CustomClassPropertyDTO](docs/CustomClassPropertyDTO.md)
 - [DbIdentitySourceDTO](docs/DbIdentitySourceDTO.md)
 - [DbIdentityVaultDTO](docs/DbIdentityVaultDTO.md)
 - [DelegatedAuthenticationDTO](docs/DelegatedAuthenticationDTO.md)
 - [DeleteBrandingReq](docs/DeleteBrandingReq.md)
 - [DeleteBrandingRes](docs/DeleteBrandingRes.md)
 - [DeleteReq](docs/DeleteReq.md)
 - [DeleteRes](docs/DeleteRes.md)
 - [DirectoryAuthenticationServiceDTO](docs/DirectoryAuthenticationServiceDTO.md)
 - [EmbeddedIdentityVaultDTO](docs/EmbeddedIdentityVaultDTO.md)
 - [EntitySelectionStrategyDTO](docs/EntitySelectionStrategyDTO.md)
 - [ExecutionEnvironmentDTO](docs/ExecutionEnvironmentDTO.md)
 - [ExtensionDTO](docs/ExtensionDTO.md)
 - [ExternalOpenIDConnectRelayingPartyDTO](docs/ExternalOpenIDConnectRelayingPartyDTO.md)
 - [ExternalSaml2ServiceProviderDTO](docs/ExternalSaml2ServiceProviderDTO.md)
 - [FacebookOpenIDConnectIdentityProviderDTO](docs/FacebookOpenIDConnectIdentityProviderDTO.md)
 - [FederatedChannelDTO](docs/FederatedChannelDTO.md)
 - [FederatedConnectionDTO](docs/FederatedConnectionDTO.md)
 - [FederatedProviderDTO](docs/FederatedProviderDTO.md)
 - [GetAllBrandingsReq](docs/GetAllBrandingsReq.md)
 - [GetAllBrandingsRes](docs/GetAllBrandingsRes.md)
 - [GetApplianceContainerRes](docs/GetApplianceContainerRes.md)
 - [GetApplianceContainersRes](docs/GetApplianceContainersRes.md)
 - [GetApplianceReq](docs/GetApplianceReq.md)
 - [GetApplianceRes](docs/GetApplianceRes.md)
 - [GetApplianceStateRes](docs/GetApplianceStateRes.md)
 - [GetAppliancesRes](docs/GetAppliancesRes.md)
 - [GetBrandingReq](docs/GetBrandingReq.md)
 - [GetBrandingRes](docs/GetBrandingRes.md)
 - [GetDbIdVaultReq](docs/GetDbIdVaultReq.md)
 - [GetDbIdVaultRes](docs/GetDbIdVaultRes.md)
 - [GetDbIdVaultsRes](docs/GetDbIdVaultsRes.md)
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
 - [GetIdpAzureReq](docs/GetIdpAzureReq.md)
 - [GetIdpAzureRes](docs/GetIdpAzureRes.md)
 - [GetIdpAzuresRes](docs/GetIdpAzuresRes.md)
 - [GetIdpFacebookReq](docs/GetIdpFacebookReq.md)
 - [GetIdpFacebookRes](docs/GetIdpFacebookRes.md)
 - [GetIdpFacebooksRes](docs/GetIdpFacebooksRes.md)
 - [GetIdpGoogleReq](docs/GetIdpGoogleReq.md)
 - [GetIdpGoogleRes](docs/GetIdpGoogleRes.md)
 - [GetIdpGooglesRes](docs/GetIdpGooglesRes.md)
 - [GetIisExecEnvReq](docs/GetIisExecEnvReq.md)
 - [GetIisExecEnvRes](docs/GetIisExecEnvRes.md)
 - [GetIisExecEnvsRes](docs/GetIisExecEnvsRes.md)
 - [GetIntSaml2SpReq](docs/GetIntSaml2SpReq.md)
 - [GetIntSaml2SpRes](docs/GetIntSaml2SpRes.md)
 - [GetIntSaml2SpsRes](docs/GetIntSaml2SpsRes.md)
 - [GetJossoRsReq](docs/GetJossoRsReq.md)
 - [GetJossoRsRes](docs/GetJossoRsRes.md)
 - [GetJossoRssRes](docs/GetJossoRssRes.md)
 - [GetOidcRpReq](docs/GetOidcRpReq.md)
 - [GetOidcRpRes](docs/GetOidcRpRes.md)
 - [GetOidcRpsRes](docs/GetOidcRpsRes.md)
 - [GetPhpExecEnvReq](docs/GetPhpExecEnvReq.md)
 - [GetPhpExecEnvRes](docs/GetPhpExecEnvRes.md)
 - [GetPhpExecEnvsRes](docs/GetPhpExecEnvsRes.md)
 - [GetProviderReq](docs/GetProviderReq.md)
 - [GetProviderRes](docs/GetProviderRes.md)
 - [GetProvidersRes](docs/GetProvidersRes.md)
 - [GetSharepointRsReq](docs/GetSharepointRsReq.md)
 - [GetSharepointRsRes](docs/GetSharepointRsRes.md)
 - [GetSharepointRssRes](docs/GetSharepointRssRes.md)
 - [GetTomcatExecEnvReq](docs/GetTomcatExecEnvReq.md)
 - [GetTomcatExecEnvRes](docs/GetTomcatExecEnvRes.md)
 - [GetTomcatExecEnvsRes](docs/GetTomcatExecEnvsRes.md)
 - [GetTypesRes](docs/GetTypesRes.md)
 - [GetVirtSaml2SpReq](docs/GetVirtSaml2SpReq.md)
 - [GetVirtSaml2SpRes](docs/GetVirtSaml2SpRes.md)
 - [GetVirtSaml2SpsRes](docs/GetVirtSaml2SpsRes.md)
 - [GetWeblogicExecEnvReq](docs/GetWeblogicExecEnvReq.md)
 - [GetWeblogicExecEnvRes](docs/GetWeblogicExecEnvRes.md)
 - [GetWeblogicExecEnvsRes](docs/GetWeblogicExecEnvsRes.md)
 - [GoogleOpenIDConnectIdentityProviderDTO](docs/GoogleOpenIDConnectIdentityProviderDTO.md)
 - [GroupDTO](docs/GroupDTO.md)
 - [IdentityApplianceContainerDTO](docs/IdentityApplianceContainerDTO.md)
 - [IdentityApplianceDTO](docs/IdentityApplianceDTO.md)
 - [IdentityApplianceDefinitionDTO](docs/IdentityApplianceDefinitionDTO.md)
 - [IdentityApplianceDeploymentDTO](docs/IdentityApplianceDeploymentDTO.md)
 - [IdentityApplianceSecurityConfigDTO](docs/IdentityApplianceSecurityConfigDTO.md)
 - [IdentityApplianceUnitDTO](docs/IdentityApplianceUnitDTO.md)
 - [IdentityLookupDTO](docs/IdentityLookupDTO.md)
 - [IdentityMappingPolicyDTO](docs/IdentityMappingPolicyDTO.md)
 - [IdentityProviderChannelDTO](docs/IdentityProviderChannelDTO.md)
 - [IdentityProviderDTO](docs/IdentityProviderDTO.md)
 - [IdentitySourceDTO](docs/IdentitySourceDTO.md)
 - [ImpersonateUserPolicyDTO](docs/ImpersonateUserPolicyDTO.md)
 - [ImportApplianceReq](docs/ImportApplianceReq.md)
 - [ImportApplianceRes](docs/ImportApplianceRes.md)
 - [InternalSaml2ServiceProviderChannelDTO](docs/InternalSaml2ServiceProviderChannelDTO.md)
 - [InternalSaml2ServiceProviderDTO](docs/InternalSaml2ServiceProviderDTO.md)
 - [JOSSO1ResourceDTO](docs/JOSSO1ResourceDTO.md)
 - [KeystoreDTO](docs/KeystoreDTO.md)
 - [LdapIdentitySourceDTO](docs/LdapIdentitySourceDTO.md)
 - [LocationDTO](docs/LocationDTO.md)
 - [OAuth2ClientDTO](docs/OAuth2ClientDTO.md)
 - [OIDCSignOnRequest](docs/OIDCSignOnRequest.md)
 - [OIDCSignOnResponse](docs/OIDCSignOnResponse.md)
 - [PHPExecutionEnvironmentDTO](docs/PHPExecutionEnvironmentDTO.md)
 - [PointDTO](docs/PointDTO.md)
 - [ProviderConfigDTO](docs/ProviderConfigDTO.md)
 - [ProviderContainerDTO](docs/ProviderContainerDTO.md)
 - [ProviderDTO](docs/ProviderDTO.md)
 - [RefreshBrandingsRes](docs/RefreshBrandingsRes.md)
 - [ResourceDTO](docs/ResourceDTO.md)
 - [SamlR2IDPConfigDTO](docs/SamlR2IDPConfigDTO.md)
 - [SamlR2SPConfigDTO](docs/SamlR2SPConfigDTO.md)
 - [ServerContext](docs/ServerContext.md)
 - [ServiceConnectionDTO](docs/ServiceConnectionDTO.md)
 - [ServiceResourceDTO](docs/ServiceResourceDTO.md)
 - [SessionManagerFactoryDTO](docs/SessionManagerFactoryDTO.md)
 - [SetApplianceStateReq](docs/SetApplianceStateReq.md)
 - [SharepointResourceDTO](docs/SharepointResourceDTO.md)
 - [StoreApplianceReq](docs/StoreApplianceReq.md)
 - [StoreApplianceRes](docs/StoreApplianceRes.md)
 - [StoreBrandingReq](docs/StoreBrandingReq.md)
 - [StoreBrandingRes](docs/StoreBrandingRes.md)
 - [StoreDbIdVaultReq](docs/StoreDbIdVaultReq.md)
 - [StoreDbIdVaultRes](docs/StoreDbIdVaultRes.md)
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
 - [StoreIdpAzureReq](docs/StoreIdpAzureReq.md)
 - [StoreIdpAzureRes](docs/StoreIdpAzureRes.md)
 - [StoreIdpFacebookReq](docs/StoreIdpFacebookReq.md)
 - [StoreIdpFacebookRes](docs/StoreIdpFacebookRes.md)
 - [StoreIdpGoogleReq](docs/StoreIdpGoogleReq.md)
 - [StoreIdpGoogleRes](docs/StoreIdpGoogleRes.md)
 - [StoreIisExecEnvReq](docs/StoreIisExecEnvReq.md)
 - [StoreIisExecEnvRes](docs/StoreIisExecEnvRes.md)
 - [StoreIntSaml2SpReq](docs/StoreIntSaml2SpReq.md)
 - [StoreIntSaml2SpRes](docs/StoreIntSaml2SpRes.md)
 - [StoreJossoRsReq](docs/StoreJossoRsReq.md)
 - [StoreJossoRsRes](docs/StoreJossoRsRes.md)
 - [StoreOidcRpReq](docs/StoreOidcRpReq.md)
 - [StoreOidcRpRes](docs/StoreOidcRpRes.md)
 - [StorePhpExecEnvReq](docs/StorePhpExecEnvReq.md)
 - [StorePhpExecEnvRes](docs/StorePhpExecEnvRes.md)
 - [StoreSharepointRsReq](docs/StoreSharepointRsReq.md)
 - [StoreSharepointRsRes](docs/StoreSharepointRsRes.md)
 - [StoreTomcatExecEnvReq](docs/StoreTomcatExecEnvReq.md)
 - [StoreTomcatExecEnvRes](docs/StoreTomcatExecEnvRes.md)
 - [StoreVirtSaml2SpReq](docs/StoreVirtSaml2SpReq.md)
 - [StoreVirtSaml2SpRes](docs/StoreVirtSaml2SpRes.md)
 - [StoreWeblogicExecEnvReq](docs/StoreWeblogicExecEnvReq.md)
 - [StoreWeblogicExecEnvRes](docs/StoreWeblogicExecEnvRes.md)
 - [SubjectAuthenticationPolicyDTO](docs/SubjectAuthenticationPolicyDTO.md)
 - [SubjectNameIdentifierPolicyDTO](docs/SubjectNameIdentifierPolicyDTO.md)
 - [TOTPAuthenticationServiceDTO](docs/TOTPAuthenticationServiceDTO.md)
 - [TomcatExecutionEnvironmentDTO](docs/TomcatExecutionEnvironmentDTO.md)
 - [UserDTO](docs/UserDTO.md)
 - [UserDashboardBrandingDTO](docs/UserDashboardBrandingDTO.md)
 - [VirtualSaml2ServiceProviderDTO](docs/VirtualSaml2ServiceProviderDTO.md)
 - [WeblogicExecutionEnvironmentDTO](docs/WeblogicExecutionEnvironmentDTO.md)
 - [WindowsIISExecutionEnvironmentDTO](docs/WindowsIISExecutionEnvironmentDTO.md)


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

