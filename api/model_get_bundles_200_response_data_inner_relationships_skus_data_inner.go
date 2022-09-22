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

// GETBundles200ResponseDataInnerRelationshipsSkusDataInner struct for GETBundles200ResponseDataInnerRelationshipsSkusDataInner
type GETBundles200ResponseDataInnerRelationshipsSkusDataInner struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETBundles200ResponseDataInnerRelationshipsSkusDataInner instantiates a new GETBundles200ResponseDataInnerRelationshipsSkusDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBundles200ResponseDataInnerRelationshipsSkusDataInner() *GETBundles200ResponseDataInnerRelationshipsSkusDataInner {
	this := GETBundles200ResponseDataInnerRelationshipsSkusDataInner{}
	return &this
}

// NewGETBundles200ResponseDataInnerRelationshipsSkusDataInnerWithDefaults instantiates a new GETBundles200ResponseDataInnerRelationshipsSkusDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBundles200ResponseDataInnerRelationshipsSkusDataInnerWithDefaults() *GETBundles200ResponseDataInnerRelationshipsSkusDataInner {
	this := GETBundles200ResponseDataInnerRelationshipsSkusDataInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) SetId(v string) {
	o.Id = &v
}

func (o GETBundles200ResponseDataInnerRelationshipsSkusDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner struct {
	value *GETBundles200ResponseDataInnerRelationshipsSkusDataInner
	isSet bool
}

func (v NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner) Get() *GETBundles200ResponseDataInnerRelationshipsSkusDataInner {
	return v.value
}

func (v *NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner) Set(val *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner(val *GETBundles200ResponseDataInnerRelationshipsSkusDataInner) *NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner {
	return &NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner{value: val, isSet: true}
}

func (v NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBundles200ResponseDataInnerRelationshipsSkusDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
