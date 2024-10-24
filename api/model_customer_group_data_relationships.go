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

// checks if the CustomerGroupDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerGroupDataRelationships{}

// CustomerGroupDataRelationships struct for CustomerGroupDataRelationships
type CustomerGroupDataRelationships struct {
	Customers   *CouponRecipientDataRelationshipsCustomer  `json:"customers,omitempty"`
	Markets     *AvalaraAccountDataRelationshipsMarkets    `json:"markets,omitempty"`
	Attachments *AuthorizationDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions    *AddressDataRelationshipsVersions          `json:"versions,omitempty"`
}

// NewCustomerGroupDataRelationships instantiates a new CustomerGroupDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGroupDataRelationships() *CustomerGroupDataRelationships {
	this := CustomerGroupDataRelationships{}
	return &this
}

// NewCustomerGroupDataRelationshipsWithDefaults instantiates a new CustomerGroupDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGroupDataRelationshipsWithDefaults() *CustomerGroupDataRelationships {
	this := CustomerGroupDataRelationships{}
	return &this
}

// GetCustomers returns the Customers field value if set, zero value otherwise.
func (o *CustomerGroupDataRelationships) GetCustomers() CouponRecipientDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customers) {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customers
}

// GetCustomersOk returns a tuple with the Customers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupDataRelationships) GetCustomersOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customers) {
		return nil, false
	}
	return o.Customers, true
}

// HasCustomers returns a boolean if a field has been set.
func (o *CustomerGroupDataRelationships) HasCustomers() bool {
	if o != nil && !IsNil(o.Customers) {
		return true
	}

	return false
}

// SetCustomers gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customers field.
func (o *CustomerGroupDataRelationships) SetCustomers(v CouponRecipientDataRelationshipsCustomer) {
	o.Customers = &v
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *CustomerGroupDataRelationships) GetMarkets() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || IsNil(o.Markets) {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupDataRelationships) GetMarketsOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *CustomerGroupDataRelationships) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Markets field.
func (o *CustomerGroupDataRelationships) SetMarkets(v AvalaraAccountDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CustomerGroupDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CustomerGroupDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *CustomerGroupDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *CustomerGroupDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *CustomerGroupDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *CustomerGroupDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o CustomerGroupDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerGroupDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customers) {
		toSerialize["customers"] = o.Customers
	}
	if !IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableCustomerGroupDataRelationships struct {
	value *CustomerGroupDataRelationships
	isSet bool
}

func (v NullableCustomerGroupDataRelationships) Get() *CustomerGroupDataRelationships {
	return v.value
}

func (v *NullableCustomerGroupDataRelationships) Set(val *CustomerGroupDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGroupDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGroupDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGroupDataRelationships(val *CustomerGroupDataRelationships) *NullableCustomerGroupDataRelationships {
	return &NullableCustomerGroupDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerGroupDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGroupDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
