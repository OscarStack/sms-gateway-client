package models

type Filter string

var (
	ALL_MESSAGES  Filter = "all"
	READ_MESSAGES Filter = "read"
	NEW_MESSAGES  Filter = "new"
)
