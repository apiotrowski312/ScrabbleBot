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

var letterRatio = map[rune]float64{
	'_': normalizeRatio(31.5, 31.5, -4),
	'S': normalizeRatio(12, 31.5, -4),
	'Z': normalizeRatio(8, 31.5, -4),
	'X': normalizeRatio(6.3, 31.5, -4),
	'E': normalizeRatio(4, 31.5, -4),
	'H': normalizeRatio(3.9, 31.5, -4),
	'C': normalizeRatio(3, 31.5, -4),
	'D': normalizeRatio(3, 31.5, -4),
	'R': normalizeRatio(2.75, 31.5, -4),
	'M': normalizeRatio(2.75, 31.5, -4),
	'A': normalizeRatio(2.5, 31.5, -4),
	'T': normalizeRatio(2.5, 31.5, -4),
	'P': normalizeRatio(2.1, 31.5, -4),
	'Y': normalizeRatio(1.95, 31.5, -4),
	'K': normalizeRatio(1.75, 31.5, -4),
	'N': normalizeRatio(1.5, 31.5, -4),
	'L': normalizeRatio(1.2, 31.5, -4),
	'J': normalizeRatio(0.6, 31.5, -4),
	'F': normalizeRatio(0.25, 31.5, -4),
	'I': normalizeRatio(0.15, 31.5, -4),
	'O': normalizeRatio(-0.1, 31.5, -4),
	'B': normalizeRatio(-0.5, 31.5, -4),
	'G': normalizeRatio(-1.5, 31.5, -4),
	'W': normalizeRatio(-1.9, 31.5, -4),
	'U': normalizeRatio(-3.2, 31.5, -4),
	'V': normalizeRatio(-3.75, 31.5, -4),
	'Q': normalizeRatio(-4.5, 31.5, -4),
}

func normalizeRatio(ratio, max, min float64) float64 {
	return 1 - (ratio-min)/(max-min)
}

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
	allTiles := []rune("__EEEEEEEEEEEEAAAAAAAAAIIIIIIIIIOOOOOOOONNNNNNRRRRRRTTTTTTLLLLSSSSUUUUDDDDGGGBBCCMMPPFFHHVVWWYYKJXQZ")

	return CreateGrabble(officialDict, officialScrabbleBoard, players, allTiles, tilePoints, 7)
}
