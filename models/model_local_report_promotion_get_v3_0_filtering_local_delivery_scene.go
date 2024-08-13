/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// LocalReportPromotionGetV30FilteringLocalDeliveryScene
type LocalReportPromotionGetV30FilteringLocalDeliveryScene string

// List of local_report_promotion_get_v3.0_filtering_local_delivery_scene
const (
	CLUE_LocalReportPromotionGetV30FilteringLocalDeliveryScene            LocalReportPromotionGetV30FilteringLocalDeliveryScene = "CLUE"
	CONTENT_HEATING_LocalReportPromotionGetV30FilteringLocalDeliveryScene LocalReportPromotionGetV30FilteringLocalDeliveryScene = "CONTENT_HEATING"
	POI_CUSTOMER_LocalReportPromotionGetV30FilteringLocalDeliveryScene    LocalReportPromotionGetV30FilteringLocalDeliveryScene = "POI_CUSTOMER"
	PURCHASE_LocalReportPromotionGetV30FilteringLocalDeliveryScene        LocalReportPromotionGetV30FilteringLocalDeliveryScene = "PURCHASE"
)

// All allowed values of LocalReportPromotionGetV30FilteringLocalDeliveryScene enum
var AllowedLocalReportPromotionGetV30FilteringLocalDeliverySceneEnumValues = []LocalReportPromotionGetV30FilteringLocalDeliveryScene{
	"CLUE",
	"CONTENT_HEATING",
	"POI_CUSTOMER",
	"PURCHASE",
}

// NewLocalReportPromotionGetV30FilteringLocalDeliverySceneFromValue returns a pointer to a valid LocalReportPromotionGetV30FilteringLocalDeliveryScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalReportPromotionGetV30FilteringLocalDeliverySceneFromValue(v string) (*LocalReportPromotionGetV30FilteringLocalDeliveryScene, error) {
	ev := LocalReportPromotionGetV30FilteringLocalDeliveryScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalReportPromotionGetV30FilteringLocalDeliveryScene: valid values are %v", v, AllowedLocalReportPromotionGetV30FilteringLocalDeliverySceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalReportPromotionGetV30FilteringLocalDeliveryScene) IsValid() bool {
	for _, existing := range AllowedLocalReportPromotionGetV30FilteringLocalDeliverySceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to local_report_promotion_get_v3.0_filtering_local_delivery_scene value
func (v LocalReportPromotionGetV30FilteringLocalDeliveryScene) Ptr() *LocalReportPromotionGetV30FilteringLocalDeliveryScene {
	return &v
}
