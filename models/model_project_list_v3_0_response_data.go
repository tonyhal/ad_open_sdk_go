/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30ResponseData
type ProjectListV30ResponseData struct {
	//
	List     []*ProjectListV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ProjectListV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
