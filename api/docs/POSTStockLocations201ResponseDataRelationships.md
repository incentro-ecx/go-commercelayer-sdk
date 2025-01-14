# POSTStockLocations201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**POSTCustomerAddresses201ResponseDataRelationshipsAddress**](POSTCustomerAddresses201ResponseDataRelationshipsAddress.md) |  | [optional] 
**InventoryStockLocations** | Pointer to [**POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations**](POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations.md) |  | [optional] 
**InventoryReturnLocations** | Pointer to [**POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations**](POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations.md) |  | [optional] 
**StockItems** | Pointer to [**POSTSkus201ResponseDataRelationshipsStockItems**](POSTSkus201ResponseDataRelationshipsStockItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockTransfers**](POSTLineItems201ResponseDataRelationshipsStockTransfers.md) |  | [optional] 
**Stores** | Pointer to [**POSTMarkets201ResponseDataRelationshipsStores**](POSTMarkets201ResponseDataRelationshipsStores.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTStockLocations201ResponseDataRelationships

`func NewPOSTStockLocations201ResponseDataRelationships() *POSTStockLocations201ResponseDataRelationships`

NewPOSTStockLocations201ResponseDataRelationships instantiates a new POSTStockLocations201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockLocations201ResponseDataRelationshipsWithDefaults

`func NewPOSTStockLocations201ResponseDataRelationshipsWithDefaults() *POSTStockLocations201ResponseDataRelationships`

NewPOSTStockLocations201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockLocations201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *POSTStockLocations201ResponseDataRelationships) GetAddress() POSTCustomerAddresses201ResponseDataRelationshipsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *POSTStockLocations201ResponseDataRelationships) GetAddressOk() (*POSTCustomerAddresses201ResponseDataRelationshipsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *POSTStockLocations201ResponseDataRelationships) SetAddress(v POSTCustomerAddresses201ResponseDataRelationshipsAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *POSTStockLocations201ResponseDataRelationships) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInventoryStockLocations

`func (o *POSTStockLocations201ResponseDataRelationships) GetInventoryStockLocations() POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations`

GetInventoryStockLocations returns the InventoryStockLocations field if non-nil, zero value otherwise.

### GetInventoryStockLocationsOk

`func (o *POSTStockLocations201ResponseDataRelationships) GetInventoryStockLocationsOk() (*POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations, bool)`

GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryStockLocations

`func (o *POSTStockLocations201ResponseDataRelationships) SetInventoryStockLocations(v POSTInventoryModels201ResponseDataRelationshipsInventoryStockLocations)`

SetInventoryStockLocations sets InventoryStockLocations field to given value.

### HasInventoryStockLocations

`func (o *POSTStockLocations201ResponseDataRelationships) HasInventoryStockLocations() bool`

HasInventoryStockLocations returns a boolean if a field has been set.

### GetInventoryReturnLocations

`func (o *POSTStockLocations201ResponseDataRelationships) GetInventoryReturnLocations() POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations`

GetInventoryReturnLocations returns the InventoryReturnLocations field if non-nil, zero value otherwise.

### GetInventoryReturnLocationsOk

`func (o *POSTStockLocations201ResponseDataRelationships) GetInventoryReturnLocationsOk() (*POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations, bool)`

GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReturnLocations

`func (o *POSTStockLocations201ResponseDataRelationships) SetInventoryReturnLocations(v POSTInventoryModels201ResponseDataRelationshipsInventoryReturnLocations)`

SetInventoryReturnLocations sets InventoryReturnLocations field to given value.

### HasInventoryReturnLocations

`func (o *POSTStockLocations201ResponseDataRelationships) HasInventoryReturnLocations() bool`

HasInventoryReturnLocations returns a boolean if a field has been set.

### GetStockItems

`func (o *POSTStockLocations201ResponseDataRelationships) GetStockItems() POSTSkus201ResponseDataRelationshipsStockItems`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *POSTStockLocations201ResponseDataRelationships) GetStockItemsOk() (*POSTSkus201ResponseDataRelationshipsStockItems, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *POSTStockLocations201ResponseDataRelationships) SetStockItems(v POSTSkus201ResponseDataRelationshipsStockItems)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *POSTStockLocations201ResponseDataRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *POSTStockLocations201ResponseDataRelationships) GetStockTransfers() POSTLineItems201ResponseDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *POSTStockLocations201ResponseDataRelationships) GetStockTransfersOk() (*POSTLineItems201ResponseDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *POSTStockLocations201ResponseDataRelationships) SetStockTransfers(v POSTLineItems201ResponseDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *POSTStockLocations201ResponseDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetStores

`func (o *POSTStockLocations201ResponseDataRelationships) GetStores() POSTMarkets201ResponseDataRelationshipsStores`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *POSTStockLocations201ResponseDataRelationships) GetStoresOk() (*POSTMarkets201ResponseDataRelationshipsStores, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *POSTStockLocations201ResponseDataRelationships) SetStores(v POSTMarkets201ResponseDataRelationshipsStores)`

SetStores sets Stores field to given value.

### HasStores

`func (o *POSTStockLocations201ResponseDataRelationships) HasStores() bool`

HasStores returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTStockLocations201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTStockLocations201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTStockLocations201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTStockLocations201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *POSTStockLocations201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTStockLocations201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTStockLocations201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTStockLocations201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


