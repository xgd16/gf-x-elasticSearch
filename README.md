# gf-x-elasticSearch
## 基于 GF 的 ES 操作库
### 注意事项使用前需要调用 [es.Init](./es/handler.go)
### 演示
```go
    // 添加
	_, err = data.SendRequest(data.ElasticSearchRequest[esapi.IndexRequest]{
		Request: esapi.IndexRequest{
			Index:      "testx",
			DocumentID: "861533#######",
			Refresh:    "true",
		},
		Body:         map[string]any{"a": 1, "b": 2, "c": 4},
		JsonCallBack: data.IndexRequestJsonCallBack,
	}.Create())
	
	fmt.Println(err) 
	// 获取
	data, err := data.SendRequest(data.ElasticSearchRequest[esapi.GetRequest]{
		Request: esapi.GetRequest{
			Index:      "testx",
			DocumentID: "861533#######",
		},
	}.Create())
	
	fmt.Println(data.MapStrAny(), err)
```