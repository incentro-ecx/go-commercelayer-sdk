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

// checks if the POSTMarkets201ResponseDataRelationshipsSubscriptionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationshipsSubscriptionModel{}

// POSTMarkets201ResponseDataRelationshipsSubscriptionModel struct for POSTMarkets201ResponseDataRelationshipsSubscriptionModel
type POSTMarkets201ResponseDataRelationshipsSubscriptionModel struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks       `json:"links,omitempty"`
	Data  *POSTMarkets201ResponseDataRelationshipsSubscriptionModelData `json:"data,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationshipsSubscriptionModel instantiates a new POSTMarkets201ResponseDataRelationshipsSubscriptionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationshipsSubscriptionModel() *POSTMarkets201ResponseDataRelationshipsSubscriptionModel {
	this := POSTMarkets201ResponseDataRelationshipsSubscriptionModel{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsSubscriptionModelWithDefaults instantiates a new POSTMarkets201ResponseDataRelationshipsSubscriptionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsSubscriptionModelWithDefaults() *POSTMarkets201ResponseDataRelationshipsSubscriptionModel {
	this := POSTMarkets201ResponseDataRelationshipsSubscriptionModel{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) GetData() POSTMarkets201ResponseDataRelationshipsSubscriptionModelData {
	if o == nil || IsNil(o.Data) {
		var ret POSTMarkets201ResponseDataRelationshipsSubscriptionModelData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) GetDataOk() (*POSTMarkets201ResponseDataRelationshipsSubscriptionModelData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTMarkets201ResponseDataRelationshipsSubscriptionModelData and assigns it to the Data field.
func (o *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) SetData(v POSTMarkets201ResponseDataRelationshipsSubscriptionModelData) {
	o.Data = &v
}

func (o POSTMarkets201ResponseDataRelationshipsSubscriptionModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationshipsSubscriptionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel struct {
	value *POSTMarkets201ResponseDataRelationshipsSubscriptionModel
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel) Get() *POSTMarkets201ResponseDataRelationshipsSubscriptionModel {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel) Set(val *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel(val *POSTMarkets201ResponseDataRelationshipsSubscriptionModel) *NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel {
	return &NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationshipsSubscriptionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
