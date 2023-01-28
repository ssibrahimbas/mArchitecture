package app

import "github.ssibrahimbas/mArchitecture/auth/cmd/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	Login    command.LoginHandler
	Register command.RegisterHandler
}
