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

// checks if the POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments{}

// POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments struct for POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments
type POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks           `json:"links,omitempty"`
	Data  *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData `json:"data,omitempty"`
}

// NewPOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments instantiates a new POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments() *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments {
	this := POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments{}
	return &this
}

// NewPOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsWithDefaults instantiates a new POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsWithDefaults() *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments {
	this := POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) GetData() POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) GetDataOk() (*POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData and assigns it to the Data field.
func (o *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) SetData(v POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPaymentsData) {
	o.Data = &v
}

func (o POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments struct {
	value *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments
	isSet bool
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) Get() *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments {
	return v.value
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) Set(val *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments(val *POSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) *NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments {
	return &NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments{value: val, isSet: true}
}

func (v NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTKlarnaGateways201ResponseDataRelationshipsKlarnaPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
