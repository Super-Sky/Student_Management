package util

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"student/common"
)

type snowResponse struct {
	Id int64 `json:"id,string"`
}

func GetSnow(ctx context.Context) int64 {
	c := common.Config()
	snowApi := c.SnowServerAddr + "/snowflake"
	resp, err := HttpGetJson(snowApi, common.GVA_LOG.Logger)
	if err != nil {
		common.Error(ctx, "GetSnow snowServerApi fail", zap.Error(err))
		return 0
	}
	var response snowResponse
	if err = json.Unmarshal(resp, &response); err != nil {
		common.Error(ctx, "GetSnow Unmarshal fail", zap.Error(err))
		return 0
	}
	common.Debug(ctx, "GetSnow from snowServer", zap.Int64("response.ID", response.Id))
	return response.Id
}
