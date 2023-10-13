/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsCommentGetV30DataCommentListCommentPermission
type ToolsCommentGetV30DataCommentListCommentPermission string

// List of tools_comment_get_v3.0_data_comment_list_comment_permission
const (
	READ_ToolsCommentGetV30DataCommentListCommentPermission  ToolsCommentGetV30DataCommentListCommentPermission = "READ"
	WRITE_ToolsCommentGetV30DataCommentListCommentPermission ToolsCommentGetV30DataCommentListCommentPermission = "WRITE"
)

// All allowed values of ToolsCommentGetV30DataCommentListCommentPermission enum
var AllowedToolsCommentGetV30DataCommentListCommentPermissionEnumValues = []ToolsCommentGetV30DataCommentListCommentPermission{
	"READ",
	"WRITE",
}

// NewToolsCommentGetV30DataCommentListCommentPermissionFromValue returns a pointer to a valid ToolsCommentGetV30DataCommentListCommentPermission
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30DataCommentListCommentPermissionFromValue(v string) (*ToolsCommentGetV30DataCommentListCommentPermission, error) {
	ev := ToolsCommentGetV30DataCommentListCommentPermission(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30DataCommentListCommentPermission: valid values are %v", v, AllowedToolsCommentGetV30DataCommentListCommentPermissionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30DataCommentListCommentPermission) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30DataCommentListCommentPermissionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_data_comment_list_comment_permission value
func (v ToolsCommentGetV30DataCommentListCommentPermission) Ptr() *ToolsCommentGetV30DataCommentListCommentPermission {
	return &v
}