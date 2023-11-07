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

// QianchuanProductAvailableGetV10FilterMarketingScene
type QianchuanProductAvailableGetV10FilterMarketingScene string

// List of qianchuan_product_available_get_v1.0_filter_marketing_scene
const (
	FEED_QianchuanProductAvailableGetV10FilterMarketingScene          QianchuanProductAvailableGetV10FilterMarketingScene = "FEED"
	SEARCH_QianchuanProductAvailableGetV10FilterMarketingScene        QianchuanProductAvailableGetV10FilterMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanProductAvailableGetV10FilterMarketingScene QianchuanProductAvailableGetV10FilterMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanProductAvailableGetV10FilterMarketingScene enum
var AllowedQianchuanProductAvailableGetV10FilterMarketingSceneEnumValues = []QianchuanProductAvailableGetV10FilterMarketingScene{
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanProductAvailableGetV10FilterMarketingSceneFromValue returns a pointer to a valid QianchuanProductAvailableGetV10FilterMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAvailableGetV10FilterMarketingSceneFromValue(v string) (*QianchuanProductAvailableGetV10FilterMarketingScene, error) {
	ev := QianchuanProductAvailableGetV10FilterMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAvailableGetV10FilterMarketingScene: valid values are %v", v, AllowedQianchuanProductAvailableGetV10FilterMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAvailableGetV10FilterMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAvailableGetV10FilterMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_available_get_v1.0_filter_marketing_scene value
func (v QianchuanProductAvailableGetV10FilterMarketingScene) Ptr() *QianchuanProductAvailableGetV10FilterMarketingScene {
	return &v
}
