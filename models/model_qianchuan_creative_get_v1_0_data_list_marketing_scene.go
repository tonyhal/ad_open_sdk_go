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

// QianchuanCreativeGetV10DataListMarketingScene
type QianchuanCreativeGetV10DataListMarketingScene string

// List of qianchuan_creative_get_v1.0_data_list_marketing_scene
const (
	ALL_QianchuanCreativeGetV10DataListMarketingScene           QianchuanCreativeGetV10DataListMarketingScene = "ALL"
	FEED_QianchuanCreativeGetV10DataListMarketingScene          QianchuanCreativeGetV10DataListMarketingScene = "FEED"
	SEARCH_QianchuanCreativeGetV10DataListMarketingScene        QianchuanCreativeGetV10DataListMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanCreativeGetV10DataListMarketingScene QianchuanCreativeGetV10DataListMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanCreativeGetV10DataListMarketingScene enum
var AllowedQianchuanCreativeGetV10DataListMarketingSceneEnumValues = []QianchuanCreativeGetV10DataListMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanCreativeGetV10DataListMarketingSceneFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListMarketingSceneFromValue(v string) (*QianchuanCreativeGetV10DataListMarketingScene, error) {
	ev := QianchuanCreativeGetV10DataListMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListMarketingScene: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_marketing_scene value
func (v QianchuanCreativeGetV10DataListMarketingScene) Ptr() *QianchuanCreativeGetV10DataListMarketingScene {
	return &v
}
