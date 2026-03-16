package state

type applicationPhaseType int

const (
	WaitingForDirectory applicationPhaseType = iota
	Ingesting
	WaitingForReport
	GeneratingReport
)

var ApplicationPhase = WaitingForDirectory
