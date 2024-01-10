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

// StarClueGetV2DataListComponentType
type StarClueGetV2DataListComponentType string

// List of star_clue_get_v2_data_list_component_type
const (
	ANCHOR_ECOM_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "ANCHOR_ECOM"
	ANCHOR_INSURANCE_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ANCHOR_INSURANCE"
	ENTERPRISE_COUPON_StarClueGetV2DataListComponentType        StarClueGetV2DataListComponentType = "ENTERPRISE_COUPON"
	ENTERPRISE_ORDER_SERVICE_StarClueGetV2DataListComponentType StarClueGetV2DataListComponentType = "ENTERPRISE_ORDER_SERVICE"
	ENTERPRISE_DOWNLOAD_APP_StarClueGetV2DataListComponentType  StarClueGetV2DataListComponentType = "ENTERPRISE_DOWNLOAD_APP"
	ENTERPRISE_SALON_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ENTERPRISE_SALON"
	ANCHOR_XIGUA_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "ANCHOR_XIGUA"
	GAME_ANCHOR_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "GAME_ANCHOR"
	ENTERPRISE_CAR_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ENTERPRISE_CAR"
	LIVE_ORDER_COMPONENT_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "LIVE_ORDER_COMPONENT"
	ANCHOR_EDUCATION_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ANCHOR_EDUCATION"
	ANCHOR_ESTATE_SERVICE_StarClueGetV2DataListComponentType    StarClueGetV2DataListComponentType = "ANCHOR_ESTATE_SERVICE"
	MICROAPP_ANCHOR_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "MICROAPP_ANCHOR"
	ANCHOR_HOME_StarClueGetV2DataListComponentType              StarClueGetV2DataListComponentType = "ANCHOR_HOME"
	ENTERPRISE_WEDDING_PHOTO_StarClueGetV2DataListComponentType StarClueGetV2DataListComponentType = "ENTERPRISE_WEDDING_PHOTO"
	ANCHOR_DOWNLOAD_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "ANCHOR_DOWNLOAD"
	ANCHOR_CAR_StarClueGetV2DataListComponentType               StarClueGetV2DataListComponentType = "ANCHOR_CAR"
	BRAND_ANCHOR_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "BRAND_ANCHOR"
	ENTERPRISE_NOVEL_StarClueGetV2DataListComponentType         StarClueGetV2DataListComponentType = "ENTERPRISE_NOVEL"
	ANCHOR_TELECOM_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ANCHOR_TELECOM"
	CART_StarClueGetV2DataListComponentType                     StarClueGetV2DataListComponentType = "CART"
	ANCHOR_TOURISM_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "ANCHOR_TOURISM"
	POI_StarClueGetV2DataListComponentType                      StarClueGetV2DataListComponentType = "POI"
	ALL_StarClueGetV2DataListComponentType                      StarClueGetV2DataListComponentType = "ALL"
	ANCHOR_INVESTMENT_StarClueGetV2DataListComponentType        StarClueGetV2DataListComponentType = "ANCHOR_INVESTMENT"
	ENTERPRISE_DOWNLOAD_StarClueGetV2DataListComponentType      StarClueGetV2DataListComponentType = "ENTERPRISE_DOWNLOAD"
	ENTERPRISE_MICRO_APP_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "ENTERPRISE_MICRO_APP"
	ANCHOR_MICRO_APP_POI_StarClueGetV2DataListComponentType     StarClueGetV2DataListComponentType = "ANCHOR_MICRO_APP_POI"
	VARIETY_ANCHOR_StarClueGetV2DataListComponentType           StarClueGetV2DataListComponentType = "VARIETY_ANCHOR"
	ENTERPRISE_ECOM_StarClueGetV2DataListComponentType          StarClueGetV2DataListComponentType = "ENTERPRISE_ECOM"
	LINK_StarClueGetV2DataListComponentType                     StarClueGetV2DataListComponentType = "LINK"
	ANCHOR_MOVIE_StarClueGetV2DataListComponentType             StarClueGetV2DataListComponentType = "ANCHOR_MOVIE"
	ANCHOR_E_GAME_StarClueGetV2DataListComponentType            StarClueGetV2DataListComponentType = "ANCHOR_E_GAME"
)

// All allowed values of StarClueGetV2DataListComponentType enum
var AllowedStarClueGetV2DataListComponentTypeEnumValues = []StarClueGetV2DataListComponentType{
	"ANCHOR_ECOM",
	"ANCHOR_INSURANCE",
	"ENTERPRISE_COUPON",
	"ENTERPRISE_ORDER_SERVICE",
	"ENTERPRISE_DOWNLOAD_APP",
	"ENTERPRISE_SALON",
	"ANCHOR_XIGUA",
	"GAME_ANCHOR",
	"ENTERPRISE_CAR",
	"LIVE_ORDER_COMPONENT",
	"ANCHOR_EDUCATION",
	"ANCHOR_ESTATE_SERVICE",
	"MICROAPP_ANCHOR",
	"ANCHOR_HOME",
	"ENTERPRISE_WEDDING_PHOTO",
	"ANCHOR_DOWNLOAD",
	"ANCHOR_CAR",
	"BRAND_ANCHOR",
	"ENTERPRISE_NOVEL",
	"ANCHOR_TELECOM",
	"CART",
	"ANCHOR_TOURISM",
	"POI",
	"ALL",
	"ANCHOR_INVESTMENT",
	"ENTERPRISE_DOWNLOAD",
	"ENTERPRISE_MICRO_APP",
	"ANCHOR_MICRO_APP_POI",
	"VARIETY_ANCHOR",
	"ENTERPRISE_ECOM",
	"LINK",
	"ANCHOR_MOVIE",
	"ANCHOR_E_GAME",
}

// NewStarClueGetV2DataListComponentTypeFromValue returns a pointer to a valid StarClueGetV2DataListComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarClueGetV2DataListComponentTypeFromValue(v string) (*StarClueGetV2DataListComponentType, error) {
	ev := StarClueGetV2DataListComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarClueGetV2DataListComponentType: valid values are %v", v, AllowedStarClueGetV2DataListComponentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarClueGetV2DataListComponentType) IsValid() bool {
	for _, existing := range AllowedStarClueGetV2DataListComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_clue_get_v2_data_list_component_type value
func (v StarClueGetV2DataListComponentType) Ptr() *StarClueGetV2DataListComponentType {
	return &v
}
