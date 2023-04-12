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

// checks if the POSTReturns201ResponseDataRelationshipsOriginAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTReturns201ResponseDataRelationshipsOriginAddress{}

// POSTReturns201ResponseDataRelationshipsOriginAddress struct for POSTReturns201ResponseDataRelationshipsOriginAddress
type POSTReturns201ResponseDataRelationshipsOriginAddress struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *POSTReturns201ResponseDataRelationshipsOriginAddressData `json:"data,omitempty"`
}

// NewPOSTReturns201ResponseDataRelationshipsOriginAddress instantiates a new POSTReturns201ResponseDataRelationshipsOriginAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTReturns201ResponseDataRelationshipsOriginAddress() *POSTReturns201ResponseDataRelationshipsOriginAddress {
	this := POSTReturns201ResponseDataRelationshipsOriginAddress{}
	return &this
}

// NewPOSTReturns201ResponseDataRelationshipsOriginAddressWithDefaults instantiates a new POSTReturns201ResponseDataRelationshipsOriginAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTReturns201ResponseDataRelationshipsOriginAddressWithDefaults() *POSTReturns201ResponseDataRelationshipsOriginAddress {
	this := POSTReturns201ResponseDataRelationshipsOriginAddress{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationshipsOriginAddress) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationshipsOriginAddress) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationshipsOriginAddress) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTReturns201ResponseDataRelationshipsOriginAddress) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationshipsOriginAddress) GetData() POSTReturns201ResponseDataRelationshipsOriginAddressData {
	if o == nil || IsNil(o.Data) {
		var ret POSTReturns201ResponseDataRelationshipsOriginAddressData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationshipsOriginAddress) GetDataOk() (*POSTReturns201ResponseDataRelationshipsOriginAddressData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationshipsOriginAddress) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTReturns201ResponseDataRelationshipsOriginAddressData and assigns it to the Data field.
func (o *POSTReturns201ResponseDataRelationshipsOriginAddress) SetData(v POSTReturns201ResponseDataRelationshipsOriginAddressData) {
	o.Data = &v
}

func (o POSTReturns201ResponseDataRelationshipsOriginAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTReturns201ResponseDataRelationshipsOriginAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTReturns201ResponseDataRelationshipsOriginAddress struct {
	value *POSTReturns201ResponseDataRelationshipsOriginAddress
	isSet bool
}

func (v NullablePOSTReturns201ResponseDataRelationshipsOriginAddress) Get() *POSTReturns201ResponseDataRelationshipsOriginAddress {
	return v.value
}

func (v *NullablePOSTReturns201ResponseDataRelationshipsOriginAddress) Set(val *POSTReturns201ResponseDataRelationshipsOriginAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTReturns201ResponseDataRelationshipsOriginAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTReturns201ResponseDataRelationshipsOriginAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTReturns201ResponseDataRelationshipsOriginAddress(val *POSTReturns201ResponseDataRelationshipsOriginAddress) *NullablePOSTReturns201ResponseDataRelationshipsOriginAddress {
	return &NullablePOSTReturns201ResponseDataRelationshipsOriginAddress{value: val, isSet: true}
}

func (v NullablePOSTReturns201ResponseDataRelationshipsOriginAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTReturns201ResponseDataRelationshipsOriginAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}