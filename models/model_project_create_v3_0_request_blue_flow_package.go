/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestBlueFlowPackage
type ProjectCreateV30RequestBlueFlowPackage struct {
	//
	BlueFlowPackageId      *int64                                                 `json:"blue_flow_package_id,omitempty"`
	BlueFlowPackageSetting *ProjectCreateV30BlueFlowPackageBlueFlowPackageSetting `json:"blue_flow_package_setting,omitempty"`
}
