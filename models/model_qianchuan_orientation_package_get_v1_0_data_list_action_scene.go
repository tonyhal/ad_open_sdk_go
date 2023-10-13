/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanOrientationPackageGetV10DataListActionScene
type QianchuanOrientationPackageGetV10DataListActionScene string

// List of qianchuan_orientation_package_get_v1.0_data_list_action_scene
const (
	APP_QianchuanOrientationPackageGetV10DataListActionScene        QianchuanOrientationPackageGetV10DataListActionScene = "APP"
	E_COMMERCE_QianchuanOrientationPackageGetV10DataListActionScene QianchuanOrientationPackageGetV10DataListActionScene = "E_COMMERCE"
	NEWS_QianchuanOrientationPackageGetV10DataListActionScene       QianchuanOrientationPackageGetV10DataListActionScene = "NEWS"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListActionScene enum
var AllowedQianchuanOrientationPackageGetV10DataListActionSceneEnumValues = []QianchuanOrientationPackageGetV10DataListActionScene{
	"APP",
	"E_COMMERCE",
	"NEWS",
}

// NewQianchuanOrientationPackageGetV10DataListActionSceneFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListActionSceneFromValue(v string) (*QianchuanOrientationPackageGetV10DataListActionScene, error) {
	ev := QianchuanOrientationPackageGetV10DataListActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListActionScene: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListActionScene) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_action_scene value
func (v QianchuanOrientationPackageGetV10DataListActionScene) Ptr() *QianchuanOrientationPackageGetV10DataListActionScene {
	return &v
}
