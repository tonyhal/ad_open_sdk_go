/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallbackV2EventDataReasonCode
type ToolsClueCallbackV2EventDataReasonCode string

// List of tools_clue_callback_v2_event_data_reason_code
const (
	CALL_THREE_TIMES_NO_RESPONSE_ToolsClueCallbackV2EventDataReasonCode         ToolsClueCallbackV2EventDataReasonCode = "CALL_THREE_TIMES_NO_RESPONSE"
	CONNECTED_NO_INTENTION_ToolsClueCallbackV2EventDataReasonCode               ToolsClueCallbackV2EventDataReasonCode = "CONNECTED_NO_INTENTION"
	EDU_FIRST_CLASS_ATTENDED_ToolsClueCallbackV2EventDataReasonCode             ToolsClueCallbackV2EventDataReasonCode = "EDU_FIRST_CLASS_ATTENDED"
	EDU_INTERESTED_ToolsClueCallbackV2EventDataReasonCode                       ToolsClueCallbackV2EventDataReasonCode = "EDU_INTERESTED"
	EDU_MISMATCH_CLASS_TIME_ToolsClueCallbackV2EventDataReasonCode              ToolsClueCallbackV2EventDataReasonCode = "EDU_MISMATCH_CLASS_TIME"
	EDU_MISMATCH_COURSE_ToolsClueCallbackV2EventDataReasonCode                  ToolsClueCallbackV2EventDataReasonCode = "EDU_MISMATCH_COURSE"
	EDU_MISMATCH_DIFFERENT_LOCATIONS_ToolsClueCallbackV2EventDataReasonCode     ToolsClueCallbackV2EventDataReasonCode = "EDU_MISMATCH_DIFFERENT_LOCATIONS"
	EDU_UNABLE_TO_ATTEND_IN_PERSON_ToolsClueCallbackV2EventDataReasonCode       ToolsClueCallbackV2EventDataReasonCode = "EDU_UNABLE_TO_ATTEND_IN_PERSON"
	HOME_APPOINTMENT_SCHEDULED_ToolsClueCallbackV2EventDataReasonCode           ToolsClueCallbackV2EventDataReasonCode = "HOME_APPOINTMENT_SCHEDULED"
	HOME_CAN_ADD_WECHAT_ToolsClueCallbackV2EventDataReasonCode                  ToolsClueCallbackV2EventDataReasonCode = "HOME_CAN_ADD_WECHAT"
	HOME_COLLABORATION_WITH_PEERS_ToolsClueCallbackV2EventDataReasonCode        ToolsClueCallbackV2EventDataReasonCode = "HOME_COLLABORATION_WITH_PEERS"
	HOME_CONNECTED_WITH_INTERESTED_PARTY_ToolsClueCallbackV2EventDataReasonCode ToolsClueCallbackV2EventDataReasonCode = "HOME_CONNECTED_WITH_INTERESTED_PARTY"
	HOME_CONTRACT_SIGNED_ToolsClueCallbackV2EventDataReasonCode                 ToolsClueCallbackV2EventDataReasonCode = "HOME_CONTRACT_SIGNED"
	HOME_MEASUREMENT_COMPLETED_ToolsClueCallbackV2EventDataReasonCode           ToolsClueCallbackV2EventDataReasonCode = "HOME_MEASUREMENT_COMPLETED"
	HOME_MISMATCH_DIFFERENT_LOCATIONS_ToolsClueCallbackV2EventDataReasonCode    ToolsClueCallbackV2EventDataReasonCode = "HOME_MISMATCH_DIFFERENT_LOCATIONS"
	HOME_MISMATCH_LOW_BUDGET_ToolsClueCallbackV2EventDataReasonCode             ToolsClueCallbackV2EventDataReasonCode = "HOME_MISMATCH_LOW_BUDGET"
	HOME_NO_SHOW_FOR_MEASUREMENT_VISIT_ToolsClueCallbackV2EventDataReasonCode   ToolsClueCallbackV2EventDataReasonCode = "HOME_NO_SHOW_FOR_MEASUREMENT_VISIT"
	HOME_ORDER_DISPATCHED_ToolsClueCallbackV2EventDataReasonCode                ToolsClueCallbackV2EventDataReasonCode = "HOME_ORDER_DISPATCHED"
	HOME_RENOVATION_TYPE_NOT_ACCEPTABLE_ToolsClueCallbackV2EventDataReasonCode  ToolsClueCallbackV2EventDataReasonCode = "HOME_RENOVATION_TYPE_NOT_ACCEPTABLE"
	HOME_SERVICE_DELIVERY_COMPLETED_ToolsClueCallbackV2EventDataReasonCode      ToolsClueCallbackV2EventDataReasonCode = "HOME_SERVICE_DELIVERY_COMPLETED"
	INVALID_NUMBER_ToolsClueCallbackV2EventDataReasonCode                       ToolsClueCallbackV2EventDataReasonCode = "INVALID_NUMBER"
	MISMATCH_DIFFERENT_LOCATIONS_ToolsClueCallbackV2EventDataReasonCode         ToolsClueCallbackV2EventDataReasonCode = "MISMATCH_DIFFERENT_LOCATIONS"
	MISMATCH_LOW_BUDGET_ToolsClueCallbackV2EventDataReasonCode                  ToolsClueCallbackV2EventDataReasonCode = "MISMATCH_LOW_BUDGET"
	NO_AD_RESPONSE_ToolsClueCallbackV2EventDataReasonCode                       ToolsClueCallbackV2EventDataReasonCode = "NO_AD_RESPONSE"
	OFFENSIVE_LANGUAGE_ToolsClueCallbackV2EventDataReasonCode                   ToolsClueCallbackV2EventDataReasonCode = "OFFENSIVE_LANGUAGE"
	OTHER_ToolsClueCallbackV2EventDataReasonCode                                ToolsClueCallbackV2EventDataReasonCode = "OTHER"
)

// Ptr returns reference to tools_clue_callback_v2_event_data_reason_code value
func (v ToolsClueCallbackV2EventDataReasonCode) Ptr() *ToolsClueCallbackV2EventDataReasonCode {
	return &v
}