# \SkuListsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESkuListsSkuListId**](SkuListsApi.md#DELETESkuListsSkuListId) | **Delete** /sku_lists/{skuListId} | Delete a SKU list
[**GETBundleIdSkuList**](SkuListsApi.md#GETBundleIdSkuList) | **Get** /bundles/{bundleId}/sku_list | Retrieve the sku list associated to the bundle
[**GETBuyXPayYPromotionIdSkuList**](SkuListsApi.md#GETBuyXPayYPromotionIdSkuList) | **Get** /buy_x_pay_y_promotions/{buyXPayYPromotionId}/sku_list | Retrieve the sku list associated to the buy x pay y promotion
[**GETCustomerIdSkuLists**](SkuListsApi.md#GETCustomerIdSkuLists) | **Get** /customers/{customerId}/sku_lists | Retrieve the sku lists associated to the customer
[**GETExternalPromotionIdSkuList**](SkuListsApi.md#GETExternalPromotionIdSkuList) | **Get** /external_promotions/{externalPromotionId}/sku_list | Retrieve the sku list associated to the external promotion
[**GETFixedAmountPromotionIdSkuList**](SkuListsApi.md#GETFixedAmountPromotionIdSkuList) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId}/sku_list | Retrieve the sku list associated to the fixed amount promotion
[**GETFixedPricePromotionIdSkuList**](SkuListsApi.md#GETFixedPricePromotionIdSkuList) | **Get** /fixed_price_promotions/{fixedPricePromotionId}/sku_list | Retrieve the sku list associated to the fixed price promotion
[**GETFreeGiftPromotionIdSkuList**](SkuListsApi.md#GETFreeGiftPromotionIdSkuList) | **Get** /free_gift_promotions/{freeGiftPromotionId}/sku_list | Retrieve the sku list associated to the free gift promotion
[**GETFreeShippingPromotionIdSkuList**](SkuListsApi.md#GETFreeShippingPromotionIdSkuList) | **Get** /free_shipping_promotions/{freeShippingPromotionId}/sku_list | Retrieve the sku list associated to the free shipping promotion
[**GETPercentageDiscountPromotionIdSkuList**](SkuListsApi.md#GETPercentageDiscountPromotionIdSkuList) | **Get** /percentage_discount_promotions/{percentageDiscountPromotionId}/sku_list | Retrieve the sku list associated to the percentage discount promotion
[**GETPromotionIdSkuList**](SkuListsApi.md#GETPromotionIdSkuList) | **Get** /promotions/{promotionId}/sku_list | Retrieve the sku list associated to the promotion
[**GETSkuIdSkuLists**](SkuListsApi.md#GETSkuIdSkuLists) | **Get** /skus/{skuId}/sku_lists | Retrieve the sku lists associated to the SKU
[**GETSkuListItemIdSkuList**](SkuListsApi.md#GETSkuListItemIdSkuList) | **Get** /sku_list_items/{skuListItemId}/sku_list | Retrieve the sku list associated to the SKU list item
[**GETSkuListPromotionRuleIdSkuList**](SkuListsApi.md#GETSkuListPromotionRuleIdSkuList) | **Get** /sku_list_promotion_rules/{skuListPromotionRuleId}/sku_list | Retrieve the sku list associated to the SKU list promotion rule
[**GETSkuLists**](SkuListsApi.md#GETSkuLists) | **Get** /sku_lists | List all SKU lists
[**GETSkuListsSkuListId**](SkuListsApi.md#GETSkuListsSkuListId) | **Get** /sku_lists/{skuListId} | Retrieve a SKU list
[**PATCHSkuListsSkuListId**](SkuListsApi.md#PATCHSkuListsSkuListId) | **Patch** /sku_lists/{skuListId} | Update a SKU list
[**POSTSkuLists**](SkuListsApi.md#POSTSkuLists) | **Post** /sku_lists | Create a SKU list



## DELETESkuListsSkuListId

> DELETESkuListsSkuListId(ctx, skuListId).Execute()

Delete a SKU list



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
    r, err := apiClient.SkuListsApi.DELETESkuListsSkuListId(context.Background(), skuListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.DELETESkuListsSkuListId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETESkuListsSkuListIdRequest struct via the builder pattern


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


## GETBundleIdSkuList

> GETBundleIdSkuList(ctx, bundleId).Execute()

Retrieve the sku list associated to the bundle



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
    bundleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SkuListsApi.GETBundleIdSkuList(context.Background(), bundleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETBundleIdSkuList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBundleIdSkuListRequest struct via the builder pattern


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


## GETBuyXPayYPromotionIdSkuList

> GETBuyXPayYPromotionIdSkuList(ctx, buyXPayYPromotionId).Execute()

Retrieve the sku list associated to the buy x pay y promotion



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
    r, err := apiClient.SkuListsApi.GETBuyXPayYPromotionIdSkuList(context.Background(), buyXPayYPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETBuyXPayYPromotionIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETBuyXPayYPromotionIdSkuListRequest struct via the builder pattern


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


## GETCustomerIdSkuLists

> GETCustomerIdSkuLists(ctx, customerId).Execute()

Retrieve the sku lists associated to the customer



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
    r, err := apiClient.SkuListsApi.GETCustomerIdSkuLists(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETCustomerIdSkuLists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETCustomerIdSkuListsRequest struct via the builder pattern


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


## GETExternalPromotionIdSkuList

> GETExternalPromotionIdSkuList(ctx, externalPromotionId).Execute()

Retrieve the sku list associated to the external promotion



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
    r, err := apiClient.SkuListsApi.GETExternalPromotionIdSkuList(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETExternalPromotionIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalPromotionIdSkuListRequest struct via the builder pattern


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


## GETFixedAmountPromotionIdSkuList

> GETFixedAmountPromotionIdSkuList(ctx, fixedAmountPromotionId).Execute()

Retrieve the sku list associated to the fixed amount promotion



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
    r, err := apiClient.SkuListsApi.GETFixedAmountPromotionIdSkuList(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETFixedAmountPromotionIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionIdSkuListRequest struct via the builder pattern


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


## GETFixedPricePromotionIdSkuList

> GETFixedPricePromotionIdSkuList(ctx, fixedPricePromotionId).Execute()

Retrieve the sku list associated to the fixed price promotion



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
    r, err := apiClient.SkuListsApi.GETFixedPricePromotionIdSkuList(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETFixedPricePromotionIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFixedPricePromotionIdSkuListRequest struct via the builder pattern


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


## GETFreeGiftPromotionIdSkuList

> GETFreeGiftPromotionIdSkuList(ctx, freeGiftPromotionId).Execute()

Retrieve the sku list associated to the free gift promotion



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
    r, err := apiClient.SkuListsApi.GETFreeGiftPromotionIdSkuList(context.Background(), freeGiftPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETFreeGiftPromotionIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeGiftPromotionIdSkuListRequest struct via the builder pattern


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


## GETFreeShippingPromotionIdSkuList

> GETFreeShippingPromotionIdSkuList(ctx, freeShippingPromotionId).Execute()

Retrieve the sku list associated to the free shipping promotion



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
    r, err := apiClient.SkuListsApi.GETFreeShippingPromotionIdSkuList(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETFreeShippingPromotionIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionIdSkuListRequest struct via the builder pattern


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


## GETPercentageDiscountPromotionIdSkuList

> GETPercentageDiscountPromotionIdSkuList(ctx, percentageDiscountPromotionId).Execute()

Retrieve the sku list associated to the percentage discount promotion



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
    r, err := apiClient.SkuListsApi.GETPercentageDiscountPromotionIdSkuList(context.Background(), percentageDiscountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETPercentageDiscountPromotionIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPercentageDiscountPromotionIdSkuListRequest struct via the builder pattern


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


## GETPromotionIdSkuList

> GETPromotionIdSkuList(ctx, promotionId).Execute()

Retrieve the sku list associated to the promotion



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
    r, err := apiClient.SkuListsApi.GETPromotionIdSkuList(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETPromotionIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPromotionIdSkuListRequest struct via the builder pattern


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


## GETSkuIdSkuLists

> GETSkuIdSkuLists(ctx, skuId).Execute()

Retrieve the sku lists associated to the SKU



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
    r, err := apiClient.SkuListsApi.GETSkuIdSkuLists(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETSkuIdSkuLists``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuIdSkuListsRequest struct via the builder pattern


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


## GETSkuListItemIdSkuList

> GETSkuListItemIdSkuList(ctx, skuListItemId).Execute()

Retrieve the sku list associated to the SKU list item



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
    skuListItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SkuListsApi.GETSkuListItemIdSkuList(context.Background(), skuListItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETSkuListItemIdSkuList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListItemIdSkuListRequest struct via the builder pattern


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


## GETSkuListPromotionRuleIdSkuList

> GETSkuListPromotionRuleIdSkuList(ctx, skuListPromotionRuleId).Execute()

Retrieve the sku list associated to the SKU list promotion rule



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
    r, err := apiClient.SkuListsApi.GETSkuListPromotionRuleIdSkuList(context.Background(), skuListPromotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETSkuListPromotionRuleIdSkuList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuListPromotionRuleIdSkuListRequest struct via the builder pattern


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


## GETSkuLists

> GETSkuLists200Response GETSkuLists(ctx).Execute()

List all SKU lists



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
    resp, r, err := apiClient.SkuListsApi.GETSkuLists(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETSkuLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuLists`: GETSkuLists200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListsApi.GETSkuLists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListsRequest struct via the builder pattern


### Return type

[**GETSkuLists200Response**](GETSkuLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSkuListsSkuListId

> GETSkuListsSkuListId200Response GETSkuListsSkuListId(ctx, skuListId).Execute()

Retrieve a SKU list



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
    resp, r, err := apiClient.SkuListsApi.GETSkuListsSkuListId(context.Background(), skuListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.GETSkuListsSkuListId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuListsSkuListId`: GETSkuListsSkuListId200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListsApi.GETSkuListsSkuListId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListsSkuListIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETSkuListsSkuListId200Response**](GETSkuListsSkuListId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHSkuListsSkuListId

> PATCHSkuListsSkuListId200Response PATCHSkuListsSkuListId(ctx, skuListId).SkuListUpdate(skuListUpdate).Execute()

Update a SKU list



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
    skuListUpdate := *openapiclient.NewSkuListUpdate(*openapiclient.NewSkuListUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHSkuListsSkuListId200ResponseDataAttributes())) // SkuListUpdate | 
    skuListId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListsApi.PATCHSkuListsSkuListId(context.Background(), skuListId).SkuListUpdate(skuListUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.PATCHSkuListsSkuListId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHSkuListsSkuListId`: PATCHSkuListsSkuListId200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListsApi.PATCHSkuListsSkuListId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHSkuListsSkuListIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuListUpdate** | [**SkuListUpdate**](SkuListUpdate.md) |  | 


### Return type

[**PATCHSkuListsSkuListId200Response**](PATCHSkuListsSkuListId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTSkuLists

> POSTSkuLists201Response POSTSkuLists(ctx).SkuListCreate(skuListCreate).Execute()

Create a SKU list



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
    skuListCreate := *openapiclient.NewSkuListCreate(*openapiclient.NewSkuListCreateData(interface{}(123), *openapiclient.NewPOSTSkuLists201ResponseDataAttributes(interface{}(Personal list)))) // SkuListCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListsApi.POSTSkuLists(context.Background()).SkuListCreate(skuListCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListsApi.POSTSkuLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTSkuLists`: POSTSkuLists201Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListsApi.POSTSkuLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTSkuListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuListCreate** | [**SkuListCreate**](SkuListCreate.md) |  | 

### Return type

[**POSTSkuLists201Response**](POSTSkuLists201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

