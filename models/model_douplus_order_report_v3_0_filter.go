/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderReportV30Filter
type DouplusOrderReportV30Filter struct {
	//
	AuthorUniqueIds []string `json:"author_unique_ids,omitempty"`
	//
	ItemIds         []int64                                     `json:"item_ids,omitempty"`
	OrderCreateTime *DouplusOrderReportV30FilterOrderCreateTime `json:"order_create_time,omitempty"`
	//
	OrderIds []int64 `json:"order_ids,omitempty"`
	//
	RoomIds []int64 `json:"room_ids,omitempty"`
}
