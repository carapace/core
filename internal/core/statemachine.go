package core

import (
	"github.com/carapace/core/api/v1/proto/generated"
)

type StateMachine interface {
	Get(version, kind, name string) State
	Generate(config v1.Config) State
}

type State interface {
	Hash() float64
	Next() State
	Prev() State
	Current() State
}
