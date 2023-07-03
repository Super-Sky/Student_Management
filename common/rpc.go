package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc"
)

var (
	rpcConn *grpc.ClientConn
)

func InitRpc() {
	var err error
	rpcConn, err = grpc.Dial(Config().RpcServer.BaseUrl+":"+Config().RpcServer.Port, grpc.WithInsecure())
	if err != nil {
		panic("rpc init error")
	}
}

// 获取Rpc
func RpcConn() *grpc.ClientConn {
	if rpcConn == nil {
		panic("please init rpc first")
	}
	return rpcConn
}

type JSON json.RawMessage

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
