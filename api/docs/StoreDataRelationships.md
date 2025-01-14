# StoreDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**Merchant** | Pointer to [**MarketDataRelationshipsMerchant**](MarketDataRelationshipsMerchant.md) |  | [optional] 
**StockLocation** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 
**Orders** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**PaymentMethods** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewStoreDataRelationships

`func NewStoreDataRelationships() *StoreDataRelationships`

NewStoreDataRelationships instantiates a new StoreDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDataRelationshipsWithDefaults

`func NewStoreDataRelationshipsWithDefaults() *StoreDataRelationships`

NewStoreDataRelationshipsWithDefaults instantiates a new StoreDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *StoreDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *StoreDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *StoreDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *StoreDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetMerchant

`func (o *StoreDataRelationships) GetMerchant() MarketDataRelationshipsMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *StoreDataRelationships) GetMerchantOk() (*MarketDataRelationshipsMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *StoreDataRelationships) SetMerchant(v MarketDataRelationshipsMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *StoreDataRelationships) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetStockLocation

`func (o *StoreDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *StoreDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *StoreDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.

### HasStockLocation

`func (o *StoreDataRelationships) HasStockLocation() bool`

HasStockLocation returns a boolean if a field has been set.

### GetOrders

`func (o *StoreDataRelationships) GetOrders() AdyenPaymentDataRelationshipsOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *StoreDataRelationships) GetOrdersOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *StoreDataRelationships) SetOrders(v AdyenPaymentDataRelationshipsOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *StoreDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *StoreDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *StoreDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *StoreDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *StoreDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetEvents

`func (o *StoreDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StoreDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StoreDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StoreDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *StoreDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *StoreDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *StoreDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *StoreDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


