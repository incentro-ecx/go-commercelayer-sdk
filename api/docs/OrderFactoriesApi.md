# \OrderFactoriesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETOrderFactories**](OrderFactoriesApi.md#GETOrderFactories) | **Get** /order_factories | List all order factories
[**GETOrderFactoriesOrderFactoryId**](OrderFactoriesApi.md#GETOrderFactoriesOrderFactoryId) | **Get** /order_factories/{orderFactoryId} | Retrieve an order factory
[**GETOrderIdOrderFactories**](OrderFactoriesApi.md#GETOrderIdOrderFactories) | **Get** /orders/{orderId}/order_factories | Retrieve the order factories associated to the order
[**GETOrderSubscriptionIdOrderFactories**](OrderFactoriesApi.md#GETOrderSubscriptionIdOrderFactories) | **Get** /order_subscriptions/{orderSubscriptionId}/order_factories | Retrieve the order factories associated to the order subscription



## GETOrderFactories

> GETOrderFactories200Response GETOrderFactories(ctx).Execute()

List all order factories



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
    resp, r, err := apiClient.OrderFactoriesApi.GETOrderFactories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderFactoriesApi.GETOrderFactories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderFactories`: GETOrderFactories200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderFactoriesApi.GETOrderFactories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderFactoriesRequest struct via the builder pattern


### Return type

[**GETOrderFactories200Response**](GETOrderFactories200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderFactoriesOrderFactoryId

> GETOrderFactoriesOrderFactoryId200Response GETOrderFactoriesOrderFactoryId(ctx, orderFactoryId).Execute()

Retrieve an order factory



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
    orderFactoryId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderFactoriesApi.GETOrderFactoriesOrderFactoryId(context.Background(), orderFactoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderFactoriesApi.GETOrderFactoriesOrderFactoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderFactoriesOrderFactoryId`: GETOrderFactoriesOrderFactoryId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderFactoriesApi.GETOrderFactoriesOrderFactoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderFactoryId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderFactoriesOrderFactoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETOrderFactoriesOrderFactoryId200Response**](GETOrderFactoriesOrderFactoryId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdOrderFactories

> GETOrderIdOrderFactories(ctx, orderId).Execute()

Retrieve the order factories associated to the order



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
    r, err := apiClient.OrderFactoriesApi.GETOrderIdOrderFactories(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderFactoriesApi.GETOrderIdOrderFactories``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdOrderFactoriesRequest struct via the builder pattern


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


## GETOrderSubscriptionIdOrderFactories

> GETOrderSubscriptionIdOrderFactories(ctx, orderSubscriptionId).Execute()

Retrieve the order factories associated to the order subscription



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
    r, err := apiClient.OrderFactoriesApi.GETOrderSubscriptionIdOrderFactories(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderFactoriesApi.GETOrderSubscriptionIdOrderFactories``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderSubscriptionIdOrderFactoriesRequest struct via the builder pattern


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

