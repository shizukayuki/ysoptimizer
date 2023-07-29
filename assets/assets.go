package assets

import (
	"encoding/json"
	"os"

	"github.com/shizukayuki/ysoptimizer/pkg/excel"
)

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	err = excel.LoadResources(func(name string, v any) error {
		d, err := os.ReadFile(homedir + "/git/GenshinData/" + name)
		if err != nil {
			return err
		}
		return json.Unmarshal(d, v)
	})
	if err != nil {
		panic(err)
	}
}
