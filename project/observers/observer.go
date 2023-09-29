package observers

type Observer interface {
	Update(event string, data interface{})
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(event string, data interface{})
}
