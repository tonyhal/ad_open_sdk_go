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

// DouplusOrderListV30DataOrderListOrderLiveScene
type DouplusOrderListV30DataOrderListOrderLiveScene string

// List of douplus_order_list_v3.0_data_order_list_order_live_scene
const (
	Enum_混合投放_DouplusOrderListV30DataOrderListOrderLiveScene    DouplusOrderListV30DataOrderListOrderLiveScene = "混合投放"
	Enum_直推直播间_DouplusOrderListV30DataOrderListOrderLiveScene   DouplusOrderListV30DataOrderListOrderLiveScene = "直推直播间"
	Enum_视频导流直播间_DouplusOrderListV30DataOrderListOrderLiveScene DouplusOrderListV30DataOrderListOrderLiveScene = "视频导流直播间"
)

// All allowed values of DouplusOrderListV30DataOrderListOrderLiveScene enum
var AllowedDouplusOrderListV30DataOrderListOrderLiveSceneEnumValues = []DouplusOrderListV30DataOrderListOrderLiveScene{
	"混合投放",
	"直推直播间",
	"视频导流直播间",
}

// NewDouplusOrderListV30DataOrderListOrderLiveSceneFromValue returns a pointer to a valid DouplusOrderListV30DataOrderListOrderLiveScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30DataOrderListOrderLiveSceneFromValue(v string) (*DouplusOrderListV30DataOrderListOrderLiveScene, error) {
	ev := DouplusOrderListV30DataOrderListOrderLiveScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30DataOrderListOrderLiveScene: valid values are %v", v, AllowedDouplusOrderListV30DataOrderListOrderLiveSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30DataOrderListOrderLiveScene) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30DataOrderListOrderLiveSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_data_order_list_order_live_scene value
func (v DouplusOrderListV30DataOrderListOrderLiveScene) Ptr() *DouplusOrderListV30DataOrderListOrderLiveScene {
	return &v
}
