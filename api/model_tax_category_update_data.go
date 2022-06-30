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

// TaxCategoryUpdateData struct for TaxCategoryUpdateData
type TaxCategoryUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                              `json:"id"`
	Attributes    TaxCategoryUpdateDataAttributes     `json:"attributes"`
	Relationships *TaxCategoryUpdateDataRelationships `json:"relationships,omitempty"`
}

// NewTaxCategoryUpdateData instantiates a new TaxCategoryUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryUpdateData(type_ string, id string, attributes TaxCategoryUpdateDataAttributes) *TaxCategoryUpdateData {
	this := TaxCategoryUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewTaxCategoryUpdateDataWithDefaults instantiates a new TaxCategoryUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryUpdateDataWithDefaults() *TaxCategoryUpdateData {
	this := TaxCategoryUpdateData{}
	var type_ string = "tax_categories"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TaxCategoryUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxCategoryUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TaxCategoryUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaxCategoryUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *TaxCategoryUpdateData) GetAttributes() TaxCategoryUpdateDataAttributes {
	if o == nil {
		var ret TaxCategoryUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryUpdateData) GetAttributesOk() (*TaxCategoryUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxCategoryUpdateData) SetAttributes(v TaxCategoryUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxCategoryUpdateData) GetRelationships() TaxCategoryUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret TaxCategoryUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryUpdateData) GetRelationshipsOk() (*TaxCategoryUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxCategoryUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxCategoryUpdateDataRelationships and assigns it to the Relationships field.
func (o *TaxCategoryUpdateData) SetRelationships(v TaxCategoryUpdateDataRelationships) {
	o.Relationships = &v
}

func (o TaxCategoryUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableTaxCategoryUpdateData struct {
	value *TaxCategoryUpdateData
	isSet bool
}

func (v NullableTaxCategoryUpdateData) Get() *TaxCategoryUpdateData {
	return v.value
}

func (v *NullableTaxCategoryUpdateData) Set(val *TaxCategoryUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryUpdateData(val *TaxCategoryUpdateData) *NullableTaxCategoryUpdateData {
	return &NullableTaxCategoryUpdateData{value: val, isSet: true}
}

func (v NullableTaxCategoryUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
