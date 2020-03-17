package main

import (
	"flag"
	"strings"

	"github.com/caarlos0/env"
	"samhofi.us/x/keybase/types/chat1"
)

// parseArgs parses command line and environment args and sets globals
func (b *bot) parseArgs(args []string) error {
	// parse the env variables into the bot config
	if err := env.Parse(&b.config); err != nil {
		return err
	}

	// then parse CLI args as overrides
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	cliConfig := botConfig{}
	flags.Int64Var(&cliConfig.MinAllowedLifetime, "min-lifetime-sec", -1, "sets the minimum exploding lifetime in seconds")
	flags.Int64Var(&cliConfig.MaxAllowedLifetime, "max-lifetime-sec", -1, "sets the maximum exploding lifetime")
	flags.BoolVar(&cliConfig.Debug, "debug", false, "enables command debugging")
	flags.StringVar(&cliConfig.EnabledTeams, "teams", "", "comma separated list of teams the bot will listen to (user must be a member)")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	// then override the environment vars if there were cli args
	if flags.NFlag() > 0 {
		if cliConfig.MinAllowedLifetime > 0 {
			b.config.MinAllowedLifetime = cliConfig.MinAllowedLifetime
		}
		if cliConfig.MaxAllowedLifetime > 0 {
			b.config.MaxAllowedLifetime = cliConfig.MaxAllowedLifetime
		}
		if cliConfig.Debug == true {
			b.config.Debug = true
		}
		if cliConfig.EnabledTeams != "" {
			b.config.EnabledTeams = cliConfig.EnabledTeams
		}

	}

	// then print the running options
	b.debug("Debug Enabled")
	if b.config.MaxAllowedLifetime < 604800 {
		b.debug("Exploding maximum allowed message life set to %d", b.config.MaxAllowedLifetime)
	}
	if b.config.MinAllowedLifetime > 0 {
		b.debug("Exploding minimum allowed message life set to %d", b.config.MinAllowedLifetime)
	}
	if b.config.EnabledTeams != "" {
		b.debug("Listening to teams: %s", b.config.EnabledTeams)
	} else {
		b.debug("no team filter provided, listening to all teams...")
	}

	return nil
}

func parseBotTeams(input string) []chat1.ChatChannel {
	fields := strings.Split(input, ",")
	var result []chat1.ChatChannel
	for _, team := range fields {
		result = append(result, chat1.ChatChannel{
			Name:        team,
			MembersType: "team",
		})
	}
	return result
}
