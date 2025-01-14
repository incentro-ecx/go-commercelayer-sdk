# \PricesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPricesPriceId**](PricesApi.md#DELETEPricesPriceId) | **Delete** /prices/{priceId} | Delete a price
[**GETPriceFrequencyTierIdPrice**](PricesApi.md#GETPriceFrequencyTierIdPrice) | **Get** /price_frequency_tiers/{priceFrequencyTierId}/price | Retrieve the price associated to the price frequency tier
[**GETPriceListIdPrices**](PricesApi.md#GETPriceListIdPrices) | **Get** /price_lists/{priceListId}/prices | Retrieve the prices associated to the price list
[**GETPriceTierIdPrice**](PricesApi.md#GETPriceTierIdPrice) | **Get** /price_tiers/{priceTierId}/price | Retrieve the price associated to the price tier
[**GETPriceVolumeTierIdPrice**](PricesApi.md#GETPriceVolumeTierIdPrice) | **Get** /price_volume_tiers/{priceVolumeTierId}/price | Retrieve the price associated to the price volume tier
[**GETPrices**](PricesApi.md#GETPrices) | **Get** /prices | List all prices
[**GETPricesPriceId**](PricesApi.md#GETPricesPriceId) | **Get** /prices/{priceId} | Retrieve a price
[**GETSkuIdPrices**](PricesApi.md#GETSkuIdPrices) | **Get** /skus/{skuId}/prices | Retrieve the prices associated to the SKU
[**PATCHPricesPriceId**](PricesApi.md#PATCHPricesPriceId) | **Patch** /prices/{priceId} | Update a price
[**POSTPrices**](PricesApi.md#POSTPrices) | **Post** /prices | Create a price



## DELETEPricesPriceId

> DELETEPricesPriceId(ctx, priceId).Execute()

Delete a price



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
    r, err := apiClient.PricesApi.DELETEPricesPriceId(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.DELETEPricesPriceId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEPricesPriceIdRequest struct via the builder pattern


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


## GETPriceFrequencyTierIdPrice

> GETPriceFrequencyTierIdPrice(ctx, priceFrequencyTierId).Execute()

Retrieve the price associated to the price frequency tier



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
    priceFrequencyTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PricesApi.GETPriceFrequencyTierIdPrice(context.Background(), priceFrequencyTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.GETPriceFrequencyTierIdPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceFrequencyTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceFrequencyTierIdPriceRequest struct via the builder pattern


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


## GETPriceListIdPrices

> GETPriceListIdPrices(ctx, priceListId).Execute()

Retrieve the prices associated to the price list



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
    priceListId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PricesApi.GETPriceListIdPrices(context.Background(), priceListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.GETPriceListIdPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceListIdPricesRequest struct via the builder pattern


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


## GETPriceTierIdPrice

> GETPriceTierIdPrice(ctx, priceTierId).Execute()

Retrieve the price associated to the price tier



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
    priceTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PricesApi.GETPriceTierIdPrice(context.Background(), priceTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.GETPriceTierIdPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceTierIdPriceRequest struct via the builder pattern


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


## GETPriceVolumeTierIdPrice

> GETPriceVolumeTierIdPrice(ctx, priceVolumeTierId).Execute()

Retrieve the price associated to the price volume tier



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
    priceVolumeTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PricesApi.GETPriceVolumeTierIdPrice(context.Background(), priceVolumeTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.GETPriceVolumeTierIdPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceVolumeTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceVolumeTierIdPriceRequest struct via the builder pattern


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


## GETPrices

> GETPrices200Response GETPrices(ctx).Execute()

List all prices



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
    resp, r, err := apiClient.PricesApi.GETPrices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.GETPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPrices`: GETPrices200Response
    fmt.Fprintf(os.Stdout, "Response from `PricesApi.GETPrices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPricesRequest struct via the builder pattern


### Return type

[**GETPrices200Response**](GETPrices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPricesPriceId

> GETPricesPriceId200Response GETPricesPriceId(ctx, priceId).Execute()

Retrieve a price



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
    resp, r, err := apiClient.PricesApi.GETPricesPriceId(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.GETPricesPriceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPricesPriceId`: GETPricesPriceId200Response
    fmt.Fprintf(os.Stdout, "Response from `PricesApi.GETPricesPriceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPricesPriceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPricesPriceId200Response**](GETPricesPriceId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSkuIdPrices

> GETSkuIdPrices(ctx, skuId).Execute()

Retrieve the prices associated to the SKU



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
    r, err := apiClient.PricesApi.GETSkuIdPrices(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.GETSkuIdPrices``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuIdPricesRequest struct via the builder pattern


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


## PATCHPricesPriceId

> PATCHPricesPriceId200Response PATCHPricesPriceId(ctx, priceId).PriceUpdate(priceUpdate).Execute()

Update a price



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
    priceUpdate := *openapiclient.NewPriceUpdate(*openapiclient.NewPriceUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPricesPriceId200ResponseDataAttributes())) // PriceUpdate | 
    priceId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesApi.PATCHPricesPriceId(context.Background(), priceId).PriceUpdate(priceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.PATCHPricesPriceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPricesPriceId`: PATCHPricesPriceId200Response
    fmt.Fprintf(os.Stdout, "Response from `PricesApi.PATCHPricesPriceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPricesPriceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceUpdate** | [**PriceUpdate**](PriceUpdate.md) |  | 


### Return type

[**PATCHPricesPriceId200Response**](PATCHPricesPriceId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPrices

> POSTPrices201Response POSTPrices(ctx).PriceCreate(priceCreate).Execute()

Create a price



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
    priceCreate := *openapiclient.NewPriceCreate(*openapiclient.NewPriceCreateData(interface{}(123), *openapiclient.NewPOSTPrices201ResponseDataAttributes(interface{}(10000)))) // PriceCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesApi.POSTPrices(context.Background()).PriceCreate(priceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.POSTPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPrices`: POSTPrices201Response
    fmt.Fprintf(os.Stdout, "Response from `PricesApi.POSTPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceCreate** | [**PriceCreate**](PriceCreate.md) |  | 

### Return type

[**POSTPrices201Response**](POSTPrices201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

