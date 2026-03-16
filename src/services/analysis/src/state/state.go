package state

// A function that takes the new state as an argument and returns a cleanup function
type ApplicationStateInterface[T any] interface {
	OnChange(func()) func()
	Set(T)
	Value() T
}

type ApplicationState[T any] struct {
	value     T
	listeners map[*func()]struct{}
}

func (applicationState *ApplicationState[T]) Set(newValue T) { // TODO make all struct methods, everywhere, take pointers
	applicationState.value = newValue
	for listener := range applicationState.listeners {
		go (*listener)() // TODO is there a reason I might regret multiple listeners firing at once?
	}
}

func (applicationState ApplicationState[T]) Value() T {
	return applicationState.value
}

// ///////////////// //
// Application Phase //
// ///////////////// //

type applicationPhaseType int

const (
	WaitingForDirectory applicationPhaseType = iota
	Ingesting
	WaitingForReport
	GeneratingReport
)

var ApplicationPhase = ApplicationState[applicationPhaseType]{
	value: WaitingForDirectory,
}
