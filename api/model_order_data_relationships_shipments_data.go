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

// OrderDataRelationshipsShipmentsData struct for OrderDataRelationshipsShipmentsData
type OrderDataRelationshipsShipmentsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource's id
	Id *string `json:"id,omitempty"`
}

// NewOrderDataRelationshipsShipmentsData instantiates a new OrderDataRelationshipsShipmentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataRelationshipsShipmentsData() *OrderDataRelationshipsShipmentsData {
	this := OrderDataRelationshipsShipmentsData{}
	return &this
}

// NewOrderDataRelationshipsShipmentsDataWithDefaults instantiates a new OrderDataRelationshipsShipmentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataRelationshipsShipmentsDataWithDefaults() *OrderDataRelationshipsShipmentsData {
	this := OrderDataRelationshipsShipmentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderDataRelationshipsShipmentsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationshipsShipmentsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderDataRelationshipsShipmentsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderDataRelationshipsShipmentsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderDataRelationshipsShipmentsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationshipsShipmentsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderDataRelationshipsShipmentsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderDataRelationshipsShipmentsData) SetId(v string) {
	o.Id = &v
}

func (o OrderDataRelationshipsShipmentsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrderDataRelationshipsShipmentsData struct {
	value *OrderDataRelationshipsShipmentsData
	isSet bool
}

func (v NullableOrderDataRelationshipsShipmentsData) Get() *OrderDataRelationshipsShipmentsData {
	return v.value
}

func (v *NullableOrderDataRelationshipsShipmentsData) Set(val *OrderDataRelationshipsShipmentsData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataRelationshipsShipmentsData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataRelationshipsShipmentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataRelationshipsShipmentsData(val *OrderDataRelationshipsShipmentsData) *NullableOrderDataRelationshipsShipmentsData {
	return &NullableOrderDataRelationshipsShipmentsData{value: val, isSet: true}
}

func (v NullableOrderDataRelationshipsShipmentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataRelationshipsShipmentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
