# \OrderSubscriptionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEOrderSubscriptionsOrderSubscriptionId**](OrderSubscriptionsApi.md#DELETEOrderSubscriptionsOrderSubscriptionId) | **Delete** /order_subscriptions/{orderSubscriptionId} | Delete an order subscription
[**GETCustomerIdOrderSubscriptions**](OrderSubscriptionsApi.md#GETCustomerIdOrderSubscriptions) | **Get** /customers/{customerId}/order_subscriptions | Retrieve the order subscriptions associated to the customer
[**GETOrderCopyIdOrderSubscription**](OrderSubscriptionsApi.md#GETOrderCopyIdOrderSubscription) | **Get** /order_copies/{orderCopyId}/order_subscription | Retrieve the order subscription associated to the order copy
[**GETOrderIdOrderSubscription**](OrderSubscriptionsApi.md#GETOrderIdOrderSubscription) | **Get** /orders/{orderId}/order_subscription | Retrieve the order subscription associated to the order
[**GETOrderIdOrderSubscriptions**](OrderSubscriptionsApi.md#GETOrderIdOrderSubscriptions) | **Get** /orders/{orderId}/order_subscriptions | Retrieve the order subscriptions associated to the order
[**GETOrderSubscriptionItemIdOrderSubscription**](OrderSubscriptionsApi.md#GETOrderSubscriptionItemIdOrderSubscription) | **Get** /order_subscription_items/{orderSubscriptionItemId}/order_subscription | Retrieve the order subscription associated to the order subscription item
[**GETOrderSubscriptions**](OrderSubscriptionsApi.md#GETOrderSubscriptions) | **Get** /order_subscriptions | List all order subscriptions
[**GETOrderSubscriptionsOrderSubscriptionId**](OrderSubscriptionsApi.md#GETOrderSubscriptionsOrderSubscriptionId) | **Get** /order_subscriptions/{orderSubscriptionId} | Retrieve an order subscription
[**GETRecurringOrderCopyIdOrderSubscription**](OrderSubscriptionsApi.md#GETRecurringOrderCopyIdOrderSubscription) | **Get** /recurring_order_copies/{recurringOrderCopyId}/order_subscription | Retrieve the order subscription associated to the recurring order copy
[**GETSubscriptionModelIdOrderSubscriptions**](OrderSubscriptionsApi.md#GETSubscriptionModelIdOrderSubscriptions) | **Get** /subscription_models/{subscriptionModelId}/order_subscriptions | Retrieve the order subscriptions associated to the subscription model
[**PATCHOrderSubscriptionsOrderSubscriptionId**](OrderSubscriptionsApi.md#PATCHOrderSubscriptionsOrderSubscriptionId) | **Patch** /order_subscriptions/{orderSubscriptionId} | Update an order subscription
[**POSTOrderSubscriptions**](OrderSubscriptionsApi.md#POSTOrderSubscriptions) | **Post** /order_subscriptions | Create an order subscription



## DELETEOrderSubscriptionsOrderSubscriptionId

> DELETEOrderSubscriptionsOrderSubscriptionId(ctx, orderSubscriptionId).Execute()

Delete an order subscription



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
    r, err := apiClient.OrderSubscriptionsApi.DELETEOrderSubscriptionsOrderSubscriptionId(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.DELETEOrderSubscriptionsOrderSubscriptionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEOrderSubscriptionsOrderSubscriptionIdRequest struct via the builder pattern


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


## GETCustomerIdOrderSubscriptions

> GETCustomerIdOrderSubscriptions(ctx, customerId).Execute()

Retrieve the order subscriptions associated to the customer



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
    customerId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderSubscriptionsApi.GETCustomerIdOrderSubscriptions(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETCustomerIdOrderSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerIdOrderSubscriptionsRequest struct via the builder pattern


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


## GETOrderCopyIdOrderSubscription

> GETOrderCopyIdOrderSubscription(ctx, orderCopyId).Execute()

Retrieve the order subscription associated to the order copy



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
    r, err := apiClient.OrderSubscriptionsApi.GETOrderCopyIdOrderSubscription(context.Background(), orderCopyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETOrderCopyIdOrderSubscription``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderCopyIdOrderSubscriptionRequest struct via the builder pattern


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


## GETOrderIdOrderSubscription

> GETOrderIdOrderSubscription(ctx, orderId).Execute()

Retrieve the order subscription associated to the order



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
    r, err := apiClient.OrderSubscriptionsApi.GETOrderIdOrderSubscription(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETOrderIdOrderSubscription``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdOrderSubscriptionRequest struct via the builder pattern


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


## GETOrderIdOrderSubscriptions

> GETOrderIdOrderSubscriptions(ctx, orderId).Execute()

Retrieve the order subscriptions associated to the order



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
    r, err := apiClient.OrderSubscriptionsApi.GETOrderIdOrderSubscriptions(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETOrderIdOrderSubscriptions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdOrderSubscriptionsRequest struct via the builder pattern


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


## GETOrderSubscriptionItemIdOrderSubscription

> GETOrderSubscriptionItemIdOrderSubscription(ctx, orderSubscriptionItemId).Execute()

Retrieve the order subscription associated to the order subscription item



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
    r, err := apiClient.OrderSubscriptionsApi.GETOrderSubscriptionItemIdOrderSubscription(context.Background(), orderSubscriptionItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETOrderSubscriptionItemIdOrderSubscription``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderSubscriptionItemIdOrderSubscriptionRequest struct via the builder pattern


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


## GETOrderSubscriptions

> GETOrderSubscriptions200Response GETOrderSubscriptions(ctx).Execute()

List all order subscriptions



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
    resp, r, err := apiClient.OrderSubscriptionsApi.GETOrderSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETOrderSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderSubscriptions`: GETOrderSubscriptions200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderSubscriptionsApi.GETOrderSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderSubscriptionsRequest struct via the builder pattern


### Return type

[**GETOrderSubscriptions200Response**](GETOrderSubscriptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderSubscriptionsOrderSubscriptionId

> GETOrderSubscriptionsOrderSubscriptionId200Response GETOrderSubscriptionsOrderSubscriptionId(ctx, orderSubscriptionId).Execute()

Retrieve an order subscription



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
    resp, r, err := apiClient.OrderSubscriptionsApi.GETOrderSubscriptionsOrderSubscriptionId(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETOrderSubscriptionsOrderSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderSubscriptionsOrderSubscriptionId`: GETOrderSubscriptionsOrderSubscriptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderSubscriptionsApi.GETOrderSubscriptionsOrderSubscriptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderSubscriptionsOrderSubscriptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETOrderSubscriptionsOrderSubscriptionId200Response**](GETOrderSubscriptionsOrderSubscriptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETRecurringOrderCopyIdOrderSubscription

> GETRecurringOrderCopyIdOrderSubscription(ctx, recurringOrderCopyId).Execute()

Retrieve the order subscription associated to the recurring order copy



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
    recurringOrderCopyId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderSubscriptionsApi.GETRecurringOrderCopyIdOrderSubscription(context.Background(), recurringOrderCopyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETRecurringOrderCopyIdOrderSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringOrderCopyId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETRecurringOrderCopyIdOrderSubscriptionRequest struct via the builder pattern


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


## GETSubscriptionModelIdOrderSubscriptions

> GETSubscriptionModelIdOrderSubscriptions(ctx, subscriptionModelId).Execute()

Retrieve the order subscriptions associated to the subscription model



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
    subscriptionModelId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderSubscriptionsApi.GETSubscriptionModelIdOrderSubscriptions(context.Background(), subscriptionModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.GETSubscriptionModelIdOrderSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionModelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSubscriptionModelIdOrderSubscriptionsRequest struct via the builder pattern


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


## PATCHOrderSubscriptionsOrderSubscriptionId

> PATCHOrderSubscriptionsOrderSubscriptionId200Response PATCHOrderSubscriptionsOrderSubscriptionId(ctx, orderSubscriptionId).OrderSubscriptionUpdate(orderSubscriptionUpdate).Execute()

Update an order subscription



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
    orderSubscriptionUpdate := *openapiclient.NewOrderSubscriptionUpdate(*openapiclient.NewOrderSubscriptionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes())) // OrderSubscriptionUpdate | 
    orderSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderSubscriptionsApi.PATCHOrderSubscriptionsOrderSubscriptionId(context.Background(), orderSubscriptionId).OrderSubscriptionUpdate(orderSubscriptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.PATCHOrderSubscriptionsOrderSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHOrderSubscriptionsOrderSubscriptionId`: PATCHOrderSubscriptionsOrderSubscriptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderSubscriptionsApi.PATCHOrderSubscriptionsOrderSubscriptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHOrderSubscriptionsOrderSubscriptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderSubscriptionUpdate** | [**OrderSubscriptionUpdate**](OrderSubscriptionUpdate.md) |  | 


### Return type

[**PATCHOrderSubscriptionsOrderSubscriptionId200Response**](PATCHOrderSubscriptionsOrderSubscriptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTOrderSubscriptions

> POSTOrderSubscriptions201Response POSTOrderSubscriptions(ctx).OrderSubscriptionCreate(orderSubscriptionCreate).Execute()

Create an order subscription



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
    orderSubscriptionCreate := *openapiclient.NewOrderSubscriptionCreate(*openapiclient.NewOrderSubscriptionCreateData(interface{}(123), *openapiclient.NewPOSTOrderSubscriptions201ResponseDataAttributes(interface{}(monthly)))) // OrderSubscriptionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderSubscriptionsApi.POSTOrderSubscriptions(context.Background()).OrderSubscriptionCreate(orderSubscriptionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderSubscriptionsApi.POSTOrderSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTOrderSubscriptions`: POSTOrderSubscriptions201Response
    fmt.Fprintf(os.Stdout, "Response from `OrderSubscriptionsApi.POSTOrderSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTOrderSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderSubscriptionCreate** | [**OrderSubscriptionCreate**](OrderSubscriptionCreate.md) |  | 

### Return type

[**POSTOrderSubscriptions201Response**](POSTOrderSubscriptions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

