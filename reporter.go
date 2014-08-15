package main

type Reporter interface {
	Status(path string, status GitStatus) error
}
