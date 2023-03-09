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

type RevisionInformation struct {
	BuildTime      string `json:"buildTime,omitempty"`
	BuildVersion   string `json:"buildVersion,omitempty"`
	CommitShaLong  string `json:"commitShaLong,omitempty"`
	CommitShaShort string `json:"commitShaShort,omitempty"`
}
