/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// This configuration allows executing raw containers in Flyte using the Flyte CoPilot system. Flyte CoPilot, eliminates the needs of flytekit or sdk inside the container. Any inputs required by the users container are side-loaded in the input_path Any outputs generated by the user container - within output_path are automatically uploaded.
type CoreDataLoadingConfig struct {
	Enabled bool `json:"enabled,omitempty"`
	InputPath string `json:"input_path,omitempty"`
	OutputPath string `json:"output_path,omitempty"`
	Format *DataLoadingConfigLiteralMapFormat `json:"format,omitempty"`
	IoStrategy *CoreIoStrategy `json:"io_strategy,omitempty"`
}
