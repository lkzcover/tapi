package types

// WebAppInfo - implement https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	URL string `json:"url"` // An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
}
