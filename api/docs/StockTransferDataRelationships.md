# StockTransferDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**OriginStockLocation** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 
**DestinationStockLocation** | Pointer to [**DeliveryLeadTimeDataRelationshipsStockLocation**](DeliveryLeadTimeDataRelationshipsStockLocation.md) |  | [optional] 
**Shipment** | Pointer to [**LineItemDataRelationshipsShipment**](LineItemDataRelationshipsShipment.md) |  | [optional] 
**LineItem** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 
**StockReservation** | Pointer to [**LineItemDataRelationshipsStockReservations**](LineItemDataRelationshipsStockReservations.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewStockTransferDataRelationships

`func NewStockTransferDataRelationships() *StockTransferDataRelationships`

NewStockTransferDataRelationships instantiates a new StockTransferDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockTransferDataRelationshipsWithDefaults

`func NewStockTransferDataRelationshipsWithDefaults() *StockTransferDataRelationships`

NewStockTransferDataRelationshipsWithDefaults instantiates a new StockTransferDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *StockTransferDataRelationships) GetSku() BundleDataRelationshipsSkus`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockTransferDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockTransferDataRelationships) SetSku(v BundleDataRelationshipsSkus)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StockTransferDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetOriginStockLocation

`func (o *StockTransferDataRelationships) GetOriginStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetOriginStockLocation returns the OriginStockLocation field if non-nil, zero value otherwise.

### GetOriginStockLocationOk

`func (o *StockTransferDataRelationships) GetOriginStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetOriginStockLocationOk returns a tuple with the OriginStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStockLocation

`func (o *StockTransferDataRelationships) SetOriginStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetOriginStockLocation sets OriginStockLocation field to given value.

### HasOriginStockLocation

`func (o *StockTransferDataRelationships) HasOriginStockLocation() bool`

HasOriginStockLocation returns a boolean if a field has been set.

### GetDestinationStockLocation

`func (o *StockTransferDataRelationships) GetDestinationStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation`

GetDestinationStockLocation returns the DestinationStockLocation field if non-nil, zero value otherwise.

### GetDestinationStockLocationOk

`func (o *StockTransferDataRelationships) GetDestinationStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool)`

GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStockLocation

`func (o *StockTransferDataRelationships) SetDestinationStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation)`

SetDestinationStockLocation sets DestinationStockLocation field to given value.

### HasDestinationStockLocation

`func (o *StockTransferDataRelationships) HasDestinationStockLocation() bool`

HasDestinationStockLocation returns a boolean if a field has been set.

### GetShipment

`func (o *StockTransferDataRelationships) GetShipment() LineItemDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *StockTransferDataRelationships) GetShipmentOk() (*LineItemDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *StockTransferDataRelationships) SetShipment(v LineItemDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *StockTransferDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetLineItem

`func (o *StockTransferDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *StockTransferDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *StockTransferDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *StockTransferDataRelationships) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetStockReservation

`func (o *StockTransferDataRelationships) GetStockReservation() LineItemDataRelationshipsStockReservations`

GetStockReservation returns the StockReservation field if non-nil, zero value otherwise.

### GetStockReservationOk

`func (o *StockTransferDataRelationships) GetStockReservationOk() (*LineItemDataRelationshipsStockReservations, bool)`

GetStockReservationOk returns a tuple with the StockReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservation

`func (o *StockTransferDataRelationships) SetStockReservation(v LineItemDataRelationshipsStockReservations)`

SetStockReservation sets StockReservation field to given value.

### HasStockReservation

`func (o *StockTransferDataRelationships) HasStockReservation() bool`

HasStockReservation returns a boolean if a field has been set.

### GetAttachments

`func (o *StockTransferDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *StockTransferDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *StockTransferDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *StockTransferDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *StockTransferDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StockTransferDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StockTransferDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StockTransferDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *StockTransferDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *StockTransferDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *StockTransferDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *StockTransferDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


