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

// AddressDataRelationshipsGeocoder struct for AddressDataRelationshipsGeocoder
type AddressDataRelationshipsGeocoder struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewAddressDataRelationshipsGeocoder instantiates a new AddressDataRelationshipsGeocoder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressDataRelationshipsGeocoder(type_ string, id string) *AddressDataRelationshipsGeocoder {
	this := AddressDataRelationshipsGeocoder{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAddressDataRelationshipsGeocoderWithDefaults instantiates a new AddressDataRelationshipsGeocoder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataRelationshipsGeocoderWithDefaults() *AddressDataRelationshipsGeocoder {
	this := AddressDataRelationshipsGeocoder{}
	var type_ string = "geocoders"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AddressDataRelationshipsGeocoder) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressDataRelationshipsGeocoder) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressDataRelationshipsGeocoder) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AddressDataRelationshipsGeocoder) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddressDataRelationshipsGeocoder) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddressDataRelationshipsGeocoder) SetId(v string) {
	o.Id = v
}

func (o AddressDataRelationshipsGeocoder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAddressDataRelationshipsGeocoder struct {
	value *AddressDataRelationshipsGeocoder
	isSet bool
}

func (v NullableAddressDataRelationshipsGeocoder) Get() *AddressDataRelationshipsGeocoder {
	return v.value
}

func (v *NullableAddressDataRelationshipsGeocoder) Set(val *AddressDataRelationshipsGeocoder) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressDataRelationshipsGeocoder) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressDataRelationshipsGeocoder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressDataRelationshipsGeocoder(val *AddressDataRelationshipsGeocoder) *NullableAddressDataRelationshipsGeocoder {
	return &NullableAddressDataRelationshipsGeocoder{value: val, isSet: true}
}

func (v NullableAddressDataRelationshipsGeocoder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressDataRelationshipsGeocoder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
