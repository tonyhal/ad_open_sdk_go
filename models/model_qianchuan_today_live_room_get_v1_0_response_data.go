/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanTodayLiveRoomGetV10ResponseData
type QianchuanTodayLiveRoomGetV10ResponseData struct {
	PageInfo *QianchuanTodayLiveRoomGetV10ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	RoomList []*QianchuanTodayLiveRoomGetV10ResponseDataRoomListInner `json:"room_list,omitempty"`
}
