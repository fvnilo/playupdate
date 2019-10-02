package http

import "github.com/nylo-andry/playupdate"

func NewUpdateProfile() *playupdate.UpdateProfile {
	return &playupdate.UpdateProfile{
		Profile: playupdate.Profile{
			Applications: []playupdate.Application{
				playupdate.Application{
					ApplicationID: "music_app",
					Version:       "v1.4.10",
				},
				playupdate.Application{
					ApplicationID: "diagnostic_app",
					Version:       "v1.2.6",
				},
				playupdate.Application{
					ApplicationID: "settings_app",
					Version:       "v1.1.5",
				},
			},
		},
	}
}
