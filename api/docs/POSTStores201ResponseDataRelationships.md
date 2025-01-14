# POSTStores201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBundles201ResponseDataRelationshipsMarket**](POSTBundles201ResponseDataRelationshipsMarket.md) |  | [optional] 
**Merchant** | Pointer to [**POSTMarkets201ResponseDataRelationshipsMerchant**](POSTMarkets201ResponseDataRelationshipsMerchant.md) |  | [optional] 
**StockLocation** | Pointer to [**POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation**](POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation.md) |  | [optional] 
**Orders** | Pointer to [**POSTCustomers201ResponseDataRelationshipsOrders**](POSTCustomers201ResponseDataRelationshipsOrders.md) |  | [optional] 
**PaymentMethods** | Pointer to [**POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods**](POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTStores201ResponseDataRelationships

`func NewPOSTStores201ResponseDataRelationships() *POSTStores201ResponseDataRelationships`

NewPOSTStores201ResponseDataRelationships instantiates a new POSTStores201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStores201ResponseDataRelationshipsWithDefaults

`func NewPOSTStores201ResponseDataRelationshipsWithDefaults() *POSTStores201ResponseDataRelationships`

NewPOSTStores201ResponseDataRelationshipsWithDefaults instantiates a new POSTStores201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTStores201ResponseDataRelationships) GetMarket() POSTBundles201ResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTStores201ResponseDataRelationships) GetMarketOk() (*POSTBundles201ResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTStores201ResponseDataRelationships) SetMarket(v POSTBundles201ResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTStores201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetMerchant

`func (o *POSTStores201ResponseDataRelationships) GetMerchant() POSTMarkets201ResponseDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *POSTStores201ResponseDataRelationships) GetMerchantOk() (*POSTMarkets201ResponseDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *POSTStores201ResponseDataRelationships) SetMerchant(v POSTMarkets201ResponseDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *POSTStores201ResponseDataRelationships) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetStockLocation

`func (o *POSTStores201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *POSTStores201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *POSTStores201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *POSTStores201ResponseDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOrders

`func (o *POSTStores201ResponseDataRelationships) GetOrders() POSTCustomers201ResponseDataRelationshipsOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *POSTStores201ResponseDataRelationships) GetOrdersOk() (*POSTCustomers201ResponseDataRelationshipsOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *POSTStores201ResponseDataRelationships) SetOrders(v POSTCustomers201ResponseDataRelationshipsOrders)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *POSTStores201ResponseDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *POSTStores201ResponseDataRelationships) GetPaymentMethods() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *POSTStores201ResponseDataRelationships) GetPaymentMethodsOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *POSTStores201ResponseDataRelationships) SetPaymentMethods(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *POSTStores201ResponseDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetEvents

`func (o *POSTStores201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTStores201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTStores201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTStores201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *POSTStores201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTStores201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTStores201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTStores201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


