// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"smartcar/pkg/models/shared"
)

type SubscribeRequest struct {
	WebhookInfo *shared.WebhookInfo `request:"mediaType=application/json"`
	VehicleID   string              `pathParam:"style=simple,explode=false,name=vehicle_id"`
	WebhookID   string              `pathParam:"style=simple,explode=false,name=webhookId"`
}

type SubscribeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return Compatibility
	SuccessResponse *shared.SuccessResponse
}
