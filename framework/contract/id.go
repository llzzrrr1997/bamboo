package contract

const IDKey = "bamboo:id"

type IDService interface {
	NewID() string
}
