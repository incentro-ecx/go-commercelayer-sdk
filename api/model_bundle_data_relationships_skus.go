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

// BundleDataRelationshipsSkus struct for BundleDataRelationshipsSkus
type BundleDataRelationshipsSkus struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewBundleDataRelationshipsSkus instantiates a new BundleDataRelationshipsSkus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleDataRelationshipsSkus(type_ string, id string) *BundleDataRelationshipsSkus {
	this := BundleDataRelationshipsSkus{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBundleDataRelationshipsSkusWithDefaults instantiates a new BundleDataRelationshipsSkus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDataRelationshipsSkusWithDefaults() *BundleDataRelationshipsSkus {
	this := BundleDataRelationshipsSkus{}
	var type_ string = "skus"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BundleDataRelationshipsSkus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BundleDataRelationshipsSkus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleDataRelationshipsSkus) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BundleDataRelationshipsSkus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BundleDataRelationshipsSkus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BundleDataRelationshipsSkus) SetId(v string) {
	o.Id = v
}

func (o BundleDataRelationshipsSkus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBundleDataRelationshipsSkus struct {
	value *BundleDataRelationshipsSkus
	isSet bool
}

func (v NullableBundleDataRelationshipsSkus) Get() *BundleDataRelationshipsSkus {
	return v.value
}

func (v *NullableBundleDataRelationshipsSkus) Set(val *BundleDataRelationshipsSkus) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleDataRelationshipsSkus) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleDataRelationshipsSkus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleDataRelationshipsSkus(val *BundleDataRelationshipsSkus) *NullableBundleDataRelationshipsSkus {
	return &NullableBundleDataRelationshipsSkus{value: val, isSet: true}
}

func (v NullableBundleDataRelationshipsSkus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleDataRelationshipsSkus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}