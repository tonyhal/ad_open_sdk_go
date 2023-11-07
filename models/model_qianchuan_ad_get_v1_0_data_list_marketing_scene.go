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

// QianchuanAdGetV10DataListMarketingScene
type QianchuanAdGetV10DataListMarketingScene string

// List of qianchuan_ad_get_v1.0_data_list_marketing_scene
const (
	ALL_QianchuanAdGetV10DataListMarketingScene           QianchuanAdGetV10DataListMarketingScene = "ALL"
	FEED_QianchuanAdGetV10DataListMarketingScene          QianchuanAdGetV10DataListMarketingScene = "FEED"
	SEARCH_QianchuanAdGetV10DataListMarketingScene        QianchuanAdGetV10DataListMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanAdGetV10DataListMarketingScene QianchuanAdGetV10DataListMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanAdGetV10DataListMarketingScene enum
var AllowedQianchuanAdGetV10DataListMarketingSceneEnumValues = []QianchuanAdGetV10DataListMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanAdGetV10DataListMarketingSceneFromValue returns a pointer to a valid QianchuanAdGetV10DataListMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListMarketingSceneFromValue(v string) (*QianchuanAdGetV10DataListMarketingScene, error) {
	ev := QianchuanAdGetV10DataListMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListMarketingScene: valid values are %v", v, AllowedQianchuanAdGetV10DataListMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_marketing_scene value
func (v QianchuanAdGetV10DataListMarketingScene) Ptr() *QianchuanAdGetV10DataListMarketingScene {
	return &v
}
