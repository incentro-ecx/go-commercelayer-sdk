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

// checks if the POSTCustomerSubscriptions201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerSubscriptions201ResponseDataRelationships{}

// POSTCustomerSubscriptions201ResponseDataRelationships struct for POSTCustomerSubscriptions201ResponseDataRelationships
type POSTCustomerSubscriptions201ResponseDataRelationships struct {
	Customer *POSTCouponRecipients201ResponseDataRelationshipsCustomer `json:"customer,omitempty"`
	Events   *POSTAddresses201ResponseDataRelationshipsEvents          `json:"events,omitempty"`
	Versions *POSTAddresses201ResponseDataRelationshipsVersions        `json:"versions,omitempty"`
}

// NewPOSTCustomerSubscriptions201ResponseDataRelationships instantiates a new POSTCustomerSubscriptions201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerSubscriptions201ResponseDataRelationships() *POSTCustomerSubscriptions201ResponseDataRelationships {
	this := POSTCustomerSubscriptions201ResponseDataRelationships{}
	return &this
}

// NewPOSTCustomerSubscriptions201ResponseDataRelationshipsWithDefaults instantiates a new POSTCustomerSubscriptions201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerSubscriptions201ResponseDataRelationshipsWithDefaults() *POSTCustomerSubscriptions201ResponseDataRelationships {
	this := POSTCustomerSubscriptions201ResponseDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret POSTCouponRecipients201ResponseDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given POSTCouponRecipients201ResponseDataRelationshipsCustomer and assigns it to the Customer field.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTCustomerSubscriptions201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTCustomerSubscriptions201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerSubscriptions201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTCustomerSubscriptions201ResponseDataRelationships struct {
	value *POSTCustomerSubscriptions201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTCustomerSubscriptions201ResponseDataRelationships) Get() *POSTCustomerSubscriptions201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTCustomerSubscriptions201ResponseDataRelationships) Set(val *POSTCustomerSubscriptions201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerSubscriptions201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerSubscriptions201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerSubscriptions201ResponseDataRelationships(val *POSTCustomerSubscriptions201ResponseDataRelationships) *NullablePOSTCustomerSubscriptions201ResponseDataRelationships {
	return &NullablePOSTCustomerSubscriptions201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTCustomerSubscriptions201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerSubscriptions201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
