package instancemetadata

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"os"
)

func NewDev() InstanceMetadata {
	return &instanceMetadata{
		tenantID:       os.Getenv("AZURE_TENANT_ID"),
		subscriptionID: os.Getenv("AZURE_SUBSCRIPTION_ID"),
		location:       os.Getenv("LOCATION"),
		resourceGroup:  os.Getenv("RESOURCEGROUP"),
	}
}
