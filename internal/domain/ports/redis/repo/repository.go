package repo

type Repository interface {
	SendStreamLog(message string)
}
