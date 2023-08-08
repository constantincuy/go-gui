package common

type Component interface {
	Core() *Core
	Mount()
	Update()
	Destroy()
}
