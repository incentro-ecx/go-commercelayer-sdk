# \GiftCardRecipientsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEGiftCardRecipientsGiftCardRecipientId**](GiftCardRecipientsApi.md#DELETEGiftCardRecipientsGiftCardRecipientId) | **Delete** /gift_card_recipients/{giftCardRecipientId} | Delete a gift card recipient
[**GETGiftCardIdGiftCardRecipient**](GiftCardRecipientsApi.md#GETGiftCardIdGiftCardRecipient) | **Get** /gift_cards/{giftCardId}/gift_card_recipient | Retrieve the gift card recipient associated to the gift card
[**GETGiftCardRecipients**](GiftCardRecipientsApi.md#GETGiftCardRecipients) | **Get** /gift_card_recipients | List all gift card recipients
[**GETGiftCardRecipientsGiftCardRecipientId**](GiftCardRecipientsApi.md#GETGiftCardRecipientsGiftCardRecipientId) | **Get** /gift_card_recipients/{giftCardRecipientId} | Retrieve a gift card recipient
[**PATCHGiftCardRecipientsGiftCardRecipientId**](GiftCardRecipientsApi.md#PATCHGiftCardRecipientsGiftCardRecipientId) | **Patch** /gift_card_recipients/{giftCardRecipientId} | Update a gift card recipient
[**POSTGiftCardRecipients**](GiftCardRecipientsApi.md#POSTGiftCardRecipients) | **Post** /gift_card_recipients | Create a gift card recipient



## DELETEGiftCardRecipientsGiftCardRecipientId

> DELETEGiftCardRecipientsGiftCardRecipientId(ctx, giftCardRecipientId).Execute()

Delete a gift card recipient



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
    giftCardRecipientId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GiftCardRecipientsApi.DELETEGiftCardRecipientsGiftCardRecipientId(context.Background(), giftCardRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardRecipientsApi.DELETEGiftCardRecipientsGiftCardRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardRecipientId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEGiftCardRecipientsGiftCardRecipientIdRequest struct via the builder pattern


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


## GETGiftCardIdGiftCardRecipient

> GETGiftCardIdGiftCardRecipient(ctx, giftCardId).Execute()

Retrieve the gift card recipient associated to the gift card



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
    r, err := apiClient.GiftCardRecipientsApi.GETGiftCardIdGiftCardRecipient(context.Background(), giftCardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardRecipientsApi.GETGiftCardIdGiftCardRecipient``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETGiftCardIdGiftCardRecipientRequest struct via the builder pattern


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


## GETGiftCardRecipients

> GETGiftCardRecipients200Response GETGiftCardRecipients(ctx).Execute()

List all gift card recipients



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
    resp, r, err := apiClient.GiftCardRecipientsApi.GETGiftCardRecipients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardRecipientsApi.GETGiftCardRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETGiftCardRecipients`: GETGiftCardRecipients200Response
    fmt.Fprintf(os.Stdout, "Response from `GiftCardRecipientsApi.GETGiftCardRecipients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETGiftCardRecipientsRequest struct via the builder pattern


### Return type

[**GETGiftCardRecipients200Response**](GETGiftCardRecipients200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETGiftCardRecipientsGiftCardRecipientId

> GETGiftCardRecipientsGiftCardRecipientId200Response GETGiftCardRecipientsGiftCardRecipientId(ctx, giftCardRecipientId).Execute()

Retrieve a gift card recipient



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
    giftCardRecipientId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GiftCardRecipientsApi.GETGiftCardRecipientsGiftCardRecipientId(context.Background(), giftCardRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardRecipientsApi.GETGiftCardRecipientsGiftCardRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETGiftCardRecipientsGiftCardRecipientId`: GETGiftCardRecipientsGiftCardRecipientId200Response
    fmt.Fprintf(os.Stdout, "Response from `GiftCardRecipientsApi.GETGiftCardRecipientsGiftCardRecipientId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardRecipientId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETGiftCardRecipientsGiftCardRecipientIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETGiftCardRecipientsGiftCardRecipientId200Response**](GETGiftCardRecipientsGiftCardRecipientId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHGiftCardRecipientsGiftCardRecipientId

> PATCHGiftCardRecipientsGiftCardRecipientId200Response PATCHGiftCardRecipientsGiftCardRecipientId(ctx, giftCardRecipientId).GiftCardRecipientUpdate(giftCardRecipientUpdate).Execute()

Update a gift card recipient



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
    giftCardRecipientUpdate := *openapiclient.NewGiftCardRecipientUpdate(*openapiclient.NewGiftCardRecipientUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes())) // GiftCardRecipientUpdate | 
    giftCardRecipientId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GiftCardRecipientsApi.PATCHGiftCardRecipientsGiftCardRecipientId(context.Background(), giftCardRecipientId).GiftCardRecipientUpdate(giftCardRecipientUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardRecipientsApi.PATCHGiftCardRecipientsGiftCardRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHGiftCardRecipientsGiftCardRecipientId`: PATCHGiftCardRecipientsGiftCardRecipientId200Response
    fmt.Fprintf(os.Stdout, "Response from `GiftCardRecipientsApi.PATCHGiftCardRecipientsGiftCardRecipientId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardRecipientId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHGiftCardRecipientsGiftCardRecipientIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardRecipientUpdate** | [**GiftCardRecipientUpdate**](GiftCardRecipientUpdate.md) |  | 


### Return type

[**PATCHGiftCardRecipientsGiftCardRecipientId200Response**](PATCHGiftCardRecipientsGiftCardRecipientId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTGiftCardRecipients

> POSTGiftCardRecipients201Response POSTGiftCardRecipients(ctx).GiftCardRecipientCreate(giftCardRecipientCreate).Execute()

Create a gift card recipient



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
    giftCardRecipientCreate := *openapiclient.NewGiftCardRecipientCreate(*openapiclient.NewGiftCardRecipientCreateData(interface{}(123), *openapiclient.NewPOSTCouponRecipients201ResponseDataAttributes(interface{}(john@example.com)))) // GiftCardRecipientCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GiftCardRecipientsApi.POSTGiftCardRecipients(context.Background()).GiftCardRecipientCreate(giftCardRecipientCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardRecipientsApi.POSTGiftCardRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTGiftCardRecipients`: POSTGiftCardRecipients201Response
    fmt.Fprintf(os.Stdout, "Response from `GiftCardRecipientsApi.POSTGiftCardRecipients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTGiftCardRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardRecipientCreate** | [**GiftCardRecipientCreate**](GiftCardRecipientCreate.md) |  | 

### Return type

[**POSTGiftCardRecipients201Response**](POSTGiftCardRecipients201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

