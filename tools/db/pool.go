package db

import (
	"fmt"

	"container/list"

	"github.com/MagicYH/stock-analysis/tools/db/connect"
)

// PoolConfig is use to store config of each connection
type PoolConfig struct {
	Name   string
	Driver string
	Config []interface{}
}

// ConnectPool is the acture struct that store connections
type ConnectPool struct {
	Driver   string
	Config   []interface{}
	Count    int
	ConnList *list.List
}

var _connectPools map[string]*ConnectPool

func init() {
	_connectPools = make(map[string]*ConnectPool)
}

// InitPool is used to init pools
func InitPool(poolConfigs []PoolConfig) {
	for _, config := range poolConfigs {
		connectPool := ConnectPool{config.Driver, config.Config, 0, list.New()}
		_connectPools[config.Name] = &connectPool
	}
}

// GetConnect is used to get a new ConnectInterface
func GetConnect(name string) (connect.ConnectInterface, error) {
	connectPool, ok := _connectPools[name]
	if !ok {
		return nil, fmt.Errorf("Connect %s not exsist", name)
	}

	var connection connect.ConnectInterface
	var err error
	err = nil
	// May add lock here
	if connectPool.ConnList.Len() > 0 {
		element := connectPool.ConnList.Front()
		connectPool.ConnList.Remove(element)
		connection = element.Value.(connect.ConnectInterface)
	} else {
		connection, err = connect.NewConnect(name, connectPool.Driver, connectPool.Config...)
		connectPool.Count = connectPool.Count + 1
	}
	return connection, err
}

// ReleaseConnect is used to recycle ConnectInterfacec
func ReleaseConnect(connection connect.ConnectInterface) {
	connectPool, _ := _connectPools[connection.GetName()]
	connectPool.ConnList.PushBack(connection)
}
