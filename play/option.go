package play

import "strconv"

type Option func(play *play) error

func StartTime(startTime int) Option {
	return func(play *play) error {
		play.StartTime = strconv.Itoa(startTime)
		return nil
	}
}

func EndTime(endTime int) Option {
	return func(play *play) error {
		play.EndTime = strconv.Itoa(endTime)
		return nil
	}
}

func Volume(volume int) Option {
	return func(play *play) error {
		play.Volume = strconv.Itoa(volume)
		return nil
	}
}

func NoReplace(noReplace bool) Option {
	return func(play *play) error {
		play.NoReplace = noReplace
		return nil
	}
}

func Pause(pause bool) Option {
	return func(play *play) error {
		play.Pause = pause
		return nil
	}
}
