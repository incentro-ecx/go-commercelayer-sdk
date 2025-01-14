# POSTNotifications201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The internal name of the notification. | 
**Flash** | Pointer to **interface{}** | Indicates if the notification is temporary, valid for the ones created by external services. | [optional] 
**Body** | Pointer to **interface{}** | An internal body of the notification. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTNotifications201ResponseDataAttributes

`func NewPOSTNotifications201ResponseDataAttributes(name interface{}, ) *POSTNotifications201ResponseDataAttributes`

NewPOSTNotifications201ResponseDataAttributes instantiates a new POSTNotifications201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTNotifications201ResponseDataAttributesWithDefaults

`func NewPOSTNotifications201ResponseDataAttributesWithDefaults() *POSTNotifications201ResponseDataAttributes`

NewPOSTNotifications201ResponseDataAttributesWithDefaults instantiates a new POSTNotifications201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTNotifications201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTNotifications201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTNotifications201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTNotifications201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTNotifications201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFlash

`func (o *POSTNotifications201ResponseDataAttributes) GetFlash() interface{}`

GetFlash returns the Flash field if non-nil, zero value otherwise.

### GetFlashOk

`func (o *POSTNotifications201ResponseDataAttributes) GetFlashOk() (*interface{}, bool)`

GetFlashOk returns a tuple with the Flash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlash

`func (o *POSTNotifications201ResponseDataAttributes) SetFlash(v interface{})`

SetFlash sets Flash field to given value.

### HasFlash

`func (o *POSTNotifications201ResponseDataAttributes) HasFlash() bool`

HasFlash returns a boolean if a field has been set.

### SetFlashNil

`func (o *POSTNotifications201ResponseDataAttributes) SetFlashNil(b bool)`

 SetFlashNil sets the value for Flash to be an explicit nil

### UnsetFlash
`func (o *POSTNotifications201ResponseDataAttributes) UnsetFlash()`

UnsetFlash ensures that no value is present for Flash, not even an explicit nil
### GetBody

`func (o *POSTNotifications201ResponseDataAttributes) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *POSTNotifications201ResponseDataAttributes) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *POSTNotifications201ResponseDataAttributes) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *POSTNotifications201ResponseDataAttributes) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *POSTNotifications201ResponseDataAttributes) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *POSTNotifications201ResponseDataAttributes) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetReference

`func (o *POSTNotifications201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTNotifications201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTNotifications201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTNotifications201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTNotifications201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTNotifications201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTNotifications201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTNotifications201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTNotifications201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTNotifications201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTNotifications201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTNotifications201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTNotifications201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTNotifications201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTNotifications201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTNotifications201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTNotifications201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTNotifications201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


