# \CustomerSubscriptionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomerSubscriptionsCustomerSubscriptionId**](CustomerSubscriptionsApi.md#DELETECustomerSubscriptionsCustomerSubscriptionId) | **Delete** /customer_subscriptions/{customerSubscriptionId} | Delete a customer subscription
[**GETCustomerIdCustomerSubscriptions**](CustomerSubscriptionsApi.md#GETCustomerIdCustomerSubscriptions) | **Get** /customers/{customerId}/customer_subscriptions | Retrieve the customer subscriptions associated to the customer
[**GETCustomerSubscriptions**](CustomerSubscriptionsApi.md#GETCustomerSubscriptions) | **Get** /customer_subscriptions | List all customer subscriptions
[**GETCustomerSubscriptionsCustomerSubscriptionId**](CustomerSubscriptionsApi.md#GETCustomerSubscriptionsCustomerSubscriptionId) | **Get** /customer_subscriptions/{customerSubscriptionId} | Retrieve a customer subscription
[**PATCHCustomerSubscriptionsCustomerSubscriptionId**](CustomerSubscriptionsApi.md#PATCHCustomerSubscriptionsCustomerSubscriptionId) | **Patch** /customer_subscriptions/{customerSubscriptionId} | Update a customer subscription
[**POSTCustomerSubscriptions**](CustomerSubscriptionsApi.md#POSTCustomerSubscriptions) | **Post** /customer_subscriptions | Create a customer subscription



## DELETECustomerSubscriptionsCustomerSubscriptionId

> DELETECustomerSubscriptionsCustomerSubscriptionId(ctx, customerSubscriptionId).Execute()

Delete a customer subscription



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
    customerSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomerSubscriptionsApi.DELETECustomerSubscriptionsCustomerSubscriptionId(context.Background(), customerSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSubscriptionsApi.DELETECustomerSubscriptionsCustomerSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest struct via the builder pattern


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


## GETCustomerIdCustomerSubscriptions

> GETCustomerIdCustomerSubscriptions(ctx, customerId).Execute()

Retrieve the customer subscriptions associated to the customer



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
    r, err := apiClient.CustomerSubscriptionsApi.GETCustomerIdCustomerSubscriptions(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSubscriptionsApi.GETCustomerIdCustomerSubscriptions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCustomerIdCustomerSubscriptionsRequest struct via the builder pattern


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


## GETCustomerSubscriptions

> GETCustomerSubscriptions200Response GETCustomerSubscriptions(ctx).Execute()

List all customer subscriptions



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
    resp, r, err := apiClient.CustomerSubscriptionsApi.GETCustomerSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSubscriptionsApi.GETCustomerSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerSubscriptions`: GETCustomerSubscriptions200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerSubscriptionsApi.GETCustomerSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerSubscriptionsRequest struct via the builder pattern


### Return type

[**GETCustomerSubscriptions200Response**](GETCustomerSubscriptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomerSubscriptionsCustomerSubscriptionId

> GETCustomerSubscriptionsCustomerSubscriptionId200Response GETCustomerSubscriptionsCustomerSubscriptionId(ctx, customerSubscriptionId).Execute()

Retrieve a customer subscription



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
    customerSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerSubscriptionsApi.GETCustomerSubscriptionsCustomerSubscriptionId(context.Background(), customerSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSubscriptionsApi.GETCustomerSubscriptionsCustomerSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerSubscriptionsCustomerSubscriptionId`: GETCustomerSubscriptionsCustomerSubscriptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerSubscriptionsApi.GETCustomerSubscriptionsCustomerSubscriptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerSubscriptionsCustomerSubscriptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCustomerSubscriptionsCustomerSubscriptionId200Response**](GETCustomerSubscriptionsCustomerSubscriptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCustomerSubscriptionsCustomerSubscriptionId

> PATCHCustomerSubscriptionsCustomerSubscriptionId200Response PATCHCustomerSubscriptionsCustomerSubscriptionId(ctx, customerSubscriptionId).CustomerSubscriptionUpdate(customerSubscriptionUpdate).Execute()

Update a customer subscription



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
    customerSubscriptionUpdate := *openapiclient.NewCustomerSubscriptionUpdate(*openapiclient.NewCustomerSubscriptionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseDataAttributes())) // CustomerSubscriptionUpdate | 
    customerSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerSubscriptionsApi.PATCHCustomerSubscriptionsCustomerSubscriptionId(context.Background(), customerSubscriptionId).CustomerSubscriptionUpdate(customerSubscriptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSubscriptionsApi.PATCHCustomerSubscriptionsCustomerSubscriptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCustomerSubscriptionsCustomerSubscriptionId`: PATCHCustomerSubscriptionsCustomerSubscriptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerSubscriptionsApi.PATCHCustomerSubscriptionsCustomerSubscriptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerSubscriptionUpdate** | [**CustomerSubscriptionUpdate**](CustomerSubscriptionUpdate.md) |  | 


### Return type

[**PATCHCustomerSubscriptionsCustomerSubscriptionId200Response**](PATCHCustomerSubscriptionsCustomerSubscriptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCustomerSubscriptions

> POSTCustomerSubscriptions201Response POSTCustomerSubscriptions(ctx).CustomerSubscriptionCreate(customerSubscriptionCreate).Execute()

Create a customer subscription



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
    customerSubscriptionCreate := *openapiclient.NewCustomerSubscriptionCreate(*openapiclient.NewCustomerSubscriptionCreateData(interface{}(123), *openapiclient.NewPOSTCustomerSubscriptions201ResponseDataAttributes(interface{}(john@example.com), interface{}(ANY-EXTERNAL-REFEFERNCE)))) // CustomerSubscriptionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerSubscriptionsApi.POSTCustomerSubscriptions(context.Background()).CustomerSubscriptionCreate(customerSubscriptionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerSubscriptionsApi.POSTCustomerSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCustomerSubscriptions`: POSTCustomerSubscriptions201Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerSubscriptionsApi.POSTCustomerSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomerSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerSubscriptionCreate** | [**CustomerSubscriptionCreate**](CustomerSubscriptionCreate.md) |  | 

### Return type

[**POSTCustomerSubscriptions201Response**](POSTCustomerSubscriptions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

