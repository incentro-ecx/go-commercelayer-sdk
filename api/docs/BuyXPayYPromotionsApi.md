# \BuyXPayYPromotionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBuyXPayYPromotionsBuyXPayYPromotionId**](BuyXPayYPromotionsApi.md#DELETEBuyXPayYPromotionsBuyXPayYPromotionId) | **Delete** /buy_x_pay_y_promotions/{buyXPayYPromotionId} | Delete a buy x pay y promotion
[**GETBuyXPayYPromotions**](BuyXPayYPromotionsApi.md#GETBuyXPayYPromotions) | **Get** /buy_x_pay_y_promotions | List all buy x pay y promotions
[**GETBuyXPayYPromotionsBuyXPayYPromotionId**](BuyXPayYPromotionsApi.md#GETBuyXPayYPromotionsBuyXPayYPromotionId) | **Get** /buy_x_pay_y_promotions/{buyXPayYPromotionId} | Retrieve a buy x pay y promotion
[**PATCHBuyXPayYPromotionsBuyXPayYPromotionId**](BuyXPayYPromotionsApi.md#PATCHBuyXPayYPromotionsBuyXPayYPromotionId) | **Patch** /buy_x_pay_y_promotions/{buyXPayYPromotionId} | Update a buy x pay y promotion
[**POSTBuyXPayYPromotions**](BuyXPayYPromotionsApi.md#POSTBuyXPayYPromotions) | **Post** /buy_x_pay_y_promotions | Create a buy x pay y promotion



## DELETEBuyXPayYPromotionsBuyXPayYPromotionId

> DELETEBuyXPayYPromotionsBuyXPayYPromotionId(ctx, buyXPayYPromotionId).Execute()

Delete a buy x pay y promotion



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
    r, err := apiClient.BuyXPayYPromotionsApi.DELETEBuyXPayYPromotionsBuyXPayYPromotionId(context.Background(), buyXPayYPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyXPayYPromotionsApi.DELETEBuyXPayYPromotionsBuyXPayYPromotionId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEBuyXPayYPromotionsBuyXPayYPromotionIdRequest struct via the builder pattern


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


## GETBuyXPayYPromotions

> GETBuyXPayYPromotions200Response GETBuyXPayYPromotions(ctx).Execute()

List all buy x pay y promotions



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
    resp, r, err := apiClient.BuyXPayYPromotionsApi.GETBuyXPayYPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyXPayYPromotionsApi.GETBuyXPayYPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBuyXPayYPromotions`: GETBuyXPayYPromotions200Response
    fmt.Fprintf(os.Stdout, "Response from `BuyXPayYPromotionsApi.GETBuyXPayYPromotions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBuyXPayYPromotionsRequest struct via the builder pattern


### Return type

[**GETBuyXPayYPromotions200Response**](GETBuyXPayYPromotions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETBuyXPayYPromotionsBuyXPayYPromotionId

> GETBuyXPayYPromotionsBuyXPayYPromotionId200Response GETBuyXPayYPromotionsBuyXPayYPromotionId(ctx, buyXPayYPromotionId).Execute()

Retrieve a buy x pay y promotion



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
    resp, r, err := apiClient.BuyXPayYPromotionsApi.GETBuyXPayYPromotionsBuyXPayYPromotionId(context.Background(), buyXPayYPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyXPayYPromotionsApi.GETBuyXPayYPromotionsBuyXPayYPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBuyXPayYPromotionsBuyXPayYPromotionId`: GETBuyXPayYPromotionsBuyXPayYPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `BuyXPayYPromotionsApi.GETBuyXPayYPromotionsBuyXPayYPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyXPayYPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBuyXPayYPromotionsBuyXPayYPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETBuyXPayYPromotionsBuyXPayYPromotionId200Response**](GETBuyXPayYPromotionsBuyXPayYPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBuyXPayYPromotionsBuyXPayYPromotionId

> PATCHBuyXPayYPromotionsBuyXPayYPromotionId200Response PATCHBuyXPayYPromotionsBuyXPayYPromotionId(ctx, buyXPayYPromotionId).BuyXPayYPromotionUpdate(buyXPayYPromotionUpdate).Execute()

Update a buy x pay y promotion



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
    buyXPayYPromotionUpdate := *openapiclient.NewBuyXPayYPromotionUpdate(*openapiclient.NewBuyXPayYPromotionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHBuyXPayYPromotionsBuyXPayYPromotionId200ResponseDataAttributes())) // BuyXPayYPromotionUpdate | 
    buyXPayYPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyXPayYPromotionsApi.PATCHBuyXPayYPromotionsBuyXPayYPromotionId(context.Background(), buyXPayYPromotionId).BuyXPayYPromotionUpdate(buyXPayYPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyXPayYPromotionsApi.PATCHBuyXPayYPromotionsBuyXPayYPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHBuyXPayYPromotionsBuyXPayYPromotionId`: PATCHBuyXPayYPromotionsBuyXPayYPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `BuyXPayYPromotionsApi.PATCHBuyXPayYPromotionsBuyXPayYPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyXPayYPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBuyXPayYPromotionsBuyXPayYPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyXPayYPromotionUpdate** | [**BuyXPayYPromotionUpdate**](BuyXPayYPromotionUpdate.md) |  | 


### Return type

[**PATCHBuyXPayYPromotionsBuyXPayYPromotionId200Response**](PATCHBuyXPayYPromotionsBuyXPayYPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTBuyXPayYPromotions

> POSTBuyXPayYPromotions201Response POSTBuyXPayYPromotions(ctx).BuyXPayYPromotionCreate(buyXPayYPromotionCreate).Execute()

Create a buy x pay y promotion



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
    buyXPayYPromotionCreate := *openapiclient.NewBuyXPayYPromotionCreate(*openapiclient.NewBuyXPayYPromotionCreateData(interface{}(123), *openapiclient.NewPOSTBuyXPayYPromotions201ResponseDataAttributes(interface{}(Personal promotion), interface{}(2018-01-01T12:00:00.000Z), interface{}(2018-01-02T12:00:00.000Z), interface{}(3), interface{}(2)))) // BuyXPayYPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyXPayYPromotionsApi.POSTBuyXPayYPromotions(context.Background()).BuyXPayYPromotionCreate(buyXPayYPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyXPayYPromotionsApi.POSTBuyXPayYPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTBuyXPayYPromotions`: POSTBuyXPayYPromotions201Response
    fmt.Fprintf(os.Stdout, "Response from `BuyXPayYPromotionsApi.POSTBuyXPayYPromotions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBuyXPayYPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyXPayYPromotionCreate** | [**BuyXPayYPromotionCreate**](BuyXPayYPromotionCreate.md) |  | 

### Return type

[**POSTBuyXPayYPromotions201Response**](POSTBuyXPayYPromotions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

