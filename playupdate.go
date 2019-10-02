package playupdate

type UpdateService interface {
	Update(string) (*UpdateProfile, error)
}

type Application struct {
	ApplicationID string `json:"applicationId"`
	Version       string `json:"version"`
}

type Profile struct {
	Applications []Application `json:"applications`
}

type UpdateProfile struct {
	Profile Profile `json:"profile"`
}
