# \PackagesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPackagesPackageId**](PackagesApi.md#DELETEPackagesPackageId) | **Delete** /packages/{packageId} | Delete a package
[**GETPackages**](PackagesApi.md#GETPackages) | **Get** /packages | List all packages
[**GETPackagesPackageId**](PackagesApi.md#GETPackagesPackageId) | **Get** /packages/{packageId} | Retrieve a package
[**GETParcelIdPackage**](PackagesApi.md#GETParcelIdPackage) | **Get** /parcels/{parcelId}/package | Retrieve the package associated to the parcel
[**PATCHPackagesPackageId**](PackagesApi.md#PATCHPackagesPackageId) | **Patch** /packages/{packageId} | Update a package
[**POSTPackages**](PackagesApi.md#POSTPackages) | **Post** /packages | Create a package



## DELETEPackagesPackageId

> DELETEPackagesPackageId(ctx, packageId).Execute()

Delete a package



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
    r, err := apiClient.PackagesApi.DELETEPackagesPackageId(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.DELETEPackagesPackageId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEPackagesPackageIdRequest struct via the builder pattern


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


## GETPackages

> GETPackages200Response GETPackages(ctx).Execute()

List all packages



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
    resp, r, err := apiClient.PackagesApi.GETPackages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.GETPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPackages`: GETPackages200Response
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.GETPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPackagesRequest struct via the builder pattern


### Return type

[**GETPackages200Response**](GETPackages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPackagesPackageId

> GETPackagesPackageId200Response GETPackagesPackageId(ctx, packageId).Execute()

Retrieve a package



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
    resp, r, err := apiClient.PackagesApi.GETPackagesPackageId(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.GETPackagesPackageId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPackagesPackageId`: GETPackagesPackageId200Response
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.GETPackagesPackageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPackagesPackageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPackagesPackageId200Response**](GETPackagesPackageId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETParcelIdPackage

> GETParcelIdPackage(ctx, parcelId).Execute()

Retrieve the package associated to the parcel



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
    r, err := apiClient.PackagesApi.GETParcelIdPackage(context.Background(), parcelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.GETParcelIdPackage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETParcelIdPackageRequest struct via the builder pattern


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


## PATCHPackagesPackageId

> PATCHPackagesPackageId200Response PATCHPackagesPackageId(ctx, packageId).PackageUpdate(packageUpdate).Execute()

Update a package



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
    packageUpdate := *openapiclient.NewPackageUpdate(*openapiclient.NewPackageUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPackagesPackageId200ResponseDataAttributes())) // PackageUpdate | 
    packageId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PATCHPackagesPackageId(context.Background(), packageId).PackageUpdate(packageUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PATCHPackagesPackageId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPackagesPackageId`: PATCHPackagesPackageId200Response
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PATCHPackagesPackageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPackagesPackageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageUpdate** | [**PackageUpdate**](PackageUpdate.md) |  | 


### Return type

[**PATCHPackagesPackageId200Response**](PATCHPackagesPackageId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPackages

> POSTPackages201Response POSTPackages(ctx).PackageCreate(packageCreate).Execute()

Create a package



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
    packageCreate := *openapiclient.NewPackageCreate(*openapiclient.NewPackageCreateData(interface{}(123), *openapiclient.NewPOSTPackages201ResponseDataAttributes(interface{}(Large (60x40x30)), interface{}(40.0), interface{}(40.0), interface{}(25.0), interface{}(gr)))) // PackageCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.POSTPackages(context.Background()).PackageCreate(packageCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.POSTPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPackages`: POSTPackages201Response
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.POSTPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageCreate** | [**PackageCreate**](PackageCreate.md) |  | 

### Return type

[**POSTPackages201Response**](POSTPackages201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

