package lib

import "context"

type MonitorControl interface {
	GetDisplays(ctx context.Context, display <-chan Display) error
	SetBrightness(ctx context.Context, display Display, brightness int) (bool, error)
	GetCapabilities(ctx context.Context, display Display) ([]Capability, error)
}

type Display struct {
	Description string
	internal    interface{}
}

type Capability struct {
	Name string
}
