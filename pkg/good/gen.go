//go:build ignore

package main

import (
	"bytes"
	"go/format"
	"html/template"
	"log"
	"os"
	"regexp"
	"strings"

	_ "github.com/shizukayuki/ysoptimizer/assets"
	"github.com/shizukayuki/ysoptimizer/pkg/excel"
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

func main() {
	d := struct {
		Name string
		Data []Keys
	}{Name: os.Args[1]}
	outPath := os.Args[2]

	var relic Keys
	relic.Type = "ArtifactSetKey"
	relic.Keys = make(map[uint32]string)
	for _, x := range excel.ReliquaryCodexExcelConfigData {
		relic.Keys[x.SuitId] = sanitizeName(x.Set().Affix(0).Name())
	}
	d.Data = append(d.Data, relic)

	var char Keys
	char.Type = "CharacterKey"
	char.Keys = make(map[uint32]string)
	for _, x := range excel.AvatarCodexExcelConfigData {
		char.Keys[x.AvatarId] = sanitizeName(x.Avatar().Name())
	}
	// traveler
	a := excel.FindAvatar(10000007)
	char.Keys[a.Id] = sanitizeName(a.Name())
	d.Data = append(d.Data, char)

	var weap Keys
	weap.Type = "WeaponKey"
	weap.Keys = make(map[uint32]string)
	for _, x := range excel.WeaponCodexExcelConfigData {
		if x.IsDisuse {
			continue
		}
		weap.Keys[x.WeaponId] = sanitizeName(x.Weapon().Name())
	}
	d.Data = append(d.Data, weap)

	var buf bytes.Buffer
	err := template.Must(template.New("").Parse(tmpl)).Execute(&buf, d)
	check(err)
	out, err := format.Source(buf.Bytes())
	check(err)
	err = os.WriteFile(outPath, out, 0o644)
	check(err)
}

var reSanitizeName = regexp.MustCompile(`[^a-zA-Z0-9]`)

func sanitizeName(s string) string {
	s = strings.ReplaceAll(s, "'", "")
	s = strings.Title(s)
	s = reSanitizeName.ReplaceAllString(s, "")
	return s
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
