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

// checks if the POSTReturns201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTReturns201ResponseDataRelationships{}

// POSTReturns201ResponseDataRelationships struct for POSTReturns201ResponseDataRelationships
type POSTReturns201ResponseDataRelationships struct {
	Order              *POSTAdyenPayments201ResponseDataRelationshipsOrder                      `json:"order,omitempty"`
	Customer           *POSTCouponRecipients201ResponseDataRelationshipsCustomer                `json:"customer,omitempty"`
	StockLocation      *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation          `json:"stock_location,omitempty"`
	OriginAddress      *POSTReturns201ResponseDataRelationshipsOriginAddress                    `json:"origin_address,omitempty"`
	DestinationAddress *POSTReturns201ResponseDataRelationshipsDestinationAddress               `json:"destination_address,omitempty"`
	ReferenceCapture   *GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture          `json:"reference_capture,omitempty"`
	ReferenceRefund    *POSTReturns201ResponseDataRelationshipsReferenceRefund                  `json:"reference_refund,omitempty"`
	ReturnLineItems    *POSTLineItems201ResponseDataRelationshipsReturnLineItems                `json:"return_line_items,omitempty"`
	Attachments        *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	ResourceErrors     *POSTOrders201ResponseDataRelationshipsResourceErrors                    `json:"resource_errors,omitempty"`
	Events             *POSTAddresses201ResponseDataRelationshipsEvents                         `json:"events,omitempty"`
	Tags               *POSTAddresses201ResponseDataRelationshipsTags                           `json:"tags,omitempty"`
	Versions           *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
}

// NewPOSTReturns201ResponseDataRelationships instantiates a new POSTReturns201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTReturns201ResponseDataRelationships() *POSTReturns201ResponseDataRelationships {
	this := POSTReturns201ResponseDataRelationships{}
	return &this
}

// NewPOSTReturns201ResponseDataRelationshipsWithDefaults instantiates a new POSTReturns201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTReturns201ResponseDataRelationshipsWithDefaults() *POSTReturns201ResponseDataRelationships {
	this := POSTReturns201ResponseDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *POSTReturns201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder) {
	o.Order = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret POSTCouponRecipients201ResponseDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given POSTCouponRecipients201ResponseDataRelationshipsCustomer and assigns it to the Customer field.
func (o *POSTReturns201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *POSTReturns201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetOriginAddress returns the OriginAddress field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetOriginAddress() POSTReturns201ResponseDataRelationshipsOriginAddress {
	if o == nil || IsNil(o.OriginAddress) {
		var ret POSTReturns201ResponseDataRelationshipsOriginAddress
		return ret
	}
	return *o.OriginAddress
}

// GetOriginAddressOk returns a tuple with the OriginAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetOriginAddressOk() (*POSTReturns201ResponseDataRelationshipsOriginAddress, bool) {
	if o == nil || IsNil(o.OriginAddress) {
		return nil, false
	}
	return o.OriginAddress, true
}

// HasOriginAddress returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasOriginAddress() bool {
	if o != nil && !IsNil(o.OriginAddress) {
		return true
	}

	return false
}

// SetOriginAddress gets a reference to the given POSTReturns201ResponseDataRelationshipsOriginAddress and assigns it to the OriginAddress field.
func (o *POSTReturns201ResponseDataRelationships) SetOriginAddress(v POSTReturns201ResponseDataRelationshipsOriginAddress) {
	o.OriginAddress = &v
}

// GetDestinationAddress returns the DestinationAddress field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetDestinationAddress() POSTReturns201ResponseDataRelationshipsDestinationAddress {
	if o == nil || IsNil(o.DestinationAddress) {
		var ret POSTReturns201ResponseDataRelationshipsDestinationAddress
		return ret
	}
	return *o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetDestinationAddressOk() (*POSTReturns201ResponseDataRelationshipsDestinationAddress, bool) {
	if o == nil || IsNil(o.DestinationAddress) {
		return nil, false
	}
	return o.DestinationAddress, true
}

// HasDestinationAddress returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasDestinationAddress() bool {
	if o != nil && !IsNil(o.DestinationAddress) {
		return true
	}

	return false
}

// SetDestinationAddress gets a reference to the given POSTReturns201ResponseDataRelationshipsDestinationAddress and assigns it to the DestinationAddress field.
func (o *POSTReturns201ResponseDataRelationships) SetDestinationAddress(v POSTReturns201ResponseDataRelationshipsDestinationAddress) {
	o.DestinationAddress = &v
}

// GetReferenceCapture returns the ReferenceCapture field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetReferenceCapture() GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture {
	if o == nil || IsNil(o.ReferenceCapture) {
		var ret GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture
		return ret
	}
	return *o.ReferenceCapture
}

// GetReferenceCaptureOk returns a tuple with the ReferenceCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetReferenceCaptureOk() (*GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture, bool) {
	if o == nil || IsNil(o.ReferenceCapture) {
		return nil, false
	}
	return o.ReferenceCapture, true
}

// HasReferenceCapture returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasReferenceCapture() bool {
	if o != nil && !IsNil(o.ReferenceCapture) {
		return true
	}

	return false
}

// SetReferenceCapture gets a reference to the given GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture and assigns it to the ReferenceCapture field.
func (o *POSTReturns201ResponseDataRelationships) SetReferenceCapture(v GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture) {
	o.ReferenceCapture = &v
}

// GetReferenceRefund returns the ReferenceRefund field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetReferenceRefund() POSTReturns201ResponseDataRelationshipsReferenceRefund {
	if o == nil || IsNil(o.ReferenceRefund) {
		var ret POSTReturns201ResponseDataRelationshipsReferenceRefund
		return ret
	}
	return *o.ReferenceRefund
}

// GetReferenceRefundOk returns a tuple with the ReferenceRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetReferenceRefundOk() (*POSTReturns201ResponseDataRelationshipsReferenceRefund, bool) {
	if o == nil || IsNil(o.ReferenceRefund) {
		return nil, false
	}
	return o.ReferenceRefund, true
}

// HasReferenceRefund returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasReferenceRefund() bool {
	if o != nil && !IsNil(o.ReferenceRefund) {
		return true
	}

	return false
}

// SetReferenceRefund gets a reference to the given POSTReturns201ResponseDataRelationshipsReferenceRefund and assigns it to the ReferenceRefund field.
func (o *POSTReturns201ResponseDataRelationships) SetReferenceRefund(v POSTReturns201ResponseDataRelationshipsReferenceRefund) {
	o.ReferenceRefund = &v
}

// GetReturnLineItems returns the ReturnLineItems field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetReturnLineItems() POSTLineItems201ResponseDataRelationshipsReturnLineItems {
	if o == nil || IsNil(o.ReturnLineItems) {
		var ret POSTLineItems201ResponseDataRelationshipsReturnLineItems
		return ret
	}
	return *o.ReturnLineItems
}

// GetReturnLineItemsOk returns a tuple with the ReturnLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetReturnLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsReturnLineItems, bool) {
	if o == nil || IsNil(o.ReturnLineItems) {
		return nil, false
	}
	return o.ReturnLineItems, true
}

// HasReturnLineItems returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasReturnLineItems() bool {
	if o != nil && !IsNil(o.ReturnLineItems) {
		return true
	}

	return false
}

// SetReturnLineItems gets a reference to the given POSTLineItems201ResponseDataRelationshipsReturnLineItems and assigns it to the ReturnLineItems field.
func (o *POSTReturns201ResponseDataRelationships) SetReturnLineItems(v POSTLineItems201ResponseDataRelationshipsReturnLineItems) {
	o.ReturnLineItems = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTReturns201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetResourceErrors returns the ResourceErrors field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetResourceErrors() POSTOrders201ResponseDataRelationshipsResourceErrors {
	if o == nil || IsNil(o.ResourceErrors) {
		var ret POSTOrders201ResponseDataRelationshipsResourceErrors
		return ret
	}
	return *o.ResourceErrors
}

// GetResourceErrorsOk returns a tuple with the ResourceErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetResourceErrorsOk() (*POSTOrders201ResponseDataRelationshipsResourceErrors, bool) {
	if o == nil || IsNil(o.ResourceErrors) {
		return nil, false
	}
	return o.ResourceErrors, true
}

// HasResourceErrors returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasResourceErrors() bool {
	if o != nil && !IsNil(o.ResourceErrors) {
		return true
	}

	return false
}

// SetResourceErrors gets a reference to the given POSTOrders201ResponseDataRelationshipsResourceErrors and assigns it to the ResourceErrors field.
func (o *POSTReturns201ResponseDataRelationships) SetResourceErrors(v POSTOrders201ResponseDataRelationshipsResourceErrors) {
	o.ResourceErrors = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTReturns201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret POSTAddresses201ResponseDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given POSTAddresses201ResponseDataRelationshipsTags and assigns it to the Tags field.
func (o *POSTReturns201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTReturns201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturns201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTReturns201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTReturns201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTReturns201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTReturns201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.OriginAddress) {
		toSerialize["origin_address"] = o.OriginAddress
	}
	if !IsNil(o.DestinationAddress) {
		toSerialize["destination_address"] = o.DestinationAddress
	}
	if !IsNil(o.ReferenceCapture) {
		toSerialize["reference_capture"] = o.ReferenceCapture
	}
	if !IsNil(o.ReferenceRefund) {
		toSerialize["reference_refund"] = o.ReferenceRefund
	}
	if !IsNil(o.ReturnLineItems) {
		toSerialize["return_line_items"] = o.ReturnLineItems
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.ResourceErrors) {
		toSerialize["resource_errors"] = o.ResourceErrors
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

type NullablePOSTReturns201ResponseDataRelationships struct {
	value *POSTReturns201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTReturns201ResponseDataRelationships) Get() *POSTReturns201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTReturns201ResponseDataRelationships) Set(val *POSTReturns201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTReturns201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTReturns201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTReturns201ResponseDataRelationships(val *POSTReturns201ResponseDataRelationships) *NullablePOSTReturns201ResponseDataRelationships {
	return &NullablePOSTReturns201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTReturns201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTReturns201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
