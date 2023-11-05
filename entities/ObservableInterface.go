package entities

type Observable interface {
	RegisterObserver(o Observer)
	notifyObservers()
}
