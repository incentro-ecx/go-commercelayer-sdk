# \CapturesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETAuthorizationIdCaptures**](CapturesApi.md#GETAuthorizationIdCaptures) | **Get** /authorizations/{authorizationId}/captures | Retrieve the captures associated to the authorization
[**GETCaptures**](CapturesApi.md#GETCaptures) | **Get** /captures | List all captures
[**GETCapturesCaptureId**](CapturesApi.md#GETCapturesCaptureId) | **Get** /captures/{captureId} | Retrieve a capture
[**GETOrderIdCaptures**](CapturesApi.md#GETOrderIdCaptures) | **Get** /orders/{orderId}/captures | Retrieve the captures associated to the order
[**GETRefundIdReferenceCapture**](CapturesApi.md#GETRefundIdReferenceCapture) | **Get** /refunds/{refundId}/reference_capture | Retrieve the reference capture associated to the refund
[**GETReturnIdReferenceCapture**](CapturesApi.md#GETReturnIdReferenceCapture) | **Get** /returns/{returnId}/reference_capture | Retrieve the reference capture associated to the return
[**PATCHCapturesCaptureId**](CapturesApi.md#PATCHCapturesCaptureId) | **Patch** /captures/{captureId} | Update a capture



## GETAuthorizationIdCaptures

> GETAuthorizationIdCaptures(ctx, authorizationId).Execute()

Retrieve the captures associated to the authorization



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
    r, err := apiClient.CapturesApi.GETAuthorizationIdCaptures(context.Background(), authorizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.GETAuthorizationIdCaptures``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETAuthorizationIdCapturesRequest struct via the builder pattern


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


## GETCaptures

> GETCaptures200Response GETCaptures(ctx).Execute()

List all captures



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
    resp, r, err := apiClient.CapturesApi.GETCaptures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.GETCaptures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCaptures`: GETCaptures200Response
    fmt.Fprintf(os.Stdout, "Response from `CapturesApi.GETCaptures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCapturesRequest struct via the builder pattern


### Return type

[**GETCaptures200Response**](GETCaptures200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCapturesCaptureId

> GETCapturesCaptureId200Response GETCapturesCaptureId(ctx, captureId).Execute()

Retrieve a capture



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
    resp, r, err := apiClient.CapturesApi.GETCapturesCaptureId(context.Background(), captureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.GETCapturesCaptureId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCapturesCaptureId`: GETCapturesCaptureId200Response
    fmt.Fprintf(os.Stdout, "Response from `CapturesApi.GETCapturesCaptureId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captureId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCapturesCaptureIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCapturesCaptureId200Response**](GETCapturesCaptureId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdCaptures

> GETOrderIdCaptures(ctx, orderId).Execute()

Retrieve the captures associated to the order



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
    r, err := apiClient.CapturesApi.GETOrderIdCaptures(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.GETOrderIdCaptures``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdCapturesRequest struct via the builder pattern


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


## GETRefundIdReferenceCapture

> GETRefundIdReferenceCapture(ctx, refundId).Execute()

Retrieve the reference capture associated to the refund



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
    refundId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CapturesApi.GETRefundIdReferenceCapture(context.Background(), refundId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.GETRefundIdReferenceCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refundId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETRefundIdReferenceCaptureRequest struct via the builder pattern


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


## GETReturnIdReferenceCapture

> GETReturnIdReferenceCapture(ctx, returnId).Execute()

Retrieve the reference capture associated to the return



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
    r, err := apiClient.CapturesApi.GETReturnIdReferenceCapture(context.Background(), returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.GETReturnIdReferenceCapture``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETReturnIdReferenceCaptureRequest struct via the builder pattern


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


## PATCHCapturesCaptureId

> PATCHCapturesCaptureId200Response PATCHCapturesCaptureId(ctx, captureId).CaptureUpdate(captureUpdate).Execute()

Update a capture



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
    captureUpdate := *openapiclient.NewCaptureUpdate(*openapiclient.NewCaptureUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCapturesCaptureId200ResponseDataAttributes())) // CaptureUpdate | 
    captureId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapturesApi.PATCHCapturesCaptureId(context.Background(), captureId).CaptureUpdate(captureUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.PATCHCapturesCaptureId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCapturesCaptureId`: PATCHCapturesCaptureId200Response
    fmt.Fprintf(os.Stdout, "Response from `CapturesApi.PATCHCapturesCaptureId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captureId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCapturesCaptureIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captureUpdate** | [**CaptureUpdate**](CaptureUpdate.md) |  | 


### Return type

[**PATCHCapturesCaptureId200Response**](PATCHCapturesCaptureId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

