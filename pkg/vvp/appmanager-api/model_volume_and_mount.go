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

type VolumeAndMount struct {
	Name        string    `json:"name,omitempty"`
	Volume      *JsonNode `json:"volume,omitempty"`
	VolumeMount *JsonNode `json:"volumeMount,omitempty"`
}
