// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RateLimitingPolicy Limit the number of requests that should be handled for the specified window using a spefic key.
type RateLimitingPolicy struct {

	// The maximum number of requests per second to allow.
	RateInRequestsPerSecond *int `mandatory:"true" json:"rateInRequestsPerSecond"`

	// The key used to group requests together.
	RateKey RateLimitingPolicyRateKeyEnum `mandatory:"true" json:"rateKey"`
}

func (m RateLimitingPolicy) String() string {
	return common.PointerString(m)
}

// RateLimitingPolicyRateKeyEnum Enum with underlying type: string
type RateLimitingPolicyRateKeyEnum string

// Set of constants representing the allowable values for RateLimitingPolicyRateKeyEnum
const (
	RateLimitingPolicyRateKeyClientIp RateLimitingPolicyRateKeyEnum = "CLIENT_IP"
	RateLimitingPolicyRateKeyTotal    RateLimitingPolicyRateKeyEnum = "TOTAL"
)

var mappingRateLimitingPolicyRateKey = map[string]RateLimitingPolicyRateKeyEnum{
	"CLIENT_IP": RateLimitingPolicyRateKeyClientIp,
	"TOTAL":     RateLimitingPolicyRateKeyTotal,
}

// GetRateLimitingPolicyRateKeyEnumValues Enumerates the set of values for RateLimitingPolicyRateKeyEnum
func GetRateLimitingPolicyRateKeyEnumValues() []RateLimitingPolicyRateKeyEnum {
	values := make([]RateLimitingPolicyRateKeyEnum, 0)
	for _, v := range mappingRateLimitingPolicyRateKey {
		values = append(values, v)
	}
	return values
}