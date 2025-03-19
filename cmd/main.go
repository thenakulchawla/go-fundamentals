package main

import (
	"context"

	"github.com/alecthomas/kong"
	"github.com/thenakulchawla/go-fundamentals/internal/channels"
	deferex "github.com/thenakulchawla/go-fundamentals/internal/defer"
	egex "github.com/thenakulchawla/go-fundamentals/internal/errgroup"
	"github.com/thenakulchawla/go-fundamentals/internal/producer"
	wgex "github.com/thenakulchawla/go-fundamentals/internal/waitgroup"
	"github.com/thenakulchawla/parchment"
)

var CLI struct {
	Defer    DeferCmd    `cmd:"defer" help:"Learn about defer"`
	ErrGrp   EGCmd       `cmd:"errgrp" help:"Learn about error groups"`
	WaitGrp  WGCmd       `cmd:"waitgrp" help:"Learn about wait groups"`
	Channels ChanCmd     `cmd:"ch" help:"Learn about channels"`
	Producer ProducerCmd `cmd:"produce" help:"Run producer and consumer"`
}

type DeferCmd struct{}

func (d *DeferCmd) Run() error {
	ctx := context.Background()
	ctx = parchment.New(ctx)
	return deferex.RunExamples(ctx)
}

type EGCmd struct{}

func (e *EGCmd) Run() error {
	ctx := context.Background()
	ctx = parchment.New(ctx)
	return egex.RunExamples(ctx)
}

type WGCmd struct{}

func (w *WGCmd) Run() error {
	ctx := context.Background()
	ctx = parchment.New(ctx)
	return wgex.RunExamples(ctx)
}

type ChanCmd struct{}

func (c *ChanCmd) Run() error {
	ctx := context.Background()
	ctx = parchment.New(ctx)
	return channels.RunExamples(ctx)
}

type ProducerCmd struct{}

func (p *ProducerCmd) Run() error {
	ctx := context.Background()
	ctx = parchment.New(ctx)
	return producer.RunExamples(ctx)
}

func main() {

	cmd := kong.Parse(&CLI)
	err := cmd.Run()
	cmd.FatalIfErrorf(err)
}
