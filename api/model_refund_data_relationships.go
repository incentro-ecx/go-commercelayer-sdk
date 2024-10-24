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

// checks if the RefundDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefundDataRelationships{}

// RefundDataRelationships struct for RefundDataRelationships
type RefundDataRelationships struct {
	Order            *AdyenPaymentDataRelationshipsOrder        `json:"order,omitempty"`
	Attachments      *AuthorizationDataRelationshipsAttachments `json:"attachments,omitempty"`
	Events           *AddressDataRelationshipsEvents            `json:"events,omitempty"`
	Versions         *AddressDataRelationshipsVersions          `json:"versions,omitempty"`
	ReferenceCapture *AuthorizationDataRelationshipsCaptures    `json:"reference_capture,omitempty"`
	Return           *CaptureDataRelationshipsReturn            `json:"return,omitempty"`
}

// NewRefundDataRelationships instantiates a new RefundDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundDataRelationships() *RefundDataRelationships {
	this := RefundDataRelationships{}
	return &this
}

// NewRefundDataRelationshipsWithDefaults instantiates a new RefundDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundDataRelationshipsWithDefaults() *RefundDataRelationships {
	this := RefundDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *RefundDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *RefundDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *RefundDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *RefundDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetReferenceCapture returns the ReferenceCapture field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetReferenceCapture() AuthorizationDataRelationshipsCaptures {
	if o == nil || IsNil(o.ReferenceCapture) {
		var ret AuthorizationDataRelationshipsCaptures
		return ret
	}
	return *o.ReferenceCapture
}

// GetReferenceCaptureOk returns a tuple with the ReferenceCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetReferenceCaptureOk() (*AuthorizationDataRelationshipsCaptures, bool) {
	if o == nil || IsNil(o.ReferenceCapture) {
		return nil, false
	}
	return o.ReferenceCapture, true
}

// HasReferenceCapture returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasReferenceCapture() bool {
	if o != nil && !IsNil(o.ReferenceCapture) {
		return true
	}

	return false
}

// SetReferenceCapture gets a reference to the given AuthorizationDataRelationshipsCaptures and assigns it to the ReferenceCapture field.
func (o *RefundDataRelationships) SetReferenceCapture(v AuthorizationDataRelationshipsCaptures) {
	o.ReferenceCapture = &v
}

// GetReturn returns the Return field value if set, zero value otherwise.
func (o *RefundDataRelationships) GetReturn() CaptureDataRelationshipsReturn {
	if o == nil || IsNil(o.Return) {
		var ret CaptureDataRelationshipsReturn
		return ret
	}
	return *o.Return
}

// GetReturnOk returns a tuple with the Return field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundDataRelationships) GetReturnOk() (*CaptureDataRelationshipsReturn, bool) {
	if o == nil || IsNil(o.Return) {
		return nil, false
	}
	return o.Return, true
}

// HasReturn returns a boolean if a field has been set.
func (o *RefundDataRelationships) HasReturn() bool {
	if o != nil && !IsNil(o.Return) {
		return true
	}

	return false
}

// SetReturn gets a reference to the given CaptureDataRelationshipsReturn and assigns it to the Return field.
func (o *RefundDataRelationships) SetReturn(v CaptureDataRelationshipsReturn) {
	o.Return = &v
}

func (o RefundDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.ReferenceCapture) {
		toSerialize["reference_capture"] = o.ReferenceCapture
	}
	if !IsNil(o.Return) {
		toSerialize["return"] = o.Return
	}
	return toSerialize, nil
}

type NullableRefundDataRelationships struct {
	value *RefundDataRelationships
	isSet bool
}

func (v NullableRefundDataRelationships) Get() *RefundDataRelationships {
	return v.value
}

func (v *NullableRefundDataRelationships) Set(val *RefundDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundDataRelationships(val *RefundDataRelationships) *NullableRefundDataRelationships {
	return &NullableRefundDataRelationships{value: val, isSet: true}
}

func (v NullableRefundDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
