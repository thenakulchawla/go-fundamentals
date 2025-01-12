package main

import (
	"github.com/alecthomas/kong"
	"github.com/rs/zerolog/log"
	"github.com/thenakulchawla/go-fundamentals/internal/channels"
	deferex "github.com/thenakulchawla/go-fundamentals/internal/defer"
	egex "github.com/thenakulchawla/go-fundamentals/internal/errgroup"
	wgex "github.com/thenakulchawla/go-fundamentals/internal/waitgroup"
)

var CLI struct {
	Defer    DeferCmd `cmd:"defer" help:"Learn about defer"`
	ErrGrp   EGCmd    `cmd:"errgrp" help:"Learn about error groups"`
	WaitGrp  WGCmd    `cmd:"waitgrp" help:"Learn about wait groups"`
	Channels ChanCmd  `cmd:"ch" help:"Learn about channels"`
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

type ChanCmd struct{}

func (c *ChanCmd) Run() error {
	return channels.RunExamples()
}

func main() {
	log.Info().Msg("starting examples")

	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
