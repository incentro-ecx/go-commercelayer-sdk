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

// checks if the POSTTaxRules201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTTaxRules201ResponseDataAttributes{}

// POSTTaxRules201ResponseDataAttributes struct for POSTTaxRules201ResponseDataAttributes
type POSTTaxRules201ResponseDataAttributes struct {
	// The tax rule internal name.
	Name interface{} `json:"name"`
	// The tax rate for this rule.
	TaxRate interface{} `json:"tax_rate,omitempty"`
	// The regex that will be evaluated to match the shipping address country code.
	CountryCodeRegex interface{} `json:"country_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address country code.
	NotCountryCodeRegex interface{} `json:"not_country_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address state code.
	StateCodeRegex interface{} `json:"state_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address state code.
	NotStateCodeRegex interface{} `json:"not_state_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address zip code.
	ZipCodeRegex interface{} `json:"zip_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping zip country code.
	NotZipCodeRegex interface{} `json:"not_zip_code_regex,omitempty"`
	// Indicates if the freight is taxable.
	FreightTaxable interface{} `json:"freight_taxable,omitempty"`
	// Indicates if the payment method is taxable.
	PaymentMethodTaxable interface{} `json:"payment_method_taxable,omitempty"`
	// Indicates if gift cards are taxable.
	GiftCardTaxable interface{} `json:"gift_card_taxable,omitempty"`
	// Indicates if adjustemnts are taxable.
	AdjustmentTaxable interface{} `json:"adjustment_taxable,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTTaxRules201ResponseDataAttributes instantiates a new POSTTaxRules201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTTaxRules201ResponseDataAttributes(name interface{}) *POSTTaxRules201ResponseDataAttributes {
	this := POSTTaxRules201ResponseDataAttributes{}
	this.Name = name
	return &this
}

// NewPOSTTaxRules201ResponseDataAttributesWithDefaults instantiates a new POSTTaxRules201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTTaxRules201ResponseDataAttributesWithDefaults() *POSTTaxRules201ResponseDataAttributes {
	this := POSTTaxRules201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTTaxRules201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetTaxRate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetTaxRateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TaxRate) {
		return nil, false
	}
	return &o.TaxRate, true
}

// HasTaxRate returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasTaxRate() bool {
	if o != nil && IsNil(o.TaxRate) {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given interface{} and assigns it to the TaxRate field.
func (o *POSTTaxRules201ResponseDataAttributes) SetTaxRate(v interface{}) {
	o.TaxRate = v
}

// GetCountryCodeRegex returns the CountryCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetCountryCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CountryCodeRegex
}

// GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetCountryCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CountryCodeRegex) {
		return nil, false
	}
	return &o.CountryCodeRegex, true
}

// HasCountryCodeRegex returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasCountryCodeRegex() bool {
	if o != nil && IsNil(o.CountryCodeRegex) {
		return true
	}

	return false
}

// SetCountryCodeRegex gets a reference to the given interface{} and assigns it to the CountryCodeRegex field.
func (o *POSTTaxRules201ResponseDataAttributes) SetCountryCodeRegex(v interface{}) {
	o.CountryCodeRegex = v
}

// GetNotCountryCodeRegex returns the NotCountryCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetNotCountryCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NotCountryCodeRegex
}

// GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetNotCountryCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NotCountryCodeRegex) {
		return nil, false
	}
	return &o.NotCountryCodeRegex, true
}

// HasNotCountryCodeRegex returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasNotCountryCodeRegex() bool {
	if o != nil && IsNil(o.NotCountryCodeRegex) {
		return true
	}

	return false
}

// SetNotCountryCodeRegex gets a reference to the given interface{} and assigns it to the NotCountryCodeRegex field.
func (o *POSTTaxRules201ResponseDataAttributes) SetNotCountryCodeRegex(v interface{}) {
	o.NotCountryCodeRegex = v
}

// GetStateCodeRegex returns the StateCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetStateCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StateCodeRegex
}

// GetStateCodeRegexOk returns a tuple with the StateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetStateCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StateCodeRegex) {
		return nil, false
	}
	return &o.StateCodeRegex, true
}

// HasStateCodeRegex returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasStateCodeRegex() bool {
	if o != nil && IsNil(o.StateCodeRegex) {
		return true
	}

	return false
}

// SetStateCodeRegex gets a reference to the given interface{} and assigns it to the StateCodeRegex field.
func (o *POSTTaxRules201ResponseDataAttributes) SetStateCodeRegex(v interface{}) {
	o.StateCodeRegex = v
}

// GetNotStateCodeRegex returns the NotStateCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetNotStateCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NotStateCodeRegex
}

// GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetNotStateCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NotStateCodeRegex) {
		return nil, false
	}
	return &o.NotStateCodeRegex, true
}

// HasNotStateCodeRegex returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasNotStateCodeRegex() bool {
	if o != nil && IsNil(o.NotStateCodeRegex) {
		return true
	}

	return false
}

// SetNotStateCodeRegex gets a reference to the given interface{} and assigns it to the NotStateCodeRegex field.
func (o *POSTTaxRules201ResponseDataAttributes) SetNotStateCodeRegex(v interface{}) {
	o.NotStateCodeRegex = v
}

// GetZipCodeRegex returns the ZipCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetZipCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ZipCodeRegex
}

// GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetZipCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ZipCodeRegex) {
		return nil, false
	}
	return &o.ZipCodeRegex, true
}

// HasZipCodeRegex returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasZipCodeRegex() bool {
	if o != nil && IsNil(o.ZipCodeRegex) {
		return true
	}

	return false
}

// SetZipCodeRegex gets a reference to the given interface{} and assigns it to the ZipCodeRegex field.
func (o *POSTTaxRules201ResponseDataAttributes) SetZipCodeRegex(v interface{}) {
	o.ZipCodeRegex = v
}

// GetNotZipCodeRegex returns the NotZipCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetNotZipCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NotZipCodeRegex
}

// GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetNotZipCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NotZipCodeRegex) {
		return nil, false
	}
	return &o.NotZipCodeRegex, true
}

// HasNotZipCodeRegex returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasNotZipCodeRegex() bool {
	if o != nil && IsNil(o.NotZipCodeRegex) {
		return true
	}

	return false
}

// SetNotZipCodeRegex gets a reference to the given interface{} and assigns it to the NotZipCodeRegex field.
func (o *POSTTaxRules201ResponseDataAttributes) SetNotZipCodeRegex(v interface{}) {
	o.NotZipCodeRegex = v
}

// GetFreightTaxable returns the FreightTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetFreightTaxable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FreightTaxable
}

// GetFreightTaxableOk returns a tuple with the FreightTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetFreightTaxableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FreightTaxable) {
		return nil, false
	}
	return &o.FreightTaxable, true
}

// HasFreightTaxable returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasFreightTaxable() bool {
	if o != nil && IsNil(o.FreightTaxable) {
		return true
	}

	return false
}

// SetFreightTaxable gets a reference to the given interface{} and assigns it to the FreightTaxable field.
func (o *POSTTaxRules201ResponseDataAttributes) SetFreightTaxable(v interface{}) {
	o.FreightTaxable = v
}

// GetPaymentMethodTaxable returns the PaymentMethodTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetPaymentMethodTaxable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentMethodTaxable
}

// GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetPaymentMethodTaxableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentMethodTaxable) {
		return nil, false
	}
	return &o.PaymentMethodTaxable, true
}

// HasPaymentMethodTaxable returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasPaymentMethodTaxable() bool {
	if o != nil && IsNil(o.PaymentMethodTaxable) {
		return true
	}

	return false
}

// SetPaymentMethodTaxable gets a reference to the given interface{} and assigns it to the PaymentMethodTaxable field.
func (o *POSTTaxRules201ResponseDataAttributes) SetPaymentMethodTaxable(v interface{}) {
	o.PaymentMethodTaxable = v
}

// GetGiftCardTaxable returns the GiftCardTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetGiftCardTaxable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GiftCardTaxable
}

// GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetGiftCardTaxableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GiftCardTaxable) {
		return nil, false
	}
	return &o.GiftCardTaxable, true
}

// HasGiftCardTaxable returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasGiftCardTaxable() bool {
	if o != nil && IsNil(o.GiftCardTaxable) {
		return true
	}

	return false
}

// SetGiftCardTaxable gets a reference to the given interface{} and assigns it to the GiftCardTaxable field.
func (o *POSTTaxRules201ResponseDataAttributes) SetGiftCardTaxable(v interface{}) {
	o.GiftCardTaxable = v
}

// GetAdjustmentTaxable returns the AdjustmentTaxable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetAdjustmentTaxable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AdjustmentTaxable
}

// GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetAdjustmentTaxableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AdjustmentTaxable) {
		return nil, false
	}
	return &o.AdjustmentTaxable, true
}

// HasAdjustmentTaxable returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasAdjustmentTaxable() bool {
	if o != nil && IsNil(o.AdjustmentTaxable) {
		return true
	}

	return false
}

// SetAdjustmentTaxable gets a reference to the given interface{} and assigns it to the AdjustmentTaxable field.
func (o *POSTTaxRules201ResponseDataAttributes) SetAdjustmentTaxable(v interface{}) {
	o.AdjustmentTaxable = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTTaxRules201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTTaxRules201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTTaxRules201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTTaxRules201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTTaxRules201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTTaxRules201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTTaxRules201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTTaxRules201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TaxRate != nil {
		toSerialize["tax_rate"] = o.TaxRate
	}
	if o.CountryCodeRegex != nil {
		toSerialize["country_code_regex"] = o.CountryCodeRegex
	}
	if o.NotCountryCodeRegex != nil {
		toSerialize["not_country_code_regex"] = o.NotCountryCodeRegex
	}
	if o.StateCodeRegex != nil {
		toSerialize["state_code_regex"] = o.StateCodeRegex
	}
	if o.NotStateCodeRegex != nil {
		toSerialize["not_state_code_regex"] = o.NotStateCodeRegex
	}
	if o.ZipCodeRegex != nil {
		toSerialize["zip_code_regex"] = o.ZipCodeRegex
	}
	if o.NotZipCodeRegex != nil {
		toSerialize["not_zip_code_regex"] = o.NotZipCodeRegex
	}
	if o.FreightTaxable != nil {
		toSerialize["freight_taxable"] = o.FreightTaxable
	}
	if o.PaymentMethodTaxable != nil {
		toSerialize["payment_method_taxable"] = o.PaymentMethodTaxable
	}
	if o.GiftCardTaxable != nil {
		toSerialize["gift_card_taxable"] = o.GiftCardTaxable
	}
	if o.AdjustmentTaxable != nil {
		toSerialize["adjustment_taxable"] = o.AdjustmentTaxable
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePOSTTaxRules201ResponseDataAttributes struct {
	value *POSTTaxRules201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTTaxRules201ResponseDataAttributes) Get() *POSTTaxRules201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTTaxRules201ResponseDataAttributes) Set(val *POSTTaxRules201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTTaxRules201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTTaxRules201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTTaxRules201ResponseDataAttributes(val *POSTTaxRules201ResponseDataAttributes) *NullablePOSTTaxRules201ResponseDataAttributes {
	return &NullablePOSTTaxRules201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTTaxRules201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTTaxRules201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
