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

// UsageNetworkHostsHour Number of active NPM hosts for each hour for a given organization.
type UsageNetworkHostsHour struct {
	// Contains the number of active NPM hosts.
	HostCount *int64 `json:"host_count,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewUsageNetworkHostsHour instantiates a new UsageNetworkHostsHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageNetworkHostsHour() *UsageNetworkHostsHour {
	this := UsageNetworkHostsHour{}
	return &this
}

// NewUsageNetworkHostsHourWithDefaults instantiates a new UsageNetworkHostsHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageNetworkHostsHourWithDefaults() *UsageNetworkHostsHour {
	this := UsageNetworkHostsHour{}
	return &this
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *UsageNetworkHostsHour) GetHostCount() int64 {
	if o == nil || o.HostCount == nil {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageNetworkHostsHour) GetHostCountOk() (*int64, bool) {
	if o == nil || o.HostCount == nil {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *UsageNetworkHostsHour) HasHostCount() bool {
	if o != nil && o.HostCount != nil {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *UsageNetworkHostsHour) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageNetworkHostsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageNetworkHostsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageNetworkHostsHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageNetworkHostsHour) SetHour(v time.Time) {
	o.Hour = &v
}

func (o UsageNetworkHostsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.HostCount != nil {
		toSerialize["host_count"] = o.HostCount
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	return json.Marshal(toSerialize)
}

func (o *UsageNetworkHostsHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		HostCount *int64     `json:"host_count,omitempty"`
		Hour      *time.Time `json:"hour,omitempty"`
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
	o.HostCount = all.HostCount
	o.Hour = all.Hour
	return nil
}

type NullableUsageNetworkHostsHour struct {
	value *UsageNetworkHostsHour
	isSet bool
}

func (v NullableUsageNetworkHostsHour) Get() *UsageNetworkHostsHour {
	return v.value
}

func (v *NullableUsageNetworkHostsHour) Set(val *UsageNetworkHostsHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageNetworkHostsHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageNetworkHostsHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageNetworkHostsHour(val *UsageNetworkHostsHour) *NullableUsageNetworkHostsHour {
	return &NullableUsageNetworkHostsHour{value: val, isSet: true}
}

func (v NullableUsageNetworkHostsHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageNetworkHostsHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
