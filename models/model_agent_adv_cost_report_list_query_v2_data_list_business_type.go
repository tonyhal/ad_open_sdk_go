/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvCostReportListQueryV2DataListBusinessType
type AgentAdvCostReportListQueryV2DataListBusinessType int64

// List of agent_adv_cost_report_list_query_v2_data_list_business_type
const (
	Enum_0_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 0
	Enum_1_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 1
	Enum_10_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 10
	Enum_102_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 102
	Enum_105_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 105
	Enum_107_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 107
	Enum_11_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 11
	Enum_12_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 12
	Enum_127_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 127
	Enum_13_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 13
	Enum_138_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 138
	Enum_14_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 14
	Enum_142_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 142
	Enum_143_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 143
	Enum_148_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 148
	Enum_149_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 149
	Enum_15_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 15
	Enum_150_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 150
	Enum_151_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 151
	Enum_152_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 152
	Enum_153_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 153
	Enum_154_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 154
	Enum_158_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 158
	Enum_159_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 159
	Enum_16_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 16
	Enum_160_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 160
	Enum_161_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 161
	Enum_162_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 162
	Enum_163_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 163
	Enum_164_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 164
	Enum_165_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 165
	Enum_166_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 166
	Enum_17_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 17
	Enum_18_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 18
	Enum_182_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 182
	Enum_183_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 183
	Enum_184_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 184
	Enum_186_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 186
	Enum_19_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 19
	Enum_2_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 2
	Enum_20_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 20
	Enum_21_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 21
	Enum_22_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 22
	Enum_221_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 221
	Enum_222_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 222
	Enum_23_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 23
	Enum_24_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 24
	Enum_25_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 25
	Enum_256_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 256
	Enum_257_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 257
	Enum_258_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 258
	Enum_26_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 26
	Enum_262_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 262
	Enum_263_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 263
	Enum_264_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 264
	Enum_265_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 265
	Enum_266_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 266
	Enum_267_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 267
	Enum_268_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 268
	Enum_269_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 269
	Enum_27_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 27
	Enum_270_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 270
	Enum_271_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 271
	Enum_28_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 28
	Enum_29_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 29
	Enum_3_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 3
	Enum_30_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 30
	Enum_31_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 31
	Enum_32_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 32
	Enum_33_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 33
	Enum_34_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 34
	Enum_35_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 35
	Enum_36_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 36
	Enum_37_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 37
	Enum_38_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 38
	Enum_39_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 39
	Enum_391_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 391
	Enum_392_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 392
	Enum_393_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 393
	Enum_394_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 394
	Enum_395_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 395
	Enum_4_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 4
	Enum_40_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 40
	Enum_41_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 41
	Enum_42_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 42
	Enum_43_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 43
	Enum_44_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 44
	Enum_45_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 45
	Enum_46_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 46
	Enum_47_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 47
	Enum_48_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 48
	Enum_49_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 49
	Enum_5_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 5
	Enum_50_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 50
	Enum_51_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 51
	Enum_52_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 52
	Enum_520_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 520
	Enum_53_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 53
	Enum_54_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 54
	Enum_55_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 55
	Enum_56_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 56
	Enum_57_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 57
	Enum_58_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 58
	Enum_59_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 59
	Enum_6_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 6
	Enum_60_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 60
	Enum_61_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 61
	Enum_62_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 62
	Enum_621_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 621
	Enum_622_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 622
	Enum_623_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 623
	Enum_63_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 63
	Enum_64_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 64
	Enum_65_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 65
	Enum_66_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 66
	Enum_7_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 7
	Enum_8_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 8
	Enum_9_AgentAdvCostReportListQueryV2DataListBusinessType    AgentAdvCostReportListQueryV2DataListBusinessType = 9
	Enum_94_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 94
	Enum_95_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 95
	Enum_96_AgentAdvCostReportListQueryV2DataListBusinessType   AgentAdvCostReportListQueryV2DataListBusinessType = 96
	Enum_999_AgentAdvCostReportListQueryV2DataListBusinessType  AgentAdvCostReportListQueryV2DataListBusinessType = 999
	Enum_9999_AgentAdvCostReportListQueryV2DataListBusinessType AgentAdvCostReportListQueryV2DataListBusinessType = 9999
)

// All allowed values of AgentAdvCostReportListQueryV2DataListBusinessType enum
var AllowedAgentAdvCostReportListQueryV2DataListBusinessTypeEnumValues = []AgentAdvCostReportListQueryV2DataListBusinessType{
	0,
	1,
	10,
	102,
	105,
	107,
	11,
	12,
	127,
	13,
	138,
	14,
	142,
	143,
	148,
	149,
	15,
	150,
	151,
	152,
	153,
	154,
	158,
	159,
	16,
	160,
	161,
	162,
	163,
	164,
	165,
	166,
	17,
	18,
	182,
	183,
	184,
	186,
	19,
	2,
	20,
	21,
	22,
	221,
	222,
	23,
	24,
	25,
	256,
	257,
	258,
	26,
	262,
	263,
	264,
	265,
	266,
	267,
	268,
	269,
	27,
	270,
	271,
	28,
	29,
	3,
	30,
	31,
	32,
	33,
	34,
	35,
	36,
	37,
	38,
	39,
	391,
	392,
	393,
	394,
	395,
	4,
	40,
	41,
	42,
	43,
	44,
	45,
	46,
	47,
	48,
	49,
	5,
	50,
	51,
	52,
	520,
	53,
	54,
	55,
	56,
	57,
	58,
	59,
	6,
	60,
	61,
	62,
	621,
	622,
	623,
	63,
	64,
	65,
	66,
	7,
	8,
	9,
	94,
	95,
	96,
	999,
	9999,
}

// NewAgentAdvCostReportListQueryV2DataListBusinessTypeFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListBusinessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListBusinessTypeFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListBusinessType, error) {
	ev := AgentAdvCostReportListQueryV2DataListBusinessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListBusinessType: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListBusinessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListBusinessType) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListBusinessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_business_type value
func (v AgentAdvCostReportListQueryV2DataListBusinessType) Ptr() *AgentAdvCostReportListQueryV2DataListBusinessType {
	return &v
}
