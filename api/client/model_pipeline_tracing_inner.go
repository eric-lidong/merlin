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
	OperationType string          `json:"operation_type,omitempty"`
	Specs         *FreeFormObject `json:"specs,omitempty"`
	Inputs        *FreeFormObject `json:"inputs,omitempty"`
	Outputs       *FreeFormObject `json:"outputs,omitempty"`
}
