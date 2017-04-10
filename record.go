package main

import "time"

type CallRecord struct {
	Id                 int           `json:"id"`
	StartTime          time.Time     `json:"startTime"`
	EndTime            time.Time     `json:"endTime"`
	Duration           time.Duration `json:"duration"`
	CallerLineIdentity string        `json:"callerLineIdentity"`
	NonChargedParty    string        `json:"nonChargedParty"`
}
