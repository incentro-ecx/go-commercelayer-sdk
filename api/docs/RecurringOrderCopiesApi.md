# \RecurringOrderCopiesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETERecurringOrderCopiesRecurringOrderCopyId**](RecurringOrderCopiesApi.md#DELETERecurringOrderCopiesRecurringOrderCopyId) | **Delete** /recurring_order_copies/{recurringOrderCopyId} | Delete a recurring order copy
[**GETOrderIdRecurringOrderCopies**](RecurringOrderCopiesApi.md#GETOrderIdRecurringOrderCopies) | **Get** /orders/{orderId}/recurring_order_copies | Retrieve the recurring order copies associated to the order
[**GETOrderSubscriptionIdRecurringOrderCopies**](RecurringOrderCopiesApi.md#GETOrderSubscriptionIdRecurringOrderCopies) | **Get** /order_subscriptions/{orderSubscriptionId}/recurring_order_copies | Retrieve the recurring order copies associated to the order subscription
[**GETRecurringOrderCopies**](RecurringOrderCopiesApi.md#GETRecurringOrderCopies) | **Get** /recurring_order_copies | List all recurring order copies
[**GETRecurringOrderCopiesRecurringOrderCopyId**](RecurringOrderCopiesApi.md#GETRecurringOrderCopiesRecurringOrderCopyId) | **Get** /recurring_order_copies/{recurringOrderCopyId} | Retrieve a recurring order copy
[**PATCHRecurringOrderCopiesRecurringOrderCopyId**](RecurringOrderCopiesApi.md#PATCHRecurringOrderCopiesRecurringOrderCopyId) | **Patch** /recurring_order_copies/{recurringOrderCopyId} | Update a recurring order copy
[**POSTRecurringOrderCopies**](RecurringOrderCopiesApi.md#POSTRecurringOrderCopies) | **Post** /recurring_order_copies | Create a recurring order copy



## DELETERecurringOrderCopiesRecurringOrderCopyId

> DELETERecurringOrderCopiesRecurringOrderCopyId(ctx, recurringOrderCopyId).Execute()

Delete a recurring order copy



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
    r, err := apiClient.RecurringOrderCopiesApi.DELETERecurringOrderCopiesRecurringOrderCopyId(context.Background(), recurringOrderCopyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringOrderCopiesApi.DELETERecurringOrderCopiesRecurringOrderCopyId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETERecurringOrderCopiesRecurringOrderCopyIdRequest struct via the builder pattern


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


## GETOrderIdRecurringOrderCopies

> GETOrderIdRecurringOrderCopies(ctx, orderId).Execute()

Retrieve the recurring order copies associated to the order



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
    r, err := apiClient.RecurringOrderCopiesApi.GETOrderIdRecurringOrderCopies(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringOrderCopiesApi.GETOrderIdRecurringOrderCopies``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdRecurringOrderCopiesRequest struct via the builder pattern


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


## GETOrderSubscriptionIdRecurringOrderCopies

> GETOrderSubscriptionIdRecurringOrderCopies(ctx, orderSubscriptionId).Execute()

Retrieve the recurring order copies associated to the order subscription



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
    r, err := apiClient.RecurringOrderCopiesApi.GETOrderSubscriptionIdRecurringOrderCopies(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringOrderCopiesApi.GETOrderSubscriptionIdRecurringOrderCopies``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderSubscriptionIdRecurringOrderCopiesRequest struct via the builder pattern


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


## GETRecurringOrderCopies

> GETRecurringOrderCopies200Response GETRecurringOrderCopies(ctx).Execute()

List all recurring order copies



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
    resp, r, err := apiClient.RecurringOrderCopiesApi.GETRecurringOrderCopies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringOrderCopiesApi.GETRecurringOrderCopies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETRecurringOrderCopies`: GETRecurringOrderCopies200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringOrderCopiesApi.GETRecurringOrderCopies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETRecurringOrderCopiesRequest struct via the builder pattern


### Return type

[**GETRecurringOrderCopies200Response**](GETRecurringOrderCopies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETRecurringOrderCopiesRecurringOrderCopyId

> GETRecurringOrderCopiesRecurringOrderCopyId200Response GETRecurringOrderCopiesRecurringOrderCopyId(ctx, recurringOrderCopyId).Execute()

Retrieve a recurring order copy



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
    resp, r, err := apiClient.RecurringOrderCopiesApi.GETRecurringOrderCopiesRecurringOrderCopyId(context.Background(), recurringOrderCopyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringOrderCopiesApi.GETRecurringOrderCopiesRecurringOrderCopyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETRecurringOrderCopiesRecurringOrderCopyId`: GETRecurringOrderCopiesRecurringOrderCopyId200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringOrderCopiesApi.GETRecurringOrderCopiesRecurringOrderCopyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringOrderCopyId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETRecurringOrderCopiesRecurringOrderCopyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETRecurringOrderCopiesRecurringOrderCopyId200Response**](GETRecurringOrderCopiesRecurringOrderCopyId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHRecurringOrderCopiesRecurringOrderCopyId

> PATCHRecurringOrderCopiesRecurringOrderCopyId200Response PATCHRecurringOrderCopiesRecurringOrderCopyId(ctx, recurringOrderCopyId).RecurringOrderCopyUpdate(recurringOrderCopyUpdate).Execute()

Update a recurring order copy



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
    recurringOrderCopyUpdate := *openapiclient.NewRecurringOrderCopyUpdate(*openapiclient.NewRecurringOrderCopyUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes())) // RecurringOrderCopyUpdate | 
    recurringOrderCopyId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringOrderCopiesApi.PATCHRecurringOrderCopiesRecurringOrderCopyId(context.Background(), recurringOrderCopyId).RecurringOrderCopyUpdate(recurringOrderCopyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringOrderCopiesApi.PATCHRecurringOrderCopiesRecurringOrderCopyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHRecurringOrderCopiesRecurringOrderCopyId`: PATCHRecurringOrderCopiesRecurringOrderCopyId200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringOrderCopiesApi.PATCHRecurringOrderCopiesRecurringOrderCopyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringOrderCopyId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHRecurringOrderCopiesRecurringOrderCopyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recurringOrderCopyUpdate** | [**RecurringOrderCopyUpdate**](RecurringOrderCopyUpdate.md) |  | 


### Return type

[**PATCHRecurringOrderCopiesRecurringOrderCopyId200Response**](PATCHRecurringOrderCopiesRecurringOrderCopyId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTRecurringOrderCopies

> POSTRecurringOrderCopies201Response POSTRecurringOrderCopies(ctx).RecurringOrderCopyCreate(recurringOrderCopyCreate).Execute()

Create a recurring order copy



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
    recurringOrderCopyCreate := *openapiclient.NewRecurringOrderCopyCreate(*openapiclient.NewRecurringOrderCopyCreateData(interface{}(123), *openapiclient.NewPOSTRecurringOrderCopies201ResponseDataAttributes())) // RecurringOrderCopyCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringOrderCopiesApi.POSTRecurringOrderCopies(context.Background()).RecurringOrderCopyCreate(recurringOrderCopyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringOrderCopiesApi.POSTRecurringOrderCopies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTRecurringOrderCopies`: POSTRecurringOrderCopies201Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringOrderCopiesApi.POSTRecurringOrderCopies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTRecurringOrderCopiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recurringOrderCopyCreate** | [**RecurringOrderCopyCreate**](RecurringOrderCopyCreate.md) |  | 

### Return type

[**POSTRecurringOrderCopies201Response**](POSTRecurringOrderCopies201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

