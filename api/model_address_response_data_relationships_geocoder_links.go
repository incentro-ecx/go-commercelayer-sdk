/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AddressResponseDataRelationshipsGeocoderLinks struct for AddressResponseDataRelationshipsGeocoderLinks
type AddressResponseDataRelationshipsGeocoderLinks struct {
	// URL
	Self *string `json:"self,omitempty"`
	// URL
	Related *string `json:"related,omitempty"`
}

// NewAddressResponseDataRelationshipsGeocoderLinks instantiates a new AddressResponseDataRelationshipsGeocoderLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressResponseDataRelationshipsGeocoderLinks() *AddressResponseDataRelationshipsGeocoderLinks {
	this := AddressResponseDataRelationshipsGeocoderLinks{}
	return &this
}

// NewAddressResponseDataRelationshipsGeocoderLinksWithDefaults instantiates a new AddressResponseDataRelationshipsGeocoderLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressResponseDataRelationshipsGeocoderLinksWithDefaults() *AddressResponseDataRelationshipsGeocoderLinks {
	this := AddressResponseDataRelationshipsGeocoderLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AddressResponseDataRelationshipsGeocoderLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseDataRelationshipsGeocoderLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AddressResponseDataRelationshipsGeocoderLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *AddressResponseDataRelationshipsGeocoderLinks) SetSelf(v string) {
	o.Self = &v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *AddressResponseDataRelationshipsGeocoderLinks) GetRelated() string {
	if o == nil || o.Related == nil {
		var ret string
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponseDataRelationshipsGeocoderLinks) GetRelatedOk() (*string, bool) {
	if o == nil || o.Related == nil {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *AddressResponseDataRelationshipsGeocoderLinks) HasRelated() bool {
	if o != nil && o.Related != nil {
		return true
	}

	return false
}

// SetRelated gets a reference to the given string and assigns it to the Related field.
func (o *AddressResponseDataRelationshipsGeocoderLinks) SetRelated(v string) {
	o.Related = &v
}

func (o AddressResponseDataRelationshipsGeocoderLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Related != nil {
		toSerialize["related"] = o.Related
	}
	return json.Marshal(toSerialize)
}

type NullableAddressResponseDataRelationshipsGeocoderLinks struct {
	value *AddressResponseDataRelationshipsGeocoderLinks
	isSet bool
}

func (v NullableAddressResponseDataRelationshipsGeocoderLinks) Get() *AddressResponseDataRelationshipsGeocoderLinks {
	return v.value
}

func (v *NullableAddressResponseDataRelationshipsGeocoderLinks) Set(val *AddressResponseDataRelationshipsGeocoderLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressResponseDataRelationshipsGeocoderLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressResponseDataRelationshipsGeocoderLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressResponseDataRelationshipsGeocoderLinks(val *AddressResponseDataRelationshipsGeocoderLinks) *NullableAddressResponseDataRelationshipsGeocoderLinks {
	return &NullableAddressResponseDataRelationshipsGeocoderLinks{value: val, isSet: true}
}

func (v NullableAddressResponseDataRelationshipsGeocoderLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressResponseDataRelationshipsGeocoderLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}