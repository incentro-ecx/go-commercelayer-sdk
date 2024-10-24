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

// checks if the POSTPrices201ResponseDataRelationshipsPriceTiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPrices201ResponseDataRelationshipsPriceTiers{}

// POSTPrices201ResponseDataRelationshipsPriceTiers struct for POSTPrices201ResponseDataRelationshipsPriceTiers
type POSTPrices201ResponseDataRelationshipsPriceTiers struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTPrices201ResponseDataRelationshipsPriceTiersData   `json:"data,omitempty"`
}

// NewPOSTPrices201ResponseDataRelationshipsPriceTiers instantiates a new POSTPrices201ResponseDataRelationshipsPriceTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPrices201ResponseDataRelationshipsPriceTiers() *POSTPrices201ResponseDataRelationshipsPriceTiers {
	this := POSTPrices201ResponseDataRelationshipsPriceTiers{}
	return &this
}

// NewPOSTPrices201ResponseDataRelationshipsPriceTiersWithDefaults instantiates a new POSTPrices201ResponseDataRelationshipsPriceTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPrices201ResponseDataRelationshipsPriceTiersWithDefaults() *POSTPrices201ResponseDataRelationshipsPriceTiers {
	this := POSTPrices201ResponseDataRelationshipsPriceTiers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiers) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiers) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiers) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiers) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiers) GetData() POSTPrices201ResponseDataRelationshipsPriceTiersData {
	if o == nil || IsNil(o.Data) {
		var ret POSTPrices201ResponseDataRelationshipsPriceTiersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiers) GetDataOk() (*POSTPrices201ResponseDataRelationshipsPriceTiersData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPrices201ResponseDataRelationshipsPriceTiersData and assigns it to the Data field.
func (o *POSTPrices201ResponseDataRelationshipsPriceTiers) SetData(v POSTPrices201ResponseDataRelationshipsPriceTiersData) {
	o.Data = &v
}

func (o POSTPrices201ResponseDataRelationshipsPriceTiers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPrices201ResponseDataRelationshipsPriceTiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTPrices201ResponseDataRelationshipsPriceTiers struct {
	value *POSTPrices201ResponseDataRelationshipsPriceTiers
	isSet bool
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceTiers) Get() *POSTPrices201ResponseDataRelationshipsPriceTiers {
	return v.value
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceTiers) Set(val *POSTPrices201ResponseDataRelationshipsPriceTiers) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceTiers) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPrices201ResponseDataRelationshipsPriceTiers(val *POSTPrices201ResponseDataRelationshipsPriceTiers) *NullablePOSTPrices201ResponseDataRelationshipsPriceTiers {
	return &NullablePOSTPrices201ResponseDataRelationshipsPriceTiers{value: val, isSet: true}
}

func (v NullablePOSTPrices201ResponseDataRelationshipsPriceTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPrices201ResponseDataRelationshipsPriceTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
