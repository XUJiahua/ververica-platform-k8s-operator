/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.1.0
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

import (
	"time"
)

type DeploymentTargetMetadata struct {
	Id              string            `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	Namespace       string            `json:"namespace,omitempty"`
	CreatedAt       time.Time         `json:"createdAt,omitempty"`
	ModifiedAt      time.Time         `json:"modifiedAt,omitempty"`
	Annotations     map[string]string `json:"annotations,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	ResourceVersion int32             `json:"resourceVersion,omitempty"`
}
