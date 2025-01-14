# \CustomPromotionRulesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomPromotionRulesCustomPromotionRuleId**](CustomPromotionRulesApi.md#DELETECustomPromotionRulesCustomPromotionRuleId) | **Delete** /custom_promotion_rules/{customPromotionRuleId} | Delete a custom promotion rule
[**GETBuyXPayYPromotionIdCustomPromotionRule**](CustomPromotionRulesApi.md#GETBuyXPayYPromotionIdCustomPromotionRule) | **Get** /buy_x_pay_y_promotions/{buyXPayYPromotionId}/custom_promotion_rule | Retrieve the custom promotion rule associated to the buy x pay y promotion
[**GETCustomPromotionRules**](CustomPromotionRulesApi.md#GETCustomPromotionRules) | **Get** /custom_promotion_rules | List all custom promotion rules
[**GETCustomPromotionRulesCustomPromotionRuleId**](CustomPromotionRulesApi.md#GETCustomPromotionRulesCustomPromotionRuleId) | **Get** /custom_promotion_rules/{customPromotionRuleId} | Retrieve a custom promotion rule
[**GETExternalPromotionIdCustomPromotionRule**](CustomPromotionRulesApi.md#GETExternalPromotionIdCustomPromotionRule) | **Get** /external_promotions/{externalPromotionId}/custom_promotion_rule | Retrieve the custom promotion rule associated to the external promotion
[**GETFixedAmountPromotionIdCustomPromotionRule**](CustomPromotionRulesApi.md#GETFixedAmountPromotionIdCustomPromotionRule) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/custom_promotion_rule | Retrieve the custom promotion rule associated to the fixed amount promotion
[**GETFixedPricePromotionIdCustomPromotionRule**](CustomPromotionRulesApi.md#GETFixedPricePromotionIdCustomPromotionRule) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/custom_promotion_rule | Retrieve the custom promotion rule associated to the fixed price promotion
[**GETFreeGiftPromotionIdCustomPromotionRule**](CustomPromotionRulesApi.md#GETFreeGiftPromotionIdCustomPromotionRule) | **Get** /free_gift_promotions/{freeGiftPromotionId}/custom_promotion_rule | Retrieve the custom promotion rule associated to the free gift promotion
[**GETFreeShippingPromotionIdCustomPromotionRule**](CustomPromotionRulesApi.md#GETFreeShippingPromotionIdCustomPromotionRule) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/custom_promotion_rule | Retrieve the custom promotion rule associated to the free shipping promotion
[**GETPercentageDiscountPromotionIdCustomPromotionRule**](CustomPromotionRulesApi.md#GETPercentageDiscountPromotionIdCustomPromotionRule) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/custom_promotion_rule | Retrieve the custom promotion rule associated to the percentage discount promotion
[**GETPromotionIdCustomPromotionRule**](CustomPromotionRulesApi.md#GETPromotionIdCustomPromotionRule) | **Get** /promotions/{promotionId}/custom_promotion_rule | Retrieve the custom promotion rule associated to the promotion
[**PATCHCustomPromotionRulesCustomPromotionRuleId**](CustomPromotionRulesApi.md#PATCHCustomPromotionRulesCustomPromotionRuleId) | **Patch** /custom_promotion_rules/{customPromotionRuleId} | Update a custom promotion rule
[**POSTCustomPromotionRules**](CustomPromotionRulesApi.md#POSTCustomPromotionRules) | **Post** /custom_promotion_rules | Create a custom promotion rule



## DELETECustomPromotionRulesCustomPromotionRuleId

> DELETECustomPromotionRulesCustomPromotionRuleId(ctx, customPromotionRuleId).Execute()

Delete a custom promotion rule



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
    customPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomPromotionRulesApi.DELETECustomPromotionRulesCustomPromotionRuleId(context.Background(), customPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.DELETECustomPromotionRulesCustomPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomPromotionRulesCustomPromotionRuleIdRequest struct via the builder pattern


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


## GETBuyXPayYPromotionIdCustomPromotionRule

> GETBuyXPayYPromotionIdCustomPromotionRule(ctx, buyXPayYPromotionId).Execute()

Retrieve the custom promotion rule associated to the buy x pay y promotion



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
    r, err := apiClient.CustomPromotionRulesApi.GETBuyXPayYPromotionIdCustomPromotionRule(context.Background(), buyXPayYPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETBuyXPayYPromotionIdCustomPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETBuyXPayYPromotionIdCustomPromotionRuleRequest struct via the builder pattern


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


## GETCustomPromotionRules

> GETCustomPromotionRules200Response GETCustomPromotionRules(ctx).Execute()

List all custom promotion rules



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
    resp, r, err := apiClient.CustomPromotionRulesApi.GETCustomPromotionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETCustomPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomPromotionRules`: GETCustomPromotionRules200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomPromotionRulesApi.GETCustomPromotionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomPromotionRulesRequest struct via the builder pattern


### Return type

[**GETCustomPromotionRules200Response**](GETCustomPromotionRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomPromotionRulesCustomPromotionRuleId

> GETCustomPromotionRulesCustomPromotionRuleId200Response GETCustomPromotionRulesCustomPromotionRuleId(ctx, customPromotionRuleId).Execute()

Retrieve a custom promotion rule



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
    customPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPromotionRulesApi.GETCustomPromotionRulesCustomPromotionRuleId(context.Background(), customPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETCustomPromotionRulesCustomPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomPromotionRulesCustomPromotionRuleId`: GETCustomPromotionRulesCustomPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomPromotionRulesApi.GETCustomPromotionRulesCustomPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomPromotionRulesCustomPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCustomPromotionRulesCustomPromotionRuleId200Response**](GETCustomPromotionRulesCustomPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalPromotionIdCustomPromotionRule

> GETExternalPromotionIdCustomPromotionRule(ctx, externalPromotionId).Execute()

Retrieve the custom promotion rule associated to the external promotion



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
    r, err := apiClient.CustomPromotionRulesApi.GETExternalPromotionIdCustomPromotionRule(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETExternalPromotionIdCustomPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPromotionIdCustomPromotionRuleRequest struct via the builder pattern


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


## GETFixedAmountPromotionIdCustomPromotionRule

> GETFixedAmountPromotionIdCustomPromotionRule(ctx, fixedAmountPromotionId).Execute()

Retrieve the custom promotion rule associated to the fixed amount promotion



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
    r, err := apiClient.CustomPromotionRulesApi.GETFixedAmountPromotionIdCustomPromotionRule(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETFixedAmountPromotionIdCustomPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdCustomPromotionRuleRequest struct via the builder pattern


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


## GETFixedPricePromotionIdCustomPromotionRule

> GETFixedPricePromotionIdCustomPromotionRule(ctx, fixedPricePromotionId).Execute()

Retrieve the custom promotion rule associated to the fixed price promotion



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
    r, err := apiClient.CustomPromotionRulesApi.GETFixedPricePromotionIdCustomPromotionRule(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETFixedPricePromotionIdCustomPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdCustomPromotionRuleRequest struct via the builder pattern


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


## GETFreeGiftPromotionIdCustomPromotionRule

> GETFreeGiftPromotionIdCustomPromotionRule(ctx, freeGiftPromotionId).Execute()

Retrieve the custom promotion rule associated to the free gift promotion



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
    r, err := apiClient.CustomPromotionRulesApi.GETFreeGiftPromotionIdCustomPromotionRule(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETFreeGiftPromotionIdCustomPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdCustomPromotionRuleRequest struct via the builder pattern


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


## GETFreeShippingPromotionIdCustomPromotionRule

> GETFreeShippingPromotionIdCustomPromotionRule(ctx, freeShippingPromotionId).Execute()

Retrieve the custom promotion rule associated to the free shipping promotion



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
    r, err := apiClient.CustomPromotionRulesApi.GETFreeShippingPromotionIdCustomPromotionRule(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETFreeShippingPromotionIdCustomPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdCustomPromotionRuleRequest struct via the builder pattern


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


## GETPercentageDiscountPromotionIdCustomPromotionRule

> GETPercentageDiscountPromotionIdCustomPromotionRule(ctx, percentageDiscountPromotionId).Execute()

Retrieve the custom promotion rule associated to the percentage discount promotion



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
    r, err := apiClient.CustomPromotionRulesApi.GETPercentageDiscountPromotionIdCustomPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETPercentageDiscountPromotionIdCustomPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdCustomPromotionRuleRequest struct via the builder pattern


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


## GETPromotionIdCustomPromotionRule

> GETPromotionIdCustomPromotionRule(ctx, promotionId).Execute()

Retrieve the custom promotion rule associated to the promotion



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
    r, err := apiClient.CustomPromotionRulesApi.GETPromotionIdCustomPromotionRule(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.GETPromotionIdCustomPromotionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPromotionIdCustomPromotionRuleRequest struct via the builder pattern


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


## PATCHCustomPromotionRulesCustomPromotionRuleId

> PATCHCustomPromotionRulesCustomPromotionRuleId200Response PATCHCustomPromotionRulesCustomPromotionRuleId(ctx, customPromotionRuleId).CustomPromotionRuleUpdate(customPromotionRuleUpdate).Execute()

Update a custom promotion rule



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
    customPromotionRuleUpdate := *openapiclient.NewCustomPromotionRuleUpdate(*openapiclient.NewCustomPromotionRuleUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes())) // CustomPromotionRuleUpdate | 
    customPromotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPromotionRulesApi.PATCHCustomPromotionRulesCustomPromotionRuleId(context.Background(), customPromotionRuleId).CustomPromotionRuleUpdate(customPromotionRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.PATCHCustomPromotionRulesCustomPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCustomPromotionRulesCustomPromotionRuleId`: PATCHCustomPromotionRulesCustomPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomPromotionRulesApi.PATCHCustomPromotionRulesCustomPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customPromotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomPromotionRulesCustomPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customPromotionRuleUpdate** | [**CustomPromotionRuleUpdate**](CustomPromotionRuleUpdate.md) |  | 


### Return type

[**PATCHCustomPromotionRulesCustomPromotionRuleId200Response**](PATCHCustomPromotionRulesCustomPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCustomPromotionRules

> POSTCustomPromotionRules201Response POSTCustomPromotionRules(ctx).CustomPromotionRuleCreate(customPromotionRuleCreate).Execute()

Create a custom promotion rule



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
    customPromotionRuleCreate := *openapiclient.NewCustomPromotionRuleCreate(*openapiclient.NewCustomPromotionRuleCreateData(interface{}(123), *openapiclient.NewPOSTCustomPromotionRules201ResponseDataAttributes())) // CustomPromotionRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPromotionRulesApi.POSTCustomPromotionRules(context.Background()).CustomPromotionRuleCreate(customPromotionRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPromotionRulesApi.POSTCustomPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCustomPromotionRules`: POSTCustomPromotionRules201Response
    fmt.Fprintf(os.Stdout, "Response from `CustomPromotionRulesApi.POSTCustomPromotionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomPromotionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customPromotionRuleCreate** | [**CustomPromotionRuleCreate**](CustomPromotionRuleCreate.md) |  | 

### Return type

[**POSTCustomPromotionRules201Response**](POSTCustomPromotionRules201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

