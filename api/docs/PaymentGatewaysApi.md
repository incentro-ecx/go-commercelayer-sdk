# \PaymentGatewaysApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETAdyenPaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETAdyenPaymentIdPaymentGateway) | **Get** /adyen_payments/{adyenPaymentId}/payment_gateway | Retrieve the payment gateway associated to the adyen payment
[**GETAxervePaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETAxervePaymentIdPaymentGateway) | **Get** /axerve_payments/{axervePaymentId}/payment_gateway | Retrieve the payment gateway associated to the axerve payment
[**GETBraintreePaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETBraintreePaymentIdPaymentGateway) | **Get** /braintree_payments/{braintreePaymentId}/payment_gateway | Retrieve the payment gateway associated to the braintree payment
[**GETCheckoutComPaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETCheckoutComPaymentIdPaymentGateway) | **Get** /checkout_com_payments/{checkoutComPaymentId}/payment_gateway | Retrieve the payment gateway associated to the checkout.com payment
[**GETExternalPaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETExternalPaymentIdPaymentGateway) | **Get** /external_payments/{externalPaymentId}/payment_gateway | Retrieve the payment gateway associated to the external payment
[**GETKlarnaPaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETKlarnaPaymentIdPaymentGateway) | **Get** /klarna_payments/{klarnaPaymentId}/payment_gateway | Retrieve the payment gateway associated to the klarna payment
[**GETPaymentGateways**](PaymentGatewaysApi.md#GETPaymentGateways) | **Get** /payment_gateways | List all payment gateways
[**GETPaymentGatewaysPaymentGatewayId**](PaymentGatewaysApi.md#GETPaymentGatewaysPaymentGatewayId) | **Get** /payment_gateways/{paymentGatewayId} | Retrieve a payment gateway
[**GETPaymentMethodIdPaymentGateway**](PaymentGatewaysApi.md#GETPaymentMethodIdPaymentGateway) | **Get** /payment_methods/{paymentMethodId}/payment_gateway | Retrieve the payment gateway associated to the payment method
[**GETPaypalPaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETPaypalPaymentIdPaymentGateway) | **Get** /paypal_payments/{paypalPaymentId}/payment_gateway | Retrieve the payment gateway associated to the paypal payment
[**GETSatispayPaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETSatispayPaymentIdPaymentGateway) | **Get** /satispay_payments/{satispayPaymentId}/payment_gateway | Retrieve the payment gateway associated to the satispay payment
[**GETStripePaymentIdPaymentGateway**](PaymentGatewaysApi.md#GETStripePaymentIdPaymentGateway) | **Get** /stripe_payments/{stripePaymentId}/payment_gateway | Retrieve the payment gateway associated to the stripe payment



## GETAdyenPaymentIdPaymentGateway

> GETAdyenPaymentIdPaymentGateway(ctx, adyenPaymentId).Execute()

Retrieve the payment gateway associated to the adyen payment



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
    adyenPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentGatewaysApi.GETAdyenPaymentIdPaymentGateway(context.Background(), adyenPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETAdyenPaymentIdPaymentGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdyenPaymentIdPaymentGatewayRequest struct via the builder pattern


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


## GETAxervePaymentIdPaymentGateway

> GETAxervePaymentIdPaymentGateway(ctx, axervePaymentId).Execute()

Retrieve the payment gateway associated to the axerve payment



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
    axervePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentGatewaysApi.GETAxervePaymentIdPaymentGateway(context.Background(), axervePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETAxervePaymentIdPaymentGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**axervePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAxervePaymentIdPaymentGatewayRequest struct via the builder pattern


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


## GETBraintreePaymentIdPaymentGateway

> GETBraintreePaymentIdPaymentGateway(ctx, braintreePaymentId).Execute()

Retrieve the payment gateway associated to the braintree payment



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
    braintreePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentGatewaysApi.GETBraintreePaymentIdPaymentGateway(context.Background(), braintreePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETBraintreePaymentIdPaymentGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreePaymentIdPaymentGatewayRequest struct via the builder pattern


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


## GETCheckoutComPaymentIdPaymentGateway

> GETCheckoutComPaymentIdPaymentGateway(ctx, checkoutComPaymentId).Execute()

Retrieve the payment gateway associated to the checkout.com payment



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
    checkoutComPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentGatewaysApi.GETCheckoutComPaymentIdPaymentGateway(context.Background(), checkoutComPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETCheckoutComPaymentIdPaymentGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutComPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCheckoutComPaymentIdPaymentGatewayRequest struct via the builder pattern


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


## GETExternalPaymentIdPaymentGateway

> GETExternalPaymentIdPaymentGateway(ctx, externalPaymentId).Execute()

Retrieve the payment gateway associated to the external payment



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
    r, err := apiClient.PaymentGatewaysApi.GETExternalPaymentIdPaymentGateway(context.Background(), externalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETExternalPaymentIdPaymentGateway``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPaymentIdPaymentGatewayRequest struct via the builder pattern


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


## GETKlarnaPaymentIdPaymentGateway

> GETKlarnaPaymentIdPaymentGateway(ctx, klarnaPaymentId).Execute()

Retrieve the payment gateway associated to the klarna payment



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
    klarnaPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentGatewaysApi.GETKlarnaPaymentIdPaymentGateway(context.Background(), klarnaPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETKlarnaPaymentIdPaymentGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaPaymentIdPaymentGatewayRequest struct via the builder pattern


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


## GETPaymentGateways

> GETPaymentGateways200Response GETPaymentGateways(ctx).Execute()

List all payment gateways



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
    resp, r, err := apiClient.PaymentGatewaysApi.GETPaymentGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETPaymentGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaymentGateways`: GETPaymentGateways200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentGatewaysApi.GETPaymentGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentGatewaysRequest struct via the builder pattern


### Return type

[**GETPaymentGateways200Response**](GETPaymentGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPaymentGatewaysPaymentGatewayId

> GETPaymentGatewaysPaymentGatewayId200Response GETPaymentGatewaysPaymentGatewayId(ctx, paymentGatewayId).Execute()

Retrieve a payment gateway



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
    paymentGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentGatewaysApi.GETPaymentGatewaysPaymentGatewayId(context.Background(), paymentGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETPaymentGatewaysPaymentGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaymentGatewaysPaymentGatewayId`: GETPaymentGatewaysPaymentGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentGatewaysApi.GETPaymentGatewaysPaymentGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentGatewaysPaymentGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPaymentGatewaysPaymentGatewayId200Response**](GETPaymentGatewaysPaymentGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPaymentMethodIdPaymentGateway

> GETPaymentMethodIdPaymentGateway(ctx, paymentMethodId).Execute()

Retrieve the payment gateway associated to the payment method



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
    paymentMethodId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentGatewaysApi.GETPaymentMethodIdPaymentGateway(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETPaymentMethodIdPaymentGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentMethodIdPaymentGatewayRequest struct via the builder pattern


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


## GETPaypalPaymentIdPaymentGateway

> GETPaypalPaymentIdPaymentGateway(ctx, paypalPaymentId).Execute()

Retrieve the payment gateway associated to the paypal payment



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
    r, err := apiClient.PaymentGatewaysApi.GETPaypalPaymentIdPaymentGateway(context.Background(), paypalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETPaypalPaymentIdPaymentGateway``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPaypalPaymentIdPaymentGatewayRequest struct via the builder pattern


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


## GETSatispayPaymentIdPaymentGateway

> GETSatispayPaymentIdPaymentGateway(ctx, satispayPaymentId).Execute()

Retrieve the payment gateway associated to the satispay payment



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
    satispayPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentGatewaysApi.GETSatispayPaymentIdPaymentGateway(context.Background(), satispayPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETSatispayPaymentIdPaymentGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSatispayPaymentIdPaymentGatewayRequest struct via the builder pattern


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


## GETStripePaymentIdPaymentGateway

> GETStripePaymentIdPaymentGateway(ctx, stripePaymentId).Execute()

Retrieve the payment gateway associated to the stripe payment



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
    r, err := apiClient.PaymentGatewaysApi.GETStripePaymentIdPaymentGateway(context.Background(), stripePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewaysApi.GETStripePaymentIdPaymentGateway``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStripePaymentIdPaymentGatewayRequest struct via the builder pattern


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

