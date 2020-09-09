package grabble

import (
	"flag"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var log = logrus.New()
var logPath = flag.String("logfile", "/tmp/grabble.log", "provide path for log file")
var logLevel = flag.String("loglevel", "ERROR", "provide log level")

func init() {

	// log to console and file
	f, err := os.OpenFile(*logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)

	level, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})
	log.SetOutput(wrt)
}

// CreateDefaultGame - it creates and return basic Grabble game.
func CreateDefaultGame(players []string) Grabble {
	officialScrabbleBoard := [15][15]rune{
		{'W', rune(0), rune(0), 'l', rune(0), rune(0), rune(0), 'W', rune(0), rune(0), rune(0), 'l', rune(0), rune(0), 'W'},
		{rune(0), 'w', rune(0), rune(0), rune(0), 'L', rune(0), rune(0), rune(0), 'L', rune(0), rune(0), rune(0), 'w', rune(0)},
		{rune(0), rune(0), 'w', rune(0), rune(0), rune(0), 'l', rune(0), 'l', rune(0), rune(0), rune(0), 'w', rune(0), rune(0)}, // 3
		{'l', rune(0), rune(0), 'w', rune(0), rune(0), rune(0), 'l', rune(0), rune(0), rune(0), 'w', rune(0), rune(0), 'l'},
		{rune(0), rune(0), rune(0), rune(0), 'w', rune(0), rune(0), rune(0), rune(0), rune(0), 'w', rune(0), rune(0), rune(0), rune(0)},
		{rune(0), 'L', rune(0), rune(0), rune(0), 'L', rune(0), rune(0), rune(0), 'L', rune(0), rune(0), rune(0), 'L', rune(0)}, // 6
		{rune(0), rune(0), 'l', rune(0), rune(0), rune(0), 'l', rune(0), 'l', rune(0), rune(0), rune(0), 'l', rune(0), rune(0)},
		{'W', rune(0), rune(0), 'l', rune(0), rune(0), rune(0), 's', rune(0), rune(0), rune(0), 'l', rune(0), rune(0), 'W'},
		{rune(0), rune(0), 'l', rune(0), rune(0), rune(0), 'l', rune(0), 'l', rune(0), rune(0), rune(0), 'l', rune(0), rune(0)}, // 9
		{rune(0), 'L', rune(0), rune(0), rune(0), 'L', rune(0), rune(0), rune(0), 'L', rune(0), rune(0), rune(0), 'L', rune(0)},
		{rune(0), rune(0), rune(0), rune(0), 'w', rune(0), rune(0), rune(0), rune(0), rune(0), 'w', rune(0), rune(0), rune(0), rune(0)},
		{'l', rune(0), rune(0), 'w', rune(0), rune(0), rune(0), 'l', rune(0), rune(0), rune(0), 'w', rune(0), rune(0), 'l'}, // 12
		{rune(0), rune(0), 'w', rune(0), rune(0), rune(0), 'l', rune(0), 'l', rune(0), rune(0), rune(0), 'w', rune(0), rune(0)},
		{rune(0), 'w', rune(0), rune(0), rune(0), 'L', rune(0), rune(0), rune(0), 'L', rune(0), rune(0), rune(0), 'w', rune(0)},
		{'W', rune(0), rune(0), 'l', rune(0), rune(0), rune(0), 'W', rune(0), rune(0), rune(0), 'l', rune(0), rune(0), 'W'}, // 15
	}
	officialDict := "../fixtures/collins_official_scrabble_2019.txt"
	tilePoints := map[rune]int{
		'_': 0,
		'E': 1, 'A': 1, 'I': 1, 'O': 1, 'N': 1, 'R': 1, 'T': 1, 'L': 1, 'S': 1, 'U': 1,
		'D': 2, 'G': 2,
		'B': 3, 'C': 3, 'M': 3, 'P': 3,
		'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
		'K': 5,
		'J': 8, 'X': 8,
		'Q': 10, 'Z': 10,
	}
	// FIXME: I removed __ from all tiles to test stuff. Blanks will require rework in many functions.
	allTiles := []rune("EEEEEEEEEEEEAAAAAAAAAIIIIIIIIIOOOOOOOONNNNNNRRRRRRTTTTTTLLLLSSSSUUUUDDDDGGGBBCCMMPPFFHHVVWWYYKJXQZ")

	return CreateGrabble(officialDict, officialScrabbleBoard, players, allTiles, tilePoints, 7)
}
