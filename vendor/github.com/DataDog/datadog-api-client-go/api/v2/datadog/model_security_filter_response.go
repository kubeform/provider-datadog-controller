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

// SecurityFilterResponse Response object which includes a single security filter.
type SecurityFilterResponse struct {
	Data *SecurityFilter     `json:"data,omitempty"`
	Meta *SecurityFilterMeta `json:"meta,omitempty"`
}

// NewSecurityFilterResponse instantiates a new SecurityFilterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityFilterResponse() *SecurityFilterResponse {
	this := SecurityFilterResponse{}
	return &this
}

// NewSecurityFilterResponseWithDefaults instantiates a new SecurityFilterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityFilterResponseWithDefaults() *SecurityFilterResponse {
	this := SecurityFilterResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SecurityFilterResponse) GetData() SecurityFilter {
	if o == nil || o.Data == nil {
		var ret SecurityFilter
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFilterResponse) GetDataOk() (*SecurityFilter, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SecurityFilterResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SecurityFilter and assigns it to the Data field.
func (o *SecurityFilterResponse) SetData(v SecurityFilter) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SecurityFilterResponse) GetMeta() SecurityFilterMeta {
	if o == nil || o.Meta == nil {
		var ret SecurityFilterMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFilterResponse) GetMetaOk() (*SecurityFilterMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SecurityFilterResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SecurityFilterMeta and assigns it to the Meta field.
func (o *SecurityFilterResponse) SetMeta(v SecurityFilterMeta) {
	o.Meta = &v
}

func (o SecurityFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityFilterResponse struct {
	value *SecurityFilterResponse
	isSet bool
}

func (v NullableSecurityFilterResponse) Get() *SecurityFilterResponse {
	return v.value
}

func (v *NullableSecurityFilterResponse) Set(val *SecurityFilterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityFilterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityFilterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityFilterResponse(val *SecurityFilterResponse) *NullableSecurityFilterResponse {
	return &NullableSecurityFilterResponse{value: val, isSet: true}
}

func (v NullableSecurityFilterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityFilterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
