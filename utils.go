package main

import (
	"encoding/json"
	"errors"
	"strconv"

	"samhofi.us/x/keybase/types/chat1"
)

// this takes a chat message and calculates the exploding lifetime
// the way keybase sends it is.... strange
func getExplodingLifetimeSeconds(m chat1.MsgSummary) (int64, error) {
	if m.IsEphemeral {
		// convert the int64 to a string so we can grab the first 10 digits
		explodingString := strconv.FormatInt(int64(m.ETime), 10)
		// truncate the string to 10 digits
		explodingTruncated := explodingString[:10]
		// then re-parse to an int64
		resultLifetime, err := strconv.ParseInt(explodingTruncated, 10, 64)
		if err != nil {
			return 0, err
		}
		// subtract the sent time from the exploding time to get the message life
		originalLifetime := resultLifetime - m.SentAt
		return originalLifetime, nil
	}
	return 0, errors.New("Unable to parse exploding lifetime on non-ephemeral message")
}

// this JSON pretty prints errors and debug
func p(b interface{}) string {
	s, _ := json.MarshalIndent(b, "", "  ")
	return string(s)
}
