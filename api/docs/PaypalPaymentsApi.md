# \PaypalPaymentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPaypalPaymentsPaypalPaymentId**](PaypalPaymentsApi.md#DELETEPaypalPaymentsPaypalPaymentId) | **Delete** /paypal_payments/{paypalPaymentId} | Delete a paypal payment
[**GETPaypalGatewayIdPaypalPayments**](PaypalPaymentsApi.md#GETPaypalGatewayIdPaypalPayments) | **Get** /paypal_gateways/{paypalGatewayId}/paypal_payments | Retrieve the paypal payments associated to the paypal gateway
[**GETPaypalPayments**](PaypalPaymentsApi.md#GETPaypalPayments) | **Get** /paypal_payments | List all paypal payments
[**GETPaypalPaymentsPaypalPaymentId**](PaypalPaymentsApi.md#GETPaypalPaymentsPaypalPaymentId) | **Get** /paypal_payments/{paypalPaymentId} | Retrieve a paypal payment
[**PATCHPaypalPaymentsPaypalPaymentId**](PaypalPaymentsApi.md#PATCHPaypalPaymentsPaypalPaymentId) | **Patch** /paypal_payments/{paypalPaymentId} | Update a paypal payment
[**POSTPaypalPayments**](PaypalPaymentsApi.md#POSTPaypalPayments) | **Post** /paypal_payments | Create a paypal payment



## DELETEPaypalPaymentsPaypalPaymentId

> DELETEPaypalPaymentsPaypalPaymentId(ctx, paypalPaymentId).Execute()

Delete a paypal payment



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
    paypalPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaypalPaymentsApi.DELETEPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.DELETEPaypalPaymentsPaypalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPaypalPaymentsPaypalPaymentIdRequest struct via the builder pattern


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


## GETPaypalGatewayIdPaypalPayments

> GETPaypalGatewayIdPaypalPayments(ctx, paypalGatewayId).Execute()

Retrieve the paypal payments associated to the paypal gateway



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
    paypalGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaypalPaymentsApi.GETPaypalGatewayIdPaypalPayments(context.Background(), paypalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.GETPaypalGatewayIdPaypalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaypalGatewayIdPaypalPaymentsRequest struct via the builder pattern


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


## GETPaypalPayments

> GETPaypalPayments200Response GETPaypalPayments(ctx).Execute()

List all paypal payments



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
    resp, r, err := apiClient.PaypalPaymentsApi.GETPaypalPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.GETPaypalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaypalPayments`: GETPaypalPayments200Response
    fmt.Fprintf(os.Stdout, "Response from `PaypalPaymentsApi.GETPaypalPayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaypalPaymentsRequest struct via the builder pattern


### Return type

[**GETPaypalPayments200Response**](GETPaypalPayments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPaypalPaymentsPaypalPaymentId

> GETPaypalPaymentsPaypalPaymentId200Response GETPaypalPaymentsPaypalPaymentId(ctx, paypalPaymentId).Execute()

Retrieve a paypal payment



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
    paypalPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalPaymentsApi.GETPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.GETPaypalPaymentsPaypalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaypalPaymentsPaypalPaymentId`: GETPaypalPaymentsPaypalPaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaypalPaymentsApi.GETPaypalPaymentsPaypalPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaypalPaymentsPaypalPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPaypalPaymentsPaypalPaymentId200Response**](GETPaypalPaymentsPaypalPaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPaypalPaymentsPaypalPaymentId

> PATCHPaypalPaymentsPaypalPaymentId200Response PATCHPaypalPaymentsPaypalPaymentId(ctx, paypalPaymentId).PaypalPaymentUpdate(paypalPaymentUpdate).Execute()

Update a paypal payment



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
    paypalPaymentUpdate := *openapiclient.NewPaypalPaymentUpdate(*openapiclient.NewPaypalPaymentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPaypalPaymentsPaypalPaymentId200ResponseDataAttributes())) // PaypalPaymentUpdate | 
    paypalPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalPaymentsApi.PATCHPaypalPaymentsPaypalPaymentId(context.Background(), paypalPaymentId).PaypalPaymentUpdate(paypalPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.PATCHPaypalPaymentsPaypalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPaypalPaymentsPaypalPaymentId`: PATCHPaypalPaymentsPaypalPaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaypalPaymentsApi.PATCHPaypalPaymentsPaypalPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPaypalPaymentsPaypalPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paypalPaymentUpdate** | [**PaypalPaymentUpdate**](PaypalPaymentUpdate.md) |  | 


### Return type

[**PATCHPaypalPaymentsPaypalPaymentId200Response**](PATCHPaypalPaymentsPaypalPaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPaypalPayments

> POSTPaypalPayments201Response POSTPaypalPayments(ctx).PaypalPaymentCreate(paypalPaymentCreate).Execute()

Create a paypal payment



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
    paypalPaymentCreate := *openapiclient.NewPaypalPaymentCreate(*openapiclient.NewPaypalPaymentCreateData(interface{}(123), *openapiclient.NewPOSTPaypalPayments201ResponseDataAttributes(interface{}(https://yourdomain.com/thankyou), interface{}(https://yourdomain.com/checkout/payment)))) // PaypalPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalPaymentsApi.POSTPaypalPayments(context.Background()).PaypalPaymentCreate(paypalPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalPaymentsApi.POSTPaypalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPaypalPayments`: POSTPaypalPayments201Response
    fmt.Fprintf(os.Stdout, "Response from `PaypalPaymentsApi.POSTPaypalPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPaypalPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paypalPaymentCreate** | [**PaypalPaymentCreate**](PaypalPaymentCreate.md) |  | 

### Return type

[**POSTPaypalPayments201Response**](POSTPaypalPayments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

