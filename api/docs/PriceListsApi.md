# \PriceListsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPriceListsPriceListId**](PriceListsApi.md#DELETEPriceListsPriceListId) | **Delete** /price_lists/{priceListId} | Delete a price list
[**GETMarketIdBasePriceList**](PriceListsApi.md#GETMarketIdBasePriceList) | **Get** /markets/{marketId}/base_price_list | Retrieve the base price list associated to the market
[**GETMarketIdPriceList**](PriceListsApi.md#GETMarketIdPriceList) | **Get** /markets/{marketId}/price_list | Retrieve the price list associated to the market
[**GETPriceIdPriceList**](PriceListsApi.md#GETPriceIdPriceList) | **Get** /prices/{priceId}/price_list | Retrieve the price list associated to the price
[**GETPriceListSchedulerIdPriceList**](PriceListsApi.md#GETPriceListSchedulerIdPriceList) | **Get** /price_list_schedulers/{priceListSchedulerId}/price_list | Retrieve the price list associated to the price list scheduler
[**GETPriceLists**](PriceListsApi.md#GETPriceLists) | **Get** /price_lists | List all price lists
[**GETPriceListsPriceListId**](PriceListsApi.md#GETPriceListsPriceListId) | **Get** /price_lists/{priceListId} | Retrieve a price list
[**PATCHPriceListsPriceListId**](PriceListsApi.md#PATCHPriceListsPriceListId) | **Patch** /price_lists/{priceListId} | Update a price list
[**POSTPriceLists**](PriceListsApi.md#POSTPriceLists) | **Post** /price_lists | Create a price list



## DELETEPriceListsPriceListId

> DELETEPriceListsPriceListId(ctx, priceListId).Execute()

Delete a price list



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
    r, err := apiClient.PriceListsApi.DELETEPriceListsPriceListId(context.Background(), priceListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.DELETEPriceListsPriceListId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEPriceListsPriceListIdRequest struct via the builder pattern


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


## GETMarketIdBasePriceList

> GETMarketIdBasePriceList(ctx, marketId).Execute()

Retrieve the base price list associated to the market



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
    marketId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PriceListsApi.GETMarketIdBasePriceList(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETMarketIdBasePriceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdBasePriceListRequest struct via the builder pattern


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


## GETMarketIdPriceList

> GETMarketIdPriceList(ctx, marketId).Execute()

Retrieve the price list associated to the market



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
    marketId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PriceListsApi.GETMarketIdPriceList(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETMarketIdPriceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdPriceListRequest struct via the builder pattern


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


## GETPriceIdPriceList

> GETPriceIdPriceList(ctx, priceId).Execute()

Retrieve the price list associated to the price



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
    r, err := apiClient.PriceListsApi.GETPriceIdPriceList(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETPriceIdPriceList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPriceIdPriceListRequest struct via the builder pattern


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


## GETPriceListSchedulerIdPriceList

> GETPriceListSchedulerIdPriceList(ctx, priceListSchedulerId).Execute()

Retrieve the price list associated to the price list scheduler



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
    priceListSchedulerId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PriceListsApi.GETPriceListSchedulerIdPriceList(context.Background(), priceListSchedulerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETPriceListSchedulerIdPriceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListSchedulerId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceListSchedulerIdPriceListRequest struct via the builder pattern


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


## GETPriceLists

> GETPriceLists200Response GETPriceLists(ctx).Execute()

List all price lists



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
    resp, r, err := apiClient.PriceListsApi.GETPriceLists(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETPriceLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceLists`: GETPriceLists200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceListsApi.GETPriceLists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceListsRequest struct via the builder pattern


### Return type

[**GETPriceLists200Response**](GETPriceLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPriceListsPriceListId

> GETPriceListsPriceListId200Response GETPriceListsPriceListId(ctx, priceListId).Execute()

Retrieve a price list



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
    resp, r, err := apiClient.PriceListsApi.GETPriceListsPriceListId(context.Background(), priceListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.GETPriceListsPriceListId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceListsPriceListId`: GETPriceListsPriceListId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceListsApi.GETPriceListsPriceListId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceListsPriceListIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPriceListsPriceListId200Response**](GETPriceListsPriceListId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPriceListsPriceListId

> PATCHPriceListsPriceListId200Response PATCHPriceListsPriceListId(ctx, priceListId).PriceListUpdate(priceListUpdate).Execute()

Update a price list



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
    priceListUpdate := *openapiclient.NewPriceListUpdate(*openapiclient.NewPriceListUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPriceListsPriceListId200ResponseDataAttributes())) // PriceListUpdate | 
    priceListId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListsApi.PATCHPriceListsPriceListId(context.Background(), priceListId).PriceListUpdate(priceListUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.PATCHPriceListsPriceListId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPriceListsPriceListId`: PATCHPriceListsPriceListId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceListsApi.PATCHPriceListsPriceListId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPriceListsPriceListIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceListUpdate** | [**PriceListUpdate**](PriceListUpdate.md) |  | 


### Return type

[**PATCHPriceListsPriceListId200Response**](PATCHPriceListsPriceListId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPriceLists

> POSTPriceLists201Response POSTPriceLists(ctx).PriceListCreate(priceListCreate).Execute()

Create a price list



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
    priceListCreate := *openapiclient.NewPriceListCreate(*openapiclient.NewPriceListCreateData(interface{}(123), *openapiclient.NewPOSTPriceLists201ResponseDataAttributes(interface{}(EU Price list), interface{}(EUR)))) // PriceListCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListsApi.POSTPriceLists(context.Background()).PriceListCreate(priceListCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListsApi.POSTPriceLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPriceLists`: POSTPriceLists201Response
    fmt.Fprintf(os.Stdout, "Response from `PriceListsApi.POSTPriceLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPriceListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceListCreate** | [**PriceListCreate**](PriceListCreate.md) |  | 

### Return type

[**POSTPriceLists201Response**](POSTPriceLists201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

