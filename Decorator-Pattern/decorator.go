package main

type CondimentDecorator interface {
	getDescription() string
	getSize() string
	setSize(string)
}
