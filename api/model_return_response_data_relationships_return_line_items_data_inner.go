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

// ReturnResponseDataRelationshipsReturnLineItemsDataInner struct for ReturnResponseDataRelationshipsReturnLineItemsDataInner
type ReturnResponseDataRelationshipsReturnLineItemsDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewReturnResponseDataRelationshipsReturnLineItemsDataInner instantiates a new ReturnResponseDataRelationshipsReturnLineItemsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnResponseDataRelationshipsReturnLineItemsDataInner() *ReturnResponseDataRelationshipsReturnLineItemsDataInner {
	this := ReturnResponseDataRelationshipsReturnLineItemsDataInner{}
	return &this
}

// NewReturnResponseDataRelationshipsReturnLineItemsDataInnerWithDefaults instantiates a new ReturnResponseDataRelationshipsReturnLineItemsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnResponseDataRelationshipsReturnLineItemsDataInnerWithDefaults() *ReturnResponseDataRelationshipsReturnLineItemsDataInner {
	this := ReturnResponseDataRelationshipsReturnLineItemsDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReturnResponseDataRelationshipsReturnLineItemsDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnResponseDataRelationshipsReturnLineItemsDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReturnResponseDataRelationshipsReturnLineItemsDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReturnResponseDataRelationshipsReturnLineItemsDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReturnResponseDataRelationshipsReturnLineItemsDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnResponseDataRelationshipsReturnLineItemsDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReturnResponseDataRelationshipsReturnLineItemsDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReturnResponseDataRelationshipsReturnLineItemsDataInner) SetId(v string) {
	o.Id = &v
}

func (o ReturnResponseDataRelationshipsReturnLineItemsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableReturnResponseDataRelationshipsReturnLineItemsDataInner struct {
	value *ReturnResponseDataRelationshipsReturnLineItemsDataInner
	isSet bool
}

func (v NullableReturnResponseDataRelationshipsReturnLineItemsDataInner) Get() *ReturnResponseDataRelationshipsReturnLineItemsDataInner {
	return v.value
}

func (v *NullableReturnResponseDataRelationshipsReturnLineItemsDataInner) Set(val *ReturnResponseDataRelationshipsReturnLineItemsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnResponseDataRelationshipsReturnLineItemsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnResponseDataRelationshipsReturnLineItemsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnResponseDataRelationshipsReturnLineItemsDataInner(val *ReturnResponseDataRelationshipsReturnLineItemsDataInner) *NullableReturnResponseDataRelationshipsReturnLineItemsDataInner {
	return &NullableReturnResponseDataRelationshipsReturnLineItemsDataInner{value: val, isSet: true}
}

func (v NullableReturnResponseDataRelationshipsReturnLineItemsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnResponseDataRelationshipsReturnLineItemsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}