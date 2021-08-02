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

// WidgetFormula Formula to be used in a widget query.
type WidgetFormula struct {
	// Expression alias.
	Alias *string `json:"alias,omitempty"`
	// String expression built from queries, formulas, and functions.
	Formula string              `json:"formula"`
	Limit   *WidgetFormulaLimit `json:"limit,omitempty"`
}

// NewWidgetFormula instantiates a new WidgetFormula object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetFormula(formula string) *WidgetFormula {
	this := WidgetFormula{}
	this.Formula = formula
	return &this
}

// NewWidgetFormulaWithDefaults instantiates a new WidgetFormula object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetFormulaWithDefaults() *WidgetFormula {
	this := WidgetFormula{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *WidgetFormula) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *WidgetFormula) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *WidgetFormula) SetAlias(v string) {
	o.Alias = &v
}

// GetFormula returns the Formula field value
func (o *WidgetFormula) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value
func (o *WidgetFormula) SetFormula(v string) {
	o.Formula = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *WidgetFormula) GetLimit() WidgetFormulaLimit {
	if o == nil || o.Limit == nil {
		var ret WidgetFormulaLimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetLimitOk() (*WidgetFormulaLimit, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *WidgetFormula) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given WidgetFormulaLimit and assigns it to the Limit field.
func (o *WidgetFormula) SetLimit(v WidgetFormulaLimit) {
	o.Limit = &v
}

func (o WidgetFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if true {
		toSerialize["formula"] = o.Formula
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

func (o *WidgetFormula) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Formula *string `json:"formula"`
	}{}
	all := struct {
		Alias   *string             `json:"alias,omitempty"`
		Formula string              `json:"formula"`
		Limit   *WidgetFormulaLimit `json:"limit,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Formula == nil {
		return fmt.Errorf("Required field formula missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Alias = all.Alias
	o.Formula = all.Formula
	o.Limit = all.Limit
	return nil
}

type NullableWidgetFormula struct {
	value *WidgetFormula
	isSet bool
}

func (v NullableWidgetFormula) Get() *WidgetFormula {
	return v.value
}

func (v *NullableWidgetFormula) Set(val *WidgetFormula) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetFormula) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetFormula) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetFormula(val *WidgetFormula) *NullableWidgetFormula {
	return &NullableWidgetFormula{value: val, isSet: true}
}

func (v NullableWidgetFormula) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetFormula) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
