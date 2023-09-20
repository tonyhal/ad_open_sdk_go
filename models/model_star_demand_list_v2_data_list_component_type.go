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

// StarDemandListV2DataListComponentType
type StarDemandListV2DataListComponentType string

// List of star_demand_list_v2_data_list_component_type
const (
	LINK_StarDemandListV2DataListComponentType                     StarDemandListV2DataListComponentType = "LINK"
	ENTERPRISE_DOWNLOAD_StarDemandListV2DataListComponentType      StarDemandListV2DataListComponentType = "ENTERPRISE_DOWNLOAD"
	MICROAPP_ANCHOR_StarDemandListV2DataListComponentType          StarDemandListV2DataListComponentType = "MICROAPP_ANCHOR"
	ENTERPRISE_CAR_StarDemandListV2DataListComponentType           StarDemandListV2DataListComponentType = "ENTERPRISE_CAR"
	ANCHOR_HOME_StarDemandListV2DataListComponentType              StarDemandListV2DataListComponentType = "ANCHOR_HOME"
	POI_StarDemandListV2DataListComponentType                      StarDemandListV2DataListComponentType = "POI"
	ENTERPRISE_WEDDING_PHOTO_StarDemandListV2DataListComponentType StarDemandListV2DataListComponentType = "ENTERPRISE_WEDDING_PHOTO"
	ENTERPRISE_DOWNLOAD_APP_StarDemandListV2DataListComponentType  StarDemandListV2DataListComponentType = "ENTERPRISE_DOWNLOAD_APP"
	VARIETY_ANCHOR_StarDemandListV2DataListComponentType           StarDemandListV2DataListComponentType = "VARIETY_ANCHOR"
	ANCHOR_MOVIE_StarDemandListV2DataListComponentType             StarDemandListV2DataListComponentType = "ANCHOR_MOVIE"
	ENTERPRISE_ORDER_SERVICE_StarDemandListV2DataListComponentType StarDemandListV2DataListComponentType = "ENTERPRISE_ORDER_SERVICE"
	ENTERPRISE_SALON_StarDemandListV2DataListComponentType         StarDemandListV2DataListComponentType = "ENTERPRISE_SALON"
	ANCHOR_ESTATE_SERVICE_StarDemandListV2DataListComponentType    StarDemandListV2DataListComponentType = "ANCHOR_ESTATE_SERVICE"
	BRAND_ANCHOR_StarDemandListV2DataListComponentType             StarDemandListV2DataListComponentType = "BRAND_ANCHOR"
	ANCHOR_MICRO_APP_POI_StarDemandListV2DataListComponentType     StarDemandListV2DataListComponentType = "ANCHOR_MICRO_APP_POI"
	ALL_StarDemandListV2DataListComponentType                      StarDemandListV2DataListComponentType = "ALL"
	ANCHOR_E_GAME_StarDemandListV2DataListComponentType            StarDemandListV2DataListComponentType = "ANCHOR_E_GAME"
	ANCHOR_XIGUA_StarDemandListV2DataListComponentType             StarDemandListV2DataListComponentType = "ANCHOR_XIGUA"
	ENTERPRISE_MICRO_APP_StarDemandListV2DataListComponentType     StarDemandListV2DataListComponentType = "ENTERPRISE_MICRO_APP"
	CART_StarDemandListV2DataListComponentType                     StarDemandListV2DataListComponentType = "CART"
	LIVE_ORDER_COMPONENT_StarDemandListV2DataListComponentType     StarDemandListV2DataListComponentType = "LIVE_ORDER_COMPONENT"
	ANCHOR_TELECOM_StarDemandListV2DataListComponentType           StarDemandListV2DataListComponentType = "ANCHOR_TELECOM"
	ENTERPRISE_NOVEL_StarDemandListV2DataListComponentType         StarDemandListV2DataListComponentType = "ENTERPRISE_NOVEL"
	ANCHOR_CAR_StarDemandListV2DataListComponentType               StarDemandListV2DataListComponentType = "ANCHOR_CAR"
	ANCHOR_INSURANCE_StarDemandListV2DataListComponentType         StarDemandListV2DataListComponentType = "ANCHOR_INSURANCE"
	ENTERPRISE_COUPON_StarDemandListV2DataListComponentType        StarDemandListV2DataListComponentType = "ENTERPRISE_COUPON"
	GAME_ANCHOR_StarDemandListV2DataListComponentType              StarDemandListV2DataListComponentType = "GAME_ANCHOR"
	ANCHOR_INVESTMENT_StarDemandListV2DataListComponentType        StarDemandListV2DataListComponentType = "ANCHOR_INVESTMENT"
	ANCHOR_TOURISM_StarDemandListV2DataListComponentType           StarDemandListV2DataListComponentType = "ANCHOR_TOURISM"
	ENTERPRISE_ECOM_StarDemandListV2DataListComponentType          StarDemandListV2DataListComponentType = "ENTERPRISE_ECOM"
	ANCHOR_DOWNLOAD_StarDemandListV2DataListComponentType          StarDemandListV2DataListComponentType = "ANCHOR_DOWNLOAD"
	ANCHOR_ECOM_StarDemandListV2DataListComponentType              StarDemandListV2DataListComponentType = "ANCHOR_ECOM"
	ANCHOR_EDUCATION_StarDemandListV2DataListComponentType         StarDemandListV2DataListComponentType = "ANCHOR_EDUCATION"
)

// All allowed values of StarDemandListV2DataListComponentType enum
var AllowedStarDemandListV2DataListComponentTypeEnumValues = []StarDemandListV2DataListComponentType{
	"LINK",
	"ENTERPRISE_DOWNLOAD",
	"MICROAPP_ANCHOR",
	"ENTERPRISE_CAR",
	"ANCHOR_HOME",
	"POI",
	"ENTERPRISE_WEDDING_PHOTO",
	"ENTERPRISE_DOWNLOAD_APP",
	"VARIETY_ANCHOR",
	"ANCHOR_MOVIE",
	"ENTERPRISE_ORDER_SERVICE",
	"ENTERPRISE_SALON",
	"ANCHOR_ESTATE_SERVICE",
	"BRAND_ANCHOR",
	"ANCHOR_MICRO_APP_POI",
	"ALL",
	"ANCHOR_E_GAME",
	"ANCHOR_XIGUA",
	"ENTERPRISE_MICRO_APP",
	"CART",
	"LIVE_ORDER_COMPONENT",
	"ANCHOR_TELECOM",
	"ENTERPRISE_NOVEL",
	"ANCHOR_CAR",
	"ANCHOR_INSURANCE",
	"ENTERPRISE_COUPON",
	"GAME_ANCHOR",
	"ANCHOR_INVESTMENT",
	"ANCHOR_TOURISM",
	"ENTERPRISE_ECOM",
	"ANCHOR_DOWNLOAD",
	"ANCHOR_ECOM",
	"ANCHOR_EDUCATION",
}

// NewStarDemandListV2DataListComponentTypeFromValue returns a pointer to a valid StarDemandListV2DataListComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandListV2DataListComponentTypeFromValue(v string) (*StarDemandListV2DataListComponentType, error) {
	ev := StarDemandListV2DataListComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandListV2DataListComponentType: valid values are %v", v, AllowedStarDemandListV2DataListComponentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandListV2DataListComponentType) IsValid() bool {
	for _, existing := range AllowedStarDemandListV2DataListComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_list_v2_data_list_component_type value
func (v StarDemandListV2DataListComponentType) Ptr() *StarDemandListV2DataListComponentType {
	return &v
}
