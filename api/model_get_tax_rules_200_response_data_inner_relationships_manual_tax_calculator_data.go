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

// GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData struct for GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData
type GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData instantiates a new GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData() *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData {
	this := GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData{}
	return &this
}

// NewGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorDataWithDefaults instantiates a new GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorDataWithDefaults() *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData {
	this := GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) SetId(v string) {
	o.Id = &v
}

func (o GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData struct {
	value *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData
	isSet bool
}

func (v NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) Get() *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData {
	return v.value
}

func (v *NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) Set(val *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData(val *GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) *NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData {
	return &NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData{value: val, isSet: true}
}

func (v NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
