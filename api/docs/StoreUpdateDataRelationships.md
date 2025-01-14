# StoreUpdateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BundleCreateDataRelationshipsMarket**](BundleCreateDataRelationshipsMarket.md) |  | [optional] 
**Merchant** | Pointer to [**MarketCreateDataRelationshipsMerchant**](MarketCreateDataRelationshipsMerchant.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeCreateDataRelationshipsStockLocation**](DeliveryLeadTimeCreateDataRelationshipsStockLocation.md) |  | [optional] 

## Methods

### NewStoreUpdateDataRelationships

`func NewStoreUpdateDataRelationships() *StoreUpdateDataRelationships`

NewStoreUpdateDataRelationships instantiates a new StoreUpdateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreUpdateDataRelationshipsWithDefaults

`func NewStoreUpdateDataRelationshipsWithDefaults() *StoreUpdateDataRelationships`

NewStoreUpdateDataRelationshipsWithDefaults instantiates a new StoreUpdateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *StoreUpdateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *StoreUpdateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *StoreUpdateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *StoreUpdateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetMerchant

`func (o *StoreUpdateDataRelationships) GetMerchant() MarketCreateDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *StoreUpdateDataRelationships) GetMerchantOk() (*MarketCreateDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *StoreUpdateDataRelationships) SetMerchant(v MarketCreateDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *StoreUpdateDataRelationships) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetStockLocation

`func (o *StoreUpdateDataRelationships) GetStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *StoreUpdateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *StoreUpdateDataRelationships) SetStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *StoreUpdateDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


