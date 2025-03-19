// Package deferex learn all about defer
package deferex

import (
	"context"
	"time"

	"github.com/thenakulchawla/parchment"
)

func RunExamples(ctx context.Context) error {

	ctx = parchment.AddToLogger(ctx, []parchment.LoggerField{
		{Key: "program", Value: "defer"},
	})

	log := parchment.FromContext(ctx)
	log.Info().Msg("running examples")
	basicDefer(ctx)
	orderedDefers(ctx)
	defersWithSleep(ctx)
	deferWithArgs(ctx)

	return nil
}

func basicDefer(ctx context.Context) {
	log := parchment.FromContext(ctx)
	log.Info().Msg("running basic defer")
	defer log.Info().Msg("this is the deferred log")

	log.Info().Msg("this is after the defer statement but will print before")

}

func orderedDefers(ctx context.Context) {
	log := parchment.FromContext(ctx)
	log.Info().Msg("running multiple defers")

	defer log.Info().Msg("exit ordered defers")

	for index := range 3 {
		defer log.Info().Int("index", index).Msg("defer in for loop")
	}

	defer log.Info().Msg("scheduling defers")
}

func defersWithSleep(ctx context.Context) {

	log := parchment.FromContext(ctx)

	log.Info().Msg("defers with sleep")

	defer log.Info().Msg("the last message")

	defer time.Sleep(time.Second * 10)
	log.Info().Msg("first sleep then print the last message")
}

func deferWithArgs(ctx context.Context) {
	log := parchment.FromContext(ctx)
	log.Info().Msg("defer with args")

	defer log.Info().Msg("exiting defer with args")

	a := "hello"
	defer func(s string) {
		log.Info().Str("a_val", a).Msg("inside deferred func")
	}(a)

	a = "world"
	log.Info().Str("a_val", a).Msg("outside of deferred func")
}
