package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = http.Client{
		Timeout: 30 * time.Minute,
	}
)

// HttpGetJson Get请求获取json数据
func HttpGetJson(url string, lg *zap.Logger) ([]byte, error) {
	defer func() {
		if e := recover(); e != nil {
			lg.Error("HttpGetJson失败", zap.Any("Err", e))
		}
	}()
	client := &http.Client{Timeout: 30 * time.Minute}
	resp, err := client.Get(url)
	if err != nil {
		lg.Error("构建get-http请求失败", zap.Error(err))
	}
	defer resp.Body.Close()
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lg.Error("获取流失败", zap.Error(err))
		return nil, err
	}
	return body1, nil
}

// HttpPostJson 发送json数据
func HttpPostJson(url string, data interface{}, header map[string]string, lg *zap.Logger) ([]byte, error) {
	defer func() {
		if e := recover(); e != nil {
			lg.Error("HttpPostJson失败", zap.Any("Err", e))
		}
	}()
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(data); err != nil {
		lg.Error("解析data错误", zap.Error(err))
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		lg.Error("构建http请求失败", zap.Error(err))
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	if header != nil {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		lg.Error("发送res失败", zap.Error(err))
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			//fmt.Println("关闭body失败", err)
			lg.Error("关闭body失败", zap.Error(err))
		}
	}(response.Body)
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		lg.Error("获取流失败", zap.Error(err))
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		lg.Error("获取错误状态码", zap.Int("response.StatusCode", response.StatusCode))
		return result, errors.New("返回错误状态码")
	}
	return result, nil
}
