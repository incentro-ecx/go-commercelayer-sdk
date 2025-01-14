# \PaymentOptionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPaymentOptionsPaymentOptionId**](PaymentOptionsApi.md#DELETEPaymentOptionsPaymentOptionId) | **Delete** /payment_options/{paymentOptionId} | Delete a payment option
[**GETOrderIdPaymentOptions**](PaymentOptionsApi.md#GETOrderIdPaymentOptions) | **Get** /orders/{orderId}/payment_options | Retrieve the payment options associated to the order
[**GETPaymentOptions**](PaymentOptionsApi.md#GETPaymentOptions) | **Get** /payment_options | List all payment options
[**GETPaymentOptionsPaymentOptionId**](PaymentOptionsApi.md#GETPaymentOptionsPaymentOptionId) | **Get** /payment_options/{paymentOptionId} | Retrieve a payment option
[**PATCHPaymentOptionsPaymentOptionId**](PaymentOptionsApi.md#PATCHPaymentOptionsPaymentOptionId) | **Patch** /payment_options/{paymentOptionId} | Update a payment option
[**POSTPaymentOptions**](PaymentOptionsApi.md#POSTPaymentOptions) | **Post** /payment_options | Create a payment option



## DELETEPaymentOptionsPaymentOptionId

> DELETEPaymentOptionsPaymentOptionId(ctx, paymentOptionId).Execute()

Delete a payment option



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
    paymentOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentOptionsApi.DELETEPaymentOptionsPaymentOptionId(context.Background(), paymentOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsApi.DELETEPaymentOptionsPaymentOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPaymentOptionsPaymentOptionIdRequest struct via the builder pattern


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


## GETOrderIdPaymentOptions

> GETOrderIdPaymentOptions(ctx, orderId).Execute()

Retrieve the payment options associated to the order



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
    r, err := apiClient.PaymentOptionsApi.GETOrderIdPaymentOptions(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsApi.GETOrderIdPaymentOptions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdPaymentOptionsRequest struct via the builder pattern


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


## GETPaymentOptions

> GETPaymentOptions200Response GETPaymentOptions(ctx).Execute()

List all payment options



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
    resp, r, err := apiClient.PaymentOptionsApi.GETPaymentOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsApi.GETPaymentOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaymentOptions`: GETPaymentOptions200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentOptionsApi.GETPaymentOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentOptionsRequest struct via the builder pattern


### Return type

[**GETPaymentOptions200Response**](GETPaymentOptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPaymentOptionsPaymentOptionId

> GETPaymentOptionsPaymentOptionId200Response GETPaymentOptionsPaymentOptionId(ctx, paymentOptionId).Execute()

Retrieve a payment option



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
    paymentOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentOptionsApi.GETPaymentOptionsPaymentOptionId(context.Background(), paymentOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsApi.GETPaymentOptionsPaymentOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaymentOptionsPaymentOptionId`: GETPaymentOptionsPaymentOptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentOptionsApi.GETPaymentOptionsPaymentOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentOptionsPaymentOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPaymentOptionsPaymentOptionId200Response**](GETPaymentOptionsPaymentOptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPaymentOptionsPaymentOptionId

> PATCHPaymentOptionsPaymentOptionId200Response PATCHPaymentOptionsPaymentOptionId(ctx, paymentOptionId).PaymentOptionUpdate(paymentOptionUpdate).Execute()

Update a payment option



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
    paymentOptionUpdate := *openapiclient.NewPaymentOptionUpdate(*openapiclient.NewPaymentOptionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPaymentOptionsPaymentOptionId200ResponseDataAttributes())) // PaymentOptionUpdate | 
    paymentOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentOptionsApi.PATCHPaymentOptionsPaymentOptionId(context.Background(), paymentOptionId).PaymentOptionUpdate(paymentOptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsApi.PATCHPaymentOptionsPaymentOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPaymentOptionsPaymentOptionId`: PATCHPaymentOptionsPaymentOptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentOptionsApi.PATCHPaymentOptionsPaymentOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPaymentOptionsPaymentOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentOptionUpdate** | [**PaymentOptionUpdate**](PaymentOptionUpdate.md) |  | 


### Return type

[**PATCHPaymentOptionsPaymentOptionId200Response**](PATCHPaymentOptionsPaymentOptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPaymentOptions

> POSTPaymentOptions201Response POSTPaymentOptions(ctx).PaymentOptionCreate(paymentOptionCreate).Execute()

Create a payment option



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
    paymentOptionCreate := *openapiclient.NewPaymentOptionCreate(*openapiclient.NewPaymentOptionCreateData(interface{}(123), *openapiclient.NewPOSTPaymentOptions201ResponseDataAttributes(interface{}(stripe_payments), interface{}({application_fee_amount=1000, on_behalf_of=pm_xxx})))) // PaymentOptionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentOptionsApi.POSTPaymentOptions(context.Background()).PaymentOptionCreate(paymentOptionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsApi.POSTPaymentOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPaymentOptions`: POSTPaymentOptions201Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentOptionsApi.POSTPaymentOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPaymentOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentOptionCreate** | [**PaymentOptionCreate**](PaymentOptionCreate.md) |  | 

### Return type

[**POSTPaymentOptions201Response**](POSTPaymentOptions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

