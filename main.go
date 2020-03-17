package main

import (
	"log"
	"os"

	"samhofi.us/x/keybase"
	"samhofi.us/x/keybase/types/chat1"
)

// Bot holds the necessary information for the bot to work.
type bot struct {
	k        *keybase.Keybase
	handlers keybase.Handlers
	opts     keybase.RunOptions
	config   botConfig
}

type botConfig struct {
	debug              bool
	minAllowedLifetime int64
	maxAllowedLifetime int64
	enabledTeams       []chat1.ChatChannel
}

// newBot returns a new empty bot
func newBot() *bot {
	var b bot
	b.k = keybase.NewKeybase()
	b.handlers = keybase.Handlers{}
	b.opts = keybase.RunOptions{}
	b.config.enabledTeams = []chat1.ChatChannel{}
	return &b
}

// Debug provides printing only when --debug flag is set or BOT_DEBUG env var is set
func (b *bot) debug(s string, a ...interface{}) {
	if b.config.debug {
		log.Printf(s, a...)
	}
}

// setOptions applies filter channels, if they are provided
func (b *bot) setOptions() {
	if len(b.config.enabledTeams) > 0 {
		b.opts = keybase.RunOptions{
			FilterChannels: b.config.enabledTeams,
		}
	}
}

// run performs a proxy main function
func (b *bot) run(args []string) error {
	// parse the arguments
	err := b.parseArgs(args)
	if err != nil {
		return err
	}

	b.setOptions()
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
