# \DefaultApi

All URIs are relative to *http://localhost:8081/atricore-res/services*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppliance**](DefaultApi.md#CreateAppliance) | **Post** /iam-deploy/appliance | 
[**CreateIdP**](DefaultApi.md#CreateIdP) | **Post** /iam-deploy/idp | 
[**DeleteAppliance**](DefaultApi.md#DeleteAppliance) | **Delete** /iam-deploy/appliance | 
[**DeleteIdP**](DefaultApi.md#DeleteIdP) | **Delete** /iam-deploy/idp | 
[**GetAppliance**](DefaultApi.md#GetAppliance) | **Get** /iam-deploy/appliance | 
[**GetAppliances**](DefaultApi.md#GetAppliances) | **Get** /iam-deploy/appliances | 
[**GetIdP**](DefaultApi.md#GetIdP) | **Get** /iam-deploy/idp | 
[**GetIdPs**](DefaultApi.md#GetIdPs) | **Get** /iam-deploy/idps | 
[**SignOn**](DefaultApi.md#SignOn) | **Post** /iam-authn/sign-on | 
[**UpdateAppliance**](DefaultApi.md#UpdateAppliance) | **Put** /iam-deploy/appliance | 
[**UpdateIdP**](DefaultApi.md#UpdateIdP) | **Put** /iam-deploy/idp | 



## CreateAppliance

> CreateApplianceRes CreateAppliance(ctx).CreateApplianceReq(createApplianceReq).Execute()



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
    createApplianceReq := *openapiclient.NewCreateApplianceReq() // CreateApplianceReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAppliance(context.Background()).CreateApplianceReq(createApplianceReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppliance`: CreateApplianceRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAppliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApplianceReq** | [**CreateApplianceReq**](CreateApplianceReq.md) |  | 

### Return type

[**CreateApplianceRes**](CreateApplianceRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdP

> CreateIdPRes CreateIdP(ctx).CreateIdPReq(createIdPReq).Execute()



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
    createIdPReq := *openapiclient.NewCreateIdPReq() // CreateIdPReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIdP(context.Background()).CreateIdPReq(createIdPReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIdP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdP`: CreateIdPRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIdP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIdPReq** | [**CreateIdPReq**](CreateIdPReq.md) |  | 

### Return type

[**CreateIdPRes**](CreateIdPRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppliance

> DeleteApplianceRes DeleteAppliance(ctx).DeleteApplianceReq(deleteApplianceReq).Execute()



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
    deleteApplianceReq := *openapiclient.NewDeleteApplianceReq() // DeleteApplianceReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAppliance(context.Background()).DeleteApplianceReq(deleteApplianceReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppliance`: DeleteApplianceRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAppliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteApplianceReq** | [**DeleteApplianceReq**](DeleteApplianceReq.md) |  | 

### Return type

[**DeleteApplianceRes**](DeleteApplianceRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdP

> DeleteIdPRes DeleteIdP(ctx).DeleteIdPReq(deleteIdPReq).Execute()



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
    deleteIdPReq := *openapiclient.NewDeleteIdPReq() // DeleteIdPReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIdP(context.Background()).DeleteIdPReq(deleteIdPReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIdP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIdP`: DeleteIdPRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIdP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteIdPReq** | [**DeleteIdPReq**](DeleteIdPReq.md) |  | 

### Return type

[**DeleteIdPRes**](DeleteIdPRes.md)

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

> GetAppliancesRes GetAppliances(ctx).GetAppliancesReq(getAppliancesReq).Execute()



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
    getAppliancesReq := *openapiclient.NewGetAppliancesReq() // GetAppliancesReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAppliances(context.Background()).GetAppliancesReq(getAppliancesReq).Execute()
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
 **getAppliancesReq** | [**GetAppliancesReq**](GetAppliancesReq.md) |  | 

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

> GetIdPsRes GetIdPs(ctx).GetIdPsReq(getIdPsReq).Execute()



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
    getIdPsReq := *openapiclient.NewGetIdPsReq() // GetIdPsReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIdPs(context.Background()).GetIdPsReq(getIdPsReq).Execute()
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
 **getIdPsReq** | [**GetIdPsReq**](GetIdPsReq.md) |  | 

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

> UpdateApplianceRes UpdateAppliance(ctx).UpdateApplianceReq(updateApplianceReq).Execute()



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
    updateApplianceReq := *openapiclient.NewUpdateApplianceReq() // UpdateApplianceReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAppliance(context.Background()).UpdateApplianceReq(updateApplianceReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAppliance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppliance`: UpdateApplianceRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAppliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateApplianceReq** | [**UpdateApplianceReq**](UpdateApplianceReq.md) |  | 

### Return type

[**UpdateApplianceRes**](UpdateApplianceRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdP

> UpdateIdPRes UpdateIdP(ctx).UpdateIdPReq(updateIdPReq).Execute()



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
    updateIdPReq := *openapiclient.NewUpdateIdPReq() // UpdateIdPReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateIdP(context.Background()).UpdateIdPReq(updateIdPReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIdP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdP`: UpdateIdPRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIdP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateIdPReq** | [**UpdateIdPReq**](UpdateIdPReq.md) |  | 

### Return type

[**UpdateIdPRes**](UpdateIdPRes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

