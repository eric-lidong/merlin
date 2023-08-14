/*
 * Merlin
 *
 * API Guide for accessing Merlin's model management, deployment, and serving functionalities
 *
 * API version: 0.14.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type PipelineTracingInner struct {
	OperationType string    `json:"operation_type,omitempty"`
	Specs         *ModelMap `json:"specs,omitempty"`
	Inputs        *ModelMap `json:"inputs,omitempty"`
	Outputs       *ModelMap `json:"outputs,omitempty"`
}
