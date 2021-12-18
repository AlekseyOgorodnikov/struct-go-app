package ports

type DbPort interface {
	CloceDcConnection()
	AddToHistory(answer int32, operation string) error
}
