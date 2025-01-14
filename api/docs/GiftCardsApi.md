# \GiftCardsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEGiftCardsGiftCardId**](GiftCardsApi.md#DELETEGiftCardsGiftCardId) | **Delete** /gift_cards/{giftCardId} | Delete a gift card
[**GETGiftCards**](GiftCardsApi.md#GETGiftCards) | **Get** /gift_cards | List all gift cards
[**GETGiftCardsGiftCardId**](GiftCardsApi.md#GETGiftCardsGiftCardId) | **Get** /gift_cards/{giftCardId} | Retrieve a gift card
[**PATCHGiftCardsGiftCardId**](GiftCardsApi.md#PATCHGiftCardsGiftCardId) | **Patch** /gift_cards/{giftCardId} | Update a gift card
[**POSTGiftCards**](GiftCardsApi.md#POSTGiftCards) | **Post** /gift_cards | Create a gift card



## DELETEGiftCardsGiftCardId

> DELETEGiftCardsGiftCardId(ctx, giftCardId).Execute()

Delete a gift card



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
    giftCardId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GiftCardsApi.DELETEGiftCardsGiftCardId(context.Background(), giftCardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.DELETEGiftCardsGiftCardId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEGiftCardsGiftCardIdRequest struct via the builder pattern


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


## GETGiftCards

> GETGiftCards200Response GETGiftCards(ctx).Execute()

List all gift cards



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
    resp, r, err := apiClient.GiftCardsApi.GETGiftCards(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.GETGiftCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETGiftCards`: GETGiftCards200Response
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.GETGiftCards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETGiftCardsRequest struct via the builder pattern


### Return type

[**GETGiftCards200Response**](GETGiftCards200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETGiftCardsGiftCardId

> GETGiftCardsGiftCardId200Response GETGiftCardsGiftCardId(ctx, giftCardId).Execute()

Retrieve a gift card



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
    giftCardId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GiftCardsApi.GETGiftCardsGiftCardId(context.Background(), giftCardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.GETGiftCardsGiftCardId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETGiftCardsGiftCardId`: GETGiftCardsGiftCardId200Response
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.GETGiftCardsGiftCardId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETGiftCardsGiftCardIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETGiftCardsGiftCardId200Response**](GETGiftCardsGiftCardId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHGiftCardsGiftCardId

> PATCHGiftCardsGiftCardId200Response PATCHGiftCardsGiftCardId(ctx, giftCardId).GiftCardUpdate(giftCardUpdate).Execute()

Update a gift card



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
    giftCardUpdate := *openapiclient.NewGiftCardUpdate(*openapiclient.NewGiftCardUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHGiftCardsGiftCardId200ResponseDataAttributes())) // GiftCardUpdate | 
    giftCardId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GiftCardsApi.PATCHGiftCardsGiftCardId(context.Background(), giftCardId).GiftCardUpdate(giftCardUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.PATCHGiftCardsGiftCardId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHGiftCardsGiftCardId`: PATCHGiftCardsGiftCardId200Response
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.PATCHGiftCardsGiftCardId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHGiftCardsGiftCardIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardUpdate** | [**GiftCardUpdate**](GiftCardUpdate.md) |  | 


### Return type

[**PATCHGiftCardsGiftCardId200Response**](PATCHGiftCardsGiftCardId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTGiftCards

> POSTGiftCards201Response POSTGiftCards(ctx).GiftCardCreate(giftCardCreate).Execute()

Create a gift card



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
    giftCardCreate := *openapiclient.NewGiftCardCreate(*openapiclient.NewGiftCardCreateData(interface{}(123), *openapiclient.NewPOSTGiftCards201ResponseDataAttributes(interface{}(15000)))) // GiftCardCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GiftCardsApi.POSTGiftCards(context.Background()).GiftCardCreate(giftCardCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.POSTGiftCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTGiftCards`: POSTGiftCards201Response
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.POSTGiftCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTGiftCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardCreate** | [**GiftCardCreate**](GiftCardCreate.md) |  | 

### Return type

[**POSTGiftCards201Response**](POSTGiftCards201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

