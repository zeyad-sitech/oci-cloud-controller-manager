// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateJobDefinitionDetails Representation of a Job Definition Resource. Job Definitions define the harvest scope and includes the list of
// objects to be harvested along with a schedule. The list of objects is usually specified through a combination of
// object type, regular expressions or specific names of objects and a sample size for the data harvested.
type CreateJobDefinitionDetails struct {

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the Job Definition.
	JobType JobTypeEnum `mandatory:"true" json:"jobType"`

	// Detailed description of the Job Definition.
	Description *string `mandatory:"false" json:"description"`

	// Specifies if the Job Definition is incremental or full.
	IsIncremental *bool `mandatory:"false" json:"isIncremental"`

	// The key of the Data Asset for which the job is defined.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// The key of the connection resource to be used for the job.
	ConnectionKey *string `mandatory:"false" json:"connectionKey"`

	// Specify if sample data to be extracted as part of this harvest
	IsSampleDataExtracted *bool `mandatory:"false" json:"isSampleDataExtracted"`

	// Specify the sample data size in MB, specified as number of rows, for this metadata harvest
	SampleDataSizeInMBs *int `mandatory:"false" json:"sampleDataSizeInMBs"`

	// A map of maps which contains the properties which are specific to the job type. Each job type
	// definition may define it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// job definitions have required properties within the "default" category.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m CreateJobDefinitionDetails) String() string {
	return common.PointerString(m)
}