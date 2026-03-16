package state

type applicationPhaseType int

const (
	WaitingForDirectory applicationPhaseType = iota
	Ingesting
	WaitingForReport
	GeneratingReport
)

var applicationPhase = WaitingForDirectory

func GetApplicationPhase() applicationPhaseType {
	return applicationPhase
}

func SetApplicationPhase(newPhase applicationPhaseType) {
	applicationPhase = newPhase
}
