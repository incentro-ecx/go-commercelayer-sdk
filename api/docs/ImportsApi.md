# \ImportsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEImportsImportId**](ImportsApi.md#DELETEImportsImportId) | **Delete** /imports/{importId} | Delete an import
[**GETImports**](ImportsApi.md#GETImports) | **Get** /imports | List all imports
[**GETImportsImportId**](ImportsApi.md#GETImportsImportId) | **Get** /imports/{importId} | Retrieve an import
[**PATCHImportsImportId**](ImportsApi.md#PATCHImportsImportId) | **Patch** /imports/{importId} | Update an import
[**POSTImports**](ImportsApi.md#POSTImports) | **Post** /imports | Create an import



## DELETEImportsImportId

> DELETEImportsImportId(ctx, importId).Execute()

Delete an import



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
    importId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImportsApi.DELETEImportsImportId(context.Background(), importId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.DELETEImportsImportId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEImportsImportIdRequest struct via the builder pattern


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


## GETImports

> GETImports200Response GETImports(ctx).Execute()

List all imports



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
    resp, r, err := apiClient.ImportsApi.GETImports(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.GETImports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETImports`: GETImports200Response
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.GETImports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETImportsRequest struct via the builder pattern


### Return type

[**GETImports200Response**](GETImports200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETImportsImportId

> GETImportsImportId200Response GETImportsImportId(ctx, importId).Execute()

Retrieve an import



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
    importId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.GETImportsImportId(context.Background(), importId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.GETImportsImportId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETImportsImportId`: GETImportsImportId200Response
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.GETImportsImportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETImportsImportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETImportsImportId200Response**](GETImportsImportId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHImportsImportId

> PATCHImportsImportId200Response PATCHImportsImportId(ctx, importId).ImportUpdate(importUpdate).Execute()

Update an import



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
    importUpdate := *openapiclient.NewImportUpdate(*openapiclient.NewImportUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCleanupsCleanupId200ResponseDataAttributes())) // ImportUpdate | 
    importId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.PATCHImportsImportId(context.Background(), importId).ImportUpdate(importUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.PATCHImportsImportId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHImportsImportId`: PATCHImportsImportId200Response
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.PATCHImportsImportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHImportsImportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importUpdate** | [**ImportUpdate**](ImportUpdate.md) |  | 


### Return type

[**PATCHImportsImportId200Response**](PATCHImportsImportId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTImports

> POSTImports201Response POSTImports(ctx).ImportCreate(importCreate).Execute()

Create an import



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
    importCreate := *openapiclient.NewImportCreate(*openapiclient.NewImportCreateData(interface{}(123), *openapiclient.NewPOSTImports201ResponseDataAttributes(interface{}(skus), interface{}([{code=ABC, name=Foo}, {code=DEF, name=Bar}])))) // ImportCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.POSTImports(context.Background()).ImportCreate(importCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.POSTImports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTImports`: POSTImports201Response
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.POSTImports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importCreate** | [**ImportCreate**](ImportCreate.md) |  | 

### Return type

[**POSTImports201Response**](POSTImports201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

