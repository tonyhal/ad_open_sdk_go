/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportCreativeGetV10FilteringCampaignScene
type QianchuanReportCreativeGetV10FilteringCampaignScene string

// List of qianchuan_report_creative_get_v1.0_filtering_campaign_scene
const (
	DAILY_SALE_QianchuanReportCreativeGetV10FilteringCampaignScene                  QianchuanReportCreativeGetV10FilteringCampaignScene = "DAILY_SALE"
	LIVE_HEAT_QianchuanReportCreativeGetV10FilteringCampaignScene                   QianchuanReportCreativeGetV10FilteringCampaignScene = "LIVE_HEAT"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanReportCreativeGetV10FilteringCampaignScene QianchuanReportCreativeGetV10FilteringCampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	PLANT_GRASS_QianchuanReportCreativeGetV10FilteringCampaignScene                 QianchuanReportCreativeGetV10FilteringCampaignScene = "PLANT_GRASS"
	PRODUCT_HEAT_QianchuanReportCreativeGetV10FilteringCampaignScene                QianchuanReportCreativeGetV10FilteringCampaignScene = "PRODUCT_HEAT"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringCampaignScene enum
var AllowedQianchuanReportCreativeGetV10FilteringCampaignSceneEnumValues = []QianchuanReportCreativeGetV10FilteringCampaignScene{
	"DAILY_SALE",
	"LIVE_HEAT",
	"NEW_CUSTOMER_TRANSFORMATION",
	"PLANT_GRASS",
	"PRODUCT_HEAT",
}

// NewQianchuanReportCreativeGetV10FilteringCampaignSceneFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringCampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringCampaignSceneFromValue(v string) (*QianchuanReportCreativeGetV10FilteringCampaignScene, error) {
	ev := QianchuanReportCreativeGetV10FilteringCampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringCampaignScene: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringCampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringCampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringCampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_campaign_scene value
func (v QianchuanReportCreativeGetV10FilteringCampaignScene) Ptr() *QianchuanReportCreativeGetV10FilteringCampaignScene {
	return &v
}
