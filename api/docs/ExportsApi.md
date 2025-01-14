# \ExportsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExportsExportId**](ExportsApi.md#DELETEExportsExportId) | **Delete** /exports/{exportId} | Delete an export
[**GETExports**](ExportsApi.md#GETExports) | **Get** /exports | List all exports
[**GETExportsExportId**](ExportsApi.md#GETExportsExportId) | **Get** /exports/{exportId} | Retrieve an export
[**PATCHExportsExportId**](ExportsApi.md#PATCHExportsExportId) | **Patch** /exports/{exportId} | Update an export
[**POSTExports**](ExportsApi.md#POSTExports) | **Post** /exports | Create an export



## DELETEExportsExportId

> DELETEExportsExportId(ctx, exportId).Execute()

Delete an export



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {
    exportId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExportsApi.DELETEExportsExportId(context.Background(), exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.DELETEExportsExportId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExportsExportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExports

> GETExports200Response GETExports(ctx).Execute()

List all exports



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportsApi.GETExports(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.GETExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExports`: GETExports200Response
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.GETExports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExportsRequest struct via the builder pattern


### Return type

[**GETExports200Response**](GETExports200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExportsExportId

> GETExportsExportId200Response GETExportsExportId(ctx, exportId).Execute()

Retrieve an export



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {
    exportId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportsApi.GETExportsExportId(context.Background(), exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.GETExportsExportId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExportsExportId`: GETExportsExportId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.GETExportsExportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExportsExportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETExportsExportId200Response**](GETExportsExportId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHExportsExportId

> PATCHExportsExportId200Response PATCHExportsExportId(ctx, exportId).ExportUpdate(exportUpdate).Execute()

Update an export



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {
    exportUpdate := *openapiclient.NewExportUpdate(*openapiclient.NewExportUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCleanupsCleanupId200ResponseDataAttributes())) // ExportUpdate | 
    exportId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportsApi.PATCHExportsExportId(context.Background(), exportId).ExportUpdate(exportUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.PATCHExportsExportId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHExportsExportId`: PATCHExportsExportId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.PATCHExportsExportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHExportsExportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportUpdate** | [**ExportUpdate**](ExportUpdate.md) |  | 


### Return type

[**PATCHExportsExportId200Response**](PATCHExportsExportId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTExports

> POSTExports201Response POSTExports(ctx).ExportCreate(exportCreate).Execute()

Create an export



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {
    exportCreate := *openapiclient.NewExportCreate(*openapiclient.NewExportCreateData(interface{}(123), *openapiclient.NewPOSTExports201ResponseDataAttributes(interface{}(skus)))) // ExportCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportsApi.POSTExports(context.Background()).ExportCreate(exportCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.POSTExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTExports`: POSTExports201Response
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.POSTExports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportCreate** | [**ExportCreate**](ExportCreate.md) |  | 

### Return type

[**POSTExports201Response**](POSTExports201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

