package constanct

type StatusTime int64

const (
	Second  StatusTime = 1
	Minutes StatusTime = 60
	Hour    StatusTime = 60 * 60
	DAY     StatusTime = 60 * 60 * 60
	Momth   StatusTime = 30 * 60 * 60
	Year    StatusTime = 365 * 60 * 60
)
