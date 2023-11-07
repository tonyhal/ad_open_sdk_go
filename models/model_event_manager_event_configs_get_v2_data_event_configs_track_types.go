/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerEventConfigsGetV2DataEventConfigsTrackTypes
type EventManagerEventConfigsGetV2DataEventConfigsTrackTypes string

// List of event_manager_event_configs_get_v2_data_event_configs_track_types
const (
	JSSDK_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes             EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "JSSDK"
	XPATH_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes             EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "XPATH"
	APPLICATION_API_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes   EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "APPLICATION_API"
	EXTERNAL_API_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes      EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "EXTERNAL_API"
	APPLICATION_SDK_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes   EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "APPLICATION_SDK"
	MINI_PROGRAME_SDK_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "MINI_PROGRAME_SDK"
	MINI_PROGRAME_API_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "MINI_PROGRAME_API"
	OFFLINE_CALLBACK_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes  EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "OFFLINE_CALLBACK"
	OFFLINE_API_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes       EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "OFFLINE_API"
	APPLOG_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes            EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "APPLOG"
	TETRIS_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes            EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "TETRIS"
	TETRIS_FE_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes         EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "TETRIS_FE"
	FLY_FISH_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes          EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "FLY_FISH"
	E_CHAT_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes            EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "E_CHAT"
	QUICK_APP_API_EventManagerEventConfigsGetV2DataEventConfigsTrackTypes     EventManagerEventConfigsGetV2DataEventConfigsTrackTypes = "QUICK_APP_API"
)

// All allowed values of EventManagerEventConfigsGetV2DataEventConfigsTrackTypes enum
var AllowedEventManagerEventConfigsGetV2DataEventConfigsTrackTypesEnumValues = []EventManagerEventConfigsGetV2DataEventConfigsTrackTypes{
	"JSSDK",
	"XPATH",
	"APPLICATION_API",
	"EXTERNAL_API",
	"APPLICATION_SDK",
	"MINI_PROGRAME_SDK",
	"MINI_PROGRAME_API",
	"OFFLINE_CALLBACK",
	"OFFLINE_API",
	"APPLOG",
	"TETRIS",
	"TETRIS_FE",
	"FLY_FISH",
	"E_CHAT",
	"QUICK_APP_API",
}

// NewEventManagerEventConfigsGetV2DataEventConfigsTrackTypesFromValue returns a pointer to a valid EventManagerEventConfigsGetV2DataEventConfigsTrackTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerEventConfigsGetV2DataEventConfigsTrackTypesFromValue(v string) (*EventManagerEventConfigsGetV2DataEventConfigsTrackTypes, error) {
	ev := EventManagerEventConfigsGetV2DataEventConfigsTrackTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerEventConfigsGetV2DataEventConfigsTrackTypes: valid values are %v", v, AllowedEventManagerEventConfigsGetV2DataEventConfigsTrackTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerEventConfigsGetV2DataEventConfigsTrackTypes) IsValid() bool {
	for _, existing := range AllowedEventManagerEventConfigsGetV2DataEventConfigsTrackTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_event_configs_get_v2_data_event_configs_track_types value
func (v EventManagerEventConfigsGetV2DataEventConfigsTrackTypes) Ptr() *EventManagerEventConfigsGetV2DataEventConfigsTrackTypes {
	return &v
}
