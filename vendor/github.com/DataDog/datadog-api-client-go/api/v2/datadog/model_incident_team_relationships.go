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

// IncidentTeamRelationships The incident team's relationships.
type IncidentTeamRelationships struct {
	CreatedBy      *RelationshipToUser `json:"created_by,omitempty"`
	LastModifiedBy *RelationshipToUser `json:"last_modified_by,omitempty"`
}

// NewIncidentTeamRelationships instantiates a new IncidentTeamRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTeamRelationships() *IncidentTeamRelationships {
	this := IncidentTeamRelationships{}
	return &this
}

// NewIncidentTeamRelationshipsWithDefaults instantiates a new IncidentTeamRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTeamRelationshipsWithDefaults() *IncidentTeamRelationships {
	this := IncidentTeamRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *IncidentTeamRelationships) GetCreatedBy() RelationshipToUser {
	if o == nil || o.CreatedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamRelationships) GetCreatedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IncidentTeamRelationships) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given RelationshipToUser and assigns it to the CreatedBy field.
func (o *IncidentTeamRelationships) SetCreatedBy(v RelationshipToUser) {
	o.CreatedBy = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *IncidentTeamRelationships) GetLastModifiedBy() RelationshipToUser {
	if o == nil || o.LastModifiedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamRelationships) GetLastModifiedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *IncidentTeamRelationships) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given RelationshipToUser and assigns it to the LastModifiedBy field.
func (o *IncidentTeamRelationships) SetLastModifiedBy(v RelationshipToUser) {
	o.LastModifiedBy = &v
}

func (o IncidentTeamRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.LastModifiedBy != nil {
		toSerialize["last_modified_by"] = o.LastModifiedBy
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentTeamRelationships struct {
	value *IncidentTeamRelationships
	isSet bool
}

func (v NullableIncidentTeamRelationships) Get() *IncidentTeamRelationships {
	return v.value
}

func (v *NullableIncidentTeamRelationships) Set(val *IncidentTeamRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTeamRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTeamRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTeamRelationships(val *IncidentTeamRelationships) *NullableIncidentTeamRelationships {
	return &NullableIncidentTeamRelationships{value: val, isSet: true}
}

func (v NullableIncidentTeamRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTeamRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
