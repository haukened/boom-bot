package main

import (
	"flag"
	"log"
	"os"
	"strconv"

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

// parseArgs parses command line and environment args and sets globals
func (b *bot) parseArgs(args []string) error {
	// first check for command line flags
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	// default 0 seconds (the keybase min)
	flags.Int64Var(&b.minAllowedLifetime, "min-lifetime-sec", 0, "sets the minimum exploding lifetime in seconds")
	// default 7 days (the keybase max)
	flags.Int64Var(&b.maxAllowedLifetime, "max-lifetime-sec", 604800, "sets the maximum exploding lifetime")
	flags.BoolVar(&b.debugEnabled, "debug", false, "enables command debugging")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	// then check the env vars
	envDebug := os.Getenv("BOT_DEBUG")
	if envDebug != "" {
		ret, err := strconv.ParseBool(envDebug)
		if err != nil {
			return err
		}

		// if flag was false but env is true, set debug
		if b.debugEnabled == false && ret == true {
			b.debugEnabled = true
		}
	}
	if b.debugEnabled {
		log.Println("Debugging enabled.")
	}

	envMinLifetime := os.Getenv("MIN_LIFETIME_SEC")
	if envMinLifetime != "" {
		ret, err := strconv.ParseInt(envMinLifetime, 10, 64)
		if err != nil {
			return err
		}

		// if flag was not set but env was
		if ret > b.minAllowedLifetime {
			b.minAllowedLifetime = ret
		}
	}
	if b.debugEnabled {
		log.Printf("exploding minimum lifetime set to %d\n", b.minAllowedLifetime)
	}

	envMaxLifetime := os.Getenv("MAX_LIFETIME_SEC")
	if envMaxLifetime != "" {
		ret, err := strconv.ParseInt(envMaxLifetime, 10, 64)
		if err != nil {
			return err
		}

		// if flag was not set but env was
		if ret < b.maxAllowedLifetime {
			b.maxAllowedLifetime = ret
		}
	}
	if b.debugEnabled {
		log.Printf("exploding maximum lifetime set to %d\n", b.maxAllowedLifetime)
	}

	return nil
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
