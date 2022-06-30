/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AdjustmentData struct for AdjustmentData
type AdjustmentData struct {
	// The resource's type
	Type          string                   `json:"type"`
	Attributes    AdjustmentDataAttributes `json:"attributes"`
	Relationships map[string]interface{}   `json:"relationships,omitempty"`
}

// NewAdjustmentData instantiates a new AdjustmentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentData(type_ string, attributes AdjustmentDataAttributes) *AdjustmentData {
	this := AdjustmentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAdjustmentDataWithDefaults instantiates a new AdjustmentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentDataWithDefaults() *AdjustmentData {
	this := AdjustmentData{}
	var type_ string = "adjustments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AdjustmentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdjustmentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdjustmentData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AdjustmentData) GetAttributes() AdjustmentDataAttributes {
	if o == nil {
		var ret AdjustmentDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdjustmentData) GetAttributesOk() (*AdjustmentDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdjustmentData) SetAttributes(v AdjustmentDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdjustmentData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdjustmentData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *AdjustmentData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o AdjustmentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableAdjustmentData struct {
	value *AdjustmentData
	isSet bool
}

func (v NullableAdjustmentData) Get() *AdjustmentData {
	return v.value
}

func (v *NullableAdjustmentData) Set(val *AdjustmentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentData(val *AdjustmentData) *NullableAdjustmentData {
	return &NullableAdjustmentData{value: val, isSet: true}
}

func (v NullableAdjustmentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustmentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
