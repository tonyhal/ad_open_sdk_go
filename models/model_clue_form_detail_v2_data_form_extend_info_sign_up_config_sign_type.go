/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType
type ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType string

// List of clue_form_detail_v2_data_form_extend_info_sign_up_config_sign_type
const (
	SIGN_TYPE_SCROLL_BAR_ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType  ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_BAR"
	SIGN_TYPE_SCROLL_WALL_ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_WALL"
)

// Ptr returns reference to clue_form_detail_v2_data_form_extend_info_sign_up_config_sign_type value
func (v ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType) Ptr() *ClueFormDetailV2DataFormExtendInfoSignUpConfigSignType {
	return &v
}
