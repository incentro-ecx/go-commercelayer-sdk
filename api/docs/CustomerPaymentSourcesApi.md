# \CustomerPaymentSourcesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomerPaymentSourcesCustomerPaymentSourceId**](CustomerPaymentSourcesApi.md#DELETECustomerPaymentSourcesCustomerPaymentSourceId) | **Delete** /customer_payment_sources/{customerPaymentSourceId} | Delete a customer payment source
[**GETCustomerIdCustomerPaymentSources**](CustomerPaymentSourcesApi.md#GETCustomerIdCustomerPaymentSources) | **Get** /customers/{customerId}/customer_payment_sources | Retrieve the customer payment sources associated to the customer
[**GETCustomerPaymentSources**](CustomerPaymentSourcesApi.md#GETCustomerPaymentSources) | **Get** /customer_payment_sources | List all customer payment sources
[**GETCustomerPaymentSourcesCustomerPaymentSourceId**](CustomerPaymentSourcesApi.md#GETCustomerPaymentSourcesCustomerPaymentSourceId) | **Get** /customer_payment_sources/{customerPaymentSourceId} | Retrieve a customer payment source
[**GETExternalPaymentIdWallet**](CustomerPaymentSourcesApi.md#GETExternalPaymentIdWallet) | **Get** /external_payments/{externalPaymentId}/wallet | Retrieve the wallet associated to the external payment
[**GETOrderIdAvailableCustomerPaymentSources**](CustomerPaymentSourcesApi.md#GETOrderIdAvailableCustomerPaymentSources) | **Get** /orders/{orderId}/available_customer_payment_sources | Retrieve the available customer payment sources associated to the order
[**GETOrderSubscriptionIdCustomerPaymentSource**](CustomerPaymentSourcesApi.md#GETOrderSubscriptionIdCustomerPaymentSource) | **Get** /order_subscriptions/{orderSubscriptionId}/customer_payment_source | Retrieve the customer payment source associated to the order subscription
[**PATCHCustomerPaymentSourcesCustomerPaymentSourceId**](CustomerPaymentSourcesApi.md#PATCHCustomerPaymentSourcesCustomerPaymentSourceId) | **Patch** /customer_payment_sources/{customerPaymentSourceId} | Update a customer payment source
[**POSTCustomerPaymentSources**](CustomerPaymentSourcesApi.md#POSTCustomerPaymentSources) | **Post** /customer_payment_sources | Create a customer payment source



## DELETECustomerPaymentSourcesCustomerPaymentSourceId

> DELETECustomerPaymentSourcesCustomerPaymentSourceId(ctx, customerPaymentSourceId).Execute()

Delete a customer payment source



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
    customerPaymentSourceId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomerPaymentSourcesApi.DELETECustomerPaymentSourcesCustomerPaymentSourceId(context.Background(), customerPaymentSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.DELETECustomerPaymentSourcesCustomerPaymentSourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPaymentSourceId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest struct via the builder pattern


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


## GETCustomerIdCustomerPaymentSources

> GETCustomerIdCustomerPaymentSources(ctx, customerId).Execute()

Retrieve the customer payment sources associated to the customer



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
    r, err := apiClient.CustomerPaymentSourcesApi.GETCustomerIdCustomerPaymentSources(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETCustomerIdCustomerPaymentSources``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCustomerIdCustomerPaymentSourcesRequest struct via the builder pattern


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


## GETCustomerPaymentSources

> GETCustomerPaymentSources200Response GETCustomerPaymentSources(ctx).Execute()

List all customer payment sources



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
    resp, r, err := apiClient.CustomerPaymentSourcesApi.GETCustomerPaymentSources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETCustomerPaymentSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerPaymentSources`: GETCustomerPaymentSources200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerPaymentSourcesApi.GETCustomerPaymentSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPaymentSourcesRequest struct via the builder pattern


### Return type

[**GETCustomerPaymentSources200Response**](GETCustomerPaymentSources200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomerPaymentSourcesCustomerPaymentSourceId

> GETCustomerPaymentSourcesCustomerPaymentSourceId200Response GETCustomerPaymentSourcesCustomerPaymentSourceId(ctx, customerPaymentSourceId).Execute()

Retrieve a customer payment source



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
    customerPaymentSourceId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.GETCustomerPaymentSourcesCustomerPaymentSourceId(context.Background(), customerPaymentSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETCustomerPaymentSourcesCustomerPaymentSourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerPaymentSourcesCustomerPaymentSourceId`: GETCustomerPaymentSourcesCustomerPaymentSourceId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerPaymentSourcesApi.GETCustomerPaymentSourcesCustomerPaymentSourceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPaymentSourceId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCustomerPaymentSourcesCustomerPaymentSourceId200Response**](GETCustomerPaymentSourcesCustomerPaymentSourceId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalPaymentIdWallet

> GETExternalPaymentIdWallet(ctx, externalPaymentId).Execute()

Retrieve the wallet associated to the external payment



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
    externalPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomerPaymentSourcesApi.GETExternalPaymentIdWallet(context.Background(), externalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETExternalPaymentIdWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPaymentIdWalletRequest struct via the builder pattern


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


## GETOrderIdAvailableCustomerPaymentSources

> GETOrderIdAvailableCustomerPaymentSources(ctx, orderId).Execute()

Retrieve the available customer payment sources associated to the order



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
    r, err := apiClient.CustomerPaymentSourcesApi.GETOrderIdAvailableCustomerPaymentSources(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETOrderIdAvailableCustomerPaymentSources``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdAvailableCustomerPaymentSourcesRequest struct via the builder pattern


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


## GETOrderSubscriptionIdCustomerPaymentSource

> GETOrderSubscriptionIdCustomerPaymentSource(ctx, orderSubscriptionId).Execute()

Retrieve the customer payment source associated to the order subscription



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
    r, err := apiClient.CustomerPaymentSourcesApi.GETOrderSubscriptionIdCustomerPaymentSource(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.GETOrderSubscriptionIdCustomerPaymentSource``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderSubscriptionIdCustomerPaymentSourceRequest struct via the builder pattern


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


## PATCHCustomerPaymentSourcesCustomerPaymentSourceId

> PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response PATCHCustomerPaymentSourcesCustomerPaymentSourceId(ctx, customerPaymentSourceId).CustomerPaymentSourceUpdate(customerPaymentSourceUpdate).Execute()

Update a customer payment source



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
    customerPaymentSourceUpdate := *openapiclient.NewCustomerPaymentSourceUpdate(*openapiclient.NewCustomerPaymentSourceUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes())) // CustomerPaymentSourceUpdate | 
    customerPaymentSourceId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.PATCHCustomerPaymentSourcesCustomerPaymentSourceId(context.Background(), customerPaymentSourceId).CustomerPaymentSourceUpdate(customerPaymentSourceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.PATCHCustomerPaymentSourcesCustomerPaymentSourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCustomerPaymentSourcesCustomerPaymentSourceId`: PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerPaymentSourcesApi.PATCHCustomerPaymentSourcesCustomerPaymentSourceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPaymentSourceId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerPaymentSourceUpdate** | [**CustomerPaymentSourceUpdate**](CustomerPaymentSourceUpdate.md) |  | 


### Return type

[**PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response**](PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCustomerPaymentSources

> POSTCustomerPaymentSources201Response POSTCustomerPaymentSources(ctx).CustomerPaymentSourceCreate(customerPaymentSourceCreate).Execute()

Create a customer payment source



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
    customerPaymentSourceCreate := *openapiclient.NewCustomerPaymentSourceCreate(*openapiclient.NewCustomerPaymentSourceCreateData(interface{}(123), *openapiclient.NewPOSTCustomerPaymentSources201ResponseDataAttributes())) // CustomerPaymentSourceCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPaymentSourcesApi.POSTCustomerPaymentSources(context.Background()).CustomerPaymentSourceCreate(customerPaymentSourceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPaymentSourcesApi.POSTCustomerPaymentSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCustomerPaymentSources`: POSTCustomerPaymentSources201Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerPaymentSourcesApi.POSTCustomerPaymentSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomerPaymentSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerPaymentSourceCreate** | [**CustomerPaymentSourceCreate**](CustomerPaymentSourceCreate.md) |  | 

### Return type

[**POSTCustomerPaymentSources201Response**](POSTCustomerPaymentSources201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

