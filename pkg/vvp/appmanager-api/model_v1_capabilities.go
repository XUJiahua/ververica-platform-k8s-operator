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

// Adds and removes POSIX capabilities from running containers.
type V1Capabilities struct {
	// Added capabilities
	Add []string `json:"add,omitempty"`
	// Removed capabilities
	Drop []string `json:"drop,omitempty"`
}
