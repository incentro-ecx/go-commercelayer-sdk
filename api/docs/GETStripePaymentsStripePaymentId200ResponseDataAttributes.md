# GETStripePaymentsStripePaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StripeId** | Pointer to **interface{}** | The Stripe payment intent ID. Required to identify a payment session on stripe. | [optional] 
**ClientSecret** | Pointer to **interface{}** | The Stripe payment intent client secret. Required to create a charge through Stripe.js. | [optional] 
**ChargeId** | Pointer to **interface{}** | The Stripe charge ID. Identifies money movement upon the payment intent confirmation. | [optional] 
**PublishableKey** | Pointer to **interface{}** | The Stripe publishable API key. | [optional] 
**Options** | Pointer to **interface{}** | Stripe payment options: &#39;customer&#39;, &#39;payment_method&#39;, &#39;return_url&#39;, etc. Check Stripe payment intent API for more details. | [optional] 
**PaymentMethod** | Pointer to **interface{}** | Stripe &#39;payment_method&#39;, set by webhook. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the created payment intent. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The URL to return to when a redirect payment is completed. | [optional] 
**ReceiptEmail** | Pointer to **interface{}** | The email address to send the receipt to. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETStripePaymentsStripePaymentId200ResponseDataAttributes

`func NewGETStripePaymentsStripePaymentId200ResponseDataAttributes() *GETStripePaymentsStripePaymentId200ResponseDataAttributes`

NewGETStripePaymentsStripePaymentId200ResponseDataAttributes instantiates a new GETStripePaymentsStripePaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults

`func NewGETStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults() *GETStripePaymentsStripePaymentId200ResponseDataAttributes`

NewGETStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults instantiates a new GETStripePaymentsStripePaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStripeId

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetStripeId() interface{}`

GetStripeId returns the StripeId field if non-nil, zero value otherwise.

### GetStripeIdOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetStripeIdOk() (*interface{}, bool)`

GetStripeIdOk returns a tuple with the StripeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeId

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetStripeId(v interface{})`

SetStripeId sets StripeId field to given value.

### HasStripeId

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasStripeId() bool`

HasStripeId returns a boolean if a field has been set.

### SetStripeIdNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetStripeIdNil(b bool)`

 SetStripeIdNil sets the value for StripeId to be an explicit nil

### UnsetStripeId
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetStripeId()`

UnsetStripeId ensures that no value is present for StripeId, not even an explicit nil
### GetClientSecret

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetChargeId

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetChargeId() interface{}`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetChargeIdOk() (*interface{}, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetChargeId(v interface{})`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### SetChargeIdNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetChargeIdNil(b bool)`

 SetChargeIdNil sets the value for ChargeId to be an explicit nil

### UnsetChargeId
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetChargeId()`

UnsetChargeId ensures that no value is present for ChargeId, not even an explicit nil
### GetPublishableKey

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPublishableKey() interface{}`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPublishableKeyOk() (*interface{}, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPublishableKey(v interface{})`

SetPublishableKey sets PublishableKey field to given value.

### HasPublishableKey

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasPublishableKey() bool`

HasPublishableKey returns a boolean if a field has been set.

### SetPublishableKeyNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPublishableKeyNil(b bool)`

 SetPublishableKeyNil sets the value for PublishableKey to be an explicit nil

### UnsetPublishableKey
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetPublishableKey()`

UnsetPublishableKey ensures that no value is present for PublishableKey, not even an explicit nil
### GetOptions

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetPaymentMethod

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPaymentMethod() interface{}`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPaymentMethodOk() (*interface{}, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPaymentMethod(v interface{})`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetReturnUrl

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetReceiptEmail

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReceiptEmail() interface{}`

GetReceiptEmail returns the ReceiptEmail field if non-nil, zero value otherwise.

### GetReceiptEmailOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReceiptEmailOk() (*interface{}, bool)`

GetReceiptEmailOk returns a tuple with the ReceiptEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptEmail

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReceiptEmail(v interface{})`

SetReceiptEmail sets ReceiptEmail field to given value.

### HasReceiptEmail

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasReceiptEmail() bool`

HasReceiptEmail returns a boolean if a field has been set.

### SetReceiptEmailNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReceiptEmailNil(b bool)`

 SetReceiptEmailNil sets the value for ReceiptEmail to be an explicit nil

### UnsetReceiptEmail
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetReceiptEmail()`

UnsetReceiptEmail ensures that no value is present for ReceiptEmail, not even an explicit nil
### GetPaymentInstrument

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


