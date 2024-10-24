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

// checks if the POSTFlexPromotions201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTFlexPromotions201ResponseDataAttributes{}

// POSTFlexPromotions201ResponseDataAttributes struct for POSTFlexPromotions201ResponseDataAttributes
type POSTFlexPromotions201ResponseDataAttributes struct {
	// The promotion's internal name.
	Name interface{} `json:"name"`
	// Indicates if the promotion will be applied exclusively, based on its priority score.
	Exclusive interface{} `json:"exclusive,omitempty"`
	// The priority assigned to the promotion (lower means higher priority).
	Priority interface{} `json:"priority,omitempty"`
	// The activation date/time of this promotion.
	StartsAt interface{} `json:"starts_at"`
	// The expiration date/time of this promotion (must be after starts_at).
	ExpiresAt interface{} `json:"expires_at"`
	// The total number of times this promotion can be applied. When 'null' it means promotion can be applied infinite times.
	TotalUsageLimit interface{} `json:"total_usage_limit,omitempty"`
	// The discount rule to be applied.
	Rules interface{} `json:"rules"`
	// Send this attribute if you want to mark this resource as disabled.
	Disable interface{} `json:"_disable,omitempty"`
	// Send this attribute if you want to mark this resource as enabled.
	Enable interface{} `json:"_enable,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTFlexPromotions201ResponseDataAttributes instantiates a new POSTFlexPromotions201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFlexPromotions201ResponseDataAttributes(name interface{}, startsAt interface{}, expiresAt interface{}, rules interface{}) *POSTFlexPromotions201ResponseDataAttributes {
	this := POSTFlexPromotions201ResponseDataAttributes{}
	this.Name = name
	this.StartsAt = startsAt
	this.ExpiresAt = expiresAt
	this.Rules = rules
	return &this
}

// NewPOSTFlexPromotions201ResponseDataAttributesWithDefaults instantiates a new POSTFlexPromotions201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFlexPromotions201ResponseDataAttributesWithDefaults() *POSTFlexPromotions201ResponseDataAttributes {
	this := POSTFlexPromotions201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTFlexPromotions201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetExclusive returns the Exclusive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFlexPromotions201ResponseDataAttributes) GetExclusive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Exclusive) {
		return nil, false
	}
	return &o.Exclusive, true
}

// HasExclusive returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataAttributes) HasExclusive() bool {
	if o != nil && IsNil(o.Exclusive) {
		return true
	}

	return false
}

// SetExclusive gets a reference to the given interface{} and assigns it to the Exclusive field.
func (o *POSTFlexPromotions201ResponseDataAttributes) SetExclusive(v interface{}) {
	o.Exclusive = v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFlexPromotions201ResponseDataAttributes) GetPriority() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetPriorityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return &o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataAttributes) HasPriority() bool {
	if o != nil && IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given interface{} and assigns it to the Priority field.
func (o *POSTFlexPromotions201ResponseDataAttributes) SetPriority(v interface{}) {
	o.Priority = v
}

// GetStartsAt returns the StartsAt field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// SetStartsAt sets field value
func (o *POSTFlexPromotions201ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *POSTFlexPromotions201ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFlexPromotions201ResponseDataAttributes) GetTotalUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageLimit) {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataAttributes) HasTotalUsageLimit() bool {
	if o != nil && IsNil(o.TotalUsageLimit) {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given interface{} and assigns it to the TotalUsageLimit field.
func (o *POSTFlexPromotions201ResponseDataAttributes) SetTotalUsageLimit(v interface{}) {
	o.TotalUsageLimit = v
}

// GetRules returns the Rules field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetRules() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetRulesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *POSTFlexPromotions201ResponseDataAttributes) SetRules(v interface{}) {
	o.Rules = v
}

// GetDisable returns the Disable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFlexPromotions201ResponseDataAttributes) GetDisable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetDisableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return &o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataAttributes) HasDisable() bool {
	if o != nil && IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given interface{} and assigns it to the Disable field.
func (o *POSTFlexPromotions201ResponseDataAttributes) SetDisable(v interface{}) {
	o.Disable = v
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFlexPromotions201ResponseDataAttributes) GetEnable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetEnableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return &o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataAttributes) HasEnable() bool {
	if o != nil && IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given interface{} and assigns it to the Enable field.
func (o *POSTFlexPromotions201ResponseDataAttributes) SetEnable(v interface{}) {
	o.Enable = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFlexPromotions201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTFlexPromotions201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFlexPromotions201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTFlexPromotions201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFlexPromotions201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFlexPromotions201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTFlexPromotions201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTFlexPromotions201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTFlexPromotions201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTFlexPromotions201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Exclusive != nil {
		toSerialize["exclusive"] = o.Exclusive
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.TotalUsageLimit != nil {
		toSerialize["total_usage_limit"] = o.TotalUsageLimit
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Disable != nil {
		toSerialize["_disable"] = o.Disable
	}
	if o.Enable != nil {
		toSerialize["_enable"] = o.Enable
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

type NullablePOSTFlexPromotions201ResponseDataAttributes struct {
	value *POSTFlexPromotions201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTFlexPromotions201ResponseDataAttributes) Get() *POSTFlexPromotions201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTFlexPromotions201ResponseDataAttributes) Set(val *POSTFlexPromotions201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFlexPromotions201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFlexPromotions201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFlexPromotions201ResponseDataAttributes(val *POSTFlexPromotions201ResponseDataAttributes) *NullablePOSTFlexPromotions201ResponseDataAttributes {
	return &NullablePOSTFlexPromotions201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTFlexPromotions201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFlexPromotions201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
