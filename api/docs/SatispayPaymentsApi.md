# \SatispayPaymentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESatispayPaymentsSatispayPaymentId**](SatispayPaymentsApi.md#DELETESatispayPaymentsSatispayPaymentId) | **Delete** /satispay_payments/{satispayPaymentId} | Delete a satispay payment
[**GETSatispayGatewayIdSatispayPayments**](SatispayPaymentsApi.md#GETSatispayGatewayIdSatispayPayments) | **Get** /satispay_gateways/{satispayGatewayId}/satispay_payments | Retrieve the satispay payments associated to the satispay gateway
[**GETSatispayPayments**](SatispayPaymentsApi.md#GETSatispayPayments) | **Get** /satispay_payments | List all satispay payments
[**GETSatispayPaymentsSatispayPaymentId**](SatispayPaymentsApi.md#GETSatispayPaymentsSatispayPaymentId) | **Get** /satispay_payments/{satispayPaymentId} | Retrieve a satispay payment
[**PATCHSatispayPaymentsSatispayPaymentId**](SatispayPaymentsApi.md#PATCHSatispayPaymentsSatispayPaymentId) | **Patch** /satispay_payments/{satispayPaymentId} | Update a satispay payment
[**POSTSatispayPayments**](SatispayPaymentsApi.md#POSTSatispayPayments) | **Post** /satispay_payments | Create a satispay payment



## DELETESatispayPaymentsSatispayPaymentId

> DELETESatispayPaymentsSatispayPaymentId(ctx, satispayPaymentId).Execute()

Delete a satispay payment



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
    satispayPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SatispayPaymentsApi.DELETESatispayPaymentsSatispayPaymentId(context.Background(), satispayPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayPaymentsApi.DELETESatispayPaymentsSatispayPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETESatispayPaymentsSatispayPaymentIdRequest struct via the builder pattern


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


## GETSatispayGatewayIdSatispayPayments

> GETSatispayGatewayIdSatispayPayments(ctx, satispayGatewayId).Execute()

Retrieve the satispay payments associated to the satispay gateway



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
    satispayGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SatispayPaymentsApi.GETSatispayGatewayIdSatispayPayments(context.Background(), satispayGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayPaymentsApi.GETSatispayGatewayIdSatispayPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSatispayGatewayIdSatispayPaymentsRequest struct via the builder pattern


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


## GETSatispayPayments

> GETSatispayPayments200Response GETSatispayPayments(ctx).Execute()

List all satispay payments



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
    resp, r, err := apiClient.SatispayPaymentsApi.GETSatispayPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayPaymentsApi.GETSatispayPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSatispayPayments`: GETSatispayPayments200Response
    fmt.Fprintf(os.Stdout, "Response from `SatispayPaymentsApi.GETSatispayPayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETSatispayPaymentsRequest struct via the builder pattern


### Return type

[**GETSatispayPayments200Response**](GETSatispayPayments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSatispayPaymentsSatispayPaymentId

> GETSatispayPaymentsSatispayPaymentId200Response GETSatispayPaymentsSatispayPaymentId(ctx, satispayPaymentId).Execute()

Retrieve a satispay payment



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
    satispayPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatispayPaymentsApi.GETSatispayPaymentsSatispayPaymentId(context.Background(), satispayPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayPaymentsApi.GETSatispayPaymentsSatispayPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSatispayPaymentsSatispayPaymentId`: GETSatispayPaymentsSatispayPaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `SatispayPaymentsApi.GETSatispayPaymentsSatispayPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSatispayPaymentsSatispayPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETSatispayPaymentsSatispayPaymentId200Response**](GETSatispayPaymentsSatispayPaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHSatispayPaymentsSatispayPaymentId

> PATCHSatispayPaymentsSatispayPaymentId200Response PATCHSatispayPaymentsSatispayPaymentId(ctx, satispayPaymentId).SatispayPaymentUpdate(satispayPaymentUpdate).Execute()

Update a satispay payment



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
    satispayPaymentUpdate := *openapiclient.NewSatispayPaymentUpdate(*openapiclient.NewSatispayPaymentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes())) // SatispayPaymentUpdate | 
    satispayPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatispayPaymentsApi.PATCHSatispayPaymentsSatispayPaymentId(context.Background(), satispayPaymentId).SatispayPaymentUpdate(satispayPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayPaymentsApi.PATCHSatispayPaymentsSatispayPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHSatispayPaymentsSatispayPaymentId`: PATCHSatispayPaymentsSatispayPaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `SatispayPaymentsApi.PATCHSatispayPaymentsSatispayPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**satispayPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHSatispayPaymentsSatispayPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **satispayPaymentUpdate** | [**SatispayPaymentUpdate**](SatispayPaymentUpdate.md) |  | 


### Return type

[**PATCHSatispayPaymentsSatispayPaymentId200Response**](PATCHSatispayPaymentsSatispayPaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTSatispayPayments

> POSTSatispayPayments201Response POSTSatispayPayments(ctx).SatispayPaymentCreate(satispayPaymentCreate).Execute()

Create a satispay payment



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
    satispayPaymentCreate := *openapiclient.NewSatispayPaymentCreate(*openapiclient.NewSatispayPaymentCreateData(interface{}(123), *openapiclient.NewPOSTSatispayPayments201ResponseDataAttributes())) // SatispayPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatispayPaymentsApi.POSTSatispayPayments(context.Background()).SatispayPaymentCreate(satispayPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatispayPaymentsApi.POSTSatispayPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTSatispayPayments`: POSTSatispayPayments201Response
    fmt.Fprintf(os.Stdout, "Response from `SatispayPaymentsApi.POSTSatispayPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTSatispayPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **satispayPaymentCreate** | [**SatispayPaymentCreate**](SatispayPaymentCreate.md) |  | 

### Return type

[**POSTSatispayPayments201Response**](POSTSatispayPayments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

