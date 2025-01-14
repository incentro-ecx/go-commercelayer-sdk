# \RefundsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETCaptureIdRefunds**](RefundsApi.md#GETCaptureIdRefunds) | **Get** /captures/{captureId}/refunds | Retrieve the refunds associated to the capture
[**GETOrderIdRefunds**](RefundsApi.md#GETOrderIdRefunds) | **Get** /orders/{orderId}/refunds | Retrieve the refunds associated to the order
[**GETRefunds**](RefundsApi.md#GETRefunds) | **Get** /refunds | List all refunds
[**GETRefundsRefundId**](RefundsApi.md#GETRefundsRefundId) | **Get** /refunds/{refundId} | Retrieve a refund
[**GETReturnIdReferenceRefund**](RefundsApi.md#GETReturnIdReferenceRefund) | **Get** /returns/{returnId}/reference_refund | Retrieve the reference refund associated to the return
[**PATCHRefundsRefundId**](RefundsApi.md#PATCHRefundsRefundId) | **Patch** /refunds/{refundId} | Update a refund



## GETCaptureIdRefunds

> GETCaptureIdRefunds(ctx, captureId).Execute()

Retrieve the refunds associated to the capture



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
    r, err := apiClient.RefundsApi.GETCaptureIdRefunds(context.Background(), captureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.GETCaptureIdRefunds``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCaptureIdRefundsRequest struct via the builder pattern


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


## GETOrderIdRefunds

> GETOrderIdRefunds(ctx, orderId).Execute()

Retrieve the refunds associated to the order



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
    r, err := apiClient.RefundsApi.GETOrderIdRefunds(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.GETOrderIdRefunds``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdRefundsRequest struct via the builder pattern


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


## GETRefunds

> GETRefunds200Response GETRefunds(ctx).Execute()

List all refunds



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
    resp, r, err := apiClient.RefundsApi.GETRefunds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.GETRefunds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETRefunds`: GETRefunds200Response
    fmt.Fprintf(os.Stdout, "Response from `RefundsApi.GETRefunds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETRefundsRequest struct via the builder pattern


### Return type

[**GETRefunds200Response**](GETRefunds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETRefundsRefundId

> GETRefundsRefundId200Response GETRefundsRefundId(ctx, refundId).Execute()

Retrieve a refund



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
    resp, r, err := apiClient.RefundsApi.GETRefundsRefundId(context.Background(), refundId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.GETRefundsRefundId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETRefundsRefundId`: GETRefundsRefundId200Response
    fmt.Fprintf(os.Stdout, "Response from `RefundsApi.GETRefundsRefundId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refundId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETRefundsRefundIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETRefundsRefundId200Response**](GETRefundsRefundId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETReturnIdReferenceRefund

> GETReturnIdReferenceRefund(ctx, returnId).Execute()

Retrieve the reference refund associated to the return



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
    r, err := apiClient.RefundsApi.GETReturnIdReferenceRefund(context.Background(), returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.GETReturnIdReferenceRefund``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETReturnIdReferenceRefundRequest struct via the builder pattern


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


## PATCHRefundsRefundId

> PATCHRefundsRefundId200Response PATCHRefundsRefundId(ctx, refundId).RefundUpdate(refundUpdate).Execute()

Update a refund



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
    refundUpdate := *openapiclient.NewRefundUpdate(*openapiclient.NewRefundUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHRefundsRefundId200ResponseDataAttributes())) // RefundUpdate | 
    refundId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RefundsApi.PATCHRefundsRefundId(context.Background(), refundId).RefundUpdate(refundUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.PATCHRefundsRefundId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHRefundsRefundId`: PATCHRefundsRefundId200Response
    fmt.Fprintf(os.Stdout, "Response from `RefundsApi.PATCHRefundsRefundId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refundId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHRefundsRefundIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundUpdate** | [**RefundUpdate**](RefundUpdate.md) |  | 


### Return type

[**PATCHRefundsRefundId200Response**](PATCHRefundsRefundId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

