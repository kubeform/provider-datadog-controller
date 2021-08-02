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

// LogsDateRemapperType Type of logs date remapper.
type LogsDateRemapperType string

// List of LogsDateRemapperType
const (
	LOGSDATEREMAPPERTYPE_DATE_REMAPPER LogsDateRemapperType = "date-remapper"
)

var allowedLogsDateRemapperTypeEnumValues = []LogsDateRemapperType{
	"date-remapper",
}

func (w *LogsDateRemapperType) GetAllowedValues() []LogsDateRemapperType {
	return allowedLogsDateRemapperTypeEnumValues
}

func (v *LogsDateRemapperType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogsDateRemapperType(value)
	for _, existing := range allowedLogsDateRemapperTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogsDateRemapperType", value)
}

// NewLogsDateRemapperTypeFromValue returns a pointer to a valid LogsDateRemapperType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsDateRemapperTypeFromValue(v string) (*LogsDateRemapperType, error) {
	ev := LogsDateRemapperType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsDateRemapperType: valid values are %v", v, allowedLogsDateRemapperTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsDateRemapperType) IsValid() bool {
	for _, existing := range allowedLogsDateRemapperTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsDateRemapperType value
func (v LogsDateRemapperType) Ptr() *LogsDateRemapperType {
	return &v
}

type NullableLogsDateRemapperType struct {
	value *LogsDateRemapperType
	isSet bool
}

func (v NullableLogsDateRemapperType) Get() *LogsDateRemapperType {
	return v.value
}

func (v *NullableLogsDateRemapperType) Set(val *LogsDateRemapperType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsDateRemapperType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsDateRemapperType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsDateRemapperType(val *LogsDateRemapperType) *NullableLogsDateRemapperType {
	return &NullableLogsDateRemapperType{value: val, isSet: true}
}

func (v NullableLogsDateRemapperType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsDateRemapperType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
