package version

type VersionService interface {
	GetAppVersion() string
}

type VersionServiceImpl struct {
}

func NewService() VersionService {
	return &VersionServiceImpl{}
}

func (vs *VersionServiceImpl) GetAppVersion() string {
	return "To be done efficiently"
}
