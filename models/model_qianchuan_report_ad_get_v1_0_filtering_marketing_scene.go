/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdGetV10FilteringMarketingScene
type QianchuanReportAdGetV10FilteringMarketingScene string

// List of qianchuan_report_ad_get_v1.0_filtering_marketing_scene
const (
	ALL_QianchuanReportAdGetV10FilteringMarketingScene           QianchuanReportAdGetV10FilteringMarketingScene = "ALL"
	FEED_QianchuanReportAdGetV10FilteringMarketingScene          QianchuanReportAdGetV10FilteringMarketingScene = "FEED"
	SEARCH_QianchuanReportAdGetV10FilteringMarketingScene        QianchuanReportAdGetV10FilteringMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanReportAdGetV10FilteringMarketingScene QianchuanReportAdGetV10FilteringMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanReportAdGetV10FilteringMarketingScene enum
var AllowedQianchuanReportAdGetV10FilteringMarketingSceneEnumValues = []QianchuanReportAdGetV10FilteringMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanReportAdGetV10FilteringMarketingSceneFromValue returns a pointer to a valid QianchuanReportAdGetV10FilteringMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdGetV10FilteringMarketingSceneFromValue(v string) (*QianchuanReportAdGetV10FilteringMarketingScene, error) {
	ev := QianchuanReportAdGetV10FilteringMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdGetV10FilteringMarketingScene: valid values are %v", v, AllowedQianchuanReportAdGetV10FilteringMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdGetV10FilteringMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdGetV10FilteringMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_get_v1.0_filtering_marketing_scene value
func (v QianchuanReportAdGetV10FilteringMarketingScene) Ptr() *QianchuanReportAdGetV10FilteringMarketingScene {
	return &v
}
