package main

import (
	"log"
	"os"

	"samhofi.us/x/keybase"
)

// Bot holds the necessary information for the bot to work.
type bot struct {
	k                  *keybase.Keybase
	handlers           keybase.Handlers
	opts               keybase.RunOptions
	debugEnabled       bool
	minAllowedLifetime int64
	maxAllowedLifetime int64
}

// newBot returns a new empty bot
func newBot() *bot {
	var b bot
	b.k = keybase.NewKeybase()
	b.handlers = keybase.Handlers{}
	b.opts = keybase.RunOptions{}
	return &b
}

// Debug provides printing only when --debug flag is set or BOT_DEBUG env var is set
func (b *bot) debug(s string, a ...interface{}) {
	if b.debugEnabled {
		log.Printf(s, a...)
	}
}

// run performs a proxy main function
func (b *bot) run(args []string) error {
	// parse the arguments
	err := b.parseArgs(args)
	if err != nil {
		return err
	}

	b.registerHandlers()

	log.Println("Starting...")
	b.k.Run(b.handlers, &b.opts)
	return nil
}

// main is a thin skeleton, proxied to bot.run()
func main() {
	b := newBot()
	if err := b.run(os.Args); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}
