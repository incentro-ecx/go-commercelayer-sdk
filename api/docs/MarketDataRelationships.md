# MarketDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | Pointer to [**MarketDataRelationshipsMerchant**](MarketDataRelationshipsMerchant.md) |  | [optional] 
**PriceList** | Pointer to [**MarketDataRelationshipsPriceList**](MarketDataRelationshipsPriceList.md) |  | [optional] 
**BasePriceList** | Pointer to [**MarketDataRelationshipsPriceList**](MarketDataRelationshipsPriceList.md) |  | [optional] 
**InventoryModel** | Pointer to [**InventoryReturnLocationDataRelationshipsInventoryModel**](InventoryReturnLocationDataRelationshipsInventoryModel.md) |  | [optional] 
**SubscriptionModel** | Pointer to [**MarketDataRelationshipsSubscriptionModel**](MarketDataRelationshipsSubscriptionModel.md) |  | [optional] 
**TaxCalculator** | Pointer to [**MarketDataRelationshipsTaxCalculator**](MarketDataRelationshipsTaxCalculator.md) |  | [optional] 
**CustomerGroup** | Pointer to [**CustomerDataRelationshipsCustomerGroup**](CustomerDataRelationshipsCustomerGroup.md) |  | [optional] 
**Geocoder** | Pointer to [**AddressDataRelationshipsGeocoder**](AddressDataRelationshipsGeocoder.md) |  | [optional] 
**Stores** | Pointer to [**MarketDataRelationshipsStores**](MarketDataRelationshipsStores.md) |  | [optional] 
**PriceListSchedulers** | Pointer to [**MarketDataRelationshipsPriceListSchedulers**](MarketDataRelationshipsPriceListSchedulers.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewMarketDataRelationships

`func NewMarketDataRelationships() *MarketDataRelationships`

NewMarketDataRelationships instantiates a new MarketDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketDataRelationshipsWithDefaults

`func NewMarketDataRelationshipsWithDefaults() *MarketDataRelationships`

NewMarketDataRelationshipsWithDefaults instantiates a new MarketDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *MarketDataRelationships) GetMerchant() MarketDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *MarketDataRelationships) GetMerchantOk() (*MarketDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *MarketDataRelationships) SetMerchant(v MarketDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *MarketDataRelationships) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPriceList

`func (o *MarketDataRelationships) GetPriceList() MarketDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *MarketDataRelationships) GetPriceListOk() (*MarketDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *MarketDataRelationships) SetPriceList(v MarketDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.

### HasPriceList

`func (o *MarketDataRelationships) HasPriceList() bool`

HasPriceList returns a boolean if a field has been set.

### GetBasePriceList

`func (o *MarketDataRelationships) GetBasePriceList() MarketDataRelationshipsPriceList`

GetBasePriceList returns the BasePriceList field if non-nil, zero value otherwise.

### GetBasePriceListOk

`func (o *MarketDataRelationships) GetBasePriceListOk() (*MarketDataRelationshipsPriceList, bool)`

GetBasePriceListOk returns a tuple with the BasePriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePriceList

`func (o *MarketDataRelationships) SetBasePriceList(v MarketDataRelationshipsPriceList)`

SetBasePriceList sets BasePriceList field to given value.

### HasBasePriceList

`func (o *MarketDataRelationships) HasBasePriceList() bool`

HasBasePriceList returns a boolean if a field has been set.

### GetInventoryModel

`func (o *MarketDataRelationships) GetInventoryModel() InventoryReturnLocationDataRelationshipsInventoryModel`

GetInventoryModel returns the InventoryModel field if non-nil, zero value otherwise.

### GetInventoryModelOk

`func (o *MarketDataRelationships) GetInventoryModelOk() (*InventoryReturnLocationDataRelationshipsInventoryModel, bool)`

GetInventoryModelOk returns a tuple with the InventoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryModel

`func (o *MarketDataRelationships) SetInventoryModel(v InventoryReturnLocationDataRelationshipsInventoryModel)`

SetInventoryModel sets InventoryModel field to given value.

### HasInventoryModel

`func (o *MarketDataRelationships) HasInventoryModel() bool`

HasInventoryModel returns a boolean if a field has been set.

### GetSubscriptionModel

`func (o *MarketDataRelationships) GetSubscriptionModel() MarketDataRelationshipsSubscriptionModel`

GetSubscriptionModel returns the SubscriptionModel field if non-nil, zero value otherwise.

### GetSubscriptionModelOk

`func (o *MarketDataRelationships) GetSubscriptionModelOk() (*MarketDataRelationshipsSubscriptionModel, bool)`

GetSubscriptionModelOk returns a tuple with the SubscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionModel

`func (o *MarketDataRelationships) SetSubscriptionModel(v MarketDataRelationshipsSubscriptionModel)`

SetSubscriptionModel sets SubscriptionModel field to given value.

### HasSubscriptionModel

`func (o *MarketDataRelationships) HasSubscriptionModel() bool`

HasSubscriptionModel returns a boolean if a field has been set.

### GetTaxCalculator

`func (o *MarketDataRelationships) GetTaxCalculator() MarketDataRelationshipsTaxCalculator`

GetTaxCalculator returns the TaxCalculator field if non-nil, zero value otherwise.

### GetTaxCalculatorOk

`func (o *MarketDataRelationships) GetTaxCalculatorOk() (*MarketDataRelationshipsTaxCalculator, bool)`

GetTaxCalculatorOk returns a tuple with the TaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculator

`func (o *MarketDataRelationships) SetTaxCalculator(v MarketDataRelationshipsTaxCalculator)`

SetTaxCalculator sets TaxCalculator field to given value.

### HasTaxCalculator

`func (o *MarketDataRelationships) HasTaxCalculator() bool`

HasTaxCalculator returns a boolean if a field has been set.

### GetCustomerGroup

`func (o *MarketDataRelationships) GetCustomerGroup() CustomerDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *MarketDataRelationships) GetCustomerGroupOk() (*CustomerDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *MarketDataRelationships) SetCustomerGroup(v CustomerDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *MarketDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.

### GetGeocoder

`func (o *MarketDataRelationships) GetGeocoder() AddressDataRelationshipsGeocoder`

GetGeocoder returns the Geocoder field if non-nil, zero value otherwise.

### GetGeocoderOk

`func (o *MarketDataRelationships) GetGeocoderOk() (*AddressDataRelationshipsGeocoder, bool)`

GetGeocoderOk returns a tuple with the Geocoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoder

`func (o *MarketDataRelationships) SetGeocoder(v AddressDataRelationshipsGeocoder)`

SetGeocoder sets Geocoder field to given value.

### HasGeocoder

`func (o *MarketDataRelationships) HasGeocoder() bool`

HasGeocoder returns a boolean if a field has been set.

### GetStores

`func (o *MarketDataRelationships) GetStores() MarketDataRelationshipsStores`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *MarketDataRelationships) GetStoresOk() (*MarketDataRelationshipsStores, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *MarketDataRelationships) SetStores(v MarketDataRelationshipsStores)`

SetStores sets Stores field to given value.

### HasStores

`func (o *MarketDataRelationships) HasStores() bool`

HasStores returns a boolean if a field has been set.

### GetPriceListSchedulers

`func (o *MarketDataRelationships) GetPriceListSchedulers() MarketDataRelationshipsPriceListSchedulers`

GetPriceListSchedulers returns the PriceListSchedulers field if non-nil, zero value otherwise.

### GetPriceListSchedulersOk

`func (o *MarketDataRelationships) GetPriceListSchedulersOk() (*MarketDataRelationshipsPriceListSchedulers, bool)`

GetPriceListSchedulersOk returns a tuple with the PriceListSchedulers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListSchedulers

`func (o *MarketDataRelationships) SetPriceListSchedulers(v MarketDataRelationshipsPriceListSchedulers)`

SetPriceListSchedulers sets PriceListSchedulers field to given value.

### HasPriceListSchedulers

`func (o *MarketDataRelationships) HasPriceListSchedulers() bool`

HasPriceListSchedulers returns a boolean if a field has been set.

### GetAttachments

`func (o *MarketDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MarketDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MarketDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MarketDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *MarketDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MarketDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MarketDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *MarketDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


