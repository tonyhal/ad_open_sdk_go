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

// QianchuanAdDetailGetV10DataCampaignScene
type QianchuanAdDetailGetV10DataCampaignScene string

// List of qianchuan_ad_detail_get_v1.0_data_campaign_scene
const (
	DAILY_SALE_QianchuanAdDetailGetV10DataCampaignScene                  QianchuanAdDetailGetV10DataCampaignScene = "DAILY_SALE"
	LIVE_HEAT_QianchuanAdDetailGetV10DataCampaignScene                   QianchuanAdDetailGetV10DataCampaignScene = "LIVE_HEAT"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanAdDetailGetV10DataCampaignScene QianchuanAdDetailGetV10DataCampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	NEW_PRODUCT_BOOST_QianchuanAdDetailGetV10DataCampaignScene           QianchuanAdDetailGetV10DataCampaignScene = "NEW_PRODUCT_BOOST"
	PLANT_GRASS_QianchuanAdDetailGetV10DataCampaignScene                 QianchuanAdDetailGetV10DataCampaignScene = "PLANT_GRASS"
	PRODUCT_HEAT_QianchuanAdDetailGetV10DataCampaignScene                QianchuanAdDetailGetV10DataCampaignScene = "PRODUCT_HEAT"
)

// All allowed values of QianchuanAdDetailGetV10DataCampaignScene enum
var AllowedQianchuanAdDetailGetV10DataCampaignSceneEnumValues = []QianchuanAdDetailGetV10DataCampaignScene{
	"DAILY_SALE",
	"LIVE_HEAT",
	"NEW_CUSTOMER_TRANSFORMATION",
	"NEW_PRODUCT_BOOST",
	"PLANT_GRASS",
	"PRODUCT_HEAT",
}

// NewQianchuanAdDetailGetV10DataCampaignSceneFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataCampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataCampaignSceneFromValue(v string) (*QianchuanAdDetailGetV10DataCampaignScene, error) {
	ev := QianchuanAdDetailGetV10DataCampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataCampaignScene: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataCampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataCampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataCampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_campaign_scene value
func (v QianchuanAdDetailGetV10DataCampaignScene) Ptr() *QianchuanAdDetailGetV10DataCampaignScene {
	return &v
}
