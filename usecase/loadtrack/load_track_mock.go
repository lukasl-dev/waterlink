package loadtrack

import "github.com/stretchr/testify/mock"

type MockedTrackLoader struct {
	mock.Mock
}

func NewMockedTrackLoader() *MockedTrackLoader {
	return new(MockedTrackLoader)
}

func (l *MockedTrackLoader) LoadTrack(identifier string) (*Response, error) {
	args := l.Called(identifier)
	return args.Get(0).(*Response), args.Error(1)
}
