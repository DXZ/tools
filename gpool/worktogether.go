package gpool

import (
	"fmt"
)

type JobOld struct {
	JobNo   int
	resutls chan<- Result
}

type Result struct {
	JobNo  int
	Answer int
}
