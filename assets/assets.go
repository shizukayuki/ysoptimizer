package assets

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/shizukayuki/excel-hk4e"
)

func Load(dir string) {
	dir = strings.TrimSuffix(dir, "/")
	err := excel.LoadResources(func(name string, v any) error {
		d, err := os.ReadFile(dir + "/" + name)
		if err != nil {
			return err
		}
		return json.Unmarshal(d, v)
	})
	if err != nil {
		panic(err)
	}
}
