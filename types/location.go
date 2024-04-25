package types

// Location - is implement https://core.telegram.org/bots/api#location
type Location struct {
	Latitude             float64  `json:"latitude"`                         // Latitude as defined by sender
	Longitude            float64  `json:"longitude"`                        // Longitude as defined by sender
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`    // The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           *int64   `json:"live_period,omitempty"`            // Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	Heading              *int64   `json:"heading,omitempty"`                // The direction in which user is moving, in degrees; 1-360. For active live locations only.
	ProximityAlertRadius *int64   `json:"proximity_alert_radius,omitempty"` // The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}
