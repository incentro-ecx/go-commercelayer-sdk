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

// checks if the AdyenPaymentDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenPaymentDataRelationships{}

// AdyenPaymentDataRelationships struct for AdyenPaymentDataRelationships
type AdyenPaymentDataRelationships struct {
	Order          *AdyenPaymentDataRelationshipsOrder          `json:"order,omitempty"`
	PaymentGateway *AdyenPaymentDataRelationshipsPaymentGateway `json:"payment_gateway,omitempty"`
	Versions       *AddressDataRelationshipsVersions            `json:"versions,omitempty"`
}

// NewAdyenPaymentDataRelationships instantiates a new AdyenPaymentDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentDataRelationships() *AdyenPaymentDataRelationships {
	this := AdyenPaymentDataRelationships{}
	return &this
}

// NewAdyenPaymentDataRelationshipsWithDefaults instantiates a new AdyenPaymentDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentDataRelationshipsWithDefaults() *AdyenPaymentDataRelationships {
	this := AdyenPaymentDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AdyenPaymentDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *AdyenPaymentDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *AdyenPaymentDataRelationships) GetPaymentGateway() AdyenPaymentDataRelationshipsPaymentGateway {
	if o == nil || IsNil(o.PaymentGateway) {
		var ret AdyenPaymentDataRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentDataRelationshipsPaymentGateway, bool) {
	if o == nil || IsNil(o.PaymentGateway) {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationships) HasPaymentGateway() bool {
	if o != nil && !IsNil(o.PaymentGateway) {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given AdyenPaymentDataRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *AdyenPaymentDataRelationships) SetPaymentGateway(v AdyenPaymentDataRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AdyenPaymentDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AdyenPaymentDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *AdyenPaymentDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o AdyenPaymentDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenPaymentDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.PaymentGateway) {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableAdyenPaymentDataRelationships struct {
	value *AdyenPaymentDataRelationships
	isSet bool
}

func (v NullableAdyenPaymentDataRelationships) Get() *AdyenPaymentDataRelationships {
	return v.value
}

func (v *NullableAdyenPaymentDataRelationships) Set(val *AdyenPaymentDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentDataRelationships(val *AdyenPaymentDataRelationships) *NullableAdyenPaymentDataRelationships {
	return &NullableAdyenPaymentDataRelationships{value: val, isSet: true}
}

func (v NullableAdyenPaymentDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
