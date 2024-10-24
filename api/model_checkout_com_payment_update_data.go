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

// checks if the CheckoutComPaymentUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutComPaymentUpdateData{}

// CheckoutComPaymentUpdateData struct for CheckoutComPaymentUpdateData
type CheckoutComPaymentUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                           `json:"id"`
	Attributes    PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships                                  `json:"relationships,omitempty"`
}

// NewCheckoutComPaymentUpdateData instantiates a new CheckoutComPaymentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComPaymentUpdateData(type_ interface{}, id interface{}, attributes PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) *CheckoutComPaymentUpdateData {
	this := CheckoutComPaymentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCheckoutComPaymentUpdateDataWithDefaults instantiates a new CheckoutComPaymentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComPaymentUpdateDataWithDefaults() *CheckoutComPaymentUpdateData {
	this := CheckoutComPaymentUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CheckoutComPaymentUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutComPaymentUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutComPaymentUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CheckoutComPaymentUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutComPaymentUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutComPaymentUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CheckoutComPaymentUpdateData) GetAttributes() PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateData) GetAttributesOk() (*PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CheckoutComPaymentUpdateData) SetAttributes(v PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *CheckoutComPaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o CheckoutComPaymentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutComPaymentUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableCheckoutComPaymentUpdateData struct {
	value *CheckoutComPaymentUpdateData
	isSet bool
}

func (v NullableCheckoutComPaymentUpdateData) Get() *CheckoutComPaymentUpdateData {
	return v.value
}

func (v *NullableCheckoutComPaymentUpdateData) Set(val *CheckoutComPaymentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComPaymentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComPaymentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComPaymentUpdateData(val *CheckoutComPaymentUpdateData) *NullableCheckoutComPaymentUpdateData {
	return &NullableCheckoutComPaymentUpdateData{value: val, isSet: true}
}

func (v NullableCheckoutComPaymentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComPaymentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
