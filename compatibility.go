// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package smartcar

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"smartcar/pkg/models/operations"
	"smartcar/pkg/models/shared"
	"smartcar/pkg/utils"
	"strings"
)

// compatibility - Operations about compatibility
type compatibility struct {
	sdkConfiguration sdkConfiguration
}

func newCompatibility(sdkConfig sdkConfiguration) *compatibility {
	return &compatibility{
		sdkConfiguration: sdkConfig,
	}
}

// ListCompatibility - Compatibility
// In the US, compatibility will return a breakdown by scope of what a car is capable of. In Europe, the check is based on the make of the car so will return only a `true` or `false`
//
// The Compatibility API lets developers determine whether a given vehicle is compatible with Smartcar and whether it is capable of the features their application requires. Using this endpoint, developers can determine whether specific vehicles are eligible before sending a user through Smartcar Connect.
//
// A compatible vehicle is one that:
// 1. Has the hardware required for internet connectivity
// 2. Belongs to the makes and models Smartcar is compatible with in the vehicle's country
//
// A vehicle is capable of given feature if:
// 1. The vehicle supports the feature (e.g., a Ford Escape supports /fuel but a Mustang Mach-e does not)
// 2. Smartcar supports the feature for the vehicle's make
//
// Note: The Compatibility API is currently available for vehicles sold in the United States. For other countries, please contact us to request early access. Our initial release for vehicles sold outside of the United States only supports checking compatibility. It does not yet support checking `capabilities`.
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  compatible|   boolean|  `true` if the vehicle is likely compatible.*, `false` otherwise.|
// reason|   string \| null|  One of the enum values described below if compatible is `false`, `null` otherwise.|
// | capabilities|   array| An array containing the set of endpoints that the provided scope value can provide authorization for. This array will be empty if `compatible` is false|
// | capabilities[].permission|   string|One of the permissions provided in the scope parameter.|
// | capabilities[].endpoint|   string| One of the endpoints that the permission authorizes access to.|
// | capabilities[].capable|   boolean|`true` if the vehicle is likely capable of this feature.*, `false` otherwise.|
// | capabilities[].reason|   string \| null|One of the enum values described below if capable is `false`, `null` otherwise.|
//
// __Note__: The compatibility of a vehicle depends on many factors including its make, model, model year, trim, package, and head unit. The Smartcar Compatibility API is optimized to return false positives rather than false negatives.
//
// __Enum Values__
//
// |  Parameter 	|Value   	|Description   	|
// |---	|---	|---	|
// |  reason|   VEHICLE_NOT_COMPATIBLE|  The vehicle does not have the hardware required for internet connectivity.|
// |  |   MAKE_NOT_COMPATIBLE|  Smartcar is not yet compatible with the vehicle's make in the specified country.|
// |  capabilities[].reason|   VEHICLE_NOT_CAPABLE|  TThe vehicle does not support this feature.|
// |  |   SMARTCAR_NOT_CAPABLE|  Smartcar is not capable of supporting the given feature on the vehicle's make.|
func (s *compatibility) ListCompatibility(ctx context.Context, request operations.ListCompatibilityRequest) (*operations.ListCompatibilityResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/compatibility"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListCompatibilityResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.CompatibilityResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.CompatibilityResponse = out
		}
	}

	return res, nil
}
