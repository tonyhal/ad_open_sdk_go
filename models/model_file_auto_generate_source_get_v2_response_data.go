/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileAutoGenerateSourceGetV2ResponseData
type FileAutoGenerateSourceGetV2ResponseData struct {
	//
	Errors []*FileAutoGenerateSourceGetV2ResponseDataErrorsInner `json:"errors,omitempty"`
	//
	List []*FileAutoGenerateSourceGetV2ResponseDataListInner `json:"list,omitempty"`
}