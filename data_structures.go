package main

type Time interface {
	Duration() int
	IsPositive() bool
	Creator() string
}

type positiveTime struct {
	duration int
	creator  string
}

type negativeTime struct {
	duration int
	creator  string
}

func (time positiveTime) Duration() int {
	return time.duration
}

func (time negativeTime) Duration() int {
	return time.duration
}

func (time positiveTime) IsPositive() bool {
	return true
}

func (time negativeTime) IsPositive() bool {
	return false
}

func (time positiveTime) Creator() string {
	return time.creator
}

func (time negativeTime) Creator() string {
	return time.creator
}
