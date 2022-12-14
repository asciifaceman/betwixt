package lifecycle

import (
	"fmt"

	"github.com/asciifaceman/betwixt/betwixt/state"
)

type Lifecycler int

const (
	None Lifecycler = iota
	AWSLifecycle
)

func StringToLifecycle(lc string) Lifecycler {
	switch lc {
	case "aws":
		return AWSLifecycle
	default:
		return None
	}
}

func (l Lifecycler) StateManager(wd string) (state.State, error) {
	switch l {
	case None:
		return nil, fmt.Errorf("no lifecycle")
	case AWSLifecycle:
		stater := &state.AwsState{}
		err := stater.Init(wd)
		return stater, err
	default:
		return nil, fmt.Errorf("no lifecycle")
	}
}

func (l Lifecycler) New() Lifecycle {
	switch l {
	case None:
		return nil
	default:
		return nil
	case AWSLifecycle:
		return &AWSRemote{}
	}
}

type Lifecycle interface {
	LoadConfig(interface{}) error
	Init() error
	Launch() error
	//Destroy() error
}

func NewLifecycle(opt Lifecycler) Lifecycle {
	switch opt {
	case None:
		return nil
	default:
		return nil
	}
}
