# \CouponRecipientsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECouponRecipientsCouponRecipientId**](CouponRecipientsApi.md#DELETECouponRecipientsCouponRecipientId) | **Delete** /coupon_recipients/{couponRecipientId} | Delete a coupon recipient
[**GETCouponIdCouponRecipient**](CouponRecipientsApi.md#GETCouponIdCouponRecipient) | **Get** /coupons/{couponId}/coupon_recipient | Retrieve the coupon recipient associated to the coupon
[**GETCouponRecipients**](CouponRecipientsApi.md#GETCouponRecipients) | **Get** /coupon_recipients | List all coupon recipients
[**GETCouponRecipientsCouponRecipientId**](CouponRecipientsApi.md#GETCouponRecipientsCouponRecipientId) | **Get** /coupon_recipients/{couponRecipientId} | Retrieve a coupon recipient
[**PATCHCouponRecipientsCouponRecipientId**](CouponRecipientsApi.md#PATCHCouponRecipientsCouponRecipientId) | **Patch** /coupon_recipients/{couponRecipientId} | Update a coupon recipient
[**POSTCouponRecipients**](CouponRecipientsApi.md#POSTCouponRecipients) | **Post** /coupon_recipients | Create a coupon recipient



## DELETECouponRecipientsCouponRecipientId

> DELETECouponRecipientsCouponRecipientId(ctx, couponRecipientId).Execute()

Delete a coupon recipient



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
    couponRecipientId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CouponRecipientsApi.DELETECouponRecipientsCouponRecipientId(context.Background(), couponRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.DELETECouponRecipientsCouponRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponRecipientId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECouponRecipientsCouponRecipientIdRequest struct via the builder pattern


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


## GETCouponIdCouponRecipient

> GETCouponIdCouponRecipient(ctx, couponId).Execute()

Retrieve the coupon recipient associated to the coupon



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
    couponId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CouponRecipientsApi.GETCouponIdCouponRecipient(context.Background(), couponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.GETCouponIdCouponRecipient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponIdCouponRecipientRequest struct via the builder pattern


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


## GETCouponRecipients

> GETCouponRecipients200Response GETCouponRecipients(ctx).Execute()

List all coupon recipients



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
    resp, r, err := apiClient.CouponRecipientsApi.GETCouponRecipients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.GETCouponRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCouponRecipients`: GETCouponRecipients200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponRecipientsApi.GETCouponRecipients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponRecipientsRequest struct via the builder pattern


### Return type

[**GETCouponRecipients200Response**](GETCouponRecipients200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCouponRecipientsCouponRecipientId

> GETCouponRecipientsCouponRecipientId200Response GETCouponRecipientsCouponRecipientId(ctx, couponRecipientId).Execute()

Retrieve a coupon recipient



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
    couponRecipientId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponRecipientsApi.GETCouponRecipientsCouponRecipientId(context.Background(), couponRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.GETCouponRecipientsCouponRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCouponRecipientsCouponRecipientId`: GETCouponRecipientsCouponRecipientId200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponRecipientsApi.GETCouponRecipientsCouponRecipientId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponRecipientId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCouponRecipientsCouponRecipientIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCouponRecipientsCouponRecipientId200Response**](GETCouponRecipientsCouponRecipientId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCouponRecipientsCouponRecipientId

> PATCHCouponRecipientsCouponRecipientId200Response PATCHCouponRecipientsCouponRecipientId(ctx, couponRecipientId).CouponRecipientUpdate(couponRecipientUpdate).Execute()

Update a coupon recipient



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
    couponRecipientUpdate := *openapiclient.NewCouponRecipientUpdate(*openapiclient.NewCouponRecipientUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes())) // CouponRecipientUpdate | 
    couponRecipientId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponRecipientsApi.PATCHCouponRecipientsCouponRecipientId(context.Background(), couponRecipientId).CouponRecipientUpdate(couponRecipientUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.PATCHCouponRecipientsCouponRecipientId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCouponRecipientsCouponRecipientId`: PATCHCouponRecipientsCouponRecipientId200Response
    fmt.Fprintf(os.Stdout, "Response from `CouponRecipientsApi.PATCHCouponRecipientsCouponRecipientId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponRecipientId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCouponRecipientsCouponRecipientIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponRecipientUpdate** | [**CouponRecipientUpdate**](CouponRecipientUpdate.md) |  | 


### Return type

[**PATCHCouponRecipientsCouponRecipientId200Response**](PATCHCouponRecipientsCouponRecipientId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCouponRecipients

> POSTCouponRecipients201Response POSTCouponRecipients(ctx).CouponRecipientCreate(couponRecipientCreate).Execute()

Create a coupon recipient



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
    couponRecipientCreate := *openapiclient.NewCouponRecipientCreate(*openapiclient.NewCouponRecipientCreateData(interface{}(123), *openapiclient.NewPOSTCouponRecipients201ResponseDataAttributes(interface{}(john@example.com)))) // CouponRecipientCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponRecipientsApi.POSTCouponRecipients(context.Background()).CouponRecipientCreate(couponRecipientCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponRecipientsApi.POSTCouponRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCouponRecipients`: POSTCouponRecipients201Response
    fmt.Fprintf(os.Stdout, "Response from `CouponRecipientsApi.POSTCouponRecipients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCouponRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponRecipientCreate** | [**CouponRecipientCreate**](CouponRecipientCreate.md) |  | 

### Return type

[**POSTCouponRecipients201Response**](POSTCouponRecipients201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

