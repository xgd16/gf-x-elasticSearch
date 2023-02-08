package test

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gogf/gf/v2/frame/g"
	"gogs.mirlowz.com/x/gf-x-elasticSearch/es"
	"testing"
)

func TestEs(t *testing.T) {
	_ = es.Elastic.Init()

	data, err := es.SendRequest(es.ElasticSearchRequest[esapi.SearchRequest]{
		Request: esapi.SearchRequest{
			Index: []string{"testa"},
			Body: bytes.NewBuffer([]byte(`
				{"query":{
					"match_all": {}
				}}
			`)),
		},
	}.Create())

	fmt.Println(data.Interface(), err)
}

// 查询修改
func TestUpdateRequest(t *testing.T) {
	_ = es.Elastic.Init()

	fmt.Println(es.UpdateRequest("fans_list", g.Map{
		"bool": g.Map{
			"must": g.List{
				{
					"match": g.Map{
						"invs.socialAccountId": 33,
					},
				},
				{
					"match": g.Map{
						"account": "uf1b53db616968a69654961c70ee67d44",
					},
				},
			},
		},
	}, g.Map{
		"avatar": "http://123.com",
		"name":   "于谦",
	}))
}
