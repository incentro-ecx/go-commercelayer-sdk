# POSTPaymentMethods201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment method&#39;s internal name. | [optional] 
**PaymentSourceType** | **interface{}** | The payment source type. One of &#39;adyen_payments&#39;, &#39;axerve_payments&#39;, &#39;braintree_payments&#39;, &#39;checkout_com_payments&#39;, &#39;credit_cards&#39;, &#39;external_payments&#39;, &#39;klarna_payments&#39;, &#39;paypal_payments&#39;, &#39;satispay_payments&#39;, &#39;stripe_payments&#39;, or &#39;wire_transfers&#39;. | 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Moto** | Pointer to **interface{}** | Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway. | [optional] 
**RequireCapture** | Pointer to **interface{}** | Send this attribute if you want to require the payment capture before fulfillment. | [optional] 
**AutoPlace** | Pointer to **interface{}** | Send this attribute if you want to automatically place the order upon authorization performed asynchronously. | [optional] 
**AutoCapture** | Pointer to **interface{}** | Send this attribute if you want to automatically capture the payment upon authorization. | [optional] 
**PriceAmountCents** | **interface{}** | The payment method&#39;s price, in cents. | 
**AutoCaptureMaxAmountCents** | Pointer to **interface{}** | Send this attribute if you want to limit automatic capture to orders for which the total amount is equal or less than the specified value, in cents. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as enabled. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTPaymentMethods201ResponseDataAttributes

`func NewPOSTPaymentMethods201ResponseDataAttributes(paymentSourceType interface{}, priceAmountCents interface{}, ) *POSTPaymentMethods201ResponseDataAttributes`

NewPOSTPaymentMethods201ResponseDataAttributes instantiates a new POSTPaymentMethods201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentMethods201ResponseDataAttributesWithDefaults

`func NewPOSTPaymentMethods201ResponseDataAttributesWithDefaults() *POSTPaymentMethods201ResponseDataAttributes`

NewPOSTPaymentMethods201ResponseDataAttributesWithDefaults instantiates a new POSTPaymentMethods201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPaymentSourceType

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetPaymentSourceType() interface{}`

GetPaymentSourceType returns the PaymentSourceType field if non-nil, zero value otherwise.

### GetPaymentSourceTypeOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetPaymentSourceTypeOk() (*interface{}, bool)`

GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceType

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetPaymentSourceType(v interface{})`

SetPaymentSourceType sets PaymentSourceType field to given value.


### SetPaymentSourceTypeNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetPaymentSourceTypeNil(b bool)`

 SetPaymentSourceTypeNil sets the value for PaymentSourceType to be an explicit nil

### UnsetPaymentSourceType
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetPaymentSourceType()`

UnsetPaymentSourceType ensures that no value is present for PaymentSourceType, not even an explicit nil
### GetCurrencyCode

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetMoto

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetMoto() interface{}`

GetMoto returns the Moto field if non-nil, zero value otherwise.

### GetMotoOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetMotoOk() (*interface{}, bool)`

GetMotoOk returns a tuple with the Moto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoto

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetMoto(v interface{})`

SetMoto sets Moto field to given value.

### HasMoto

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasMoto() bool`

HasMoto returns a boolean if a field has been set.

### SetMotoNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetMotoNil(b bool)`

 SetMotoNil sets the value for Moto to be an explicit nil

### UnsetMoto
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetMoto()`

UnsetMoto ensures that no value is present for Moto, not even an explicit nil
### GetRequireCapture

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetRequireCapture() interface{}`

GetRequireCapture returns the RequireCapture field if non-nil, zero value otherwise.

### GetRequireCaptureOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetRequireCaptureOk() (*interface{}, bool)`

GetRequireCaptureOk returns a tuple with the RequireCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCapture

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetRequireCapture(v interface{})`

SetRequireCapture sets RequireCapture field to given value.

### HasRequireCapture

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasRequireCapture() bool`

HasRequireCapture returns a boolean if a field has been set.

### SetRequireCaptureNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetRequireCaptureNil(b bool)`

 SetRequireCaptureNil sets the value for RequireCapture to be an explicit nil

### UnsetRequireCapture
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetRequireCapture()`

UnsetRequireCapture ensures that no value is present for RequireCapture, not even an explicit nil
### GetAutoPlace

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoPlace() interface{}`

GetAutoPlace returns the AutoPlace field if non-nil, zero value otherwise.

### GetAutoPlaceOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoPlaceOk() (*interface{}, bool)`

GetAutoPlaceOk returns a tuple with the AutoPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlace

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoPlace(v interface{})`

SetAutoPlace sets AutoPlace field to given value.

### HasAutoPlace

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasAutoPlace() bool`

HasAutoPlace returns a boolean if a field has been set.

### SetAutoPlaceNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoPlaceNil(b bool)`

 SetAutoPlaceNil sets the value for AutoPlace to be an explicit nil

### UnsetAutoPlace
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetAutoPlace()`

UnsetAutoPlace ensures that no value is present for AutoPlace, not even an explicit nil
### GetAutoCapture

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoCapture() interface{}`

GetAutoCapture returns the AutoCapture field if non-nil, zero value otherwise.

### GetAutoCaptureOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoCaptureOk() (*interface{}, bool)`

GetAutoCaptureOk returns a tuple with the AutoCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapture

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoCapture(v interface{})`

SetAutoCapture sets AutoCapture field to given value.

### HasAutoCapture

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasAutoCapture() bool`

HasAutoCapture returns a boolean if a field has been set.

### SetAutoCaptureNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoCaptureNil(b bool)`

 SetAutoCaptureNil sets the value for AutoCapture to be an explicit nil

### UnsetAutoCapture
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetAutoCapture()`

UnsetAutoCapture ensures that no value is present for AutoCapture, not even an explicit nil
### GetPriceAmountCents

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.


### SetPriceAmountCentsNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetAutoCaptureMaxAmountCents

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoCaptureMaxAmountCents() interface{}`

GetAutoCaptureMaxAmountCents returns the AutoCaptureMaxAmountCents field if non-nil, zero value otherwise.

### GetAutoCaptureMaxAmountCentsOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetAutoCaptureMaxAmountCentsOk() (*interface{}, bool)`

GetAutoCaptureMaxAmountCentsOk returns a tuple with the AutoCaptureMaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCaptureMaxAmountCents

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoCaptureMaxAmountCents(v interface{})`

SetAutoCaptureMaxAmountCents sets AutoCaptureMaxAmountCents field to given value.

### HasAutoCaptureMaxAmountCents

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasAutoCaptureMaxAmountCents() bool`

HasAutoCaptureMaxAmountCents returns a boolean if a field has been set.

### SetAutoCaptureMaxAmountCentsNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetAutoCaptureMaxAmountCentsNil(b bool)`

 SetAutoCaptureMaxAmountCentsNil sets the value for AutoCaptureMaxAmountCents to be an explicit nil

### UnsetAutoCaptureMaxAmountCents
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetAutoCaptureMaxAmountCents()`

UnsetAutoCaptureMaxAmountCents ensures that no value is present for AutoCaptureMaxAmountCents, not even an explicit nil
### GetDisable

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetReference

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPaymentMethods201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPaymentMethods201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTPaymentMethods201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTPaymentMethods201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


