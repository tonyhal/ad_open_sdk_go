/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordBatchGetV30ResponseData
type ToolsPrivativeWordBatchGetV30ResponseData struct {
	//
	ProjectsPrivative []*ToolsPrivativeWordBatchGetV30ResponseDataProjectsPrivativeInner `json:"projects_privative,omitempty"`
	//
	PromotionsPrivative []*ToolsPrivativeWordBatchGetV30ResponseDataPromotionsPrivativeInner `json:"promotions_privative,omitempty"`
}
