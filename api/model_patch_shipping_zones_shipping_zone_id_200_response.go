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

// PATCHShippingZonesShippingZoneId200Response struct for PATCHShippingZonesShippingZoneId200Response
type PATCHShippingZonesShippingZoneId200Response struct {
	Data *PATCHShippingZonesShippingZoneId200ResponseData `json:"data,omitempty"`
}

// NewPATCHShippingZonesShippingZoneId200Response instantiates a new PATCHShippingZonesShippingZoneId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShippingZonesShippingZoneId200Response() *PATCHShippingZonesShippingZoneId200Response {
	this := PATCHShippingZonesShippingZoneId200Response{}
	return &this
}

// NewPATCHShippingZonesShippingZoneId200ResponseWithDefaults instantiates a new PATCHShippingZonesShippingZoneId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShippingZonesShippingZoneId200ResponseWithDefaults() *PATCHShippingZonesShippingZoneId200Response {
	this := PATCHShippingZonesShippingZoneId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200Response) GetData() PATCHShippingZonesShippingZoneId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHShippingZonesShippingZoneId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200Response) GetDataOk() (*PATCHShippingZonesShippingZoneId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHShippingZonesShippingZoneId200ResponseData and assigns it to the Data field.
func (o *PATCHShippingZonesShippingZoneId200Response) SetData(v PATCHShippingZonesShippingZoneId200ResponseData) {
	o.Data = &v
}

func (o PATCHShippingZonesShippingZoneId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHShippingZonesShippingZoneId200Response struct {
	value *PATCHShippingZonesShippingZoneId200Response
	isSet bool
}

func (v NullablePATCHShippingZonesShippingZoneId200Response) Get() *PATCHShippingZonesShippingZoneId200Response {
	return v.value
}

func (v *NullablePATCHShippingZonesShippingZoneId200Response) Set(val *PATCHShippingZonesShippingZoneId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShippingZonesShippingZoneId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShippingZonesShippingZoneId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShippingZonesShippingZoneId200Response(val *PATCHShippingZonesShippingZoneId200Response) *NullablePATCHShippingZonesShippingZoneId200Response {
	return &NullablePATCHShippingZonesShippingZoneId200Response{value: val, isSet: true}
}

func (v NullablePATCHShippingZonesShippingZoneId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShippingZonesShippingZoneId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
