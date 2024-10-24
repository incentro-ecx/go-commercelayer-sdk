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

// checks if the AvalaraAccountCreateDataRelationshipsTaxCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvalaraAccountCreateDataRelationshipsTaxCategories{}

// AvalaraAccountCreateDataRelationshipsTaxCategories struct for AvalaraAccountCreateDataRelationshipsTaxCategories
type AvalaraAccountCreateDataRelationshipsTaxCategories struct {
	Data AvalaraAccountDataRelationshipsTaxCategoriesData `json:"data"`
}

// NewAvalaraAccountCreateDataRelationshipsTaxCategories instantiates a new AvalaraAccountCreateDataRelationshipsTaxCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountCreateDataRelationshipsTaxCategories(data AvalaraAccountDataRelationshipsTaxCategoriesData) *AvalaraAccountCreateDataRelationshipsTaxCategories {
	this := AvalaraAccountCreateDataRelationshipsTaxCategories{}
	this.Data = data
	return &this
}

// NewAvalaraAccountCreateDataRelationshipsTaxCategoriesWithDefaults instantiates a new AvalaraAccountCreateDataRelationshipsTaxCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountCreateDataRelationshipsTaxCategoriesWithDefaults() *AvalaraAccountCreateDataRelationshipsTaxCategories {
	this := AvalaraAccountCreateDataRelationshipsTaxCategories{}
	return &this
}

// GetData returns the Data field value
func (o *AvalaraAccountCreateDataRelationshipsTaxCategories) GetData() AvalaraAccountDataRelationshipsTaxCategoriesData {
	if o == nil {
		var ret AvalaraAccountDataRelationshipsTaxCategoriesData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccountCreateDataRelationshipsTaxCategories) GetDataOk() (*AvalaraAccountDataRelationshipsTaxCategoriesData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AvalaraAccountCreateDataRelationshipsTaxCategories) SetData(v AvalaraAccountDataRelationshipsTaxCategoriesData) {
	o.Data = v
}

func (o AvalaraAccountCreateDataRelationshipsTaxCategories) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvalaraAccountCreateDataRelationshipsTaxCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAvalaraAccountCreateDataRelationshipsTaxCategories struct {
	value *AvalaraAccountCreateDataRelationshipsTaxCategories
	isSet bool
}

func (v NullableAvalaraAccountCreateDataRelationshipsTaxCategories) Get() *AvalaraAccountCreateDataRelationshipsTaxCategories {
	return v.value
}

func (v *NullableAvalaraAccountCreateDataRelationshipsTaxCategories) Set(val *AvalaraAccountCreateDataRelationshipsTaxCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountCreateDataRelationshipsTaxCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountCreateDataRelationshipsTaxCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountCreateDataRelationshipsTaxCategories(val *AvalaraAccountCreateDataRelationshipsTaxCategories) *NullableAvalaraAccountCreateDataRelationshipsTaxCategories {
	return &NullableAvalaraAccountCreateDataRelationshipsTaxCategories{value: val, isSet: true}
}

func (v NullableAvalaraAccountCreateDataRelationshipsTaxCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountCreateDataRelationshipsTaxCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
