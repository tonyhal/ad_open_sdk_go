/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialListV2ResponseData
type FileMaterialListV2ResponseData struct {
	//
	Materials []*FileMaterialListV2ResponseDataMaterialsInner `json:"materials,omitempty"`
	PageInfo  *FileMaterialListV2ResponseDataPageInfo         `json:"page_info,omitempty"`
}
