package main

import "flag"

type cmdArgs struct {
	port     int
	botToken string
	ownerId  int64
}

func getArgs() *cmdArgs {
	var output cmdArgs

	flag.IntVar(&output.port, "port", 9000, "port to serve")
	flag.StringVar(&output.botToken, "bot_token", "", "telegram bot token")
	flag.Int64Var(&output.ownerId, "owner_id", 0, "owner telegram id")
	flag.Parse()
	return &output
}
