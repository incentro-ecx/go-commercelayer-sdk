# \ParcelLineItemsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEParcelLineItemsParcelLineItemId**](ParcelLineItemsApi.md#DELETEParcelLineItemsParcelLineItemId) | **Delete** /parcel_line_items/{parcelLineItemId} | Delete a parcel line item
[**GETParcelIdParcelLineItems**](ParcelLineItemsApi.md#GETParcelIdParcelLineItems) | **Get** /parcels/{parcelId}/parcel_line_items | Retrieve the parcel line items associated to the parcel
[**GETParcelLineItems**](ParcelLineItemsApi.md#GETParcelLineItems) | **Get** /parcel_line_items | List all parcel line items
[**GETParcelLineItemsParcelLineItemId**](ParcelLineItemsApi.md#GETParcelLineItemsParcelLineItemId) | **Get** /parcel_line_items/{parcelLineItemId} | Retrieve a parcel line item
[**PATCHParcelLineItemsParcelLineItemId**](ParcelLineItemsApi.md#PATCHParcelLineItemsParcelLineItemId) | **Patch** /parcel_line_items/{parcelLineItemId} | Update a parcel line item
[**POSTParcelLineItems**](ParcelLineItemsApi.md#POSTParcelLineItems) | **Post** /parcel_line_items | Create a parcel line item



## DELETEParcelLineItemsParcelLineItemId

> DELETEParcelLineItemsParcelLineItemId(ctx, parcelLineItemId).Execute()

Delete a parcel line item



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
    parcelLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ParcelLineItemsApi.DELETEParcelLineItemsParcelLineItemId(context.Background(), parcelLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelLineItemsApi.DELETEParcelLineItemsParcelLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parcelLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEParcelLineItemsParcelLineItemIdRequest struct via the builder pattern


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


## GETParcelIdParcelLineItems

> GETParcelIdParcelLineItems(ctx, parcelId).Execute()

Retrieve the parcel line items associated to the parcel



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
    parcelId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ParcelLineItemsApi.GETParcelIdParcelLineItems(context.Background(), parcelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelLineItemsApi.GETParcelIdParcelLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parcelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETParcelIdParcelLineItemsRequest struct via the builder pattern


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


## GETParcelLineItems

> GETParcelLineItems200Response GETParcelLineItems(ctx).Execute()

List all parcel line items



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
    resp, r, err := apiClient.ParcelLineItemsApi.GETParcelLineItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelLineItemsApi.GETParcelLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETParcelLineItems`: GETParcelLineItems200Response
    fmt.Fprintf(os.Stdout, "Response from `ParcelLineItemsApi.GETParcelLineItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETParcelLineItemsRequest struct via the builder pattern


### Return type

[**GETParcelLineItems200Response**](GETParcelLineItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETParcelLineItemsParcelLineItemId

> GETParcelLineItemsParcelLineItemId200Response GETParcelLineItemsParcelLineItemId(ctx, parcelLineItemId).Execute()

Retrieve a parcel line item



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
    parcelLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParcelLineItemsApi.GETParcelLineItemsParcelLineItemId(context.Background(), parcelLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelLineItemsApi.GETParcelLineItemsParcelLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETParcelLineItemsParcelLineItemId`: GETParcelLineItemsParcelLineItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `ParcelLineItemsApi.GETParcelLineItemsParcelLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parcelLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETParcelLineItemsParcelLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETParcelLineItemsParcelLineItemId200Response**](GETParcelLineItemsParcelLineItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHParcelLineItemsParcelLineItemId

> PATCHParcelLineItemsParcelLineItemId200Response PATCHParcelLineItemsParcelLineItemId(ctx, parcelLineItemId).ParcelLineItemUpdate(parcelLineItemUpdate).Execute()

Update a parcel line item



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
    parcelLineItemUpdate := *openapiclient.NewParcelLineItemUpdate(*openapiclient.NewParcelLineItemUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes())) // ParcelLineItemUpdate | 
    parcelLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParcelLineItemsApi.PATCHParcelLineItemsParcelLineItemId(context.Background(), parcelLineItemId).ParcelLineItemUpdate(parcelLineItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelLineItemsApi.PATCHParcelLineItemsParcelLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHParcelLineItemsParcelLineItemId`: PATCHParcelLineItemsParcelLineItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `ParcelLineItemsApi.PATCHParcelLineItemsParcelLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parcelLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHParcelLineItemsParcelLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parcelLineItemUpdate** | [**ParcelLineItemUpdate**](ParcelLineItemUpdate.md) |  | 


### Return type

[**PATCHParcelLineItemsParcelLineItemId200Response**](PATCHParcelLineItemsParcelLineItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTParcelLineItems

> POSTParcelLineItems201Response POSTParcelLineItems(ctx).ParcelLineItemCreate(parcelLineItemCreate).Execute()

Create a parcel line item



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
    parcelLineItemCreate := *openapiclient.NewParcelLineItemCreate(*openapiclient.NewParcelLineItemCreateData(interface{}(123), *openapiclient.NewPOSTParcelLineItems201ResponseDataAttributes(interface{}(4)))) // ParcelLineItemCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParcelLineItemsApi.POSTParcelLineItems(context.Background()).ParcelLineItemCreate(parcelLineItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelLineItemsApi.POSTParcelLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTParcelLineItems`: POSTParcelLineItems201Response
    fmt.Fprintf(os.Stdout, "Response from `ParcelLineItemsApi.POSTParcelLineItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTParcelLineItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parcelLineItemCreate** | [**ParcelLineItemCreate**](ParcelLineItemCreate.md) |  | 

### Return type

[**POSTParcelLineItems201Response**](POSTParcelLineItems201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

