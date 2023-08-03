package component

type Component interface {
	Core() *Core
	Mount()
	Update()
	Destroy()
}
