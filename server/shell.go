package main

import (
	"bytes"
	"log"
	"os/exec"
)

type ShellService struct{}

func (ShellService) Run(s ShellService_RunServer) error {
	log.Println("func Run:")
	req, err := s.Recv()
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(s.Context(), req.GetProg(), req.GetArgs()...)
	log.Printf("cmd: %v \n", cmd.Args)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return err
	}
	s.Send(&Output{
		Stdout: outb.Bytes(),
		Stderr: outb.Bytes(),
	})
	return err

}

func (ShellService) CMsg(c ShellService_CMsgServer) error {
	return nil
}

func (ShellService) SMsg(e *CmdRequest, s ShellService_SMsgServer) error {
	return nil
}
