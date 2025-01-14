# \AdyenGatewaysApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAdyenGatewaysAdyenGatewayId**](AdyenGatewaysApi.md#DELETEAdyenGatewaysAdyenGatewayId) | **Delete** /adyen_gateways/{adyenGatewayId} | Delete an adyen gateway
[**GETAdyenGateways**](AdyenGatewaysApi.md#GETAdyenGateways) | **Get** /adyen_gateways | List all adyen gateways
[**GETAdyenGatewaysAdyenGatewayId**](AdyenGatewaysApi.md#GETAdyenGatewaysAdyenGatewayId) | **Get** /adyen_gateways/{adyenGatewayId} | Retrieve an adyen gateway
[**PATCHAdyenGatewaysAdyenGatewayId**](AdyenGatewaysApi.md#PATCHAdyenGatewaysAdyenGatewayId) | **Patch** /adyen_gateways/{adyenGatewayId} | Update an adyen gateway
[**POSTAdyenGateways**](AdyenGatewaysApi.md#POSTAdyenGateways) | **Post** /adyen_gateways | Create an adyen gateway



## DELETEAdyenGatewaysAdyenGatewayId

> DELETEAdyenGatewaysAdyenGatewayId(ctx, adyenGatewayId).Execute()

Delete an adyen gateway



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
    adyenGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdyenGatewaysApi.DELETEAdyenGatewaysAdyenGatewayId(context.Background(), adyenGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenGatewaysApi.DELETEAdyenGatewaysAdyenGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEAdyenGatewaysAdyenGatewayIdRequest struct via the builder pattern


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


## GETAdyenGateways

> GETAdyenGateways200Response GETAdyenGateways(ctx).Execute()

List all adyen gateways



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
    resp, r, err := apiClient.AdyenGatewaysApi.GETAdyenGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenGatewaysApi.GETAdyenGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAdyenGateways`: GETAdyenGateways200Response
    fmt.Fprintf(os.Stdout, "Response from `AdyenGatewaysApi.GETAdyenGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdyenGatewaysRequest struct via the builder pattern


### Return type

[**GETAdyenGateways200Response**](GETAdyenGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETAdyenGatewaysAdyenGatewayId

> GETAdyenGatewaysAdyenGatewayId200Response GETAdyenGatewaysAdyenGatewayId(ctx, adyenGatewayId).Execute()

Retrieve an adyen gateway



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
    adyenGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdyenGatewaysApi.GETAdyenGatewaysAdyenGatewayId(context.Background(), adyenGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenGatewaysApi.GETAdyenGatewaysAdyenGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAdyenGatewaysAdyenGatewayId`: GETAdyenGatewaysAdyenGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `AdyenGatewaysApi.GETAdyenGatewaysAdyenGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAdyenGatewaysAdyenGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETAdyenGatewaysAdyenGatewayId200Response**](GETAdyenGatewaysAdyenGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHAdyenGatewaysAdyenGatewayId

> PATCHAdyenGatewaysAdyenGatewayId200Response PATCHAdyenGatewaysAdyenGatewayId(ctx, adyenGatewayId).AdyenGatewayUpdate(adyenGatewayUpdate).Execute()

Update an adyen gateway



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
    adyenGatewayUpdate := *openapiclient.NewAdyenGatewayUpdate(*openapiclient.NewAdyenGatewayUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes())) // AdyenGatewayUpdate | 
    adyenGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdyenGatewaysApi.PATCHAdyenGatewaysAdyenGatewayId(context.Background(), adyenGatewayId).AdyenGatewayUpdate(adyenGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenGatewaysApi.PATCHAdyenGatewaysAdyenGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHAdyenGatewaysAdyenGatewayId`: PATCHAdyenGatewaysAdyenGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `AdyenGatewaysApi.PATCHAdyenGatewaysAdyenGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adyenGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAdyenGatewaysAdyenGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adyenGatewayUpdate** | [**AdyenGatewayUpdate**](AdyenGatewayUpdate.md) |  | 


### Return type

[**PATCHAdyenGatewaysAdyenGatewayId200Response**](PATCHAdyenGatewaysAdyenGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTAdyenGateways

> POSTAdyenGateways201Response POSTAdyenGateways(ctx).AdyenGatewayCreate(adyenGatewayCreate).Execute()

Create an adyen gateway



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
    adyenGatewayCreate := *openapiclient.NewAdyenGatewayCreate(*openapiclient.NewAdyenGatewayCreateData(interface{}(123), *openapiclient.NewPOSTAdyenGateways201ResponseDataAttributes(interface{}(US payment gateway), interface{}(xxxx-yyyy-zzzz), interface{}(xxxx-yyyy-zzzz), interface{}(1797a841fbb37ca7-AdyenDemo)))) // AdyenGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdyenGatewaysApi.POSTAdyenGateways(context.Background()).AdyenGatewayCreate(adyenGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdyenGatewaysApi.POSTAdyenGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTAdyenGateways`: POSTAdyenGateways201Response
    fmt.Fprintf(os.Stdout, "Response from `AdyenGatewaysApi.POSTAdyenGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTAdyenGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adyenGatewayCreate** | [**AdyenGatewayCreate**](AdyenGatewayCreate.md) |  | 

### Return type

[**POSTAdyenGateways201Response**](POSTAdyenGateways201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

