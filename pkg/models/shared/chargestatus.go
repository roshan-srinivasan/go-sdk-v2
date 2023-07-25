// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ChargeStatusState string

const (
	ChargeStatusStateCharging     ChargeStatusState = "CHARGING"
	ChargeStatusStateFullyCharged ChargeStatusState = "FULLY_CHARGED"
	ChargeStatusStateNotCharging  ChargeStatusState = "NOT_CHARGING"
)

func (e ChargeStatusState) ToPointer() *ChargeStatusState {
	return &e
}

func (e *ChargeStatusState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CHARGING":
		fallthrough
	case "FULLY_CHARGED":
		fallthrough
	case "NOT_CHARGING":
		*e = ChargeStatusState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChargeStatusState: %v", v)
	}
}

type ChargeStatus struct {
	// Indicates whether a charging cable is currently plugged into the vehicle’s charge port.
	IsPluggedIn *bool              `json:"isPluggedIn,omitempty"`
	State       *ChargeStatusState `json:"state,omitempty"`
}

func (o *ChargeStatus) GetIsPluggedIn() *bool {
	if o == nil {
		return nil
	}
	return o.IsPluggedIn
}

func (o *ChargeStatus) GetState() *ChargeStatusState {
	if o == nil {
		return nil
	}
	return o.State
}
