package models

import (
	"time"

	"github.com/google/uuid"
)

type Answer struct {
	Id        uuid.UUID
	Answer    string
	TimeStamp time.Time
}

func getAnswer(answer string) Answer {
	return Answer{
		Id:        uuid.UUID{},
		Answer:    answer,
		TimeStamp: time.Now(),
	}
}
