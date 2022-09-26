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

// GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData struct for GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData
type GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData instantiates a new GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData() *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData {
	this := GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData{}
	return &this
}

// NewGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorDataWithDefaults instantiates a new GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorDataWithDefaults() *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData {
	this := GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) SetId(v string) {
	o.Id = &v
}

func (o GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData struct {
	value *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) Get() *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) Set(val *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData(val *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) *NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData {
	return &NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}