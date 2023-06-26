// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"smartcar/pkg/models/shared"
)

type GetTeslaSpeedRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetTeslaSpeedResponse struct {
	ContentType string
	// returns the speed of a Tesla.
	Speed       *shared.Speed
	StatusCode  int
	RawResponse *http.Response
}