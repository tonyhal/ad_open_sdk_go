/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestAdInfoAudienceGeolocationInner struct for AdlabGroupCreateV30RequestAdInfoAudienceGeolocationInner
type AdlabGroupCreateV30RequestAdInfoAudienceGeolocationInner struct {
	// 纬度
	Lat *float64 `json:"lat,omitempty"`
	// 经度
	Long *float64 `json:"long,omitempty"`
	// 地点名称
	Name *string `json:"name,omitempty"`
	// 半径，单位为m，允许值为：3000-15000m
	Radius *int64 `json:"radius,omitempty"`
}
