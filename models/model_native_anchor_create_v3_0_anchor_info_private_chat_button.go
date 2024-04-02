/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorCreateV30AnchorInfoPrivateChatButton
type NativeAnchorCreateV30AnchorInfoPrivateChatButton string

// List of native_anchor_create_v3.0_anchor_info_private_chat_button
const (
	PRIVATE_MESSAGE_NativeAnchorCreateV30AnchorInfoPrivateChatButton  NativeAnchorCreateV30AnchorInfoPrivateChatButton = "PRIVATE_MESSAGE"
	CONSULT_NOW_NativeAnchorCreateV30AnchorInfoPrivateChatButton      NativeAnchorCreateV30AnchorInfoPrivateChatButton = "CONSULT_NOW"
	CONSULT_ADVISOR_NativeAnchorCreateV30AnchorInfoPrivateChatButton  NativeAnchorCreateV30AnchorInfoPrivateChatButton = "CONSULT_ADVISOR"
	CONSULT_DESIGNER_NativeAnchorCreateV30AnchorInfoPrivateChatButton NativeAnchorCreateV30AnchorInfoPrivateChatButton = "CONSULT_DESIGNER"
	ASK_TEACHER_NativeAnchorCreateV30AnchorInfoPrivateChatButton      NativeAnchorCreateV30AnchorInfoPrivateChatButton = "ASK_TEACHER"
)

// All allowed values of NativeAnchorCreateV30AnchorInfoPrivateChatButton enum
var AllowedNativeAnchorCreateV30AnchorInfoPrivateChatButtonEnumValues = []NativeAnchorCreateV30AnchorInfoPrivateChatButton{
	"PRIVATE_MESSAGE",
	"CONSULT_NOW",
	"CONSULT_ADVISOR",
	"CONSULT_DESIGNER",
	"ASK_TEACHER",
}

// NewNativeAnchorCreateV30AnchorInfoPrivateChatButtonFromValue returns a pointer to a valid NativeAnchorCreateV30AnchorInfoPrivateChatButton
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorCreateV30AnchorInfoPrivateChatButtonFromValue(v string) (*NativeAnchorCreateV30AnchorInfoPrivateChatButton, error) {
	ev := NativeAnchorCreateV30AnchorInfoPrivateChatButton(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorCreateV30AnchorInfoPrivateChatButton: valid values are %v", v, AllowedNativeAnchorCreateV30AnchorInfoPrivateChatButtonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorCreateV30AnchorInfoPrivateChatButton) IsValid() bool {
	for _, existing := range AllowedNativeAnchorCreateV30AnchorInfoPrivateChatButtonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_create_v3.0_anchor_info_private_chat_button value
func (v NativeAnchorCreateV30AnchorInfoPrivateChatButton) Ptr() *NativeAnchorCreateV30AnchorInfoPrivateChatButton {
	return &v
}
