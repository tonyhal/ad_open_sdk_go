/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpDataSourceCreateV2Request struct for DmpDataSourceCreateV2Request
type DmpDataSourceCreateV2Request struct {
	//
	AdvertiserId int64                           `json:"advertiser_id"`
	DataFormat   DmpDataSourceCreateV2DataFormat `json:"data_format"`
	//
	DataSourceName string                               `json:"data_source_name"`
	DataSourceType *DmpDataSourceCreateV2DataSourceType `json:"data_source_type,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	FilePaths       []string                             `json:"file_paths"`
	FileStorageType DmpDataSourceCreateV2FileStorageType `json:"file_storage_type"`
}
