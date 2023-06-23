// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"smartcar/pkg/models/shared"
)

type GetBatteryCapacityRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetBatteryCapacityResponse struct {
	// return EV Battery Capacity reading
	BatteryCapacity *shared.BatteryCapacity
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
}
