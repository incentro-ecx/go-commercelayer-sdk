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

// PriceDataRelationshipsPriceTiers struct for PriceDataRelationshipsPriceTiers
type PriceDataRelationshipsPriceTiers struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewPriceDataRelationshipsPriceTiers instantiates a new PriceDataRelationshipsPriceTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceDataRelationshipsPriceTiers(type_ string, id string) *PriceDataRelationshipsPriceTiers {
	this := PriceDataRelationshipsPriceTiers{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewPriceDataRelationshipsPriceTiersWithDefaults instantiates a new PriceDataRelationshipsPriceTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceDataRelationshipsPriceTiersWithDefaults() *PriceDataRelationshipsPriceTiers {
	this := PriceDataRelationshipsPriceTiers{}
	var type_ string = "price_tiers"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PriceDataRelationshipsPriceTiers) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceDataRelationshipsPriceTiers) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceDataRelationshipsPriceTiers) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PriceDataRelationshipsPriceTiers) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PriceDataRelationshipsPriceTiers) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PriceDataRelationshipsPriceTiers) SetId(v string) {
	o.Id = v
}

func (o PriceDataRelationshipsPriceTiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePriceDataRelationshipsPriceTiers struct {
	value *PriceDataRelationshipsPriceTiers
	isSet bool
}

func (v NullablePriceDataRelationshipsPriceTiers) Get() *PriceDataRelationshipsPriceTiers {
	return v.value
}

func (v *NullablePriceDataRelationshipsPriceTiers) Set(val *PriceDataRelationshipsPriceTiers) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceDataRelationshipsPriceTiers) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceDataRelationshipsPriceTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceDataRelationshipsPriceTiers(val *PriceDataRelationshipsPriceTiers) *NullablePriceDataRelationshipsPriceTiers {
	return &NullablePriceDataRelationshipsPriceTiers{value: val, isSet: true}
}

func (v NullablePriceDataRelationshipsPriceTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceDataRelationshipsPriceTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}