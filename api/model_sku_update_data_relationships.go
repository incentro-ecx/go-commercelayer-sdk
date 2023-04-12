/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SkuUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuUpdateDataRelationships{}

// SkuUpdateDataRelationships struct for SkuUpdateDataRelationships
type SkuUpdateDataRelationships struct {
	ShippingCategory *ShippingMethodCreateDataRelationshipsShippingCategory `json:"shipping_category,omitempty"`
}

// NewSkuUpdateDataRelationships instantiates a new SkuUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuUpdateDataRelationships() *SkuUpdateDataRelationships {
	this := SkuUpdateDataRelationships{}
	return &this
}

// NewSkuUpdateDataRelationshipsWithDefaults instantiates a new SkuUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuUpdateDataRelationshipsWithDefaults() *SkuUpdateDataRelationships {
	this := SkuUpdateDataRelationships{}
	return &this
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *SkuUpdateDataRelationships) GetShippingCategory() ShippingMethodCreateDataRelationshipsShippingCategory {
	if o == nil || IsNil(o.ShippingCategory) {
		var ret ShippingMethodCreateDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataRelationships) GetShippingCategoryOk() (*ShippingMethodCreateDataRelationshipsShippingCategory, bool) {
	if o == nil || IsNil(o.ShippingCategory) {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *SkuUpdateDataRelationships) HasShippingCategory() bool {
	if o != nil && !IsNil(o.ShippingCategory) {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShippingMethodCreateDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *SkuUpdateDataRelationships) SetShippingCategory(v ShippingMethodCreateDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

func (o SkuUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShippingCategory) {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	return toSerialize, nil
}

type NullableSkuUpdateDataRelationships struct {
	value *SkuUpdateDataRelationships
	isSet bool
}

func (v NullableSkuUpdateDataRelationships) Get() *SkuUpdateDataRelationships {
	return v.value
}

func (v *NullableSkuUpdateDataRelationships) Set(val *SkuUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuUpdateDataRelationships(val *SkuUpdateDataRelationships) *NullableSkuUpdateDataRelationships {
	return &NullableSkuUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableSkuUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
