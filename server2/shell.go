package main

import (
	"bytes"
	"log"
	"os/exec"
)

type ShellService struct{}

func (ShellService) Run2(s ShellService_Run2Server) error {
	log.Println("func Run:")
	req, err := s.Recv()
	defer func(e *error) {
		if err != nil {
			log.Println("error[Run]: ", *e)
		}
	}(&err)
	if err != nil {
		return err
	}
	var cmd *exec.Cmd
	if req.GetArgs()[0] != "" {
		cmd = exec.CommandContext(s.Context(), req.GetProg(), req.GetArgs()...)
	} else {
		cmd = exec.CommandContext(s.Context(), req.GetProg())
	}

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
