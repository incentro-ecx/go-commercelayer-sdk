# StoreCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | [**BundleCreateDataRelationshipsMarket**](BundleCreateDataRelationshipsMarket.md) |  | 
**Merchant** | Pointer to [**MarketCreateDataRelationshipsMerchant**](MarketCreateDataRelationshipsMerchant.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeCreateDataRelationshipsStockLocation**](DeliveryLeadTimeCreateDataRelationshipsStockLocation.md) |  | [optional] 

## Methods

### NewStoreCreateDataRelationships

`func NewStoreCreateDataRelationships(market BundleCreateDataRelationshipsMarket, ) *StoreCreateDataRelationships`

NewStoreCreateDataRelationships instantiates a new StoreCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreCreateDataRelationshipsWithDefaults

`func NewStoreCreateDataRelationshipsWithDefaults() *StoreCreateDataRelationships`

NewStoreCreateDataRelationshipsWithDefaults instantiates a new StoreCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *StoreCreateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *StoreCreateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *StoreCreateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket)`

SetMarket sets Market field to given value.


### GetMerchant

`func (o *StoreCreateDataRelationships) GetMerchant() MarketCreateDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *StoreCreateDataRelationships) GetMerchantOk() (*MarketCreateDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *StoreCreateDataRelationships) SetMerchant(v MarketCreateDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *StoreCreateDataRelationships) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetStockLocation

`func (o *StoreCreateDataRelationships) GetStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *StoreCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *StoreCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *StoreCreateDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


