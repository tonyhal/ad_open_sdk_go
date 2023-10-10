/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpDataSourceReadV2ResponseDataDataListInner struct for DmpDataSourceReadV2ResponseDataDataListInner
type DmpDataSourceReadV2ResponseDataDataListInner struct {
	//
	ChangeLogs []*DmpDataSourceReadV2ResponseDataDataListInnerChangeLogsInner `json:"change_logs,omitempty"`
	//
	CoverNum *int64 `json:"cover_num,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	DataSourceId    *string                                                      `json:"data_source_id,omitempty"`
	DataSourceType  *DmpDataSourceReadV2DataDataListDataSourceType               `json:"data_source_type,omitempty"`
	DefaultAudience *DmpDataSourceReadV2ResponseDataDataListInnerDefaultAudience `json:"default_audience,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	LastestPublishedChangeLogId *int64 `json:"lastest_published_change_log_id,omitempty"`
	//
	LastestPublishedTime *string `json:"lastest_published_time,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Status *int64 `json:"status,omitempty"`
	//
	UploadNum *int64 `json:"upload_num,omitempty"`
}
