/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRegionUpdateV10District
type QianchuanAdRegionUpdateV10District string

// List of qianchuan_ad_region_update_v1.0_district
const (
	CITY_QianchuanAdRegionUpdateV10District   QianchuanAdRegionUpdateV10District = "CITY"
	COUNTY_QianchuanAdRegionUpdateV10District QianchuanAdRegionUpdateV10District = "COUNTY"
	NONE_QianchuanAdRegionUpdateV10District   QianchuanAdRegionUpdateV10District = "NONE"
)

// Ptr returns reference to qianchuan_ad_region_update_v1.0_district value
func (v QianchuanAdRegionUpdateV10District) Ptr() *QianchuanAdRegionUpdateV10District {
	return &v
}
