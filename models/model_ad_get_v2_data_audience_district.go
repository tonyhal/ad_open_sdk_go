/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceDistrict
type AdGetV2DataAudienceDistrict string

// List of ad_get_v2_data_audience_district
const (
	BUSINESS_DISTRICT_AdGetV2DataAudienceDistrict AdGetV2DataAudienceDistrict = "BUSINESS_DISTRICT"
	NONE_AdGetV2DataAudienceDistrict              AdGetV2DataAudienceDistrict = "NONE"
	REGION_AdGetV2DataAudienceDistrict            AdGetV2DataAudienceDistrict = "REGION"
	CITY_AdGetV2DataAudienceDistrict              AdGetV2DataAudienceDistrict = "CITY"
	COUNTY_AdGetV2DataAudienceDistrict            AdGetV2DataAudienceDistrict = "COUNTY"
	OVERSEA_AdGetV2DataAudienceDistrict           AdGetV2DataAudienceDistrict = "OVERSEA"
)

// Ptr returns reference to ad_get_v2_data_audience_district value
func (v AdGetV2DataAudienceDistrict) Ptr() *AdGetV2DataAudienceDistrict {
	return &v
}
