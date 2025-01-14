# \SkuListPromotionRulesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESkuListPromotionRulesSkuListPromotionRuleId**](SkuListPromotionRulesApi.md#DELETESkuListPromotionRulesSkuListPromotionRuleId) | **Delete** /sku_list_promotion_rules/{skuListPromotionRuleId} | Delete a SKU list promotion rule
[**GETBuyXPayYPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETBuyXPayYPromotionIdSkuListPromotionRule) | **Get** /buy_x_pay_y_promotions/{buyXPayYPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the buy x pay y promotion
[**GETExternalPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETExternalPromotionIdSkuListPromotionRule) | **Get** /external_promotions/{externalPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the external promotion
[**GETFixedAmountPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETFixedAmountPromotionIdSkuListPromotionRule) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the fixed amount promotion
[**GETFixedPricePromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETFixedPricePromotionIdSkuListPromotionRule) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the fixed price promotion
[**GETFreeGiftPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETFreeGiftPromotionIdSkuListPromotionRule) | **Get** /free_gift_promotions/{freeGiftPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the free gift promotion
[**GETFreeShippingPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETFreeShippingPromotionIdSkuListPromotionRule) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the free shipping promotion
[**GETPercentageDiscountPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETPercentageDiscountPromotionIdSkuListPromotionRule) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the percentage discount promotion
[**GETPromotionIdSkuListPromotionRule**](SkuListPromotionRulesApi.md#GETPromotionIdSkuListPromotionRule) | **Get** /promotions/{promotionId}/sku_list_promotion_rule | Retrieve the sku list promotion rule associated to the promotion
[**GETSkuListPromotionRules**](SkuListPromotionRulesApi.md#GETSkuListPromotionRules) | **Get** /sku_list_promotion_rules | List all SKU list promotion rules
[**GETSkuListPromotionRulesSkuListPromotionRuleId**](SkuListPromotionRulesApi.md#GETSkuListPromotionRulesSkuListPromotionRuleId) | **Get** /sku_list_promotion_rules/{skuListPromotionRuleId} | Retrieve a SKU list promotion rule
[**PATCHSkuListPromotionRulesSkuListPromotionRuleId**](SkuListPromotionRulesApi.md#PATCHSkuListPromotionRulesSkuListPromotionRuleId) | **Patch** /sku_list_promotion_rules/{skuListPromotionRuleId} | Update a SKU list promotion rule
[**POSTSkuListPromotionRules**](SkuListPromotionRulesApi.md#POSTSkuListPromotionRules) | **Post** /sku_list_promotion_rules | Create a SKU list promotion rule



## DELETESkuListPromotionRulesSkuListPromotionRuleId

> DELETESkuListPromotionRulesSkuListPromotionRuleId(ctx, skuListPromotionRuleId).Execute()

Delete a SKU list promotion rule



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
    skuListPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SkuListPromotionRulesApi.DELETESkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.DELETESkuListPromotionRulesSkuListPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETESkuListPromotionRulesSkuListPromotionRuleIdRequest struct via the builder pattern


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


## GETBuyXPayYPromotionIdSkuListPromotionRule

> GETBuyXPayYPromotionIdSkuListPromotionRule(ctx, buyXPayYPromotionId).Execute()

Retrieve the sku list promotion rule associated to the buy x pay y promotion



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
    r, err := apiClient.SkuListPromotionRulesApi.GETBuyXPayYPromotionIdSkuListPromotionRule(context.Background(), buyXPayYPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETBuyXPayYPromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETBuyXPayYPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETExternalPromotionIdSkuListPromotionRule

> GETExternalPromotionIdSkuListPromotionRule(ctx, externalPromotionId).Execute()

Retrieve the sku list promotion rule associated to the external promotion



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
    r, err := apiClient.SkuListPromotionRulesApi.GETExternalPromotionIdSkuListPromotionRule(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETExternalPromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETFixedAmountPromotionIdSkuListPromotionRule

> GETFixedAmountPromotionIdSkuListPromotionRule(ctx, fixedAmountPromotionId).Execute()

Retrieve the sku list promotion rule associated to the fixed amount promotion



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
    r, err := apiClient.SkuListPromotionRulesApi.GETFixedAmountPromotionIdSkuListPromotionRule(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETFixedAmountPromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETFixedPricePromotionIdSkuListPromotionRule

> GETFixedPricePromotionIdSkuListPromotionRule(ctx, fixedPricePromotionId).Execute()

Retrieve the sku list promotion rule associated to the fixed price promotion



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
    r, err := apiClient.SkuListPromotionRulesApi.GETFixedPricePromotionIdSkuListPromotionRule(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETFixedPricePromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETFreeGiftPromotionIdSkuListPromotionRule

> GETFreeGiftPromotionIdSkuListPromotionRule(ctx, freeGiftPromotionId).Execute()

Retrieve the sku list promotion rule associated to the free gift promotion



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
    r, err := apiClient.SkuListPromotionRulesApi.GETFreeGiftPromotionIdSkuListPromotionRule(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETFreeGiftPromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETFreeShippingPromotionIdSkuListPromotionRule

> GETFreeShippingPromotionIdSkuListPromotionRule(ctx, freeShippingPromotionId).Execute()

Retrieve the sku list promotion rule associated to the free shipping promotion



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
    r, err := apiClient.SkuListPromotionRulesApi.GETFreeShippingPromotionIdSkuListPromotionRule(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETFreeShippingPromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETPercentageDiscountPromotionIdSkuListPromotionRule

> GETPercentageDiscountPromotionIdSkuListPromotionRule(ctx, percentageDiscountPromotionId).Execute()

Retrieve the sku list promotion rule associated to the percentage discount promotion



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
    r, err := apiClient.SkuListPromotionRulesApi.GETPercentageDiscountPromotionIdSkuListPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETPercentageDiscountPromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETPromotionIdSkuListPromotionRule

> GETPromotionIdSkuListPromotionRule(ctx, promotionId).Execute()

Retrieve the sku list promotion rule associated to the promotion



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
    r, err := apiClient.SkuListPromotionRulesApi.GETPromotionIdSkuListPromotionRule(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETPromotionIdSkuListPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPromotionIdSkuListPromotionRuleRequest struct via the builder pattern


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


## GETSkuListPromotionRules

> GETSkuListPromotionRules200Response GETSkuListPromotionRules(ctx).Execute()

List all SKU list promotion rules



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
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETSkuListPromotionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETSkuListPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuListPromotionRules`: GETSkuListPromotionRules200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListPromotionRulesApi.GETSkuListPromotionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListPromotionRulesRequest struct via the builder pattern


### Return type

[**GETSkuListPromotionRules200Response**](GETSkuListPromotionRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSkuListPromotionRulesSkuListPromotionRuleId

> GETSkuListPromotionRulesSkuListPromotionRuleId200Response GETSkuListPromotionRulesSkuListPromotionRuleId(ctx, skuListPromotionRuleId).Execute()

Retrieve a SKU list promotion rule



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
    skuListPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.GETSkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.GETSkuListPromotionRulesSkuListPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuListPromotionRulesSkuListPromotionRuleId`: GETSkuListPromotionRulesSkuListPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListPromotionRulesApi.GETSkuListPromotionRulesSkuListPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListPromotionRulesSkuListPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETSkuListPromotionRulesSkuListPromotionRuleId200Response**](GETSkuListPromotionRulesSkuListPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHSkuListPromotionRulesSkuListPromotionRuleId

> PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response PATCHSkuListPromotionRulesSkuListPromotionRuleId(ctx, skuListPromotionRuleId).SkuListPromotionRuleUpdate(skuListPromotionRuleUpdate).Execute()

Update a SKU list promotion rule



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
    skuListPromotionRuleUpdate := *openapiclient.NewSkuListPromotionRuleUpdate(*openapiclient.NewSkuListPromotionRuleUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes())) // SkuListPromotionRuleUpdate | 
    skuListPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.PATCHSkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).SkuListPromotionRuleUpdate(skuListPromotionRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.PATCHSkuListPromotionRulesSkuListPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHSkuListPromotionRulesSkuListPromotionRuleId`: PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListPromotionRulesApi.PATCHSkuListPromotionRulesSkuListPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHSkuListPromotionRulesSkuListPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuListPromotionRuleUpdate** | [**SkuListPromotionRuleUpdate**](SkuListPromotionRuleUpdate.md) |  | 


### Return type

[**PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response**](PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTSkuListPromotionRules

> POSTSkuListPromotionRules201Response POSTSkuListPromotionRules(ctx).SkuListPromotionRuleCreate(skuListPromotionRuleCreate).Execute()

Create a SKU list promotion rule



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
    skuListPromotionRuleCreate := *openapiclient.NewSkuListPromotionRuleCreate(*openapiclient.NewSkuListPromotionRuleCreateData(interface{}(123), *openapiclient.NewPOSTSkuListPromotionRules201ResponseDataAttributes())) // SkuListPromotionRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListPromotionRulesApi.POSTSkuListPromotionRules(context.Background()).SkuListPromotionRuleCreate(skuListPromotionRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListPromotionRulesApi.POSTSkuListPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTSkuListPromotionRules`: POSTSkuListPromotionRules201Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListPromotionRulesApi.POSTSkuListPromotionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTSkuListPromotionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuListPromotionRuleCreate** | [**SkuListPromotionRuleCreate**](SkuListPromotionRuleCreate.md) |  | 

### Return type

[**POSTSkuListPromotionRules201Response**](POSTSkuListPromotionRules201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

