/*
 * PI Web API 2018 Swagger Spec
 *
 * Swagger Spec file that describes PI Web API
 *
 * API version: 1.11.0.640
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gowebapi

type Value struct {
	Value *interface{} `json:"Value,omitempty"`

	Exception *Errors `json:"Exception,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
