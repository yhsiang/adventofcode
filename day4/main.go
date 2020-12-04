package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Height struct {
	Value int64
	Unit  string
}

func NewHeight(str string) *Height {
	var h = &Height{}
	switch str[len(str)-2:] {
	case "cm":
		i, err := strconv.ParseInt(str[:len(str)-2], 10, 64)
		if err != nil {
			return h
		}
		h.Value = i
		h.Unit = "cm"
	case "in":
		i, err := strconv.ParseInt(str[:len(str)-2], 10, 64)
		if err != nil {
			return h
		}
		h.Value = i
		h.Unit = "in"
	}
	return h
}

type Passport struct {
	BirthYear      int64
	IssueYear      int64
	ExpirationYear int64
	Height         Height
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
	Valid          bool
}

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func checkEyeColor(str string) bool {
	for _, e := range eyeColors {
		if e == str {
			return true
		}
	}

	return false
}

func NewPassport(data []string) *Passport {
	var p = &Passport{}
	for _, d := range data {
		dd := strings.Split(strings.TrimSpace(d), ":")

		switch dd[0] {
		case "byr":
			i, err := strconv.ParseInt(dd[1], 10, 64)
			if err != nil {
				return p
			}
			p.BirthYear = i
		case "iyr":
			i, err := strconv.ParseInt(dd[1], 10, 64)
			if err != nil {
				return p
			}
			p.IssueYear = i
		case "eyr":
			i, err := strconv.ParseInt(dd[1], 10, 64)
			if err != nil {
				return p
			}
			p.ExpirationYear = i
		case "hgt":
			p.Height = *NewHeight(dd[1])
		case "hcl":
			p.HairColor = dd[1]
		case "ecl":
			p.EyeColor = dd[1]
		case "pid":
			p.PassportID = dd[1]
		case "cid":
			p.CountryID = dd[1]
		}
	}

	if len(data) == 8 || (len(data) == 7 && p.CountryID == "") {
		p.Valid = true
	}

	if p.BirthYear < 1920 || p.BirthYear > 2002 {
		p.Valid = false
	}
	// fmt.Printf("b %+v\n", p)
	if p.IssueYear < 2010 || p.IssueYear > 2020 {
		p.Valid = false
	}
	// fmt.Printf("i %+v\n", p)
	if p.ExpirationYear < 2020 || p.ExpirationYear > 2030 {
		p.Valid = false
	}
	// fmt.Printf("e %+v\n", p)
	switch p.Height.Unit {
	case "cm":
		if p.Height.Value < 150 || p.Height.Value > 193 {
			p.Valid = false
		}
	case "in":
		if p.Height.Value < 59 || p.Height.Value > 76 {
			p.Valid = false
		}
	default:
		p.Valid = false
	}
	// fmt.Printf("h %+v\n", p)
	matched, err := regexp.Match(`^#[0-9a-f]{6}`, []byte(p.HairColor))
	if err != nil || !matched {
		p.Valid = false
	}
	// fmt.Printf("ha %+v\n", p)
	if !checkEyeColor(p.EyeColor) {
		p.Valid = false
	}
	// fmt.Printf("eye %+v\n", p)
	matched, err = regexp.Match(`^[0-9]{9}$`, []byte(p.PassportID))
	if err != nil || !matched {
		p.Valid = false
	}
	// fmt.Printf("p %+v\n", p)
	return p
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	var passportData []string
	var passports []*Passport
	for i, line := range lines {
		if line == "" {
			passports = append(passports, NewPassport(passportData))
			passportData = []string{}
			continue
		}
		passportData = append(passportData, strings.Split(line, " ")...)
		if i == len(lines)-1 {
			passports = append(passports, NewPassport(passportData))
			break
		}
	}

	var num int
	for _, p := range passports {
		if p.Valid {
			// fmt.Printf("%+v\n", p)
			num += 1
		}
		// fmt.Printf("%+v\n", p)
	}
	fmt.Println(num)
}
