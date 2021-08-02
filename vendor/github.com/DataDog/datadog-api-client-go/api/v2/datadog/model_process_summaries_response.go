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

// ProcessSummariesResponse List of process summaries.
type ProcessSummariesResponse struct {
	// Array of process summary objects.
	Data *[]ProcessSummary     `json:"data,omitempty"`
	Meta *ProcessSummariesMeta `json:"meta,omitempty"`
}

// NewProcessSummariesResponse instantiates a new ProcessSummariesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessSummariesResponse() *ProcessSummariesResponse {
	this := ProcessSummariesResponse{}
	return &this
}

// NewProcessSummariesResponseWithDefaults instantiates a new ProcessSummariesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessSummariesResponseWithDefaults() *ProcessSummariesResponse {
	this := ProcessSummariesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProcessSummariesResponse) GetData() []ProcessSummary {
	if o == nil || o.Data == nil {
		var ret []ProcessSummary
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessSummariesResponse) GetDataOk() (*[]ProcessSummary, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProcessSummariesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ProcessSummary and assigns it to the Data field.
func (o *ProcessSummariesResponse) SetData(v []ProcessSummary) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProcessSummariesResponse) GetMeta() ProcessSummariesMeta {
	if o == nil || o.Meta == nil {
		var ret ProcessSummariesMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessSummariesResponse) GetMetaOk() (*ProcessSummariesMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProcessSummariesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ProcessSummariesMeta and assigns it to the Meta field.
func (o *ProcessSummariesResponse) SetMeta(v ProcessSummariesMeta) {
	o.Meta = &v
}

func (o ProcessSummariesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableProcessSummariesResponse struct {
	value *ProcessSummariesResponse
	isSet bool
}

func (v NullableProcessSummariesResponse) Get() *ProcessSummariesResponse {
	return v.value
}

func (v *NullableProcessSummariesResponse) Set(val *ProcessSummariesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessSummariesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessSummariesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessSummariesResponse(val *ProcessSummariesResponse) *NullableProcessSummariesResponse {
	return &NullableProcessSummariesResponse{value: val, isSet: true}
}

func (v NullableProcessSummariesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessSummariesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
