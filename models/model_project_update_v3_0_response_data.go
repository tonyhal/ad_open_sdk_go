/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30ResponseData
type ProjectUpdateV30ResponseData struct {
	//
	ErrorList []*ProjectUpdateV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
}
