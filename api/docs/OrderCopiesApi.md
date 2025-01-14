# \OrderCopiesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEOrderCopiesOrderCopyId**](OrderCopiesApi.md#DELETEOrderCopiesOrderCopyId) | **Delete** /order_copies/{orderCopyId} | Delete an order copy
[**GETOrderCopies**](OrderCopiesApi.md#GETOrderCopies) | **Get** /order_copies | List all order copies
[**GETOrderCopiesOrderCopyId**](OrderCopiesApi.md#GETOrderCopiesOrderCopyId) | **Get** /order_copies/{orderCopyId} | Retrieve an order copy
[**GETOrderIdOrderCopies**](OrderCopiesApi.md#GETOrderIdOrderCopies) | **Get** /orders/{orderId}/order_copies | Retrieve the order copies associated to the order
[**PATCHOrderCopiesOrderCopyId**](OrderCopiesApi.md#PATCHOrderCopiesOrderCopyId) | **Patch** /order_copies/{orderCopyId} | Update an order copy
[**POSTOrderCopies**](OrderCopiesApi.md#POSTOrderCopies) | **Post** /order_copies | Create an order copy



## DELETEOrderCopiesOrderCopyId

> DELETEOrderCopiesOrderCopyId(ctx, orderCopyId).Execute()

Delete an order copy



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
    orderCopyId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderCopiesApi.DELETEOrderCopiesOrderCopyId(context.Background(), orderCopyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderCopiesApi.DELETEOrderCopiesOrderCopyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderCopyId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEOrderCopiesOrderCopyIdRequest struct via the builder pattern


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


## GETOrderCopies

> GETOrderCopies200Response GETOrderCopies(ctx).Execute()

List all order copies



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
    resp, r, err := apiClient.OrderCopiesApi.GETOrderCopies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderCopiesApi.GETOrderCopies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderCopies`: GETOrderCopies200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderCopiesApi.GETOrderCopies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderCopiesRequest struct via the builder pattern


### Return type

[**GETOrderCopies200Response**](GETOrderCopies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderCopiesOrderCopyId

> GETOrderCopiesOrderCopyId200Response GETOrderCopiesOrderCopyId(ctx, orderCopyId).Execute()

Retrieve an order copy



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
    orderCopyId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderCopiesApi.GETOrderCopiesOrderCopyId(context.Background(), orderCopyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderCopiesApi.GETOrderCopiesOrderCopyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderCopiesOrderCopyId`: GETOrderCopiesOrderCopyId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderCopiesApi.GETOrderCopiesOrderCopyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderCopyId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderCopiesOrderCopyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETOrderCopiesOrderCopyId200Response**](GETOrderCopiesOrderCopyId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdOrderCopies

> GETOrderIdOrderCopies(ctx, orderId).Execute()

Retrieve the order copies associated to the order



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
    r, err := apiClient.OrderCopiesApi.GETOrderIdOrderCopies(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderCopiesApi.GETOrderIdOrderCopies``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdOrderCopiesRequest struct via the builder pattern


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


## PATCHOrderCopiesOrderCopyId

> PATCHOrderCopiesOrderCopyId200Response PATCHOrderCopiesOrderCopyId(ctx, orderCopyId).OrderCopyUpdate(orderCopyUpdate).Execute()

Update an order copy



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
    orderCopyUpdate := *openapiclient.NewOrderCopyUpdate(*openapiclient.NewOrderCopyUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes())) // OrderCopyUpdate | 
    orderCopyId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderCopiesApi.PATCHOrderCopiesOrderCopyId(context.Background(), orderCopyId).OrderCopyUpdate(orderCopyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderCopiesApi.PATCHOrderCopiesOrderCopyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHOrderCopiesOrderCopyId`: PATCHOrderCopiesOrderCopyId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderCopiesApi.PATCHOrderCopiesOrderCopyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderCopyId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHOrderCopiesOrderCopyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCopyUpdate** | [**OrderCopyUpdate**](OrderCopyUpdate.md) |  | 


### Return type

[**PATCHOrderCopiesOrderCopyId200Response**](PATCHOrderCopiesOrderCopyId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTOrderCopies

> POSTOrderCopies201Response POSTOrderCopies(ctx).OrderCopyCreate(orderCopyCreate).Execute()

Create an order copy



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
    orderCopyCreate := *openapiclient.NewOrderCopyCreate(*openapiclient.NewOrderCopyCreateData(interface{}(123), *openapiclient.NewPOSTOrderCopies201ResponseDataAttributes())) // OrderCopyCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderCopiesApi.POSTOrderCopies(context.Background()).OrderCopyCreate(orderCopyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderCopiesApi.POSTOrderCopies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTOrderCopies`: POSTOrderCopies201Response
    fmt.Fprintf(os.Stdout, "Response from `OrderCopiesApi.POSTOrderCopies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTOrderCopiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCopyCreate** | [**OrderCopyCreate**](OrderCopyCreate.md) |  | 

### Return type

[**POSTOrderCopies201Response**](POSTOrderCopies201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

