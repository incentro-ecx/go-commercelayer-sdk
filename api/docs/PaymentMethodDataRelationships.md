# PaymentMethodDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**PaymentGateway** | Pointer to [**AdyenPaymentDataRelationshipsPaymentGateway**](AdyenPaymentDataRelationshipsPaymentGateway.md) |  | [optional] 
**Store** | Pointer to [**MarketDataRelationshipsStores**](MarketDataRelationshipsStores.md) |  | [optional] 
**Orders** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPaymentMethodDataRelationships

`func NewPaymentMethodDataRelationships() *PaymentMethodDataRelationships`

NewPaymentMethodDataRelationships instantiates a new PaymentMethodDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodDataRelationshipsWithDefaults

`func NewPaymentMethodDataRelationshipsWithDefaults() *PaymentMethodDataRelationships`

NewPaymentMethodDataRelationshipsWithDefaults instantiates a new PaymentMethodDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *PaymentMethodDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *PaymentMethodDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *PaymentMethodDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *PaymentMethodDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *PaymentMethodDataRelationships) GetPaymentGateway() AdyenPaymentDataRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *PaymentMethodDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentDataRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *PaymentMethodDataRelationships) SetPaymentGateway(v AdyenPaymentDataRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.

### HasPaymentGateway

`func (o *PaymentMethodDataRelationships) HasPaymentGateway() bool`

HasPaymentGateway returns a boolean if a field has been set.

### GetStore

`func (o *PaymentMethodDataRelationships) GetStore() MarketDataRelationshipsStores`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PaymentMethodDataRelationships) GetStoreOk() (*MarketDataRelationshipsStores, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PaymentMethodDataRelationships) SetStore(v MarketDataRelationshipsStores)`

SetStore sets Store field to given value.

### HasStore

`func (o *PaymentMethodDataRelationships) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetOrders

`func (o *PaymentMethodDataRelationships) GetOrders() AdyenPaymentDataRelationshipsOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PaymentMethodDataRelationships) GetOrdersOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PaymentMethodDataRelationships) SetOrders(v AdyenPaymentDataRelationshipsOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *PaymentMethodDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetAttachments

`func (o *PaymentMethodDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PaymentMethodDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PaymentMethodDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PaymentMethodDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *PaymentMethodDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PaymentMethodDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PaymentMethodDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PaymentMethodDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


