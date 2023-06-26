// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type SetChargingLimitRequest struct {
	ChargeLimit *shared.ChargeLimit `request:"mediaType=application/json"`
	VehicleID   string              `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type SetChargingLimitResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return Success Response
	SuccessResponse *shared.SuccessResponse
}
