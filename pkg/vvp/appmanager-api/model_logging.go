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

type Logging struct {
	LoggingProfile              string            `json:"loggingProfile,omitempty"`
	Log4j2ConfigurationTemplate string            `json:"log4j2ConfigurationTemplate,omitempty"`
	Log4jLoggers                map[string]string `json:"log4jLoggers,omitempty"`
}
