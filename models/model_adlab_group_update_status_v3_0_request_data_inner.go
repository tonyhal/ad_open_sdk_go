/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateStatusV30RequestDataInner struct for AdlabGroupUpdateStatusV30RequestDataInner
type AdlabGroupUpdateStatusV30RequestDataInner struct {
	// 项目id
	Id        int64                                  `json:"id"`
	OptStatus AdlabGroupUpdateStatusV30DataOptStatus `json:"opt_status"`
}
