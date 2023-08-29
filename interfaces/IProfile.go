// this interface will provide the requried methods
package interfaces

import "grpc-mongo/models"

type IProfile interface {
	CreateProfile(profile *models.Profile) (*models.DBResponse, error)
	EditProfile(profile *models.Profile)
	DeleteProfile(profileId int)
	GetProfileById(profile int)
	GetProfileBySearch()
}
