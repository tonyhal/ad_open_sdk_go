/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorQrcodePreviewGetV30ResponseData
type NativeAnchorQrcodePreviewGetV30ResponseData struct {
	//
	ErrorList []*NativeAnchorQrcodePreviewGetV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []*NativeAnchorQrcodePreviewGetV30ResponseDataSuccessListInner `json:"success_list,omitempty"`
}
