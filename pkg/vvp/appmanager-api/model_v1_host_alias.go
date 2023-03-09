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

// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
type V1HostAlias struct {
	// Hostnames for the above IP address.
	Hostnames []string `json:"hostnames,omitempty"`
	// IP address of the host file entry.
	Ip string `json:"ip,omitempty"`
}
