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

// ExternalTaxCalculatorUpdateData struct for ExternalTaxCalculatorUpdateData
type ExternalTaxCalculatorUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                    `json:"id"`
	Attributes    ExternalTaxCalculatorUpdateDataAttributes `json:"attributes"`
	Relationships *AvalaraAccountCreateDataRelationships    `json:"relationships,omitempty"`
}

// NewExternalTaxCalculatorUpdateData instantiates a new ExternalTaxCalculatorUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaxCalculatorUpdateData(type_ string, id string, attributes ExternalTaxCalculatorUpdateDataAttributes) *ExternalTaxCalculatorUpdateData {
	this := ExternalTaxCalculatorUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewExternalTaxCalculatorUpdateDataWithDefaults instantiates a new ExternalTaxCalculatorUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaxCalculatorUpdateDataWithDefaults() *ExternalTaxCalculatorUpdateData {
	this := ExternalTaxCalculatorUpdateData{}
	var type_ string = "external_tax_calculators"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ExternalTaxCalculatorUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalTaxCalculatorUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ExternalTaxCalculatorUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalTaxCalculatorUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalTaxCalculatorUpdateData) GetAttributes() ExternalTaxCalculatorUpdateDataAttributes {
	if o == nil {
		var ret ExternalTaxCalculatorUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorUpdateData) GetAttributesOk() (*ExternalTaxCalculatorUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalTaxCalculatorUpdateData) SetAttributes(v ExternalTaxCalculatorUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalTaxCalculatorUpdateData) GetRelationships() AvalaraAccountCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AvalaraAccountCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorUpdateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalTaxCalculatorUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountCreateDataRelationships and assigns it to the Relationships field.
func (o *ExternalTaxCalculatorUpdateData) SetRelationships(v AvalaraAccountCreateDataRelationships) {
	o.Relationships = &v
}

func (o ExternalTaxCalculatorUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableExternalTaxCalculatorUpdateData struct {
	value *ExternalTaxCalculatorUpdateData
	isSet bool
}

func (v NullableExternalTaxCalculatorUpdateData) Get() *ExternalTaxCalculatorUpdateData {
	return v.value
}

func (v *NullableExternalTaxCalculatorUpdateData) Set(val *ExternalTaxCalculatorUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaxCalculatorUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaxCalculatorUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaxCalculatorUpdateData(val *ExternalTaxCalculatorUpdateData) *NullableExternalTaxCalculatorUpdateData {
	return &NullableExternalTaxCalculatorUpdateData{value: val, isSet: true}
}

func (v NullableExternalTaxCalculatorUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaxCalculatorUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
