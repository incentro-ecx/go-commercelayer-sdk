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

// checks if the MerchantUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantUpdateDataRelationships{}

// MerchantUpdateDataRelationships struct for MerchantUpdateDataRelationships
type MerchantUpdateDataRelationships struct {
	Address *CustomerAddressCreateDataRelationshipsAddress `json:"address,omitempty"`
}

// NewMerchantUpdateDataRelationships instantiates a new MerchantUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUpdateDataRelationships() *MerchantUpdateDataRelationships {
	this := MerchantUpdateDataRelationships{}
	return &this
}

// NewMerchantUpdateDataRelationshipsWithDefaults instantiates a new MerchantUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUpdateDataRelationshipsWithDefaults() *MerchantUpdateDataRelationships {
	this := MerchantUpdateDataRelationships{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MerchantUpdateDataRelationships) GetAddress() CustomerAddressCreateDataRelationshipsAddress {
	if o == nil || IsNil(o.Address) {
		var ret CustomerAddressCreateDataRelationshipsAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUpdateDataRelationships) GetAddressOk() (*CustomerAddressCreateDataRelationshipsAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MerchantUpdateDataRelationships) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CustomerAddressCreateDataRelationshipsAddress and assigns it to the Address field.
func (o *MerchantUpdateDataRelationships) SetAddress(v CustomerAddressCreateDataRelationshipsAddress) {
	o.Address = &v
}

func (o MerchantUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableMerchantUpdateDataRelationships struct {
	value *MerchantUpdateDataRelationships
	isSet bool
}

func (v NullableMerchantUpdateDataRelationships) Get() *MerchantUpdateDataRelationships {
	return v.value
}

func (v *NullableMerchantUpdateDataRelationships) Set(val *MerchantUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUpdateDataRelationships(val *MerchantUpdateDataRelationships) *NullableMerchantUpdateDataRelationships {
	return &NullableMerchantUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableMerchantUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
