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

// checks if the POSTOrders201ResponseDataRelationshipsTransactions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsTransactions{}

// POSTOrders201ResponseDataRelationshipsTransactions struct for POSTOrders201ResponseDataRelationshipsTransactions
type POSTOrders201ResponseDataRelationshipsTransactions struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsTransactionsData `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsTransactions instantiates a new POSTOrders201ResponseDataRelationshipsTransactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsTransactions() *POSTOrders201ResponseDataRelationshipsTransactions {
	this := POSTOrders201ResponseDataRelationshipsTransactions{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsTransactionsWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsTransactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsTransactionsWithDefaults() *POSTOrders201ResponseDataRelationshipsTransactions {
	this := POSTOrders201ResponseDataRelationshipsTransactions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsTransactions) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsTransactions) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsTransactions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsTransactions) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsTransactions) GetData() POSTOrders201ResponseDataRelationshipsTransactionsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsTransactionsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsTransactions) GetDataOk() (*POSTOrders201ResponseDataRelationshipsTransactionsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsTransactions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsTransactionsData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsTransactions) SetData(v POSTOrders201ResponseDataRelationshipsTransactionsData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsTransactions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsTransactions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsTransactions struct {
	value *POSTOrders201ResponseDataRelationshipsTransactions
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsTransactions) Get() *POSTOrders201ResponseDataRelationshipsTransactions {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsTransactions) Set(val *POSTOrders201ResponseDataRelationshipsTransactions) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsTransactions(val *POSTOrders201ResponseDataRelationshipsTransactions) *NullablePOSTOrders201ResponseDataRelationshipsTransactions {
	return &NullablePOSTOrders201ResponseDataRelationshipsTransactions{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}