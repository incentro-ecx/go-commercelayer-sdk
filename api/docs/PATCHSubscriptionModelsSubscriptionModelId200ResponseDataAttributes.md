# PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The subscription model&#39;s internal name. | [optional] 
**Strategy** | Pointer to **interface{}** | The subscription model&#39;s strategy used to generate order subscriptions: one between &#39;by_frequency&#39; (default) and &#39;by_line_items&#39;. | [optional] 
**Frequencies** | Pointer to **interface{}** | An object that contains the frequencies available for this subscription model. Supported ones are &#39;hourly&#39;, &#39;daily&#39;, &#39;weekly&#39;, &#39;monthly&#39;, &#39;two-month&#39;, &#39;three-month&#39;, &#39;four-month&#39;, &#39;six-month&#39;, &#39;yearly&#39;, or your custom crontab expression (min unit is hour). | [optional] 
**AutoActivate** | Pointer to **interface{}** | Indicates if the created subscriptions will be activated considering the placed source order as its first run, default to true. | [optional] 
**AutoCancel** | Pointer to **interface{}** | Indicates if the created subscriptions will be cancelled in case the source order is cancelled, default to false. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes

`func NewPATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes() *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes`

NewPATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes instantiates a new PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributesWithDefaults

`func NewPATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributesWithDefaults() *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes`

NewPATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributesWithDefaults instantiates a new PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStrategy

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetStrategy() interface{}`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetStrategyOk() (*interface{}, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetStrategy(v interface{})`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategyNil

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetStrategyNil(b bool)`

 SetStrategyNil sets the value for Strategy to be an explicit nil

### UnsetStrategy
`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) UnsetStrategy()`

UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil
### GetFrequencies

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetFrequencies() interface{}`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetFrequenciesOk() (*interface{}, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetFrequencies(v interface{})`

SetFrequencies sets Frequencies field to given value.

### HasFrequencies

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) HasFrequencies() bool`

HasFrequencies returns a boolean if a field has been set.

### SetFrequenciesNil

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetFrequenciesNil(b bool)`

 SetFrequenciesNil sets the value for Frequencies to be an explicit nil

### UnsetFrequencies
`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) UnsetFrequencies()`

UnsetFrequencies ensures that no value is present for Frequencies, not even an explicit nil
### GetAutoActivate

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetAutoActivate() interface{}`

GetAutoActivate returns the AutoActivate field if non-nil, zero value otherwise.

### GetAutoActivateOk

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetAutoActivateOk() (*interface{}, bool)`

GetAutoActivateOk returns a tuple with the AutoActivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoActivate

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetAutoActivate(v interface{})`

SetAutoActivate sets AutoActivate field to given value.

### HasAutoActivate

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) HasAutoActivate() bool`

HasAutoActivate returns a boolean if a field has been set.

### SetAutoActivateNil

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetAutoActivateNil(b bool)`

 SetAutoActivateNil sets the value for AutoActivate to be an explicit nil

### UnsetAutoActivate
`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) UnsetAutoActivate()`

UnsetAutoActivate ensures that no value is present for AutoActivate, not even an explicit nil
### GetAutoCancel

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetAutoCancel() interface{}`

GetAutoCancel returns the AutoCancel field if non-nil, zero value otherwise.

### GetAutoCancelOk

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetAutoCancelOk() (*interface{}, bool)`

GetAutoCancelOk returns a tuple with the AutoCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancel

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetAutoCancel(v interface{})`

SetAutoCancel sets AutoCancel field to given value.

### HasAutoCancel

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) HasAutoCancel() bool`

HasAutoCancel returns a boolean if a field has been set.

### SetAutoCancelNil

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetAutoCancelNil(b bool)`

 SetAutoCancelNil sets the value for AutoCancel to be an explicit nil

### UnsetAutoCancel
`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) UnsetAutoCancel()`

UnsetAutoCancel ensures that no value is present for AutoCancel, not even an explicit nil
### GetReference

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


