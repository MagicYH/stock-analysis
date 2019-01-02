package main

import (
	"fmt"

	"github.com/MagicYH/stock-analysis/tools"
)

func main() {
	tools.LocdGlobalConfig("/Users/magic/Project/Source/gopath/src/github.com/MagicYH/stock-analysis/config/config.toml")
	config := tools.GetGlobalConfig()
	fmt.Printf("%v", config)
}
