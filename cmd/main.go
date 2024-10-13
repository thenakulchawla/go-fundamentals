package main

import (
	"github.com/alecthomas/kong"
	"github.com/rs/zerolog/log"
	deferex "github.com/thenakulchawla/go-fundamentals/internal/defer"
	egex "github.com/thenakulchawla/go-fundamentals/internal/errgroup"
	wgex "github.com/thenakulchawla/go-fundamentals/internal/waitgroup"
)

var CLI struct {
	Defer   DeferCmd `cmd:"defer" help:"Learn about defer"`
	ErrGrp  EGCmd    `cmd:"errgrp" help:"Learn about error groups"`
	WaitGrp WGCmd    `cmd:"waitgrp" help:"Learn about error groups"`
}

type DeferCmd struct{}

func (d *DeferCmd) Run() error {
	return deferex.RunExamples()
}

type EGCmd struct{}

func (e *EGCmd) Run() error {
	return egex.RunExamples()
}

type WGCmd struct{}

func (w *WGCmd) Run() error {
	return wgex.RunExamples()
}

func main() {
	log.Info().Msg("starting examples")

	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
