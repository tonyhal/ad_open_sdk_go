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

// QueryRebateBalanceV2FilteringMonthQuarter
type QueryRebateBalanceV2FilteringMonthQuarter string

// List of query_rebate_balance_v2_filtering_month_quarter
const (
	ALL_YEAR_QueryRebateBalanceV2FilteringMonthQuarter QueryRebateBalanceV2FilteringMonthQuarter = "AllYear"
	APR_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Apr"
	AUG_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Aug"
	DEC_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Dec"
	FEB_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Feb"
	FIR_HALF_QueryRebateBalanceV2FilteringMonthQuarter QueryRebateBalanceV2FilteringMonthQuarter = "FirHalf"
	JAN_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Jan"
	JUL_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Jul"
	JUN_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Jun"
	MAR_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Mar"
	MAY_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "May"
	NOV_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Nov"
	OCT_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Oct"
	Q1_QueryRebateBalanceV2FilteringMonthQuarter       QueryRebateBalanceV2FilteringMonthQuarter = "Q1"
	Q2_QueryRebateBalanceV2FilteringMonthQuarter       QueryRebateBalanceV2FilteringMonthQuarter = "Q2"
	Q3_QueryRebateBalanceV2FilteringMonthQuarter       QueryRebateBalanceV2FilteringMonthQuarter = "Q3"
	Q4_QueryRebateBalanceV2FilteringMonthQuarter       QueryRebateBalanceV2FilteringMonthQuarter = "Q4"
	SEC_HALF_QueryRebateBalanceV2FilteringMonthQuarter QueryRebateBalanceV2FilteringMonthQuarter = "SecHalf"
	SEP_QueryRebateBalanceV2FilteringMonthQuarter      QueryRebateBalanceV2FilteringMonthQuarter = "Sep"
)

// All allowed values of QueryRebateBalanceV2FilteringMonthQuarter enum
var AllowedQueryRebateBalanceV2FilteringMonthQuarterEnumValues = []QueryRebateBalanceV2FilteringMonthQuarter{
	"AllYear",
	"Apr",
	"Aug",
	"Dec",
	"Feb",
	"FirHalf",
	"Jan",
	"Jul",
	"Jun",
	"Mar",
	"May",
	"Nov",
	"Oct",
	"Q1",
	"Q2",
	"Q3",
	"Q4",
	"SecHalf",
	"Sep",
}

// NewQueryRebateBalanceV2FilteringMonthQuarterFromValue returns a pointer to a valid QueryRebateBalanceV2FilteringMonthQuarter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateBalanceV2FilteringMonthQuarterFromValue(v string) (*QueryRebateBalanceV2FilteringMonthQuarter, error) {
	ev := QueryRebateBalanceV2FilteringMonthQuarter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateBalanceV2FilteringMonthQuarter: valid values are %v", v, AllowedQueryRebateBalanceV2FilteringMonthQuarterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateBalanceV2FilteringMonthQuarter) IsValid() bool {
	for _, existing := range AllowedQueryRebateBalanceV2FilteringMonthQuarterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_balance_v2_filtering_month_quarter value
func (v QueryRebateBalanceV2FilteringMonthQuarter) Ptr() *QueryRebateBalanceV2FilteringMonthQuarter {
	return &v
}
