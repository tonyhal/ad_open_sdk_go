/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueSmartphoneGetV2DataListValidateType
type ClueSmartphoneGetV2DataListValidateType string

// List of clue_smartphone_get_v2_data_list_validate_type
const (
	VALIDITY_PRIORITY_ClueSmartphoneGetV2DataListValidateType ClueSmartphoneGetV2DataListValidateType = "VALIDITY_PRIORITY"
	ALL_VERIFICATION_ClueSmartphoneGetV2DataListValidateType  ClueSmartphoneGetV2DataListValidateType = "ALL_VERIFICATION"
	NONE_VERIFICATION_ClueSmartphoneGetV2DataListValidateType ClueSmartphoneGetV2DataListValidateType = "NONE_VERIFICATION"
	AUTO_VERIFICATION_ClueSmartphoneGetV2DataListValidateType ClueSmartphoneGetV2DataListValidateType = "AUTO_VERIFICATION"
	CLUE_PRIORITY_ClueSmartphoneGetV2DataListValidateType     ClueSmartphoneGetV2DataListValidateType = "CLUE_PRIORITY"
)

// Ptr returns reference to clue_smartphone_get_v2_data_list_validate_type value
func (v ClueSmartphoneGetV2DataListValidateType) Ptr() *ClueSmartphoneGetV2DataListValidateType {
	return &v
}
