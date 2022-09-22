/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner struct for GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner
type GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInnerWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInnerWithDefaults() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner struct {
	value *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) Get() *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) Set(val *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner(val *GETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner {
	return &NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableFreeSkusDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
