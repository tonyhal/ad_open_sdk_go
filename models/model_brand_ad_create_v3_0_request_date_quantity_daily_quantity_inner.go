/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdCreateV30RequestDateQuantityDailyQuantityInner struct for BrandAdCreateV30RequestDateQuantityDailyQuantityInner
type BrandAdCreateV30RequestDateQuantityDailyQuantityInner struct {
	// 日期
	Date string `json:"date"`
	// 时段
	Period []string `json:"period,omitempty"`
	// 预订量
	Quantity int64 `json:"quantity"`
}
