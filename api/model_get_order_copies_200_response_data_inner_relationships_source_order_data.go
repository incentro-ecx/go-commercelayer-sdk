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

// GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData struct for GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData
type GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData instantiates a new GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData() *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData {
	this := GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData{}
	return &this
}

// NewGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderDataWithDefaults instantiates a new GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderDataWithDefaults() *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData {
	this := GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) SetId(v string) {
	o.Id = &v
}

func (o GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData struct {
	value *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData
	isSet bool
}

func (v NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) Get() *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData {
	return v.value
}

func (v *NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) Set(val *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData(val *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) *NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData {
	return &NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData{value: val, isSet: true}
}

func (v NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderCopies200ResponseDataInnerRelationshipsSourceOrderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
