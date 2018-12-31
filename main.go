package main

import (
	"github.com/MagicYH/stock-analysis/tools/db"
)

func main() {
	poolConfigs := make([]db.PoolConfig, 0)
	db.InitPool(poolConfigs)
}
