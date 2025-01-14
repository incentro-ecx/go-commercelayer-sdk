# LineItemDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**Item** | Pointer to [**LineItemDataRelationshipsItem**](LineItemDataRelationshipsItem.md) |  | [optional] 
**Sku** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**Bundle** | Pointer to [**LineItemDataRelationshipsBundle**](LineItemDataRelationshipsBundle.md) |  | [optional] 
**Adjustment** | Pointer to [**LineItemDataRelationshipsAdjustment**](LineItemDataRelationshipsAdjustment.md) |  | [optional] 
**GiftCard** | Pointer to [**LineItemDataRelationshipsGiftCard**](LineItemDataRelationshipsGiftCard.md) |  | [optional] 
**Shipment** | Pointer to [**LineItemDataRelationshipsShipment**](LineItemDataRelationshipsShipment.md) |  | [optional] 
**PaymentMethod** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**LineItemOptions** | Pointer to [**LineItemDataRelationshipsLineItemOptions**](LineItemDataRelationshipsLineItemOptions.md) |  | [optional] 
**ReturnLineItems** | Pointer to [**LineItemDataRelationshipsReturnLineItems**](LineItemDataRelationshipsReturnLineItems.md) |  | [optional] 
**StockReservations** | Pointer to [**LineItemDataRelationshipsStockReservations**](LineItemDataRelationshipsStockReservations.md) |  | [optional] 
**StockLineItems** | Pointer to [**LineItemDataRelationshipsStockLineItems**](LineItemDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**LineItemDataRelationshipsStockTransfers**](LineItemDataRelationshipsStockTransfers.md) |  | [optional] 
**Notifications** | Pointer to [**LineItemDataRelationshipsNotifications**](LineItemDataRelationshipsNotifications.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**AddressDataRelationshipsTags**](AddressDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewLineItemDataRelationships

`func NewLineItemDataRelationships() *LineItemDataRelationships`

NewLineItemDataRelationships instantiates a new LineItemDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemDataRelationshipsWithDefaults

`func NewLineItemDataRelationshipsWithDefaults() *LineItemDataRelationships`

NewLineItemDataRelationshipsWithDefaults instantiates a new LineItemDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *LineItemDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LineItemDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *LineItemDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *LineItemDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetItem

`func (o *LineItemDataRelationships) GetItem() LineItemDataRelationshipsItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *LineItemDataRelationships) GetItemOk() (*LineItemDataRelationshipsItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *LineItemDataRelationships) SetItem(v LineItemDataRelationshipsItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *LineItemDataRelationships) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetSku

`func (o *LineItemDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *LineItemDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *LineItemDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.

### HasSku

`func (o *LineItemDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetBundle

`func (o *LineItemDataRelationships) GetBundle() LineItemDataRelationshipsBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *LineItemDataRelationships) GetBundleOk() (*LineItemDataRelationshipsBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *LineItemDataRelationships) SetBundle(v LineItemDataRelationshipsBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *LineItemDataRelationships) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetAdjustment

`func (o *LineItemDataRelationships) GetAdjustment() LineItemDataRelationshipsAdjustment`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *LineItemDataRelationships) GetAdjustmentOk() (*LineItemDataRelationshipsAdjustment, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *LineItemDataRelationships) SetAdjustment(v LineItemDataRelationshipsAdjustment)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *LineItemDataRelationships) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetGiftCard

`func (o *LineItemDataRelationships) GetGiftCard() LineItemDataRelationshipsGiftCard`

GetGiftCard returns the GiftCard field if non-nil, zero value otherwise.

### GetGiftCardOk

`func (o *LineItemDataRelationships) GetGiftCardOk() (*LineItemDataRelationshipsGiftCard, bool)`

GetGiftCardOk returns a tuple with the GiftCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCard

`func (o *LineItemDataRelationships) SetGiftCard(v LineItemDataRelationshipsGiftCard)`

SetGiftCard sets GiftCard field to given value.

### HasGiftCard

`func (o *LineItemDataRelationships) HasGiftCard() bool`

HasGiftCard returns a boolean if a field has been set.

### GetShipment

`func (o *LineItemDataRelationships) GetShipment() LineItemDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *LineItemDataRelationships) GetShipmentOk() (*LineItemDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *LineItemDataRelationships) SetShipment(v LineItemDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *LineItemDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *LineItemDataRelationships) GetPaymentMethod() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *LineItemDataRelationships) GetPaymentMethodOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *LineItemDataRelationships) SetPaymentMethod(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *LineItemDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetLineItemOptions

`func (o *LineItemDataRelationships) GetLineItemOptions() LineItemDataRelationshipsLineItemOptions`

GetLineItemOptions returns the LineItemOptions field if non-nil, zero value otherwise.

### GetLineItemOptionsOk

`func (o *LineItemDataRelationships) GetLineItemOptionsOk() (*LineItemDataRelationshipsLineItemOptions, bool)`

GetLineItemOptionsOk returns a tuple with the LineItemOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptions

`func (o *LineItemDataRelationships) SetLineItemOptions(v LineItemDataRelationshipsLineItemOptions)`

SetLineItemOptions sets LineItemOptions field to given value.

### HasLineItemOptions

`func (o *LineItemDataRelationships) HasLineItemOptions() bool`

HasLineItemOptions returns a boolean if a field has been set.

### GetReturnLineItems

`func (o *LineItemDataRelationships) GetReturnLineItems() LineItemDataRelationshipsReturnLineItems`

GetReturnLineItems returns the ReturnLineItems field if non-nil, zero value otherwise.

### GetReturnLineItemsOk

`func (o *LineItemDataRelationships) GetReturnLineItemsOk() (*LineItemDataRelationshipsReturnLineItems, bool)`

GetReturnLineItemsOk returns a tuple with the ReturnLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLineItems

`func (o *LineItemDataRelationships) SetReturnLineItems(v LineItemDataRelationshipsReturnLineItems)`

SetReturnLineItems sets ReturnLineItems field to given value.

### HasReturnLineItems

`func (o *LineItemDataRelationships) HasReturnLineItems() bool`

HasReturnLineItems returns a boolean if a field has been set.

### GetStockReservations

`func (o *LineItemDataRelationships) GetStockReservations() LineItemDataRelationshipsStockReservations`

GetStockReservations returns the StockReservations field if non-nil, zero value otherwise.

### GetStockReservationsOk

`func (o *LineItemDataRelationships) GetStockReservationsOk() (*LineItemDataRelationshipsStockReservations, bool)`

GetStockReservationsOk returns a tuple with the StockReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservations

`func (o *LineItemDataRelationships) SetStockReservations(v LineItemDataRelationshipsStockReservations)`

SetStockReservations sets StockReservations field to given value.

### HasStockReservations

`func (o *LineItemDataRelationships) HasStockReservations() bool`

HasStockReservations returns a boolean if a field has been set.

### GetStockLineItems

`func (o *LineItemDataRelationships) GetStockLineItems() LineItemDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *LineItemDataRelationships) GetStockLineItemsOk() (*LineItemDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *LineItemDataRelationships) SetStockLineItems(v LineItemDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *LineItemDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *LineItemDataRelationships) GetStockTransfers() LineItemDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *LineItemDataRelationships) GetStockTransfersOk() (*LineItemDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *LineItemDataRelationships) SetStockTransfers(v LineItemDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *LineItemDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetNotifications

`func (o *LineItemDataRelationships) GetNotifications() LineItemDataRelationshipsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *LineItemDataRelationships) GetNotificationsOk() (*LineItemDataRelationshipsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *LineItemDataRelationships) SetNotifications(v LineItemDataRelationshipsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *LineItemDataRelationships) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetEvents

`func (o *LineItemDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *LineItemDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *LineItemDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *LineItemDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *LineItemDataRelationships) GetTags() AddressDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LineItemDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LineItemDataRelationships) SetTags(v AddressDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LineItemDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


