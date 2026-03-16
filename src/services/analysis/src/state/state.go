package state

// A function that takes the new state as an argument and returns a cleanup function
type ApplicationStateInterface[T comparable] interface {
	OnChange(func()) func()
	Set(T)
	Value() T
}

type ApplicationState[T comparable] struct {
	value     T
	listeners map[*func()]struct{}
}

func (applicationState *ApplicationState[T]) Set(newValue T) { // TODO make all struct methods, everywhere, take pointers
	oldValue := applicationState.value
	if oldValue == newValue {
		return
	}

	applicationState.value = newValue

	for listener := range applicationState.listeners {
		(*listener)() // TODO is there a reason I might regret multiple listeners firing at once?
	}
}

func (applicationState *ApplicationState[T]) Value() T {
	return applicationState.value
}

func (applicationState *ApplicationState[T]) OnChange(callback func()) func() {
	if applicationState.listeners == nil {
		applicationState.listeners = make(map[*func()]struct{})
	}
	// add to listeners list
	applicationState.listeners[&callback] = struct{}{}

	// cleanup function
	cleanup := func() {
		delete(applicationState.listeners, &callback)
	}
	return cleanup
}

// ///////////////// //
// Application Phase //
// ///////////////// //

type applicationPhaseType string

const (
	WaitingForDirectory applicationPhaseType = "WaitingForDirectory"
	Ingesting           applicationPhaseType = "Ingesting"
	WaitingForReport    applicationPhaseType = "WaitingForReport"
	GeneratingReport    applicationPhaseType = "GeneratingReport"
)

var ApplicationPhase = ApplicationState[applicationPhaseType]{
	value: WaitingForDirectory,
}
