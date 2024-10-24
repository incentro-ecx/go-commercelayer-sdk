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

// checks if the POSTMarkets201ResponseDataRelationshipsPriceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationshipsPriceList{}

// POSTMarkets201ResponseDataRelationshipsPriceList struct for POSTMarkets201ResponseDataRelationshipsPriceList
type POSTMarkets201ResponseDataRelationshipsPriceList struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTMarkets201ResponseDataRelationshipsPriceListData   `json:"data,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationshipsPriceList instantiates a new POSTMarkets201ResponseDataRelationshipsPriceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationshipsPriceList() *POSTMarkets201ResponseDataRelationshipsPriceList {
	this := POSTMarkets201ResponseDataRelationshipsPriceList{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsPriceListWithDefaults instantiates a new POSTMarkets201ResponseDataRelationshipsPriceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsPriceListWithDefaults() *POSTMarkets201ResponseDataRelationshipsPriceList {
	this := POSTMarkets201ResponseDataRelationshipsPriceList{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationshipsPriceList) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationshipsPriceList) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsPriceList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTMarkets201ResponseDataRelationshipsPriceList) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationshipsPriceList) GetData() POSTMarkets201ResponseDataRelationshipsPriceListData {
	if o == nil || IsNil(o.Data) {
		var ret POSTMarkets201ResponseDataRelationshipsPriceListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationshipsPriceList) GetDataOk() (*POSTMarkets201ResponseDataRelationshipsPriceListData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsPriceList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTMarkets201ResponseDataRelationshipsPriceListData and assigns it to the Data field.
func (o *POSTMarkets201ResponseDataRelationshipsPriceList) SetData(v POSTMarkets201ResponseDataRelationshipsPriceListData) {
	o.Data = &v
}

func (o POSTMarkets201ResponseDataRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationshipsPriceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationshipsPriceList struct {
	value *POSTMarkets201ResponseDataRelationshipsPriceList
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsPriceList) Get() *POSTMarkets201ResponseDataRelationshipsPriceList {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsPriceList) Set(val *POSTMarkets201ResponseDataRelationshipsPriceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsPriceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsPriceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationshipsPriceList(val *POSTMarkets201ResponseDataRelationshipsPriceList) *NullablePOSTMarkets201ResponseDataRelationshipsPriceList {
	return &NullablePOSTMarkets201ResponseDataRelationshipsPriceList{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsPriceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
