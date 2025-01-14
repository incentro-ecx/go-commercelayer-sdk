# \PriceTiersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETPriceIdPriceTiers**](PriceTiersApi.md#GETPriceIdPriceTiers) | **Get** /prices/{priceId}/price_tiers | Retrieve the price tiers associated to the price
[**GETPriceTiers**](PriceTiersApi.md#GETPriceTiers) | **Get** /price_tiers | List all price tiers
[**GETPriceTiersPriceTierId**](PriceTiersApi.md#GETPriceTiersPriceTierId) | **Get** /price_tiers/{priceTierId} | Retrieve a price tier



## GETPriceIdPriceTiers

> GETPriceIdPriceTiers(ctx, priceId).Execute()

Retrieve the price tiers associated to the price



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
    r, err := apiClient.PriceTiersApi.GETPriceIdPriceTiers(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceTiersApi.GETPriceIdPriceTiers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPriceIdPriceTiersRequest struct via the builder pattern


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


## GETPriceTiers

> GETPriceTiers200Response GETPriceTiers(ctx).Execute()

List all price tiers



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
    resp, r, err := apiClient.PriceTiersApi.GETPriceTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceTiersApi.GETPriceTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceTiers`: GETPriceTiers200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceTiersApi.GETPriceTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceTiersRequest struct via the builder pattern


### Return type

[**GETPriceTiers200Response**](GETPriceTiers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPriceTiersPriceTierId

> GETPriceTiersPriceTierId200Response GETPriceTiersPriceTierId(ctx, priceTierId).Execute()

Retrieve a price tier



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
    resp, r, err := apiClient.PriceTiersApi.GETPriceTiersPriceTierId(context.Background(), priceTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceTiersApi.GETPriceTiersPriceTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceTiersPriceTierId`: GETPriceTiersPriceTierId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceTiersApi.GETPriceTiersPriceTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceTiersPriceTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPriceTiersPriceTierId200Response**](GETPriceTiersPriceTierId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

