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

// checks if the PATCHStockLocationsStockLocationId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHStockLocationsStockLocationId200Response{}

// PATCHStockLocationsStockLocationId200Response struct for PATCHStockLocationsStockLocationId200Response
type PATCHStockLocationsStockLocationId200Response struct {
	Data *PATCHStockLocationsStockLocationId200ResponseData `json:"data,omitempty"`
}

// NewPATCHStockLocationsStockLocationId200Response instantiates a new PATCHStockLocationsStockLocationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHStockLocationsStockLocationId200Response() *PATCHStockLocationsStockLocationId200Response {
	this := PATCHStockLocationsStockLocationId200Response{}
	return &this
}

// NewPATCHStockLocationsStockLocationId200ResponseWithDefaults instantiates a new PATCHStockLocationsStockLocationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHStockLocationsStockLocationId200ResponseWithDefaults() *PATCHStockLocationsStockLocationId200Response {
	this := PATCHStockLocationsStockLocationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHStockLocationsStockLocationId200Response) GetData() PATCHStockLocationsStockLocationId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHStockLocationsStockLocationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStockLocationsStockLocationId200Response) GetDataOk() (*PATCHStockLocationsStockLocationId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHStockLocationsStockLocationId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHStockLocationsStockLocationId200ResponseData and assigns it to the Data field.
func (o *PATCHStockLocationsStockLocationId200Response) SetData(v PATCHStockLocationsStockLocationId200ResponseData) {
	o.Data = &v
}

func (o PATCHStockLocationsStockLocationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHStockLocationsStockLocationId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHStockLocationsStockLocationId200Response struct {
	value *PATCHStockLocationsStockLocationId200Response
	isSet bool
}

func (v NullablePATCHStockLocationsStockLocationId200Response) Get() *PATCHStockLocationsStockLocationId200Response {
	return v.value
}

func (v *NullablePATCHStockLocationsStockLocationId200Response) Set(val *PATCHStockLocationsStockLocationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHStockLocationsStockLocationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHStockLocationsStockLocationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHStockLocationsStockLocationId200Response(val *PATCHStockLocationsStockLocationId200Response) *NullablePATCHStockLocationsStockLocationId200Response {
	return &NullablePATCHStockLocationsStockLocationId200Response{value: val, isSet: true}
}

func (v NullablePATCHStockLocationsStockLocationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHStockLocationsStockLocationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
