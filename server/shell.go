package main

import context "context"

type ShellService struct{}

func (c ShellService) Run(ctx context.Context, in *Msg) (*Std, error) {
	return &Std{
		Stdout: "",
		Stderr: "",
	}, nil
}
