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

// OrderResponseDataRelationshipsShipments struct for OrderResponseDataRelationshipsShipments
type OrderResponseDataRelationshipsShipments struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  []OrderResponseDataRelationshipsShipmentsDataInner `json:"data,omitempty"`
}

// NewOrderResponseDataRelationshipsShipments instantiates a new OrderResponseDataRelationshipsShipments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseDataRelationshipsShipments() *OrderResponseDataRelationshipsShipments {
	this := OrderResponseDataRelationshipsShipments{}
	return &this
}

// NewOrderResponseDataRelationshipsShipmentsWithDefaults instantiates a new OrderResponseDataRelationshipsShipments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseDataRelationshipsShipmentsWithDefaults() *OrderResponseDataRelationshipsShipments {
	this := OrderResponseDataRelationshipsShipments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsShipments) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsShipments) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsShipments) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *OrderResponseDataRelationshipsShipments) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsShipments) GetData() []OrderResponseDataRelationshipsShipmentsDataInner {
	if o == nil || o.Data == nil {
		var ret []OrderResponseDataRelationshipsShipmentsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsShipments) GetDataOk() ([]OrderResponseDataRelationshipsShipmentsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsShipments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []OrderResponseDataRelationshipsShipmentsDataInner and assigns it to the Data field.
func (o *OrderResponseDataRelationshipsShipments) SetData(v []OrderResponseDataRelationshipsShipmentsDataInner) {
	o.Data = v
}

func (o OrderResponseDataRelationshipsShipments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponseDataRelationshipsShipments struct {
	value *OrderResponseDataRelationshipsShipments
	isSet bool
}

func (v NullableOrderResponseDataRelationshipsShipments) Get() *OrderResponseDataRelationshipsShipments {
	return v.value
}

func (v *NullableOrderResponseDataRelationshipsShipments) Set(val *OrderResponseDataRelationshipsShipments) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseDataRelationshipsShipments) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseDataRelationshipsShipments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseDataRelationshipsShipments(val *OrderResponseDataRelationshipsShipments) *NullableOrderResponseDataRelationshipsShipments {
	return &NullableOrderResponseDataRelationshipsShipments{value: val, isSet: true}
}

func (v NullableOrderResponseDataRelationshipsShipments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseDataRelationshipsShipments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}