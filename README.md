[![Build](https://github.com/haukened/boom-bot/workflows/Build/badge.svg)](https://github.com/haukened/boom-bot/actions?query=workflow%3ABuild)
# boom-bot
A bot for Keybase users who hate exploding (ephemeral) messages.

This package requires the keybase binary installed on your system, and works on linux, macOS, and Windows 10

#### Tested on:
 - Ubuntu Latest
 - macOS Latest
 - Windows Latest

## Running on the command line:
#### Installation:
`go get -u github.com/haukened/boom-bot`

`go install github.com/haukened/boom-bot`
#### Running:
```
  -debug
        enables command debugging
  -max-lifetime-sec int
        sets the maximum exploding lifetime (0 to 604800 seconds, default 604800)
  -min-lifetime-sec int
        sets the minimum exploding lifetime in seconds (0 to 604800 seconds, default 0)
  -teams string
        comma separated list of teams the bot will listen to (user must be a member)
        if not set the program will listen to every team you are a member of,
        and explode all messages in teams you have permissions.
```

#### Example: 
`boom-bot --debug --min-lifetime-sec 100 --max-lifetime-sec 500 --teams keybasefriends,mkbot,kbtui`

## Running in the docker container (coming soon):
You need to set ENV vars instead of passing command line flags:

Required by keybase: (Must set all of these)
 - `KEYBASE_USERNAME=foo`
 - `KEYBASE_PAPERKEY="bar baz ..."`
 - `KEYBASE_SERVICE=1`
 
Required by this package: (Set the values you feel like, if you don't set them they won't be used)
 - `BOT_DEBUG=true`
 - `BOT_MIN_LIFETIME_SEC=100`
 - `BOT_MAX_LIFETIME_SEC=500`
 - `BOT_TEAMS="keybasefriends,mkbot,kbtui`
