# \CouponCodesPromotionRulesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId**](CouponCodesPromotionRulesApi.md#DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId) | **Delete** /coupon_codes_promotion_rules/{couponCodesPromotionRuleId} | Delete a coupon codes promotion rule
[**GETBuyXPayYPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETBuyXPayYPromotionIdCouponCodesPromotionRule) | **Get** /buy_x_pay_y_promotions/{buyXPayYPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the buy x pay y promotion
[**GETCouponCodesPromotionRules**](CouponCodesPromotionRulesApi.md#GETCouponCodesPromotionRules) | **Get** /coupon_codes_promotion_rules | List all coupon codes promotion rules
[**GETCouponCodesPromotionRulesCouponCodesPromotionRuleId**](CouponCodesPromotionRulesApi.md#GETCouponCodesPromotionRulesCouponCodesPromotionRuleId) | **Get** /coupon_codes_promotion_rules/{couponCodesPromotionRuleId} | Retrieve a coupon codes promotion rule
[**GETCouponIdPromotionRule**](CouponCodesPromotionRulesApi.md#GETCouponIdPromotionRule) | **Get** /coupons/{couponId}/promotion_rule | Retrieve the promotion rule associated to the coupon
[**GETExternalPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETExternalPromotionIdCouponCodesPromotionRule) | **Get** /external_promotions/{externalPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the external promotion
[**GETFixedAmountPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFixedAmountPromotionIdCouponCodesPromotionRule) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the fixed amount promotion
[**GETFixedPricePromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFixedPricePromotionIdCouponCodesPromotionRule) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the fixed price promotion
[**GETFlexPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFlexPromotionIdCouponCodesPromotionRule) | **Get** /flex_promotions/{flexPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the flex promotion
[**GETFreeGiftPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFreeGiftPromotionIdCouponCodesPromotionRule) | **Get** /free_gift_promotions/{freeGiftPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the free gift promotion
[**GETFreeShippingPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETFreeShippingPromotionIdCouponCodesPromotionRule) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the free shipping promotion
[**GETPercentageDiscountPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETPercentageDiscountPromotionIdCouponCodesPromotionRule) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the percentage discount promotion
[**GETPromotionIdCouponCodesPromotionRule**](CouponCodesPromotionRulesApi.md#GETPromotionIdCouponCodesPromotionRule) | **Get** /promotions/{promotionId}/coupon_codes_promotion_rule | Retrieve the coupon codes promotion rule associated to the promotion
[**PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId**](CouponCodesPromotionRulesApi.md#PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId) | **Patch** /coupon_codes_promotion_rules/{couponCodesPromotionRuleId} | Update a coupon codes promotion rule
[**POSTCouponCodesPromotionRules**](CouponCodesPromotionRulesApi.md#POSTCouponCodesPromotionRules) | **Post** /coupon_codes_promotion_rules | Create a coupon codes promotion rule



## DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId

> DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId(ctx, couponCodesPromotionRuleId).Execute()

Delete a coupon codes promotion rule



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
    couponCodesPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CouponCodesPromotionRulesApi.DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId(context.Background(), couponCodesPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.DELETECouponCodesPromotionRulesCouponCodesPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCodesPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECouponCodesPromotionRulesCouponCodesPromotionRuleIdRequest struct via the builder pattern


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


## GETBuyXPayYPromotionIdCouponCodesPromotionRule

> GETBuyXPayYPromotionIdCouponCodesPromotionRule(ctx, buyXPayYPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the buy x pay y promotion



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
    r, err := apiClient.CouponCodesPromotionRulesApi.GETBuyXPayYPromotionIdCouponCodesPromotionRule(context.Background(), buyXPayYPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETBuyXPayYPromotionIdCouponCodesPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETBuyXPayYPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## GETCouponCodesPromotionRules

> GETCouponCodesPromotionRules200Response GETCouponCodesPromotionRules(ctx).Execute()

List all coupon codes promotion rules



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
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETCouponCodesPromotionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETCouponCodesPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCouponCodesPromotionRules`: GETCouponCodesPromotionRules200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponCodesPromotionRulesApi.GETCouponCodesPromotionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponCodesPromotionRulesRequest struct via the builder pattern


### Return type

[**GETCouponCodesPromotionRules200Response**](GETCouponCodesPromotionRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCouponCodesPromotionRulesCouponCodesPromotionRuleId

> GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response GETCouponCodesPromotionRulesCouponCodesPromotionRuleId(ctx, couponCodesPromotionRuleId).Execute()

Retrieve a coupon codes promotion rule



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
    couponCodesPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.GETCouponCodesPromotionRulesCouponCodesPromotionRuleId(context.Background(), couponCodesPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETCouponCodesPromotionRulesCouponCodesPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCouponCodesPromotionRulesCouponCodesPromotionRuleId`: GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponCodesPromotionRulesApi.GETCouponCodesPromotionRulesCouponCodesPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCodesPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponCodesPromotionRulesCouponCodesPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response**](GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCouponIdPromotionRule

> GETCouponIdPromotionRule(ctx, couponId).Execute()

Retrieve the promotion rule associated to the coupon



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
    couponId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CouponCodesPromotionRulesApi.GETCouponIdPromotionRule(context.Background(), couponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETCouponIdPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponIdPromotionRuleRequest struct via the builder pattern


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


## GETExternalPromotionIdCouponCodesPromotionRule

> GETExternalPromotionIdCouponCodesPromotionRule(ctx, externalPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the external promotion



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
    r, err := apiClient.CouponCodesPromotionRulesApi.GETExternalPromotionIdCouponCodesPromotionRule(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETExternalPromotionIdCouponCodesPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## GETFixedAmountPromotionIdCouponCodesPromotionRule

> GETFixedAmountPromotionIdCouponCodesPromotionRule(ctx, fixedAmountPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the fixed amount promotion



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
    r, err := apiClient.CouponCodesPromotionRulesApi.GETFixedAmountPromotionIdCouponCodesPromotionRule(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFixedAmountPromotionIdCouponCodesPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## GETFixedPricePromotionIdCouponCodesPromotionRule

> GETFixedPricePromotionIdCouponCodesPromotionRule(ctx, fixedPricePromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the fixed price promotion



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
    r, err := apiClient.CouponCodesPromotionRulesApi.GETFixedPricePromotionIdCouponCodesPromotionRule(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFixedPricePromotionIdCouponCodesPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## GETFlexPromotionIdCouponCodesPromotionRule

> GETFlexPromotionIdCouponCodesPromotionRule(ctx, flexPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the flex promotion



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
    flexPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CouponCodesPromotionRulesApi.GETFlexPromotionIdCouponCodesPromotionRule(context.Background(), flexPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFlexPromotionIdCouponCodesPromotionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flexPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFlexPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## GETFreeGiftPromotionIdCouponCodesPromotionRule

> GETFreeGiftPromotionIdCouponCodesPromotionRule(ctx, freeGiftPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the free gift promotion



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
    r, err := apiClient.CouponCodesPromotionRulesApi.GETFreeGiftPromotionIdCouponCodesPromotionRule(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFreeGiftPromotionIdCouponCodesPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## GETFreeShippingPromotionIdCouponCodesPromotionRule

> GETFreeShippingPromotionIdCouponCodesPromotionRule(ctx, freeShippingPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the free shipping promotion



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
    r, err := apiClient.CouponCodesPromotionRulesApi.GETFreeShippingPromotionIdCouponCodesPromotionRule(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETFreeShippingPromotionIdCouponCodesPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## GETPercentageDiscountPromotionIdCouponCodesPromotionRule

> GETPercentageDiscountPromotionIdCouponCodesPromotionRule(ctx, percentageDiscountPromotionId).Execute()

Retrieve the coupon codes promotion rule associated to the percentage discount promotion



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
    r, err := apiClient.CouponCodesPromotionRulesApi.GETPercentageDiscountPromotionIdCouponCodesPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETPercentageDiscountPromotionIdCouponCodesPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## GETPromotionIdCouponCodesPromotionRule

> GETPromotionIdCouponCodesPromotionRule(ctx, promotionId).Execute()

Retrieve the coupon codes promotion rule associated to the promotion



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
    r, err := apiClient.CouponCodesPromotionRulesApi.GETPromotionIdCouponCodesPromotionRule(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.GETPromotionIdCouponCodesPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPromotionIdCouponCodesPromotionRuleRequest struct via the builder pattern


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


## PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId

> PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId(ctx, couponCodesPromotionRuleId).CouponCodesPromotionRuleUpdate(couponCodesPromotionRuleUpdate).Execute()

Update a coupon codes promotion rule



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
    couponCodesPromotionRuleUpdate := *openapiclient.NewCouponCodesPromotionRuleUpdate(*openapiclient.NewCouponCodesPromotionRuleUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes())) // CouponCodesPromotionRuleUpdate | 
    couponCodesPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId(context.Background(), couponCodesPromotionRuleId).CouponCodesPromotionRuleUpdate(couponCodesPromotionRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId`: PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponCodesPromotionRulesApi.PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCodesPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCouponCodesPromotionRulesCouponCodesPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponCodesPromotionRuleUpdate** | [**CouponCodesPromotionRuleUpdate**](CouponCodesPromotionRuleUpdate.md) |  | 


### Return type

[**PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response**](PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCouponCodesPromotionRules

> POSTCouponCodesPromotionRules201Response POSTCouponCodesPromotionRules(ctx).CouponCodesPromotionRuleCreate(couponCodesPromotionRuleCreate).Execute()

Create a coupon codes promotion rule



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
    couponCodesPromotionRuleCreate := *openapiclient.NewCouponCodesPromotionRuleCreate(*openapiclient.NewCouponCodesPromotionRuleCreateData(interface{}(123), *openapiclient.NewPOSTAdyenPayments201ResponseDataAttributes())) // CouponCodesPromotionRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponCodesPromotionRulesApi.POSTCouponCodesPromotionRules(context.Background()).CouponCodesPromotionRuleCreate(couponCodesPromotionRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponCodesPromotionRulesApi.POSTCouponCodesPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCouponCodesPromotionRules`: POSTCouponCodesPromotionRules201Response
    fmt.Fprintf(os.Stdout, "Response from `CouponCodesPromotionRulesApi.POSTCouponCodesPromotionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCouponCodesPromotionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponCodesPromotionRuleCreate** | [**CouponCodesPromotionRuleCreate**](CouponCodesPromotionRuleCreate.md) |  | 

### Return type

[**POSTCouponCodesPromotionRules201Response**](POSTCouponCodesPromotionRules201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

