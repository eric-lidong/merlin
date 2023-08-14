/*
 * Merlin
 *
 * API Guide for accessing Merlin's model management, deployment, and serving functionalities
 *
 * API version: 0.14.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type PredictionJobConfigModelResult struct {
	Type_    *ResultType `json:"type,omitempty"`
	ItemType *ResultType `json:"item_type,omitempty"`
}
