/*
 * Kyma Gateway Metadata API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

// A Publish request
type PublishRequest struct {
	// Type of the event.
	EventType string `json:"event-type"`
	// The version of the event-type. This is applicable to the data payload alone.
	EventTypeVersion string `json:"event-type-version"`
	// Optional publisher provided ID (UUID v4) of the to-be-published event. When omitted, one will be automatically generated.
	EventId string `json:"event-id,omitempty"`
	// RFC 3339 timestamp of when the event happened.
	EventTime time.Time `json:"event-time"`
	Data *AnyValue `json:"data"`
}