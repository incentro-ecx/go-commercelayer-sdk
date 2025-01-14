# \CarrierAccountsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECarrierAccountsCarrierAccountId**](CarrierAccountsApi.md#DELETECarrierAccountsCarrierAccountId) | **Delete** /carrier_accounts/{carrierAccountId} | Delete a carrier account
[**GETCarrierAccounts**](CarrierAccountsApi.md#GETCarrierAccounts) | **Get** /carrier_accounts | List all carrier accounts
[**GETCarrierAccountsCarrierAccountId**](CarrierAccountsApi.md#GETCarrierAccountsCarrierAccountId) | **Get** /carrier_accounts/{carrierAccountId} | Retrieve a carrier account
[**GETShipmentIdCarrierAccounts**](CarrierAccountsApi.md#GETShipmentIdCarrierAccounts) | **Get** /shipments/{shipmentId}/carrier_accounts | Retrieve the carrier accounts associated to the shipment
[**PATCHCarrierAccountsCarrierAccountId**](CarrierAccountsApi.md#PATCHCarrierAccountsCarrierAccountId) | **Patch** /carrier_accounts/{carrierAccountId} | Update a carrier account
[**POSTCarrierAccounts**](CarrierAccountsApi.md#POSTCarrierAccounts) | **Post** /carrier_accounts | Create a carrier account



## DELETECarrierAccountsCarrierAccountId

> DELETECarrierAccountsCarrierAccountId(ctx, carrierAccountId).Execute()

Delete a carrier account



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
    carrierAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CarrierAccountsApi.DELETECarrierAccountsCarrierAccountId(context.Background(), carrierAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.DELETECarrierAccountsCarrierAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**carrierAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECarrierAccountsCarrierAccountIdRequest struct via the builder pattern


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


## GETCarrierAccounts

> GETCarrierAccounts200Response GETCarrierAccounts(ctx).Execute()

List all carrier accounts



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
    resp, r, err := apiClient.CarrierAccountsApi.GETCarrierAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.GETCarrierAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCarrierAccounts`: GETCarrierAccounts200Response
    fmt.Fprintf(os.Stdout, "Response from `CarrierAccountsApi.GETCarrierAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCarrierAccountsRequest struct via the builder pattern


### Return type

[**GETCarrierAccounts200Response**](GETCarrierAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCarrierAccountsCarrierAccountId

> GETCarrierAccountsCarrierAccountId200Response GETCarrierAccountsCarrierAccountId(ctx, carrierAccountId).Execute()

Retrieve a carrier account



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
    carrierAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CarrierAccountsApi.GETCarrierAccountsCarrierAccountId(context.Background(), carrierAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.GETCarrierAccountsCarrierAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCarrierAccountsCarrierAccountId`: GETCarrierAccountsCarrierAccountId200Response
    fmt.Fprintf(os.Stdout, "Response from `CarrierAccountsApi.GETCarrierAccountsCarrierAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**carrierAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCarrierAccountsCarrierAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCarrierAccountsCarrierAccountId200Response**](GETCarrierAccountsCarrierAccountId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShipmentIdCarrierAccounts

> GETShipmentIdCarrierAccounts(ctx, shipmentId).Execute()

Retrieve the carrier accounts associated to the shipment



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
    r, err := apiClient.CarrierAccountsApi.GETShipmentIdCarrierAccounts(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.GETShipmentIdCarrierAccounts``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShipmentIdCarrierAccountsRequest struct via the builder pattern


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


## PATCHCarrierAccountsCarrierAccountId

> PATCHCarrierAccountsCarrierAccountId200Response PATCHCarrierAccountsCarrierAccountId(ctx, carrierAccountId).CarrierAccountUpdate(carrierAccountUpdate).Execute()

Update a carrier account



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
    carrierAccountUpdate := *openapiclient.NewCarrierAccountUpdate(*openapiclient.NewCarrierAccountUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes())) // CarrierAccountUpdate | 
    carrierAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CarrierAccountsApi.PATCHCarrierAccountsCarrierAccountId(context.Background(), carrierAccountId).CarrierAccountUpdate(carrierAccountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.PATCHCarrierAccountsCarrierAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCarrierAccountsCarrierAccountId`: PATCHCarrierAccountsCarrierAccountId200Response
    fmt.Fprintf(os.Stdout, "Response from `CarrierAccountsApi.PATCHCarrierAccountsCarrierAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**carrierAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCarrierAccountsCarrierAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carrierAccountUpdate** | [**CarrierAccountUpdate**](CarrierAccountUpdate.md) |  | 


### Return type

[**PATCHCarrierAccountsCarrierAccountId200Response**](PATCHCarrierAccountsCarrierAccountId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCarrierAccounts

> POSTCarrierAccounts201Response POSTCarrierAccounts(ctx).CarrierAccountCreate(carrierAccountCreate).Execute()

Create a carrier account



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
    carrierAccountCreate := *openapiclient.NewCarrierAccountCreate(*openapiclient.NewCarrierAccountCreateData(interface{}(123), *openapiclient.NewPOSTCarrierAccounts201ResponseDataAttributes(interface{}(Accurate), interface{}(AccurateAccount), interface{}({username=xxxx, password=secret})))) // CarrierAccountCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CarrierAccountsApi.POSTCarrierAccounts(context.Background()).CarrierAccountCreate(carrierAccountCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierAccountsApi.POSTCarrierAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCarrierAccounts`: POSTCarrierAccounts201Response
    fmt.Fprintf(os.Stdout, "Response from `CarrierAccountsApi.POSTCarrierAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCarrierAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carrierAccountCreate** | [**CarrierAccountCreate**](CarrierAccountCreate.md) |  | 

### Return type

[**POSTCarrierAccounts201Response**](POSTCarrierAccounts201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

