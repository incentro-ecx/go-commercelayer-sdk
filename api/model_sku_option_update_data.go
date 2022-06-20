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

// SkuOptionUpdateData struct for SkuOptionUpdateData
type SkuOptionUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                      `json:"id"`
	Attributes    SkuOptionUpdateDataAttributes               `json:"attributes"`
	Relationships *BillingInfoValidationRuleDataRelationships `json:"relationships,omitempty"`
}

// NewSkuOptionUpdateData instantiates a new SkuOptionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionUpdateData(type_ string, id string, attributes SkuOptionUpdateDataAttributes) *SkuOptionUpdateData {
	this := SkuOptionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewSkuOptionUpdateDataWithDefaults instantiates a new SkuOptionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionUpdateDataWithDefaults() *SkuOptionUpdateData {
	this := SkuOptionUpdateData{}
	var type_ string = "sku_options"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SkuOptionUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuOptionUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SkuOptionUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkuOptionUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SkuOptionUpdateData) GetAttributes() SkuOptionUpdateDataAttributes {
	if o == nil {
		var ret SkuOptionUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateData) GetAttributesOk() (*SkuOptionUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuOptionUpdateData) SetAttributes(v SkuOptionUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuOptionUpdateData) GetRelationships() BillingInfoValidationRuleDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret BillingInfoValidationRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateData) GetRelationshipsOk() (*BillingInfoValidationRuleDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuOptionUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BillingInfoValidationRuleDataRelationships and assigns it to the Relationships field.
func (o *SkuOptionUpdateData) SetRelationships(v BillingInfoValidationRuleDataRelationships) {
	o.Relationships = &v
}

func (o SkuOptionUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableSkuOptionUpdateData struct {
	value *SkuOptionUpdateData
	isSet bool
}

func (v NullableSkuOptionUpdateData) Get() *SkuOptionUpdateData {
	return v.value
}

func (v *NullableSkuOptionUpdateData) Set(val *SkuOptionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionUpdateData(val *SkuOptionUpdateData) *NullableSkuOptionUpdateData {
	return &NullableSkuOptionUpdateData{value: val, isSet: true}
}

func (v NullableSkuOptionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}