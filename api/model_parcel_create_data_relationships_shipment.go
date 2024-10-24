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

// checks if the ParcelCreateDataRelationshipsShipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParcelCreateDataRelationshipsShipment{}

// ParcelCreateDataRelationshipsShipment struct for ParcelCreateDataRelationshipsShipment
type ParcelCreateDataRelationshipsShipment struct {
	Data LineItemDataRelationshipsShipmentData `json:"data"`
}

// NewParcelCreateDataRelationshipsShipment instantiates a new ParcelCreateDataRelationshipsShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelCreateDataRelationshipsShipment(data LineItemDataRelationshipsShipmentData) *ParcelCreateDataRelationshipsShipment {
	this := ParcelCreateDataRelationshipsShipment{}
	this.Data = data
	return &this
}

// NewParcelCreateDataRelationshipsShipmentWithDefaults instantiates a new ParcelCreateDataRelationshipsShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelCreateDataRelationshipsShipmentWithDefaults() *ParcelCreateDataRelationshipsShipment {
	this := ParcelCreateDataRelationshipsShipment{}
	return &this
}

// GetData returns the Data field value
func (o *ParcelCreateDataRelationshipsShipment) GetData() LineItemDataRelationshipsShipmentData {
	if o == nil {
		var ret LineItemDataRelationshipsShipmentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParcelCreateDataRelationshipsShipment) GetDataOk() (*LineItemDataRelationshipsShipmentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ParcelCreateDataRelationshipsShipment) SetData(v LineItemDataRelationshipsShipmentData) {
	o.Data = v
}

func (o ParcelCreateDataRelationshipsShipment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParcelCreateDataRelationshipsShipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableParcelCreateDataRelationshipsShipment struct {
	value *ParcelCreateDataRelationshipsShipment
	isSet bool
}

func (v NullableParcelCreateDataRelationshipsShipment) Get() *ParcelCreateDataRelationshipsShipment {
	return v.value
}

func (v *NullableParcelCreateDataRelationshipsShipment) Set(val *ParcelCreateDataRelationshipsShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelCreateDataRelationshipsShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelCreateDataRelationshipsShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelCreateDataRelationshipsShipment(val *ParcelCreateDataRelationshipsShipment) *NullableParcelCreateDataRelationshipsShipment {
	return &NullableParcelCreateDataRelationshipsShipment{value: val, isSet: true}
}

func (v NullableParcelCreateDataRelationshipsShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelCreateDataRelationshipsShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
