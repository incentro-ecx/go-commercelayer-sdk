# \PriceListSchedulersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPriceListSchedulersPriceListSchedulerId**](PriceListSchedulersApi.md#DELETEPriceListSchedulersPriceListSchedulerId) | **Delete** /price_list_schedulers/{priceListSchedulerId} | Delete a price list scheduler
[**GETMarketIdPriceListSchedulers**](PriceListSchedulersApi.md#GETMarketIdPriceListSchedulers) | **Get** /markets/{marketId}/price_list_schedulers | Retrieve the price list schedulers associated to the market
[**GETPriceListIdPriceListSchedulers**](PriceListSchedulersApi.md#GETPriceListIdPriceListSchedulers) | **Get** /price_lists/{priceListId}/price_list_schedulers | Retrieve the price list schedulers associated to the price list
[**GETPriceListSchedulers**](PriceListSchedulersApi.md#GETPriceListSchedulers) | **Get** /price_list_schedulers | List all price list schedulers
[**GETPriceListSchedulersPriceListSchedulerId**](PriceListSchedulersApi.md#GETPriceListSchedulersPriceListSchedulerId) | **Get** /price_list_schedulers/{priceListSchedulerId} | Retrieve a price list scheduler
[**PATCHPriceListSchedulersPriceListSchedulerId**](PriceListSchedulersApi.md#PATCHPriceListSchedulersPriceListSchedulerId) | **Patch** /price_list_schedulers/{priceListSchedulerId} | Update a price list scheduler
[**POSTPriceListSchedulers**](PriceListSchedulersApi.md#POSTPriceListSchedulers) | **Post** /price_list_schedulers | Create a price list scheduler



## DELETEPriceListSchedulersPriceListSchedulerId

> DELETEPriceListSchedulersPriceListSchedulerId(ctx, priceListSchedulerId).Execute()

Delete a price list scheduler



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
    r, err := apiClient.PriceListSchedulersApi.DELETEPriceListSchedulersPriceListSchedulerId(context.Background(), priceListSchedulerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListSchedulersApi.DELETEPriceListSchedulersPriceListSchedulerId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEPriceListSchedulersPriceListSchedulerIdRequest struct via the builder pattern


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


## GETMarketIdPriceListSchedulers

> GETMarketIdPriceListSchedulers(ctx, marketId).Execute()

Retrieve the price list schedulers associated to the market



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
    r, err := apiClient.PriceListSchedulersApi.GETMarketIdPriceListSchedulers(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListSchedulersApi.GETMarketIdPriceListSchedulers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETMarketIdPriceListSchedulersRequest struct via the builder pattern


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


## GETPriceListIdPriceListSchedulers

> GETPriceListIdPriceListSchedulers(ctx, priceListId).Execute()

Retrieve the price list schedulers associated to the price list



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
    r, err := apiClient.PriceListSchedulersApi.GETPriceListIdPriceListSchedulers(context.Background(), priceListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListSchedulersApi.GETPriceListIdPriceListSchedulers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETPriceListIdPriceListSchedulersRequest struct via the builder pattern


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


## GETPriceListSchedulers

> GETPriceListSchedulers200Response GETPriceListSchedulers(ctx).Execute()

List all price list schedulers



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
    resp, r, err := apiClient.PriceListSchedulersApi.GETPriceListSchedulers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListSchedulersApi.GETPriceListSchedulers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceListSchedulers`: GETPriceListSchedulers200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceListSchedulersApi.GETPriceListSchedulers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceListSchedulersRequest struct via the builder pattern


### Return type

[**GETPriceListSchedulers200Response**](GETPriceListSchedulers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPriceListSchedulersPriceListSchedulerId

> GETPriceListSchedulersPriceListSchedulerId200Response GETPriceListSchedulersPriceListSchedulerId(ctx, priceListSchedulerId).Execute()

Retrieve a price list scheduler



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
    resp, r, err := apiClient.PriceListSchedulersApi.GETPriceListSchedulersPriceListSchedulerId(context.Background(), priceListSchedulerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListSchedulersApi.GETPriceListSchedulersPriceListSchedulerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceListSchedulersPriceListSchedulerId`: GETPriceListSchedulersPriceListSchedulerId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceListSchedulersApi.GETPriceListSchedulersPriceListSchedulerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListSchedulerId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceListSchedulersPriceListSchedulerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPriceListSchedulersPriceListSchedulerId200Response**](GETPriceListSchedulersPriceListSchedulerId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPriceListSchedulersPriceListSchedulerId

> PATCHPriceListSchedulersPriceListSchedulerId200Response PATCHPriceListSchedulersPriceListSchedulerId(ctx, priceListSchedulerId).PriceListSchedulerUpdate(priceListSchedulerUpdate).Execute()

Update a price list scheduler



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
    priceListSchedulerUpdate := *openapiclient.NewPriceListSchedulerUpdate(*openapiclient.NewPriceListSchedulerUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes())) // PriceListSchedulerUpdate | 
    priceListSchedulerId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListSchedulersApi.PATCHPriceListSchedulersPriceListSchedulerId(context.Background(), priceListSchedulerId).PriceListSchedulerUpdate(priceListSchedulerUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListSchedulersApi.PATCHPriceListSchedulersPriceListSchedulerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPriceListSchedulersPriceListSchedulerId`: PATCHPriceListSchedulersPriceListSchedulerId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceListSchedulersApi.PATCHPriceListSchedulersPriceListSchedulerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListSchedulerId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPriceListSchedulersPriceListSchedulerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceListSchedulerUpdate** | [**PriceListSchedulerUpdate**](PriceListSchedulerUpdate.md) |  | 


### Return type

[**PATCHPriceListSchedulersPriceListSchedulerId200Response**](PATCHPriceListSchedulersPriceListSchedulerId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPriceListSchedulers

> POSTPriceListSchedulers201Response POSTPriceListSchedulers(ctx).PriceListSchedulerCreate(priceListSchedulerCreate).Execute()

Create a price list scheduler



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
    priceListSchedulerCreate := *openapiclient.NewPriceListSchedulerCreate(*openapiclient.NewPriceListSchedulerCreateData(interface{}(123), *openapiclient.NewPOSTPriceListSchedulers201ResponseDataAttributes(interface{}(FW SALE 2023), interface{}(2018-01-01T12:00:00.000Z), interface{}(2018-01-02T12:00:00.000Z)))) // PriceListSchedulerCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceListSchedulersApi.POSTPriceListSchedulers(context.Background()).PriceListSchedulerCreate(priceListSchedulerCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceListSchedulersApi.POSTPriceListSchedulers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPriceListSchedulers`: POSTPriceListSchedulers201Response
    fmt.Fprintf(os.Stdout, "Response from `PriceListSchedulersApi.POSTPriceListSchedulers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPriceListSchedulersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceListSchedulerCreate** | [**PriceListSchedulerCreate**](PriceListSchedulerCreate.md) |  | 

### Return type

[**POSTPriceListSchedulers201Response**](POSTPriceListSchedulers201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

