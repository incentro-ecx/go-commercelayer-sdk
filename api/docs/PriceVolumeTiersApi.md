# \PriceVolumeTiersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPriceVolumeTiersPriceVolumeTierId**](PriceVolumeTiersApi.md#DELETEPriceVolumeTiersPriceVolumeTierId) | **Delete** /price_volume_tiers/{priceVolumeTierId} | Delete a price volume tier
[**GETPriceIdPriceVolumeTiers**](PriceVolumeTiersApi.md#GETPriceIdPriceVolumeTiers) | **Get** /prices/{priceId}/price_volume_tiers | Retrieve the price volume tiers associated to the price
[**GETPriceVolumeTiers**](PriceVolumeTiersApi.md#GETPriceVolumeTiers) | **Get** /price_volume_tiers | List all price volume tiers
[**GETPriceVolumeTiersPriceVolumeTierId**](PriceVolumeTiersApi.md#GETPriceVolumeTiersPriceVolumeTierId) | **Get** /price_volume_tiers/{priceVolumeTierId} | Retrieve a price volume tier
[**PATCHPriceVolumeTiersPriceVolumeTierId**](PriceVolumeTiersApi.md#PATCHPriceVolumeTiersPriceVolumeTierId) | **Patch** /price_volume_tiers/{priceVolumeTierId} | Update a price volume tier
[**POSTPriceVolumeTiers**](PriceVolumeTiersApi.md#POSTPriceVolumeTiers) | **Post** /price_volume_tiers | Create a price volume tier



## DELETEPriceVolumeTiersPriceVolumeTierId

> DELETEPriceVolumeTiersPriceVolumeTierId(ctx, priceVolumeTierId).Execute()

Delete a price volume tier



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
    r, err := apiClient.PriceVolumeTiersApi.DELETEPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.DELETEPriceVolumeTiersPriceVolumeTierId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEPriceVolumeTiersPriceVolumeTierIdRequest struct via the builder pattern


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


## GETPriceIdPriceVolumeTiers

> GETPriceIdPriceVolumeTiers(ctx, priceId).Execute()

Retrieve the price volume tiers associated to the price



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
    r, err := apiClient.PriceVolumeTiersApi.GETPriceIdPriceVolumeTiers(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.GETPriceIdPriceVolumeTiers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPriceIdPriceVolumeTiersRequest struct via the builder pattern


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


## GETPriceVolumeTiers

> GETPriceVolumeTiers200Response GETPriceVolumeTiers(ctx).Execute()

List all price volume tiers



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
    resp, r, err := apiClient.PriceVolumeTiersApi.GETPriceVolumeTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.GETPriceVolumeTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceVolumeTiers`: GETPriceVolumeTiers200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceVolumeTiersApi.GETPriceVolumeTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceVolumeTiersRequest struct via the builder pattern


### Return type

[**GETPriceVolumeTiers200Response**](GETPriceVolumeTiers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPriceVolumeTiersPriceVolumeTierId

> GETPriceVolumeTiersPriceVolumeTierId200Response GETPriceVolumeTiersPriceVolumeTierId(ctx, priceVolumeTierId).Execute()

Retrieve a price volume tier



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
    resp, r, err := apiClient.PriceVolumeTiersApi.GETPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.GETPriceVolumeTiersPriceVolumeTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceVolumeTiersPriceVolumeTierId`: GETPriceVolumeTiersPriceVolumeTierId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceVolumeTiersApi.GETPriceVolumeTiersPriceVolumeTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceVolumeTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceVolumeTiersPriceVolumeTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPriceVolumeTiersPriceVolumeTierId200Response**](GETPriceVolumeTiersPriceVolumeTierId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPriceVolumeTiersPriceVolumeTierId

> PATCHPriceVolumeTiersPriceVolumeTierId200Response PATCHPriceVolumeTiersPriceVolumeTierId(ctx, priceVolumeTierId).PriceVolumeTierUpdate(priceVolumeTierUpdate).Execute()

Update a price volume tier



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
    priceVolumeTierUpdate := *openapiclient.NewPriceVolumeTierUpdate(*openapiclient.NewPriceVolumeTierUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes())) // PriceVolumeTierUpdate | 
    priceVolumeTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceVolumeTiersApi.PATCHPriceVolumeTiersPriceVolumeTierId(context.Background(), priceVolumeTierId).PriceVolumeTierUpdate(priceVolumeTierUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.PATCHPriceVolumeTiersPriceVolumeTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPriceVolumeTiersPriceVolumeTierId`: PATCHPriceVolumeTiersPriceVolumeTierId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceVolumeTiersApi.PATCHPriceVolumeTiersPriceVolumeTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceVolumeTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPriceVolumeTiersPriceVolumeTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceVolumeTierUpdate** | [**PriceVolumeTierUpdate**](PriceVolumeTierUpdate.md) |  | 


### Return type

[**PATCHPriceVolumeTiersPriceVolumeTierId200Response**](PATCHPriceVolumeTiersPriceVolumeTierId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPriceVolumeTiers

> POSTPriceVolumeTiers201Response POSTPriceVolumeTiers(ctx).PriceVolumeTierCreate(priceVolumeTierCreate).Execute()

Create a price volume tier



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
    priceVolumeTierCreate := *openapiclient.NewPriceVolumeTierCreate(*openapiclient.NewPriceVolumeTierCreateData(interface{}(123), *openapiclient.NewPOSTPriceVolumeTiers201ResponseDataAttributes(interface{}(six pack), interface{}(1000)))) // PriceVolumeTierCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceVolumeTiersApi.POSTPriceVolumeTiers(context.Background()).PriceVolumeTierCreate(priceVolumeTierCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceVolumeTiersApi.POSTPriceVolumeTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPriceVolumeTiers`: POSTPriceVolumeTiers201Response
    fmt.Fprintf(os.Stdout, "Response from `PriceVolumeTiersApi.POSTPriceVolumeTiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPriceVolumeTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceVolumeTierCreate** | [**PriceVolumeTierCreate**](PriceVolumeTierCreate.md) |  | 

### Return type

[**POSTPriceVolumeTiers201Response**](POSTPriceVolumeTiers201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

