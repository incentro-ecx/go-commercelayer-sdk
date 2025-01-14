# POSTMarkets201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | Pointer to [**POSTMarkets201ResponseDataRelationshipsMerchant**](POSTMarkets201ResponseDataRelationshipsMerchant.md) |  | [optional] 
**PriceList** | Pointer to [**POSTMarkets201ResponseDataRelationshipsPriceList**](POSTMarkets201ResponseDataRelationshipsPriceList.md) |  | [optional] 
**BasePriceList** | Pointer to [**POSTMarkets201ResponseDataRelationshipsBasePriceList**](POSTMarkets201ResponseDataRelationshipsBasePriceList.md) |  | [optional] 
**InventoryModel** | Pointer to [**POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel**](POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel.md) |  | [optional] 
**SubscriptionModel** | Pointer to [**POSTMarkets201ResponseDataRelationshipsSubscriptionModel**](POSTMarkets201ResponseDataRelationshipsSubscriptionModel.md) |  | [optional] 
**TaxCalculator** | Pointer to [**POSTMarkets201ResponseDataRelationshipsTaxCalculator**](POSTMarkets201ResponseDataRelationshipsTaxCalculator.md) |  | [optional] 
**CustomerGroup** | Pointer to [**POSTCustomers201ResponseDataRelationshipsCustomerGroup**](POSTCustomers201ResponseDataRelationshipsCustomerGroup.md) |  | [optional] 
**Geocoder** | Pointer to [**POSTAddresses201ResponseDataRelationshipsGeocoder**](POSTAddresses201ResponseDataRelationshipsGeocoder.md) |  | [optional] 
**Stores** | Pointer to [**POSTMarkets201ResponseDataRelationshipsStores**](POSTMarkets201ResponseDataRelationshipsStores.md) |  | [optional] 
**PriceListSchedulers** | Pointer to [**POSTMarkets201ResponseDataRelationshipsPriceListSchedulers**](POSTMarkets201ResponseDataRelationshipsPriceListSchedulers.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTMarkets201ResponseDataRelationships

`func NewPOSTMarkets201ResponseDataRelationships() *POSTMarkets201ResponseDataRelationships`

NewPOSTMarkets201ResponseDataRelationships instantiates a new POSTMarkets201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTMarkets201ResponseDataRelationshipsWithDefaults

`func NewPOSTMarkets201ResponseDataRelationshipsWithDefaults() *POSTMarkets201ResponseDataRelationships`

NewPOSTMarkets201ResponseDataRelationshipsWithDefaults instantiates a new POSTMarkets201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *POSTMarkets201ResponseDataRelationships) GetMerchant() POSTMarkets201ResponseDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *POSTMarkets201ResponseDataRelationships) GetMerchantOk() (*POSTMarkets201ResponseDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *POSTMarkets201ResponseDataRelationships) SetMerchant(v POSTMarkets201ResponseDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *POSTMarkets201ResponseDataRelationships) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPriceList

`func (o *POSTMarkets201ResponseDataRelationships) GetPriceList() POSTMarkets201ResponseDataRelationshipsPriceList`

GetPriceList returns the PriceList field if non-nil, zero value otherwise.

### GetPriceListOk

`func (o *POSTMarkets201ResponseDataRelationships) GetPriceListOk() (*POSTMarkets201ResponseDataRelationshipsPriceList, bool)`

GetPriceListOk returns a tuple with the PriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceList

`func (o *POSTMarkets201ResponseDataRelationships) SetPriceList(v POSTMarkets201ResponseDataRelationshipsPriceList)`

SetPriceList sets PriceList field to given value.

### HasPriceList

`func (o *POSTMarkets201ResponseDataRelationships) HasPriceList() bool`

HasPriceList returns a boolean if a field has been set.

### GetBasePriceList

`func (o *POSTMarkets201ResponseDataRelationships) GetBasePriceList() POSTMarkets201ResponseDataRelationshipsBasePriceList`

GetBasePriceList returns the BasePriceList field if non-nil, zero value otherwise.

### GetBasePriceListOk

`func (o *POSTMarkets201ResponseDataRelationships) GetBasePriceListOk() (*POSTMarkets201ResponseDataRelationshipsBasePriceList, bool)`

GetBasePriceListOk returns a tuple with the BasePriceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePriceList

`func (o *POSTMarkets201ResponseDataRelationships) SetBasePriceList(v POSTMarkets201ResponseDataRelationshipsBasePriceList)`

SetBasePriceList sets BasePriceList field to given value.

### HasBasePriceList

`func (o *POSTMarkets201ResponseDataRelationships) HasBasePriceList() bool`

HasBasePriceList returns a boolean if a field has been set.

### GetInventoryModel

`func (o *POSTMarkets201ResponseDataRelationships) GetInventoryModel() POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel`

GetInventoryModel returns the InventoryModel field if non-nil, zero value otherwise.

### GetInventoryModelOk

`func (o *POSTMarkets201ResponseDataRelationships) GetInventoryModelOk() (*POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel, bool)`

GetInventoryModelOk returns a tuple with the InventoryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryModel

`func (o *POSTMarkets201ResponseDataRelationships) SetInventoryModel(v POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel)`

SetInventoryModel sets InventoryModel field to given value.

### HasInventoryModel

`func (o *POSTMarkets201ResponseDataRelationships) HasInventoryModel() bool`

HasInventoryModel returns a boolean if a field has been set.

### GetSubscriptionModel

`func (o *POSTMarkets201ResponseDataRelationships) GetSubscriptionModel() POSTMarkets201ResponseDataRelationshipsSubscriptionModel`

GetSubscriptionModel returns the SubscriptionModel field if non-nil, zero value otherwise.

### GetSubscriptionModelOk

`func (o *POSTMarkets201ResponseDataRelationships) GetSubscriptionModelOk() (*POSTMarkets201ResponseDataRelationshipsSubscriptionModel, bool)`

GetSubscriptionModelOk returns a tuple with the SubscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionModel

`func (o *POSTMarkets201ResponseDataRelationships) SetSubscriptionModel(v POSTMarkets201ResponseDataRelationshipsSubscriptionModel)`

SetSubscriptionModel sets SubscriptionModel field to given value.

### HasSubscriptionModel

`func (o *POSTMarkets201ResponseDataRelationships) HasSubscriptionModel() bool`

HasSubscriptionModel returns a boolean if a field has been set.

### GetTaxCalculator

`func (o *POSTMarkets201ResponseDataRelationships) GetTaxCalculator() POSTMarkets201ResponseDataRelationshipsTaxCalculator`

GetTaxCalculator returns the TaxCalculator field if non-nil, zero value otherwise.

### GetTaxCalculatorOk

`func (o *POSTMarkets201ResponseDataRelationships) GetTaxCalculatorOk() (*POSTMarkets201ResponseDataRelationshipsTaxCalculator, bool)`

GetTaxCalculatorOk returns a tuple with the TaxCalculator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculator

`func (o *POSTMarkets201ResponseDataRelationships) SetTaxCalculator(v POSTMarkets201ResponseDataRelationshipsTaxCalculator)`

SetTaxCalculator sets TaxCalculator field to given value.

### HasTaxCalculator

`func (o *POSTMarkets201ResponseDataRelationships) HasTaxCalculator() bool`

HasTaxCalculator returns a boolean if a field has been set.

### GetCustomerGroup

`func (o *POSTMarkets201ResponseDataRelationships) GetCustomerGroup() POSTCustomers201ResponseDataRelationshipsCustomerGroup`

GetCustomerGroup returns the CustomerGroup field if non-nil, zero value otherwise.

### GetCustomerGroupOk

`func (o *POSTMarkets201ResponseDataRelationships) GetCustomerGroupOk() (*POSTCustomers201ResponseDataRelationshipsCustomerGroup, bool)`

GetCustomerGroupOk returns a tuple with the CustomerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroup

`func (o *POSTMarkets201ResponseDataRelationships) SetCustomerGroup(v POSTCustomers201ResponseDataRelationshipsCustomerGroup)`

SetCustomerGroup sets CustomerGroup field to given value.

### HasCustomerGroup

`func (o *POSTMarkets201ResponseDataRelationships) HasCustomerGroup() bool`

HasCustomerGroup returns a boolean if a field has been set.

### GetGeocoder

`func (o *POSTMarkets201ResponseDataRelationships) GetGeocoder() POSTAddresses201ResponseDataRelationshipsGeocoder`

GetGeocoder returns the Geocoder field if non-nil, zero value otherwise.

### GetGeocoderOk

`func (o *POSTMarkets201ResponseDataRelationships) GetGeocoderOk() (*POSTAddresses201ResponseDataRelationshipsGeocoder, bool)`

GetGeocoderOk returns a tuple with the Geocoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoder

`func (o *POSTMarkets201ResponseDataRelationships) SetGeocoder(v POSTAddresses201ResponseDataRelationshipsGeocoder)`

SetGeocoder sets Geocoder field to given value.

### HasGeocoder

`func (o *POSTMarkets201ResponseDataRelationships) HasGeocoder() bool`

HasGeocoder returns a boolean if a field has been set.

### GetStores

`func (o *POSTMarkets201ResponseDataRelationships) GetStores() POSTMarkets201ResponseDataRelationshipsStores`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *POSTMarkets201ResponseDataRelationships) GetStoresOk() (*POSTMarkets201ResponseDataRelationshipsStores, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *POSTMarkets201ResponseDataRelationships) SetStores(v POSTMarkets201ResponseDataRelationshipsStores)`

SetStores sets Stores field to given value.

### HasStores

`func (o *POSTMarkets201ResponseDataRelationships) HasStores() bool`

HasStores returns a boolean if a field has been set.

### GetPriceListSchedulers

`func (o *POSTMarkets201ResponseDataRelationships) GetPriceListSchedulers() POSTMarkets201ResponseDataRelationshipsPriceListSchedulers`

GetPriceListSchedulers returns the PriceListSchedulers field if non-nil, zero value otherwise.

### GetPriceListSchedulersOk

`func (o *POSTMarkets201ResponseDataRelationships) GetPriceListSchedulersOk() (*POSTMarkets201ResponseDataRelationshipsPriceListSchedulers, bool)`

GetPriceListSchedulersOk returns a tuple with the PriceListSchedulers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListSchedulers

`func (o *POSTMarkets201ResponseDataRelationships) SetPriceListSchedulers(v POSTMarkets201ResponseDataRelationshipsPriceListSchedulers)`

SetPriceListSchedulers sets PriceListSchedulers field to given value.

### HasPriceListSchedulers

`func (o *POSTMarkets201ResponseDataRelationships) HasPriceListSchedulers() bool`

HasPriceListSchedulers returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTMarkets201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTMarkets201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTMarkets201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTMarkets201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *POSTMarkets201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTMarkets201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTMarkets201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTMarkets201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


