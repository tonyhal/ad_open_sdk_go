/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaProductAvailablesV2DataListProductIndustry
type DpaProductAvailablesV2DataListProductIndustry string

// List of dpa_product_availables_v2_data_list_product_industry
const (
	MEDICINE_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "MEDICINE"
	FURNITURE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "FURNITURE"
	GENERAL_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "GENERAL"
	AUTO_NEW_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_NEW"
	CREDIT_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "CREDIT"
	AUTO_OLD_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_OLD"
	OTHER_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "OTHER"
	TOUR_ROUTE_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_ROUTE"
	VIDEO_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "VIDEO"
	RECRUITMENT_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "RECRUITMENT"
	TOUR_HOTEL_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_HOTEL"
	TOUR_TICKET_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "TOUR_TICKET"
	FINANCE_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "FINANCE"
	GAME_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "GAME"
	MERCHANTS_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "MERCHANTS"
	COMMUNICATION_DpaProductAvailablesV2DataListProductIndustry    DpaProductAvailablesV2DataListProductIndustry = "COMMUNICATION"
	NOVEL_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "NOVEL"
	ESTATE_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "ESTATE"
	TRANSPORT_TICKET_DpaProductAvailablesV2DataListProductIndustry DpaProductAvailablesV2DataListProductIndustry = "TRANSPORT_TICKET"
	EDUCATION_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "EDUCATION"
	ECOMMERCE_V2_DpaProductAvailablesV2DataListProductIndustry     DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE_V2"
	NEW_HOUSE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "NEW_HOUSE"
	WEALTH_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "WEALTH"
	MEDICAL_SERVICE_DpaProductAvailablesV2DataListProductIndustry  DpaProductAvailablesV2DataListProductIndustry = "MEDICAL_SERVICE"
	ECOMMERCE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE"
	LIVE_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "LIVE"
)

// All allowed values of DpaProductAvailablesV2DataListProductIndustry enum
var AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues = []DpaProductAvailablesV2DataListProductIndustry{
	"MEDICINE",
	"FURNITURE",
	"GENERAL",
	"AUTO_NEW",
	"CREDIT",
	"AUTO_OLD",
	"OTHER",
	"TOUR_ROUTE",
	"VIDEO",
	"RECRUITMENT",
	"TOUR_HOTEL",
	"TOUR_TICKET",
	"FINANCE",
	"GAME",
	"MERCHANTS",
	"COMMUNICATION",
	"NOVEL",
	"ESTATE",
	"TRANSPORT_TICKET",
	"EDUCATION",
	"ECOMMERCE_V2",
	"NEW_HOUSE",
	"WEALTH",
	"MEDICAL_SERVICE",
	"ECOMMERCE",
	"LIVE",
}

// NewDpaProductAvailablesV2DataListProductIndustryFromValue returns a pointer to a valid DpaProductAvailablesV2DataListProductIndustry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaProductAvailablesV2DataListProductIndustryFromValue(v string) (*DpaProductAvailablesV2DataListProductIndustry, error) {
	ev := DpaProductAvailablesV2DataListProductIndustry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaProductAvailablesV2DataListProductIndustry: valid values are %v", v, AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaProductAvailablesV2DataListProductIndustry) IsValid() bool {
	for _, existing := range AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_product_availables_v2_data_list_product_industry value
func (v DpaProductAvailablesV2DataListProductIndustry) Ptr() *DpaProductAvailablesV2DataListProductIndustry {
	return &v
}
