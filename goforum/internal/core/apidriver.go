package core

type ApiDriver interface {
	Init() error
	Run() error
}
