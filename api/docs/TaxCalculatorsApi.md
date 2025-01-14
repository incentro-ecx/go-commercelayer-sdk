# \TaxCalculatorsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETMarketIdTaxCalculator**](TaxCalculatorsApi.md#GETMarketIdTaxCalculator) | **Get** /markets/{marketId}/tax_calculator | Retrieve the tax calculator associated to the market
[**GETTaxCalculators**](TaxCalculatorsApi.md#GETTaxCalculators) | **Get** /tax_calculators | List all tax calculators
[**GETTaxCalculatorsTaxCalculatorId**](TaxCalculatorsApi.md#GETTaxCalculatorsTaxCalculatorId) | **Get** /tax_calculators/{taxCalculatorId} | Retrieve a tax calculator



## GETMarketIdTaxCalculator

> GETMarketIdTaxCalculator(ctx, marketId).Execute()

Retrieve the tax calculator associated to the market



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
    marketId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaxCalculatorsApi.GETMarketIdTaxCalculator(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCalculatorsApi.GETMarketIdTaxCalculator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdTaxCalculatorRequest struct via the builder pattern


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


## GETTaxCalculators

> GETTaxCalculators200Response GETTaxCalculators(ctx).Execute()

List all tax calculators



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
    resp, r, err := apiClient.TaxCalculatorsApi.GETTaxCalculators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCalculatorsApi.GETTaxCalculators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxCalculators`: GETTaxCalculators200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxCalculatorsApi.GETTaxCalculators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxCalculatorsRequest struct via the builder pattern


### Return type

[**GETTaxCalculators200Response**](GETTaxCalculators200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTaxCalculatorsTaxCalculatorId

> GETTaxCalculatorsTaxCalculatorId200Response GETTaxCalculatorsTaxCalculatorId(ctx, taxCalculatorId).Execute()

Retrieve a tax calculator



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
    taxCalculatorId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCalculatorsApi.GETTaxCalculatorsTaxCalculatorId(context.Background(), taxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCalculatorsApi.GETTaxCalculatorsTaxCalculatorId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxCalculatorsTaxCalculatorId`: GETTaxCalculatorsTaxCalculatorId200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxCalculatorsApi.GETTaxCalculatorsTaxCalculatorId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCalculatorId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxCalculatorsTaxCalculatorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETTaxCalculatorsTaxCalculatorId200Response**](GETTaxCalculatorsTaxCalculatorId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

