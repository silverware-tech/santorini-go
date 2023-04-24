package utility

const X_SIZE = 5
const Y_SIZE = 5
const MIN_PLAYERS = 2
const MAX_PLAYERS = 4
const MAX_CHARACTERS_PER_PLAYER = 2
const MAX_CHARACTERS = MAX_PLAYERS * MAX_CHARACTERS_PER_PLAYER

var COLORS = map[uint8]string{
	0: "\x1b[1;31m",
	1: "\x1b[1;34m",
	2: "\x1b[1;32m",
	3: "\x1b[1;33m",
}
