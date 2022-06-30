/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CouponCreateDataAttributes struct for CouponCreateDataAttributes
type CouponCreateDataAttributes struct {
	// The coupon code, that uniquely identifies the coupon within the promotion rule.
	Code string `json:"code"`
	// Indicates if the coupon can be used just once per customer.
	CustomerSingleUse *bool `json:"customer_single_use,omitempty"`
	// The total number of times this coupon can be used.
	UsageLimit int32 `json:"usage_limit"`
	// The email address of the associated recipient. When creating or updating a coupon, this is a shortcut to find or create the associated recipient by email.
	RecipientEmail *string `json:"recipient_email,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewCouponCreateDataAttributes instantiates a new CouponCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCreateDataAttributes(code string, usageLimit int32) *CouponCreateDataAttributes {
	this := CouponCreateDataAttributes{}
	this.Code = code
	this.UsageLimit = usageLimit
	return &this
}

// NewCouponCreateDataAttributesWithDefaults instantiates a new CouponCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCreateDataAttributesWithDefaults() *CouponCreateDataAttributes {
	this := CouponCreateDataAttributes{}
	return &this
}

// GetCode returns the Code field value
func (o *CouponCreateDataAttributes) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CouponCreateDataAttributes) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CouponCreateDataAttributes) SetCode(v string) {
	o.Code = v
}

// GetCustomerSingleUse returns the CustomerSingleUse field value if set, zero value otherwise.
func (o *CouponCreateDataAttributes) GetCustomerSingleUse() bool {
	if o == nil || o.CustomerSingleUse == nil {
		var ret bool
		return ret
	}
	return *o.CustomerSingleUse
}

// GetCustomerSingleUseOk returns a tuple with the CustomerSingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreateDataAttributes) GetCustomerSingleUseOk() (*bool, bool) {
	if o == nil || o.CustomerSingleUse == nil {
		return nil, false
	}
	return o.CustomerSingleUse, true
}

// HasCustomerSingleUse returns a boolean if a field has been set.
func (o *CouponCreateDataAttributes) HasCustomerSingleUse() bool {
	if o != nil && o.CustomerSingleUse != nil {
		return true
	}

	return false
}

// SetCustomerSingleUse gets a reference to the given bool and assigns it to the CustomerSingleUse field.
func (o *CouponCreateDataAttributes) SetCustomerSingleUse(v bool) {
	o.CustomerSingleUse = &v
}

// GetUsageLimit returns the UsageLimit field value
func (o *CouponCreateDataAttributes) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value
// and a boolean to check if the value has been set.
func (o *CouponCreateDataAttributes) GetUsageLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageLimit, true
}

// SetUsageLimit sets field value
func (o *CouponCreateDataAttributes) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise.
func (o *CouponCreateDataAttributes) GetRecipientEmail() string {
	if o == nil || o.RecipientEmail == nil {
		var ret string
		return ret
	}
	return *o.RecipientEmail
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreateDataAttributes) GetRecipientEmailOk() (*string, bool) {
	if o == nil || o.RecipientEmail == nil {
		return nil, false
	}
	return o.RecipientEmail, true
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *CouponCreateDataAttributes) HasRecipientEmail() bool {
	if o != nil && o.RecipientEmail != nil {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given string and assigns it to the RecipientEmail field.
func (o *CouponCreateDataAttributes) SetRecipientEmail(v string) {
	o.RecipientEmail = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CouponCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CouponCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CouponCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *CouponCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *CouponCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *CouponCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CouponCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CouponCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CouponCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CouponCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.CustomerSingleUse != nil {
		toSerialize["customer_single_use"] = o.CustomerSingleUse
	}
	if true {
		toSerialize["usage_limit"] = o.UsageLimit
	}
	if o.RecipientEmail != nil {
		toSerialize["recipient_email"] = o.RecipientEmail
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
	return json.Marshal(toSerialize)
}

type NullableCouponCreateDataAttributes struct {
	value *CouponCreateDataAttributes
	isSet bool
}

func (v NullableCouponCreateDataAttributes) Get() *CouponCreateDataAttributes {
	return v.value
}

func (v *NullableCouponCreateDataAttributes) Set(val *CouponCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCreateDataAttributes(val *CouponCreateDataAttributes) *NullableCouponCreateDataAttributes {
	return &NullableCouponCreateDataAttributes{value: val, isSet: true}
}

func (v NullableCouponCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
