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

// PodTemplateSpec describes the data a pod should have when created from a template
type V1PodTemplateSpec struct {
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Spec     *V1PodSpec    `json:"spec,omitempty"`
}
