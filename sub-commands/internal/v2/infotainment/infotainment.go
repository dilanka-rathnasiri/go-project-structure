package infotainment

type InfotainmentSystemV2 interface {
	DisplayNavigation(destination string) (string, error)
}
