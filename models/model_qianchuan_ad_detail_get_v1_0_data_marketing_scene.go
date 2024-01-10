/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataMarketingScene
type QianchuanAdDetailGetV10DataMarketingScene string

// List of qianchuan_ad_detail_get_v1.0_data_marketing_scene
const (
	ALL_QianchuanAdDetailGetV10DataMarketingScene           QianchuanAdDetailGetV10DataMarketingScene = "ALL"
	FEED_QianchuanAdDetailGetV10DataMarketingScene          QianchuanAdDetailGetV10DataMarketingScene = "FEED"
	SEARCH_QianchuanAdDetailGetV10DataMarketingScene        QianchuanAdDetailGetV10DataMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanAdDetailGetV10DataMarketingScene QianchuanAdDetailGetV10DataMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanAdDetailGetV10DataMarketingScene enum
var AllowedQianchuanAdDetailGetV10DataMarketingSceneEnumValues = []QianchuanAdDetailGetV10DataMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanAdDetailGetV10DataMarketingSceneFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataMarketingSceneFromValue(v string) (*QianchuanAdDetailGetV10DataMarketingScene, error) {
	ev := QianchuanAdDetailGetV10DataMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataMarketingScene: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_marketing_scene value
func (v QianchuanAdDetailGetV10DataMarketingScene) Ptr() *QianchuanAdDetailGetV10DataMarketingScene {
	return &v
}
