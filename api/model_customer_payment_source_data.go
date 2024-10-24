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

// checks if the CustomerPaymentSourceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPaymentSourceData{}

// CustomerPaymentSourceData struct for CustomerPaymentSourceData
type CustomerPaymentSourceData struct {
	// The resource's type
	Type          interface{}                                                               `json:"type"`
	Attributes    GETCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes `json:"attributes"`
	Relationships *CustomerPaymentSourceDataRelationships                                   `json:"relationships,omitempty"`
}

// NewCustomerPaymentSourceData instantiates a new CustomerPaymentSourceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceData(type_ interface{}, attributes GETCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes) *CustomerPaymentSourceData {
	this := CustomerPaymentSourceData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerPaymentSourceDataWithDefaults instantiates a new CustomerPaymentSourceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceDataWithDefaults() *CustomerPaymentSourceData {
	this := CustomerPaymentSourceData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerPaymentSourceData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerPaymentSourceData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerPaymentSourceData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerPaymentSourceData) GetAttributes() GETCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes {
	if o == nil {
		var ret GETCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceData) GetAttributesOk() (*GETCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerPaymentSourceData) SetAttributes(v GETCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerPaymentSourceData) GetRelationships() CustomerPaymentSourceDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CustomerPaymentSourceDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceData) GetRelationshipsOk() (*CustomerPaymentSourceDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerPaymentSourceData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerPaymentSourceDataRelationships and assigns it to the Relationships field.
func (o *CustomerPaymentSourceData) SetRelationships(v CustomerPaymentSourceDataRelationships) {
	o.Relationships = &v
}

func (o CustomerPaymentSourceData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPaymentSourceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableCustomerPaymentSourceData struct {
	value *CustomerPaymentSourceData
	isSet bool
}

func (v NullableCustomerPaymentSourceData) Get() *CustomerPaymentSourceData {
	return v.value
}

func (v *NullableCustomerPaymentSourceData) Set(val *CustomerPaymentSourceData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceData(val *CustomerPaymentSourceData) *NullableCustomerPaymentSourceData {
	return &NullableCustomerPaymentSourceData{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
