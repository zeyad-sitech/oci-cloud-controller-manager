// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Batch Service
//
// This is a Oracle Batch Service. You can find out more about at
// wiki (https://confluence.oraclecorp.com/confluence/display/C9QA/OCI+Batch+Service+-+Core+Functionality+Technical+Design+and+Explanation+for+Phase+I).
//

package batch

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateBatchInstanceDetails Details for creating a new batch instance.
// Batch instance is a definition that associate an OKE cluster with Batch Service,
// customer can create Compute environment, Job definition and submit Job upon Batch instance.
type CreateBatchInstanceDetails struct {

	// The OCID of the cmpartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the cluster.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// The name of the batch instance must consist of lower case alphanumeric characters or ‘-’,
	// start with an alphabetic character, and end with an alphanumeric character,
	// (e.g.’my-name’ or ‘abc-123’ or 'batchinstance')
	// When not provided, the system generate value using the format "<resourceType><timestamp>", example: batchinstance20181211220642.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBatchInstanceDetails) String() string {
	return common.PointerString(m)
}