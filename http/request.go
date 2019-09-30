package http

type Application struct {
	ApplicationID string `json:"applicationId"`
	Version       string `json:"version"`
}

type Profile struct {
	Applications []Application `json:"applications`
}

type RequestBody struct {
	Profile Profile `json:"profile"`
}

func NewRequestBody() *RequestBody {
	return &RequestBody{
		Profile: Profile{
			Applications: []Application{
				Application{
					ApplicationID: "music_app",
					Version:       "v1.4.10",
				},
				Application{
					ApplicationID: "diagnostic_app",
					Version:       "v1.2.6",
				},
				Application{
					ApplicationID: "settings_app",
					Version:       "v1.1.5",
				},
			},
		},
	}
}
