package main

import (
	"log"

	"samhofi.us/x/keybase"
	"samhofi.us/x/keybase/types/chat1"
)

// RegisterHandlers is called by main to map these handler funcs to events
func (b *bot) registerHandlers() {
	chat := b.chatHandler
	err := b.errHandler

	b.handlers = keybase.Handlers{
		ChatHandler:  &chat,
		ErrorHandler: &err,
	}
}

// chatHandler should handle all messages coming from the chat
func (b *bot) chatHandler(m chat1.MsgSummary) {
	if m.IsEphemeral {
		explodingLifetime, err := getExplodingLifetimeSeconds(m)
		if err != nil {
			log.Println(err)
			return
		}
		b.debug("Ephemeral message recieved with %ds lifetime", explodingLifetime)
		if explodingLifetime < b.minAllowedLifetime || explodingLifetime > b.maxAllowedLifetime {
			b.debug("message exploded")
			b.k.DeleteByConvID(m.ConvID, m.Id)
		}
	}
}

// this handles all errors returned from the keybase binary
func (b *bot) errHandler(m error) {
	log.Println("---[ error ]---")
	log.Println(p(m))
}
