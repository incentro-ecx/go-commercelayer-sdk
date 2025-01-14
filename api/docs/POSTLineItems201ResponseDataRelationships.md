# POSTLineItems201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**Item** | Pointer to [**POSTLineItems201ResponseDataRelationshipsItem**](POSTLineItems201ResponseDataRelationshipsItem.md) |  | [optional] 
**Sku** | Pointer to [**POSTInStockSubscriptions201ResponseDataRelationshipsSku**](POSTInStockSubscriptions201ResponseDataRelationshipsSku.md) |  | [optional] 
**Bundle** | Pointer to [**POSTLineItems201ResponseDataRelationshipsBundle**](POSTLineItems201ResponseDataRelationshipsBundle.md) |  | [optional] 
**Adjustment** | Pointer to [**POSTLineItems201ResponseDataRelationshipsAdjustment**](POSTLineItems201ResponseDataRelationshipsAdjustment.md) |  | [optional] 
**GiftCard** | Pointer to [**POSTLineItems201ResponseDataRelationshipsGiftCard**](POSTLineItems201ResponseDataRelationshipsGiftCard.md) |  | [optional] 
**Shipment** | Pointer to [**POSTLineItems201ResponseDataRelationshipsShipment**](POSTLineItems201ResponseDataRelationshipsShipment.md) |  | [optional] 
**PaymentMethod** | Pointer to [**POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod**](POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod.md) |  | [optional] 
**LineItemOptions** | Pointer to [**POSTLineItems201ResponseDataRelationshipsLineItemOptions**](POSTLineItems201ResponseDataRelationshipsLineItemOptions.md) |  | [optional] 
**ReturnLineItems** | Pointer to [**POSTLineItems201ResponseDataRelationshipsReturnLineItems**](POSTLineItems201ResponseDataRelationshipsReturnLineItems.md) |  | [optional] 
**StockReservations** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockReservations**](POSTLineItems201ResponseDataRelationshipsStockReservations.md) |  | [optional] 
**StockLineItems** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockLineItems**](POSTLineItems201ResponseDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockTransfers**](POSTLineItems201ResponseDataRelationshipsStockTransfers.md) |  | [optional] 
**Notifications** | Pointer to [**POSTLineItems201ResponseDataRelationshipsNotifications**](POSTLineItems201ResponseDataRelationshipsNotifications.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**POSTAddresses201ResponseDataRelationshipsTags**](POSTAddresses201ResponseDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewPOSTLineItems201ResponseDataRelationships

`func NewPOSTLineItems201ResponseDataRelationships() *POSTLineItems201ResponseDataRelationships`

NewPOSTLineItems201ResponseDataRelationships instantiates a new POSTLineItems201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTLineItems201ResponseDataRelationshipsWithDefaults

`func NewPOSTLineItems201ResponseDataRelationshipsWithDefaults() *POSTLineItems201ResponseDataRelationships`

NewPOSTLineItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTLineItems201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *POSTLineItems201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *POSTLineItems201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *POSTLineItems201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *POSTLineItems201ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetItem

`func (o *POSTLineItems201ResponseDataRelationships) GetItem() POSTLineItems201ResponseDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *POSTLineItems201ResponseDataRelationships) GetItemOk() (*POSTLineItems201ResponseDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *POSTLineItems201ResponseDataRelationships) SetItem(v POSTLineItems201ResponseDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *POSTLineItems201ResponseDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetSku

`func (o *POSTLineItems201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *POSTLineItems201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *POSTLineItems201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *POSTLineItems201ResponseDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetBundle

`func (o *POSTLineItems201ResponseDataRelationships) GetBundle() POSTLineItems201ResponseDataRelationshipsBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *POSTLineItems201ResponseDataRelationships) GetBundleOk() (*POSTLineItems201ResponseDataRelationshipsBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *POSTLineItems201ResponseDataRelationships) SetBundle(v POSTLineItems201ResponseDataRelationshipsBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *POSTLineItems201ResponseDataRelationships) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetAdjustment

`func (o *POSTLineItems201ResponseDataRelationships) GetAdjustment() POSTLineItems201ResponseDataRelationshipsAdjustment`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *POSTLineItems201ResponseDataRelationships) GetAdjustmentOk() (*POSTLineItems201ResponseDataRelationshipsAdjustment, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *POSTLineItems201ResponseDataRelationships) SetAdjustment(v POSTLineItems201ResponseDataRelationshipsAdjustment)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *POSTLineItems201ResponseDataRelationships) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetGiftCard

`func (o *POSTLineItems201ResponseDataRelationships) GetGiftCard() POSTLineItems201ResponseDataRelationshipsGiftCard`

GetGiftCard returns the GiftCard field if non-nil, zero value otherwise.

### GetGiftCardOk

`func (o *POSTLineItems201ResponseDataRelationships) GetGiftCardOk() (*POSTLineItems201ResponseDataRelationshipsGiftCard, bool)`

GetGiftCardOk returns a tuple with the GiftCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCard

`func (o *POSTLineItems201ResponseDataRelationships) SetGiftCard(v POSTLineItems201ResponseDataRelationshipsGiftCard)`

SetGiftCard sets GiftCard field to given value.

### HasGiftCard

`func (o *POSTLineItems201ResponseDataRelationships) HasGiftCard() bool`

HasGiftCard returns a boolean if a field has been set.

### GetShipment

`func (o *POSTLineItems201ResponseDataRelationships) GetShipment() POSTLineItems201ResponseDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *POSTLineItems201ResponseDataRelationships) GetShipmentOk() (*POSTLineItems201ResponseDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *POSTLineItems201ResponseDataRelationships) SetShipment(v POSTLineItems201ResponseDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *POSTLineItems201ResponseDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *POSTLineItems201ResponseDataRelationships) GetPaymentMethod() POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *POSTLineItems201ResponseDataRelationships) GetPaymentMethodOk() (*POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *POSTLineItems201ResponseDataRelationships) SetPaymentMethod(v POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *POSTLineItems201ResponseDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetLineItemOptions

`func (o *POSTLineItems201ResponseDataRelationships) GetLineItemOptions() POSTLineItems201ResponseDataRelationshipsLineItemOptions`

GetLineItemOptions returns the LineItemOptions field if non-nil, zero value otherwise.

### GetLineItemOptionsOk

`func (o *POSTLineItems201ResponseDataRelationships) GetLineItemOptionsOk() (*POSTLineItems201ResponseDataRelationshipsLineItemOptions, bool)`

GetLineItemOptionsOk returns a tuple with the LineItemOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptions

`func (o *POSTLineItems201ResponseDataRelationships) SetLineItemOptions(v POSTLineItems201ResponseDataRelationshipsLineItemOptions)`

SetLineItemOptions sets LineItemOptions field to given value.

### HasLineItemOptions

`func (o *POSTLineItems201ResponseDataRelationships) HasLineItemOptions() bool`

HasLineItemOptions returns a boolean if a field has been set.

### GetReturnLineItems

`func (o *POSTLineItems201ResponseDataRelationships) GetReturnLineItems() POSTLineItems201ResponseDataRelationshipsReturnLineItems`

GetReturnLineItems returns the ReturnLineItems field if non-nil, zero value otherwise.

### GetReturnLineItemsOk

`func (o *POSTLineItems201ResponseDataRelationships) GetReturnLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsReturnLineItems, bool)`

GetReturnLineItemsOk returns a tuple with the ReturnLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineItems

`func (o *POSTLineItems201ResponseDataRelationships) SetReturnLineItems(v POSTLineItems201ResponseDataRelationshipsReturnLineItems)`

SetReturnLineItems sets ReturnLineItems field to given value.

### HasReturnLineItems

`func (o *POSTLineItems201ResponseDataRelationships) HasReturnLineItems() bool`

HasReturnLineItems returns a boolean if a field has been set.

### GetStockReservations

`func (o *POSTLineItems201ResponseDataRelationships) GetStockReservations() POSTLineItems201ResponseDataRelationshipsStockReservations`

GetStockReservations returns the StockReservations field if non-nil, zero value otherwise.

### GetStockReservationsOk

`func (o *POSTLineItems201ResponseDataRelationships) GetStockReservationsOk() (*POSTLineItems201ResponseDataRelationshipsStockReservations, bool)`

GetStockReservationsOk returns a tuple with the StockReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservations

`func (o *POSTLineItems201ResponseDataRelationships) SetStockReservations(v POSTLineItems201ResponseDataRelationshipsStockReservations)`

SetStockReservations sets StockReservations field to given value.

### HasStockReservations

`func (o *POSTLineItems201ResponseDataRelationships) HasStockReservations() bool`

HasStockReservations returns a boolean if a field has been set.

### GetStockLineItems

`func (o *POSTLineItems201ResponseDataRelationships) GetStockLineItems() POSTLineItems201ResponseDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *POSTLineItems201ResponseDataRelationships) GetStockLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *POSTLineItems201ResponseDataRelationships) SetStockLineItems(v POSTLineItems201ResponseDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *POSTLineItems201ResponseDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *POSTLineItems201ResponseDataRelationships) GetStockTransfers() POSTLineItems201ResponseDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *POSTLineItems201ResponseDataRelationships) GetStockTransfersOk() (*POSTLineItems201ResponseDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *POSTLineItems201ResponseDataRelationships) SetStockTransfers(v POSTLineItems201ResponseDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *POSTLineItems201ResponseDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetNotifications

`func (o *POSTLineItems201ResponseDataRelationships) GetNotifications() POSTLineItems201ResponseDataRelationshipsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *POSTLineItems201ResponseDataRelationships) GetNotificationsOk() (*POSTLineItems201ResponseDataRelationshipsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *POSTLineItems201ResponseDataRelationships) SetNotifications(v POSTLineItems201ResponseDataRelationshipsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *POSTLineItems201ResponseDataRelationships) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetEvents

`func (o *POSTLineItems201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTLineItems201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTLineItems201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTLineItems201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *POSTLineItems201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *POSTLineItems201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *POSTLineItems201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *POSTLineItems201ResponseDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


