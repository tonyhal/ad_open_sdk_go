/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderUserDistributionGetV2ResponseData
type StarReportOrderUserDistributionGetV2ResponseData struct {
	//
	Activity []*StarReportOrderUserDistributionGetV2ResponseDataActivityInner `json:"activity,omitempty"`
	//
	Age []*StarReportOrderUserDistributionGetV2ResponseDataActivityInner `json:"age,omitempty"`
	//
	City []*StarReportOrderUserDistributionGetV2ResponseDataActivityInner `json:"city,omitempty"`
	//
	Device []*StarReportOrderUserDistributionGetV2ResponseDataActivityInner `json:"device,omitempty"`
	//
	Gender []*StarReportOrderUserDistributionGetV2ResponseDataActivityInner `json:"gender,omitempty"`
	//
	Interest []*StarReportOrderUserDistributionGetV2ResponseDataActivityInner `json:"interest,omitempty"`
	//
	Province []*StarReportOrderUserDistributionGetV2ResponseDataActivityInner `json:"province,omitempty"`
	//
	UpdateTime *string `json:"update_time,omitempty"`
}
