# \TaxRulesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETETaxRulesTaxRuleId**](TaxRulesApi.md#DELETETaxRulesTaxRuleId) | **Delete** /tax_rules/{taxRuleId} | Delete a tax rule
[**GETManualTaxCalculatorIdTaxRules**](TaxRulesApi.md#GETManualTaxCalculatorIdTaxRules) | **Get** /manual_tax_calculators/{manualTaxCalculatorId}/tax_rules | Retrieve the tax rules associated to the manual tax calculator
[**GETTaxRules**](TaxRulesApi.md#GETTaxRules) | **Get** /tax_rules | List all tax rules
[**GETTaxRulesTaxRuleId**](TaxRulesApi.md#GETTaxRulesTaxRuleId) | **Get** /tax_rules/{taxRuleId} | Retrieve a tax rule
[**PATCHTaxRulesTaxRuleId**](TaxRulesApi.md#PATCHTaxRulesTaxRuleId) | **Patch** /tax_rules/{taxRuleId} | Update a tax rule
[**POSTTaxRules**](TaxRulesApi.md#POSTTaxRules) | **Post** /tax_rules | Create a tax rule



## DELETETaxRulesTaxRuleId

> DELETETaxRulesTaxRuleId(ctx, taxRuleId).Execute()

Delete a tax rule



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
    taxRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaxRulesApi.DELETETaxRulesTaxRuleId(context.Background(), taxRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.DELETETaxRulesTaxRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETETaxRulesTaxRuleIdRequest struct via the builder pattern


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


## GETManualTaxCalculatorIdTaxRules

> GETManualTaxCalculatorIdTaxRules(ctx, manualTaxCalculatorId).Execute()

Retrieve the tax rules associated to the manual tax calculator



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
    manualTaxCalculatorId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaxRulesApi.GETManualTaxCalculatorIdTaxRules(context.Background(), manualTaxCalculatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.GETManualTaxCalculatorIdTaxRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualTaxCalculatorId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETManualTaxCalculatorIdTaxRulesRequest struct via the builder pattern


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


## GETTaxRules

> GETTaxRules200Response GETTaxRules(ctx).Execute()

List all tax rules



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
    resp, r, err := apiClient.TaxRulesApi.GETTaxRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.GETTaxRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxRules`: GETTaxRules200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxRulesApi.GETTaxRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxRulesRequest struct via the builder pattern


### Return type

[**GETTaxRules200Response**](GETTaxRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTaxRulesTaxRuleId

> GETTaxRulesTaxRuleId200Response GETTaxRulesTaxRuleId(ctx, taxRuleId).Execute()

Retrieve a tax rule



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
    taxRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxRulesApi.GETTaxRulesTaxRuleId(context.Background(), taxRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.GETTaxRulesTaxRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxRulesTaxRuleId`: GETTaxRulesTaxRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxRulesApi.GETTaxRulesTaxRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxRulesTaxRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETTaxRulesTaxRuleId200Response**](GETTaxRulesTaxRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHTaxRulesTaxRuleId

> PATCHTaxRulesTaxRuleId200Response PATCHTaxRulesTaxRuleId(ctx, taxRuleId).TaxRuleUpdate(taxRuleUpdate).Execute()

Update a tax rule



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
    taxRuleUpdate := *openapiclient.NewTaxRuleUpdate(*openapiclient.NewTaxRuleUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHTaxRulesTaxRuleId200ResponseDataAttributes())) // TaxRuleUpdate | 
    taxRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxRulesApi.PATCHTaxRulesTaxRuleId(context.Background(), taxRuleId).TaxRuleUpdate(taxRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.PATCHTaxRulesTaxRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHTaxRulesTaxRuleId`: PATCHTaxRulesTaxRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxRulesApi.PATCHTaxRulesTaxRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHTaxRulesTaxRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxRuleUpdate** | [**TaxRuleUpdate**](TaxRuleUpdate.md) |  | 


### Return type

[**PATCHTaxRulesTaxRuleId200Response**](PATCHTaxRulesTaxRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTTaxRules

> POSTTaxRules201Response POSTTaxRules(ctx).TaxRuleCreate(taxRuleCreate).Execute()

Create a tax rule



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
    taxRuleCreate := *openapiclient.NewTaxRuleCreate(*openapiclient.NewTaxRuleCreateData(interface{}(123), *openapiclient.NewPOSTTaxRules201ResponseDataAttributes(interface{}(Fixed 22%)))) // TaxRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxRulesApi.POSTTaxRules(context.Background()).TaxRuleCreate(taxRuleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesApi.POSTTaxRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTTaxRules`: POSTTaxRules201Response
    fmt.Fprintf(os.Stdout, "Response from `TaxRulesApi.POSTTaxRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTaxRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxRuleCreate** | [**TaxRuleCreate**](TaxRuleCreate.md) |  | 

### Return type

[**POSTTaxRules201Response**](POSTTaxRules201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

