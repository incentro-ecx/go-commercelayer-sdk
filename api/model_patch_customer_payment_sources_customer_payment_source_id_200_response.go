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

// checks if the PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response{}

// PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response struct for PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response
type PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response struct {
	Data *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response instantiates a new PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response() *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response {
	this := PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response{}
	return &this
}

// NewPATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseWithDefaults instantiates a new PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseWithDefaults() *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response {
	this := PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) GetData() PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) GetDataOk() (*PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseData and assigns it to the Data field.
func (o *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) SetData(v PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseData) {
	o.Data = &v
}

func (o PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response struct {
	value *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response
	isSet bool
}

func (v NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) Get() *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response {
	return v.value
}

func (v *NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) Set(val *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response(val *PATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) *NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response {
	return &NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response{value: val, isSet: true}
}

func (v NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
