# \AdjustmentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAdjustmentsAdjustmentId**](AdjustmentsApi.md#DELETEAdjustmentsAdjustmentId) | **Delete** /adjustments/{adjustmentId} | Delete an adjustment
[**GETAdjustments**](AdjustmentsApi.md#GETAdjustments) | **Get** /adjustments | List all adjustments
[**GETAdjustmentsAdjustmentId**](AdjustmentsApi.md#GETAdjustmentsAdjustmentId) | **Get** /adjustments/{adjustmentId} | Retrieve an adjustment
[**PATCHAdjustmentsAdjustmentId**](AdjustmentsApi.md#PATCHAdjustmentsAdjustmentId) | **Patch** /adjustments/{adjustmentId} | Update an adjustment
[**POSTAdjustments**](AdjustmentsApi.md#POSTAdjustments) | **Post** /adjustments | Create an adjustment



## DELETEAdjustmentsAdjustmentId

> DELETEAdjustmentsAdjustmentId(ctx, adjustmentId).Execute()

Delete an adjustment



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
    adjustmentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdjustmentsApi.DELETEAdjustmentsAdjustmentId(context.Background(), adjustmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.DELETEAdjustmentsAdjustmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adjustmentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEAdjustmentsAdjustmentIdRequest struct via the builder pattern


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


## GETAdjustments

> GETAdjustments200Response GETAdjustments(ctx).Execute()

List all adjustments



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
    resp, r, err := apiClient.AdjustmentsApi.GETAdjustments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.GETAdjustments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAdjustments`: GETAdjustments200Response
    fmt.Fprintf(os.Stdout, "Response from `AdjustmentsApi.GETAdjustments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdjustmentsRequest struct via the builder pattern


### Return type

[**GETAdjustments200Response**](GETAdjustments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETAdjustmentsAdjustmentId

> GETAdjustmentsAdjustmentId200Response GETAdjustmentsAdjustmentId(ctx, adjustmentId).Execute()

Retrieve an adjustment



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
    adjustmentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsApi.GETAdjustmentsAdjustmentId(context.Background(), adjustmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.GETAdjustmentsAdjustmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAdjustmentsAdjustmentId`: GETAdjustmentsAdjustmentId200Response
    fmt.Fprintf(os.Stdout, "Response from `AdjustmentsApi.GETAdjustmentsAdjustmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adjustmentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdjustmentsAdjustmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETAdjustmentsAdjustmentId200Response**](GETAdjustmentsAdjustmentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHAdjustmentsAdjustmentId

> PATCHAdjustmentsAdjustmentId200Response PATCHAdjustmentsAdjustmentId(ctx, adjustmentId).AdjustmentUpdate(adjustmentUpdate).Execute()

Update an adjustment



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
    adjustmentUpdate := *openapiclient.NewAdjustmentUpdate(*openapiclient.NewAdjustmentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHAdjustmentsAdjustmentId200ResponseDataAttributes())) // AdjustmentUpdate | 
    adjustmentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsApi.PATCHAdjustmentsAdjustmentId(context.Background(), adjustmentId).AdjustmentUpdate(adjustmentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.PATCHAdjustmentsAdjustmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHAdjustmentsAdjustmentId`: PATCHAdjustmentsAdjustmentId200Response
    fmt.Fprintf(os.Stdout, "Response from `AdjustmentsApi.PATCHAdjustmentsAdjustmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adjustmentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAdjustmentsAdjustmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adjustmentUpdate** | [**AdjustmentUpdate**](AdjustmentUpdate.md) |  | 


### Return type

[**PATCHAdjustmentsAdjustmentId200Response**](PATCHAdjustmentsAdjustmentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTAdjustments

> POSTAdjustments201Response POSTAdjustments(ctx).AdjustmentCreate(adjustmentCreate).Execute()

Create an adjustment



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
    adjustmentCreate := *openapiclient.NewAdjustmentCreate(*openapiclient.NewAdjustmentCreateData(interface{}(123), *openapiclient.NewPOSTAdjustments201ResponseDataAttributes(interface{}(Additional service), interface{}(EUR), interface{}(1500)))) // AdjustmentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsApi.POSTAdjustments(context.Background()).AdjustmentCreate(adjustmentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsApi.POSTAdjustments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTAdjustments`: POSTAdjustments201Response
    fmt.Fprintf(os.Stdout, "Response from `AdjustmentsApi.POSTAdjustments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTAdjustmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adjustmentCreate** | [**AdjustmentCreate**](AdjustmentCreate.md) |  | 

### Return type

[**POSTAdjustments201Response**](POSTAdjustments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

