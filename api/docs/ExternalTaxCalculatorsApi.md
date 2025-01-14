# \ExternalTaxCalculatorsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExternalTaxCalculatorsExternalTaxCalculatorId**](ExternalTaxCalculatorsApi.md#DELETEExternalTaxCalculatorsExternalTaxCalculatorId) | **Delete** /external_tax_calculators/{externalTaxCalculatorId} | Delete an external tax calculator
[**GETExternalTaxCalculators**](ExternalTaxCalculatorsApi.md#GETExternalTaxCalculators) | **Get** /external_tax_calculators | List all external tax calculators
[**GETExternalTaxCalculatorsExternalTaxCalculatorId**](ExternalTaxCalculatorsApi.md#GETExternalTaxCalculatorsExternalTaxCalculatorId) | **Get** /external_tax_calculators/{externalTaxCalculatorId} | Retrieve an external tax calculator
[**PATCHExternalTaxCalculatorsExternalTaxCalculatorId**](ExternalTaxCalculatorsApi.md#PATCHExternalTaxCalculatorsExternalTaxCalculatorId) | **Patch** /external_tax_calculators/{externalTaxCalculatorId} | Update an external tax calculator
[**POSTExternalTaxCalculators**](ExternalTaxCalculatorsApi.md#POSTExternalTaxCalculators) | **Post** /external_tax_calculators | Create an external tax calculator



## DELETEExternalTaxCalculatorsExternalTaxCalculatorId

> DELETEExternalTaxCalculatorsExternalTaxCalculatorId(ctx, externalTaxCalculatorId).Execute()

Delete an external tax calculator



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
    externalTaxCalculatorId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExternalTaxCalculatorsApi.DELETEExternalTaxCalculatorsExternalTaxCalculatorId(context.Background(), externalTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaxCalculatorsApi.DELETEExternalTaxCalculatorsExternalTaxCalculatorId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalTaxCalculatorId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct via the builder pattern


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


## GETExternalTaxCalculators

> GETExternalTaxCalculators200Response GETExternalTaxCalculators(ctx).Execute()

List all external tax calculators



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
    resp, r, err := apiClient.ExternalTaxCalculatorsApi.GETExternalTaxCalculators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaxCalculatorsApi.GETExternalTaxCalculators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalTaxCalculators`: GETExternalTaxCalculators200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalTaxCalculatorsApi.GETExternalTaxCalculators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalTaxCalculatorsRequest struct via the builder pattern


### Return type

[**GETExternalTaxCalculators200Response**](GETExternalTaxCalculators200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalTaxCalculatorsExternalTaxCalculatorId

> GETExternalTaxCalculatorsExternalTaxCalculatorId200Response GETExternalTaxCalculatorsExternalTaxCalculatorId(ctx, externalTaxCalculatorId).Execute()

Retrieve an external tax calculator



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
    externalTaxCalculatorId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalTaxCalculatorsApi.GETExternalTaxCalculatorsExternalTaxCalculatorId(context.Background(), externalTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaxCalculatorsApi.GETExternalTaxCalculatorsExternalTaxCalculatorId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalTaxCalculatorsExternalTaxCalculatorId`: GETExternalTaxCalculatorsExternalTaxCalculatorId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalTaxCalculatorsApi.GETExternalTaxCalculatorsExternalTaxCalculatorId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalTaxCalculatorId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETExternalTaxCalculatorsExternalTaxCalculatorId200Response**](GETExternalTaxCalculatorsExternalTaxCalculatorId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHExternalTaxCalculatorsExternalTaxCalculatorId

> PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response PATCHExternalTaxCalculatorsExternalTaxCalculatorId(ctx, externalTaxCalculatorId).ExternalTaxCalculatorUpdate(externalTaxCalculatorUpdate).Execute()

Update an external tax calculator



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
    externalTaxCalculatorUpdate := *openapiclient.NewExternalTaxCalculatorUpdate(*openapiclient.NewExternalTaxCalculatorUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseDataAttributes())) // ExternalTaxCalculatorUpdate | 
    externalTaxCalculatorId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalTaxCalculatorsApi.PATCHExternalTaxCalculatorsExternalTaxCalculatorId(context.Background(), externalTaxCalculatorId).ExternalTaxCalculatorUpdate(externalTaxCalculatorUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaxCalculatorsApi.PATCHExternalTaxCalculatorsExternalTaxCalculatorId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHExternalTaxCalculatorsExternalTaxCalculatorId`: PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalTaxCalculatorsApi.PATCHExternalTaxCalculatorsExternalTaxCalculatorId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalTaxCalculatorId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalTaxCalculatorUpdate** | [**ExternalTaxCalculatorUpdate**](ExternalTaxCalculatorUpdate.md) |  | 


### Return type

[**PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response**](PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTExternalTaxCalculators

> POSTExternalTaxCalculators201Response POSTExternalTaxCalculators(ctx).ExternalTaxCalculatorCreate(externalTaxCalculatorCreate).Execute()

Create an external tax calculator



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
    externalTaxCalculatorCreate := *openapiclient.NewExternalTaxCalculatorCreate(*openapiclient.NewExternalTaxCalculatorCreateData(interface{}(123), *openapiclient.NewPOSTExternalTaxCalculators201ResponseDataAttributes(interface{}(Personal tax calculator), interface{}(https://external_calculator.yourbrand.com)))) // ExternalTaxCalculatorCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalTaxCalculatorsApi.POSTExternalTaxCalculators(context.Background()).ExternalTaxCalculatorCreate(externalTaxCalculatorCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTaxCalculatorsApi.POSTExternalTaxCalculators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTExternalTaxCalculators`: POSTExternalTaxCalculators201Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalTaxCalculatorsApi.POSTExternalTaxCalculators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExternalTaxCalculatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalTaxCalculatorCreate** | [**ExternalTaxCalculatorCreate**](ExternalTaxCalculatorCreate.md) |  | 

### Return type

[**POSTExternalTaxCalculators201Response**](POSTExternalTaxCalculators201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

