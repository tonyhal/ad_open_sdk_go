/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportCustomConfigGetV10ResponseDataCustomConfigDatasInner struct for QianchuanReportCustomConfigGetV10ResponseDataCustomConfigDatasInner
type QianchuanReportCustomConfigGetV10ResponseDataCustomConfigDatasInner struct {
	//
	DataTopic *string `json:"data_topic,omitempty"`
	//
	Dimensions []*QianchuanReportCustomConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInner `json:"dimensions,omitempty"`
	//
	Metrics []*QianchuanReportCustomConfigGetV10ResponseDataCustomConfigDatasInnerMetricsInner `json:"metrics,omitempty"`
}
