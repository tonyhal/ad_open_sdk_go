/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectListV30DataProjectListBidType
type LocalProjectListV30DataProjectListBidType string

// List of local_project_list_v3.0_data_project_list_bid_type
const (
	MANUAL_LocalProjectListV30DataProjectListBidType LocalProjectListV30DataProjectListBidType = "MANUAL"
	SMART_LocalProjectListV30DataProjectListBidType  LocalProjectListV30DataProjectListBidType = "SMART"
)

// Ptr returns reference to local_project_list_v3.0_data_project_list_bid_type value
func (v LocalProjectListV30DataProjectListBidType) Ptr() *LocalProjectListV30DataProjectListBidType {
	return &v
}