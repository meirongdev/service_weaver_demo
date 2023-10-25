package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/metrics"
)

var metricsCounter = metrics.NewCounter("calls", "Number of calls")

type Reverser interface {
	Reverse(ctx context.Context, str string) (string, error)
}

type reverser struct {
	weaver.Implements[Reverser]
}

func (r *reverser) Reverse(ctx context.Context, str string) (string, error) {
	defer metricsCounter.Inc()
	r.Logger(ctx).Debug("reversing", "str", str)
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes), nil
}
