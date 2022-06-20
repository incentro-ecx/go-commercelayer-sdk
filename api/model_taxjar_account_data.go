/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TaxjarAccountData struct for TaxjarAccountData
type TaxjarAccountData struct {
	// The resource's type
	Type          string                            `json:"type"`
	Attributes    ManualTaxCalculatorDataAttributes `json:"attributes"`
	Relationships *AvalaraAccountDataRelationships  `json:"relationships,omitempty"`
}

// NewTaxjarAccountData instantiates a new TaxjarAccountData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxjarAccountData(type_ string, attributes ManualTaxCalculatorDataAttributes) *TaxjarAccountData {
	this := TaxjarAccountData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxjarAccountDataWithDefaults instantiates a new TaxjarAccountData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxjarAccountDataWithDefaults() *TaxjarAccountData {
	this := TaxjarAccountData{}
	var type_ string = "taxjar_accounts"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TaxjarAccountData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxjarAccountData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxjarAccountData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxjarAccountData) GetAttributes() ManualTaxCalculatorDataAttributes {
	if o == nil {
		var ret ManualTaxCalculatorDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxjarAccountData) GetAttributesOk() (*ManualTaxCalculatorDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxjarAccountData) SetAttributes(v ManualTaxCalculatorDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxjarAccountData) GetRelationships() AvalaraAccountDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AvalaraAccountDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxjarAccountData) GetRelationshipsOk() (*AvalaraAccountDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxjarAccountData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountDataRelationships and assigns it to the Relationships field.
func (o *TaxjarAccountData) SetRelationships(v AvalaraAccountDataRelationships) {
	o.Relationships = &v
}

func (o TaxjarAccountData) MarshalJSON() ([]byte, error) {
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

type NullableTaxjarAccountData struct {
	value *TaxjarAccountData
	isSet bool
}

func (v NullableTaxjarAccountData) Get() *TaxjarAccountData {
	return v.value
}

func (v *NullableTaxjarAccountData) Set(val *TaxjarAccountData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxjarAccountData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxjarAccountData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxjarAccountData(val *TaxjarAccountData) *NullableTaxjarAccountData {
	return &NullableTaxjarAccountData{value: val, isSet: true}
}

func (v NullableTaxjarAccountData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxjarAccountData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}