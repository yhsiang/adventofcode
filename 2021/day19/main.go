package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

type beaconPair struct {
	p1 string
	p2 string
}

type twoPairs struct {
	bp1 beaconPair
	bp2 beaconPair
}

type scanner struct {
	number         int
	position       string
	beacons        []string
	rotatedBeacons [][]string
	fingerprint    map[beaconPair]float64
}

// learn from https://www.reddit.com/r/adventofcode/comments/rjpf7f/comment/hpk7i8c/?utm_source=reddit&utm_medium=web2x&context=3
func initScanners(input string) (scanners []*scanner) {
	data := strings.Split(input, "\n\n")
	for i, d := range data {
		coords := strings.Split(d, "\n")
		scanner := &scanner{
			number: i,
		}

		if i == 0 {
			scanner.position = "0,0,0"
		}
		for _, coord := range coords[1:] {
			scanner.addBeacon(coord)
		}
		// scanner.beacons = append(scanner.beacons, coords[1:]...)
		scanners = append(scanners, scanner)
	}

	return
}

func (s *scanner) calculateDistanceFingerprint(recalculate bool) map[beaconPair]float64 {
	if len(s.fingerprint) > 0 && !recalculate {
		return s.fingerprint
	}
	var fp map[beaconPair]float64
	if recalculate {
		fp = s.fingerprint
	} else {
		fp = make(map[beaconPair]float64)
	}

	for _, b1 := range s.beacons {
		for _, b2 := range s.beacons {
			if b1 == b2 {
				continue
			}
			_, fpb1b2 := fp[beaconPair{b1, b2}]
			_, fpb2b1 := fp[beaconPair{b2, b1}]
			if fpb1b2 || fpb2b1 {
				continue
			}
			d := dist(b1, b2)
			fp[beaconPair{b1, b2}] = d
		}
	}
	s.fingerprint = fp
	return fp
}

func (s *scanner) getAllPossibleRotatedBeacons() (beacons [][]string) {
	if len(s.rotatedBeacons) > 0 {
		return s.rotatedBeacons
	}
	for i := 0; i < 24; i++ {
		beacons = append(beacons, []string{})
	}
	for _, b := range s.beacons {
		// positive x
		for i := 0; i < 24; i++ {
			beacons[i] = append(beacons[i], rotateByType(b, i))
		}
	}
	s.rotatedBeacons = beacons
	return
}

func (s *scanner) addBeacon(b string) {
	for _, b2 := range s.beacons {
		if b == b2 {
			return
		}
	}
	s.beacons = append(s.beacons, b)
}

func rotateByType(b string, rotType int) string {
	x, y, z := util.Coord3d(b)
	switch rotType {
	case 0:
		return fmt.Sprintf("%d,%d,%d", +x, +y, +z)
	case 1:
		return fmt.Sprintf("%d,%d,%d", +x, -z, +y)
	case 2:
		return fmt.Sprintf("%d,%d,%d", +x, -y, -z)
	case 3:
		return fmt.Sprintf("%d,%d,%d", +x, +z, -y)
	// negative x
	case 4:
		return fmt.Sprintf("%d,%d,%d", -x, -y, +z)
	case 5:
		return fmt.Sprintf("%d,%d,%d", -x, +z, +y)
	case 6:
		return fmt.Sprintf("%d,%d,%d", -x, +y, -z)
	case 7:
		return fmt.Sprintf("%d,%d,%d", -x, -z, -y)
	// positive y
	case 8:
		return fmt.Sprintf("%d,%d,%d", +y, +z, +x)
	case 9:
		return fmt.Sprintf("%d,%d,%d", +y, -x, +z)
	case 10:
		return fmt.Sprintf("%d,%d,%d", +y, -z, -x)
	case 11:
		return fmt.Sprintf("%d,%d,%d", +y, +x, -z)
	// negative y
	case 12:
		return fmt.Sprintf("%d,%d,%d", -y, -z, +x)
	case 13:
		return fmt.Sprintf("%d,%d,%d", -y, +x, +z)
	case 14:
		return fmt.Sprintf("%d,%d,%d", -y, +z, -x)
	case 15:
		return fmt.Sprintf("%d,%d,%d", -y, -x, -z)
	// positive z
	case 16:
		return fmt.Sprintf("%d,%d,%d", +z, +x, +y)
	case 17:
		return fmt.Sprintf("%d,%d,%d", +z, -y, +x)
	case 18:
		return fmt.Sprintf("%d,%d,%d", +z, -x, -y)
	case 19:
		return fmt.Sprintf("%d,%d,%d", +z, +y, -x)
	// negative z
	case 20:
		return fmt.Sprintf("%d,%d,%d", -z, -x, +y)
	case 21:
		return fmt.Sprintf("%d,%d,%d", -z, +y, +x)
	case 22:
		return fmt.Sprintf("%d,%d,%d", -z, +x, -y)
	case 23:
		return fmt.Sprintf("%d,%d,%d", -z, -y, -x)
	}
	return b
}

func dist(p1 string, p2 string) float64 {
	x1, y1, z1 := util.Coord3d(p1)
	x2, y2, z2 := util.Coord3d(p2)

	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2) + math.Pow(float64(z1-z2), 2))
}

func setFingerprints(scanners []*scanner) {
	for _, s := range scanners {
		s.calculateDistanceFingerprint(false)
	}
}

func matchFingerprints(f0, f1 map[beaconPair]float64, overlap int) (bool, []twoPairs) {
	matchCount := 0
	var matches []twoPairs
	// a fingerprint is a map of about b^2 pairs
	// it's a match if there are 12 distances in f0 that are also in f1
	for f0p, f0f := range f0 {
		if f0f > 252 {
			continue // f0f comes from scanner 0 that could have fingerprints all over the place, but if they're far apart they won't be in F1
		}
		for f1p, f1f := range f1 {
			if f0f == f1f {
				matchCount++
				// a,b in f0p  and c,d in f1p - could be a=c and b=d, or a=c and b=d
				matches = append(matches, twoPairs{f0p, f1p})
			}
		}
	}
	//log.Printf("Fingerprint matches: %d", matchCount)
	return matchCount >= overlap, matches
}

func findBestScannerMatchingFingerprint(f0 map[beaconPair]float64, scanners []*scanner) (matchingScanner *scanner, matches []twoPairs) {
	var match bool
	type mcount struct {
		matches int
		s       *scanner
		m       []twoPairs
	}
	matchers := make(map[int]mcount)
	setFingerprints(scanners)
	for _, s := range scanners {
		match, matches = matchFingerprints(f0, s.fingerprint, 12)
		if match {
			matchers[s.number] = mcount{len(matches), s, matches}
			matchingScanner = s
		}
	}
	var bestms *scanner
	var bestmatches []twoPairs
	max := 0
	for _, mc := range matchers {
		if mc.matches > max {
			max = mc.matches
			bestms = mc.s
			bestmatches = mc.m
		}
	}
	return bestms, bestmatches
	//	return matchingScanner, matches
}

func translateBeacon(b string, x1, y1, z1 int) string {
	x, y, z := util.Coord3d(b)
	return fmt.Sprintf("%d,%d,%d", x+x1, y+y1, z+z1)
}

func calculateBeaconDiff(b1, b2 string) (x, y, z int) {
	x1, y1, z1 := util.Coord3d(b1)
	x2, y2, z2 := util.Coord3d(b2)
	x = x1 - x2
	y = y1 - y2
	z = z1 - z2
	return
}

func translateBeacons(beacons []string, bp beaconPair) []string {
	ourBeacon := bp.p1
	theirBeacon := bp.p2
	x, y, z := calculateBeaconDiff(ourBeacon, theirBeacon)
	//log.Printf("translation: %d,%d,%d", x, y, z)

	translatedBeacons := []string{}
	for _, b := range beacons {
		translatedBeacons = append(translatedBeacons, translateBeacon(b, x, y, z))
	}
	return translatedBeacons
}

func checkBeaconMatch(b0, b1 []string, overlap int) bool {
	matchCount := 0
	for _, b0b := range b0 {
		x1, y1, z1 := util.Coord3d(b0b)

		for _, b1b := range b1 {
			//if b0b == b1b {
			x2, y2, z2 := util.Coord3d(b1b)
			if x1 == x2 && y1 == y2 && z1 == z2 {
				matchCount++
			}
		}
	}
	return matchCount >= overlap
}

func matchScanner(s0 *scanner, matchingScanner *scanner, matches []twoPairs) (int, int, int, int, []string) {

	// key is rotation type, value is []beacon
	rotations := matchingScanner.getAllPossibleRotatedBeacons()
	// try the translations from the first match- maybe we don't have to loop them all
	// bp1.p1 = bp2.p1 and bp1.p2 = bp2.p2, or bp1.p2 = bp2.p1 and bp1.p1 = bp2.p2
	// four transforms to try ( * 24 rotations...)
	for matchToTry := range matches {
		bpToTry := []beaconPair{
			{matches[matchToTry].bp1.p1, matches[matchToTry].bp2.p1},
			{matches[matchToTry].bp1.p1, matches[matchToTry].bp2.p2},
			{matches[matchToTry].bp1.p2, matches[matchToTry].bp2.p1},
			{matches[matchToTry].bp1.p2, matches[matchToTry].bp2.p2},
		}
		for rotType, rotatedBeacons := range rotations {
			for _, b := range bpToTry {
				newBp := beaconPair{b.p1, rotateByType(b.p2, rotType)}
				translatedBeacons := translateBeacons(rotatedBeacons, newBp)
				// if this is true, we know how to align s0 and matchingscanner
				match := checkBeaconMatch(s0.beacons, translatedBeacons, 12)
				if match {
					x, y, z := calculateBeaconDiff(newBp.p1, newBp.p2)
					log.Printf("Beacons aligned: translation %d,%d,%d", x, y, z)
					// So now we know the rotType and x,y,z translation to line up matchingScanner with s0
					return rotType, x, y, z, rotatedBeacons
				}
			}
		}
	}
	log.Panicf("matchScanner should have found something")
	return 0, 0, 0, 0, nil
}

func removeMatchedScanner(s *scanner, scanners []*scanner) (filteredScanners []*scanner) {
	for _, s2 := range scanners {
		if s2.number != s.number {
			filteredScanners = append(filteredScanners, s2)
		}
	}
	return filteredScanners
}

func calcMaxManhattanDistance(points []string) int {
	max := 0
	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}
			d := manhattanDistance(p1, p2)
			if d > max {
				max = d
			}
		}
	}

	return max
}

func manhattanDistance(p1, p2 string) int {
	x1, y1, z1 := util.Coord3d(p1)
	x2, y2, z2 := util.Coord3d(p2)
	return util.Abs(x1-x2) + util.Abs(y1-y2) + util.Abs(z1-z2)
}

func run(input string) (int, int) {
	scanners := initScanners(input)
	s0 := scanners[0]
	s0.calculateDistanceFingerprint(false)
	matches := scanners[1:]
	var locatedScanners []string
	locatedScanners = append(locatedScanners, s0.position)
	for len(matches) > 0 {
		matchingScanner, twoPairs := findBestScannerMatchingFingerprint(s0.fingerprint, matches)
		if matchingScanner == nil {
			log.Panicf("no matching beacon sets, with %d scanners left", len(matches))
		}
		log.Printf("Scanner %d matched- lining up beacons", matchingScanner.number)
		_, x, y, z, rotatedBeacons := matchScanner(s0, matchingScanner, twoPairs)
		// using these, we can add all of the beacons on matchingscanner to s0
		for _, rb := range rotatedBeacons {
			s0.addBeacon(translateBeacon(rb, x, y, z))
		}
		s0.calculateDistanceFingerprint(true)
		log.Printf("s0 now has %d beacons", len(s0.beacons))
		// matchingScanner.location.translate(x, y, z)
		matchingScanner.position = fmt.Sprintf("%d,%d,%d", x, y, z)
		locatedScanners = append(locatedScanners, matchingScanner.position)
		matches = removeMatchedScanner(matchingScanner, matches)
	}

	return len(s0.beacons), calcMaxManhattanDistance(locatedScanners)

}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	beaconCount, maxScannerDist := run(file)
	fmt.Printf("part1: %d\n", beaconCount)
	fmt.Printf("part2: %d\n", maxScannerDist)
}
