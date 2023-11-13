package main

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func zlibtest() {
	/*err := queryCache.Set("k1", []byte("v1"), 1)
	fmt.Printf("err:%+v\n", err)
	v, err := queryCache.Get("k1")
	fmt.Printf("%+v, %+v\n", v, err)
	v, err = queryCache.Get("k2")
	fmt.Printf("%+v, %+v\n", v, err)*/

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte(`{"Req":{"sql":"select","dbName":"PUSH_STATIS_R","cacheValidTime":20000,"cacheSaveTime":20000,"dataBaseName":"db_push_news_recomm"},"data_vec":[{"data":{"k1":"v1","k2":"v2"}}],"store_time":"2022-12-23T11:16:49.2329833+08:00","expire_duration":1000000000}`))
	w.Close()
	bs := b.Bytes()
	fmt.Printf("%+v\n%+v", string(bs), bs)

	dbValue := []*DBValue{}
	dbValue = append(dbValue, &DBValue{Data: map[string]string{"k1": "v1", "k2": "v2"}})

	v1Byte := dbData2Bytes(&SqlRequest{
		Sql:            "select",
		DbName:         "PUSH_STATIS_R",
		CacheValidTime: 20000,
		CacheSaveTime:  20000,
		DataBaseName:   "db_push_news_recomm",
	}, dbValue, 1*time.Second)
	fmt.Printf("%+v\n", v1Byte)
	// queryCache.Set("k1", v1Byte, 2)
	// v1, err := queryCache.Get("k1")
	// s1, err1 := bytesToDBData(v1)
	// fmt.Printf("%+v, %+v, %+v\n", s1, err1, err)

	v2Byte := dbData2Bytes(&SqlRequest{
		Sql:            "select",
		DbName:         "PUSH_STATIS_R",
		CacheValidTime: 20000,
		CacheSaveTime:  20000,
		DataBaseName:   "db_push_news_recomm",
	}, getDataVec(), 1*time.Second)
	s2, err1 := bytesToDBData(v2Byte)
	fmt.Printf("%+v, %+v\n", s2, err1)
}

type value struct {
	// Req sql请求
	Req *SqlRequest
	// DataVec 数据
	DataVec []*DBValue `json:"data_vec"`
	// StoreTime 插入时间
	StoreTime time.Time `json:"store_time"`
	// ExpireDuration 过期时间
	ExpireDuration time.Duration `json:"expire_duration"`
}

func dbData2Bytes(req *SqlRequest, dataVec []*DBValue, expireDuration time.Duration) []byte {
	start := time.Now()
	// defer func() { metrics.IncrCounter("dbData2Bytes_cost", float64(time.Since(start).Milliseconds())) }()
	b, err := jsoniter.Marshal(value{
		Req:            req,
		DataVec:        dataVec,
		StoreTime:      time.Now(),
		ExpireDuration: expireDuration,
	})
	// fmt.Printf("json.Marshal err:%+v, len:%d, %+v", err, len(b), string(b))
	var buffer bytes.Buffer
	w := zlib.NewWriter(&buffer)
	fmt.Printf("NewWriter success:%+v, json:%s\n", err, string(b))
	n, err := w.Write(b)
	fmt.Printf("NewWriter write:%d, %+v\n", n, err)
	w.Close()
	if err != nil {
		fmt.Printf("NewWriter err:%+v, %+v\n", err, start)
		return []byte{}
	}
	fmt.Printf("NewWriter write success:%+v, %+v\n", err, start)

	saveBytes := buffer.Bytes()
	fmt.Printf("len: json:%d, zlib:%d\n", len(b), len(saveBytes))
	return saveBytes
}
func bytesToDBData(v []byte) (dataVec value, err error) {
	// start := time.Now()
	// defer func() { metrics.IncrCounter("bytesToDBData_cost", float64(time.Since(start).Milliseconds())) }()
	buffer := bytes.NewReader(v)
	r, err := zlib.NewReader(buffer)
	if err != nil {
		return value{}, err
	}
	defer r.Close()
	var out bytes.Buffer
	_, err = io.Copy(&out, r)
	if err != nil {
		return value{}, err
	}
	b := out.Bytes()

	err = json.Unmarshal(b, &dataVec)
	if err != nil {
		return value{}, fmt.Errorf("value not value:%T", v)
	}
	return dataVec, nil
}

type SqlRequest struct {
	// tRPC数据校验模块 开启自动数据校验。限制传入参数只能为tsecstr默认安全类型
	Sql            string `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`                        // 预编译sql，可带?
	DbName         string `protobuf:"bytes,3,opt,name=dbName,proto3" json:"dbName,omitempty"`                  // db名，可以传数据库全名（newcpush.mdb.mig:15075），也可以传七彩石db key，不过都要在七彩石进行配置帐号密码
	CacheValidTime int64  `protobuf:"varint,4,opt,name=cacheValidTime,proto3" json:"cacheValidTime,omitempty"` // 缓存在cacheValidTime内直接返回，单位ms，0表示直接查数据库
	CacheSaveTime  int64  `protobuf:"varint,5,opt,name=cacheSaveTime,proto3" json:"cacheSaveTime,omitempty"`   // 本次查询缓存有效期，单位ms，0表示本次数据库查询不更新缓存。建议设置2 * cacheValidTime
	DataBaseName   string `protobuf:"bytes,6,opt,name=dataBaseName,proto3" json:"dataBaseName,omitempty"`      // dbName 传数据库全名时务必传一下dataBaseName，保证不会连接到错误的库表 dbname和database是跟db_proxy的mysql comm配置相对应的
}

type DBValue struct {
	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
