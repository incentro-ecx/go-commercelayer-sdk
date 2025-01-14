# StockLocationDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**InventoryStockLocations** | Pointer to [**InventoryModelDataRelationshipsInventoryStockLocations**](InventoryModelDataRelationshipsInventoryStockLocations.md) |  | [optional] 
**InventoryReturnLocations** | Pointer to [**InventoryModelDataRelationshipsInventoryReturnLocations**](InventoryModelDataRelationshipsInventoryReturnLocations.md) |  | [optional] 
**StockItems** | Pointer to [**ReservedStockDataRelationshipsStockItem**](ReservedStockDataRelationshipsStockItem.md) |  | [optional] 
**StockTransfers** | Pointer to [**LineItemDataRelationshipsStockTransfers**](LineItemDataRelationshipsStockTransfers.md) |  | [optional] 
**Stores** | Pointer to [**MarketDataRelationshipsStores**](MarketDataRelationshipsStores.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewStockLocationDataRelationships

`func NewStockLocationDataRelationships() *StockLocationDataRelationships`

NewStockLocationDataRelationships instantiates a new StockLocationDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLocationDataRelationshipsWithDefaults

`func NewStockLocationDataRelationshipsWithDefaults() *StockLocationDataRelationships`

NewStockLocationDataRelationshipsWithDefaults instantiates a new StockLocationDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *StockLocationDataRelationships) GetAddress() BingGeocoderDataRelationshipsAddresses`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StockLocationDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StockLocationDataRelationships) SetAddress(v BingGeocoderDataRelationshipsAddresses)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StockLocationDataRelationships) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInventoryStockLocations

`func (o *StockLocationDataRelationships) GetInventoryStockLocations() InventoryModelDataRelationshipsInventoryStockLocations`

GetInventoryStockLocations returns the InventoryStockLocations field if non-nil, zero value otherwise.

### GetInventoryStockLocationsOk

`func (o *StockLocationDataRelationships) GetInventoryStockLocationsOk() (*InventoryModelDataRelationshipsInventoryStockLocations, bool)`

GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryStockLocations

`func (o *StockLocationDataRelationships) SetInventoryStockLocations(v InventoryModelDataRelationshipsInventoryStockLocations)`

SetInventoryStockLocations sets InventoryStockLocations field to given value.

### HasInventoryStockLocations

`func (o *StockLocationDataRelationships) HasInventoryStockLocations() bool`

HasInventoryStockLocations returns a boolean if a field has been set.

### GetInventoryReturnLocations

`func (o *StockLocationDataRelationships) GetInventoryReturnLocations() InventoryModelDataRelationshipsInventoryReturnLocations`

GetInventoryReturnLocations returns the InventoryReturnLocations field if non-nil, zero value otherwise.

### GetInventoryReturnLocationsOk

`func (o *StockLocationDataRelationships) GetInventoryReturnLocationsOk() (*InventoryModelDataRelationshipsInventoryReturnLocations, bool)`

GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReturnLocations

`func (o *StockLocationDataRelationships) SetInventoryReturnLocations(v InventoryModelDataRelationshipsInventoryReturnLocations)`

SetInventoryReturnLocations sets InventoryReturnLocations field to given value.

### HasInventoryReturnLocations

`func (o *StockLocationDataRelationships) HasInventoryReturnLocations() bool`

HasInventoryReturnLocations returns a boolean if a field has been set.

### GetStockItems

`func (o *StockLocationDataRelationships) GetStockItems() ReservedStockDataRelationshipsStockItem`

GetStockItems returns the StockItems field if non-nil, zero value otherwise.

### GetStockItemsOk

`func (o *StockLocationDataRelationships) GetStockItemsOk() (*ReservedStockDataRelationshipsStockItem, bool)`

GetStockItemsOk returns a tuple with the StockItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockItems

`func (o *StockLocationDataRelationships) SetStockItems(v ReservedStockDataRelationshipsStockItem)`

SetStockItems sets StockItems field to given value.

### HasStockItems

`func (o *StockLocationDataRelationships) HasStockItems() bool`

HasStockItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *StockLocationDataRelationships) GetStockTransfers() LineItemDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *StockLocationDataRelationships) GetStockTransfersOk() (*LineItemDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *StockLocationDataRelationships) SetStockTransfers(v LineItemDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *StockLocationDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetStores

`func (o *StockLocationDataRelationships) GetStores() MarketDataRelationshipsStores`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *StockLocationDataRelationships) GetStoresOk() (*MarketDataRelationshipsStores, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *StockLocationDataRelationships) SetStores(v MarketDataRelationshipsStores)`

SetStores sets Stores field to given value.

### HasStores

`func (o *StockLocationDataRelationships) HasStores() bool`

HasStores returns a boolean if a field has been set.

### GetAttachments

`func (o *StockLocationDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *StockLocationDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *StockLocationDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *StockLocationDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *StockLocationDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *StockLocationDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *StockLocationDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *StockLocationDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


