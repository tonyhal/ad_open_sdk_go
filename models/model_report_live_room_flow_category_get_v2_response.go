/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportLiveRoomFlowCategoryGetV2Response struct for ReportLiveRoomFlowCategoryGetV2Response
type ReportLiveRoomFlowCategoryGetV2Response struct {
	//
	Code *int64                                       `json:"code,omitempty"`
	Data *ReportLiveRoomFlowCategoryGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
