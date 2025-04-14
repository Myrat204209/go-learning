package main

import "fmt"

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

func (e email) cost() int {
	if !e.isSubscribed {
		return 5 * len(e.body)
	}
	return 2 * len(e.body)
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	}
	return fmt.Sprintf("'%s' | Not Subscribed", e.body)
}

type email struct {
	isSubscribed bool
	body         string
}
