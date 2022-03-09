/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// EventCreateResponse Object containing an event response.
type EventCreateResponse struct {
	AlertType *EventAlertType `json:"alert_type,omitempty"`
	// POSIX timestamp of the event. Must be sent as an integer (that is no quotes). Limited to events no older than 7 days.
	DateHappened *int64 `json:"date_happened,omitempty"`
	// A device name.
	DeviceName *string `json:"device_name,omitempty"`
	// Host name to associate with the event. Any tags associated with the host are also applied to this event.
	Host *string `json:"host,omitempty"`
	// Integer ID of the event.
	Id *int64 `json:"id,omitempty"`
	// Payload of the event.
	Payload  *string        `json:"payload,omitempty"`
	Priority *EventPriority `json:"priority,omitempty"`
	// ID of the parent event. Must be sent as an integer (that is no quotes).
	RelatedEventId *int64 `json:"related_event_id,omitempty"`
	// The type of event being posted. Option examples include nagios, hudson, jenkins, my_apps, chef, puppet, git, bitbucket, etc. A complete list of source attribute values [available here](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value).
	SourceTypeName *string `json:"source_type_name,omitempty"`
	// A status.
	Status *string `json:"status,omitempty"`
	// A list of tags to apply to the event.
	Tags *[]string `json:"tags,omitempty"`
	// The body of the event. Limited to 4000 characters. The text supports markdown. Use `msg_text` with the Datadog Ruby library.
	Text *string `json:"text,omitempty"`
	// The event title. Limited to 100 characters. Use `msg_title` with the Datadog Ruby library.
	Title *string `json:"title,omitempty"`
	// URL of the event.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewEventCreateResponse instantiates a new EventCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCreateResponse() *EventCreateResponse {
	this := EventCreateResponse{}
	return &this
}

// NewEventCreateResponseWithDefaults instantiates a new EventCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCreateResponseWithDefaults() *EventCreateResponse {
	this := EventCreateResponse{}
	return &this
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *EventCreateResponse) GetAlertType() EventAlertType {
	if o == nil || o.AlertType == nil {
		var ret EventAlertType
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetAlertTypeOk() (*EventAlertType, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *EventCreateResponse) HasAlertType() bool {
	if o != nil && o.AlertType != nil {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given EventAlertType and assigns it to the AlertType field.
func (o *EventCreateResponse) SetAlertType(v EventAlertType) {
	o.AlertType = &v
}

// GetDateHappened returns the DateHappened field value if set, zero value otherwise.
func (o *EventCreateResponse) GetDateHappened() int64 {
	if o == nil || o.DateHappened == nil {
		var ret int64
		return ret
	}
	return *o.DateHappened
}

// GetDateHappenedOk returns a tuple with the DateHappened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetDateHappenedOk() (*int64, bool) {
	if o == nil || o.DateHappened == nil {
		return nil, false
	}
	return o.DateHappened, true
}

// HasDateHappened returns a boolean if a field has been set.
func (o *EventCreateResponse) HasDateHappened() bool {
	if o != nil && o.DateHappened != nil {
		return true
	}

	return false
}

// SetDateHappened gets a reference to the given int64 and assigns it to the DateHappened field.
func (o *EventCreateResponse) SetDateHappened(v int64) {
	o.DateHappened = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *EventCreateResponse) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *EventCreateResponse) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *EventCreateResponse) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *EventCreateResponse) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *EventCreateResponse) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *EventCreateResponse) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventCreateResponse) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventCreateResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *EventCreateResponse) SetId(v int64) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *EventCreateResponse) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *EventCreateResponse) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *EventCreateResponse) SetPayload(v string) {
	o.Payload = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *EventCreateResponse) GetPriority() EventPriority {
	if o == nil || o.Priority == nil {
		var ret EventPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetPriorityOk() (*EventPriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *EventCreateResponse) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given EventPriority and assigns it to the Priority field.
func (o *EventCreateResponse) SetPriority(v EventPriority) {
	o.Priority = &v
}

// GetRelatedEventId returns the RelatedEventId field value if set, zero value otherwise.
func (o *EventCreateResponse) GetRelatedEventId() int64 {
	if o == nil || o.RelatedEventId == nil {
		var ret int64
		return ret
	}
	return *o.RelatedEventId
}

// GetRelatedEventIdOk returns a tuple with the RelatedEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetRelatedEventIdOk() (*int64, bool) {
	if o == nil || o.RelatedEventId == nil {
		return nil, false
	}
	return o.RelatedEventId, true
}

// HasRelatedEventId returns a boolean if a field has been set.
func (o *EventCreateResponse) HasRelatedEventId() bool {
	if o != nil && o.RelatedEventId != nil {
		return true
	}

	return false
}

// SetRelatedEventId gets a reference to the given int64 and assigns it to the RelatedEventId field.
func (o *EventCreateResponse) SetRelatedEventId(v int64) {
	o.RelatedEventId = &v
}

// GetSourceTypeName returns the SourceTypeName field value if set, zero value otherwise.
func (o *EventCreateResponse) GetSourceTypeName() string {
	if o == nil || o.SourceTypeName == nil {
		var ret string
		return ret
	}
	return *o.SourceTypeName
}

// GetSourceTypeNameOk returns a tuple with the SourceTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetSourceTypeNameOk() (*string, bool) {
	if o == nil || o.SourceTypeName == nil {
		return nil, false
	}
	return o.SourceTypeName, true
}

// HasSourceTypeName returns a boolean if a field has been set.
func (o *EventCreateResponse) HasSourceTypeName() bool {
	if o != nil && o.SourceTypeName != nil {
		return true
	}

	return false
}

// SetSourceTypeName gets a reference to the given string and assigns it to the SourceTypeName field.
func (o *EventCreateResponse) SetSourceTypeName(v string) {
	o.SourceTypeName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventCreateResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventCreateResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventCreateResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventCreateResponse) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventCreateResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EventCreateResponse) SetTags(v []string) {
	o.Tags = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *EventCreateResponse) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *EventCreateResponse) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *EventCreateResponse) SetText(v string) {
	o.Text = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EventCreateResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EventCreateResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EventCreateResponse) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EventCreateResponse) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponse) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EventCreateResponse) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EventCreateResponse) SetUrl(v string) {
	o.Url = &v
}

func (o EventCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AlertType != nil {
		toSerialize["alert_type"] = o.AlertType
	}
	if o.DateHappened != nil {
		toSerialize["date_happened"] = o.DateHappened
	}
	if o.DeviceName != nil {
		toSerialize["device_name"] = o.DeviceName
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.RelatedEventId != nil {
		toSerialize["related_event_id"] = o.RelatedEventId
	}
	if o.SourceTypeName != nil {
		toSerialize["source_type_name"] = o.SourceTypeName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

func (o *EventCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AlertType      *EventAlertType `json:"alert_type,omitempty"`
		DateHappened   *int64          `json:"date_happened,omitempty"`
		DeviceName     *string         `json:"device_name,omitempty"`
		Host           *string         `json:"host,omitempty"`
		Id             *int64          `json:"id,omitempty"`
		Payload        *string         `json:"payload,omitempty"`
		Priority       *EventPriority  `json:"priority,omitempty"`
		RelatedEventId *int64          `json:"related_event_id,omitempty"`
		SourceTypeName *string         `json:"source_type_name,omitempty"`
		Status         *string         `json:"status,omitempty"`
		Tags           *[]string       `json:"tags,omitempty"`
		Text           *string         `json:"text,omitempty"`
		Title          *string         `json:"title,omitempty"`
		Url            *string         `json:"url,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.AlertType; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Priority; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.AlertType = all.AlertType
	o.DateHappened = all.DateHappened
	o.DeviceName = all.DeviceName
	o.Host = all.Host
	o.Id = all.Id
	o.Payload = all.Payload
	o.Priority = all.Priority
	o.RelatedEventId = all.RelatedEventId
	o.SourceTypeName = all.SourceTypeName
	o.Status = all.Status
	o.Tags = all.Tags
	o.Text = all.Text
	o.Title = all.Title
	o.Url = all.Url
	return nil
}

type NullableEventCreateResponse struct {
	value *EventCreateResponse
	isSet bool
}

func (v NullableEventCreateResponse) Get() *EventCreateResponse {
	return v.value
}

func (v *NullableEventCreateResponse) Set(val *EventCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCreateResponse(val *EventCreateResponse) *NullableEventCreateResponse {
	return &NullableEventCreateResponse{value: val, isSet: true}
}

func (v NullableEventCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
