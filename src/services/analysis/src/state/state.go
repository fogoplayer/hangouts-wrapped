package state

import "zarinloosli.com/hangouts-wrapped/util"

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

// /////////////////// //
// SetApplicationState //
// /////////////////// //

// TODO is there a way to keep users from calling state.Value().Add? Or state.Set()

type SetApplicationState[T comparable] struct {
	ApplicationState[*util.Set[T]]
}

func (applicationState *SetApplicationState[T]) getInternalSet() *util.Set[T] {
	set := applicationState.getInternalSet()
	if set == nil {
		set = &util.Set[T]{}
		applicationState.value = set // do not update listeners for initializing internal data structure
	}
	return set
}

func (applicationState *SetApplicationState[T]) Add(newValue T) {
	set := applicationState.getInternalSet()
	exists := set.Includes(newValue)
	if exists {
		return
	}

	set.Add(newValue)

	applicationState.updateListeners()
}

func (applicationState *SetApplicationState[T]) Delete(newValue T) {
	set := applicationState.value
	exists := set.Includes(newValue)
	if !exists {
		return
	}

	set.Delete(newValue)

	applicationState.updateListeners()
}

func (applicationState *SetApplicationState[T]) updateListeners() {
	for listener := range applicationState.listeners {
		(*listener)() // TODO is there a reason I might regret multiple listeners firing at once?
	}
}

func (applicationState *SetApplicationState[T]) Value() util.Set[T] {
	return *applicationState.value
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
