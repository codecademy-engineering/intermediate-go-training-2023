package sos

import (
	"github.com/codecademy-engineering/intermediate-go-training-2023/models"
)

type Beacon interface {
	AlertAllAgents(sq *models.Squirrel) error
}

type mockBeacon struct {
	err error
}

func NewMockBeacon(err error) *mockBeacon {
	return &mockBeacon{
		err: err,
	}
}

func (m *mockBeacon) AlertAllAgents(sq *models.Squirrel) error {
	return m.err
}
