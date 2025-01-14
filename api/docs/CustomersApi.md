# \CustomersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomersCustomerId**](CustomersApi.md#DELETECustomersCustomerId) | **Delete** /customers/{customerId} | Delete a customer
[**GETCouponRecipientIdCustomer**](CustomersApi.md#GETCouponRecipientIdCustomer) | **Get** /coupon_recipients/{couponRecipientId}/customer | Retrieve the customer associated to the coupon recipient
[**GETCustomerAddressIdCustomer**](CustomersApi.md#GETCustomerAddressIdCustomer) | **Get** /customer_addresses/{customerAddressId}/customer | Retrieve the customer associated to the customer address
[**GETCustomerGroupIdCustomers**](CustomersApi.md#GETCustomerGroupIdCustomers) | **Get** /customer_groups/{customerGroupId}/customers | Retrieve the customers associated to the customer group
[**GETCustomerPasswordResetIdCustomer**](CustomersApi.md#GETCustomerPasswordResetIdCustomer) | **Get** /customer_password_resets/{customerPasswordResetId}/customer | Retrieve the customer associated to the customer password reset
[**GETCustomerPaymentSourceIdCustomer**](CustomersApi.md#GETCustomerPaymentSourceIdCustomer) | **Get** /customer_payment_sources/{customerPaymentSourceId}/customer | Retrieve the customer associated to the customer payment source
[**GETCustomerSubscriptionIdCustomer**](CustomersApi.md#GETCustomerSubscriptionIdCustomer) | **Get** /customer_subscriptions/{customerSubscriptionId}/customer | Retrieve the customer associated to the customer subscription
[**GETCustomers**](CustomersApi.md#GETCustomers) | **Get** /customers | List all customers
[**GETCustomersCustomerId**](CustomersApi.md#GETCustomersCustomerId) | **Get** /customers/{customerId} | Retrieve a customer
[**GETGiftCardRecipientIdCustomer**](CustomersApi.md#GETGiftCardRecipientIdCustomer) | **Get** /gift_card_recipients/{giftCardRecipientId}/customer | Retrieve the customer associated to the gift card recipient
[**GETInStockSubscriptionIdCustomer**](CustomersApi.md#GETInStockSubscriptionIdCustomer) | **Get** /in_stock_subscriptions/{inStockSubscriptionId}/customer | Retrieve the customer associated to the in stock subscription
[**GETOrderIdCustomer**](CustomersApi.md#GETOrderIdCustomer) | **Get** /orders/{orderId}/customer | Retrieve the customer associated to the order
[**GETOrderSubscriptionIdCustomer**](CustomersApi.md#GETOrderSubscriptionIdCustomer) | **Get** /order_subscriptions/{orderSubscriptionId}/customer | Retrieve the customer associated to the order subscription
[**GETPriceIdJwtCustomer**](CustomersApi.md#GETPriceIdJwtCustomer) | **Get** /prices/{priceId}/jwt_customer | Retrieve the jwt customer associated to the price
[**GETReturnIdCustomer**](CustomersApi.md#GETReturnIdCustomer) | **Get** /returns/{returnId}/customer | Retrieve the customer associated to the return
[**GETSkuIdJwtCustomer**](CustomersApi.md#GETSkuIdJwtCustomer) | **Get** /skus/{skuId}/jwt_customer | Retrieve the jwt customer associated to the SKU
[**GETSkuListIdCustomer**](CustomersApi.md#GETSkuListIdCustomer) | **Get** /sku_lists/{skuListId}/customer | Retrieve the customer associated to the SKU list
[**PATCHCustomersCustomerId**](CustomersApi.md#PATCHCustomersCustomerId) | **Patch** /customers/{customerId} | Update a customer
[**POSTCustomers**](CustomersApi.md#POSTCustomers) | **Post** /customers | Create a customer



## DELETECustomersCustomerId

> DELETECustomersCustomerId(ctx, customerId).Execute()

Delete a customer



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
    r, err := apiClient.CustomersApi.DELETECustomersCustomerId(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.DELETECustomersCustomerId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETECustomersCustomerIdRequest struct via the builder pattern


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


## GETCouponRecipientIdCustomer

> GETCouponRecipientIdCustomer(ctx, couponRecipientId).Execute()

Retrieve the customer associated to the coupon recipient



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
    couponRecipientId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETCouponRecipientIdCustomer(context.Background(), couponRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETCouponRecipientIdCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponRecipientId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponRecipientIdCustomerRequest struct via the builder pattern


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


## GETCustomerAddressIdCustomer

> GETCustomerAddressIdCustomer(ctx, customerAddressId).Execute()

Retrieve the customer associated to the customer address



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
    customerAddressId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETCustomerAddressIdCustomer(context.Background(), customerAddressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETCustomerAddressIdCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerAddressId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerAddressIdCustomerRequest struct via the builder pattern


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


## GETCustomerGroupIdCustomers

> GETCustomerGroupIdCustomers(ctx, customerGroupId).Execute()

Retrieve the customers associated to the customer group



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
    customerGroupId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETCustomerGroupIdCustomers(context.Background(), customerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETCustomerGroupIdCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerGroupId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerGroupIdCustomersRequest struct via the builder pattern


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


## GETCustomerPasswordResetIdCustomer

> GETCustomerPasswordResetIdCustomer(ctx, customerPasswordResetId).Execute()

Retrieve the customer associated to the customer password reset



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
    customerPasswordResetId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETCustomerPasswordResetIdCustomer(context.Background(), customerPasswordResetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETCustomerPasswordResetIdCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPasswordResetId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPasswordResetIdCustomerRequest struct via the builder pattern


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


## GETCustomerPaymentSourceIdCustomer

> GETCustomerPaymentSourceIdCustomer(ctx, customerPaymentSourceId).Execute()

Retrieve the customer associated to the customer payment source



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
    r, err := apiClient.CustomersApi.GETCustomerPaymentSourceIdCustomer(context.Background(), customerPaymentSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETCustomerPaymentSourceIdCustomer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCustomerPaymentSourceIdCustomerRequest struct via the builder pattern


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


## GETCustomerSubscriptionIdCustomer

> GETCustomerSubscriptionIdCustomer(ctx, customerSubscriptionId).Execute()

Retrieve the customer associated to the customer subscription



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
    r, err := apiClient.CustomersApi.GETCustomerSubscriptionIdCustomer(context.Background(), customerSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETCustomerSubscriptionIdCustomer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCustomerSubscriptionIdCustomerRequest struct via the builder pattern


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


## GETCustomers

> GETCustomers200Response GETCustomers(ctx).Execute()

List all customers



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
    resp, r, err := apiClient.CustomersApi.GETCustomers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomers`: GETCustomers200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GETCustomers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomersRequest struct via the builder pattern


### Return type

[**GETCustomers200Response**](GETCustomers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomersCustomerId

> GETCustomersCustomerId200Response GETCustomersCustomerId(ctx, customerId).Execute()

Retrieve a customer



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
    resp, r, err := apiClient.CustomersApi.GETCustomersCustomerId(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETCustomersCustomerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomersCustomerId`: GETCustomersCustomerId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GETCustomersCustomerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomersCustomerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCustomersCustomerId200Response**](GETCustomersCustomerId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETGiftCardRecipientIdCustomer

> GETGiftCardRecipientIdCustomer(ctx, giftCardRecipientId).Execute()

Retrieve the customer associated to the gift card recipient



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
    giftCardRecipientId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETGiftCardRecipientIdCustomer(context.Background(), giftCardRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETGiftCardRecipientIdCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardRecipientId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETGiftCardRecipientIdCustomerRequest struct via the builder pattern


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


## GETInStockSubscriptionIdCustomer

> GETInStockSubscriptionIdCustomer(ctx, inStockSubscriptionId).Execute()

Retrieve the customer associated to the in stock subscription



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
    inStockSubscriptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETInStockSubscriptionIdCustomer(context.Background(), inStockSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETInStockSubscriptionIdCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inStockSubscriptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInStockSubscriptionIdCustomerRequest struct via the builder pattern


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


## GETOrderIdCustomer

> GETOrderIdCustomer(ctx, orderId).Execute()

Retrieve the customer associated to the order



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
    r, err := apiClient.CustomersApi.GETOrderIdCustomer(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETOrderIdCustomer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdCustomerRequest struct via the builder pattern


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


## GETOrderSubscriptionIdCustomer

> GETOrderSubscriptionIdCustomer(ctx, orderSubscriptionId).Execute()

Retrieve the customer associated to the order subscription



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
    r, err := apiClient.CustomersApi.GETOrderSubscriptionIdCustomer(context.Background(), orderSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETOrderSubscriptionIdCustomer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderSubscriptionIdCustomerRequest struct via the builder pattern


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


## GETPriceIdJwtCustomer

> GETPriceIdJwtCustomer(ctx, priceId).Execute()

Retrieve the jwt customer associated to the price



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
    priceId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETPriceIdJwtCustomer(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETPriceIdJwtCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceIdJwtCustomerRequest struct via the builder pattern


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


## GETReturnIdCustomer

> GETReturnIdCustomer(ctx, returnId).Execute()

Retrieve the customer associated to the return



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
    returnId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETReturnIdCustomer(context.Background(), returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETReturnIdCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnIdCustomerRequest struct via the builder pattern


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


## GETSkuIdJwtCustomer

> GETSkuIdJwtCustomer(ctx, skuId).Execute()

Retrieve the jwt customer associated to the SKU



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
    skuId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETSkuIdJwtCustomer(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETSkuIdJwtCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuIdJwtCustomerRequest struct via the builder pattern


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


## GETSkuListIdCustomer

> GETSkuListIdCustomer(ctx, skuListId).Execute()

Retrieve the customer associated to the SKU list



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
    skuListId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomersApi.GETSkuListIdCustomer(context.Background(), skuListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GETSkuListIdCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListIdCustomerRequest struct via the builder pattern


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


## PATCHCustomersCustomerId

> PATCHCustomersCustomerId200Response PATCHCustomersCustomerId(ctx, customerId).CustomerUpdate(customerUpdate).Execute()

Update a customer



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
    customerUpdate := *openapiclient.NewCustomerUpdate(*openapiclient.NewCustomerUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCustomersCustomerId200ResponseDataAttributes())) // CustomerUpdate | 
    customerId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.PATCHCustomersCustomerId(context.Background(), customerId).CustomerUpdate(customerUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.PATCHCustomersCustomerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCustomersCustomerId`: PATCHCustomersCustomerId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.PATCHCustomersCustomerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomersCustomerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerUpdate** | [**CustomerUpdate**](CustomerUpdate.md) |  | 


### Return type

[**PATCHCustomersCustomerId200Response**](PATCHCustomersCustomerId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCustomers

> POSTCustomers201Response POSTCustomers(ctx).CustomerCreate(customerCreate).Execute()

Create a customer



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
    customerCreate := *openapiclient.NewCustomerCreate(*openapiclient.NewCustomerCreateData(interface{}(123), *openapiclient.NewPOSTCustomers201ResponseDataAttributes(interface{}(john@example.com)))) // CustomerCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.POSTCustomers(context.Background()).CustomerCreate(customerCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.POSTCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCustomers`: POSTCustomers201Response
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.POSTCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerCreate** | [**CustomerCreate**](CustomerCreate.md) |  | 

### Return type

[**POSTCustomers201Response**](POSTCustomers201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

