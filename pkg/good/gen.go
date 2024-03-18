//go:build ignore

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/shizukayuki/excel-hk4e"
)

type Keys struct {
	Type string
	Keys map[uint32]string
}

var tmpl = `package {{ .Name }}
{{- range .Data }}

{{ $type := .Type -}}
type {{ $type }} uint32

const (
	Unknown{{ $type }} {{ $type }} = 0 //
{{- range $id, $name := .Keys }}
	{{ $name }} {{ $type }} = {{ $id }}
{{- end }}
)
{{- end }}
`

var (
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
	err := excel.LoadResources(func(name string, v any) error {
		d, err := os.ReadFile(filepath.Join(repo, name))
		if err != nil {
			return err
		}
		return json.Unmarshal(d, v)
	})
	check(err)

	d := struct {
		Name string
		Data []Keys
	}{Name: os.Args[1]}
	outPath := os.Args[2]

	var relic Keys
	relic.Type = "ArtifactSetKey"
	relic.Keys = make(map[uint32]string)
	for _, x := range excel.ReliquaryCodexExcelConfigData {
		relic.Keys[x.SuitId] = excel.Slug(x.Set().Affix(0).Name())
	}
	d.Data = append(d.Data, relic)

	var char Keys
	char.Type = "CharacterKey"
	char.Keys = make(map[uint32]string)
	for _, x := range excel.AvatarCodexExcelConfigData {
		char.Keys[x.AvatarId] = excel.Slug(x.Avatar().Name())
	}
	// traveler
	a := excel.FindAvatar(10000007)
	char.Keys[a.Id] = excel.Slug(a.Name())
	d.Data = append(d.Data, char)

	var weap Keys
	weap.Type = "WeaponKey"
	weap.Keys = make(map[uint32]string)
	for _, x := range excel.WeaponCodexExcelConfigData {
		if x.Weapon().StoryId == 0 {
			continue
		}
		weap.Keys[x.WeaponId] = excel.Slug(x.Weapon().Name())
	}
	d.Data = append(d.Data, weap)

	var buf bytes.Buffer
	err = template.Must(template.New("").Parse(tmpl)).Execute(&buf, d)
	check(err)
	out, err := format.Source(buf.Bytes())
	check(err)
	err = os.WriteFile(outPath, out, 0o644)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
