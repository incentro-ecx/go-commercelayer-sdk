# \ResourceErrorsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETOrderIdResourceErrors**](ResourceErrorsApi.md#GETOrderIdResourceErrors) | **Get** /orders/{orderId}/resource_errors | Retrieve the resource errors associated to the order
[**GETResourceErrors**](ResourceErrorsApi.md#GETResourceErrors) | **Get** /resource_errors | List all resource errors
[**GETResourceErrorsResourceErrorId**](ResourceErrorsApi.md#GETResourceErrorsResourceErrorId) | **Get** /resource_errors/{resourceErrorId} | Retrieve a resource error
[**GETReturnIdResourceErrors**](ResourceErrorsApi.md#GETReturnIdResourceErrors) | **Get** /returns/{returnId}/resource_errors | Retrieve the resource errors associated to the return



## GETOrderIdResourceErrors

> GETOrderIdResourceErrors(ctx, orderId).Execute()

Retrieve the resource errors associated to the order



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
    orderId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceErrorsApi.GETOrderIdResourceErrors(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceErrorsApi.GETOrderIdResourceErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderIdResourceErrorsRequest struct via the builder pattern


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


## GETResourceErrors

> GETResourceErrors200Response GETResourceErrors(ctx).Execute()

List all resource errors



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
    resp, r, err := apiClient.ResourceErrorsApi.GETResourceErrors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceErrorsApi.GETResourceErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETResourceErrors`: GETResourceErrors200Response
    fmt.Fprintf(os.Stdout, "Response from `ResourceErrorsApi.GETResourceErrors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETResourceErrorsRequest struct via the builder pattern


### Return type

[**GETResourceErrors200Response**](GETResourceErrors200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETResourceErrorsResourceErrorId

> GETResourceErrorsResourceErrorId200Response GETResourceErrorsResourceErrorId(ctx, resourceErrorId).Execute()

Retrieve a resource error



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
    resourceErrorId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceErrorsApi.GETResourceErrorsResourceErrorId(context.Background(), resourceErrorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceErrorsApi.GETResourceErrorsResourceErrorId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETResourceErrorsResourceErrorId`: GETResourceErrorsResourceErrorId200Response
    fmt.Fprintf(os.Stdout, "Response from `ResourceErrorsApi.GETResourceErrorsResourceErrorId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceErrorId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETResourceErrorsResourceErrorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETResourceErrorsResourceErrorId200Response**](GETResourceErrorsResourceErrorId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETReturnIdResourceErrors

> GETReturnIdResourceErrors(ctx, returnId).Execute()

Retrieve the resource errors associated to the return



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
    returnId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceErrorsApi.GETReturnIdResourceErrors(context.Background(), returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceErrorsApi.GETReturnIdResourceErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnIdResourceErrorsRequest struct via the builder pattern


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

