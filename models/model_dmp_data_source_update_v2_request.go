/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpDataSourceUpdateV2Request struct for DmpDataSourceUpdateV2Request
type DmpDataSourceUpdateV2Request struct {
	//
	AdvertiserId int64                           `json:"advertiser_id"`
	DataFormat   DmpDataSourceUpdateV2DataFormat `json:"data_format"`
	//
	DataSourceId string `json:"data_source_id"`
	//
	FilePaths       []string                             `json:"file_paths"`
	FileStorageType DmpDataSourceUpdateV2FileStorageType `json:"file_storage_type"`
	OperationType   DmpDataSourceUpdateV2OperationType   `json:"operation_type"`
}
