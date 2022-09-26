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

// LineItemResponseDataRelationships struct for LineItemResponseDataRelationships
type LineItemResponseDataRelationships struct {
	Order             *AdyenPaymentResponseDataRelationshipsOrder         `json:"order,omitempty"`
	Item              *LineItemResponseDataRelationshipsItem              `json:"item,omitempty"`
	LineItemOptions   *LineItemResponseDataRelationshipsLineItemOptions   `json:"line_item_options,omitempty"`
	ShipmentLineItems *LineItemResponseDataRelationshipsShipmentLineItems `json:"shipment_line_items,omitempty"`
	StockLineItems    *LineItemResponseDataRelationshipsStockLineItems    `json:"stock_line_items,omitempty"`
	StockTransfers    *LineItemResponseDataRelationshipsStockTransfers    `json:"stock_transfers,omitempty"`
}

// NewLineItemResponseDataRelationships instantiates a new LineItemResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemResponseDataRelationships() *LineItemResponseDataRelationships {
	this := LineItemResponseDataRelationships{}
	return &this
}

// NewLineItemResponseDataRelationshipsWithDefaults instantiates a new LineItemResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemResponseDataRelationshipsWithDefaults() *LineItemResponseDataRelationships {
	this := LineItemResponseDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *LineItemResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder) {
	o.Order = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationships) GetItem() LineItemResponseDataRelationshipsItem {
	if o == nil || o.Item == nil {
		var ret LineItemResponseDataRelationshipsItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationships) GetItemOk() (*LineItemResponseDataRelationshipsItem, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationships) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given LineItemResponseDataRelationshipsItem and assigns it to the Item field.
func (o *LineItemResponseDataRelationships) SetItem(v LineItemResponseDataRelationshipsItem) {
	o.Item = &v
}

// GetLineItemOptions returns the LineItemOptions field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationships) GetLineItemOptions() LineItemResponseDataRelationshipsLineItemOptions {
	if o == nil || o.LineItemOptions == nil {
		var ret LineItemResponseDataRelationshipsLineItemOptions
		return ret
	}
	return *o.LineItemOptions
}

// GetLineItemOptionsOk returns a tuple with the LineItemOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationships) GetLineItemOptionsOk() (*LineItemResponseDataRelationshipsLineItemOptions, bool) {
	if o == nil || o.LineItemOptions == nil {
		return nil, false
	}
	return o.LineItemOptions, true
}

// HasLineItemOptions returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationships) HasLineItemOptions() bool {
	if o != nil && o.LineItemOptions != nil {
		return true
	}

	return false
}

// SetLineItemOptions gets a reference to the given LineItemResponseDataRelationshipsLineItemOptions and assigns it to the LineItemOptions field.
func (o *LineItemResponseDataRelationships) SetLineItemOptions(v LineItemResponseDataRelationshipsLineItemOptions) {
	o.LineItemOptions = &v
}

// GetShipmentLineItems returns the ShipmentLineItems field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationships) GetShipmentLineItems() LineItemResponseDataRelationshipsShipmentLineItems {
	if o == nil || o.ShipmentLineItems == nil {
		var ret LineItemResponseDataRelationshipsShipmentLineItems
		return ret
	}
	return *o.ShipmentLineItems
}

// GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationships) GetShipmentLineItemsOk() (*LineItemResponseDataRelationshipsShipmentLineItems, bool) {
	if o == nil || o.ShipmentLineItems == nil {
		return nil, false
	}
	return o.ShipmentLineItems, true
}

// HasShipmentLineItems returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationships) HasShipmentLineItems() bool {
	if o != nil && o.ShipmentLineItems != nil {
		return true
	}

	return false
}

// SetShipmentLineItems gets a reference to the given LineItemResponseDataRelationshipsShipmentLineItems and assigns it to the ShipmentLineItems field.
func (o *LineItemResponseDataRelationships) SetShipmentLineItems(v LineItemResponseDataRelationshipsShipmentLineItems) {
	o.ShipmentLineItems = &v
}

// GetStockLineItems returns the StockLineItems field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationships) GetStockLineItems() LineItemResponseDataRelationshipsStockLineItems {
	if o == nil || o.StockLineItems == nil {
		var ret LineItemResponseDataRelationshipsStockLineItems
		return ret
	}
	return *o.StockLineItems
}

// GetStockLineItemsOk returns a tuple with the StockLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationships) GetStockLineItemsOk() (*LineItemResponseDataRelationshipsStockLineItems, bool) {
	if o == nil || o.StockLineItems == nil {
		return nil, false
	}
	return o.StockLineItems, true
}

// HasStockLineItems returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationships) HasStockLineItems() bool {
	if o != nil && o.StockLineItems != nil {
		return true
	}

	return false
}

// SetStockLineItems gets a reference to the given LineItemResponseDataRelationshipsStockLineItems and assigns it to the StockLineItems field.
func (o *LineItemResponseDataRelationships) SetStockLineItems(v LineItemResponseDataRelationshipsStockLineItems) {
	o.StockLineItems = &v
}

// GetStockTransfers returns the StockTransfers field value if set, zero value otherwise.
func (o *LineItemResponseDataRelationships) GetStockTransfers() LineItemResponseDataRelationshipsStockTransfers {
	if o == nil || o.StockTransfers == nil {
		var ret LineItemResponseDataRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfers
}

// GetStockTransfersOk returns a tuple with the StockTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemResponseDataRelationships) GetStockTransfersOk() (*LineItemResponseDataRelationshipsStockTransfers, bool) {
	if o == nil || o.StockTransfers == nil {
		return nil, false
	}
	return o.StockTransfers, true
}

// HasStockTransfers returns a boolean if a field has been set.
func (o *LineItemResponseDataRelationships) HasStockTransfers() bool {
	if o != nil && o.StockTransfers != nil {
		return true
	}

	return false
}

// SetStockTransfers gets a reference to the given LineItemResponseDataRelationshipsStockTransfers and assigns it to the StockTransfers field.
func (o *LineItemResponseDataRelationships) SetStockTransfers(v LineItemResponseDataRelationshipsStockTransfers) {
	o.StockTransfers = &v
}

func (o LineItemResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	if o.LineItemOptions != nil {
		toSerialize["line_item_options"] = o.LineItemOptions
	}
	if o.ShipmentLineItems != nil {
		toSerialize["shipment_line_items"] = o.ShipmentLineItems
	}
	if o.StockLineItems != nil {
		toSerialize["stock_line_items"] = o.StockLineItems
	}
	if o.StockTransfers != nil {
		toSerialize["stock_transfers"] = o.StockTransfers
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemResponseDataRelationships struct {
	value *LineItemResponseDataRelationships
	isSet bool
}

func (v NullableLineItemResponseDataRelationships) Get() *LineItemResponseDataRelationships {
	return v.value
}

func (v *NullableLineItemResponseDataRelationships) Set(val *LineItemResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemResponseDataRelationships(val *LineItemResponseDataRelationships) *NullableLineItemResponseDataRelationships {
	return &NullableLineItemResponseDataRelationships{value: val, isSet: true}
}

func (v NullableLineItemResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}