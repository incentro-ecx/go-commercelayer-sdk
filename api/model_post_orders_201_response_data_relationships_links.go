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

// checks if the POSTOrders201ResponseDataRelationshipsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsLinks{}

// POSTOrders201ResponseDataRelationshipsLinks struct for POSTOrders201ResponseDataRelationshipsLinks
type POSTOrders201ResponseDataRelationshipsLinks struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsLinksData        `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsLinks instantiates a new POSTOrders201ResponseDataRelationshipsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsLinks() *POSTOrders201ResponseDataRelationshipsLinks {
	this := POSTOrders201ResponseDataRelationshipsLinks{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsLinksWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsLinksWithDefaults() *POSTOrders201ResponseDataRelationshipsLinks {
	this := POSTOrders201ResponseDataRelationshipsLinks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsLinks) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsLinks) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsLinks) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsLinks) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsLinks) GetData() POSTOrders201ResponseDataRelationshipsLinksData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsLinksData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsLinks) GetDataOk() (*POSTOrders201ResponseDataRelationshipsLinksData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsLinks) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsLinksData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsLinks) SetData(v POSTOrders201ResponseDataRelationshipsLinksData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsLinks struct {
	value *POSTOrders201ResponseDataRelationshipsLinks
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsLinks) Get() *POSTOrders201ResponseDataRelationshipsLinks {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsLinks) Set(val *POSTOrders201ResponseDataRelationshipsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsLinks(val *POSTOrders201ResponseDataRelationshipsLinks) *NullablePOSTOrders201ResponseDataRelationshipsLinks {
	return &NullablePOSTOrders201ResponseDataRelationshipsLinks{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}