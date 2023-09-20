/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsAllowCouponV10CampaignScene
type QianchuanToolsAllowCouponV10CampaignScene string

// List of qianchuan_tools_allow_coupon_v1.0_campaign_scene
const (
	DAILY_SALE_QianchuanToolsAllowCouponV10CampaignScene                  QianchuanToolsAllowCouponV10CampaignScene = "DAILY_SALE"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanToolsAllowCouponV10CampaignScene QianchuanToolsAllowCouponV10CampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
)

// All allowed values of QianchuanToolsAllowCouponV10CampaignScene enum
var AllowedQianchuanToolsAllowCouponV10CampaignSceneEnumValues = []QianchuanToolsAllowCouponV10CampaignScene{
	"DAILY_SALE",
	"NEW_CUSTOMER_TRANSFORMATION",
}

// NewQianchuanToolsAllowCouponV10CampaignSceneFromValue returns a pointer to a valid QianchuanToolsAllowCouponV10CampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsAllowCouponV10CampaignSceneFromValue(v string) (*QianchuanToolsAllowCouponV10CampaignScene, error) {
	ev := QianchuanToolsAllowCouponV10CampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsAllowCouponV10CampaignScene: valid values are %v", v, AllowedQianchuanToolsAllowCouponV10CampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsAllowCouponV10CampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsAllowCouponV10CampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_allow_coupon_v1.0_campaign_scene value
func (v QianchuanToolsAllowCouponV10CampaignScene) Ptr() *QianchuanToolsAllowCouponV10CampaignScene {
	return &v
}
