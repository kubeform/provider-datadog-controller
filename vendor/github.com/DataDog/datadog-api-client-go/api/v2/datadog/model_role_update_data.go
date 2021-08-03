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

// RoleUpdateData Data related to the update of a role.
type RoleUpdateData struct {
	Attributes RoleUpdateAttributes `json:"attributes"`
	// ID of the role.
	Id   string    `json:"id"`
	Type RolesType `json:"type"`
}

// NewRoleUpdateData instantiates a new RoleUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleUpdateData(attributes RoleUpdateAttributes, id string, type_ RolesType) *RoleUpdateData {
	this := RoleUpdateData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = type_
	return &this
}

// NewRoleUpdateDataWithDefaults instantiates a new RoleUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleUpdateDataWithDefaults() *RoleUpdateData {
	this := RoleUpdateData{}
	var type_ RolesType = ROLESTYPE_ROLES
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *RoleUpdateData) GetAttributes() RoleUpdateAttributes {
	if o == nil {
		var ret RoleUpdateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RoleUpdateData) GetAttributesOk() (*RoleUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *RoleUpdateData) SetAttributes(v RoleUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value
func (o *RoleUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RoleUpdateData) GetType() RolesType {
	if o == nil {
		var ret RolesType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoleUpdateData) GetTypeOk() (*RolesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoleUpdateData) SetType(v RolesType) {
	o.Type = v
}

func (o RoleUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *RoleUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Attributes *RoleUpdateAttributes `json:"attributes"`
		Id         *string               `json:"id"`
		Type       *RolesType            `json:"type"`
	}{}
	all := struct {
		Attributes RoleUpdateAttributes `json:"attributes"`
		Id         string               `json:"id"`
		Type       RolesType            `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
	}
	if required.Id == nil {
		return fmt.Errorf("Required field id missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	o.Type = all.Type
	return nil
}

type NullableRoleUpdateData struct {
	value *RoleUpdateData
	isSet bool
}

func (v NullableRoleUpdateData) Get() *RoleUpdateData {
	return v.value
}

func (v *NullableRoleUpdateData) Set(val *RoleUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleUpdateData(val *RoleUpdateData) *NullableRoleUpdateData {
	return &NullableRoleUpdateData{value: val, isSet: true}
}

func (v NullableRoleUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}