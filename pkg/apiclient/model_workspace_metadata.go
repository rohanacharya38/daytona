/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WorkspaceMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceMetadata{}

// WorkspaceMetadata struct for WorkspaceMetadata
type WorkspaceMetadata struct {
	GitStatus GitStatus `json:"gitStatus"`
	UpdatedAt string    `json:"updatedAt"`
	Uptime    int32     `json:"uptime"`
}

type _WorkspaceMetadata WorkspaceMetadata

// NewWorkspaceMetadata instantiates a new WorkspaceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceMetadata(gitStatus GitStatus, updatedAt string, uptime int32) *WorkspaceMetadata {
	this := WorkspaceMetadata{}
	this.GitStatus = gitStatus
	this.UpdatedAt = updatedAt
	this.Uptime = uptime
	return &this
}

// NewWorkspaceMetadataWithDefaults instantiates a new WorkspaceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceMetadataWithDefaults() *WorkspaceMetadata {
	this := WorkspaceMetadata{}
	return &this
}

// GetGitStatus returns the GitStatus field value
func (o *WorkspaceMetadata) GetGitStatus() GitStatus {
	if o == nil {
		var ret GitStatus
		return ret
	}

	return o.GitStatus
}

// GetGitStatusOk returns a tuple with the GitStatus field value
// and a boolean to check if the value has been set.
func (o *WorkspaceMetadata) GetGitStatusOk() (*GitStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitStatus, true
}

// SetGitStatus sets field value
func (o *WorkspaceMetadata) SetGitStatus(v GitStatus) {
	o.GitStatus = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WorkspaceMetadata) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkspaceMetadata) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WorkspaceMetadata) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUptime returns the Uptime field value
func (o *WorkspaceMetadata) GetUptime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value
// and a boolean to check if the value has been set.
func (o *WorkspaceMetadata) GetUptimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uptime, true
}

// SetUptime sets field value
func (o *WorkspaceMetadata) SetUptime(v int32) {
	o.Uptime = v
}

func (o WorkspaceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gitStatus"] = o.GitStatus
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["uptime"] = o.Uptime
	return toSerialize, nil
}

func (o *WorkspaceMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gitStatus",
		"updatedAt",
		"uptime",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWorkspaceMetadata := _WorkspaceMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkspaceMetadata)

	if err != nil {
		return err
	}

	*o = WorkspaceMetadata(varWorkspaceMetadata)

	return err
}

type NullableWorkspaceMetadata struct {
	value *WorkspaceMetadata
	isSet bool
}

func (v NullableWorkspaceMetadata) Get() *WorkspaceMetadata {
	return v.value
}

func (v *NullableWorkspaceMetadata) Set(val *WorkspaceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceMetadata(val *WorkspaceMetadata) *NullableWorkspaceMetadata {
	return &NullableWorkspaceMetadata{value: val, isSet: true}
}

func (v NullableWorkspaceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
