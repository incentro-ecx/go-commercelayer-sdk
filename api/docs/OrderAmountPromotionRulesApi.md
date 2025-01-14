# \OrderAmountPromotionRulesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId**](OrderAmountPromotionRulesApi.md#DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId) | **Delete** /order_amount_promotion_rules/{orderAmountPromotionRuleId} | Delete an order amount promotion rule
[**GETBuyXPayYPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETBuyXPayYPromotionIdOrderAmountPromotionRule) | **Get** /buy_x_pay_y_promotions/{buyXPayYPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the buy x pay y promotion
[**GETExternalPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETExternalPromotionIdOrderAmountPromotionRule) | **Get** /external_promotions/{externalPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the external promotion
[**GETFixedAmountPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETFixedAmountPromotionIdOrderAmountPromotionRule) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the fixed amount promotion
[**GETFixedPricePromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETFixedPricePromotionIdOrderAmountPromotionRule) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the fixed price promotion
[**GETFreeGiftPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETFreeGiftPromotionIdOrderAmountPromotionRule) | **Get** /free_gift_promotions/{freeGiftPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the free gift promotion
[**GETFreeShippingPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETFreeShippingPromotionIdOrderAmountPromotionRule) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the free shipping promotion
[**GETOrderAmountPromotionRules**](OrderAmountPromotionRulesApi.md#GETOrderAmountPromotionRules) | **Get** /order_amount_promotion_rules | List all order amount promotion rules
[**GETOrderAmountPromotionRulesOrderAmountPromotionRuleId**](OrderAmountPromotionRulesApi.md#GETOrderAmountPromotionRulesOrderAmountPromotionRuleId) | **Get** /order_amount_promotion_rules/{orderAmountPromotionRuleId} | Retrieve an order amount promotion rule
[**GETPercentageDiscountPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETPercentageDiscountPromotionIdOrderAmountPromotionRule) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the percentage discount promotion
[**GETPromotionIdOrderAmountPromotionRule**](OrderAmountPromotionRulesApi.md#GETPromotionIdOrderAmountPromotionRule) | **Get** /promotions/{promotionId}/order_amount_promotion_rule | Retrieve the order amount promotion rule associated to the promotion
[**PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId**](OrderAmountPromotionRulesApi.md#PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId) | **Patch** /order_amount_promotion_rules/{orderAmountPromotionRuleId} | Update an order amount promotion rule
[**POSTOrderAmountPromotionRules**](OrderAmountPromotionRulesApi.md#POSTOrderAmountPromotionRules) | **Post** /order_amount_promotion_rules | Create an order amount promotion rule



## DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId

> DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId(ctx, orderAmountPromotionRuleId).Execute()

Delete an order amount promotion rule



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
    orderAmountPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderAmountPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct via the builder pattern


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


## GETBuyXPayYPromotionIdOrderAmountPromotionRule

> GETBuyXPayYPromotionIdOrderAmountPromotionRule(ctx, buyXPayYPromotionId).Execute()

Retrieve the order amount promotion rule associated to the buy x pay y promotion



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
    buyXPayYPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.GETBuyXPayYPromotionIdOrderAmountPromotionRule(context.Background(), buyXPayYPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETBuyXPayYPromotionIdOrderAmountPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyXPayYPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBuyXPayYPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


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


## GETExternalPromotionIdOrderAmountPromotionRule

> GETExternalPromotionIdOrderAmountPromotionRule(ctx, externalPromotionId).Execute()

Retrieve the order amount promotion rule associated to the external promotion



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
    externalPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.GETExternalPromotionIdOrderAmountPromotionRule(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETExternalPromotionIdOrderAmountPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


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


## GETFixedAmountPromotionIdOrderAmountPromotionRule

> GETFixedAmountPromotionIdOrderAmountPromotionRule(ctx, fixedAmountPromotionId).Execute()

Retrieve the order amount promotion rule associated to the fixed amount promotion



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
    fixedAmountPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.GETFixedAmountPromotionIdOrderAmountPromotionRule(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETFixedAmountPromotionIdOrderAmountPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedAmountPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


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


## GETFixedPricePromotionIdOrderAmountPromotionRule

> GETFixedPricePromotionIdOrderAmountPromotionRule(ctx, fixedPricePromotionId).Execute()

Retrieve the order amount promotion rule associated to the fixed price promotion



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
    fixedPricePromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.GETFixedPricePromotionIdOrderAmountPromotionRule(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETFixedPricePromotionIdOrderAmountPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedPricePromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


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


## GETFreeGiftPromotionIdOrderAmountPromotionRule

> GETFreeGiftPromotionIdOrderAmountPromotionRule(ctx, freeGiftPromotionId).Execute()

Retrieve the order amount promotion rule associated to the free gift promotion



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
    freeGiftPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.GETFreeGiftPromotionIdOrderAmountPromotionRule(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETFreeGiftPromotionIdOrderAmountPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeGiftPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


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


## GETFreeShippingPromotionIdOrderAmountPromotionRule

> GETFreeShippingPromotionIdOrderAmountPromotionRule(ctx, freeShippingPromotionId).Execute()

Retrieve the order amount promotion rule associated to the free shipping promotion



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
    freeShippingPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.GETFreeShippingPromotionIdOrderAmountPromotionRule(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETFreeShippingPromotionIdOrderAmountPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeShippingPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


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


## GETOrderAmountPromotionRules

> GETOrderAmountPromotionRules200Response GETOrderAmountPromotionRules(ctx).Execute()

List all order amount promotion rules



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
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETOrderAmountPromotionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETOrderAmountPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderAmountPromotionRules`: GETOrderAmountPromotionRules200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAmountPromotionRulesApi.GETOrderAmountPromotionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderAmountPromotionRulesRequest struct via the builder pattern


### Return type

[**GETOrderAmountPromotionRules200Response**](GETOrderAmountPromotionRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderAmountPromotionRulesOrderAmountPromotionRuleId

> GETOrderAmountPromotionRulesOrderAmountPromotionRuleId200Response GETOrderAmountPromotionRulesOrderAmountPromotionRuleId(ctx, orderAmountPromotionRuleId).Execute()

Retrieve an order amount promotion rule



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
    orderAmountPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.GETOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETOrderAmountPromotionRulesOrderAmountPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETOrderAmountPromotionRulesOrderAmountPromotionRuleId`: GETOrderAmountPromotionRulesOrderAmountPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAmountPromotionRulesApi.GETOrderAmountPromotionRulesOrderAmountPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderAmountPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETOrderAmountPromotionRulesOrderAmountPromotionRuleId200Response**](GETOrderAmountPromotionRulesOrderAmountPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPercentageDiscountPromotionIdOrderAmountPromotionRule

> GETPercentageDiscountPromotionIdOrderAmountPromotionRule(ctx, percentageDiscountPromotionId).Execute()

Retrieve the order amount promotion rule associated to the percentage discount promotion



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
    percentageDiscountPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.GETPercentageDiscountPromotionIdOrderAmountPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETPercentageDiscountPromotionIdOrderAmountPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**percentageDiscountPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


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


## GETPromotionIdOrderAmountPromotionRule

> GETPromotionIdOrderAmountPromotionRule(ctx, promotionId).Execute()

Retrieve the order amount promotion rule associated to the promotion



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
    promotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrderAmountPromotionRulesApi.GETPromotionIdOrderAmountPromotionRule(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.GETPromotionIdOrderAmountPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPromotionIdOrderAmountPromotionRuleRequest struct via the builder pattern


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


## PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId

> PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId200Response PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId(ctx, orderAmountPromotionRuleId).OrderAmountPromotionRuleUpdate(orderAmountPromotionRuleUpdate).Execute()

Update an order amount promotion rule



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
    orderAmountPromotionRuleUpdate := *openapiclient.NewOrderAmountPromotionRuleUpdate(*openapiclient.NewOrderAmountPromotionRuleUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId200ResponseDataAttributes())) // OrderAmountPromotionRuleUpdate | 
    orderAmountPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).OrderAmountPromotionRuleUpdate(orderAmountPromotionRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId`: PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAmountPromotionRulesApi.PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderAmountPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderAmountPromotionRuleUpdate** | [**OrderAmountPromotionRuleUpdate**](OrderAmountPromotionRuleUpdate.md) |  | 


### Return type

[**PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId200Response**](PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTOrderAmountPromotionRules

> POSTOrderAmountPromotionRules201Response POSTOrderAmountPromotionRules(ctx).OrderAmountPromotionRuleCreate(orderAmountPromotionRuleCreate).Execute()

Create an order amount promotion rule



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
    orderAmountPromotionRuleCreate := *openapiclient.NewOrderAmountPromotionRuleCreate(*openapiclient.NewOrderAmountPromotionRuleCreateData(interface{}(123), *openapiclient.NewPOSTOrderAmountPromotionRules201ResponseDataAttributes())) // OrderAmountPromotionRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderAmountPromotionRulesApi.POSTOrderAmountPromotionRules(context.Background()).OrderAmountPromotionRuleCreate(orderAmountPromotionRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderAmountPromotionRulesApi.POSTOrderAmountPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTOrderAmountPromotionRules`: POSTOrderAmountPromotionRules201Response
    fmt.Fprintf(os.Stdout, "Response from `OrderAmountPromotionRulesApi.POSTOrderAmountPromotionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTOrderAmountPromotionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderAmountPromotionRuleCreate** | [**OrderAmountPromotionRuleCreate**](OrderAmountPromotionRuleCreate.md) |  | 

### Return type

[**POSTOrderAmountPromotionRules201Response**](POSTOrderAmountPromotionRules201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

