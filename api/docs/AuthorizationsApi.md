# \AuthorizationsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETAuthorizations**](AuthorizationsApi.md#GETAuthorizations) | **Get** /authorizations | List all authorizations
[**GETAuthorizationsAuthorizationId**](AuthorizationsApi.md#GETAuthorizationsAuthorizationId) | **Get** /authorizations/{authorizationId} | Retrieve an authorization
[**GETCaptureIdReferenceAuthorization**](AuthorizationsApi.md#GETCaptureIdReferenceAuthorization) | **Get** /captures/{captureId}/reference_authorization | Retrieve the reference authorization associated to the capture
[**GETOrderIdAuthorizations**](AuthorizationsApi.md#GETOrderIdAuthorizations) | **Get** /orders/{orderId}/authorizations | Retrieve the authorizations associated to the order
[**GETVoidIdReferenceAuthorization**](AuthorizationsApi.md#GETVoidIdReferenceAuthorization) | **Get** /voids/{voidId}/reference_authorization | Retrieve the reference authorization associated to the void
[**PATCHAuthorizationsAuthorizationId**](AuthorizationsApi.md#PATCHAuthorizationsAuthorizationId) | **Patch** /authorizations/{authorizationId} | Update an authorization



## GETAuthorizations

> GETAuthorizations200Response GETAuthorizations(ctx).Execute()

List all authorizations



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
    resp, r, err := apiClient.AuthorizationsApi.GETAuthorizations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.GETAuthorizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAuthorizations`: GETAuthorizations200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationsApi.GETAuthorizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAuthorizationsRequest struct via the builder pattern


### Return type

[**GETAuthorizations200Response**](GETAuthorizations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETAuthorizationsAuthorizationId

> GETAuthorizationsAuthorizationId200Response GETAuthorizationsAuthorizationId(ctx, authorizationId).Execute()

Retrieve an authorization



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
    resp, r, err := apiClient.AuthorizationsApi.GETAuthorizationsAuthorizationId(context.Background(), authorizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.GETAuthorizationsAuthorizationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAuthorizationsAuthorizationId`: GETAuthorizationsAuthorizationId200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationsApi.GETAuthorizationsAuthorizationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAuthorizationsAuthorizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETAuthorizationsAuthorizationId200Response**](GETAuthorizationsAuthorizationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCaptureIdReferenceAuthorization

> GETCaptureIdReferenceAuthorization(ctx, captureId).Execute()

Retrieve the reference authorization associated to the capture



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
    captureId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationsApi.GETCaptureIdReferenceAuthorization(context.Background(), captureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.GETCaptureIdReferenceAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captureId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCaptureIdReferenceAuthorizationRequest struct via the builder pattern


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


## GETOrderIdAuthorizations

> GETOrderIdAuthorizations(ctx, orderId).Execute()

Retrieve the authorizations associated to the order



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
    r, err := apiClient.AuthorizationsApi.GETOrderIdAuthorizations(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.GETOrderIdAuthorizations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdAuthorizationsRequest struct via the builder pattern


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


## GETVoidIdReferenceAuthorization

> GETVoidIdReferenceAuthorization(ctx, voidId).Execute()

Retrieve the reference authorization associated to the void



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
    r, err := apiClient.AuthorizationsApi.GETVoidIdReferenceAuthorization(context.Background(), voidId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.GETVoidIdReferenceAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voidId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETVoidIdReferenceAuthorizationRequest struct via the builder pattern


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


## PATCHAuthorizationsAuthorizationId

> PATCHAuthorizationsAuthorizationId200Response PATCHAuthorizationsAuthorizationId(ctx, authorizationId).AuthorizationUpdate(authorizationUpdate).Execute()

Update an authorization



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
    authorizationUpdate := *openapiclient.NewAuthorizationUpdate(*openapiclient.NewAuthorizationUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHAuthorizationsAuthorizationId200ResponseDataAttributes())) // AuthorizationUpdate | 
    authorizationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationsApi.PATCHAuthorizationsAuthorizationId(context.Background(), authorizationId).AuthorizationUpdate(authorizationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationsApi.PATCHAuthorizationsAuthorizationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHAuthorizationsAuthorizationId`: PATCHAuthorizationsAuthorizationId200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationsApi.PATCHAuthorizationsAuthorizationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAuthorizationsAuthorizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationUpdate** | [**AuthorizationUpdate**](AuthorizationUpdate.md) |  | 


### Return type

[**PATCHAuthorizationsAuthorizationId200Response**](PATCHAuthorizationsAuthorizationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

