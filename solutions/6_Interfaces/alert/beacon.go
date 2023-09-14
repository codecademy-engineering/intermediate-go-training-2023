package alert

import (
	"github.com/codecademy-engineering/intermediate-go-training-2023/models"
)

type Beacon interface {
	AlertAllAgents(sq *models.Squirrel) error
}

type mockBeacon struct{}

func NewMockBeacon() *mockBeacon {
	return &mockBeacon{}
}

func (m *mockBeacon) AlertAllAgents(sq *models.Squirrel) error {
	return nil
}
