package models

import "time"

type Request struct {
	URL         string        `json:"url" binding:"required"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type Response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"x-rate-limit"`
	XRateLimitReset time.Duration `json:"rate-limit-reset"`
}
