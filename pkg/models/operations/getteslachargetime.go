// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type GetTeslaChargeTimeRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetTeslaChargeTimeResponse struct {
	// returns the date and time the vehicle expects to "complete" this charging session.
	ChargeTime  *shared.ChargeTime
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
