# \StripePaymentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStripePaymentsStripePaymentId**](StripePaymentsApi.md#DELETEStripePaymentsStripePaymentId) | **Delete** /stripe_payments/{stripePaymentId} | Delete a stripe payment
[**GETStripeGatewayIdStripePayments**](StripePaymentsApi.md#GETStripeGatewayIdStripePayments) | **Get** /stripe_gateways/{stripeGatewayId}/stripe_payments | Retrieve the stripe payments associated to the stripe gateway
[**GETStripePayments**](StripePaymentsApi.md#GETStripePayments) | **Get** /stripe_payments | List all stripe payments
[**GETStripePaymentsStripePaymentId**](StripePaymentsApi.md#GETStripePaymentsStripePaymentId) | **Get** /stripe_payments/{stripePaymentId} | Retrieve a stripe payment
[**PATCHStripePaymentsStripePaymentId**](StripePaymentsApi.md#PATCHStripePaymentsStripePaymentId) | **Patch** /stripe_payments/{stripePaymentId} | Update a stripe payment
[**POSTStripePayments**](StripePaymentsApi.md#POSTStripePayments) | **Post** /stripe_payments | Create a stripe payment



## DELETEStripePaymentsStripePaymentId

> DELETEStripePaymentsStripePaymentId(ctx, stripePaymentId).Execute()

Delete a stripe payment



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
    stripePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StripePaymentsApi.DELETEStripePaymentsStripePaymentId(context.Background(), stripePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.DELETEStripePaymentsStripePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEStripePaymentsStripePaymentIdRequest struct via the builder pattern


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


## GETStripeGatewayIdStripePayments

> GETStripeGatewayIdStripePayments(ctx, stripeGatewayId).Execute()

Retrieve the stripe payments associated to the stripe gateway



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
    stripeGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StripePaymentsApi.GETStripeGatewayIdStripePayments(context.Background(), stripeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.GETStripeGatewayIdStripePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripeGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripeGatewayIdStripePaymentsRequest struct via the builder pattern


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


## GETStripePayments

> GETStripePayments200Response GETStripePayments(ctx).Execute()

List all stripe payments



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
    resp, r, err := apiClient.StripePaymentsApi.GETStripePayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.GETStripePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStripePayments`: GETStripePayments200Response
    fmt.Fprintf(os.Stdout, "Response from `StripePaymentsApi.GETStripePayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripePaymentsRequest struct via the builder pattern


### Return type

[**GETStripePayments200Response**](GETStripePayments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStripePaymentsStripePaymentId

> GETStripePaymentsStripePaymentId200Response GETStripePaymentsStripePaymentId(ctx, stripePaymentId).Execute()

Retrieve a stripe payment



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
    stripePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripePaymentsApi.GETStripePaymentsStripePaymentId(context.Background(), stripePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.GETStripePaymentsStripePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStripePaymentsStripePaymentId`: GETStripePaymentsStripePaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `StripePaymentsApi.GETStripePaymentsStripePaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripePaymentsStripePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETStripePaymentsStripePaymentId200Response**](GETStripePaymentsStripePaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHStripePaymentsStripePaymentId

> PATCHStripePaymentsStripePaymentId200Response PATCHStripePaymentsStripePaymentId(ctx, stripePaymentId).StripePaymentUpdate(stripePaymentUpdate).Execute()

Update a stripe payment



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
    stripePaymentUpdate := *openapiclient.NewStripePaymentUpdate(*openapiclient.NewStripePaymentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributes())) // StripePaymentUpdate | 
    stripePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripePaymentsApi.PATCHStripePaymentsStripePaymentId(context.Background(), stripePaymentId).StripePaymentUpdate(stripePaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.PATCHStripePaymentsStripePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHStripePaymentsStripePaymentId`: PATCHStripePaymentsStripePaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `StripePaymentsApi.PATCHStripePaymentsStripePaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStripePaymentsStripePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripePaymentUpdate** | [**StripePaymentUpdate**](StripePaymentUpdate.md) |  | 


### Return type

[**PATCHStripePaymentsStripePaymentId200Response**](PATCHStripePaymentsStripePaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTStripePayments

> POSTStripePayments201Response POSTStripePayments(ctx).StripePaymentCreate(stripePaymentCreate).Execute()

Create a stripe payment



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
    stripePaymentCreate := *openapiclient.NewStripePaymentCreate(*openapiclient.NewStripePaymentCreateData(interface{}(123), *openapiclient.NewPOSTStripePayments201ResponseDataAttributes())) // StripePaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripePaymentsApi.POSTStripePayments(context.Background()).StripePaymentCreate(stripePaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripePaymentsApi.POSTStripePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTStripePayments`: POSTStripePayments201Response
    fmt.Fprintf(os.Stdout, "Response from `StripePaymentsApi.POSTStripePayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStripePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripePaymentCreate** | [**StripePaymentCreate**](StripePaymentCreate.md) |  | 

### Return type

[**POSTStripePayments201Response**](POSTStripePayments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

