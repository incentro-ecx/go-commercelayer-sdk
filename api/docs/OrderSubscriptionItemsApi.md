# \OrderSubscriptionItemsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEOrderSubscriptionItemsOrderSubscriptionItemId**](OrderSubscriptionItemsApi.md#DELETEOrderSubscriptionItemsOrderSubscriptionItemId) | **Delete** /order_subscription_items/{orderSubscriptionItemId} | Delete an order subscription item
[**GETOrderSubscriptionIdOrderSubscriptionItems**](OrderSubscriptionItemsApi.md#GETOrderSubscriptionIdOrderSubscriptionItems) | **Get** /order_subscriptions/{orderSubscriptionId}/order_subscription_items | Retrieve the order subscription items associated to the order subscription
[**GETOrderSubscriptionItems**](OrderSubscriptionItemsApi.md#GETOrderSubscriptionItems) | **Get** /order_subscription_items | List all order subscription items
[**GETOrderSubscriptionItemsOrderSubscriptionItemId**](OrderSubscriptionItemsApi.md#GETOrderSubscriptionItemsOrderSubscriptionItemId) | **Get** /order_subscription_items/{orderSubscriptionItemId} | Retrieve an order subscription item
[**PATCHOrderSubscriptionItemsOrderSubscriptionItemId**](OrderSubscriptionItemsApi.md#PATCHOrderSubscriptionItemsOrderSubscriptionItemId) | **Patch** /order_subscription_items/{orderSubscriptionItemId} | Update an order subscription item
[**POSTOrderSubscriptionItems**](OrderSubscriptionItemsApi.md#POSTOrderSubscriptionItems) | **Post** /order_subscription_items | Create an order subscription item



## DELETEOrderSubscriptionItemsOrderSubscriptionItemId

> DELETEOrderSubscriptionItemsOrderSubscriptionItemId(ctx, orderSubscriptionItemId).Execute()

Delete an order subscription item



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
    orderSubscriptionItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderSubscriptionItemsApi.DELETEOrderSubscriptionItemsOrderSubscriptionItemId(context.Background(), orderSubscriptionItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionItemsApi.DELETEOrderSubscriptionItemsOrderSubscriptionItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderSubscriptionItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEOrderSubscriptionItemsOrderSubscriptionItemIdRequest struct via the builder pattern


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


## GETOrderSubscriptionIdOrderSubscriptionItems

> GETOrderSubscriptionIdOrderSubscriptionItems(ctx, orderSubscriptionId).Execute()

Retrieve the order subscription items associated to the order subscription



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
    orderSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderSubscriptionItemsApi.GETOrderSubscriptionIdOrderSubscriptionItems(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionItemsApi.GETOrderSubscriptionIdOrderSubscriptionItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderSubscriptionIdOrderSubscriptionItemsRequest struct via the builder pattern


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


## GETOrderSubscriptionItems

> GETOrderSubscriptionItems200Response GETOrderSubscriptionItems(ctx).Execute()

List all order subscription items



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
    resp, r, err := apiClient.OrderSubscriptionItemsApi.GETOrderSubscriptionItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionItemsApi.GETOrderSubscriptionItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderSubscriptionItems`: GETOrderSubscriptionItems200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderSubscriptionItemsApi.GETOrderSubscriptionItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderSubscriptionItemsRequest struct via the builder pattern


### Return type

[**GETOrderSubscriptionItems200Response**](GETOrderSubscriptionItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderSubscriptionItemsOrderSubscriptionItemId

> GETOrderSubscriptionItemsOrderSubscriptionItemId200Response GETOrderSubscriptionItemsOrderSubscriptionItemId(ctx, orderSubscriptionItemId).Execute()

Retrieve an order subscription item



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
    orderSubscriptionItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderSubscriptionItemsApi.GETOrderSubscriptionItemsOrderSubscriptionItemId(context.Background(), orderSubscriptionItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionItemsApi.GETOrderSubscriptionItemsOrderSubscriptionItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderSubscriptionItemsOrderSubscriptionItemId`: GETOrderSubscriptionItemsOrderSubscriptionItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderSubscriptionItemsApi.GETOrderSubscriptionItemsOrderSubscriptionItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderSubscriptionItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderSubscriptionItemsOrderSubscriptionItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETOrderSubscriptionItemsOrderSubscriptionItemId200Response**](GETOrderSubscriptionItemsOrderSubscriptionItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHOrderSubscriptionItemsOrderSubscriptionItemId

> PATCHOrderSubscriptionItemsOrderSubscriptionItemId200Response PATCHOrderSubscriptionItemsOrderSubscriptionItemId(ctx, orderSubscriptionItemId).OrderSubscriptionItemUpdate(orderSubscriptionItemUpdate).Execute()

Update an order subscription item



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
    orderSubscriptionItemUpdate := *openapiclient.NewOrderSubscriptionItemUpdate(*openapiclient.NewOrderSubscriptionItemUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes())) // OrderSubscriptionItemUpdate | 
    orderSubscriptionItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderSubscriptionItemsApi.PATCHOrderSubscriptionItemsOrderSubscriptionItemId(context.Background(), orderSubscriptionItemId).OrderSubscriptionItemUpdate(orderSubscriptionItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionItemsApi.PATCHOrderSubscriptionItemsOrderSubscriptionItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHOrderSubscriptionItemsOrderSubscriptionItemId`: PATCHOrderSubscriptionItemsOrderSubscriptionItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderSubscriptionItemsApi.PATCHOrderSubscriptionItemsOrderSubscriptionItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderSubscriptionItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHOrderSubscriptionItemsOrderSubscriptionItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderSubscriptionItemUpdate** | [**OrderSubscriptionItemUpdate**](OrderSubscriptionItemUpdate.md) |  | 


### Return type

[**PATCHOrderSubscriptionItemsOrderSubscriptionItemId200Response**](PATCHOrderSubscriptionItemsOrderSubscriptionItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTOrderSubscriptionItems

> POSTOrderSubscriptionItems201Response POSTOrderSubscriptionItems(ctx).OrderSubscriptionItemCreate(orderSubscriptionItemCreate).Execute()

Create an order subscription item



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
    orderSubscriptionItemCreate := *openapiclient.NewOrderSubscriptionItemCreate(*openapiclient.NewOrderSubscriptionItemCreateData(interface{}(123), *openapiclient.NewPOSTOrderSubscriptionItems201ResponseDataAttributes(interface{}(4)))) // OrderSubscriptionItemCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderSubscriptionItemsApi.POSTOrderSubscriptionItems(context.Background()).OrderSubscriptionItemCreate(orderSubscriptionItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionItemsApi.POSTOrderSubscriptionItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTOrderSubscriptionItems`: POSTOrderSubscriptionItems201Response
    fmt.Fprintf(os.Stdout, "Response from `OrderSubscriptionItemsApi.POSTOrderSubscriptionItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTOrderSubscriptionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderSubscriptionItemCreate** | [**OrderSubscriptionItemCreate**](OrderSubscriptionItemCreate.md) |  | 

### Return type

[**POSTOrderSubscriptionItems201Response**](POSTOrderSubscriptionItems201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

