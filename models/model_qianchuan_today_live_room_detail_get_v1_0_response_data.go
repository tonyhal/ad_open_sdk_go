/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanTodayLiveRoomDetailGetV10ResponseData
type QianchuanTodayLiveRoomDetailGetV10ResponseData struct {
	//
	AwemeAvatar []string `json:"aweme_avatar,omitempty"`
	//
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	AwemeName *string `json:"aweme_name,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	RoomCover []string `json:"room_cover,omitempty"`
	//
	RoomDelivery *float64 `json:"room_delivery,omitempty"`
	//
	RoomId     *int64                                            `json:"room_id,omitempty"`
	RoomStatus *QianchuanTodayLiveRoomDetailGetV10DataRoomStatus `json:"room_status,omitempty"`
	//
	RoomTitle *string `json:"room_title,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
