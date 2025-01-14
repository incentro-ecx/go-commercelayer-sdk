# GETOrdersOrderId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | The order identifier. Can be specified if unique within the organization (for enterprise plans only), default to numeric ID otherwise. Cannot be passed by sales channels. | [optional] 
**AffiliateCode** | Pointer to **interface{}** | The affiliate code, if any, to track commissions using any third party services. | [optional] 
**Autorefresh** | Pointer to **interface{}** | Save this attribute as &#39;false&#39; if you want prevent the order to be refreshed automatically at each change (much faster). | [optional] 
**PlaceAsync** | Pointer to **interface{}** | Save this attribute as &#39;true&#39; if you want perform the place asynchronously. Payment errors, if any, will be collected afterwards. | [optional] 
**Status** | Pointer to **interface{}** | The order status. One of &#39;draft&#39; (default), &#39;pending&#39;, &#39;editing&#39;, &#39;placing&#39;, &#39;placed&#39;, &#39;approved&#39;, or &#39;cancelled&#39;. | [optional] 
**PaymentStatus** | Pointer to **interface{}** | The order payment status. One of &#39;unpaid&#39; (default), &#39;authorized&#39;, &#39;partially_authorized&#39;, &#39;paid&#39;, &#39;partially_paid&#39;, &#39;voided&#39;, &#39;partially_voided&#39;, &#39;refunded&#39;, &#39;partially_refunded&#39;, or &#39;free&#39;. | [optional] 
**FulfillmentStatus** | Pointer to **interface{}** | The order fulfillment status. One of &#39;unfulfilled&#39; (default), &#39;in_progress&#39;, &#39;fulfilled&#39;, or &#39;not_required&#39;. | [optional] 
**Guest** | Pointer to **interface{}** | Indicates if the order has been placed as guest. | [optional] 
**Editable** | Pointer to **interface{}** | Indicates if the order can be edited. | [optional] 
**CustomerEmail** | Pointer to **interface{}** | The email address of the associated customer. When creating or updating an order, this is a shortcut to find or create the associated customer by email. | [optional] 
**LanguageCode** | Pointer to **interface{}** | The preferred language code (ISO 639-1) to be used when communicating with the customer. This can be useful when sending the order to 3rd party marketing tools and CRMs. If the language is supported, the hosted checkout will be localized accordingly. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order&#39;s market. | [optional] 
**TaxIncluded** | Pointer to **interface{}** | Indicates if taxes are included in the order amounts, automatically inherited from the order&#39;s price list. | [optional] 
**TaxRate** | Pointer to **interface{}** | The tax rate for this order (if calculated). | [optional] 
**FreightTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to shipping costs. | [optional] 
**PaymentMethodTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to payment methods costs. | [optional] 
**AdjustmentTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to positive adjustments. | [optional] 
**GiftCardTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to purchased gift cards. | [optional] 
**RequiresBillingInfo** | Pointer to **interface{}** | Indicates if the billing address associated to this order requires billing info to be present. | [optional] 
**CountryCode** | Pointer to **interface{}** | The international 2-letter country code as defined by the ISO 3166-1 standard, automatically inherited from the order&#39;s shipping or billing addresses. | [optional] 
**ShippingCountryCodeLock** | Pointer to **interface{}** | The country code that you want the shipping address to be locked to. This can be useful to make sure the shipping address belongs to a given shipping country, e.g. the one selected in a country selector page. Not relevant if order contains only digital products. | [optional] 
**CouponCode** | Pointer to **interface{}** | The coupon code to be used for the order. If valid, it triggers a promotion adding a discount line item to the order. | [optional] 
**GiftCardCode** | Pointer to **interface{}** | The gift card code (at least the first 8 characters) to be used for the order. If valid, it uses the gift card balance to pay for the order. | [optional] 
**SubtotalAmountCents** | Pointer to **interface{}** | The sum of all the SKU line items total amounts, in cents. | [optional] 
**SubtotalAmountFloat** | Pointer to **interface{}** | The sum of all the SKU line items total amounts, float. | [optional] 
**FormattedSubtotalAmount** | Pointer to **interface{}** | The sum of all the SKU line items total amounts, formatted. | [optional] 
**ShippingAmountCents** | Pointer to **interface{}** | The sum of all the shipping costs, in cents. | [optional] 
**ShippingAmountFloat** | Pointer to **interface{}** | The sum of all the shipping costs, float. | [optional] 
**FormattedShippingAmount** | Pointer to **interface{}** | The sum of all the shipping costs, formatted. | [optional] 
**PaymentMethodAmountCents** | Pointer to **interface{}** | The payment method costs, in cents. | [optional] 
**PaymentMethodAmountFloat** | Pointer to **interface{}** | The payment method costs, float. | [optional] 
**FormattedPaymentMethodAmount** | Pointer to **interface{}** | The payment method costs, formatted. | [optional] 
**DiscountAmountCents** | Pointer to **interface{}** | The sum of all the discounts applied to the order, in cents (negative amount). | [optional] 
**DiscountAmountFloat** | Pointer to **interface{}** | The sum of all the discounts applied to the order, float. | [optional] 
**FormattedDiscountAmount** | Pointer to **interface{}** | The sum of all the discounts applied to the order, formatted. | [optional] 
**AdjustmentAmountCents** | Pointer to **interface{}** | The sum of all the adjustments applied to the order, in cents. | [optional] 
**AdjustmentAmountFloat** | Pointer to **interface{}** | The sum of all the adjustments applied to the order, float. | [optional] 
**FormattedAdjustmentAmount** | Pointer to **interface{}** | The sum of all the adjustments applied to the order, formatted. | [optional] 
**GiftCardAmountCents** | Pointer to **interface{}** | The sum of all the gift_cards applied to the order, in cents. | [optional] 
**GiftCardAmountFloat** | Pointer to **interface{}** | The sum of all the gift_cards applied to the order, float. | [optional] 
**FormattedGiftCardAmount** | Pointer to **interface{}** | The sum of all the gift_cards applied to the order, formatted. | [optional] 
**TotalTaxAmountCents** | Pointer to **interface{}** | The sum of all the taxes applied to the order, in cents. | [optional] 
**TotalTaxAmountFloat** | Pointer to **interface{}** | The sum of all the taxes applied to the order, float. | [optional] 
**FormattedTotalTaxAmount** | Pointer to **interface{}** | The sum of all the taxes applied to the order, formatted. | [optional] 
**SubtotalTaxAmountCents** | Pointer to **interface{}** | The taxes applied to the order&#39;s subtotal, in cents. | [optional] 
**SubtotalTaxAmountFloat** | Pointer to **interface{}** | The taxes applied to the order&#39;s subtotal, float. | [optional] 
**FormattedSubtotalTaxAmount** | Pointer to **interface{}** | The taxes applied to the order&#39;s subtotal, formatted. | [optional] 
**ShippingTaxAmountCents** | Pointer to **interface{}** | The taxes applied to the order&#39;s shipping costs, in cents. | [optional] 
**ShippingTaxAmountFloat** | Pointer to **interface{}** | The taxes applied to the order&#39;s shipping costs, float. | [optional] 
**FormattedShippingTaxAmount** | Pointer to **interface{}** | The taxes applied to the order&#39;s shipping costs, formatted. | [optional] 
**PaymentMethodTaxAmountCents** | Pointer to **interface{}** | The taxes applied to the order&#39;s payment method costs, in cents. | [optional] 
**PaymentMethodTaxAmountFloat** | Pointer to **interface{}** | The taxes applied to the order&#39;s payment method costs, float. | [optional] 
**FormattedPaymentMethodTaxAmount** | Pointer to **interface{}** | The taxes applied to the order&#39;s payment method costs, formatted. | [optional] 
**AdjustmentTaxAmountCents** | Pointer to **interface{}** | The taxes applied to the order adjustments, in cents. | [optional] 
**AdjustmentTaxAmountFloat** | Pointer to **interface{}** | The taxes applied to the order adjustments, float. | [optional] 
**FormattedAdjustmentTaxAmount** | Pointer to **interface{}** | The taxes applied to the order adjustments, formatted. | [optional] 
**TotalAmountCents** | Pointer to **interface{}** | The order&#39;s total amount, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **interface{}** | The order&#39;s total amount, float. | [optional] 
**FormattedTotalAmount** | Pointer to **interface{}** | The order&#39;s total amount, formatted. | [optional] 
**TotalTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s total taxable amount, in cents (without discounts). | [optional] 
**TotalTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s total taxable amount, float. | [optional] 
**FormattedTotalTaxableAmount** | Pointer to **interface{}** | The order&#39;s total taxable amount, formatted. | [optional] 
**SubtotalTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s subtotal taxable amount, in cents (equal to subtotal_amount_cents when prices don&#39;t include taxes). | [optional] 
**SubtotalTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s subtotal taxable amount, float. | [optional] 
**FormattedSubtotalTaxableAmount** | Pointer to **interface{}** | The order&#39;s subtotal taxable amount, formatted. | [optional] 
**ShippingTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s shipping taxable amount, in cents (equal to shipping_amount_cents when prices don&#39;t include taxes). | [optional] 
**ShippingTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s shipping taxable amount, float. | [optional] 
**FormattedShippingTaxableAmount** | Pointer to **interface{}** | The order&#39;s shipping taxable amount, formatted. | [optional] 
**PaymentMethodTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s payment method taxable amount, in cents (equal to payment_method_amount_cents when prices don&#39;t include taxes). | [optional] 
**PaymentMethodTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s payment method taxable amount, float. | [optional] 
**FormattedPaymentMethodTaxableAmount** | Pointer to **interface{}** | The order&#39;s payment method taxable amount, formatted. | [optional] 
**AdjustmentTaxableAmountCents** | Pointer to **interface{}** | The order&#39;s adjustment taxable amount, in cents (equal to discount_adjustment_cents when prices don&#39;t include taxes). | [optional] 
**AdjustmentTaxableAmountFloat** | Pointer to **interface{}** | The order&#39;s adjustment taxable amount, float. | [optional] 
**FormattedAdjustmentTaxableAmount** | Pointer to **interface{}** | The order&#39;s adjustment taxable amount, formatted. | [optional] 
**TotalAmountWithTaxesCents** | Pointer to **interface{}** | The order&#39;s total amount (when prices include taxes) or the order&#39;s total + taxes amount (when prices don&#39;t include taxes, e.g. US Markets or B2B). | [optional] 
**TotalAmountWithTaxesFloat** | Pointer to **interface{}** | The order&#39;s total amount with taxes, float. | [optional] 
**FormattedTotalAmountWithTaxes** | Pointer to **interface{}** | The order&#39;s total amount with taxes, formatted. | [optional] 
**FeesAmountCents** | Pointer to **interface{}** | The fees amount that is applied by Commerce Layer, in cents. | [optional] 
**FeesAmountFloat** | Pointer to **interface{}** | The fees amount that is applied by Commerce Layer, float. | [optional] 
**FormattedFeesAmount** | Pointer to **interface{}** | The fees amount that is applied by Commerce Layer, formatted. | [optional] 
**DutyAmountCents** | Pointer to **interface{}** | The duty amount that is calculated by external services, in cents. | [optional] 
**DutyAmountFloat** | Pointer to **interface{}** | The duty amount that is calculated by external services, float. | [optional] 
**FormattedDutyAmount** | Pointer to **interface{}** | The duty amount that is calculated by external services, formatted. | [optional] 
**PlaceTotalAmountCents** | Pointer to **interface{}** | The total amount at place time, in cents, which is used internally for editing. | [optional] 
**PlaceTotalAmountFloat** | Pointer to **interface{}** | The total amount at place time, float. | [optional] 
**FormattedPlaceTotalAmount** | Pointer to **interface{}** | The total amount at place time, formatted. | [optional] 
**SkusCount** | Pointer to **interface{}** | The total number of SKUs in the order&#39;s line items. This can be useful to display a preview of the customer shopping cart content. | [optional] 
**LineItemOptionsCount** | Pointer to **interface{}** | The total number of line item options. This can be useful to display a preview of the customer shopping cart content. | [optional] 
**ShipmentsCount** | Pointer to **interface{}** | The total number of shipments. This can be useful to manage the shipping method(s) selection during checkout. | [optional] 
**TaxCalculationsCount** | Pointer to **interface{}** | The total number of tax calculations. This can be useful to monitor external tax service usage. | [optional] 
**ValidationsCount** | Pointer to **interface{}** | The total number of external validation performed. This can be useful to monitor if external validation has been triggered. | [optional] 
**ErrorsCount** | Pointer to **interface{}** | The total number of resource errors. | [optional] 
**PaymentSourceDetails** | Pointer to **interface{}** | An object that contains the shareable details of the order&#39;s payment source. | [optional] 
**Token** | Pointer to **interface{}** | A unique token that can be shared more securely instead of the order&#39;s id. | [optional] 
**CartUrl** | Pointer to **interface{}** | The cart url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The return url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**TermsUrl** | Pointer to **interface{}** | The terms and conditions url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**PrivacyUrl** | Pointer to **interface{}** | The privacy policy url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**CheckoutUrl** | Pointer to **interface{}** | The checkout url that was automatically generated for the order. Send the customers to this url to let them checkout the order securely on our hosted checkout application. | [optional] 
**PlacedAt** | Pointer to **interface{}** | Time at which the order was placed. | [optional] 
**ApprovedAt** | Pointer to **interface{}** | Time at which the order was approved. | [optional] 
**CancelledAt** | Pointer to **interface{}** | Time at which the order was cancelled. | [optional] 
**PaymentUpdatedAt** | Pointer to **interface{}** | Time at which the order&#39;s payment status was last updated. | [optional] 
**FulfillmentUpdatedAt** | Pointer to **interface{}** | Time at which the order&#39;s fulfillment status was last updated. | [optional] 
**RefreshedAt** | Pointer to **interface{}** | Last time at which an order was manually refreshed. | [optional] 
**ArchivedAt** | Pointer to **interface{}** | Time at which the resource has been archived. | [optional] 
**SubscriptionCreatedAt** | Pointer to **interface{}** | Time at which the order has been marked to create a subscription from its recurring line items. | [optional] 
**CircuitState** | Pointer to **interface{}** | The circuit breaker state, by default it is &#39;closed&#39;. It can become &#39;open&#39; once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made. | [optional] 
**CircuitFailureCount** | Pointer to **interface{}** | The number of consecutive failures recorded by the circuit breaker associated to this resource, will be reset on first successful call to callback. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETOrdersOrderId200ResponseDataAttributes

`func NewGETOrdersOrderId200ResponseDataAttributes() *GETOrdersOrderId200ResponseDataAttributes`

NewGETOrdersOrderId200ResponseDataAttributes instantiates a new GETOrdersOrderId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrdersOrderId200ResponseDataAttributesWithDefaults

`func NewGETOrdersOrderId200ResponseDataAttributesWithDefaults() *GETOrdersOrderId200ResponseDataAttributes`

NewGETOrdersOrderId200ResponseDataAttributesWithDefaults instantiates a new GETOrdersOrderId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetAffiliateCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAffiliateCode() interface{}`

GetAffiliateCode returns the AffiliateCode field if non-nil, zero value otherwise.

### GetAffiliateCodeOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAffiliateCodeOk() (*interface{}, bool)`

GetAffiliateCodeOk returns a tuple with the AffiliateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAffiliateCode(v interface{})`

SetAffiliateCode sets AffiliateCode field to given value.

### HasAffiliateCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAffiliateCode() bool`

HasAffiliateCode returns a boolean if a field has been set.

### SetAffiliateCodeNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAffiliateCodeNil(b bool)`

 SetAffiliateCodeNil sets the value for AffiliateCode to be an explicit nil

### UnsetAffiliateCode
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAffiliateCode()`

UnsetAffiliateCode ensures that no value is present for AffiliateCode, not even an explicit nil
### GetAutorefresh

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAutorefresh() interface{}`

GetAutorefresh returns the Autorefresh field if non-nil, zero value otherwise.

### GetAutorefreshOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAutorefreshOk() (*interface{}, bool)`

GetAutorefreshOk returns a tuple with the Autorefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorefresh

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAutorefresh(v interface{})`

SetAutorefresh sets Autorefresh field to given value.

### HasAutorefresh

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAutorefresh() bool`

HasAutorefresh returns a boolean if a field has been set.

### SetAutorefreshNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAutorefreshNil(b bool)`

 SetAutorefreshNil sets the value for Autorefresh to be an explicit nil

### UnsetAutorefresh
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAutorefresh()`

UnsetAutorefresh ensures that no value is present for Autorefresh, not even an explicit nil
### GetPlaceAsync

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPlaceAsync() interface{}`

GetPlaceAsync returns the PlaceAsync field if non-nil, zero value otherwise.

### GetPlaceAsyncOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPlaceAsyncOk() (*interface{}, bool)`

GetPlaceAsyncOk returns a tuple with the PlaceAsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceAsync

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPlaceAsync(v interface{})`

SetPlaceAsync sets PlaceAsync field to given value.

### HasPlaceAsync

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPlaceAsync() bool`

HasPlaceAsync returns a boolean if a field has been set.

### SetPlaceAsyncNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPlaceAsyncNil(b bool)`

 SetPlaceAsyncNil sets the value for PlaceAsync to be an explicit nil

### UnsetPlaceAsync
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPlaceAsync()`

UnsetPlaceAsync ensures that no value is present for PlaceAsync, not even an explicit nil
### GetStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPaymentStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentStatus() interface{}`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentStatusOk() (*interface{}, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentStatus(v interface{})`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### SetPaymentStatusNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentStatusNil(b bool)`

 SetPaymentStatusNil sets the value for PaymentStatus to be an explicit nil

### UnsetPaymentStatus
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentStatus()`

UnsetPaymentStatus ensures that no value is present for PaymentStatus, not even an explicit nil
### GetFulfillmentStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFulfillmentStatus() interface{}`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFulfillmentStatusOk() (*interface{}, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFulfillmentStatus(v interface{})`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### SetFulfillmentStatusNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFulfillmentStatusNil(b bool)`

 SetFulfillmentStatusNil sets the value for FulfillmentStatus to be an explicit nil

### UnsetFulfillmentStatus
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFulfillmentStatus()`

UnsetFulfillmentStatus ensures that no value is present for FulfillmentStatus, not even an explicit nil
### GetGuest

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGuest() interface{}`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGuestOk() (*interface{}, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGuest(v interface{})`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### SetGuestNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGuestNil(b bool)`

 SetGuestNil sets the value for Guest to be an explicit nil

### UnsetGuest
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetGuest()`

UnsetGuest ensures that no value is present for Guest, not even an explicit nil
### GetEditable

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetEditable() interface{}`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetEditableOk() (*interface{}, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetEditable(v interface{})`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetCustomerEmail

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetLanguageCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetLanguageCode() interface{}`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetLanguageCodeOk() (*interface{}, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetLanguageCode(v interface{})`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetCurrencyCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetTaxIncluded

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTaxIncluded() interface{}`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTaxIncludedOk() (*interface{}, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTaxIncluded(v interface{})`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### SetTaxIncludedNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTaxIncludedNil(b bool)`

 SetTaxIncludedNil sets the value for TaxIncluded to be an explicit nil

### UnsetTaxIncluded
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTaxIncluded()`

UnsetTaxIncluded ensures that no value is present for TaxIncluded, not even an explicit nil
### GetTaxRate

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTaxRate() interface{}`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTaxRateOk() (*interface{}, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTaxRate(v interface{})`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetFreightTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFreightTaxable() interface{}`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFreightTaxableOk() (*interface{}, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFreightTaxable(v interface{})`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### SetFreightTaxableNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFreightTaxableNil(b bool)`

 SetFreightTaxableNil sets the value for FreightTaxable to be an explicit nil

### UnsetFreightTaxable
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFreightTaxable()`

UnsetFreightTaxable ensures that no value is present for FreightTaxable, not even an explicit nil
### GetPaymentMethodTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxable() interface{}`

GetPaymentMethodTaxable returns the PaymentMethodTaxable field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxableOk() (*interface{}, bool)`

GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxable(v interface{})`

SetPaymentMethodTaxable sets PaymentMethodTaxable field to given value.

### HasPaymentMethodTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentMethodTaxable() bool`

HasPaymentMethodTaxable returns a boolean if a field has been set.

### SetPaymentMethodTaxableNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxableNil(b bool)`

 SetPaymentMethodTaxableNil sets the value for PaymentMethodTaxable to be an explicit nil

### UnsetPaymentMethodTaxable
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentMethodTaxable()`

UnsetPaymentMethodTaxable ensures that no value is present for PaymentMethodTaxable, not even an explicit nil
### GetAdjustmentTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxable() interface{}`

GetAdjustmentTaxable returns the AdjustmentTaxable field if non-nil, zero value otherwise.

### GetAdjustmentTaxableOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxableOk() (*interface{}, bool)`

GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxable(v interface{})`

SetAdjustmentTaxable sets AdjustmentTaxable field to given value.

### HasAdjustmentTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAdjustmentTaxable() bool`

HasAdjustmentTaxable returns a boolean if a field has been set.

### SetAdjustmentTaxableNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxableNil(b bool)`

 SetAdjustmentTaxableNil sets the value for AdjustmentTaxable to be an explicit nil

### UnsetAdjustmentTaxable
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAdjustmentTaxable()`

UnsetAdjustmentTaxable ensures that no value is present for AdjustmentTaxable, not even an explicit nil
### GetGiftCardTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGiftCardTaxable() interface{}`

GetGiftCardTaxable returns the GiftCardTaxable field if non-nil, zero value otherwise.

### GetGiftCardTaxableOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGiftCardTaxableOk() (*interface{}, bool)`

GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGiftCardTaxable(v interface{})`

SetGiftCardTaxable sets GiftCardTaxable field to given value.

### HasGiftCardTaxable

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasGiftCardTaxable() bool`

HasGiftCardTaxable returns a boolean if a field has been set.

### SetGiftCardTaxableNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGiftCardTaxableNil(b bool)`

 SetGiftCardTaxableNil sets the value for GiftCardTaxable to be an explicit nil

### UnsetGiftCardTaxable
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetGiftCardTaxable()`

UnsetGiftCardTaxable ensures that no value is present for GiftCardTaxable, not even an explicit nil
### GetRequiresBillingInfo

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetRequiresBillingInfo() interface{}`

GetRequiresBillingInfo returns the RequiresBillingInfo field if non-nil, zero value otherwise.

### GetRequiresBillingInfoOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetRequiresBillingInfoOk() (*interface{}, bool)`

GetRequiresBillingInfoOk returns a tuple with the RequiresBillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresBillingInfo

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetRequiresBillingInfo(v interface{})`

SetRequiresBillingInfo sets RequiresBillingInfo field to given value.

### HasRequiresBillingInfo

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasRequiresBillingInfo() bool`

HasRequiresBillingInfo returns a boolean if a field has been set.

### SetRequiresBillingInfoNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetRequiresBillingInfoNil(b bool)`

 SetRequiresBillingInfoNil sets the value for RequiresBillingInfo to be an explicit nil

### UnsetRequiresBillingInfo
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetRequiresBillingInfo()`

UnsetRequiresBillingInfo ensures that no value is present for RequiresBillingInfo, not even an explicit nil
### GetCountryCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCountryCode() interface{}`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCountryCodeOk() (*interface{}, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCountryCode(v interface{})`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetShippingCountryCodeLock

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingCountryCodeLock() interface{}`

GetShippingCountryCodeLock returns the ShippingCountryCodeLock field if non-nil, zero value otherwise.

### GetShippingCountryCodeLockOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingCountryCodeLockOk() (*interface{}, bool)`

GetShippingCountryCodeLockOk returns a tuple with the ShippingCountryCodeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryCodeLock

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingCountryCodeLock(v interface{})`

SetShippingCountryCodeLock sets ShippingCountryCodeLock field to given value.

### HasShippingCountryCodeLock

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasShippingCountryCodeLock() bool`

HasShippingCountryCodeLock returns a boolean if a field has been set.

### SetShippingCountryCodeLockNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingCountryCodeLockNil(b bool)`

 SetShippingCountryCodeLockNil sets the value for ShippingCountryCodeLock to be an explicit nil

### UnsetShippingCountryCodeLock
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetShippingCountryCodeLock()`

UnsetShippingCountryCodeLock ensures that no value is present for ShippingCountryCodeLock, not even an explicit nil
### GetCouponCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCouponCode() interface{}`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCouponCodeOk() (*interface{}, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCouponCode(v interface{})`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### SetCouponCodeNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCouponCodeNil(b bool)`

 SetCouponCodeNil sets the value for CouponCode to be an explicit nil

### UnsetCouponCode
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCouponCode()`

UnsetCouponCode ensures that no value is present for CouponCode, not even an explicit nil
### GetGiftCardCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGiftCardCode() interface{}`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGiftCardCodeOk() (*interface{}, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGiftCardCode(v interface{})`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### SetGiftCardCodeNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGiftCardCodeNil(b bool)`

 SetGiftCardCodeNil sets the value for GiftCardCode to be an explicit nil

### UnsetGiftCardCode
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetGiftCardCode()`

UnsetGiftCardCode ensures that no value is present for GiftCardCode, not even an explicit nil
### GetSubtotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalAmountCents() interface{}`

GetSubtotalAmountCents returns the SubtotalAmountCents field if non-nil, zero value otherwise.

### GetSubtotalAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalAmountCentsOk() (*interface{}, bool)`

GetSubtotalAmountCentsOk returns a tuple with the SubtotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalAmountCents(v interface{})`

SetSubtotalAmountCents sets SubtotalAmountCents field to given value.

### HasSubtotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasSubtotalAmountCents() bool`

HasSubtotalAmountCents returns a boolean if a field has been set.

### SetSubtotalAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalAmountCentsNil(b bool)`

 SetSubtotalAmountCentsNil sets the value for SubtotalAmountCents to be an explicit nil

### UnsetSubtotalAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetSubtotalAmountCents()`

UnsetSubtotalAmountCents ensures that no value is present for SubtotalAmountCents, not even an explicit nil
### GetSubtotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalAmountFloat() interface{}`

GetSubtotalAmountFloat returns the SubtotalAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalAmountFloatOk() (*interface{}, bool)`

GetSubtotalAmountFloatOk returns a tuple with the SubtotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalAmountFloat(v interface{})`

SetSubtotalAmountFloat sets SubtotalAmountFloat field to given value.

### HasSubtotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasSubtotalAmountFloat() bool`

HasSubtotalAmountFloat returns a boolean if a field has been set.

### SetSubtotalAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalAmountFloatNil(b bool)`

 SetSubtotalAmountFloatNil sets the value for SubtotalAmountFloat to be an explicit nil

### UnsetSubtotalAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetSubtotalAmountFloat()`

UnsetSubtotalAmountFloat ensures that no value is present for SubtotalAmountFloat, not even an explicit nil
### GetFormattedSubtotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedSubtotalAmount() interface{}`

GetFormattedSubtotalAmount returns the FormattedSubtotalAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedSubtotalAmountOk() (*interface{}, bool)`

GetFormattedSubtotalAmountOk returns a tuple with the FormattedSubtotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedSubtotalAmount(v interface{})`

SetFormattedSubtotalAmount sets FormattedSubtotalAmount field to given value.

### HasFormattedSubtotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedSubtotalAmount() bool`

HasFormattedSubtotalAmount returns a boolean if a field has been set.

### SetFormattedSubtotalAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedSubtotalAmountNil(b bool)`

 SetFormattedSubtotalAmountNil sets the value for FormattedSubtotalAmount to be an explicit nil

### UnsetFormattedSubtotalAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedSubtotalAmount()`

UnsetFormattedSubtotalAmount ensures that no value is present for FormattedSubtotalAmount, not even an explicit nil
### GetShippingAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingAmountCents() interface{}`

GetShippingAmountCents returns the ShippingAmountCents field if non-nil, zero value otherwise.

### GetShippingAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingAmountCentsOk() (*interface{}, bool)`

GetShippingAmountCentsOk returns a tuple with the ShippingAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingAmountCents(v interface{})`

SetShippingAmountCents sets ShippingAmountCents field to given value.

### HasShippingAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasShippingAmountCents() bool`

HasShippingAmountCents returns a boolean if a field has been set.

### SetShippingAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingAmountCentsNil(b bool)`

 SetShippingAmountCentsNil sets the value for ShippingAmountCents to be an explicit nil

### UnsetShippingAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetShippingAmountCents()`

UnsetShippingAmountCents ensures that no value is present for ShippingAmountCents, not even an explicit nil
### GetShippingAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingAmountFloat() interface{}`

GetShippingAmountFloat returns the ShippingAmountFloat field if non-nil, zero value otherwise.

### GetShippingAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingAmountFloatOk() (*interface{}, bool)`

GetShippingAmountFloatOk returns a tuple with the ShippingAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingAmountFloat(v interface{})`

SetShippingAmountFloat sets ShippingAmountFloat field to given value.

### HasShippingAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasShippingAmountFloat() bool`

HasShippingAmountFloat returns a boolean if a field has been set.

### SetShippingAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingAmountFloatNil(b bool)`

 SetShippingAmountFloatNil sets the value for ShippingAmountFloat to be an explicit nil

### UnsetShippingAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetShippingAmountFloat()`

UnsetShippingAmountFloat ensures that no value is present for ShippingAmountFloat, not even an explicit nil
### GetFormattedShippingAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedShippingAmount() interface{}`

GetFormattedShippingAmount returns the FormattedShippingAmount field if non-nil, zero value otherwise.

### GetFormattedShippingAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedShippingAmountOk() (*interface{}, bool)`

GetFormattedShippingAmountOk returns a tuple with the FormattedShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedShippingAmount(v interface{})`

SetFormattedShippingAmount sets FormattedShippingAmount field to given value.

### HasFormattedShippingAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedShippingAmount() bool`

HasFormattedShippingAmount returns a boolean if a field has been set.

### SetFormattedShippingAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedShippingAmountNil(b bool)`

 SetFormattedShippingAmountNil sets the value for FormattedShippingAmount to be an explicit nil

### UnsetFormattedShippingAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedShippingAmount()`

UnsetFormattedShippingAmount ensures that no value is present for FormattedShippingAmount, not even an explicit nil
### GetPaymentMethodAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodAmountCents() interface{}`

GetPaymentMethodAmountCents returns the PaymentMethodAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodAmountCentsOk() (*interface{}, bool)`

GetPaymentMethodAmountCentsOk returns a tuple with the PaymentMethodAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodAmountCents(v interface{})`

SetPaymentMethodAmountCents sets PaymentMethodAmountCents field to given value.

### HasPaymentMethodAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentMethodAmountCents() bool`

HasPaymentMethodAmountCents returns a boolean if a field has been set.

### SetPaymentMethodAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodAmountCentsNil(b bool)`

 SetPaymentMethodAmountCentsNil sets the value for PaymentMethodAmountCents to be an explicit nil

### UnsetPaymentMethodAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentMethodAmountCents()`

UnsetPaymentMethodAmountCents ensures that no value is present for PaymentMethodAmountCents, not even an explicit nil
### GetPaymentMethodAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodAmountFloat() interface{}`

GetPaymentMethodAmountFloat returns the PaymentMethodAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodAmountFloatOk() (*interface{}, bool)`

GetPaymentMethodAmountFloatOk returns a tuple with the PaymentMethodAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodAmountFloat(v interface{})`

SetPaymentMethodAmountFloat sets PaymentMethodAmountFloat field to given value.

### HasPaymentMethodAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentMethodAmountFloat() bool`

HasPaymentMethodAmountFloat returns a boolean if a field has been set.

### SetPaymentMethodAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodAmountFloatNil(b bool)`

 SetPaymentMethodAmountFloatNil sets the value for PaymentMethodAmountFloat to be an explicit nil

### UnsetPaymentMethodAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentMethodAmountFloat()`

UnsetPaymentMethodAmountFloat ensures that no value is present for PaymentMethodAmountFloat, not even an explicit nil
### GetFormattedPaymentMethodAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedPaymentMethodAmount() interface{}`

GetFormattedPaymentMethodAmount returns the FormattedPaymentMethodAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedPaymentMethodAmountOk() (*interface{}, bool)`

GetFormattedPaymentMethodAmountOk returns a tuple with the FormattedPaymentMethodAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedPaymentMethodAmount(v interface{})`

SetFormattedPaymentMethodAmount sets FormattedPaymentMethodAmount field to given value.

### HasFormattedPaymentMethodAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedPaymentMethodAmount() bool`

HasFormattedPaymentMethodAmount returns a boolean if a field has been set.

### SetFormattedPaymentMethodAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedPaymentMethodAmountNil(b bool)`

 SetFormattedPaymentMethodAmountNil sets the value for FormattedPaymentMethodAmount to be an explicit nil

### UnsetFormattedPaymentMethodAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedPaymentMethodAmount()`

UnsetFormattedPaymentMethodAmount ensures that no value is present for FormattedPaymentMethodAmount, not even an explicit nil
### GetDiscountAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetDiscountAmountCents() interface{}`

GetDiscountAmountCents returns the DiscountAmountCents field if non-nil, zero value otherwise.

### GetDiscountAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetDiscountAmountCentsOk() (*interface{}, bool)`

GetDiscountAmountCentsOk returns a tuple with the DiscountAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetDiscountAmountCents(v interface{})`

SetDiscountAmountCents sets DiscountAmountCents field to given value.

### HasDiscountAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasDiscountAmountCents() bool`

HasDiscountAmountCents returns a boolean if a field has been set.

### SetDiscountAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetDiscountAmountCentsNil(b bool)`

 SetDiscountAmountCentsNil sets the value for DiscountAmountCents to be an explicit nil

### UnsetDiscountAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetDiscountAmountCents()`

UnsetDiscountAmountCents ensures that no value is present for DiscountAmountCents, not even an explicit nil
### GetDiscountAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetDiscountAmountFloat() interface{}`

GetDiscountAmountFloat returns the DiscountAmountFloat field if non-nil, zero value otherwise.

### GetDiscountAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetDiscountAmountFloatOk() (*interface{}, bool)`

GetDiscountAmountFloatOk returns a tuple with the DiscountAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetDiscountAmountFloat(v interface{})`

SetDiscountAmountFloat sets DiscountAmountFloat field to given value.

### HasDiscountAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasDiscountAmountFloat() bool`

HasDiscountAmountFloat returns a boolean if a field has been set.

### SetDiscountAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetDiscountAmountFloatNil(b bool)`

 SetDiscountAmountFloatNil sets the value for DiscountAmountFloat to be an explicit nil

### UnsetDiscountAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetDiscountAmountFloat()`

UnsetDiscountAmountFloat ensures that no value is present for DiscountAmountFloat, not even an explicit nil
### GetFormattedDiscountAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedDiscountAmount() interface{}`

GetFormattedDiscountAmount returns the FormattedDiscountAmount field if non-nil, zero value otherwise.

### GetFormattedDiscountAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedDiscountAmountOk() (*interface{}, bool)`

GetFormattedDiscountAmountOk returns a tuple with the FormattedDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDiscountAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedDiscountAmount(v interface{})`

SetFormattedDiscountAmount sets FormattedDiscountAmount field to given value.

### HasFormattedDiscountAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedDiscountAmount() bool`

HasFormattedDiscountAmount returns a boolean if a field has been set.

### SetFormattedDiscountAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedDiscountAmountNil(b bool)`

 SetFormattedDiscountAmountNil sets the value for FormattedDiscountAmount to be an explicit nil

### UnsetFormattedDiscountAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedDiscountAmount()`

UnsetFormattedDiscountAmount ensures that no value is present for FormattedDiscountAmount, not even an explicit nil
### GetAdjustmentAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentAmountCents() interface{}`

GetAdjustmentAmountCents returns the AdjustmentAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentAmountCentsOk() (*interface{}, bool)`

GetAdjustmentAmountCentsOk returns a tuple with the AdjustmentAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentAmountCents(v interface{})`

SetAdjustmentAmountCents sets AdjustmentAmountCents field to given value.

### HasAdjustmentAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAdjustmentAmountCents() bool`

HasAdjustmentAmountCents returns a boolean if a field has been set.

### SetAdjustmentAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentAmountCentsNil(b bool)`

 SetAdjustmentAmountCentsNil sets the value for AdjustmentAmountCents to be an explicit nil

### UnsetAdjustmentAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAdjustmentAmountCents()`

UnsetAdjustmentAmountCents ensures that no value is present for AdjustmentAmountCents, not even an explicit nil
### GetAdjustmentAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentAmountFloat() interface{}`

GetAdjustmentAmountFloat returns the AdjustmentAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentAmountFloatOk() (*interface{}, bool)`

GetAdjustmentAmountFloatOk returns a tuple with the AdjustmentAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentAmountFloat(v interface{})`

SetAdjustmentAmountFloat sets AdjustmentAmountFloat field to given value.

### HasAdjustmentAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAdjustmentAmountFloat() bool`

HasAdjustmentAmountFloat returns a boolean if a field has been set.

### SetAdjustmentAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentAmountFloatNil(b bool)`

 SetAdjustmentAmountFloatNil sets the value for AdjustmentAmountFloat to be an explicit nil

### UnsetAdjustmentAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAdjustmentAmountFloat()`

UnsetAdjustmentAmountFloat ensures that no value is present for AdjustmentAmountFloat, not even an explicit nil
### GetFormattedAdjustmentAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedAdjustmentAmount() interface{}`

GetFormattedAdjustmentAmount returns the FormattedAdjustmentAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedAdjustmentAmountOk() (*interface{}, bool)`

GetFormattedAdjustmentAmountOk returns a tuple with the FormattedAdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedAdjustmentAmount(v interface{})`

SetFormattedAdjustmentAmount sets FormattedAdjustmentAmount field to given value.

### HasFormattedAdjustmentAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedAdjustmentAmount() bool`

HasFormattedAdjustmentAmount returns a boolean if a field has been set.

### SetFormattedAdjustmentAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedAdjustmentAmountNil(b bool)`

 SetFormattedAdjustmentAmountNil sets the value for FormattedAdjustmentAmount to be an explicit nil

### UnsetFormattedAdjustmentAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedAdjustmentAmount()`

UnsetFormattedAdjustmentAmount ensures that no value is present for FormattedAdjustmentAmount, not even an explicit nil
### GetGiftCardAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGiftCardAmountCents() interface{}`

GetGiftCardAmountCents returns the GiftCardAmountCents field if non-nil, zero value otherwise.

### GetGiftCardAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGiftCardAmountCentsOk() (*interface{}, bool)`

GetGiftCardAmountCentsOk returns a tuple with the GiftCardAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGiftCardAmountCents(v interface{})`

SetGiftCardAmountCents sets GiftCardAmountCents field to given value.

### HasGiftCardAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasGiftCardAmountCents() bool`

HasGiftCardAmountCents returns a boolean if a field has been set.

### SetGiftCardAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGiftCardAmountCentsNil(b bool)`

 SetGiftCardAmountCentsNil sets the value for GiftCardAmountCents to be an explicit nil

### UnsetGiftCardAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetGiftCardAmountCents()`

UnsetGiftCardAmountCents ensures that no value is present for GiftCardAmountCents, not even an explicit nil
### GetGiftCardAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGiftCardAmountFloat() interface{}`

GetGiftCardAmountFloat returns the GiftCardAmountFloat field if non-nil, zero value otherwise.

### GetGiftCardAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetGiftCardAmountFloatOk() (*interface{}, bool)`

GetGiftCardAmountFloatOk returns a tuple with the GiftCardAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGiftCardAmountFloat(v interface{})`

SetGiftCardAmountFloat sets GiftCardAmountFloat field to given value.

### HasGiftCardAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasGiftCardAmountFloat() bool`

HasGiftCardAmountFloat returns a boolean if a field has been set.

### SetGiftCardAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetGiftCardAmountFloatNil(b bool)`

 SetGiftCardAmountFloatNil sets the value for GiftCardAmountFloat to be an explicit nil

### UnsetGiftCardAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetGiftCardAmountFloat()`

UnsetGiftCardAmountFloat ensures that no value is present for GiftCardAmountFloat, not even an explicit nil
### GetFormattedGiftCardAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedGiftCardAmount() interface{}`

GetFormattedGiftCardAmount returns the FormattedGiftCardAmount field if non-nil, zero value otherwise.

### GetFormattedGiftCardAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedGiftCardAmountOk() (*interface{}, bool)`

GetFormattedGiftCardAmountOk returns a tuple with the FormattedGiftCardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedGiftCardAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedGiftCardAmount(v interface{})`

SetFormattedGiftCardAmount sets FormattedGiftCardAmount field to given value.

### HasFormattedGiftCardAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedGiftCardAmount() bool`

HasFormattedGiftCardAmount returns a boolean if a field has been set.

### SetFormattedGiftCardAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedGiftCardAmountNil(b bool)`

 SetFormattedGiftCardAmountNil sets the value for FormattedGiftCardAmount to be an explicit nil

### UnsetFormattedGiftCardAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedGiftCardAmount()`

UnsetFormattedGiftCardAmount ensures that no value is present for FormattedGiftCardAmount, not even an explicit nil
### GetTotalTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalTaxAmountCents() interface{}`

GetTotalTaxAmountCents returns the TotalTaxAmountCents field if non-nil, zero value otherwise.

### GetTotalTaxAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalTaxAmountCentsOk() (*interface{}, bool)`

GetTotalTaxAmountCentsOk returns a tuple with the TotalTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalTaxAmountCents(v interface{})`

SetTotalTaxAmountCents sets TotalTaxAmountCents field to given value.

### HasTotalTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTotalTaxAmountCents() bool`

HasTotalTaxAmountCents returns a boolean if a field has been set.

### SetTotalTaxAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalTaxAmountCentsNil(b bool)`

 SetTotalTaxAmountCentsNil sets the value for TotalTaxAmountCents to be an explicit nil

### UnsetTotalTaxAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTotalTaxAmountCents()`

UnsetTotalTaxAmountCents ensures that no value is present for TotalTaxAmountCents, not even an explicit nil
### GetTotalTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalTaxAmountFloat() interface{}`

GetTotalTaxAmountFloat returns the TotalTaxAmountFloat field if non-nil, zero value otherwise.

### GetTotalTaxAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalTaxAmountFloatOk() (*interface{}, bool)`

GetTotalTaxAmountFloatOk returns a tuple with the TotalTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalTaxAmountFloat(v interface{})`

SetTotalTaxAmountFloat sets TotalTaxAmountFloat field to given value.

### HasTotalTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTotalTaxAmountFloat() bool`

HasTotalTaxAmountFloat returns a boolean if a field has been set.

### SetTotalTaxAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalTaxAmountFloatNil(b bool)`

 SetTotalTaxAmountFloatNil sets the value for TotalTaxAmountFloat to be an explicit nil

### UnsetTotalTaxAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTotalTaxAmountFloat()`

UnsetTotalTaxAmountFloat ensures that no value is present for TotalTaxAmountFloat, not even an explicit nil
### GetFormattedTotalTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedTotalTaxAmount() interface{}`

GetFormattedTotalTaxAmount returns the FormattedTotalTaxAmount field if non-nil, zero value otherwise.

### GetFormattedTotalTaxAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedTotalTaxAmountOk() (*interface{}, bool)`

GetFormattedTotalTaxAmountOk returns a tuple with the FormattedTotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedTotalTaxAmount(v interface{})`

SetFormattedTotalTaxAmount sets FormattedTotalTaxAmount field to given value.

### HasFormattedTotalTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedTotalTaxAmount() bool`

HasFormattedTotalTaxAmount returns a boolean if a field has been set.

### SetFormattedTotalTaxAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedTotalTaxAmountNil(b bool)`

 SetFormattedTotalTaxAmountNil sets the value for FormattedTotalTaxAmount to be an explicit nil

### UnsetFormattedTotalTaxAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedTotalTaxAmount()`

UnsetFormattedTotalTaxAmount ensures that no value is present for FormattedTotalTaxAmount, not even an explicit nil
### GetSubtotalTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalTaxAmountCents() interface{}`

GetSubtotalTaxAmountCents returns the SubtotalTaxAmountCents field if non-nil, zero value otherwise.

### GetSubtotalTaxAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalTaxAmountCentsOk() (*interface{}, bool)`

GetSubtotalTaxAmountCentsOk returns a tuple with the SubtotalTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalTaxAmountCents(v interface{})`

SetSubtotalTaxAmountCents sets SubtotalTaxAmountCents field to given value.

### HasSubtotalTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasSubtotalTaxAmountCents() bool`

HasSubtotalTaxAmountCents returns a boolean if a field has been set.

### SetSubtotalTaxAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalTaxAmountCentsNil(b bool)`

 SetSubtotalTaxAmountCentsNil sets the value for SubtotalTaxAmountCents to be an explicit nil

### UnsetSubtotalTaxAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetSubtotalTaxAmountCents()`

UnsetSubtotalTaxAmountCents ensures that no value is present for SubtotalTaxAmountCents, not even an explicit nil
### GetSubtotalTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalTaxAmountFloat() interface{}`

GetSubtotalTaxAmountFloat returns the SubtotalTaxAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalTaxAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalTaxAmountFloatOk() (*interface{}, bool)`

GetSubtotalTaxAmountFloatOk returns a tuple with the SubtotalTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalTaxAmountFloat(v interface{})`

SetSubtotalTaxAmountFloat sets SubtotalTaxAmountFloat field to given value.

### HasSubtotalTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasSubtotalTaxAmountFloat() bool`

HasSubtotalTaxAmountFloat returns a boolean if a field has been set.

### SetSubtotalTaxAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalTaxAmountFloatNil(b bool)`

 SetSubtotalTaxAmountFloatNil sets the value for SubtotalTaxAmountFloat to be an explicit nil

### UnsetSubtotalTaxAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetSubtotalTaxAmountFloat()`

UnsetSubtotalTaxAmountFloat ensures that no value is present for SubtotalTaxAmountFloat, not even an explicit nil
### GetFormattedSubtotalTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedSubtotalTaxAmount() interface{}`

GetFormattedSubtotalTaxAmount returns the FormattedSubtotalTaxAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalTaxAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedSubtotalTaxAmountOk() (*interface{}, bool)`

GetFormattedSubtotalTaxAmountOk returns a tuple with the FormattedSubtotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedSubtotalTaxAmount(v interface{})`

SetFormattedSubtotalTaxAmount sets FormattedSubtotalTaxAmount field to given value.

### HasFormattedSubtotalTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedSubtotalTaxAmount() bool`

HasFormattedSubtotalTaxAmount returns a boolean if a field has been set.

### SetFormattedSubtotalTaxAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedSubtotalTaxAmountNil(b bool)`

 SetFormattedSubtotalTaxAmountNil sets the value for FormattedSubtotalTaxAmount to be an explicit nil

### UnsetFormattedSubtotalTaxAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedSubtotalTaxAmount()`

UnsetFormattedSubtotalTaxAmount ensures that no value is present for FormattedSubtotalTaxAmount, not even an explicit nil
### GetShippingTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingTaxAmountCents() interface{}`

GetShippingTaxAmountCents returns the ShippingTaxAmountCents field if non-nil, zero value otherwise.

### GetShippingTaxAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingTaxAmountCentsOk() (*interface{}, bool)`

GetShippingTaxAmountCentsOk returns a tuple with the ShippingTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingTaxAmountCents(v interface{})`

SetShippingTaxAmountCents sets ShippingTaxAmountCents field to given value.

### HasShippingTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasShippingTaxAmountCents() bool`

HasShippingTaxAmountCents returns a boolean if a field has been set.

### SetShippingTaxAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingTaxAmountCentsNil(b bool)`

 SetShippingTaxAmountCentsNil sets the value for ShippingTaxAmountCents to be an explicit nil

### UnsetShippingTaxAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetShippingTaxAmountCents()`

UnsetShippingTaxAmountCents ensures that no value is present for ShippingTaxAmountCents, not even an explicit nil
### GetShippingTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingTaxAmountFloat() interface{}`

GetShippingTaxAmountFloat returns the ShippingTaxAmountFloat field if non-nil, zero value otherwise.

### GetShippingTaxAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingTaxAmountFloatOk() (*interface{}, bool)`

GetShippingTaxAmountFloatOk returns a tuple with the ShippingTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingTaxAmountFloat(v interface{})`

SetShippingTaxAmountFloat sets ShippingTaxAmountFloat field to given value.

### HasShippingTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasShippingTaxAmountFloat() bool`

HasShippingTaxAmountFloat returns a boolean if a field has been set.

### SetShippingTaxAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingTaxAmountFloatNil(b bool)`

 SetShippingTaxAmountFloatNil sets the value for ShippingTaxAmountFloat to be an explicit nil

### UnsetShippingTaxAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetShippingTaxAmountFloat()`

UnsetShippingTaxAmountFloat ensures that no value is present for ShippingTaxAmountFloat, not even an explicit nil
### GetFormattedShippingTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedShippingTaxAmount() interface{}`

GetFormattedShippingTaxAmount returns the FormattedShippingTaxAmount field if non-nil, zero value otherwise.

### GetFormattedShippingTaxAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedShippingTaxAmountOk() (*interface{}, bool)`

GetFormattedShippingTaxAmountOk returns a tuple with the FormattedShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedShippingTaxAmount(v interface{})`

SetFormattedShippingTaxAmount sets FormattedShippingTaxAmount field to given value.

### HasFormattedShippingTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedShippingTaxAmount() bool`

HasFormattedShippingTaxAmount returns a boolean if a field has been set.

### SetFormattedShippingTaxAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedShippingTaxAmountNil(b bool)`

 SetFormattedShippingTaxAmountNil sets the value for FormattedShippingTaxAmount to be an explicit nil

### UnsetFormattedShippingTaxAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedShippingTaxAmount()`

UnsetFormattedShippingTaxAmount ensures that no value is present for FormattedShippingTaxAmount, not even an explicit nil
### GetPaymentMethodTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxAmountCents() interface{}`

GetPaymentMethodTaxAmountCents returns the PaymentMethodTaxAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodTaxAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxAmountCentsOk() (*interface{}, bool)`

GetPaymentMethodTaxAmountCentsOk returns a tuple with the PaymentMethodTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxAmountCents(v interface{})`

SetPaymentMethodTaxAmountCents sets PaymentMethodTaxAmountCents field to given value.

### HasPaymentMethodTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentMethodTaxAmountCents() bool`

HasPaymentMethodTaxAmountCents returns a boolean if a field has been set.

### SetPaymentMethodTaxAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxAmountCentsNil(b bool)`

 SetPaymentMethodTaxAmountCentsNil sets the value for PaymentMethodTaxAmountCents to be an explicit nil

### UnsetPaymentMethodTaxAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentMethodTaxAmountCents()`

UnsetPaymentMethodTaxAmountCents ensures that no value is present for PaymentMethodTaxAmountCents, not even an explicit nil
### GetPaymentMethodTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxAmountFloat() interface{}`

GetPaymentMethodTaxAmountFloat returns the PaymentMethodTaxAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodTaxAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxAmountFloatOk() (*interface{}, bool)`

GetPaymentMethodTaxAmountFloatOk returns a tuple with the PaymentMethodTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxAmountFloat(v interface{})`

SetPaymentMethodTaxAmountFloat sets PaymentMethodTaxAmountFloat field to given value.

### HasPaymentMethodTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentMethodTaxAmountFloat() bool`

HasPaymentMethodTaxAmountFloat returns a boolean if a field has been set.

### SetPaymentMethodTaxAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxAmountFloatNil(b bool)`

 SetPaymentMethodTaxAmountFloatNil sets the value for PaymentMethodTaxAmountFloat to be an explicit nil

### UnsetPaymentMethodTaxAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentMethodTaxAmountFloat()`

UnsetPaymentMethodTaxAmountFloat ensures that no value is present for PaymentMethodTaxAmountFloat, not even an explicit nil
### GetFormattedPaymentMethodTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedPaymentMethodTaxAmount() interface{}`

GetFormattedPaymentMethodTaxAmount returns the FormattedPaymentMethodTaxAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodTaxAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedPaymentMethodTaxAmountOk() (*interface{}, bool)`

GetFormattedPaymentMethodTaxAmountOk returns a tuple with the FormattedPaymentMethodTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedPaymentMethodTaxAmount(v interface{})`

SetFormattedPaymentMethodTaxAmount sets FormattedPaymentMethodTaxAmount field to given value.

### HasFormattedPaymentMethodTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedPaymentMethodTaxAmount() bool`

HasFormattedPaymentMethodTaxAmount returns a boolean if a field has been set.

### SetFormattedPaymentMethodTaxAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedPaymentMethodTaxAmountNil(b bool)`

 SetFormattedPaymentMethodTaxAmountNil sets the value for FormattedPaymentMethodTaxAmount to be an explicit nil

### UnsetFormattedPaymentMethodTaxAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedPaymentMethodTaxAmount()`

UnsetFormattedPaymentMethodTaxAmount ensures that no value is present for FormattedPaymentMethodTaxAmount, not even an explicit nil
### GetAdjustmentTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxAmountCents() interface{}`

GetAdjustmentTaxAmountCents returns the AdjustmentTaxAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentTaxAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxAmountCentsOk() (*interface{}, bool)`

GetAdjustmentTaxAmountCentsOk returns a tuple with the AdjustmentTaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxAmountCents(v interface{})`

SetAdjustmentTaxAmountCents sets AdjustmentTaxAmountCents field to given value.

### HasAdjustmentTaxAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAdjustmentTaxAmountCents() bool`

HasAdjustmentTaxAmountCents returns a boolean if a field has been set.

### SetAdjustmentTaxAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxAmountCentsNil(b bool)`

 SetAdjustmentTaxAmountCentsNil sets the value for AdjustmentTaxAmountCents to be an explicit nil

### UnsetAdjustmentTaxAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAdjustmentTaxAmountCents()`

UnsetAdjustmentTaxAmountCents ensures that no value is present for AdjustmentTaxAmountCents, not even an explicit nil
### GetAdjustmentTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxAmountFloat() interface{}`

GetAdjustmentTaxAmountFloat returns the AdjustmentTaxAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentTaxAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxAmountFloatOk() (*interface{}, bool)`

GetAdjustmentTaxAmountFloatOk returns a tuple with the AdjustmentTaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxAmountFloat(v interface{})`

SetAdjustmentTaxAmountFloat sets AdjustmentTaxAmountFloat field to given value.

### HasAdjustmentTaxAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAdjustmentTaxAmountFloat() bool`

HasAdjustmentTaxAmountFloat returns a boolean if a field has been set.

### SetAdjustmentTaxAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxAmountFloatNil(b bool)`

 SetAdjustmentTaxAmountFloatNil sets the value for AdjustmentTaxAmountFloat to be an explicit nil

### UnsetAdjustmentTaxAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAdjustmentTaxAmountFloat()`

UnsetAdjustmentTaxAmountFloat ensures that no value is present for AdjustmentTaxAmountFloat, not even an explicit nil
### GetFormattedAdjustmentTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedAdjustmentTaxAmount() interface{}`

GetFormattedAdjustmentTaxAmount returns the FormattedAdjustmentTaxAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentTaxAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedAdjustmentTaxAmountOk() (*interface{}, bool)`

GetFormattedAdjustmentTaxAmountOk returns a tuple with the FormattedAdjustmentTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedAdjustmentTaxAmount(v interface{})`

SetFormattedAdjustmentTaxAmount sets FormattedAdjustmentTaxAmount field to given value.

### HasFormattedAdjustmentTaxAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedAdjustmentTaxAmount() bool`

HasFormattedAdjustmentTaxAmount returns a boolean if a field has been set.

### SetFormattedAdjustmentTaxAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedAdjustmentTaxAmountNil(b bool)`

 SetFormattedAdjustmentTaxAmountNil sets the value for FormattedAdjustmentTaxAmount to be an explicit nil

### UnsetFormattedAdjustmentTaxAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedAdjustmentTaxAmount()`

UnsetFormattedAdjustmentTaxAmount ensures that no value is present for FormattedAdjustmentTaxAmount, not even an explicit nil
### GetTotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalAmountCents() interface{}`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalAmountCentsOk() (*interface{}, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalAmountCents(v interface{})`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### SetTotalAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalAmountCentsNil(b bool)`

 SetTotalAmountCentsNil sets the value for TotalAmountCents to be an explicit nil

### UnsetTotalAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTotalAmountCents()`

UnsetTotalAmountCents ensures that no value is present for TotalAmountCents, not even an explicit nil
### GetTotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalAmountFloat() interface{}`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalAmountFloatOk() (*interface{}, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalAmountFloat(v interface{})`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### SetTotalAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalAmountFloatNil(b bool)`

 SetTotalAmountFloatNil sets the value for TotalAmountFloat to be an explicit nil

### UnsetTotalAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTotalAmountFloat()`

UnsetTotalAmountFloat ensures that no value is present for TotalAmountFloat, not even an explicit nil
### GetFormattedTotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedTotalAmount() interface{}`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedTotalAmountOk() (*interface{}, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedTotalAmount(v interface{})`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### SetFormattedTotalAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedTotalAmountNil(b bool)`

 SetFormattedTotalAmountNil sets the value for FormattedTotalAmount to be an explicit nil

### UnsetFormattedTotalAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedTotalAmount()`

UnsetFormattedTotalAmount ensures that no value is present for FormattedTotalAmount, not even an explicit nil
### GetTotalTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalTaxableAmountCents() interface{}`

GetTotalTaxableAmountCents returns the TotalTaxableAmountCents field if non-nil, zero value otherwise.

### GetTotalTaxableAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalTaxableAmountCentsOk() (*interface{}, bool)`

GetTotalTaxableAmountCentsOk returns a tuple with the TotalTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalTaxableAmountCents(v interface{})`

SetTotalTaxableAmountCents sets TotalTaxableAmountCents field to given value.

### HasTotalTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTotalTaxableAmountCents() bool`

HasTotalTaxableAmountCents returns a boolean if a field has been set.

### SetTotalTaxableAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalTaxableAmountCentsNil(b bool)`

 SetTotalTaxableAmountCentsNil sets the value for TotalTaxableAmountCents to be an explicit nil

### UnsetTotalTaxableAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTotalTaxableAmountCents()`

UnsetTotalTaxableAmountCents ensures that no value is present for TotalTaxableAmountCents, not even an explicit nil
### GetTotalTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalTaxableAmountFloat() interface{}`

GetTotalTaxableAmountFloat returns the TotalTaxableAmountFloat field if non-nil, zero value otherwise.

### GetTotalTaxableAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalTaxableAmountFloatOk() (*interface{}, bool)`

GetTotalTaxableAmountFloatOk returns a tuple with the TotalTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalTaxableAmountFloat(v interface{})`

SetTotalTaxableAmountFloat sets TotalTaxableAmountFloat field to given value.

### HasTotalTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTotalTaxableAmountFloat() bool`

HasTotalTaxableAmountFloat returns a boolean if a field has been set.

### SetTotalTaxableAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalTaxableAmountFloatNil(b bool)`

 SetTotalTaxableAmountFloatNil sets the value for TotalTaxableAmountFloat to be an explicit nil

### UnsetTotalTaxableAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTotalTaxableAmountFloat()`

UnsetTotalTaxableAmountFloat ensures that no value is present for TotalTaxableAmountFloat, not even an explicit nil
### GetFormattedTotalTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedTotalTaxableAmount() interface{}`

GetFormattedTotalTaxableAmount returns the FormattedTotalTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedTotalTaxableAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedTotalTaxableAmountOk() (*interface{}, bool)`

GetFormattedTotalTaxableAmountOk returns a tuple with the FormattedTotalTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedTotalTaxableAmount(v interface{})`

SetFormattedTotalTaxableAmount sets FormattedTotalTaxableAmount field to given value.

### HasFormattedTotalTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedTotalTaxableAmount() bool`

HasFormattedTotalTaxableAmount returns a boolean if a field has been set.

### SetFormattedTotalTaxableAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedTotalTaxableAmountNil(b bool)`

 SetFormattedTotalTaxableAmountNil sets the value for FormattedTotalTaxableAmount to be an explicit nil

### UnsetFormattedTotalTaxableAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedTotalTaxableAmount()`

UnsetFormattedTotalTaxableAmount ensures that no value is present for FormattedTotalTaxableAmount, not even an explicit nil
### GetSubtotalTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalTaxableAmountCents() interface{}`

GetSubtotalTaxableAmountCents returns the SubtotalTaxableAmountCents field if non-nil, zero value otherwise.

### GetSubtotalTaxableAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalTaxableAmountCentsOk() (*interface{}, bool)`

GetSubtotalTaxableAmountCentsOk returns a tuple with the SubtotalTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalTaxableAmountCents(v interface{})`

SetSubtotalTaxableAmountCents sets SubtotalTaxableAmountCents field to given value.

### HasSubtotalTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasSubtotalTaxableAmountCents() bool`

HasSubtotalTaxableAmountCents returns a boolean if a field has been set.

### SetSubtotalTaxableAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalTaxableAmountCentsNil(b bool)`

 SetSubtotalTaxableAmountCentsNil sets the value for SubtotalTaxableAmountCents to be an explicit nil

### UnsetSubtotalTaxableAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetSubtotalTaxableAmountCents()`

UnsetSubtotalTaxableAmountCents ensures that no value is present for SubtotalTaxableAmountCents, not even an explicit nil
### GetSubtotalTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalTaxableAmountFloat() interface{}`

GetSubtotalTaxableAmountFloat returns the SubtotalTaxableAmountFloat field if non-nil, zero value otherwise.

### GetSubtotalTaxableAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubtotalTaxableAmountFloatOk() (*interface{}, bool)`

GetSubtotalTaxableAmountFloatOk returns a tuple with the SubtotalTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalTaxableAmountFloat(v interface{})`

SetSubtotalTaxableAmountFloat sets SubtotalTaxableAmountFloat field to given value.

### HasSubtotalTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasSubtotalTaxableAmountFloat() bool`

HasSubtotalTaxableAmountFloat returns a boolean if a field has been set.

### SetSubtotalTaxableAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubtotalTaxableAmountFloatNil(b bool)`

 SetSubtotalTaxableAmountFloatNil sets the value for SubtotalTaxableAmountFloat to be an explicit nil

### UnsetSubtotalTaxableAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetSubtotalTaxableAmountFloat()`

UnsetSubtotalTaxableAmountFloat ensures that no value is present for SubtotalTaxableAmountFloat, not even an explicit nil
### GetFormattedSubtotalTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedSubtotalTaxableAmount() interface{}`

GetFormattedSubtotalTaxableAmount returns the FormattedSubtotalTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedSubtotalTaxableAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedSubtotalTaxableAmountOk() (*interface{}, bool)`

GetFormattedSubtotalTaxableAmountOk returns a tuple with the FormattedSubtotalTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedSubtotalTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedSubtotalTaxableAmount(v interface{})`

SetFormattedSubtotalTaxableAmount sets FormattedSubtotalTaxableAmount field to given value.

### HasFormattedSubtotalTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedSubtotalTaxableAmount() bool`

HasFormattedSubtotalTaxableAmount returns a boolean if a field has been set.

### SetFormattedSubtotalTaxableAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedSubtotalTaxableAmountNil(b bool)`

 SetFormattedSubtotalTaxableAmountNil sets the value for FormattedSubtotalTaxableAmount to be an explicit nil

### UnsetFormattedSubtotalTaxableAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedSubtotalTaxableAmount()`

UnsetFormattedSubtotalTaxableAmount ensures that no value is present for FormattedSubtotalTaxableAmount, not even an explicit nil
### GetShippingTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingTaxableAmountCents() interface{}`

GetShippingTaxableAmountCents returns the ShippingTaxableAmountCents field if non-nil, zero value otherwise.

### GetShippingTaxableAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingTaxableAmountCentsOk() (*interface{}, bool)`

GetShippingTaxableAmountCentsOk returns a tuple with the ShippingTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingTaxableAmountCents(v interface{})`

SetShippingTaxableAmountCents sets ShippingTaxableAmountCents field to given value.

### HasShippingTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasShippingTaxableAmountCents() bool`

HasShippingTaxableAmountCents returns a boolean if a field has been set.

### SetShippingTaxableAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingTaxableAmountCentsNil(b bool)`

 SetShippingTaxableAmountCentsNil sets the value for ShippingTaxableAmountCents to be an explicit nil

### UnsetShippingTaxableAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetShippingTaxableAmountCents()`

UnsetShippingTaxableAmountCents ensures that no value is present for ShippingTaxableAmountCents, not even an explicit nil
### GetShippingTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingTaxableAmountFloat() interface{}`

GetShippingTaxableAmountFloat returns the ShippingTaxableAmountFloat field if non-nil, zero value otherwise.

### GetShippingTaxableAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShippingTaxableAmountFloatOk() (*interface{}, bool)`

GetShippingTaxableAmountFloatOk returns a tuple with the ShippingTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingTaxableAmountFloat(v interface{})`

SetShippingTaxableAmountFloat sets ShippingTaxableAmountFloat field to given value.

### HasShippingTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasShippingTaxableAmountFloat() bool`

HasShippingTaxableAmountFloat returns a boolean if a field has been set.

### SetShippingTaxableAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShippingTaxableAmountFloatNil(b bool)`

 SetShippingTaxableAmountFloatNil sets the value for ShippingTaxableAmountFloat to be an explicit nil

### UnsetShippingTaxableAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetShippingTaxableAmountFloat()`

UnsetShippingTaxableAmountFloat ensures that no value is present for ShippingTaxableAmountFloat, not even an explicit nil
### GetFormattedShippingTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedShippingTaxableAmount() interface{}`

GetFormattedShippingTaxableAmount returns the FormattedShippingTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedShippingTaxableAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedShippingTaxableAmountOk() (*interface{}, bool)`

GetFormattedShippingTaxableAmountOk returns a tuple with the FormattedShippingTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedShippingTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedShippingTaxableAmount(v interface{})`

SetFormattedShippingTaxableAmount sets FormattedShippingTaxableAmount field to given value.

### HasFormattedShippingTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedShippingTaxableAmount() bool`

HasFormattedShippingTaxableAmount returns a boolean if a field has been set.

### SetFormattedShippingTaxableAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedShippingTaxableAmountNil(b bool)`

 SetFormattedShippingTaxableAmountNil sets the value for FormattedShippingTaxableAmount to be an explicit nil

### UnsetFormattedShippingTaxableAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedShippingTaxableAmount()`

UnsetFormattedShippingTaxableAmount ensures that no value is present for FormattedShippingTaxableAmount, not even an explicit nil
### GetPaymentMethodTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxableAmountCents() interface{}`

GetPaymentMethodTaxableAmountCents returns the PaymentMethodTaxableAmountCents field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxableAmountCentsOk() (*interface{}, bool)`

GetPaymentMethodTaxableAmountCentsOk returns a tuple with the PaymentMethodTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxableAmountCents(v interface{})`

SetPaymentMethodTaxableAmountCents sets PaymentMethodTaxableAmountCents field to given value.

### HasPaymentMethodTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentMethodTaxableAmountCents() bool`

HasPaymentMethodTaxableAmountCents returns a boolean if a field has been set.

### SetPaymentMethodTaxableAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxableAmountCentsNil(b bool)`

 SetPaymentMethodTaxableAmountCentsNil sets the value for PaymentMethodTaxableAmountCents to be an explicit nil

### UnsetPaymentMethodTaxableAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentMethodTaxableAmountCents()`

UnsetPaymentMethodTaxableAmountCents ensures that no value is present for PaymentMethodTaxableAmountCents, not even an explicit nil
### GetPaymentMethodTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxableAmountFloat() interface{}`

GetPaymentMethodTaxableAmountFloat returns the PaymentMethodTaxableAmountFloat field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentMethodTaxableAmountFloatOk() (*interface{}, bool)`

GetPaymentMethodTaxableAmountFloatOk returns a tuple with the PaymentMethodTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxableAmountFloat(v interface{})`

SetPaymentMethodTaxableAmountFloat sets PaymentMethodTaxableAmountFloat field to given value.

### HasPaymentMethodTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentMethodTaxableAmountFloat() bool`

HasPaymentMethodTaxableAmountFloat returns a boolean if a field has been set.

### SetPaymentMethodTaxableAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentMethodTaxableAmountFloatNil(b bool)`

 SetPaymentMethodTaxableAmountFloatNil sets the value for PaymentMethodTaxableAmountFloat to be an explicit nil

### UnsetPaymentMethodTaxableAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentMethodTaxableAmountFloat()`

UnsetPaymentMethodTaxableAmountFloat ensures that no value is present for PaymentMethodTaxableAmountFloat, not even an explicit nil
### GetFormattedPaymentMethodTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedPaymentMethodTaxableAmount() interface{}`

GetFormattedPaymentMethodTaxableAmount returns the FormattedPaymentMethodTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedPaymentMethodTaxableAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedPaymentMethodTaxableAmountOk() (*interface{}, bool)`

GetFormattedPaymentMethodTaxableAmountOk returns a tuple with the FormattedPaymentMethodTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPaymentMethodTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedPaymentMethodTaxableAmount(v interface{})`

SetFormattedPaymentMethodTaxableAmount sets FormattedPaymentMethodTaxableAmount field to given value.

### HasFormattedPaymentMethodTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedPaymentMethodTaxableAmount() bool`

HasFormattedPaymentMethodTaxableAmount returns a boolean if a field has been set.

### SetFormattedPaymentMethodTaxableAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedPaymentMethodTaxableAmountNil(b bool)`

 SetFormattedPaymentMethodTaxableAmountNil sets the value for FormattedPaymentMethodTaxableAmount to be an explicit nil

### UnsetFormattedPaymentMethodTaxableAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedPaymentMethodTaxableAmount()`

UnsetFormattedPaymentMethodTaxableAmount ensures that no value is present for FormattedPaymentMethodTaxableAmount, not even an explicit nil
### GetAdjustmentTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxableAmountCents() interface{}`

GetAdjustmentTaxableAmountCents returns the AdjustmentTaxableAmountCents field if non-nil, zero value otherwise.

### GetAdjustmentTaxableAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxableAmountCentsOk() (*interface{}, bool)`

GetAdjustmentTaxableAmountCentsOk returns a tuple with the AdjustmentTaxableAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxableAmountCents(v interface{})`

SetAdjustmentTaxableAmountCents sets AdjustmentTaxableAmountCents field to given value.

### HasAdjustmentTaxableAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAdjustmentTaxableAmountCents() bool`

HasAdjustmentTaxableAmountCents returns a boolean if a field has been set.

### SetAdjustmentTaxableAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxableAmountCentsNil(b bool)`

 SetAdjustmentTaxableAmountCentsNil sets the value for AdjustmentTaxableAmountCents to be an explicit nil

### UnsetAdjustmentTaxableAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAdjustmentTaxableAmountCents()`

UnsetAdjustmentTaxableAmountCents ensures that no value is present for AdjustmentTaxableAmountCents, not even an explicit nil
### GetAdjustmentTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxableAmountFloat() interface{}`

GetAdjustmentTaxableAmountFloat returns the AdjustmentTaxableAmountFloat field if non-nil, zero value otherwise.

### GetAdjustmentTaxableAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetAdjustmentTaxableAmountFloatOk() (*interface{}, bool)`

GetAdjustmentTaxableAmountFloatOk returns a tuple with the AdjustmentTaxableAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxableAmountFloat(v interface{})`

SetAdjustmentTaxableAmountFloat sets AdjustmentTaxableAmountFloat field to given value.

### HasAdjustmentTaxableAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasAdjustmentTaxableAmountFloat() bool`

HasAdjustmentTaxableAmountFloat returns a boolean if a field has been set.

### SetAdjustmentTaxableAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetAdjustmentTaxableAmountFloatNil(b bool)`

 SetAdjustmentTaxableAmountFloatNil sets the value for AdjustmentTaxableAmountFloat to be an explicit nil

### UnsetAdjustmentTaxableAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetAdjustmentTaxableAmountFloat()`

UnsetAdjustmentTaxableAmountFloat ensures that no value is present for AdjustmentTaxableAmountFloat, not even an explicit nil
### GetFormattedAdjustmentTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedAdjustmentTaxableAmount() interface{}`

GetFormattedAdjustmentTaxableAmount returns the FormattedAdjustmentTaxableAmount field if non-nil, zero value otherwise.

### GetFormattedAdjustmentTaxableAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedAdjustmentTaxableAmountOk() (*interface{}, bool)`

GetFormattedAdjustmentTaxableAmountOk returns a tuple with the FormattedAdjustmentTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAdjustmentTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedAdjustmentTaxableAmount(v interface{})`

SetFormattedAdjustmentTaxableAmount sets FormattedAdjustmentTaxableAmount field to given value.

### HasFormattedAdjustmentTaxableAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedAdjustmentTaxableAmount() bool`

HasFormattedAdjustmentTaxableAmount returns a boolean if a field has been set.

### SetFormattedAdjustmentTaxableAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedAdjustmentTaxableAmountNil(b bool)`

 SetFormattedAdjustmentTaxableAmountNil sets the value for FormattedAdjustmentTaxableAmount to be an explicit nil

### UnsetFormattedAdjustmentTaxableAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedAdjustmentTaxableAmount()`

UnsetFormattedAdjustmentTaxableAmount ensures that no value is present for FormattedAdjustmentTaxableAmount, not even an explicit nil
### GetTotalAmountWithTaxesCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalAmountWithTaxesCents() interface{}`

GetTotalAmountWithTaxesCents returns the TotalAmountWithTaxesCents field if non-nil, zero value otherwise.

### GetTotalAmountWithTaxesCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalAmountWithTaxesCentsOk() (*interface{}, bool)`

GetTotalAmountWithTaxesCentsOk returns a tuple with the TotalAmountWithTaxesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithTaxesCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalAmountWithTaxesCents(v interface{})`

SetTotalAmountWithTaxesCents sets TotalAmountWithTaxesCents field to given value.

### HasTotalAmountWithTaxesCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTotalAmountWithTaxesCents() bool`

HasTotalAmountWithTaxesCents returns a boolean if a field has been set.

### SetTotalAmountWithTaxesCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalAmountWithTaxesCentsNil(b bool)`

 SetTotalAmountWithTaxesCentsNil sets the value for TotalAmountWithTaxesCents to be an explicit nil

### UnsetTotalAmountWithTaxesCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTotalAmountWithTaxesCents()`

UnsetTotalAmountWithTaxesCents ensures that no value is present for TotalAmountWithTaxesCents, not even an explicit nil
### GetTotalAmountWithTaxesFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalAmountWithTaxesFloat() interface{}`

GetTotalAmountWithTaxesFloat returns the TotalAmountWithTaxesFloat field if non-nil, zero value otherwise.

### GetTotalAmountWithTaxesFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTotalAmountWithTaxesFloatOk() (*interface{}, bool)`

GetTotalAmountWithTaxesFloatOk returns a tuple with the TotalAmountWithTaxesFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountWithTaxesFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalAmountWithTaxesFloat(v interface{})`

SetTotalAmountWithTaxesFloat sets TotalAmountWithTaxesFloat field to given value.

### HasTotalAmountWithTaxesFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTotalAmountWithTaxesFloat() bool`

HasTotalAmountWithTaxesFloat returns a boolean if a field has been set.

### SetTotalAmountWithTaxesFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTotalAmountWithTaxesFloatNil(b bool)`

 SetTotalAmountWithTaxesFloatNil sets the value for TotalAmountWithTaxesFloat to be an explicit nil

### UnsetTotalAmountWithTaxesFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTotalAmountWithTaxesFloat()`

UnsetTotalAmountWithTaxesFloat ensures that no value is present for TotalAmountWithTaxesFloat, not even an explicit nil
### GetFormattedTotalAmountWithTaxes

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedTotalAmountWithTaxes() interface{}`

GetFormattedTotalAmountWithTaxes returns the FormattedTotalAmountWithTaxes field if non-nil, zero value otherwise.

### GetFormattedTotalAmountWithTaxesOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedTotalAmountWithTaxesOk() (*interface{}, bool)`

GetFormattedTotalAmountWithTaxesOk returns a tuple with the FormattedTotalAmountWithTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmountWithTaxes

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedTotalAmountWithTaxes(v interface{})`

SetFormattedTotalAmountWithTaxes sets FormattedTotalAmountWithTaxes field to given value.

### HasFormattedTotalAmountWithTaxes

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedTotalAmountWithTaxes() bool`

HasFormattedTotalAmountWithTaxes returns a boolean if a field has been set.

### SetFormattedTotalAmountWithTaxesNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedTotalAmountWithTaxesNil(b bool)`

 SetFormattedTotalAmountWithTaxesNil sets the value for FormattedTotalAmountWithTaxes to be an explicit nil

### UnsetFormattedTotalAmountWithTaxes
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedTotalAmountWithTaxes()`

UnsetFormattedTotalAmountWithTaxes ensures that no value is present for FormattedTotalAmountWithTaxes, not even an explicit nil
### GetFeesAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFeesAmountCents() interface{}`

GetFeesAmountCents returns the FeesAmountCents field if non-nil, zero value otherwise.

### GetFeesAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFeesAmountCentsOk() (*interface{}, bool)`

GetFeesAmountCentsOk returns a tuple with the FeesAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFeesAmountCents(v interface{})`

SetFeesAmountCents sets FeesAmountCents field to given value.

### HasFeesAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFeesAmountCents() bool`

HasFeesAmountCents returns a boolean if a field has been set.

### SetFeesAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFeesAmountCentsNil(b bool)`

 SetFeesAmountCentsNil sets the value for FeesAmountCents to be an explicit nil

### UnsetFeesAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFeesAmountCents()`

UnsetFeesAmountCents ensures that no value is present for FeesAmountCents, not even an explicit nil
### GetFeesAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFeesAmountFloat() interface{}`

GetFeesAmountFloat returns the FeesAmountFloat field if non-nil, zero value otherwise.

### GetFeesAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFeesAmountFloatOk() (*interface{}, bool)`

GetFeesAmountFloatOk returns a tuple with the FeesAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFeesAmountFloat(v interface{})`

SetFeesAmountFloat sets FeesAmountFloat field to given value.

### HasFeesAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFeesAmountFloat() bool`

HasFeesAmountFloat returns a boolean if a field has been set.

### SetFeesAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFeesAmountFloatNil(b bool)`

 SetFeesAmountFloatNil sets the value for FeesAmountFloat to be an explicit nil

### UnsetFeesAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFeesAmountFloat()`

UnsetFeesAmountFloat ensures that no value is present for FeesAmountFloat, not even an explicit nil
### GetFormattedFeesAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedFeesAmount() interface{}`

GetFormattedFeesAmount returns the FormattedFeesAmount field if non-nil, zero value otherwise.

### GetFormattedFeesAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedFeesAmountOk() (*interface{}, bool)`

GetFormattedFeesAmountOk returns a tuple with the FormattedFeesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFeesAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedFeesAmount(v interface{})`

SetFormattedFeesAmount sets FormattedFeesAmount field to given value.

### HasFormattedFeesAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedFeesAmount() bool`

HasFormattedFeesAmount returns a boolean if a field has been set.

### SetFormattedFeesAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedFeesAmountNil(b bool)`

 SetFormattedFeesAmountNil sets the value for FormattedFeesAmount to be an explicit nil

### UnsetFormattedFeesAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedFeesAmount()`

UnsetFormattedFeesAmount ensures that no value is present for FormattedFeesAmount, not even an explicit nil
### GetDutyAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetDutyAmountCents() interface{}`

GetDutyAmountCents returns the DutyAmountCents field if non-nil, zero value otherwise.

### GetDutyAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetDutyAmountCentsOk() (*interface{}, bool)`

GetDutyAmountCentsOk returns a tuple with the DutyAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetDutyAmountCents(v interface{})`

SetDutyAmountCents sets DutyAmountCents field to given value.

### HasDutyAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasDutyAmountCents() bool`

HasDutyAmountCents returns a boolean if a field has been set.

### SetDutyAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetDutyAmountCentsNil(b bool)`

 SetDutyAmountCentsNil sets the value for DutyAmountCents to be an explicit nil

### UnsetDutyAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetDutyAmountCents()`

UnsetDutyAmountCents ensures that no value is present for DutyAmountCents, not even an explicit nil
### GetDutyAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetDutyAmountFloat() interface{}`

GetDutyAmountFloat returns the DutyAmountFloat field if non-nil, zero value otherwise.

### GetDutyAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetDutyAmountFloatOk() (*interface{}, bool)`

GetDutyAmountFloatOk returns a tuple with the DutyAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetDutyAmountFloat(v interface{})`

SetDutyAmountFloat sets DutyAmountFloat field to given value.

### HasDutyAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasDutyAmountFloat() bool`

HasDutyAmountFloat returns a boolean if a field has been set.

### SetDutyAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetDutyAmountFloatNil(b bool)`

 SetDutyAmountFloatNil sets the value for DutyAmountFloat to be an explicit nil

### UnsetDutyAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetDutyAmountFloat()`

UnsetDutyAmountFloat ensures that no value is present for DutyAmountFloat, not even an explicit nil
### GetFormattedDutyAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedDutyAmount() interface{}`

GetFormattedDutyAmount returns the FormattedDutyAmount field if non-nil, zero value otherwise.

### GetFormattedDutyAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedDutyAmountOk() (*interface{}, bool)`

GetFormattedDutyAmountOk returns a tuple with the FormattedDutyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDutyAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedDutyAmount(v interface{})`

SetFormattedDutyAmount sets FormattedDutyAmount field to given value.

### HasFormattedDutyAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedDutyAmount() bool`

HasFormattedDutyAmount returns a boolean if a field has been set.

### SetFormattedDutyAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedDutyAmountNil(b bool)`

 SetFormattedDutyAmountNil sets the value for FormattedDutyAmount to be an explicit nil

### UnsetFormattedDutyAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedDutyAmount()`

UnsetFormattedDutyAmount ensures that no value is present for FormattedDutyAmount, not even an explicit nil
### GetPlaceTotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPlaceTotalAmountCents() interface{}`

GetPlaceTotalAmountCents returns the PlaceTotalAmountCents field if non-nil, zero value otherwise.

### GetPlaceTotalAmountCentsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPlaceTotalAmountCentsOk() (*interface{}, bool)`

GetPlaceTotalAmountCentsOk returns a tuple with the PlaceTotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPlaceTotalAmountCents(v interface{})`

SetPlaceTotalAmountCents sets PlaceTotalAmountCents field to given value.

### HasPlaceTotalAmountCents

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPlaceTotalAmountCents() bool`

HasPlaceTotalAmountCents returns a boolean if a field has been set.

### SetPlaceTotalAmountCentsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPlaceTotalAmountCentsNil(b bool)`

 SetPlaceTotalAmountCentsNil sets the value for PlaceTotalAmountCents to be an explicit nil

### UnsetPlaceTotalAmountCents
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPlaceTotalAmountCents()`

UnsetPlaceTotalAmountCents ensures that no value is present for PlaceTotalAmountCents, not even an explicit nil
### GetPlaceTotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPlaceTotalAmountFloat() interface{}`

GetPlaceTotalAmountFloat returns the PlaceTotalAmountFloat field if non-nil, zero value otherwise.

### GetPlaceTotalAmountFloatOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPlaceTotalAmountFloatOk() (*interface{}, bool)`

GetPlaceTotalAmountFloatOk returns a tuple with the PlaceTotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPlaceTotalAmountFloat(v interface{})`

SetPlaceTotalAmountFloat sets PlaceTotalAmountFloat field to given value.

### HasPlaceTotalAmountFloat

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPlaceTotalAmountFloat() bool`

HasPlaceTotalAmountFloat returns a boolean if a field has been set.

### SetPlaceTotalAmountFloatNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPlaceTotalAmountFloatNil(b bool)`

 SetPlaceTotalAmountFloatNil sets the value for PlaceTotalAmountFloat to be an explicit nil

### UnsetPlaceTotalAmountFloat
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPlaceTotalAmountFloat()`

UnsetPlaceTotalAmountFloat ensures that no value is present for PlaceTotalAmountFloat, not even an explicit nil
### GetFormattedPlaceTotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedPlaceTotalAmount() interface{}`

GetFormattedPlaceTotalAmount returns the FormattedPlaceTotalAmount field if non-nil, zero value otherwise.

### GetFormattedPlaceTotalAmountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFormattedPlaceTotalAmountOk() (*interface{}, bool)`

GetFormattedPlaceTotalAmountOk returns a tuple with the FormattedPlaceTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPlaceTotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedPlaceTotalAmount(v interface{})`

SetFormattedPlaceTotalAmount sets FormattedPlaceTotalAmount field to given value.

### HasFormattedPlaceTotalAmount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFormattedPlaceTotalAmount() bool`

HasFormattedPlaceTotalAmount returns a boolean if a field has been set.

### SetFormattedPlaceTotalAmountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFormattedPlaceTotalAmountNil(b bool)`

 SetFormattedPlaceTotalAmountNil sets the value for FormattedPlaceTotalAmount to be an explicit nil

### UnsetFormattedPlaceTotalAmount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFormattedPlaceTotalAmount()`

UnsetFormattedPlaceTotalAmount ensures that no value is present for FormattedPlaceTotalAmount, not even an explicit nil
### GetSkusCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSkusCount() interface{}`

GetSkusCount returns the SkusCount field if non-nil, zero value otherwise.

### GetSkusCountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSkusCountOk() (*interface{}, bool)`

GetSkusCountOk returns a tuple with the SkusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSkusCount(v interface{})`

SetSkusCount sets SkusCount field to given value.

### HasSkusCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasSkusCount() bool`

HasSkusCount returns a boolean if a field has been set.

### SetSkusCountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSkusCountNil(b bool)`

 SetSkusCountNil sets the value for SkusCount to be an explicit nil

### UnsetSkusCount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetSkusCount()`

UnsetSkusCount ensures that no value is present for SkusCount, not even an explicit nil
### GetLineItemOptionsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetLineItemOptionsCount() interface{}`

GetLineItemOptionsCount returns the LineItemOptionsCount field if non-nil, zero value otherwise.

### GetLineItemOptionsCountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetLineItemOptionsCountOk() (*interface{}, bool)`

GetLineItemOptionsCountOk returns a tuple with the LineItemOptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptionsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetLineItemOptionsCount(v interface{})`

SetLineItemOptionsCount sets LineItemOptionsCount field to given value.

### HasLineItemOptionsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasLineItemOptionsCount() bool`

HasLineItemOptionsCount returns a boolean if a field has been set.

### SetLineItemOptionsCountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetLineItemOptionsCountNil(b bool)`

 SetLineItemOptionsCountNil sets the value for LineItemOptionsCount to be an explicit nil

### UnsetLineItemOptionsCount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetLineItemOptionsCount()`

UnsetLineItemOptionsCount ensures that no value is present for LineItemOptionsCount, not even an explicit nil
### GetShipmentsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShipmentsCount() interface{}`

GetShipmentsCount returns the ShipmentsCount field if non-nil, zero value otherwise.

### GetShipmentsCountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetShipmentsCountOk() (*interface{}, bool)`

GetShipmentsCountOk returns a tuple with the ShipmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShipmentsCount(v interface{})`

SetShipmentsCount sets ShipmentsCount field to given value.

### HasShipmentsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasShipmentsCount() bool`

HasShipmentsCount returns a boolean if a field has been set.

### SetShipmentsCountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetShipmentsCountNil(b bool)`

 SetShipmentsCountNil sets the value for ShipmentsCount to be an explicit nil

### UnsetShipmentsCount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetShipmentsCount()`

UnsetShipmentsCount ensures that no value is present for ShipmentsCount, not even an explicit nil
### GetTaxCalculationsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTaxCalculationsCount() interface{}`

GetTaxCalculationsCount returns the TaxCalculationsCount field if non-nil, zero value otherwise.

### GetTaxCalculationsCountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTaxCalculationsCountOk() (*interface{}, bool)`

GetTaxCalculationsCountOk returns a tuple with the TaxCalculationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTaxCalculationsCount(v interface{})`

SetTaxCalculationsCount sets TaxCalculationsCount field to given value.

### HasTaxCalculationsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTaxCalculationsCount() bool`

HasTaxCalculationsCount returns a boolean if a field has been set.

### SetTaxCalculationsCountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTaxCalculationsCountNil(b bool)`

 SetTaxCalculationsCountNil sets the value for TaxCalculationsCount to be an explicit nil

### UnsetTaxCalculationsCount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTaxCalculationsCount()`

UnsetTaxCalculationsCount ensures that no value is present for TaxCalculationsCount, not even an explicit nil
### GetValidationsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetValidationsCount() interface{}`

GetValidationsCount returns the ValidationsCount field if non-nil, zero value otherwise.

### GetValidationsCountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetValidationsCountOk() (*interface{}, bool)`

GetValidationsCountOk returns a tuple with the ValidationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetValidationsCount(v interface{})`

SetValidationsCount sets ValidationsCount field to given value.

### HasValidationsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasValidationsCount() bool`

HasValidationsCount returns a boolean if a field has been set.

### SetValidationsCountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetValidationsCountNil(b bool)`

 SetValidationsCountNil sets the value for ValidationsCount to be an explicit nil

### UnsetValidationsCount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetValidationsCount()`

UnsetValidationsCount ensures that no value is present for ValidationsCount, not even an explicit nil
### GetErrorsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetErrorsCount() interface{}`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetErrorsCountOk() (*interface{}, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetErrorsCount(v interface{})`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### SetErrorsCountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetErrorsCountNil(b bool)`

 SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil

### UnsetErrorsCount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetErrorsCount()`

UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
### GetPaymentSourceDetails

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentSourceDetails() interface{}`

GetPaymentSourceDetails returns the PaymentSourceDetails field if non-nil, zero value otherwise.

### GetPaymentSourceDetailsOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentSourceDetailsOk() (*interface{}, bool)`

GetPaymentSourceDetailsOk returns a tuple with the PaymentSourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceDetails

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentSourceDetails(v interface{})`

SetPaymentSourceDetails sets PaymentSourceDetails field to given value.

### HasPaymentSourceDetails

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentSourceDetails() bool`

HasPaymentSourceDetails returns a boolean if a field has been set.

### SetPaymentSourceDetailsNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentSourceDetailsNil(b bool)`

 SetPaymentSourceDetailsNil sets the value for PaymentSourceDetails to be an explicit nil

### UnsetPaymentSourceDetails
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentSourceDetails()`

UnsetPaymentSourceDetails ensures that no value is present for PaymentSourceDetails, not even an explicit nil
### GetToken

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetCartUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCartUrl() interface{}`

GetCartUrl returns the CartUrl field if non-nil, zero value otherwise.

### GetCartUrlOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCartUrlOk() (*interface{}, bool)`

GetCartUrlOk returns a tuple with the CartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCartUrl(v interface{})`

SetCartUrl sets CartUrl field to given value.

### HasCartUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCartUrl() bool`

HasCartUrl returns a boolean if a field has been set.

### SetCartUrlNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCartUrlNil(b bool)`

 SetCartUrlNil sets the value for CartUrl to be an explicit nil

### UnsetCartUrl
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCartUrl()`

UnsetCartUrl ensures that no value is present for CartUrl, not even an explicit nil
### GetReturnUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetTermsUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTermsUrl() interface{}`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetTermsUrlOk() (*interface{}, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTermsUrl(v interface{})`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### SetTermsUrlNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetTermsUrlNil(b bool)`

 SetTermsUrlNil sets the value for TermsUrl to be an explicit nil

### UnsetTermsUrl
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetTermsUrl()`

UnsetTermsUrl ensures that no value is present for TermsUrl, not even an explicit nil
### GetPrivacyUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPrivacyUrl() interface{}`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPrivacyUrlOk() (*interface{}, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPrivacyUrl(v interface{})`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrlNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetCheckoutUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCheckoutUrl() interface{}`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCheckoutUrlOk() (*interface{}, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCheckoutUrl(v interface{})`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### SetCheckoutUrlNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCheckoutUrlNil(b bool)`

 SetCheckoutUrlNil sets the value for CheckoutUrl to be an explicit nil

### UnsetCheckoutUrl
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCheckoutUrl()`

UnsetCheckoutUrl ensures that no value is present for CheckoutUrl, not even an explicit nil
### GetPlacedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPlacedAt() interface{}`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPlacedAtOk() (*interface{}, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPlacedAt(v interface{})`

SetPlacedAt sets PlacedAt field to given value.

### HasPlacedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPlacedAt() bool`

HasPlacedAt returns a boolean if a field has been set.

### SetPlacedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPlacedAtNil(b bool)`

 SetPlacedAtNil sets the value for PlacedAt to be an explicit nil

### UnsetPlacedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPlacedAt()`

UnsetPlacedAt ensures that no value is present for PlacedAt, not even an explicit nil
### GetApprovedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetApprovedAt() interface{}`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetApprovedAtOk() (*interface{}, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetApprovedAt(v interface{})`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### SetApprovedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetApprovedAtNil(b bool)`

 SetApprovedAtNil sets the value for ApprovedAt to be an explicit nil

### UnsetApprovedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetApprovedAt()`

UnsetApprovedAt ensures that no value is present for ApprovedAt, not even an explicit nil
### GetCancelledAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCancelledAt() interface{}`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCancelledAtOk() (*interface{}, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCancelledAt(v interface{})`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetPaymentUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentUpdatedAt() interface{}`

GetPaymentUpdatedAt returns the PaymentUpdatedAt field if non-nil, zero value otherwise.

### GetPaymentUpdatedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetPaymentUpdatedAtOk() (*interface{}, bool)`

GetPaymentUpdatedAtOk returns a tuple with the PaymentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentUpdatedAt(v interface{})`

SetPaymentUpdatedAt sets PaymentUpdatedAt field to given value.

### HasPaymentUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasPaymentUpdatedAt() bool`

HasPaymentUpdatedAt returns a boolean if a field has been set.

### SetPaymentUpdatedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetPaymentUpdatedAtNil(b bool)`

 SetPaymentUpdatedAtNil sets the value for PaymentUpdatedAt to be an explicit nil

### UnsetPaymentUpdatedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetPaymentUpdatedAt()`

UnsetPaymentUpdatedAt ensures that no value is present for PaymentUpdatedAt, not even an explicit nil
### GetFulfillmentUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFulfillmentUpdatedAt() interface{}`

GetFulfillmentUpdatedAt returns the FulfillmentUpdatedAt field if non-nil, zero value otherwise.

### GetFulfillmentUpdatedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetFulfillmentUpdatedAtOk() (*interface{}, bool)`

GetFulfillmentUpdatedAtOk returns a tuple with the FulfillmentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFulfillmentUpdatedAt(v interface{})`

SetFulfillmentUpdatedAt sets FulfillmentUpdatedAt field to given value.

### HasFulfillmentUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasFulfillmentUpdatedAt() bool`

HasFulfillmentUpdatedAt returns a boolean if a field has been set.

### SetFulfillmentUpdatedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetFulfillmentUpdatedAtNil(b bool)`

 SetFulfillmentUpdatedAtNil sets the value for FulfillmentUpdatedAt to be an explicit nil

### UnsetFulfillmentUpdatedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetFulfillmentUpdatedAt()`

UnsetFulfillmentUpdatedAt ensures that no value is present for FulfillmentUpdatedAt, not even an explicit nil
### GetRefreshedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetRefreshedAt() interface{}`

GetRefreshedAt returns the RefreshedAt field if non-nil, zero value otherwise.

### GetRefreshedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetRefreshedAtOk() (*interface{}, bool)`

GetRefreshedAtOk returns a tuple with the RefreshedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetRefreshedAt(v interface{})`

SetRefreshedAt sets RefreshedAt field to given value.

### HasRefreshedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasRefreshedAt() bool`

HasRefreshedAt returns a boolean if a field has been set.

### SetRefreshedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetRefreshedAtNil(b bool)`

 SetRefreshedAtNil sets the value for RefreshedAt to be an explicit nil

### UnsetRefreshedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetRefreshedAt()`

UnsetRefreshedAt ensures that no value is present for RefreshedAt, not even an explicit nil
### GetArchivedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetArchivedAt() interface{}`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetArchivedAtOk() (*interface{}, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetArchivedAt(v interface{})`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetSubscriptionCreatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubscriptionCreatedAt() interface{}`

GetSubscriptionCreatedAt returns the SubscriptionCreatedAt field if non-nil, zero value otherwise.

### GetSubscriptionCreatedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetSubscriptionCreatedAtOk() (*interface{}, bool)`

GetSubscriptionCreatedAtOk returns a tuple with the SubscriptionCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCreatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubscriptionCreatedAt(v interface{})`

SetSubscriptionCreatedAt sets SubscriptionCreatedAt field to given value.

### HasSubscriptionCreatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasSubscriptionCreatedAt() bool`

HasSubscriptionCreatedAt returns a boolean if a field has been set.

### SetSubscriptionCreatedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetSubscriptionCreatedAtNil(b bool)`

 SetSubscriptionCreatedAtNil sets the value for SubscriptionCreatedAt to be an explicit nil

### UnsetSubscriptionCreatedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetSubscriptionCreatedAt()`

UnsetSubscriptionCreatedAt ensures that no value is present for SubscriptionCreatedAt, not even an explicit nil
### GetCircuitState

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCircuitState() interface{}`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCircuitStateOk() (*interface{}, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCircuitState(v interface{})`

SetCircuitState sets CircuitState field to given value.

### HasCircuitState

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCircuitState() bool`

HasCircuitState returns a boolean if a field has been set.

### SetCircuitStateNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCircuitStateNil(b bool)`

 SetCircuitStateNil sets the value for CircuitState to be an explicit nil

### UnsetCircuitState
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCircuitState()`

UnsetCircuitState ensures that no value is present for CircuitState, not even an explicit nil
### GetCircuitFailureCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCircuitFailureCount() interface{}`

GetCircuitFailureCount returns the CircuitFailureCount field if non-nil, zero value otherwise.

### GetCircuitFailureCountOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCircuitFailureCountOk() (*interface{}, bool)`

GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitFailureCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCircuitFailureCount(v interface{})`

SetCircuitFailureCount sets CircuitFailureCount field to given value.

### HasCircuitFailureCount

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCircuitFailureCount() bool`

HasCircuitFailureCount returns a boolean if a field has been set.

### SetCircuitFailureCountNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCircuitFailureCountNil(b bool)`

 SetCircuitFailureCountNil sets the value for CircuitFailureCount to be an explicit nil

### UnsetCircuitFailureCount
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCircuitFailureCount()`

UnsetCircuitFailureCount ensures that no value is present for CircuitFailureCount, not even an explicit nil
### GetCreatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrdersOrderId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrdersOrderId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrdersOrderId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrdersOrderId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


