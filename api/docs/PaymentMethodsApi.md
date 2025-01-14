# \PaymentMethodsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPaymentMethodsPaymentMethodId**](PaymentMethodsApi.md#DELETEPaymentMethodsPaymentMethodId) | **Delete** /payment_methods/{paymentMethodId} | Delete a payment method
[**GETAdyenGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETAdyenGatewayIdPaymentMethods) | **Get** /adyen_gateways/{adyenGatewayId}/payment_methods | Retrieve the payment methods associated to the adyen gateway
[**GETAxerveGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETAxerveGatewayIdPaymentMethods) | **Get** /axerve_gateways/{axerveGatewayId}/payment_methods | Retrieve the payment methods associated to the axerve gateway
[**GETBraintreeGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETBraintreeGatewayIdPaymentMethods) | **Get** /braintree_gateways/{braintreeGatewayId}/payment_methods | Retrieve the payment methods associated to the braintree gateway
[**GETCheckoutComGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETCheckoutComGatewayIdPaymentMethods) | **Get** /checkout_com_gateways/{checkoutComGatewayId}/payment_methods | Retrieve the payment methods associated to the checkout.com gateway
[**GETCustomerPaymentSourceIdPaymentMethod**](PaymentMethodsApi.md#GETCustomerPaymentSourceIdPaymentMethod) | **Get** /customer_payment_sources/{customerPaymentSourceId}/payment_method | Retrieve the payment method associated to the customer payment source
[**GETExternalGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETExternalGatewayIdPaymentMethods) | **Get** /external_gateways/{externalGatewayId}/payment_methods | Retrieve the payment methods associated to the external gateway
[**GETKlarnaGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETKlarnaGatewayIdPaymentMethods) | **Get** /klarna_gateways/{klarnaGatewayId}/payment_methods | Retrieve the payment methods associated to the klarna gateway
[**GETManualGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETManualGatewayIdPaymentMethods) | **Get** /manual_gateways/{manualGatewayId}/payment_methods | Retrieve the payment methods associated to the manual gateway
[**GETOrderIdAvailablePaymentMethods**](PaymentMethodsApi.md#GETOrderIdAvailablePaymentMethods) | **Get** /orders/{orderId}/available_payment_methods | Retrieve the available payment methods associated to the order
[**GETOrderIdPaymentMethod**](PaymentMethodsApi.md#GETOrderIdPaymentMethod) | **Get** /orders/{orderId}/payment_method | Retrieve the payment method associated to the order
[**GETPaymentGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETPaymentGatewayIdPaymentMethods) | **Get** /payment_gateways/{paymentGatewayId}/payment_methods | Retrieve the payment methods associated to the payment gateway
[**GETPaymentMethods**](PaymentMethodsApi.md#GETPaymentMethods) | **Get** /payment_methods | List all payment methods
[**GETPaymentMethodsPaymentMethodId**](PaymentMethodsApi.md#GETPaymentMethodsPaymentMethodId) | **Get** /payment_methods/{paymentMethodId} | Retrieve a payment method
[**GETPaypalGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETPaypalGatewayIdPaymentMethods) | **Get** /paypal_gateways/{paypalGatewayId}/payment_methods | Retrieve the payment methods associated to the paypal gateway
[**GETSatispayGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETSatispayGatewayIdPaymentMethods) | **Get** /satispay_gateways/{satispayGatewayId}/payment_methods | Retrieve the payment methods associated to the satispay gateway
[**GETStoreIdPaymentMethods**](PaymentMethodsApi.md#GETStoreIdPaymentMethods) | **Get** /stores/{storeId}/payment_methods | Retrieve the payment methods associated to the store
[**GETStripeGatewayIdPaymentMethods**](PaymentMethodsApi.md#GETStripeGatewayIdPaymentMethods) | **Get** /stripe_gateways/{stripeGatewayId}/payment_methods | Retrieve the payment methods associated to the stripe gateway
[**PATCHPaymentMethodsPaymentMethodId**](PaymentMethodsApi.md#PATCHPaymentMethodsPaymentMethodId) | **Patch** /payment_methods/{paymentMethodId} | Update a payment method
[**POSTPaymentMethods**](PaymentMethodsApi.md#POSTPaymentMethods) | **Post** /payment_methods | Create a payment method



## DELETEPaymentMethodsPaymentMethodId

> DELETEPaymentMethodsPaymentMethodId(ctx, paymentMethodId).Execute()

Delete a payment method



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
    r, err := apiClient.PaymentMethodsApi.DELETEPaymentMethodsPaymentMethodId(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.DELETEPaymentMethodsPaymentMethodId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEPaymentMethodsPaymentMethodIdRequest struct via the builder pattern


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


## GETAdyenGatewayIdPaymentMethods

> GETAdyenGatewayIdPaymentMethods(ctx, adyenGatewayId).Execute()

Retrieve the payment methods associated to the adyen gateway



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
    adyenGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETAdyenGatewayIdPaymentMethods(context.Background(), adyenGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETAdyenGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdyenGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETAxerveGatewayIdPaymentMethods

> GETAxerveGatewayIdPaymentMethods(ctx, axerveGatewayId).Execute()

Retrieve the payment methods associated to the axerve gateway



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
    axerveGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETAxerveGatewayIdPaymentMethods(context.Background(), axerveGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETAxerveGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**axerveGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAxerveGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETBraintreeGatewayIdPaymentMethods

> GETBraintreeGatewayIdPaymentMethods(ctx, braintreeGatewayId).Execute()

Retrieve the payment methods associated to the braintree gateway



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
    braintreeGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETBraintreeGatewayIdPaymentMethods(context.Background(), braintreeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETBraintreeGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreeGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreeGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETCheckoutComGatewayIdPaymentMethods

> GETCheckoutComGatewayIdPaymentMethods(ctx, checkoutComGatewayId).Execute()

Retrieve the payment methods associated to the checkout.com gateway



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
    checkoutComGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETCheckoutComGatewayIdPaymentMethods(context.Background(), checkoutComGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETCheckoutComGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutComGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCheckoutComGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETCustomerPaymentSourceIdPaymentMethod

> GETCustomerPaymentSourceIdPaymentMethod(ctx, customerPaymentSourceId).Execute()

Retrieve the payment method associated to the customer payment source



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
    r, err := apiClient.PaymentMethodsApi.GETCustomerPaymentSourceIdPaymentMethod(context.Background(), customerPaymentSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETCustomerPaymentSourceIdPaymentMethod``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCustomerPaymentSourceIdPaymentMethodRequest struct via the builder pattern


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


## GETExternalGatewayIdPaymentMethods

> GETExternalGatewayIdPaymentMethods(ctx, externalGatewayId).Execute()

Retrieve the payment methods associated to the external gateway



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
    externalGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETExternalGatewayIdPaymentMethods(context.Background(), externalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETExternalGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETKlarnaGatewayIdPaymentMethods

> GETKlarnaGatewayIdPaymentMethods(ctx, klarnaGatewayId).Execute()

Retrieve the payment methods associated to the klarna gateway



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
    klarnaGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETKlarnaGatewayIdPaymentMethods(context.Background(), klarnaGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETKlarnaGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETManualGatewayIdPaymentMethods

> GETManualGatewayIdPaymentMethods(ctx, manualGatewayId).Execute()

Retrieve the payment methods associated to the manual gateway



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
    manualGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETManualGatewayIdPaymentMethods(context.Background(), manualGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETManualGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETManualGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETOrderIdAvailablePaymentMethods

> GETOrderIdAvailablePaymentMethods(ctx, orderId).Execute()

Retrieve the available payment methods associated to the order



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
    r, err := apiClient.PaymentMethodsApi.GETOrderIdAvailablePaymentMethods(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETOrderIdAvailablePaymentMethods``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdAvailablePaymentMethodsRequest struct via the builder pattern


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


## GETOrderIdPaymentMethod

> GETOrderIdPaymentMethod(ctx, orderId).Execute()

Retrieve the payment method associated to the order



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
    r, err := apiClient.PaymentMethodsApi.GETOrderIdPaymentMethod(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETOrderIdPaymentMethod``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdPaymentMethodRequest struct via the builder pattern


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


## GETPaymentGatewayIdPaymentMethods

> GETPaymentGatewayIdPaymentMethods(ctx, paymentGatewayId).Execute()

Retrieve the payment methods associated to the payment gateway



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
    r, err := apiClient.PaymentMethodsApi.GETPaymentGatewayIdPaymentMethods(context.Background(), paymentGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETPaymentGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETPaymentMethods

> GETPaymentMethods200Response GETPaymentMethods(ctx).Execute()

List all payment methods



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
    resp, r, err := apiClient.PaymentMethodsApi.GETPaymentMethods(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaymentMethods`: GETPaymentMethods200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsApi.GETPaymentMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentMethodsRequest struct via the builder pattern


### Return type

[**GETPaymentMethods200Response**](GETPaymentMethods200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPaymentMethodsPaymentMethodId

> GETPaymentMethodsPaymentMethodId200Response GETPaymentMethodsPaymentMethodId(ctx, paymentMethodId).Execute()

Retrieve a payment method



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
    resp, r, err := apiClient.PaymentMethodsApi.GETPaymentMethodsPaymentMethodId(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETPaymentMethodsPaymentMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaymentMethodsPaymentMethodId`: GETPaymentMethodsPaymentMethodId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsApi.GETPaymentMethodsPaymentMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaymentMethodsPaymentMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPaymentMethodsPaymentMethodId200Response**](GETPaymentMethodsPaymentMethodId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPaypalGatewayIdPaymentMethods

> GETPaypalGatewayIdPaymentMethods(ctx, paypalGatewayId).Execute()

Retrieve the payment methods associated to the paypal gateway



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
    r, err := apiClient.PaymentMethodsApi.GETPaypalGatewayIdPaymentMethods(context.Background(), paypalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETPaypalGatewayIdPaymentMethods``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPaypalGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETSatispayGatewayIdPaymentMethods

> GETSatispayGatewayIdPaymentMethods(ctx, satispayGatewayId).Execute()

Retrieve the payment methods associated to the satispay gateway



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
    satispayGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETSatispayGatewayIdPaymentMethods(context.Background(), satispayGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETSatispayGatewayIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSatispayGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## GETStoreIdPaymentMethods

> GETStoreIdPaymentMethods(ctx, storeId).Execute()

Retrieve the payment methods associated to the store



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
    storeId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentMethodsApi.GETStoreIdPaymentMethods(context.Background(), storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETStoreIdPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStoreIdPaymentMethodsRequest struct via the builder pattern


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


## GETStripeGatewayIdPaymentMethods

> GETStripeGatewayIdPaymentMethods(ctx, stripeGatewayId).Execute()

Retrieve the payment methods associated to the stripe gateway



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
    r, err := apiClient.PaymentMethodsApi.GETStripeGatewayIdPaymentMethods(context.Background(), stripeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.GETStripeGatewayIdPaymentMethods``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStripeGatewayIdPaymentMethodsRequest struct via the builder pattern


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


## PATCHPaymentMethodsPaymentMethodId

> PATCHPaymentMethodsPaymentMethodId200Response PATCHPaymentMethodsPaymentMethodId(ctx, paymentMethodId).PaymentMethodUpdate(paymentMethodUpdate).Execute()

Update a payment method



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
    paymentMethodUpdate := *openapiclient.NewPaymentMethodUpdate(*openapiclient.NewPaymentMethodUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes())) // PaymentMethodUpdate | 
    paymentMethodId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsApi.PATCHPaymentMethodsPaymentMethodId(context.Background(), paymentMethodId).PaymentMethodUpdate(paymentMethodUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.PATCHPaymentMethodsPaymentMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPaymentMethodsPaymentMethodId`: PATCHPaymentMethodsPaymentMethodId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsApi.PATCHPaymentMethodsPaymentMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPaymentMethodsPaymentMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodUpdate** | [**PaymentMethodUpdate**](PaymentMethodUpdate.md) |  | 


### Return type

[**PATCHPaymentMethodsPaymentMethodId200Response**](PATCHPaymentMethodsPaymentMethodId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPaymentMethods

> POSTPaymentMethods201Response POSTPaymentMethods(ctx).PaymentMethodCreate(paymentMethodCreate).Execute()

Create a payment method



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
    paymentMethodCreate := *openapiclient.NewPaymentMethodCreate(*openapiclient.NewPaymentMethodCreateData(interface{}(123), *openapiclient.NewPOSTPaymentMethods201ResponseDataAttributes(interface{}(stripe_payments), interface{}(0)))) // PaymentMethodCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsApi.POSTPaymentMethods(context.Background()).PaymentMethodCreate(paymentMethodCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsApi.POSTPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPaymentMethods`: POSTPaymentMethods201Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsApi.POSTPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodCreate** | [**PaymentMethodCreate**](PaymentMethodCreate.md) |  | 

### Return type

[**POSTPaymentMethods201Response**](POSTPaymentMethods201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

