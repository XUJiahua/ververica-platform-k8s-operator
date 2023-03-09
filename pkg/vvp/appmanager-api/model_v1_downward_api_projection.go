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

// Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.
type V1DownwardApiProjection struct {
	// Items is a list of DownwardAPIVolume file
	Items []V1DownwardApiVolumeFile `json:"items,omitempty"`
}
