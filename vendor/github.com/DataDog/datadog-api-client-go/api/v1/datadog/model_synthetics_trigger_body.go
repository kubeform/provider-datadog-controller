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

// SyntheticsTriggerBody Object describing the synthetics tests to trigger.
type SyntheticsTriggerBody struct {
	// Individual synthetics test.
	Tests []SyntheticsTriggerTest `json:"tests"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSyntheticsTriggerBody instantiates a new SyntheticsTriggerBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTriggerBody(tests []SyntheticsTriggerTest) *SyntheticsTriggerBody {
	this := SyntheticsTriggerBody{}
	this.Tests = tests
	return &this
}

// NewSyntheticsTriggerBodyWithDefaults instantiates a new SyntheticsTriggerBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTriggerBodyWithDefaults() *SyntheticsTriggerBody {
	this := SyntheticsTriggerBody{}
	return &this
}

// GetTests returns the Tests field value
func (o *SyntheticsTriggerBody) GetTests() []SyntheticsTriggerTest {
	if o == nil {
		var ret []SyntheticsTriggerTest
		return ret
	}

	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTriggerBody) GetTestsOk() (*[]SyntheticsTriggerTest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tests, true
}

// SetTests sets field value
func (o *SyntheticsTriggerBody) SetTests(v []SyntheticsTriggerTest) {
	o.Tests = v
}

func (o SyntheticsTriggerBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["tests"] = o.Tests
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsTriggerBody) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Tests *[]SyntheticsTriggerTest `json:"tests"`
	}{}
	all := struct {
		Tests []SyntheticsTriggerTest `json:"tests"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Tests == nil {
		return fmt.Errorf("Required field tests missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Tests = all.Tests
	return nil
}

type NullableSyntheticsTriggerBody struct {
	value *SyntheticsTriggerBody
	isSet bool
}

func (v NullableSyntheticsTriggerBody) Get() *SyntheticsTriggerBody {
	return v.value
}

func (v *NullableSyntheticsTriggerBody) Set(val *SyntheticsTriggerBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTriggerBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTriggerBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTriggerBody(val *SyntheticsTriggerBody) *NullableSyntheticsTriggerBody {
	return &NullableSyntheticsTriggerBody{value: val, isSet: true}
}

func (v NullableSyntheticsTriggerBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTriggerBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
