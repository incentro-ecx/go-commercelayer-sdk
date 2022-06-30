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

// CarrierAccountData struct for CarrierAccountData
type CarrierAccountData struct {
	// The resource's type
	Type          string                           `json:"type"`
	Attributes    CarrierAccountDataAttributes     `json:"attributes"`
	Relationships *CarrierAccountDataRelationships `json:"relationships,omitempty"`
}

// NewCarrierAccountData instantiates a new CarrierAccountData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierAccountData(type_ string, attributes CarrierAccountDataAttributes) *CarrierAccountData {
	this := CarrierAccountData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCarrierAccountDataWithDefaults instantiates a new CarrierAccountData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierAccountDataWithDefaults() *CarrierAccountData {
	this := CarrierAccountData{}
	var type_ string = "carrier_accounts"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CarrierAccountData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CarrierAccountData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CarrierAccountData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CarrierAccountData) GetAttributes() CarrierAccountDataAttributes {
	if o == nil {
		var ret CarrierAccountDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CarrierAccountData) GetAttributesOk() (*CarrierAccountDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CarrierAccountData) SetAttributes(v CarrierAccountDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CarrierAccountData) GetRelationships() CarrierAccountDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CarrierAccountDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierAccountData) GetRelationshipsOk() (*CarrierAccountDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CarrierAccountData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CarrierAccountDataRelationships and assigns it to the Relationships field.
func (o *CarrierAccountData) SetRelationships(v CarrierAccountDataRelationships) {
	o.Relationships = &v
}

func (o CarrierAccountData) MarshalJSON() ([]byte, error) {
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

type NullableCarrierAccountData struct {
	value *CarrierAccountData
	isSet bool
}

func (v NullableCarrierAccountData) Get() *CarrierAccountData {
	return v.value
}

func (v *NullableCarrierAccountData) Set(val *CarrierAccountData) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierAccountData) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierAccountData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierAccountData(val *CarrierAccountData) *NullableCarrierAccountData {
	return &NullableCarrierAccountData{value: val, isSet: true}
}

func (v NullableCarrierAccountData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrierAccountData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
