# \NotificationsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETENotificationsNotificationId**](NotificationsApi.md#DELETENotificationsNotificationId) | **Delete** /notifications/{notificationId} | Delete a notification
[**GETLineItemIdNotifications**](NotificationsApi.md#GETLineItemIdNotifications) | **Get** /line_items/{lineItemId}/notifications | Retrieve the notifications associated to the line item
[**GETNotifications**](NotificationsApi.md#GETNotifications) | **Get** /notifications | List all notifications
[**GETNotificationsNotificationId**](NotificationsApi.md#GETNotificationsNotificationId) | **Get** /notifications/{notificationId} | Retrieve a notification
[**GETOrderIdNotifications**](NotificationsApi.md#GETOrderIdNotifications) | **Get** /orders/{orderId}/notifications | Retrieve the notifications associated to the order
[**PATCHNotificationsNotificationId**](NotificationsApi.md#PATCHNotificationsNotificationId) | **Patch** /notifications/{notificationId} | Update a notification
[**POSTNotifications**](NotificationsApi.md#POSTNotifications) | **Post** /notifications | Create a notification



## DELETENotificationsNotificationId

> DELETENotificationsNotificationId(ctx, notificationId).Execute()

Delete a notification



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
    notificationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsApi.DELETENotificationsNotificationId(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DELETENotificationsNotificationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETENotificationsNotificationIdRequest struct via the builder pattern


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


## GETLineItemIdNotifications

> GETLineItemIdNotifications(ctx, lineItemId).Execute()

Retrieve the notifications associated to the line item



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
    r, err := apiClient.NotificationsApi.GETLineItemIdNotifications(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GETLineItemIdNotifications``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETLineItemIdNotificationsRequest struct via the builder pattern


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


## GETNotifications

> GETNotifications200Response GETNotifications(ctx).Execute()

List all notifications



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
    resp, r, err := apiClient.NotificationsApi.GETNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GETNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETNotifications`: GETNotifications200Response
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GETNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETNotificationsRequest struct via the builder pattern


### Return type

[**GETNotifications200Response**](GETNotifications200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETNotificationsNotificationId

> GETNotificationsNotificationId200Response GETNotificationsNotificationId(ctx, notificationId).Execute()

Retrieve a notification



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
    notificationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GETNotificationsNotificationId(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GETNotificationsNotificationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETNotificationsNotificationId`: GETNotificationsNotificationId200Response
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GETNotificationsNotificationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETNotificationsNotificationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETNotificationsNotificationId200Response**](GETNotificationsNotificationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdNotifications

> GETOrderIdNotifications(ctx, orderId).Execute()

Retrieve the notifications associated to the order



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
    r, err := apiClient.NotificationsApi.GETOrderIdNotifications(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GETOrderIdNotifications``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdNotificationsRequest struct via the builder pattern


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


## PATCHNotificationsNotificationId

> PATCHNotificationsNotificationId200Response PATCHNotificationsNotificationId(ctx, notificationId).NotificationUpdate(notificationUpdate).Execute()

Update a notification



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
    notificationUpdate := *openapiclient.NewNotificationUpdate(*openapiclient.NewNotificationUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHNotificationsNotificationId200ResponseDataAttributes())) // NotificationUpdate | 
    notificationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.PATCHNotificationsNotificationId(context.Background(), notificationId).NotificationUpdate(notificationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.PATCHNotificationsNotificationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHNotificationsNotificationId`: PATCHNotificationsNotificationId200Response
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.PATCHNotificationsNotificationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHNotificationsNotificationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationUpdate** | [**NotificationUpdate**](NotificationUpdate.md) |  | 


### Return type

[**PATCHNotificationsNotificationId200Response**](PATCHNotificationsNotificationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTNotifications

> POSTNotifications201Response POSTNotifications(ctx).NotificationCreate(notificationCreate).Execute()

Create a notification



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
    notificationCreate := *openapiclient.NewNotificationCreate(*openapiclient.NewNotificationCreateData(interface{}(123), *openapiclient.NewPOSTNotifications201ResponseDataAttributes(interface{}(DDT transport document)))) // NotificationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.POSTNotifications(context.Background()).NotificationCreate(notificationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.POSTNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTNotifications`: POSTNotifications201Response
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.POSTNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationCreate** | [**NotificationCreate**](NotificationCreate.md) |  | 

### Return type

[**POSTNotifications201Response**](POSTNotifications201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

