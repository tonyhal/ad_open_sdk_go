/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanTrackUrlCheckV10ResponseData
type QianchuanTrackUrlCheckV10ResponseData struct {
	//
	FailTrackUrl []*QianchuanTrackUrlCheckV10ResponseDataFailTrackUrlInner `json:"fail_track_url,omitempty"`
	//
	SucTrackUrl []string `json:"suc_track_url"`
}
