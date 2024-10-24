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

// checks if the POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories{}

// POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories struct for POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories
type POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks           `json:"links,omitempty"`
	Data  *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategoriesData `json:"data,omitempty"`
}

// NewPOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories instantiates a new POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories() *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories {
	this := POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories{}
	return &this
}

// NewPOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategoriesWithDefaults instantiates a new POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategoriesWithDefaults() *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories {
	this := POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) GetData() POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategoriesData {
	if o == nil || IsNil(o.Data) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategoriesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) GetDataOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategoriesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategoriesData and assigns it to the Data field.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) SetData(v POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategoriesData) {
	o.Data = &v
}

func (o POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories struct {
	value *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories
	isSet bool
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) Get() *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories {
	return v.value
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) Set(val *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories(val *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories {
	return &NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories{value: val, isSet: true}
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
