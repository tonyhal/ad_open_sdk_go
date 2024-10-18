/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormDetailV2DataFormElementsElementType
type ClueFormDetailV2DataFormElementsElementType string

// List of clue_form_detail_v2_data_form_elements_element_type
const (
	NUMBER_ClueFormDetailV2DataFormElementsElementType     ClueFormDetailV2DataFormElementsElementType = "NUMBER"
	RADIO_ClueFormDetailV2DataFormElementsElementType      ClueFormDetailV2DataFormElementsElementType = "RADIO"
	CITY_ClueFormDetailV2DataFormElementsElementType       ClueFormDetailV2DataFormElementsElementType = "CITY"
	SELECT_ClueFormDetailV2DataFormElementsElementType     ClueFormDetailV2DataFormElementsElementType = "SELECT"
	DATE_ClueFormDetailV2DataFormElementsElementType       ClueFormDetailV2DataFormElementsElementType = "DATE"
	CALCULATOR_ClueFormDetailV2DataFormElementsElementType ClueFormDetailV2DataFormElementsElementType = "CALCULATOR"
	CHECKBOX_ClueFormDetailV2DataFormElementsElementType   ClueFormDetailV2DataFormElementsElementType = "CHECKBOX"
	NAME_ClueFormDetailV2DataFormElementsElementType       ClueFormDetailV2DataFormElementsElementType = "NAME"
	SEX_ClueFormDetailV2DataFormElementsElementType        ClueFormDetailV2DataFormElementsElementType = "SEX"
	TELEPHONE_ClueFormDetailV2DataFormElementsElementType  ClueFormDetailV2DataFormElementsElementType = "TELEPHONE"
	EMAIL_ClueFormDetailV2DataFormElementsElementType      ClueFormDetailV2DataFormElementsElementType = "EMAIL"
	TEXT_ClueFormDetailV2DataFormElementsElementType       ClueFormDetailV2DataFormElementsElementType = "TEXT"
	TEXTAREA_ClueFormDetailV2DataFormElementsElementType   ClueFormDetailV2DataFormElementsElementType = "TEXTAREA"
)

// Ptr returns reference to clue_form_detail_v2_data_form_elements_element_type value
func (v ClueFormDetailV2DataFormElementsElementType) Ptr() *ClueFormDetailV2DataFormElementsElementType {
	return &v
}
