package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Accident struct {
	gorm.Model
	AccidentId       string `gorm:"unique;not null"`
	AccidentTime     time.Time
	PersonInChange   string
	Participant      datatypes.JSON //list of participant
	Bug              bool
	Recover          bool
	Level            string
	RecoverMethod    string
	AccidentDesc     string
	ScopeOfInfluence string
	Review           bool
	ReviewTime       time.Time
	Cause            datatypes.JSONType[Cause] //struct
	Rectification    datatypes.JSONType[Rectification]
}

type Cause struct {
	DirectCause string
	RootCause   string
}

type Rectification struct {
	Technology   []string
	Officers     []string
	Organization []string
}

func (m *Accident) TableName() string {
	return "accident"
}
