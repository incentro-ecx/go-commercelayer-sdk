# \SubscriptionModelsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESubscriptionModelsSubscriptionModelId**](SubscriptionModelsApi.md#DELETESubscriptionModelsSubscriptionModelId) | **Delete** /subscription_models/{subscriptionModelId} | Delete a subscription model
[**GETMarketIdSubscriptionModel**](SubscriptionModelsApi.md#GETMarketIdSubscriptionModel) | **Get** /markets/{marketId}/subscription_model | Retrieve the subscription model associated to the market
[**GETOrderSubscriptionIdSubscriptionModel**](SubscriptionModelsApi.md#GETOrderSubscriptionIdSubscriptionModel) | **Get** /order_subscriptions/{orderSubscriptionId}/subscription_model | Retrieve the subscription model associated to the order subscription
[**GETSubscriptionModels**](SubscriptionModelsApi.md#GETSubscriptionModels) | **Get** /subscription_models | List all subscription models
[**GETSubscriptionModelsSubscriptionModelId**](SubscriptionModelsApi.md#GETSubscriptionModelsSubscriptionModelId) | **Get** /subscription_models/{subscriptionModelId} | Retrieve a subscription model
[**PATCHSubscriptionModelsSubscriptionModelId**](SubscriptionModelsApi.md#PATCHSubscriptionModelsSubscriptionModelId) | **Patch** /subscription_models/{subscriptionModelId} | Update a subscription model
[**POSTSubscriptionModels**](SubscriptionModelsApi.md#POSTSubscriptionModels) | **Post** /subscription_models | Create a subscription model



## DELETESubscriptionModelsSubscriptionModelId

> DELETESubscriptionModelsSubscriptionModelId(ctx, subscriptionModelId).Execute()

Delete a subscription model



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
    r, err := apiClient.SubscriptionModelsApi.DELETESubscriptionModelsSubscriptionModelId(context.Background(), subscriptionModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModelsApi.DELETESubscriptionModelsSubscriptionModelId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETESubscriptionModelsSubscriptionModelIdRequest struct via the builder pattern


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


## GETMarketIdSubscriptionModel

> GETMarketIdSubscriptionModel(ctx, marketId).Execute()

Retrieve the subscription model associated to the market



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
    marketId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionModelsApi.GETMarketIdSubscriptionModel(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModelsApi.GETMarketIdSubscriptionModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdSubscriptionModelRequest struct via the builder pattern


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


## GETOrderSubscriptionIdSubscriptionModel

> GETOrderSubscriptionIdSubscriptionModel(ctx, orderSubscriptionId).Execute()

Retrieve the subscription model associated to the order subscription



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
    r, err := apiClient.SubscriptionModelsApi.GETOrderSubscriptionIdSubscriptionModel(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModelsApi.GETOrderSubscriptionIdSubscriptionModel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderSubscriptionIdSubscriptionModelRequest struct via the builder pattern


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


## GETSubscriptionModels

> GETSubscriptionModels200Response GETSubscriptionModels(ctx).Execute()

List all subscription models



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
    resp, r, err := apiClient.SubscriptionModelsApi.GETSubscriptionModels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModelsApi.GETSubscriptionModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSubscriptionModels`: GETSubscriptionModels200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionModelsApi.GETSubscriptionModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETSubscriptionModelsRequest struct via the builder pattern


### Return type

[**GETSubscriptionModels200Response**](GETSubscriptionModels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSubscriptionModelsSubscriptionModelId

> GETSubscriptionModelsSubscriptionModelId200Response GETSubscriptionModelsSubscriptionModelId(ctx, subscriptionModelId).Execute()

Retrieve a subscription model



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
    resp, r, err := apiClient.SubscriptionModelsApi.GETSubscriptionModelsSubscriptionModelId(context.Background(), subscriptionModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModelsApi.GETSubscriptionModelsSubscriptionModelId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSubscriptionModelsSubscriptionModelId`: GETSubscriptionModelsSubscriptionModelId200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionModelsApi.GETSubscriptionModelsSubscriptionModelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionModelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSubscriptionModelsSubscriptionModelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETSubscriptionModelsSubscriptionModelId200Response**](GETSubscriptionModelsSubscriptionModelId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHSubscriptionModelsSubscriptionModelId

> PATCHSubscriptionModelsSubscriptionModelId200Response PATCHSubscriptionModelsSubscriptionModelId(ctx, subscriptionModelId).SubscriptionModelUpdate(subscriptionModelUpdate).Execute()

Update a subscription model



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
    subscriptionModelUpdate := *openapiclient.NewSubscriptionModelUpdate(*openapiclient.NewSubscriptionModelUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes())) // SubscriptionModelUpdate | 
    subscriptionModelId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionModelsApi.PATCHSubscriptionModelsSubscriptionModelId(context.Background(), subscriptionModelId).SubscriptionModelUpdate(subscriptionModelUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModelsApi.PATCHSubscriptionModelsSubscriptionModelId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHSubscriptionModelsSubscriptionModelId`: PATCHSubscriptionModelsSubscriptionModelId200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionModelsApi.PATCHSubscriptionModelsSubscriptionModelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionModelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHSubscriptionModelsSubscriptionModelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionModelUpdate** | [**SubscriptionModelUpdate**](SubscriptionModelUpdate.md) |  | 


### Return type

[**PATCHSubscriptionModelsSubscriptionModelId200Response**](PATCHSubscriptionModelsSubscriptionModelId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTSubscriptionModels

> POSTSubscriptionModels201Response POSTSubscriptionModels(ctx).SubscriptionModelCreate(subscriptionModelCreate).Execute()

Create a subscription model



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
    subscriptionModelCreate := *openapiclient.NewSubscriptionModelCreate(*openapiclient.NewSubscriptionModelCreateData(interface{}(123), *openapiclient.NewPOSTSubscriptionModels201ResponseDataAttributes(interface{}(EU Subscription Model), interface{}([hourly, 10 * * * *, weekly, monthly, two-month])))) // SubscriptionModelCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionModelsApi.POSTSubscriptionModels(context.Background()).SubscriptionModelCreate(subscriptionModelCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModelsApi.POSTSubscriptionModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTSubscriptionModels`: POSTSubscriptionModels201Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionModelsApi.POSTSubscriptionModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTSubscriptionModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionModelCreate** | [**SubscriptionModelCreate**](SubscriptionModelCreate.md) |  | 

### Return type

[**POSTSubscriptionModels201Response**](POSTSubscriptionModels201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

