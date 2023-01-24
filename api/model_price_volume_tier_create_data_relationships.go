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

// PriceVolumeTierCreateDataRelationships struct for PriceVolumeTierCreateDataRelationships
type PriceVolumeTierCreateDataRelationships struct {
	Price PriceVolumeTierCreateDataRelationshipsPrice `json:"price"`
}

// NewPriceVolumeTierCreateDataRelationships instantiates a new PriceVolumeTierCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolumeTierCreateDataRelationships(price PriceVolumeTierCreateDataRelationshipsPrice) *PriceVolumeTierCreateDataRelationships {
	this := PriceVolumeTierCreateDataRelationships{}
	this.Price = price
	return &this
}

// NewPriceVolumeTierCreateDataRelationshipsWithDefaults instantiates a new PriceVolumeTierCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeTierCreateDataRelationshipsWithDefaults() *PriceVolumeTierCreateDataRelationships {
	this := PriceVolumeTierCreateDataRelationships{}
	return &this
}

// GetPrice returns the Price field value
func (o *PriceVolumeTierCreateDataRelationships) GetPrice() PriceVolumeTierCreateDataRelationshipsPrice {
	if o == nil {
		var ret PriceVolumeTierCreateDataRelationshipsPrice
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierCreateDataRelationships) GetPriceOk() (*PriceVolumeTierCreateDataRelationshipsPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *PriceVolumeTierCreateDataRelationships) SetPrice(v PriceVolumeTierCreateDataRelationshipsPrice) {
	o.Price = v
}

func (o PriceVolumeTierCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullablePriceVolumeTierCreateDataRelationships struct {
	value *PriceVolumeTierCreateDataRelationships
	isSet bool
}

func (v NullablePriceVolumeTierCreateDataRelationships) Get() *PriceVolumeTierCreateDataRelationships {
	return v.value
}

func (v *NullablePriceVolumeTierCreateDataRelationships) Set(val *PriceVolumeTierCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolumeTierCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolumeTierCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolumeTierCreateDataRelationships(val *PriceVolumeTierCreateDataRelationships) *NullablePriceVolumeTierCreateDataRelationships {
	return &NullablePriceVolumeTierCreateDataRelationships{value: val, isSet: true}
}

func (v NullablePriceVolumeTierCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolumeTierCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
