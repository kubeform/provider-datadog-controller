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

// Widget Information about widget.  **Note**: The `layout` property is required for widgets in dashboards with `free` `layout_type`.       For the **new dashboard layout**, the `layout` property depends on the `reflow_type` of the dashboard.       - If `reflow_type` is `fixed`, `layout` is required.       - If `reflow_type` is `auto`, `layout` should not be set.
type Widget struct {
	Definition WidgetDefinition `json:"definition"`
	// ID of the widget.
	Id     *int64        `json:"id,omitempty"`
	Layout *WidgetLayout `json:"layout,omitempty"`
}

// NewWidget instantiates a new Widget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidget(definition WidgetDefinition) *Widget {
	this := Widget{}
	this.Definition = definition
	return &this
}

// NewWidgetWithDefaults instantiates a new Widget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetWithDefaults() *Widget {
	this := Widget{}
	return &this
}

// GetDefinition returns the Definition field value
func (o *Widget) GetDefinition() WidgetDefinition {
	if o == nil {
		var ret WidgetDefinition
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *Widget) GetDefinitionOk() (*WidgetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *Widget) SetDefinition(v WidgetDefinition) {
	o.Definition = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Widget) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Widget) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Widget) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Widget) SetId(v int64) {
	o.Id = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *Widget) GetLayout() WidgetLayout {
	if o == nil || o.Layout == nil {
		var ret WidgetLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Widget) GetLayoutOk() (*WidgetLayout, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *Widget) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given WidgetLayout and assigns it to the Layout field.
func (o *Widget) SetLayout(v WidgetLayout) {
	o.Layout = &v
}

func (o Widget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["definition"] = o.Definition
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	return json.Marshal(toSerialize)
}

func (o *Widget) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Definition *WidgetDefinition `json:"definition"`
	}{}
	all := struct {
		Definition WidgetDefinition `json:"definition"`
		Id         *int64           `json:"id,omitempty"`
		Layout     *WidgetLayout    `json:"layout,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Definition == nil {
		return fmt.Errorf("Required field definition missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Definition = all.Definition
	o.Id = all.Id
	o.Layout = all.Layout
	return nil
}

type NullableWidget struct {
	value *Widget
	isSet bool
}

func (v NullableWidget) Get() *Widget {
	return v.value
}

func (v *NullableWidget) Set(val *Widget) {
	v.value = val
	v.isSet = true
}

func (v NullableWidget) IsSet() bool {
	return v.isSet
}

func (v *NullableWidget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidget(val *Widget) *NullableWidget {
	return &NullableWidget{value: val, isSet: true}
}

func (v NullableWidget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
