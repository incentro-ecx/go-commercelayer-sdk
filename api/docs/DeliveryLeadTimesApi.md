# \DeliveryLeadTimesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEDeliveryLeadTimesDeliveryLeadTimeId**](DeliveryLeadTimesApi.md#DELETEDeliveryLeadTimesDeliveryLeadTimeId) | **Delete** /delivery_lead_times/{deliveryLeadTimeId} | Delete a delivery lead time
[**GETDeliveryLeadTimes**](DeliveryLeadTimesApi.md#GETDeliveryLeadTimes) | **Get** /delivery_lead_times | List all delivery lead times
[**GETDeliveryLeadTimesDeliveryLeadTimeId**](DeliveryLeadTimesApi.md#GETDeliveryLeadTimesDeliveryLeadTimeId) | **Get** /delivery_lead_times/{deliveryLeadTimeId} | Retrieve a delivery lead time
[**GETShipmentIdDeliveryLeadTime**](DeliveryLeadTimesApi.md#GETShipmentIdDeliveryLeadTime) | **Get** /shipments/{shipmentId}/delivery_lead_time | Retrieve the delivery lead time associated to the shipment
[**GETShippingMethodIdDeliveryLeadTimeForShipment**](DeliveryLeadTimesApi.md#GETShippingMethodIdDeliveryLeadTimeForShipment) | **Get** /shipping_methods/{shippingMethodId}/delivery_lead_time_for_shipment | Retrieve the delivery lead time for shipment associated to the shipping method
[**GETSkuIdDeliveryLeadTimes**](DeliveryLeadTimesApi.md#GETSkuIdDeliveryLeadTimes) | **Get** /skus/{skuId}/delivery_lead_times | Retrieve the delivery lead times associated to the SKU
[**PATCHDeliveryLeadTimesDeliveryLeadTimeId**](DeliveryLeadTimesApi.md#PATCHDeliveryLeadTimesDeliveryLeadTimeId) | **Patch** /delivery_lead_times/{deliveryLeadTimeId} | Update a delivery lead time
[**POSTDeliveryLeadTimes**](DeliveryLeadTimesApi.md#POSTDeliveryLeadTimes) | **Post** /delivery_lead_times | Create a delivery lead time



## DELETEDeliveryLeadTimesDeliveryLeadTimeId

> DELETEDeliveryLeadTimesDeliveryLeadTimeId(ctx, deliveryLeadTimeId).Execute()

Delete a delivery lead time



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
    deliveryLeadTimeId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryLeadTimesApi.DELETEDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.DELETEDeliveryLeadTimesDeliveryLeadTimeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryLeadTimeId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEDeliveryLeadTimesDeliveryLeadTimeIdRequest struct via the builder pattern


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


## GETDeliveryLeadTimes

> GETDeliveryLeadTimes200Response GETDeliveryLeadTimes(ctx).Execute()

List all delivery lead times



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
    resp, r, err := apiClient.DeliveryLeadTimesApi.GETDeliveryLeadTimes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETDeliveryLeadTimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETDeliveryLeadTimes`: GETDeliveryLeadTimes200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryLeadTimesApi.GETDeliveryLeadTimes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETDeliveryLeadTimesRequest struct via the builder pattern


### Return type

[**GETDeliveryLeadTimes200Response**](GETDeliveryLeadTimes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETDeliveryLeadTimesDeliveryLeadTimeId

> GETDeliveryLeadTimesDeliveryLeadTimeId200Response GETDeliveryLeadTimesDeliveryLeadTimeId(ctx, deliveryLeadTimeId).Execute()

Retrieve a delivery lead time



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
    deliveryLeadTimeId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.GETDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETDeliveryLeadTimesDeliveryLeadTimeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETDeliveryLeadTimesDeliveryLeadTimeId`: GETDeliveryLeadTimesDeliveryLeadTimeId200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryLeadTimesApi.GETDeliveryLeadTimesDeliveryLeadTimeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryLeadTimeId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETDeliveryLeadTimesDeliveryLeadTimeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETDeliveryLeadTimesDeliveryLeadTimeId200Response**](GETDeliveryLeadTimesDeliveryLeadTimeId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShipmentIdDeliveryLeadTime

> GETShipmentIdDeliveryLeadTime(ctx, shipmentId).Execute()

Retrieve the delivery lead time associated to the shipment



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
    shipmentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryLeadTimesApi.GETShipmentIdDeliveryLeadTime(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETShipmentIdDeliveryLeadTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShipmentIdDeliveryLeadTimeRequest struct via the builder pattern


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


## GETShippingMethodIdDeliveryLeadTimeForShipment

> GETShippingMethodIdDeliveryLeadTimeForShipment(ctx, shippingMethodId).Execute()

Retrieve the delivery lead time for shipment associated to the shipping method



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
    shippingMethodId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeliveryLeadTimesApi.GETShippingMethodIdDeliveryLeadTimeForShipment(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETShippingMethodIdDeliveryLeadTimeForShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodIdDeliveryLeadTimeForShipmentRequest struct via the builder pattern


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


## GETSkuIdDeliveryLeadTimes

> GETSkuIdDeliveryLeadTimes(ctx, skuId).Execute()

Retrieve the delivery lead times associated to the SKU



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
    r, err := apiClient.DeliveryLeadTimesApi.GETSkuIdDeliveryLeadTimes(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.GETSkuIdDeliveryLeadTimes``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuIdDeliveryLeadTimesRequest struct via the builder pattern


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


## PATCHDeliveryLeadTimesDeliveryLeadTimeId

> PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response PATCHDeliveryLeadTimesDeliveryLeadTimeId(ctx, deliveryLeadTimeId).DeliveryLeadTimeUpdate(deliveryLeadTimeUpdate).Execute()

Update a delivery lead time



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
    deliveryLeadTimeUpdate := *openapiclient.NewDeliveryLeadTimeUpdate(*openapiclient.NewDeliveryLeadTimeUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes())) // DeliveryLeadTimeUpdate | 
    deliveryLeadTimeId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.PATCHDeliveryLeadTimesDeliveryLeadTimeId(context.Background(), deliveryLeadTimeId).DeliveryLeadTimeUpdate(deliveryLeadTimeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.PATCHDeliveryLeadTimesDeliveryLeadTimeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHDeliveryLeadTimesDeliveryLeadTimeId`: PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryLeadTimesApi.PATCHDeliveryLeadTimesDeliveryLeadTimeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryLeadTimeId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryLeadTimeUpdate** | [**DeliveryLeadTimeUpdate**](DeliveryLeadTimeUpdate.md) |  | 


### Return type

[**PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response**](PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTDeliveryLeadTimes

> POSTDeliveryLeadTimes201Response POSTDeliveryLeadTimes(ctx).DeliveryLeadTimeCreate(deliveryLeadTimeCreate).Execute()

Create a delivery lead time



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
    deliveryLeadTimeCreate := *openapiclient.NewDeliveryLeadTimeCreate(*openapiclient.NewDeliveryLeadTimeCreateData(interface{}(123), *openapiclient.NewPOSTDeliveryLeadTimes201ResponseDataAttributes(interface{}(48), interface{}(72)))) // DeliveryLeadTimeCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryLeadTimesApi.POSTDeliveryLeadTimes(context.Background()).DeliveryLeadTimeCreate(deliveryLeadTimeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryLeadTimesApi.POSTDeliveryLeadTimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTDeliveryLeadTimes`: POSTDeliveryLeadTimes201Response
    fmt.Fprintf(os.Stdout, "Response from `DeliveryLeadTimesApi.POSTDeliveryLeadTimes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTDeliveryLeadTimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryLeadTimeCreate** | [**DeliveryLeadTimeCreate**](DeliveryLeadTimeCreate.md) |  | 

### Return type

[**POSTDeliveryLeadTimes201Response**](POSTDeliveryLeadTimes201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

