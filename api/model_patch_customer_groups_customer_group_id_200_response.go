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

// checks if the PATCHCustomerGroupsCustomerGroupId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCustomerGroupsCustomerGroupId200Response{}

// PATCHCustomerGroupsCustomerGroupId200Response struct for PATCHCustomerGroupsCustomerGroupId200Response
type PATCHCustomerGroupsCustomerGroupId200Response struct {
	Data *PATCHCustomerGroupsCustomerGroupId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCustomerGroupsCustomerGroupId200Response instantiates a new PATCHCustomerGroupsCustomerGroupId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomerGroupsCustomerGroupId200Response() *PATCHCustomerGroupsCustomerGroupId200Response {
	this := PATCHCustomerGroupsCustomerGroupId200Response{}
	return &this
}

// NewPATCHCustomerGroupsCustomerGroupId200ResponseWithDefaults instantiates a new PATCHCustomerGroupsCustomerGroupId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomerGroupsCustomerGroupId200ResponseWithDefaults() *PATCHCustomerGroupsCustomerGroupId200Response {
	this := PATCHCustomerGroupsCustomerGroupId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCustomerGroupsCustomerGroupId200Response) GetData() PATCHCustomerGroupsCustomerGroupId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHCustomerGroupsCustomerGroupId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerGroupsCustomerGroupId200Response) GetDataOk() (*PATCHCustomerGroupsCustomerGroupId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCustomerGroupsCustomerGroupId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCustomerGroupsCustomerGroupId200ResponseData and assigns it to the Data field.
func (o *PATCHCustomerGroupsCustomerGroupId200Response) SetData(v PATCHCustomerGroupsCustomerGroupId200ResponseData) {
	o.Data = &v
}

func (o PATCHCustomerGroupsCustomerGroupId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCustomerGroupsCustomerGroupId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHCustomerGroupsCustomerGroupId200Response struct {
	value *PATCHCustomerGroupsCustomerGroupId200Response
	isSet bool
}

func (v NullablePATCHCustomerGroupsCustomerGroupId200Response) Get() *PATCHCustomerGroupsCustomerGroupId200Response {
	return v.value
}

func (v *NullablePATCHCustomerGroupsCustomerGroupId200Response) Set(val *PATCHCustomerGroupsCustomerGroupId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomerGroupsCustomerGroupId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomerGroupsCustomerGroupId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomerGroupsCustomerGroupId200Response(val *PATCHCustomerGroupsCustomerGroupId200Response) *NullablePATCHCustomerGroupsCustomerGroupId200Response {
	return &NullablePATCHCustomerGroupsCustomerGroupId200Response{value: val, isSet: true}
}

func (v NullablePATCHCustomerGroupsCustomerGroupId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomerGroupsCustomerGroupId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
