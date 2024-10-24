/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTStockReservations201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockReservations201ResponseDataRelationships{}

// POSTStockReservations201ResponseDataRelationships struct for POSTStockReservations201ResponseDataRelationships
type POSTStockReservations201ResponseDataRelationships struct {
	LineItem      *POSTLineItemOptions201ResponseDataRelationshipsLineItem               `json:"line_item,omitempty"`
	Order         *POSTAdyenPayments201ResponseDataRelationshipsOrder                    `json:"order,omitempty"`
	StockLineItem *POSTParcelLineItems201ResponseDataRelationshipsStockLineItem          `json:"stock_line_item,omitempty"`
	StockTransfer *POSTStockReservations201ResponseDataRelationshipsStockTransfer        `json:"stock_transfer,omitempty"`
	StockItem     *GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem `json:"stock_item,omitempty"`
	ReservedStock *POSTStockItems201ResponseDataRelationshipsReservedStock               `json:"reserved_stock,omitempty"`
	Sku           *POSTInStockSubscriptions201ResponseDataRelationshipsSku               `json:"sku,omitempty"`
}

// NewPOSTStockReservations201ResponseDataRelationships instantiates a new POSTStockReservations201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockReservations201ResponseDataRelationships() *POSTStockReservations201ResponseDataRelationships {
	this := POSTStockReservations201ResponseDataRelationships{}
	return &this
}

// NewPOSTStockReservations201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockReservations201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockReservations201ResponseDataRelationshipsWithDefaults() *POSTStockReservations201ResponseDataRelationships {
	this := POSTStockReservations201ResponseDataRelationships{}
	return &this
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseDataRelationships) GetLineItem() POSTLineItemOptions201ResponseDataRelationshipsLineItem {
	if o == nil || IsNil(o.LineItem) {
		var ret POSTLineItemOptions201ResponseDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseDataRelationships) GetLineItemOk() (*POSTLineItemOptions201ResponseDataRelationshipsLineItem, bool) {
	if o == nil || IsNil(o.LineItem) {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationships) HasLineItem() bool {
	if o != nil && !IsNil(o.LineItem) {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given POSTLineItemOptions201ResponseDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *POSTStockReservations201ResponseDataRelationships) SetLineItem(v POSTLineItemOptions201ResponseDataRelationshipsLineItem) {
	o.LineItem = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *POSTStockReservations201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder) {
	o.Order = &v
}

// GetStockLineItem returns the StockLineItem field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseDataRelationships) GetStockLineItem() POSTParcelLineItems201ResponseDataRelationshipsStockLineItem {
	if o == nil || IsNil(o.StockLineItem) {
		var ret POSTParcelLineItems201ResponseDataRelationshipsStockLineItem
		return ret
	}
	return *o.StockLineItem
}

// GetStockLineItemOk returns a tuple with the StockLineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseDataRelationships) GetStockLineItemOk() (*POSTParcelLineItems201ResponseDataRelationshipsStockLineItem, bool) {
	if o == nil || IsNil(o.StockLineItem) {
		return nil, false
	}
	return o.StockLineItem, true
}

// HasStockLineItem returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationships) HasStockLineItem() bool {
	if o != nil && !IsNil(o.StockLineItem) {
		return true
	}

	return false
}

// SetStockLineItem gets a reference to the given POSTParcelLineItems201ResponseDataRelationshipsStockLineItem and assigns it to the StockLineItem field.
func (o *POSTStockReservations201ResponseDataRelationships) SetStockLineItem(v POSTParcelLineItems201ResponseDataRelationshipsStockLineItem) {
	o.StockLineItem = &v
}

// GetStockTransfer returns the StockTransfer field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseDataRelationships) GetStockTransfer() POSTStockReservations201ResponseDataRelationshipsStockTransfer {
	if o == nil || IsNil(o.StockTransfer) {
		var ret POSTStockReservations201ResponseDataRelationshipsStockTransfer
		return ret
	}
	return *o.StockTransfer
}

// GetStockTransferOk returns a tuple with the StockTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseDataRelationships) GetStockTransferOk() (*POSTStockReservations201ResponseDataRelationshipsStockTransfer, bool) {
	if o == nil || IsNil(o.StockTransfer) {
		return nil, false
	}
	return o.StockTransfer, true
}

// HasStockTransfer returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationships) HasStockTransfer() bool {
	if o != nil && !IsNil(o.StockTransfer) {
		return true
	}

	return false
}

// SetStockTransfer gets a reference to the given POSTStockReservations201ResponseDataRelationshipsStockTransfer and assigns it to the StockTransfer field.
func (o *POSTStockReservations201ResponseDataRelationships) SetStockTransfer(v POSTStockReservations201ResponseDataRelationshipsStockTransfer) {
	o.StockTransfer = &v
}

// GetStockItem returns the StockItem field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseDataRelationships) GetStockItem() GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem {
	if o == nil || IsNil(o.StockItem) {
		var ret GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem
		return ret
	}
	return *o.StockItem
}

// GetStockItemOk returns a tuple with the StockItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseDataRelationships) GetStockItemOk() (*GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem, bool) {
	if o == nil || IsNil(o.StockItem) {
		return nil, false
	}
	return o.StockItem, true
}

// HasStockItem returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationships) HasStockItem() bool {
	if o != nil && !IsNil(o.StockItem) {
		return true
	}

	return false
}

// SetStockItem gets a reference to the given GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem and assigns it to the StockItem field.
func (o *POSTStockReservations201ResponseDataRelationships) SetStockItem(v GETReservedStocksReservedStockId200ResponseDataRelationshipsStockItem) {
	o.StockItem = &v
}

// GetReservedStock returns the ReservedStock field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseDataRelationships) GetReservedStock() POSTStockItems201ResponseDataRelationshipsReservedStock {
	if o == nil || IsNil(o.ReservedStock) {
		var ret POSTStockItems201ResponseDataRelationshipsReservedStock
		return ret
	}
	return *o.ReservedStock
}

// GetReservedStockOk returns a tuple with the ReservedStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseDataRelationships) GetReservedStockOk() (*POSTStockItems201ResponseDataRelationshipsReservedStock, bool) {
	if o == nil || IsNil(o.ReservedStock) {
		return nil, false
	}
	return o.ReservedStock, true
}

// HasReservedStock returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationships) HasReservedStock() bool {
	if o != nil && !IsNil(o.ReservedStock) {
		return true
	}

	return false
}

// SetReservedStock gets a reference to the given POSTStockItems201ResponseDataRelationshipsReservedStock and assigns it to the ReservedStock field.
func (o *POSTStockReservations201ResponseDataRelationships) SetReservedStock(v POSTStockItems201ResponseDataRelationshipsReservedStock) {
	o.ReservedStock = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptions201ResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptions201ResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *POSTStockReservations201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku) {
	o.Sku = &v
}

func (o POSTStockReservations201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockReservations201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineItem) {
		toSerialize["line_item"] = o.LineItem
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.StockLineItem) {
		toSerialize["stock_line_item"] = o.StockLineItem
	}
	if !IsNil(o.StockTransfer) {
		toSerialize["stock_transfer"] = o.StockTransfer
	}
	if !IsNil(o.StockItem) {
		toSerialize["stock_item"] = o.StockItem
	}
	if !IsNil(o.ReservedStock) {
		toSerialize["reserved_stock"] = o.ReservedStock
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullablePOSTStockReservations201ResponseDataRelationships struct {
	value *POSTStockReservations201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTStockReservations201ResponseDataRelationships) Get() *POSTStockReservations201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTStockReservations201ResponseDataRelationships) Set(val *POSTStockReservations201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockReservations201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockReservations201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockReservations201ResponseDataRelationships(val *POSTStockReservations201ResponseDataRelationships) *NullablePOSTStockReservations201ResponseDataRelationships {
	return &NullablePOSTStockReservations201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTStockReservations201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockReservations201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
