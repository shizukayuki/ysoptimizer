package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shizukayuki/excel-hk4e"
	"github.com/shizukayuki/ysoptimizer/pkg/good"
)

var (
	optimized = map[good.CharacterKey]OptimizeState{}
	slow      = flag.Bool("slow", false, "Run in slow mode. Don't filter artifacts that have dead stats")
	path      = flag.String("good-file", "./GOOD.json", "Location of GOOD.json")
	excelPath = flag.String("genshin-data", "", "Specify GenshinData path. Defaults to $HOME/git/GenshinData or $GENSHIN_DATA_REPO if set")
)

func main() {
	flag.Parse()

	repo := os.Getenv("GENSHIN_DATA_REPO")
	if *excelPath != "" {
		repo = *excelPath
	}
	if repo == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			check(err)
		}
		repo = filepath.Join(home, "git", "GenshinData")
	}
	err := excel.LoadResources(excel.LoaderConfig{
		Root:      repo,
		Languages: []string{"en"},
	})
	check(err)

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

	db := ParseGODatebase(*path)
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

	data, err = handleTraveler(data)
	check(err)

	var db good.Datebase
	err = json.Unmarshal(data, &db)
	check(err)

	return &db
}

func handleTraveler(data []byte) ([]byte, error) {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	var chars []map[string]json.RawMessage
	if err := json.Unmarshal(obj["characters"], &chars); err != nil {
		return nil, err
	}

	for _, c := range chars {
		for _, name := range []string{
			"Manekin",
			"Manekina",
			"Traveler",
		} {
			if strings.HasPrefix(string(c["key"]), "\""+name) {
				c["key"] = json.RawMessage("\"" + name + "\"")
				break
			}
		}
	}

	var err error
	obj["characters"], err = json.Marshal(chars)
	if err != nil {
		return nil, err
	}

	return json.Marshal(obj)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
