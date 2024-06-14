/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInnerAudienceInfoGeolocationInner struct for BrandAdGetV30ResponseDataAdsInnerAudienceInfoGeolocationInner
type BrandAdGetV30ResponseDataAdsInnerAudienceInfoGeolocationInner struct {
	// 城市
	City *string `json:"city,omitempty"`
	// 地区
	District *string `json:"district,omitempty"`
	// 维度
	Lat *int64 `json:"lat,omitempty"`
	// 精度
	Long *int64 `json:"long,omitempty"`
	// 名称
	Name *string `json:"name,omitempty"`
	// 省份
	Province *string `json:"province,omitempty"`
	// 半径,单位米
	Radius *int64 `json:"radius,omitempty"`
	// 街道
	Street *string `json:"street,omitempty"`
	// 门牌号码
	StreetNumber *string `json:"street_number,omitempty"`
}
