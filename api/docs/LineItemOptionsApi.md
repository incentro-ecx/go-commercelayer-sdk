# \LineItemOptionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETELineItemOptionsLineItemOptionId**](LineItemOptionsApi.md#DELETELineItemOptionsLineItemOptionId) | **Delete** /line_item_options/{lineItemOptionId} | Delete a line item option
[**GETLineItemIdLineItemOptions**](LineItemOptionsApi.md#GETLineItemIdLineItemOptions) | **Get** /line_items/{lineItemId}/line_item_options | Retrieve the line item options associated to the line item
[**GETLineItemOptions**](LineItemOptionsApi.md#GETLineItemOptions) | **Get** /line_item_options | List all line item options
[**GETLineItemOptionsLineItemOptionId**](LineItemOptionsApi.md#GETLineItemOptionsLineItemOptionId) | **Get** /line_item_options/{lineItemOptionId} | Retrieve a line item option
[**GETOrderIdLineItemOptions**](LineItemOptionsApi.md#GETOrderIdLineItemOptions) | **Get** /orders/{orderId}/line_item_options | Retrieve the line item options associated to the order
[**PATCHLineItemOptionsLineItemOptionId**](LineItemOptionsApi.md#PATCHLineItemOptionsLineItemOptionId) | **Patch** /line_item_options/{lineItemOptionId} | Update a line item option
[**POSTLineItemOptions**](LineItemOptionsApi.md#POSTLineItemOptions) | **Post** /line_item_options | Create a line item option



## DELETELineItemOptionsLineItemOptionId

> DELETELineItemOptionsLineItemOptionId(ctx, lineItemOptionId).Execute()

Delete a line item option



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
    lineItemOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LineItemOptionsApi.DELETELineItemOptionsLineItemOptionId(context.Background(), lineItemOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.DELETELineItemOptionsLineItemOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETELineItemOptionsLineItemOptionIdRequest struct via the builder pattern


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


## GETLineItemIdLineItemOptions

> GETLineItemIdLineItemOptions(ctx, lineItemId).Execute()

Retrieve the line item options associated to the line item



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
    lineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LineItemOptionsApi.GETLineItemIdLineItemOptions(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.GETLineItemIdLineItemOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemIdLineItemOptionsRequest struct via the builder pattern


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


## GETLineItemOptions

> GETLineItemOptions200Response GETLineItemOptions(ctx).Execute()

List all line item options



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
    resp, r, err := apiClient.LineItemOptionsApi.GETLineItemOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.GETLineItemOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLineItemOptions`: GETLineItemOptions200Response
    fmt.Fprintf(os.Stdout, "Response from `LineItemOptionsApi.GETLineItemOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemOptionsRequest struct via the builder pattern


### Return type

[**GETLineItemOptions200Response**](GETLineItemOptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETLineItemOptionsLineItemOptionId

> GETLineItemOptionsLineItemOptionId200Response GETLineItemOptionsLineItemOptionId(ctx, lineItemOptionId).Execute()

Retrieve a line item option



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
    lineItemOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemOptionsApi.GETLineItemOptionsLineItemOptionId(context.Background(), lineItemOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.GETLineItemOptionsLineItemOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLineItemOptionsLineItemOptionId`: GETLineItemOptionsLineItemOptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `LineItemOptionsApi.GETLineItemOptionsLineItemOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemOptionsLineItemOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETLineItemOptionsLineItemOptionId200Response**](GETLineItemOptionsLineItemOptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdLineItemOptions

> GETOrderIdLineItemOptions(ctx, orderId).Execute()

Retrieve the line item options associated to the order



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
    orderId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LineItemOptionsApi.GETOrderIdLineItemOptions(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.GETOrderIdLineItemOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderIdLineItemOptionsRequest struct via the builder pattern


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


## PATCHLineItemOptionsLineItemOptionId

> PATCHLineItemOptionsLineItemOptionId200Response PATCHLineItemOptionsLineItemOptionId(ctx, lineItemOptionId).LineItemOptionUpdate(lineItemOptionUpdate).Execute()

Update a line item option



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
    lineItemOptionUpdate := *openapiclient.NewLineItemOptionUpdate(*openapiclient.NewLineItemOptionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes())) // LineItemOptionUpdate | 
    lineItemOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemOptionsApi.PATCHLineItemOptionsLineItemOptionId(context.Background(), lineItemOptionId).LineItemOptionUpdate(lineItemOptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.PATCHLineItemOptionsLineItemOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHLineItemOptionsLineItemOptionId`: PATCHLineItemOptionsLineItemOptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `LineItemOptionsApi.PATCHLineItemOptionsLineItemOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHLineItemOptionsLineItemOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lineItemOptionUpdate** | [**LineItemOptionUpdate**](LineItemOptionUpdate.md) |  | 


### Return type

[**PATCHLineItemOptionsLineItemOptionId200Response**](PATCHLineItemOptionsLineItemOptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTLineItemOptions

> POSTLineItemOptions201Response POSTLineItemOptions(ctx).LineItemOptionCreate(lineItemOptionCreate).Execute()

Create a line item option



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
    lineItemOptionCreate := *openapiclient.NewLineItemOptionCreate(*openapiclient.NewLineItemOptionCreateData(interface{}(123), *openapiclient.NewPOSTLineItemOptions201ResponseDataAttributes(interface{}(2), interface{}({embossing_text=Happy Birthday!})))) // LineItemOptionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemOptionsApi.POSTLineItemOptions(context.Background()).LineItemOptionCreate(lineItemOptionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemOptionsApi.POSTLineItemOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTLineItemOptions`: POSTLineItemOptions201Response
    fmt.Fprintf(os.Stdout, "Response from `LineItemOptionsApi.POSTLineItemOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTLineItemOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lineItemOptionCreate** | [**LineItemOptionCreate**](LineItemOptionCreate.md) |  | 

### Return type

[**POSTLineItemOptions201Response**](POSTLineItemOptions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

