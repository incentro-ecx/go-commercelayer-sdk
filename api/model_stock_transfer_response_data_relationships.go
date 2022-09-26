/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StockTransferResponseDataRelationships struct for StockTransferResponseDataRelationships
type StockTransferResponseDataRelationships struct {
	Sku                      *InStockSubscriptionResponseDataRelationshipsSku                `json:"sku,omitempty"`
	OriginStockLocation      *StockTransferResponseDataRelationshipsOriginStockLocation      `json:"origin_stock_location,omitempty"`
	DestinationStockLocation *StockTransferResponseDataRelationshipsDestinationStockLocation `json:"destination_stock_location,omitempty"`
	Shipment                 *ParcelResponseDataRelationshipsShipment                        `json:"shipment,omitempty"`
	LineItem                 *LineItemOptionResponseDataRelationshipsLineItem                `json:"line_item,omitempty"`
	Events                   *CustomerAddressResponseDataRelationshipsEvents                 `json:"events,omitempty"`
}

// NewStockTransferResponseDataRelationships instantiates a new StockTransferResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferResponseDataRelationships() *StockTransferResponseDataRelationships {
	this := StockTransferResponseDataRelationships{}
	return &this
}

// NewStockTransferResponseDataRelationshipsWithDefaults instantiates a new StockTransferResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferResponseDataRelationshipsWithDefaults() *StockTransferResponseDataRelationships {
	this := StockTransferResponseDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockTransferResponseDataRelationships) GetSku() InStockSubscriptionResponseDataRelationshipsSku {
	if o == nil || o.Sku == nil {
		var ret InStockSubscriptionResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferResponseDataRelationships) GetSkuOk() (*InStockSubscriptionResponseDataRelationshipsSku, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockTransferResponseDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *StockTransferResponseDataRelationships) SetSku(v InStockSubscriptionResponseDataRelationshipsSku) {
	o.Sku = &v
}

// GetOriginStockLocation returns the OriginStockLocation field value if set, zero value otherwise.
func (o *StockTransferResponseDataRelationships) GetOriginStockLocation() StockTransferResponseDataRelationshipsOriginStockLocation {
	if o == nil || o.OriginStockLocation == nil {
		var ret StockTransferResponseDataRelationshipsOriginStockLocation
		return ret
	}
	return *o.OriginStockLocation
}

// GetOriginStockLocationOk returns a tuple with the OriginStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferResponseDataRelationships) GetOriginStockLocationOk() (*StockTransferResponseDataRelationshipsOriginStockLocation, bool) {
	if o == nil || o.OriginStockLocation == nil {
		return nil, false
	}
	return o.OriginStockLocation, true
}

// HasOriginStockLocation returns a boolean if a field has been set.
func (o *StockTransferResponseDataRelationships) HasOriginStockLocation() bool {
	if o != nil && o.OriginStockLocation != nil {
		return true
	}

	return false
}

// SetOriginStockLocation gets a reference to the given StockTransferResponseDataRelationshipsOriginStockLocation and assigns it to the OriginStockLocation field.
func (o *StockTransferResponseDataRelationships) SetOriginStockLocation(v StockTransferResponseDataRelationshipsOriginStockLocation) {
	o.OriginStockLocation = &v
}

// GetDestinationStockLocation returns the DestinationStockLocation field value if set, zero value otherwise.
func (o *StockTransferResponseDataRelationships) GetDestinationStockLocation() StockTransferResponseDataRelationshipsDestinationStockLocation {
	if o == nil || o.DestinationStockLocation == nil {
		var ret StockTransferResponseDataRelationshipsDestinationStockLocation
		return ret
	}
	return *o.DestinationStockLocation
}

// GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferResponseDataRelationships) GetDestinationStockLocationOk() (*StockTransferResponseDataRelationshipsDestinationStockLocation, bool) {
	if o == nil || o.DestinationStockLocation == nil {
		return nil, false
	}
	return o.DestinationStockLocation, true
}

// HasDestinationStockLocation returns a boolean if a field has been set.
func (o *StockTransferResponseDataRelationships) HasDestinationStockLocation() bool {
	if o != nil && o.DestinationStockLocation != nil {
		return true
	}

	return false
}

// SetDestinationStockLocation gets a reference to the given StockTransferResponseDataRelationshipsDestinationStockLocation and assigns it to the DestinationStockLocation field.
func (o *StockTransferResponseDataRelationships) SetDestinationStockLocation(v StockTransferResponseDataRelationshipsDestinationStockLocation) {
	o.DestinationStockLocation = &v
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *StockTransferResponseDataRelationships) GetShipment() ParcelResponseDataRelationshipsShipment {
	if o == nil || o.Shipment == nil {
		var ret ParcelResponseDataRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferResponseDataRelationships) GetShipmentOk() (*ParcelResponseDataRelationshipsShipment, bool) {
	if o == nil || o.Shipment == nil {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *StockTransferResponseDataRelationships) HasShipment() bool {
	if o != nil && o.Shipment != nil {
		return true
	}

	return false
}

// SetShipment gets a reference to the given ParcelResponseDataRelationshipsShipment and assigns it to the Shipment field.
func (o *StockTransferResponseDataRelationships) SetShipment(v ParcelResponseDataRelationshipsShipment) {
	o.Shipment = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *StockTransferResponseDataRelationships) GetLineItem() LineItemOptionResponseDataRelationshipsLineItem {
	if o == nil || o.LineItem == nil {
		var ret LineItemOptionResponseDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferResponseDataRelationships) GetLineItemOk() (*LineItemOptionResponseDataRelationshipsLineItem, bool) {
	if o == nil || o.LineItem == nil {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *StockTransferResponseDataRelationships) HasLineItem() bool {
	if o != nil && o.LineItem != nil {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given LineItemOptionResponseDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *StockTransferResponseDataRelationships) SetLineItem(v LineItemOptionResponseDataRelationshipsLineItem) {
	o.LineItem = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *StockTransferResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StockTransferResponseDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *StockTransferResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents) {
	o.Events = &v
}

func (o StockTransferResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.OriginStockLocation != nil {
		toSerialize["origin_stock_location"] = o.OriginStockLocation
	}
	if o.DestinationStockLocation != nil {
		toSerialize["destination_stock_location"] = o.DestinationStockLocation
	}
	if o.Shipment != nil {
		toSerialize["shipment"] = o.Shipment
	}
	if o.LineItem != nil {
		toSerialize["line_item"] = o.LineItem
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableStockTransferResponseDataRelationships struct {
	value *StockTransferResponseDataRelationships
	isSet bool
}

func (v NullableStockTransferResponseDataRelationships) Get() *StockTransferResponseDataRelationships {
	return v.value
}

func (v *NullableStockTransferResponseDataRelationships) Set(val *StockTransferResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferResponseDataRelationships(val *StockTransferResponseDataRelationships) *NullableStockTransferResponseDataRelationships {
	return &NullableStockTransferResponseDataRelationships{value: val, isSet: true}
}

func (v NullableStockTransferResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}