/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AssetsCreativeComponentUpdateV2RequestComponentInfo
type AssetsCreativeComponentUpdateV2RequestComponentInfo struct {
	//
	ComponentData map[string]interface{} `json:"component_data"`
	//
	ComponentName string                                                    `json:"component_name"`
	ComponentType AssetsCreativeComponentUpdateV2ComponentInfoComponentType `json:"component_type"`
}
