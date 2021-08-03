/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// UsageIoTHour IoT usage for a given organization for a given hour.
type UsageIoTHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// The total number of IoT devices during a given hour.
	IotDeviceCount *int64 `json:"iot_device_count,omitempty"`
}

// NewUsageIoTHour instantiates a new UsageIoTHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageIoTHour() *UsageIoTHour {
	this := UsageIoTHour{}
	return &this
}

// NewUsageIoTHourWithDefaults instantiates a new UsageIoTHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageIoTHourWithDefaults() *UsageIoTHour {
	this := UsageIoTHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageIoTHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIoTHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageIoTHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageIoTHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetIotDeviceCount returns the IotDeviceCount field value if set, zero value otherwise.
func (o *UsageIoTHour) GetIotDeviceCount() int64 {
	if o == nil || o.IotDeviceCount == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceCount
}

// GetIotDeviceCountOk returns a tuple with the IotDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIoTHour) GetIotDeviceCountOk() (*int64, bool) {
	if o == nil || o.IotDeviceCount == nil {
		return nil, false
	}
	return o.IotDeviceCount, true
}

// HasIotDeviceCount returns a boolean if a field has been set.
func (o *UsageIoTHour) HasIotDeviceCount() bool {
	if o != nil && o.IotDeviceCount != nil {
		return true
	}

	return false
}

// SetIotDeviceCount gets a reference to the given int64 and assigns it to the IotDeviceCount field.
func (o *UsageIoTHour) SetIotDeviceCount(v int64) {
	o.IotDeviceCount = &v
}

func (o UsageIoTHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.IotDeviceCount != nil {
		toSerialize["iot_device_count"] = o.IotDeviceCount
	}
	return json.Marshal(toSerialize)
}

type NullableUsageIoTHour struct {
	value *UsageIoTHour
	isSet bool
}

func (v NullableUsageIoTHour) Get() *UsageIoTHour {
	return v.value
}

func (v *NullableUsageIoTHour) Set(val *UsageIoTHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageIoTHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageIoTHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageIoTHour(val *UsageIoTHour) *NullableUsageIoTHour {
	return &NullableUsageIoTHour{value: val, isSet: true}
}

func (v NullableUsageIoTHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageIoTHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}