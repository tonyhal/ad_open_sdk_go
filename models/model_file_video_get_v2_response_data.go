/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoGetV2ResponseData
type FileVideoGetV2ResponseData struct {
	// json返回值
	List     []*FileVideoGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *FileVideoGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
