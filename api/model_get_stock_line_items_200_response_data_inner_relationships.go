/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETStockLineItems200ResponseDataInnerRelationships struct for GETStockLineItems200ResponseDataInnerRelationships
type GETStockLineItems200ResponseDataInnerRelationships struct {
	Shipment  *GETParcels200ResponseDataInnerRelationshipsShipment         `json:"shipment,omitempty"`
	LineItem  *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem `json:"line_item,omitempty"`
	StockItem *GETStockLineItems200ResponseDataInnerRelationshipsStockItem `json:"stock_item,omitempty"`
}

// NewGETStockLineItems200ResponseDataInnerRelationships instantiates a new GETStockLineItems200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockLineItems200ResponseDataInnerRelationships() *GETStockLineItems200ResponseDataInnerRelationships {
	this := GETStockLineItems200ResponseDataInnerRelationships{}
	return &this
}

// NewGETStockLineItems200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETStockLineItems200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockLineItems200ResponseDataInnerRelationshipsWithDefaults() *GETStockLineItems200ResponseDataInnerRelationships {
	this := GETStockLineItems200ResponseDataInnerRelationships{}
	return &this
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *GETStockLineItems200ResponseDataInnerRelationships) GetShipment() GETParcels200ResponseDataInnerRelationshipsShipment {
	if o == nil || o.Shipment == nil {
		var ret GETParcels200ResponseDataInnerRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationships) GetShipmentOk() (*GETParcels200ResponseDataInnerRelationshipsShipment, bool) {
	if o == nil || o.Shipment == nil {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationships) HasShipment() bool {
	if o != nil && o.Shipment != nil {
		return true
	}

	return false
}

// SetShipment gets a reference to the given GETParcels200ResponseDataInnerRelationshipsShipment and assigns it to the Shipment field.
func (o *GETStockLineItems200ResponseDataInnerRelationships) SetShipment(v GETParcels200ResponseDataInnerRelationshipsShipment) {
	o.Shipment = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *GETStockLineItems200ResponseDataInnerRelationships) GetLineItem() GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	if o == nil || o.LineItem == nil {
		var ret GETLineItemOptions200ResponseDataInnerRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationships) GetLineItemOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsLineItem, bool) {
	if o == nil || o.LineItem == nil {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationships) HasLineItem() bool {
	if o != nil && o.LineItem != nil {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given GETLineItemOptions200ResponseDataInnerRelationshipsLineItem and assigns it to the LineItem field.
func (o *GETStockLineItems200ResponseDataInnerRelationships) SetLineItem(v GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) {
	o.LineItem = &v
}

// GetStockItem returns the StockItem field value if set, zero value otherwise.
func (o *GETStockLineItems200ResponseDataInnerRelationships) GetStockItem() GETStockLineItems200ResponseDataInnerRelationshipsStockItem {
	if o == nil || o.StockItem == nil {
		var ret GETStockLineItems200ResponseDataInnerRelationshipsStockItem
		return ret
	}
	return *o.StockItem
}

// GetStockItemOk returns a tuple with the StockItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationships) GetStockItemOk() (*GETStockLineItems200ResponseDataInnerRelationshipsStockItem, bool) {
	if o == nil || o.StockItem == nil {
		return nil, false
	}
	return o.StockItem, true
}

// HasStockItem returns a boolean if a field has been set.
func (o *GETStockLineItems200ResponseDataInnerRelationships) HasStockItem() bool {
	if o != nil && o.StockItem != nil {
		return true
	}

	return false
}

// SetStockItem gets a reference to the given GETStockLineItems200ResponseDataInnerRelationshipsStockItem and assigns it to the StockItem field.
func (o *GETStockLineItems200ResponseDataInnerRelationships) SetStockItem(v GETStockLineItems200ResponseDataInnerRelationshipsStockItem) {
	o.StockItem = &v
}

func (o GETStockLineItems200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shipment != nil {
		toSerialize["shipment"] = o.Shipment
	}
	if o.LineItem != nil {
		toSerialize["line_item"] = o.LineItem
	}
	if o.StockItem != nil {
		toSerialize["stock_item"] = o.StockItem
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockLineItems200ResponseDataInnerRelationships struct {
	value *GETStockLineItems200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETStockLineItems200ResponseDataInnerRelationships) Get() *GETStockLineItems200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETStockLineItems200ResponseDataInnerRelationships) Set(val *GETStockLineItems200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockLineItems200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockLineItems200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockLineItems200ResponseDataInnerRelationships(val *GETStockLineItems200ResponseDataInnerRelationships) *NullableGETStockLineItems200ResponseDataInnerRelationships {
	return &NullableGETStockLineItems200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETStockLineItems200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockLineItems200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
