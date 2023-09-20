/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseOperationLogGetV10ResponseDataListInner struct for EnterpriseOperationLogGetV10ResponseDataListInner
type EnterpriseOperationLogGetV10ResponseDataListInner struct {
	//
	AdId []int64 `json:"ad_id,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Budget                    *float64                                                       `json:"budget,omitempty"`
	BusinessPageOperationType *EnterpriseOperationLogGetV10DataListBusinessPageOperationType `json:"business_page_operation_type,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	RoomCover *string `json:"room_cover,omitempty"`
	//
	RoomId *int64 `json:"room_id,omitempty"`
	//
	RoomTitle *string `json:"room_title,omitempty"`
	//
	VideoId *int64 `json:"video_id,omitempty"`
	//
	VideoImg *string `json:"video_img,omitempty"`
}
