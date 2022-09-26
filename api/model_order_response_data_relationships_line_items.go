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

// OrderResponseDataRelationshipsLineItems struct for OrderResponseDataRelationshipsLineItems
type OrderResponseDataRelationshipsLineItems struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  []OrderResponseDataRelationshipsLineItemsDataInner `json:"data,omitempty"`
}

// NewOrderResponseDataRelationshipsLineItems instantiates a new OrderResponseDataRelationshipsLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseDataRelationshipsLineItems() *OrderResponseDataRelationshipsLineItems {
	this := OrderResponseDataRelationshipsLineItems{}
	return &this
}

// NewOrderResponseDataRelationshipsLineItemsWithDefaults instantiates a new OrderResponseDataRelationshipsLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseDataRelationshipsLineItemsWithDefaults() *OrderResponseDataRelationshipsLineItems {
	this := OrderResponseDataRelationshipsLineItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsLineItems) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsLineItems) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsLineItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *OrderResponseDataRelationshipsLineItems) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsLineItems) GetData() []OrderResponseDataRelationshipsLineItemsDataInner {
	if o == nil || o.Data == nil {
		var ret []OrderResponseDataRelationshipsLineItemsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsLineItems) GetDataOk() ([]OrderResponseDataRelationshipsLineItemsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsLineItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []OrderResponseDataRelationshipsLineItemsDataInner and assigns it to the Data field.
func (o *OrderResponseDataRelationshipsLineItems) SetData(v []OrderResponseDataRelationshipsLineItemsDataInner) {
	o.Data = v
}

func (o OrderResponseDataRelationshipsLineItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponseDataRelationshipsLineItems struct {
	value *OrderResponseDataRelationshipsLineItems
	isSet bool
}

func (v NullableOrderResponseDataRelationshipsLineItems) Get() *OrderResponseDataRelationshipsLineItems {
	return v.value
}

func (v *NullableOrderResponseDataRelationshipsLineItems) Set(val *OrderResponseDataRelationshipsLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseDataRelationshipsLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseDataRelationshipsLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseDataRelationshipsLineItems(val *OrderResponseDataRelationshipsLineItems) *NullableOrderResponseDataRelationshipsLineItems {
	return &NullableOrderResponseDataRelationshipsLineItems{value: val, isSet: true}
}

func (v NullableOrderResponseDataRelationshipsLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseDataRelationshipsLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}