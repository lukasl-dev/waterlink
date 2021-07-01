package loadtrack

type TrackLoader interface {
	LoadTrack(identifier string) (*Response, error)
}
