/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2ActionScene
type ToolsBidSuggestV2ActionScene string

// List of tools_bid_suggest_v2_action_scene
const (
	SEARCH_ToolsBidSuggestV2ActionScene     ToolsBidSuggestV2ActionScene = "SEARCH"
	NEWS_ToolsBidSuggestV2ActionScene       ToolsBidSuggestV2ActionScene = "NEWS"
	AD_ToolsBidSuggestV2ActionScene         ToolsBidSuggestV2ActionScene = "AD"
	E_COMMERCE_ToolsBidSuggestV2ActionScene ToolsBidSuggestV2ActionScene = "E-COMMERCE"
	APP_ToolsBidSuggestV2ActionScene        ToolsBidSuggestV2ActionScene = "APP"
)

// All allowed values of ToolsBidSuggestV2ActionScene enum
var AllowedToolsBidSuggestV2ActionSceneEnumValues = []ToolsBidSuggestV2ActionScene{
	"SEARCH",
	"NEWS",
	"AD",
	"E-COMMERCE",
	"APP",
}

// NewToolsBidSuggestV2ActionSceneFromValue returns a pointer to a valid ToolsBidSuggestV2ActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2ActionSceneFromValue(v string) (*ToolsBidSuggestV2ActionScene, error) {
	ev := ToolsBidSuggestV2ActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2ActionScene: valid values are %v", v, AllowedToolsBidSuggestV2ActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2ActionScene) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2ActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_action_scene value
func (v ToolsBidSuggestV2ActionScene) Ptr() *ToolsBidSuggestV2ActionScene {
	return &v
}
