/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ParcelCreateDataRelationshipsPackage struct for ParcelCreateDataRelationshipsPackage
type ParcelCreateDataRelationshipsPackage struct {
	Data ParcelDataRelationshipsPackageData `json:"data"`
}

// NewParcelCreateDataRelationshipsPackage instantiates a new ParcelCreateDataRelationshipsPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelCreateDataRelationshipsPackage(data ParcelDataRelationshipsPackageData) *ParcelCreateDataRelationshipsPackage {
	this := ParcelCreateDataRelationshipsPackage{}
	this.Data = data
	return &this
}

// NewParcelCreateDataRelationshipsPackageWithDefaults instantiates a new ParcelCreateDataRelationshipsPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelCreateDataRelationshipsPackageWithDefaults() *ParcelCreateDataRelationshipsPackage {
	this := ParcelCreateDataRelationshipsPackage{}
	return &this
}

// GetData returns the Data field value
func (o *ParcelCreateDataRelationshipsPackage) GetData() ParcelDataRelationshipsPackageData {
	if o == nil {
		var ret ParcelDataRelationshipsPackageData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParcelCreateDataRelationshipsPackage) GetDataOk() (*ParcelDataRelationshipsPackageData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ParcelCreateDataRelationshipsPackage) SetData(v ParcelDataRelationshipsPackageData) {
	o.Data = v
}

func (o ParcelCreateDataRelationshipsPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableParcelCreateDataRelationshipsPackage struct {
	value *ParcelCreateDataRelationshipsPackage
	isSet bool
}

func (v NullableParcelCreateDataRelationshipsPackage) Get() *ParcelCreateDataRelationshipsPackage {
	return v.value
}

func (v *NullableParcelCreateDataRelationshipsPackage) Set(val *ParcelCreateDataRelationshipsPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelCreateDataRelationshipsPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelCreateDataRelationshipsPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelCreateDataRelationshipsPackage(val *ParcelCreateDataRelationshipsPackage) *NullableParcelCreateDataRelationshipsPackage {
	return &NullableParcelCreateDataRelationshipsPackage{value: val, isSet: true}
}

func (v NullableParcelCreateDataRelationshipsPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelCreateDataRelationshipsPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}