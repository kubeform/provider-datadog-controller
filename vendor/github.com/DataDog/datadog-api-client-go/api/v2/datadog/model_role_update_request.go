/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// RoleUpdateRequest Update a role.
type RoleUpdateRequest struct {
	Data RoleUpdateData `json:"data"`
}

// NewRoleUpdateRequest instantiates a new RoleUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleUpdateRequest(data RoleUpdateData) *RoleUpdateRequest {
	this := RoleUpdateRequest{}
	this.Data = data
	return &this
}

// NewRoleUpdateRequestWithDefaults instantiates a new RoleUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleUpdateRequestWithDefaults() *RoleUpdateRequest {
	this := RoleUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *RoleUpdateRequest) GetData() RoleUpdateData {
	if o == nil {
		var ret RoleUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RoleUpdateRequest) GetDataOk() (*RoleUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RoleUpdateRequest) SetData(v RoleUpdateData) {
	o.Data = v
}

func (o RoleUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *RoleUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Data *RoleUpdateData `json:"data"`
	}{}
	all := struct {
		Data RoleUpdateData `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Data = all.Data
	return nil
}

type NullableRoleUpdateRequest struct {
	value *RoleUpdateRequest
	isSet bool
}

func (v NullableRoleUpdateRequest) Get() *RoleUpdateRequest {
	return v.value
}

func (v *NullableRoleUpdateRequest) Set(val *RoleUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleUpdateRequest(val *RoleUpdateRequest) *NullableRoleUpdateRequest {
	return &NullableRoleUpdateRequest{value: val, isSet: true}
}

func (v NullableRoleUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
