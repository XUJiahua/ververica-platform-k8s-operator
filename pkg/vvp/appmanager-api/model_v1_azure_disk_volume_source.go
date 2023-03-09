/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.6.2
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type V1AzureDiskVolumeSource struct {
	// Host Caching mode: None, Read Only, Read Write.
	CachingMode string `json:"cachingMode,omitempty"`
	// The Name of the data disk in the blob storage
	DiskName string `json:"diskName"`
	// The URI the data disk in the blob storage
	DiskURI string `json:"diskURI"`
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType string `json:"fsType,omitempty"`
	// Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
	Kind string `json:"kind,omitempty"`
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`
}
