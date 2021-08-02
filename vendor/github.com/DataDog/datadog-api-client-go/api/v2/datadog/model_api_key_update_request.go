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

// APIKeyUpdateRequest Request used to update an API key.
type APIKeyUpdateRequest struct {
	Data APIKeyUpdateData `json:"data"`
}

// NewAPIKeyUpdateRequest instantiates a new APIKeyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIKeyUpdateRequest(data APIKeyUpdateData) *APIKeyUpdateRequest {
	this := APIKeyUpdateRequest{}
	this.Data = data
	return &this
}

// NewAPIKeyUpdateRequestWithDefaults instantiates a new APIKeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIKeyUpdateRequestWithDefaults() *APIKeyUpdateRequest {
	this := APIKeyUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *APIKeyUpdateRequest) GetData() APIKeyUpdateData {
	if o == nil {
		var ret APIKeyUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *APIKeyUpdateRequest) GetDataOk() (*APIKeyUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *APIKeyUpdateRequest) SetData(v APIKeyUpdateData) {
	o.Data = v
}

func (o APIKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *APIKeyUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Data *APIKeyUpdateData `json:"data"`
	}{}
	all := struct {
		Data APIKeyUpdateData `json:"data"`
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

type NullableAPIKeyUpdateRequest struct {
	value *APIKeyUpdateRequest
	isSet bool
}

func (v NullableAPIKeyUpdateRequest) Get() *APIKeyUpdateRequest {
	return v.value
}

func (v *NullableAPIKeyUpdateRequest) Set(val *APIKeyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIKeyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIKeyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIKeyUpdateRequest(val *APIKeyUpdateRequest) *NullableAPIKeyUpdateRequest {
	return &NullableAPIKeyUpdateRequest{value: val, isSet: true}
}

func (v NullableAPIKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIKeyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
