# PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The payment gateway&#39;s internal name. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**ConnectedAccount** | Pointer to **interface{}** | The account (if any) for which the funds of the PaymentIntent are intended. | [optional] 
**AutoPayments** | Pointer to **interface{}** | Indicates if the gateway will accept payment methods enabled in the Stripe dashboard. | [optional] 

## Methods

### NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes

`func NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes() *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes`

NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes instantiates a new PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults

`func NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults() *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes`

NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReference

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetConnectedAccount

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetConnectedAccount() interface{}`

GetConnectedAccount returns the ConnectedAccount field if non-nil, zero value otherwise.

### GetConnectedAccountOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetConnectedAccountOk() (*interface{}, bool)`

GetConnectedAccountOk returns a tuple with the ConnectedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAccount

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetConnectedAccount(v interface{})`

SetConnectedAccount sets ConnectedAccount field to given value.

### HasConnectedAccount

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasConnectedAccount() bool`

HasConnectedAccount returns a boolean if a field has been set.

### SetConnectedAccountNil

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetConnectedAccountNil(b bool)`

 SetConnectedAccountNil sets the value for ConnectedAccount to be an explicit nil

### UnsetConnectedAccount
`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetConnectedAccount()`

UnsetConnectedAccount ensures that no value is present for ConnectedAccount, not even an explicit nil
### GetAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetAutoPayments() interface{}`

GetAutoPayments returns the AutoPayments field if non-nil, zero value otherwise.

### GetAutoPaymentsOk

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) GetAutoPaymentsOk() (*interface{}, bool)`

GetAutoPaymentsOk returns a tuple with the AutoPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetAutoPayments(v interface{})`

SetAutoPayments sets AutoPayments field to given value.

### HasAutoPayments

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) HasAutoPayments() bool`

HasAutoPayments returns a boolean if a field has been set.

### SetAutoPaymentsNil

`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) SetAutoPaymentsNil(b bool)`

 SetAutoPaymentsNil sets the value for AutoPayments to be an explicit nil

### UnsetAutoPayments
`func (o *PATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes) UnsetAutoPayments()`

UnsetAutoPayments ensures that no value is present for AutoPayments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


