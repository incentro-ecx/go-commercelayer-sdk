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

// checks if the PATCHShippingWeightTiersShippingWeightTierId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHShippingWeightTiersShippingWeightTierId200Response{}

// PATCHShippingWeightTiersShippingWeightTierId200Response struct for PATCHShippingWeightTiersShippingWeightTierId200Response
type PATCHShippingWeightTiersShippingWeightTierId200Response struct {
	Data *PATCHShippingWeightTiersShippingWeightTierId200ResponseData `json:"data,omitempty"`
}

// NewPATCHShippingWeightTiersShippingWeightTierId200Response instantiates a new PATCHShippingWeightTiersShippingWeightTierId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShippingWeightTiersShippingWeightTierId200Response() *PATCHShippingWeightTiersShippingWeightTierId200Response {
	this := PATCHShippingWeightTiersShippingWeightTierId200Response{}
	return &this
}

// NewPATCHShippingWeightTiersShippingWeightTierId200ResponseWithDefaults instantiates a new PATCHShippingWeightTiersShippingWeightTierId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShippingWeightTiersShippingWeightTierId200ResponseWithDefaults() *PATCHShippingWeightTiersShippingWeightTierId200Response {
	this := PATCHShippingWeightTiersShippingWeightTierId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHShippingWeightTiersShippingWeightTierId200Response) GetData() PATCHShippingWeightTiersShippingWeightTierId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHShippingWeightTiersShippingWeightTierId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingWeightTiersShippingWeightTierId200Response) GetDataOk() (*PATCHShippingWeightTiersShippingWeightTierId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHShippingWeightTiersShippingWeightTierId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHShippingWeightTiersShippingWeightTierId200ResponseData and assigns it to the Data field.
func (o *PATCHShippingWeightTiersShippingWeightTierId200Response) SetData(v PATCHShippingWeightTiersShippingWeightTierId200ResponseData) {
	o.Data = &v
}

func (o PATCHShippingWeightTiersShippingWeightTierId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHShippingWeightTiersShippingWeightTierId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHShippingWeightTiersShippingWeightTierId200Response struct {
	value *PATCHShippingWeightTiersShippingWeightTierId200Response
	isSet bool
}

func (v NullablePATCHShippingWeightTiersShippingWeightTierId200Response) Get() *PATCHShippingWeightTiersShippingWeightTierId200Response {
	return v.value
}

func (v *NullablePATCHShippingWeightTiersShippingWeightTierId200Response) Set(val *PATCHShippingWeightTiersShippingWeightTierId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShippingWeightTiersShippingWeightTierId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShippingWeightTiersShippingWeightTierId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShippingWeightTiersShippingWeightTierId200Response(val *PATCHShippingWeightTiersShippingWeightTierId200Response) *NullablePATCHShippingWeightTiersShippingWeightTierId200Response {
	return &NullablePATCHShippingWeightTiersShippingWeightTierId200Response{value: val, isSet: true}
}

func (v NullablePATCHShippingWeightTiersShippingWeightTierId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShippingWeightTiersShippingWeightTierId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
