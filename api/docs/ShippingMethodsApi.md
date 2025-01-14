# \ShippingMethodsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEShippingMethodsShippingMethodId**](ShippingMethodsApi.md#DELETEShippingMethodsShippingMethodId) | **Delete** /shipping_methods/{shippingMethodId} | Delete a shipping method
[**GETDeliveryLeadTimeIdShippingMethod**](ShippingMethodsApi.md#GETDeliveryLeadTimeIdShippingMethod) | **Get** /delivery_lead_times/{deliveryLeadTimeId}/shipping_method | Retrieve the shipping method associated to the delivery lead time
[**GETShipmentIdAvailableShippingMethods**](ShippingMethodsApi.md#GETShipmentIdAvailableShippingMethods) | **Get** /shipments/{shipmentId}/available_shipping_methods | Retrieve the available shipping methods associated to the shipment
[**GETShipmentIdShippingMethod**](ShippingMethodsApi.md#GETShipmentIdShippingMethod) | **Get** /shipments/{shipmentId}/shipping_method | Retrieve the shipping method associated to the shipment
[**GETShippingMethodTierIdShippingMethod**](ShippingMethodsApi.md#GETShippingMethodTierIdShippingMethod) | **Get** /shipping_method_tiers/{shippingMethodTierId}/shipping_method | Retrieve the shipping method associated to the shipping method tier
[**GETShippingMethods**](ShippingMethodsApi.md#GETShippingMethods) | **Get** /shipping_methods | List all shipping methods
[**GETShippingMethodsShippingMethodId**](ShippingMethodsApi.md#GETShippingMethodsShippingMethodId) | **Get** /shipping_methods/{shippingMethodId} | Retrieve a shipping method
[**GETShippingWeightTierIdShippingMethod**](ShippingMethodsApi.md#GETShippingWeightTierIdShippingMethod) | **Get** /shipping_weight_tiers/{shippingWeightTierId}/shipping_method | Retrieve the shipping method associated to the shipping weight tier
[**PATCHShippingMethodsShippingMethodId**](ShippingMethodsApi.md#PATCHShippingMethodsShippingMethodId) | **Patch** /shipping_methods/{shippingMethodId} | Update a shipping method
[**POSTShippingMethods**](ShippingMethodsApi.md#POSTShippingMethods) | **Post** /shipping_methods | Create a shipping method



## DELETEShippingMethodsShippingMethodId

> DELETEShippingMethodsShippingMethodId(ctx, shippingMethodId).Execute()

Delete a shipping method



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
    r, err := apiClient.ShippingMethodsApi.DELETEShippingMethodsShippingMethodId(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.DELETEShippingMethodsShippingMethodId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEShippingMethodsShippingMethodIdRequest struct via the builder pattern


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


## GETDeliveryLeadTimeIdShippingMethod

> GETDeliveryLeadTimeIdShippingMethod(ctx, deliveryLeadTimeId).Execute()

Retrieve the shipping method associated to the delivery lead time



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
    r, err := apiClient.ShippingMethodsApi.GETDeliveryLeadTimeIdShippingMethod(context.Background(), deliveryLeadTimeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.GETDeliveryLeadTimeIdShippingMethod``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETDeliveryLeadTimeIdShippingMethodRequest struct via the builder pattern


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


## GETShipmentIdAvailableShippingMethods

> GETShipmentIdAvailableShippingMethods(ctx, shipmentId).Execute()

Retrieve the available shipping methods associated to the shipment



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
    r, err := apiClient.ShippingMethodsApi.GETShipmentIdAvailableShippingMethods(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.GETShipmentIdAvailableShippingMethods``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShipmentIdAvailableShippingMethodsRequest struct via the builder pattern


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


## GETShipmentIdShippingMethod

> GETShipmentIdShippingMethod(ctx, shipmentId).Execute()

Retrieve the shipping method associated to the shipment



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
    r, err := apiClient.ShippingMethodsApi.GETShipmentIdShippingMethod(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.GETShipmentIdShippingMethod``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShipmentIdShippingMethodRequest struct via the builder pattern


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


## GETShippingMethodTierIdShippingMethod

> GETShippingMethodTierIdShippingMethod(ctx, shippingMethodTierId).Execute()

Retrieve the shipping method associated to the shipping method tier



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
    shippingMethodTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShippingMethodsApi.GETShippingMethodTierIdShippingMethod(context.Background(), shippingMethodTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.GETShippingMethodTierIdShippingMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodTierIdShippingMethodRequest struct via the builder pattern


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


## GETShippingMethods

> GETShippingMethods200Response GETShippingMethods(ctx).Execute()

List all shipping methods



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
    resp, r, err := apiClient.ShippingMethodsApi.GETShippingMethods(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.GETShippingMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingMethods`: GETShippingMethods200Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingMethodsApi.GETShippingMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodsRequest struct via the builder pattern


### Return type

[**GETShippingMethods200Response**](GETShippingMethods200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShippingMethodsShippingMethodId

> GETShippingMethodsShippingMethodId200Response GETShippingMethodsShippingMethodId(ctx, shippingMethodId).Execute()

Retrieve a shipping method



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
    resp, r, err := apiClient.ShippingMethodsApi.GETShippingMethodsShippingMethodId(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.GETShippingMethodsShippingMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingMethodsShippingMethodId`: GETShippingMethodsShippingMethodId200Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingMethodsApi.GETShippingMethodsShippingMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodsShippingMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETShippingMethodsShippingMethodId200Response**](GETShippingMethodsShippingMethodId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShippingWeightTierIdShippingMethod

> GETShippingWeightTierIdShippingMethod(ctx, shippingWeightTierId).Execute()

Retrieve the shipping method associated to the shipping weight tier



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
    shippingWeightTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShippingMethodsApi.GETShippingWeightTierIdShippingMethod(context.Background(), shippingWeightTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.GETShippingWeightTierIdShippingMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingWeightTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingWeightTierIdShippingMethodRequest struct via the builder pattern


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


## PATCHShippingMethodsShippingMethodId

> PATCHShippingMethodsShippingMethodId200Response PATCHShippingMethodsShippingMethodId(ctx, shippingMethodId).ShippingMethodUpdate(shippingMethodUpdate).Execute()

Update a shipping method



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
    shippingMethodUpdate := *openapiclient.NewShippingMethodUpdate(*openapiclient.NewShippingMethodUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHShippingMethodsShippingMethodId200ResponseDataAttributes())) // ShippingMethodUpdate | 
    shippingMethodId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingMethodsApi.PATCHShippingMethodsShippingMethodId(context.Background(), shippingMethodId).ShippingMethodUpdate(shippingMethodUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.PATCHShippingMethodsShippingMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHShippingMethodsShippingMethodId`: PATCHShippingMethodsShippingMethodId200Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingMethodsApi.PATCHShippingMethodsShippingMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHShippingMethodsShippingMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingMethodUpdate** | [**ShippingMethodUpdate**](ShippingMethodUpdate.md) |  | 


### Return type

[**PATCHShippingMethodsShippingMethodId200Response**](PATCHShippingMethodsShippingMethodId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTShippingMethods

> POSTShippingMethods201Response POSTShippingMethods(ctx).ShippingMethodCreate(shippingMethodCreate).Execute()

Create a shipping method



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
    shippingMethodCreate := *openapiclient.NewShippingMethodCreate(*openapiclient.NewShippingMethodCreateData(interface{}(123), *openapiclient.NewPOSTShippingMethods201ResponseDataAttributes(interface{}(Standard shipping), interface{}(1000)))) // ShippingMethodCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingMethodsApi.POSTShippingMethods(context.Background()).ShippingMethodCreate(shippingMethodCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsApi.POSTShippingMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTShippingMethods`: POSTShippingMethods201Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingMethodsApi.POSTShippingMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTShippingMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingMethodCreate** | [**ShippingMethodCreate**](ShippingMethodCreate.md) |  | 

### Return type

[**POSTShippingMethods201Response**](POSTShippingMethods201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

