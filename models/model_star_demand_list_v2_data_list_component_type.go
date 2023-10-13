/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
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
	ENTERPRISE_COUPON_StarDemandListV2DataListComponentType        StarDemandListV2DataListComponentType = "ENTERPRISE_COUPON"
	ANCHOR_MICRO_APP_POI_StarDemandListV2DataListComponentType     StarDemandListV2DataListComponentType = "ANCHOR_MICRO_APP_POI"
	ENTERPRISE_MICRO_APP_StarDemandListV2DataListComponentType     StarDemandListV2DataListComponentType = "ENTERPRISE_MICRO_APP"
	GAME_ANCHOR_StarDemandListV2DataListComponentType              StarDemandListV2DataListComponentType = "GAME_ANCHOR"
	ANCHOR_INVESTMENT_StarDemandListV2DataListComponentType        StarDemandListV2DataListComponentType = "ANCHOR_INVESTMENT"
	ENTERPRISE_ECOM_StarDemandListV2DataListComponentType          StarDemandListV2DataListComponentType = "ENTERPRISE_ECOM"
	ENTERPRISE_NOVEL_StarDemandListV2DataListComponentType         StarDemandListV2DataListComponentType = "ENTERPRISE_NOVEL"
	LIVE_ORDER_COMPONENT_StarDemandListV2DataListComponentType     StarDemandListV2DataListComponentType = "LIVE_ORDER_COMPONENT"
	LINK_StarDemandListV2DataListComponentType                     StarDemandListV2DataListComponentType = "LINK"
	BRAND_ANCHOR_StarDemandListV2DataListComponentType             StarDemandListV2DataListComponentType = "BRAND_ANCHOR"
	ANCHOR_XIGUA_StarDemandListV2DataListComponentType             StarDemandListV2DataListComponentType = "ANCHOR_XIGUA"
	ENTERPRISE_SALON_StarDemandListV2DataListComponentType         StarDemandListV2DataListComponentType = "ENTERPRISE_SALON"
	ANCHOR_MOVIE_StarDemandListV2DataListComponentType             StarDemandListV2DataListComponentType = "ANCHOR_MOVIE"
	ANCHOR_CAR_StarDemandListV2DataListComponentType               StarDemandListV2DataListComponentType = "ANCHOR_CAR"
	ALL_StarDemandListV2DataListComponentType                      StarDemandListV2DataListComponentType = "ALL"
	POI_StarDemandListV2DataListComponentType                      StarDemandListV2DataListComponentType = "POI"
	CART_StarDemandListV2DataListComponentType                     StarDemandListV2DataListComponentType = "CART"
	ANCHOR_TOURISM_StarDemandListV2DataListComponentType           StarDemandListV2DataListComponentType = "ANCHOR_TOURISM"
	ENTERPRISE_CAR_StarDemandListV2DataListComponentType           StarDemandListV2DataListComponentType = "ENTERPRISE_CAR"
	ENTERPRISE_DOWNLOAD_APP_StarDemandListV2DataListComponentType  StarDemandListV2DataListComponentType = "ENTERPRISE_DOWNLOAD_APP"
	ANCHOR_HOME_StarDemandListV2DataListComponentType              StarDemandListV2DataListComponentType = "ANCHOR_HOME"
	ANCHOR_E_GAME_StarDemandListV2DataListComponentType            StarDemandListV2DataListComponentType = "ANCHOR_E_GAME"
	ANCHOR_ESTATE_SERVICE_StarDemandListV2DataListComponentType    StarDemandListV2DataListComponentType = "ANCHOR_ESTATE_SERVICE"
	ENTERPRISE_DOWNLOAD_StarDemandListV2DataListComponentType      StarDemandListV2DataListComponentType = "ENTERPRISE_DOWNLOAD"
	ANCHOR_TELECOM_StarDemandListV2DataListComponentType           StarDemandListV2DataListComponentType = "ANCHOR_TELECOM"
	ENTERPRISE_ORDER_SERVICE_StarDemandListV2DataListComponentType StarDemandListV2DataListComponentType = "ENTERPRISE_ORDER_SERVICE"
	ENTERPRISE_WEDDING_PHOTO_StarDemandListV2DataListComponentType StarDemandListV2DataListComponentType = "ENTERPRISE_WEDDING_PHOTO"
	ANCHOR_ECOM_StarDemandListV2DataListComponentType              StarDemandListV2DataListComponentType = "ANCHOR_ECOM"
	ANCHOR_INSURANCE_StarDemandListV2DataListComponentType         StarDemandListV2DataListComponentType = "ANCHOR_INSURANCE"
	VARIETY_ANCHOR_StarDemandListV2DataListComponentType           StarDemandListV2DataListComponentType = "VARIETY_ANCHOR"
	ANCHOR_DOWNLOAD_StarDemandListV2DataListComponentType          StarDemandListV2DataListComponentType = "ANCHOR_DOWNLOAD"
	MICROAPP_ANCHOR_StarDemandListV2DataListComponentType          StarDemandListV2DataListComponentType = "MICROAPP_ANCHOR"
	ANCHOR_EDUCATION_StarDemandListV2DataListComponentType         StarDemandListV2DataListComponentType = "ANCHOR_EDUCATION"
)

// All allowed values of StarDemandListV2DataListComponentType enum
var AllowedStarDemandListV2DataListComponentTypeEnumValues = []StarDemandListV2DataListComponentType{
	"ENTERPRISE_COUPON",
	"ANCHOR_MICRO_APP_POI",
	"ENTERPRISE_MICRO_APP",
	"GAME_ANCHOR",
	"ANCHOR_INVESTMENT",
	"ENTERPRISE_ECOM",
	"ENTERPRISE_NOVEL",
	"LIVE_ORDER_COMPONENT",
	"LINK",
	"BRAND_ANCHOR",
	"ANCHOR_XIGUA",
	"ENTERPRISE_SALON",
	"ANCHOR_MOVIE",
	"ANCHOR_CAR",
	"ALL",
	"POI",
	"CART",
	"ANCHOR_TOURISM",
	"ENTERPRISE_CAR",
	"ENTERPRISE_DOWNLOAD_APP",
	"ANCHOR_HOME",
	"ANCHOR_E_GAME",
	"ANCHOR_ESTATE_SERVICE",
	"ENTERPRISE_DOWNLOAD",
	"ANCHOR_TELECOM",
	"ENTERPRISE_ORDER_SERVICE",
	"ENTERPRISE_WEDDING_PHOTO",
	"ANCHOR_ECOM",
	"ANCHOR_INSURANCE",
	"VARIETY_ANCHOR",
	"ANCHOR_DOWNLOAD",
	"MICROAPP_ANCHOR",
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
