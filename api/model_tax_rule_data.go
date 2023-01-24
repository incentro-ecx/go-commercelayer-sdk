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

// TaxRuleData struct for TaxRuleData
type TaxRuleData struct {
	// The resource's type
	Type          string                                    `json:"type"`
	Attributes    GETTaxRules200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *TaxRuleDataRelationships                 `json:"relationships,omitempty"`
}

// NewTaxRuleData instantiates a new TaxRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleData(type_ string, attributes GETTaxRules200ResponseDataInnerAttributes) *TaxRuleData {
	this := TaxRuleData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxRuleDataWithDefaults instantiates a new TaxRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleDataWithDefaults() *TaxRuleData {
	this := TaxRuleData{}
	return &this
}

// GetType returns the Type field value
func (o *TaxRuleData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxRuleData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxRuleData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxRuleData) GetAttributes() GETTaxRules200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETTaxRules200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxRuleData) GetAttributesOk() (*GETTaxRules200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxRuleData) SetAttributes(v GETTaxRules200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxRuleData) GetRelationships() TaxRuleDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret TaxRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleData) GetRelationshipsOk() (*TaxRuleDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxRuleData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxRuleDataRelationships and assigns it to the Relationships field.
func (o *TaxRuleData) SetRelationships(v TaxRuleDataRelationships) {
	o.Relationships = &v
}

func (o TaxRuleData) MarshalJSON() ([]byte, error) {
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

type NullableTaxRuleData struct {
	value *TaxRuleData
	isSet bool
}

func (v NullableTaxRuleData) Get() *TaxRuleData {
	return v.value
}

func (v *NullableTaxRuleData) Set(val *TaxRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleData(val *TaxRuleData) *NullableTaxRuleData {
	return &NullableTaxRuleData{value: val, isSet: true}
}

func (v NullableTaxRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
