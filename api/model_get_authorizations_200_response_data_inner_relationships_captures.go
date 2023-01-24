/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETAuthorizations200ResponseDataInnerRelationshipsCaptures struct for GETAuthorizations200ResponseDataInnerRelationshipsCaptures
type GETAuthorizations200ResponseDataInnerRelationshipsCaptures struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *GETAuthorizations200ResponseDataInnerRelationshipsCapturesData `json:"data,omitempty"`
}

// NewGETAuthorizations200ResponseDataInnerRelationshipsCaptures instantiates a new GETAuthorizations200ResponseDataInnerRelationshipsCaptures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAuthorizations200ResponseDataInnerRelationshipsCaptures() *GETAuthorizations200ResponseDataInnerRelationshipsCaptures {
	this := GETAuthorizations200ResponseDataInnerRelationshipsCaptures{}
	return &this
}

// NewGETAuthorizations200ResponseDataInnerRelationshipsCapturesWithDefaults instantiates a new GETAuthorizations200ResponseDataInnerRelationshipsCaptures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAuthorizations200ResponseDataInnerRelationshipsCapturesWithDefaults() *GETAuthorizations200ResponseDataInnerRelationshipsCaptures {
	this := GETAuthorizations200ResponseDataInnerRelationshipsCaptures{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) GetData() GETAuthorizations200ResponseDataInnerRelationshipsCapturesData {
	if o == nil || o.Data == nil {
		var ret GETAuthorizations200ResponseDataInnerRelationshipsCapturesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) GetDataOk() (*GETAuthorizations200ResponseDataInnerRelationshipsCapturesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAuthorizations200ResponseDataInnerRelationshipsCapturesData and assigns it to the Data field.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) SetData(v GETAuthorizations200ResponseDataInnerRelationshipsCapturesData) {
	o.Data = &v
}

func (o GETAuthorizations200ResponseDataInnerRelationshipsCaptures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures struct {
	value *GETAuthorizations200ResponseDataInnerRelationshipsCaptures
	isSet bool
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures) Get() *GETAuthorizations200ResponseDataInnerRelationshipsCaptures {
	return v.value
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures) Set(val *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures(val *GETAuthorizations200ResponseDataInnerRelationshipsCaptures) *NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures {
	return &NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures{value: val, isSet: true}
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsCaptures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
