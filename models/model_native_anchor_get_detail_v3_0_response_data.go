/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseData
type NativeAnchorGetDetailV30ResponseData struct {
	//
	ErrorList []*NativeAnchorGetDetailV30ResponseDataErrorListInner `json:"error_list"`
	//
	List []*NativeAnchorGetDetailV30ResponseDataListInner `json:"list"`
}
