package es

import (
	"bytes"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	// IndexRequestJsonCallBack 创建索引数据预制回调
	IndexRequestJsonCallBack = func(req any, json []byte) error {
		req.(*esapi.IndexRequest).Body = bytes.NewBuffer(json)

		return nil
	}
)

// ElasticSearchRequest ES 请求
type ElasticSearchRequest[T esapi.Request] struct {
	Request      T
	Body         any
	JsonCallBack func(req any, json []byte) error
	error        error
}

// Create 创建 ES 请求
func (e ElasticSearchRequest[T]) Create() *ElasticSearchRequest[T] {
	// 存在body时执行
	if e.Body != nil {
		// 定义异常
		var err error
		// 转换为json
		encode, err := gjson.Encode(e.Body)
		// 存在回调时执行
		if e.JsonCallBack != nil {
			// 回调json解析
			if callbackErr := e.JsonCallBack(&e.Request, encode); callbackErr != nil && err == nil {
				err = callbackErr
			}
		}
	}
	// 返回请求数据
	return &e
}

func (e ElasticSearchRequest[T]) GetRequest() (esapi.Request, *elasticsearch.Client, error) {
	return e.Request, Elastic.Client, e.error
}

var Elastic = &ElasticSearch{}

// ElasticSearch ES查询引擎
type ElasticSearch struct {
	Client *elasticsearch.Client
}

// Init 初始化
func (e *ElasticSearch) Init() error {
	var err error
	// 获取配置信息
	esCfg, err := g.Cfg().Get(gctx.New(), "es")

	if err != nil {
		return err
	}

	esData := esCfg.MapStrVar()

	// 初始化配置
	e.Client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: esData["address"].Strings(),
		Username:  esData["username"].String(),
		Password:  esData["password"].String(),
	})

	if err != nil {
		return err
	}

	return nil
}

// SendRequest 创建索引数据
func SendRequest[T esapi.Request](ESRequest *ElasticSearchRequest[T]) (*gjson.Json, error) {
	request, client, err := ESRequest.GetRequest()
	// 处理获取数据错误
	if err != nil {
		return nil, err
	}
	// 发起请求
	res, err := request.Do(gctx.New(), client)
	// 发起请求失败返回
	if err != nil {
		return nil, err
	}
	// 判断是否请求错误
	if res.IsError() {
		return nil, errors.New("请求失败: " + res.String())
	}
	// 转换为string
	buf := new(bytes.Buffer)

	if _, err := buf.ReadFrom(res.Body); err != nil {
		return nil, err
	}

	jsonData, err := gjson.DecodeToJson(buf.String())
	// 处理解析json失败
	if err != nil {
		return nil, err
	}
	// 关闭请求释放资源
	if err := res.Body.Close(); err != nil {
		return nil, err
	}

	return jsonData, nil
}
