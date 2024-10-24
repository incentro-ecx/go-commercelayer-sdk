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

// checks if the GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource{}

// GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource struct for GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource
type GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                   `json:"links,omitempty"`
	Data  *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResourceData `json:"data,omitempty"`
}

// NewGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource instantiates a new GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource() *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource {
	this := GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource{}
	return &this
}

// NewGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResourceWithDefaults instantiates a new GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResourceWithDefaults() *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource {
	this := GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) GetData() GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResourceData {
	if o == nil || IsNil(o.Data) {
		var ret GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResourceData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) GetDataOk() (*GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResourceData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResourceData and assigns it to the Data field.
func (o *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) SetData(v GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResourceData) {
	o.Data = &v
}

func (o GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource struct {
	value *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource
	isSet bool
}

func (v NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) Get() *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource {
	return v.value
}

func (v *NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) Set(val *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource(val *GETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) *NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource {
	return &NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource{value: val, isSet: true}
}

func (v NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETResourceErrorsResourceErrorId200ResponseDataRelationshipsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}