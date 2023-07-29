package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/shizukayuki/ysoptimizer/assets"
	"github.com/shizukayuki/ysoptimizer/pkg/good"
)

var optimized = map[good.CharacterKey]OptimizeState{}

func main() {
	flag.Parse()

	var prio []good.CharacterKey
	for _, s := range flag.Args() {
		char, err := good.CharacterKeyString(s)
		if err != nil {
			fmt.Println(priority)
			panic(err)
		}
		prio = append(prio, char)
	}
	if len(prio) == 0 {
		prio = priority
	}

	db := ParseGODatebase("./GOOD.json")
	for char, t := range config {
		t.Datebase = db

		for _, v := range db.Characters {
			if v.Key == char {
				t.Character = v
			}
		}
		for _, v := range db.Weapons {
			if v.Location == char {
				t.Weapon = v
			}
		}
		for _, v := range db.Artifacts {
			if v.Location == char {
				t.CurArtifacts[v.SlotKey] = v
			}
		}
	}

	for _, v := range db.Artifacts {
		v.Location = good.UnknownCharacterKey
		err := v.Process()
		check(err)
	}

	var total time.Duration
	for _, char := range prio {
		now := time.Now()

		cfg := config[char]
		cur := cfg.CurrentState()
		opt := cfg.Optimize()
		optimized[char] = opt

		dur := time.Since(now)
		total += dur

		change := false
		var a, b strings.Builder
		for i, x := range cur.Build {
			y := opt.Build[i]
			if y == nil {
				log.Fatalf("no build found for %s", cfg.Character.Key)
			}
			if x != y {
				change = true
				slot := strings.ToUpper(string(y.SlotKey.String()[0]))
				if x != nil {
					fmt.Fprintf(&a, "- [%s] %s\n", slot, x)
				}
				fmt.Fprintf(&b, "+ [%s] %s\n", slot, y)
			}
		}
		for _, a := range opt.Build {
			a.Location = char
		}

		result := fmt.Sprintf("%.0f->%.0f", cur.Result, opt.Result)
		if int(opt.Result) > int(cur.Result) {
			result = G(result)
		} else if int(opt.Result) < int(cur.Result) {
			result = R(result)
		}

		fmt.Printf("> %s / %s @ %s ; %.2f\n", cfg.Character.Key, cfg.Weapon.Key, result, dur.Seconds())
		if change {
			fmt.Print(R(a.String()))
			fmt.Print(G(b.String()))
			fmt.Println(opt)
		}
	}
	fmt.Printf("total %.2f\n", total.Seconds())
}

func R(s string) string {
	return "\x1b[31m" + s + "\x1b[1;0m"
}

func G(s string) string {
	return "\x1b[32m" + s + "\x1b[1;0m"
}

func ParseGODatebase(filename string) *good.Datebase {
	data, err := os.ReadFile(filename)
	check(err)

	var db good.Datebase
	err = json.Unmarshal(data, &db)
	check(err)

	return &db
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
