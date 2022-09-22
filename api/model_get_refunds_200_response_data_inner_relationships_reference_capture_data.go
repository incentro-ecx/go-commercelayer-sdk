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

// GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData struct for GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData
type GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData instantiates a new GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData() *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData {
	this := GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData{}
	return &this
}

// NewGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureDataWithDefaults instantiates a new GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureDataWithDefaults() *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData {
	this := GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) SetId(v string) {
	o.Id = &v
}

func (o GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData struct {
	value *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData
	isSet bool
}

func (v NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) Get() *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData {
	return v.value
}

func (v *NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) Set(val *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData(val *GETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) *NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData {
	return &NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData{value: val, isSet: true}
}

func (v NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETRefunds200ResponseDataInnerRelationshipsReferenceCaptureData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
