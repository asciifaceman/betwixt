package state

// State interface defines the actions for managing a state file
// relating to betwixt instance tracking
type State interface {
	Load() error
	Write() error
	GetID() string
	SetID(string) error
}
