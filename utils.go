package main

import (
	"encoding/json"
	"errors"
	"strconv"

	"samhofi.us/x/keybase/types/chat1"
)

func getExplodingLifetimeSeconds(m chat1.MsgSummary) (int64, error) {
	if m.IsEphemeral {
		explodingString := strconv.FormatInt(int64(m.ETime), 10)
		explodingTruncated := explodingString[:10]
		resultLifetime, err := strconv.ParseInt(explodingTruncated, 10, 64)
		if err != nil {
			return 0, err
		}
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
