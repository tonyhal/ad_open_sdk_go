/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AssetsCreativeComponentGetV2Filtering
type AssetsCreativeComponentGetV2Filtering struct {
	//
	ComponentId *int64 `json:"component_id,omitempty"`
	//
	ComponentName *string `json:"component_name,omitempty"`
	//
	ComponentTypes []*AssetsCreativeComponentGetV2FilteringComponentTypes `json:"component_types,omitempty"`
	//
	Status []*AssetsCreativeComponentGetV2FilteringStatus `json:"status,omitempty"`
}
