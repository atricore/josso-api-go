# \DefaultApi

All URIs are relative to *http://localhost:8081/atricore-res/services*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppliance**](DefaultApi.md#CreateAppliance) | **Post** /iam-deploy/appliance | 
[**CreateIdP**](DefaultApi.md#CreateIdP) | **Post** /iam-deploy/idp | 
[**CreateIdSourceLdap**](DefaultApi.md#CreateIdSourceLdap) | **Post** /iam-deploy/idsourceldap | 
[**CreateIdVault**](DefaultApi.md#CreateIdVault) | **Post** /iam-deploy/idvault | 
[**DeleteAppliance**](DefaultApi.md#DeleteAppliance) | **Delete** /iam-deploy/appliance | 
[**DeleteIdP**](DefaultApi.md#DeleteIdP) | **Delete** /iam-deploy/idp | 
[**DeleteIdSourceLdap**](DefaultApi.md#DeleteIdSourceLdap) | **Delete** /iam-deploy/idsourceldap | 
[**DeleteIdVault**](DefaultApi.md#DeleteIdVault) | **Delete** /iam-deploy/idvault | 
[**GetAppliance**](DefaultApi.md#GetAppliance) | **Get** /iam-deploy/appliance | 
[**GetAppliances**](DefaultApi.md#GetAppliances) | **Get** /iam-deploy/appliances | 
[**GetIdP**](DefaultApi.md#GetIdP) | **Get** /iam-deploy/idp | 
[**GetIdPs**](DefaultApi.md#GetIdPs) | **Get** /iam-deploy/idps | 
[**GetIdSourceLdap**](DefaultApi.md#GetIdSourceLdap) | **Get** /iam-deploy/idsourceldap | 
[**GetIdSourceLdaps**](DefaultApi.md#GetIdSourceLdaps) | **Get** /iam-deploy/idsourceldaps | 
[**GetIdVault**](DefaultApi.md#GetIdVault) | **Get** /iam-deploy/idvault | 
[**GetIdVaults**](DefaultApi.md#GetIdVaults) | **Get** /iam-deploy/idvaults | 
[**SignOn**](DefaultApi.md#SignOn) | **Post** /iam-authn/sign-on | 
[**UpdateAppliance**](DefaultApi.md#UpdateAppliance) | **Put** /iam-deploy/appliance | 
[**UpdateIdP**](DefaultApi.md#UpdateIdP) | **Put** /iam-deploy/idp | 
[**UpdateIdSourceLdap**](DefaultApi.md#UpdateIdSourceLdap) | **Put** /iam-deploy/idsourceldap | 
[**UpdateIdVault**](DefaultApi.md#UpdateIdVault) | **Put** /iam-deploy/idvault | 



## CreateAppliance

> StoreApplianceRes CreateAppliance(ctx).StoreApplianceReq(storeApplianceReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeApplianceReq := *openapiclient.NewStoreApplianceReq() // StoreApplianceReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAppliance(context.Background()).StoreApplianceReq(storeApplianceReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppliance`: StoreApplianceRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAppliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeApplianceReq** | [**StoreApplianceReq**](StoreApplianceReq.md) |  | 

### Return type

[**StoreApplianceRes**](StoreApplianceRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdP

> StoreIdPRes CreateIdP(ctx).StoreIdPReq(storeIdPReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeIdPReq := *openapiclient.NewStoreIdPReq() // StoreIdPReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIdP(context.Background()).StoreIdPReq(storeIdPReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIdP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdP`: StoreIdPRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIdP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeIdPReq** | [**StoreIdPReq**](StoreIdPReq.md) |  | 

### Return type

[**StoreIdPRes**](StoreIdPRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdSourceLdap

> StoreIdSourceLdapRes CreateIdSourceLdap(ctx).StoreIdSourceLdapReq(storeIdSourceLdapReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeIdSourceLdapReq := *openapiclient.NewStoreIdSourceLdapReq() // StoreIdSourceLdapReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIdSourceLdap(context.Background()).StoreIdSourceLdapReq(storeIdSourceLdapReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIdSourceLdap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdSourceLdap`: StoreIdSourceLdapRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIdSourceLdap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdSourceLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeIdSourceLdapReq** | [**StoreIdSourceLdapReq**](StoreIdSourceLdapReq.md) |  | 

### Return type

[**StoreIdSourceLdapRes**](StoreIdSourceLdapRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdVault

> StoreIdVaultRes CreateIdVault(ctx).StoreIdVaultReq(storeIdVaultReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeIdVaultReq := *openapiclient.NewStoreIdVaultReq() // StoreIdVaultReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIdVault(context.Background()).StoreIdVaultReq(storeIdVaultReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIdVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdVault`: StoreIdVaultRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIdVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeIdVaultReq** | [**StoreIdVaultReq**](StoreIdVaultReq.md) |  | 

### Return type

[**StoreIdVaultRes**](StoreIdVaultRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppliance

> DeleteRes DeleteAppliance(ctx).DeleteReq(deleteReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteReq := *openapiclient.NewDeleteReq() // DeleteReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAppliance(context.Background()).DeleteReq(deleteReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppliance`: DeleteRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAppliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteReq** | [**DeleteReq**](DeleteReq.md) |  | 

### Return type

[**DeleteRes**](DeleteRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdP

> DeleteRes DeleteIdP(ctx).DeleteReq(deleteReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteReq := *openapiclient.NewDeleteReq() // DeleteReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIdP(context.Background()).DeleteReq(deleteReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIdP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdP`: DeleteRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIdP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteReq** | [**DeleteReq**](DeleteReq.md) |  | 

### Return type

[**DeleteRes**](DeleteRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdSourceLdap

> DeleteRes DeleteIdSourceLdap(ctx).DeleteReq(deleteReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteReq := *openapiclient.NewDeleteReq() // DeleteReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIdSourceLdap(context.Background()).DeleteReq(deleteReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIdSourceLdap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdSourceLdap`: DeleteRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIdSourceLdap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdSourceLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteReq** | [**DeleteReq**](DeleteReq.md) |  | 

### Return type

[**DeleteRes**](DeleteRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdVault

> DeleteRes DeleteIdVault(ctx).DeleteReq(deleteReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteReq := *openapiclient.NewDeleteReq() // DeleteReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIdVault(context.Background()).DeleteReq(deleteReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIdVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdVault`: DeleteRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIdVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteReq** | [**DeleteReq**](DeleteReq.md) |  | 

### Return type

[**DeleteRes**](DeleteRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliance

> GetApplianceRes GetAppliance(ctx).GetApplianceReq(getApplianceReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getApplianceReq := *openapiclient.NewGetApplianceReq() // GetApplianceReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAppliance(context.Background()).GetApplianceReq(getApplianceReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppliance`: GetApplianceRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAppliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getApplianceReq** | [**GetApplianceReq**](GetApplianceReq.md) |  | 

### Return type

[**GetApplianceRes**](GetApplianceRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliances

> GetAppliancesRes GetAppliances(ctx).GetApplianceReq(getApplianceReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getApplianceReq := *openapiclient.NewGetApplianceReq() // GetApplianceReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAppliances(context.Background()).GetApplianceReq(getApplianceReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAppliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppliances`: GetAppliancesRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAppliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getApplianceReq** | [**GetApplianceReq**](GetApplianceReq.md) |  | 

### Return type

[**GetAppliancesRes**](GetAppliancesRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdP

> GetIdPRes GetIdP(ctx).GetIdPReq(getIdPReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getIdPReq := *openapiclient.NewGetIdPReq() // GetIdPReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIdP(context.Background()).GetIdPReq(getIdPReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIdP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdP`: GetIdPRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIdP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getIdPReq** | [**GetIdPReq**](GetIdPReq.md) |  | 

### Return type

[**GetIdPRes**](GetIdPRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdPs

> GetIdPsRes GetIdPs(ctx).GetIdPReq(getIdPReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getIdPReq := *openapiclient.NewGetIdPReq() // GetIdPReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIdPs(context.Background()).GetIdPReq(getIdPReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIdPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdPs`: GetIdPsRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIdPs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getIdPReq** | [**GetIdPReq**](GetIdPReq.md) |  | 

### Return type

[**GetIdPsRes**](GetIdPsRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdSourceLdap

> GetIdSourceLdapRes GetIdSourceLdap(ctx).GetIdSourceLdapReq(getIdSourceLdapReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getIdSourceLdapReq := *openapiclient.NewGetIdSourceLdapReq() // GetIdSourceLdapReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIdSourceLdap(context.Background()).GetIdSourceLdapReq(getIdSourceLdapReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIdSourceLdap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdSourceLdap`: GetIdSourceLdapRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIdSourceLdap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdSourceLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getIdSourceLdapReq** | [**GetIdSourceLdapReq**](GetIdSourceLdapReq.md) |  | 

### Return type

[**GetIdSourceLdapRes**](GetIdSourceLdapRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdSourceLdaps

> GetIdSourceLdapsRes GetIdSourceLdaps(ctx).GetIdSourceLdapReq(getIdSourceLdapReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getIdSourceLdapReq := *openapiclient.NewGetIdSourceLdapReq() // GetIdSourceLdapReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIdSourceLdaps(context.Background()).GetIdSourceLdapReq(getIdSourceLdapReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIdSourceLdaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdSourceLdaps`: GetIdSourceLdapsRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIdSourceLdaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdSourceLdapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getIdSourceLdapReq** | [**GetIdSourceLdapReq**](GetIdSourceLdapReq.md) |  | 

### Return type

[**GetIdSourceLdapsRes**](GetIdSourceLdapsRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdVault

> GetIdVaultRes GetIdVault(ctx).GetIdVaultReq(getIdVaultReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getIdVaultReq := *openapiclient.NewGetIdVaultReq() // GetIdVaultReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIdVault(context.Background()).GetIdVaultReq(getIdVaultReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIdVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdVault`: GetIdVaultRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIdVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getIdVaultReq** | [**GetIdVaultReq**](GetIdVaultReq.md) |  | 

### Return type

[**GetIdVaultRes**](GetIdVaultRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdVaults

> GetIdVaultsRes GetIdVaults(ctx).GetIdVaultReq(getIdVaultReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getIdVaultReq := *openapiclient.NewGetIdVaultReq() // GetIdVaultReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIdVaults(context.Background()).GetIdVaultReq(getIdVaultReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIdVaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdVaults`: GetIdVaultsRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIdVaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdVaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getIdVaultReq** | [**GetIdVaultReq**](GetIdVaultReq.md) |  | 

### Return type

[**GetIdVaultsRes**](GetIdVaultsRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignOn

> OIDCSignOnResponse SignOn(ctx).OIDCSignOnRequest(oIDCSignOnRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    oIDCSignOnRequest := *openapiclient.NewOIDCSignOnRequest() // OIDCSignOnRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SignOn(context.Background()).OIDCSignOnRequest(oIDCSignOnRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SignOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignOn`: OIDCSignOnResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SignOn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCSignOnRequest** | [**OIDCSignOnRequest**](OIDCSignOnRequest.md) |  | 

### Return type

[**OIDCSignOnResponse**](OIDCSignOnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppliance

> StoreApplianceRes UpdateAppliance(ctx).StoreApplianceReq(storeApplianceReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeApplianceReq := *openapiclient.NewStoreApplianceReq() // StoreApplianceReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAppliance(context.Background()).StoreApplianceReq(storeApplianceReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppliance`: StoreApplianceRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAppliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeApplianceReq** | [**StoreApplianceReq**](StoreApplianceReq.md) |  | 

### Return type

[**StoreApplianceRes**](StoreApplianceRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdP

> StoreIdPRes UpdateIdP(ctx).StoreIdPReq(storeIdPReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeIdPReq := *openapiclient.NewStoreIdPReq() // StoreIdPReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateIdP(context.Background()).StoreIdPReq(storeIdPReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIdP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdP`: StoreIdPRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIdP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeIdPReq** | [**StoreIdPReq**](StoreIdPReq.md) |  | 

### Return type

[**StoreIdPRes**](StoreIdPRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdSourceLdap

> StoreIdSourceLdapRes UpdateIdSourceLdap(ctx).StoreIdSourceLdapReq(storeIdSourceLdapReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeIdSourceLdapReq := *openapiclient.NewStoreIdSourceLdapReq() // StoreIdSourceLdapReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateIdSourceLdap(context.Background()).StoreIdSourceLdapReq(storeIdSourceLdapReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIdSourceLdap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdSourceLdap`: StoreIdSourceLdapRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIdSourceLdap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdSourceLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeIdSourceLdapReq** | [**StoreIdSourceLdapReq**](StoreIdSourceLdapReq.md) |  | 

### Return type

[**StoreIdSourceLdapRes**](StoreIdSourceLdapRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdVault

> StoreIdVaultRes UpdateIdVault(ctx).StoreIdVaultReq(storeIdVaultReq).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeIdVaultReq := *openapiclient.NewStoreIdVaultReq() // StoreIdVaultReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateIdVault(context.Background()).StoreIdVaultReq(storeIdVaultReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIdVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdVault`: StoreIdVaultRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIdVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeIdVaultReq** | [**StoreIdVaultReq**](StoreIdVaultReq.md) |  | 

### Return type

[**StoreIdVaultRes**](StoreIdVaultRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

