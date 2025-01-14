# \VoidsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETAuthorizationIdVoids**](VoidsApi.md#GETAuthorizationIdVoids) | **Get** /authorizations/{authorizationId}/voids | Retrieve the voids associated to the authorization
[**GETOrderIdVoids**](VoidsApi.md#GETOrderIdVoids) | **Get** /orders/{orderId}/voids | Retrieve the voids associated to the order
[**GETVoids**](VoidsApi.md#GETVoids) | **Get** /voids | List all voids
[**GETVoidsVoidId**](VoidsApi.md#GETVoidsVoidId) | **Get** /voids/{voidId} | Retrieve a void
[**PATCHVoidsVoidId**](VoidsApi.md#PATCHVoidsVoidId) | **Patch** /voids/{voidId} | Update a void



## GETAuthorizationIdVoids

> GETAuthorizationIdVoids(ctx, authorizationId).Execute()

Retrieve the voids associated to the authorization



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
    authorizationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VoidsApi.GETAuthorizationIdVoids(context.Background(), authorizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoidsApi.GETAuthorizationIdVoids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAuthorizationIdVoidsRequest struct via the builder pattern


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


## GETOrderIdVoids

> GETOrderIdVoids(ctx, orderId).Execute()

Retrieve the voids associated to the order



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
    r, err := apiClient.VoidsApi.GETOrderIdVoids(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoidsApi.GETOrderIdVoids``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdVoidsRequest struct via the builder pattern


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


## GETVoids

> GETVoids200Response GETVoids(ctx).Execute()

List all voids



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
    resp, r, err := apiClient.VoidsApi.GETVoids(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoidsApi.GETVoids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETVoids`: GETVoids200Response
    fmt.Fprintf(os.Stdout, "Response from `VoidsApi.GETVoids`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETVoidsRequest struct via the builder pattern


### Return type

[**GETVoids200Response**](GETVoids200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETVoidsVoidId

> GETVoidsVoidId200Response GETVoidsVoidId(ctx, voidId).Execute()

Retrieve a void



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
    voidId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoidsApi.GETVoidsVoidId(context.Background(), voidId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoidsApi.GETVoidsVoidId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETVoidsVoidId`: GETVoidsVoidId200Response
    fmt.Fprintf(os.Stdout, "Response from `VoidsApi.GETVoidsVoidId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voidId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETVoidsVoidIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETVoidsVoidId200Response**](GETVoidsVoidId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHVoidsVoidId

> PATCHVoidsVoidId200Response PATCHVoidsVoidId(ctx, voidId).VoidUpdate(voidUpdate).Execute()

Update a void



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
    voidUpdate := *openapiclient.NewVoidUpdate(*openapiclient.NewVoidUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHRefundsRefundId200ResponseDataAttributes())) // VoidUpdate | 
    voidId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoidsApi.PATCHVoidsVoidId(context.Background(), voidId).VoidUpdate(voidUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoidsApi.PATCHVoidsVoidId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHVoidsVoidId`: PATCHVoidsVoidId200Response
    fmt.Fprintf(os.Stdout, "Response from `VoidsApi.PATCHVoidsVoidId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voidId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHVoidsVoidIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voidUpdate** | [**VoidUpdate**](VoidUpdate.md) |  | 


### Return type

[**PATCHVoidsVoidId200Response**](PATCHVoidsVoidId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

