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

// GETReturnLineItems200ResponseDataInnerRelationshipsReturnData struct for GETReturnLineItems200ResponseDataInnerRelationshipsReturnData
type GETReturnLineItems200ResponseDataInnerRelationshipsReturnData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETReturnLineItems200ResponseDataInnerRelationshipsReturnData instantiates a new GETReturnLineItems200ResponseDataInnerRelationshipsReturnData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturnLineItems200ResponseDataInnerRelationshipsReturnData() *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData {
	this := GETReturnLineItems200ResponseDataInnerRelationshipsReturnData{}
	return &this
}

// NewGETReturnLineItems200ResponseDataInnerRelationshipsReturnDataWithDefaults instantiates a new GETReturnLineItems200ResponseDataInnerRelationshipsReturnData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturnLineItems200ResponseDataInnerRelationshipsReturnDataWithDefaults() *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData {
	this := GETReturnLineItems200ResponseDataInnerRelationshipsReturnData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) SetId(v string) {
	o.Id = &v
}

func (o GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData struct {
	value *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData
	isSet bool
}

func (v NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData) Get() *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData {
	return v.value
}

func (v *NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData) Set(val *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData(val *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) *NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData {
	return &NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData{value: val, isSet: true}
}

func (v NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturnData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
