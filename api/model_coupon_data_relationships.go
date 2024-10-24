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

// checks if the CouponDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponDataRelationships{}

// CouponDataRelationships struct for CouponDataRelationships
type CouponDataRelationships struct {
	PromotionRule   *BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule `json:"promotion_rule,omitempty"`
	CouponRecipient *CouponDataRelationshipsCouponRecipient                     `json:"coupon_recipient,omitempty"`
	Events          *AddressDataRelationshipsEvents                             `json:"events,omitempty"`
	Tags            *AddressDataRelationshipsTags                               `json:"tags,omitempty"`
	Versions        *AddressDataRelationshipsVersions                           `json:"versions,omitempty"`
}

// NewCouponDataRelationships instantiates a new CouponDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponDataRelationships() *CouponDataRelationships {
	this := CouponDataRelationships{}
	return &this
}

// NewCouponDataRelationshipsWithDefaults instantiates a new CouponDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponDataRelationshipsWithDefaults() *CouponDataRelationships {
	this := CouponDataRelationships{}
	return &this
}

// GetPromotionRule returns the PromotionRule field value if set, zero value otherwise.
func (o *CouponDataRelationships) GetPromotionRule() BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule {
	if o == nil || IsNil(o.PromotionRule) {
		var ret BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule
		return ret
	}
	return *o.PromotionRule
}

// GetPromotionRuleOk returns a tuple with the PromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponDataRelationships) GetPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule, bool) {
	if o == nil || IsNil(o.PromotionRule) {
		return nil, false
	}
	return o.PromotionRule, true
}

// HasPromotionRule returns a boolean if a field has been set.
func (o *CouponDataRelationships) HasPromotionRule() bool {
	if o != nil && !IsNil(o.PromotionRule) {
		return true
	}

	return false
}

// SetPromotionRule gets a reference to the given BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule and assigns it to the PromotionRule field.
func (o *CouponDataRelationships) SetPromotionRule(v BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule) {
	o.PromotionRule = &v
}

// GetCouponRecipient returns the CouponRecipient field value if set, zero value otherwise.
func (o *CouponDataRelationships) GetCouponRecipient() CouponDataRelationshipsCouponRecipient {
	if o == nil || IsNil(o.CouponRecipient) {
		var ret CouponDataRelationshipsCouponRecipient
		return ret
	}
	return *o.CouponRecipient
}

// GetCouponRecipientOk returns a tuple with the CouponRecipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponDataRelationships) GetCouponRecipientOk() (*CouponDataRelationshipsCouponRecipient, bool) {
	if o == nil || IsNil(o.CouponRecipient) {
		return nil, false
	}
	return o.CouponRecipient, true
}

// HasCouponRecipient returns a boolean if a field has been set.
func (o *CouponDataRelationships) HasCouponRecipient() bool {
	if o != nil && !IsNil(o.CouponRecipient) {
		return true
	}

	return false
}

// SetCouponRecipient gets a reference to the given CouponDataRelationshipsCouponRecipient and assigns it to the CouponRecipient field.
func (o *CouponDataRelationships) SetCouponRecipient(v CouponDataRelationshipsCouponRecipient) {
	o.CouponRecipient = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CouponDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CouponDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *CouponDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CouponDataRelationships) GetTags() AddressDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CouponDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressDataRelationshipsTags and assigns it to the Tags field.
func (o *CouponDataRelationships) SetTags(v AddressDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *CouponDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *CouponDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *CouponDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o CouponDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PromotionRule) {
		toSerialize["promotion_rule"] = o.PromotionRule
	}
	if !IsNil(o.CouponRecipient) {
		toSerialize["coupon_recipient"] = o.CouponRecipient
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableCouponDataRelationships struct {
	value *CouponDataRelationships
	isSet bool
}

func (v NullableCouponDataRelationships) Get() *CouponDataRelationships {
	return v.value
}

func (v *NullableCouponDataRelationships) Set(val *CouponDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponDataRelationships(val *CouponDataRelationships) *NullableCouponDataRelationships {
	return &NullableCouponDataRelationships{value: val, isSet: true}
}

func (v NullableCouponDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
