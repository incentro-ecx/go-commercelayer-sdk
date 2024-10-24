/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TaxRuleCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRuleCreateData{}

// TaxRuleCreateData struct for TaxRuleCreateData
type TaxRuleCreateData struct {
	// The resource's type
	Type          interface{}                           `json:"type"`
	Attributes    POSTTaxRules201ResponseDataAttributes `json:"attributes"`
	Relationships *TaxRuleCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewTaxRuleCreateData instantiates a new TaxRuleCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleCreateData(type_ interface{}, attributes POSTTaxRules201ResponseDataAttributes) *TaxRuleCreateData {
	this := TaxRuleCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxRuleCreateDataWithDefaults instantiates a new TaxRuleCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleCreateDataWithDefaults() *TaxRuleCreateData {
	this := TaxRuleCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TaxRuleCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxRuleCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxRuleCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxRuleCreateData) GetAttributes() POSTTaxRules201ResponseDataAttributes {
	if o == nil {
		var ret POSTTaxRules201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxRuleCreateData) GetAttributesOk() (*POSTTaxRules201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxRuleCreateData) SetAttributes(v POSTTaxRules201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxRuleCreateData) GetRelationships() TaxRuleCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret TaxRuleCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleCreateData) GetRelationshipsOk() (*TaxRuleCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxRuleCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxRuleCreateDataRelationships and assigns it to the Relationships field.
func (o *TaxRuleCreateData) SetRelationships(v TaxRuleCreateDataRelationships) {
	o.Relationships = &v
}

func (o TaxRuleCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxRuleCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableTaxRuleCreateData struct {
	value *TaxRuleCreateData
	isSet bool
}

func (v NullableTaxRuleCreateData) Get() *TaxRuleCreateData {
	return v.value
}

func (v *NullableTaxRuleCreateData) Set(val *TaxRuleCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleCreateData(val *TaxRuleCreateData) *NullableTaxRuleCreateData {
	return &NullableTaxRuleCreateData{value: val, isSet: true}
}

func (v NullableTaxRuleCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
