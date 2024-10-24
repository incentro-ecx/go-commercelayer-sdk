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

// checks if the SkuListPromotionRuleCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuListPromotionRuleCreateData{}

// SkuListPromotionRuleCreateData struct for SkuListPromotionRuleCreateData
type SkuListPromotionRuleCreateData struct {
	// The resource's type
	Type          interface{}                                        `json:"type"`
	Attributes    POSTSkuListPromotionRules201ResponseDataAttributes `json:"attributes"`
	Relationships *SkuListPromotionRuleCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewSkuListPromotionRuleCreateData instantiates a new SkuListPromotionRuleCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleCreateData(type_ interface{}, attributes POSTSkuListPromotionRules201ResponseDataAttributes) *SkuListPromotionRuleCreateData {
	this := SkuListPromotionRuleCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuListPromotionRuleCreateDataWithDefaults instantiates a new SkuListPromotionRuleCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleCreateDataWithDefaults() *SkuListPromotionRuleCreateData {
	this := SkuListPromotionRuleCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuListPromotionRuleCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuListPromotionRuleCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListPromotionRuleCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListPromotionRuleCreateData) GetAttributes() POSTSkuListPromotionRules201ResponseDataAttributes {
	if o == nil {
		var ret POSTSkuListPromotionRules201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleCreateData) GetAttributesOk() (*POSTSkuListPromotionRules201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListPromotionRuleCreateData) SetAttributes(v POSTSkuListPromotionRules201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListPromotionRuleCreateData) GetRelationships() SkuListPromotionRuleCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SkuListPromotionRuleCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleCreateData) GetRelationshipsOk() (*SkuListPromotionRuleCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListPromotionRuleCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuListPromotionRuleCreateDataRelationships and assigns it to the Relationships field.
func (o *SkuListPromotionRuleCreateData) SetRelationships(v SkuListPromotionRuleCreateDataRelationships) {
	o.Relationships = &v
}

func (o SkuListPromotionRuleCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuListPromotionRuleCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableSkuListPromotionRuleCreateData struct {
	value *SkuListPromotionRuleCreateData
	isSet bool
}

func (v NullableSkuListPromotionRuleCreateData) Get() *SkuListPromotionRuleCreateData {
	return v.value
}

func (v *NullableSkuListPromotionRuleCreateData) Set(val *SkuListPromotionRuleCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleCreateData(val *SkuListPromotionRuleCreateData) *NullableSkuListPromotionRuleCreateData {
	return &NullableSkuListPromotionRuleCreateData{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
