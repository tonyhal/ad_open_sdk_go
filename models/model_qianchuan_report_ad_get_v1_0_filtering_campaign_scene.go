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

// QianchuanReportAdGetV10FilteringCampaignScene
type QianchuanReportAdGetV10FilteringCampaignScene string

// List of qianchuan_report_ad_get_v1.0_filtering_campaign_scene
const (
	DAILY_SALE_QianchuanReportAdGetV10FilteringCampaignScene                  QianchuanReportAdGetV10FilteringCampaignScene = "DAILY_SALE"
	LIVE_HEAT_QianchuanReportAdGetV10FilteringCampaignScene                   QianchuanReportAdGetV10FilteringCampaignScene = "LIVE_HEAT"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanReportAdGetV10FilteringCampaignScene QianchuanReportAdGetV10FilteringCampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	NEW_PRODUCT_BOOST_QianchuanReportAdGetV10FilteringCampaignScene           QianchuanReportAdGetV10FilteringCampaignScene = "NEW_PRODUCT_BOOST"
	PLANT_GRASS_QianchuanReportAdGetV10FilteringCampaignScene                 QianchuanReportAdGetV10FilteringCampaignScene = "PLANT_GRASS"
	PRODUCT_HEAT_QianchuanReportAdGetV10FilteringCampaignScene                QianchuanReportAdGetV10FilteringCampaignScene = "PRODUCT_HEAT"
)

// All allowed values of QianchuanReportAdGetV10FilteringCampaignScene enum
var AllowedQianchuanReportAdGetV10FilteringCampaignSceneEnumValues = []QianchuanReportAdGetV10FilteringCampaignScene{
	"DAILY_SALE",
	"LIVE_HEAT",
	"NEW_CUSTOMER_TRANSFORMATION",
	"NEW_PRODUCT_BOOST",
	"PLANT_GRASS",
	"PRODUCT_HEAT",
}

// NewQianchuanReportAdGetV10FilteringCampaignSceneFromValue returns a pointer to a valid QianchuanReportAdGetV10FilteringCampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdGetV10FilteringCampaignSceneFromValue(v string) (*QianchuanReportAdGetV10FilteringCampaignScene, error) {
	ev := QianchuanReportAdGetV10FilteringCampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdGetV10FilteringCampaignScene: valid values are %v", v, AllowedQianchuanReportAdGetV10FilteringCampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdGetV10FilteringCampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdGetV10FilteringCampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_get_v1.0_filtering_campaign_scene value
func (v QianchuanReportAdGetV10FilteringCampaignScene) Ptr() *QianchuanReportAdGetV10FilteringCampaignScene {
	return &v
}
