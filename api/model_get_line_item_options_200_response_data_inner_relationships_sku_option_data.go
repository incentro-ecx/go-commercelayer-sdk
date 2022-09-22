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

// GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData struct for GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData
type GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData instantiates a new GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData() *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData {
	this := GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData{}
	return &this
}

// NewGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionDataWithDefaults instantiates a new GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionDataWithDefaults() *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData {
	this := GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) SetId(v string) {
	o.Id = &v
}

func (o GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData struct {
	value *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData
	isSet bool
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) Get() *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData {
	return v.value
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) Set(val *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData(val *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) *NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData {
	return &NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData{value: val, isSet: true}
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsSkuOptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
