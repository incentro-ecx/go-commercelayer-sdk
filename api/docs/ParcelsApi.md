# \ParcelsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEParcelsParcelId**](ParcelsApi.md#DELETEParcelsParcelId) | **Delete** /parcels/{parcelId} | Delete a parcel
[**GETPackageIdParcels**](ParcelsApi.md#GETPackageIdParcels) | **Get** /packages/{packageId}/parcels | Retrieve the parcels associated to the package
[**GETParcelLineItemIdParcel**](ParcelsApi.md#GETParcelLineItemIdParcel) | **Get** /parcel_line_items/{parcelLineItemId}/parcel | Retrieve the parcel associated to the parcel line item
[**GETParcels**](ParcelsApi.md#GETParcels) | **Get** /parcels | List all parcels
[**GETParcelsParcelId**](ParcelsApi.md#GETParcelsParcelId) | **Get** /parcels/{parcelId} | Retrieve a parcel
[**GETShipmentIdParcels**](ParcelsApi.md#GETShipmentIdParcels) | **Get** /shipments/{shipmentId}/parcels | Retrieve the parcels associated to the shipment
[**PATCHParcelsParcelId**](ParcelsApi.md#PATCHParcelsParcelId) | **Patch** /parcels/{parcelId} | Update a parcel
[**POSTParcels**](ParcelsApi.md#POSTParcels) | **Post** /parcels | Create a parcel



## DELETEParcelsParcelId

> DELETEParcelsParcelId(ctx, parcelId).Execute()

Delete a parcel



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
    r, err := apiClient.ParcelsApi.DELETEParcelsParcelId(context.Background(), parcelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelsApi.DELETEParcelsParcelId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEParcelsParcelIdRequest struct via the builder pattern


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


## GETPackageIdParcels

> GETPackageIdParcels(ctx, packageId).Execute()

Retrieve the parcels associated to the package



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
    packageId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ParcelsApi.GETPackageIdParcels(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelsApi.GETPackageIdParcels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPackageIdParcelsRequest struct via the builder pattern


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


## GETParcelLineItemIdParcel

> GETParcelLineItemIdParcel(ctx, parcelLineItemId).Execute()

Retrieve the parcel associated to the parcel line item



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
    r, err := apiClient.ParcelsApi.GETParcelLineItemIdParcel(context.Background(), parcelLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelsApi.GETParcelLineItemIdParcel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETParcelLineItemIdParcelRequest struct via the builder pattern


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


## GETParcels

> GETParcels200Response GETParcels(ctx).Execute()

List all parcels



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
    resp, r, err := apiClient.ParcelsApi.GETParcels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelsApi.GETParcels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETParcels`: GETParcels200Response
    fmt.Fprintf(os.Stdout, "Response from `ParcelsApi.GETParcels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETParcelsRequest struct via the builder pattern


### Return type

[**GETParcels200Response**](GETParcels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETParcelsParcelId

> GETParcelsParcelId200Response GETParcelsParcelId(ctx, parcelId).Execute()

Retrieve a parcel



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
    resp, r, err := apiClient.ParcelsApi.GETParcelsParcelId(context.Background(), parcelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelsApi.GETParcelsParcelId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETParcelsParcelId`: GETParcelsParcelId200Response
    fmt.Fprintf(os.Stdout, "Response from `ParcelsApi.GETParcelsParcelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parcelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETParcelsParcelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETParcelsParcelId200Response**](GETParcelsParcelId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShipmentIdParcels

> GETShipmentIdParcels(ctx, shipmentId).Execute()

Retrieve the parcels associated to the shipment



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
    r, err := apiClient.ParcelsApi.GETShipmentIdParcels(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelsApi.GETShipmentIdParcels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShipmentIdParcelsRequest struct via the builder pattern


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


## PATCHParcelsParcelId

> PATCHParcelsParcelId200Response PATCHParcelsParcelId(ctx, parcelId).ParcelUpdate(parcelUpdate).Execute()

Update a parcel



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
    parcelUpdate := *openapiclient.NewParcelUpdate(*openapiclient.NewParcelUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHParcelsParcelId200ResponseDataAttributes())) // ParcelUpdate | 
    parcelId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParcelsApi.PATCHParcelsParcelId(context.Background(), parcelId).ParcelUpdate(parcelUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelsApi.PATCHParcelsParcelId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHParcelsParcelId`: PATCHParcelsParcelId200Response
    fmt.Fprintf(os.Stdout, "Response from `ParcelsApi.PATCHParcelsParcelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parcelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHParcelsParcelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parcelUpdate** | [**ParcelUpdate**](ParcelUpdate.md) |  | 


### Return type

[**PATCHParcelsParcelId200Response**](PATCHParcelsParcelId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTParcels

> POSTParcels201Response POSTParcels(ctx).ParcelCreate(parcelCreate).Execute()

Create a parcel



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
    parcelCreate := *openapiclient.NewParcelCreate(*openapiclient.NewParcelCreateData(interface{}(123), *openapiclient.NewPOSTParcels201ResponseDataAttributes(interface{}(1000.0), interface{}(gr)))) // ParcelCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParcelsApi.POSTParcels(context.Background()).ParcelCreate(parcelCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParcelsApi.POSTParcels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTParcels`: POSTParcels201Response
    fmt.Fprintf(os.Stdout, "Response from `ParcelsApi.POSTParcels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTParcelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parcelCreate** | [**ParcelCreate**](ParcelCreate.md) |  | 

### Return type

[**POSTParcels201Response**](POSTParcels201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

