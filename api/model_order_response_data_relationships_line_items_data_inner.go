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

// OrderResponseDataRelationshipsLineItemsDataInner struct for OrderResponseDataRelationshipsLineItemsDataInner
type OrderResponseDataRelationshipsLineItemsDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewOrderResponseDataRelationshipsLineItemsDataInner instantiates a new OrderResponseDataRelationshipsLineItemsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseDataRelationshipsLineItemsDataInner() *OrderResponseDataRelationshipsLineItemsDataInner {
	this := OrderResponseDataRelationshipsLineItemsDataInner{}
	return &this
}

// NewOrderResponseDataRelationshipsLineItemsDataInnerWithDefaults instantiates a new OrderResponseDataRelationshipsLineItemsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseDataRelationshipsLineItemsDataInnerWithDefaults() *OrderResponseDataRelationshipsLineItemsDataInner {
	this := OrderResponseDataRelationshipsLineItemsDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsLineItemsDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsLineItemsDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsLineItemsDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderResponseDataRelationshipsLineItemsDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderResponseDataRelationshipsLineItemsDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationshipsLineItemsDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderResponseDataRelationshipsLineItemsDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderResponseDataRelationshipsLineItemsDataInner) SetId(v string) {
	o.Id = &v
}

func (o OrderResponseDataRelationshipsLineItemsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponseDataRelationshipsLineItemsDataInner struct {
	value *OrderResponseDataRelationshipsLineItemsDataInner
	isSet bool
}

func (v NullableOrderResponseDataRelationshipsLineItemsDataInner) Get() *OrderResponseDataRelationshipsLineItemsDataInner {
	return v.value
}

func (v *NullableOrderResponseDataRelationshipsLineItemsDataInner) Set(val *OrderResponseDataRelationshipsLineItemsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseDataRelationshipsLineItemsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseDataRelationshipsLineItemsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseDataRelationshipsLineItemsDataInner(val *OrderResponseDataRelationshipsLineItemsDataInner) *NullableOrderResponseDataRelationshipsLineItemsDataInner {
	return &NullableOrderResponseDataRelationshipsLineItemsDataInner{value: val, isSet: true}
}

func (v NullableOrderResponseDataRelationshipsLineItemsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseDataRelationshipsLineItemsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}