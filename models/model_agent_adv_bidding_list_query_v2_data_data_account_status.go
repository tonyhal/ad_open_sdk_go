/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvBiddingListQueryV2DataDataAccountStatus
type AgentAdvBiddingListQueryV2DataDataAccountStatus string

// List of agent_adv_bidding_list_query_v2_data_data_account_status
const (
	STATUS_DISABLE_AgentAdvBiddingListQueryV2DataDataAccountStatus                    AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_DISABLE"
	STATUS_PENDING_CONFIRM_AgentAdvBiddingListQueryV2DataDataAccountStatus            AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_PENDING_CONFIRM"
	STATUS_PENDING_VERIFIED_AgentAdvBiddingListQueryV2DataDataAccountStatus           AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_PENDING_VERIFIED"
	STATUS_CONFIRM_FAIL_AgentAdvBiddingListQueryV2DataDataAccountStatus               AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_CONFIRM_FAIL"
	STATUS_ENABLE_AgentAdvBiddingListQueryV2DataDataAccountStatus                     AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_ENABLE"
	STATUS_CONFIRM_FAIL_END_AgentAdvBiddingListQueryV2DataDataAccountStatus           AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_CONFIRM_FAIL_END"
	STATUS_PENDING_CONFIRM_MODIFY_AgentAdvBiddingListQueryV2DataDataAccountStatus     AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	STATUS_CONFIRM_MODIFY_FAIL_AgentAdvBiddingListQueryV2DataDataAccountStatus        AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	STATUS_PUNISH_AgentAdvBiddingListQueryV2DataDataAccountStatus                     AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_PUNISH"
	STATUS_WAIT_FOR_BPM_AUDIT_AgentAdvBiddingListQueryV2DataDataAccountStatus         AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	STATUS_SELF_SERVICE_UNAUDITED_AgentAdvBiddingListQueryV2DataDataAccountStatus     AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_ENABLE_AND_AVATAR_AUDITING_AgentAdvBiddingListQueryV2DataDataAccountStatus AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_ENABLE_AND_AVATAR_AUDITING"
	STATUS_WAIT_FOR_BPM_FILE_CONTACT_AgentAdvBiddingListQueryV2DataDataAccountStatus  AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_WAIT_FOR_BPM_FILE_CONTACT"
	STATUS_WAIT_FOR_ACCOUNT_FEE_AgentAdvBiddingListQueryV2DataDataAccountStatus       AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_WAIT_FOR_ACCOUNT_FEE"
	STATUS_WAIT_FOR_PUBLIC_AUTH_AgentAdvBiddingListQueryV2DataDataAccountStatus       AgentAdvBiddingListQueryV2DataDataAccountStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
)

// Ptr returns reference to agent_adv_bidding_list_query_v2_data_data_account_status value
func (v AgentAdvBiddingListQueryV2DataDataAccountStatus) Ptr() *AgentAdvBiddingListQueryV2DataDataAccountStatus {
	return &v
}