package basic

import (
	"context"
	"fmt"
	"github.com/tikv/client-go/v2/txnkv"
	"github.com/tikv/client-go/v2/txnkv/transaction"
	"strings"
	"time"
)

var (
	OS     = "OS"
	SLASH  = "/"
	BUCKET = "Bucket"
	LIST   = "List"
	NAME   = "Name"
	DATA   = "Data"
	LS     = "LS"
)

type MetaStore struct {
	tikvAddr           string
	defaultTikvTimeOut time.Duration
	txnClient          *txnkv.Client
}

func NewMetaStore(tikvAddr string) (*MetaStore, error) {
	tikvAddrs := strings.Split(tikvAddr, ",")
	client, err := txnkv.NewClient(tikvAddrs)
	if err != nil {
		return nil, err
	}

	ms := &MetaStore{
		txnClient:          client,
		defaultTikvTimeOut: time.Second * 60,
	}
	return ms, nil
}

// OS/T03/Bucket/List/Name/test1
func GetPathKey(tenantId, bucket string) []byte {
	return []byte(OS + SLASH + tenantId + SLASH + BUCKET + SLASH + LIST + SLASH + NAME + SLASH + bucket)
}

// OS/T03/123456789987/Data/LS/test/
func GetObjKey(tenantID, pathKey, obj string) []byte {
	return []byte(OS + SLASH + tenantID + SLASH + pathKey + SLASH + DATA + SLASH + LS + SLASH + obj + SLASH)
}

// 事务执行方法
func (ms *MetaStore) ExecuteTxn(fn func(txn *transaction.KVTxn) error) error {
	txn, err := ms.txnClient.Begin()
	if err != nil {
		return fmt.Errorf("事务启动失败: %w", err)
	}
	if err := fn(txn); err != nil {
		_ = txn.Rollback()
		return err
	}
	if err := txn.Commit(context.Background()); err != nil {
		return fmt.Errorf("事务提交失败: %w", err)
	}
	return nil
}

type Obj struct {
	Uuid string `json:"uuid"`
}
type LockMeta struct {
	LockTime    int64 `json:"lockTime"`
	MaxDuration int64 `json:"maxDuration"`
}

func ExpireTimeTest() {
	//metaStore, _ := basic.NewMetaStore("10.0.11.33:2379")
	//_ = metaStore.ExecuteTxn(func(txn *transaction.KVTxn) error {
	//	//txn.Delete([]byte("OS/T03/1010121741588318174/Data/Lock/b3b5a095-3c52-4b15-8b4e-f4486e0a1a58"))
	//	ctx := context.Background()
	//	get, _ := txn.Get(ctx, []byte("OS/T03/1010121741588318174/Data/LS/1-4GB.zip"))
	//	var obj Obj
	//	err := json.Unmarshal(get, &obj)
	//	if err != nil {
	//		fmt.Println(err)
	//		return err
	//	}
	//	lookMeta := &LockMeta{
	//		LockTime:    time.Now().UnixMilli(),
	//		MaxDuration: 100 * time.Minute.Milliseconds(),
	//	}
	//	bytes, err := json.Marshal(lookMeta)
	//	if err != nil {
	//		fmt.Println("json marshal err:", err)
	//		return err
	//	}
	//	err = txn.Set([]byte("OS/T03/1010121741588318174/Data/Lock/"+obj.Uuid), bytes)
	//	if err != nil {
	//		fmt.Println("set failed, err:", err)
	//		return err
	//	}
	//	return nil
	//})
}

//func get(name string) string {
//	var result string
//	_ = ExecuteTxn(func(txn *transaction.KVTxn) error {
//		//name := "OS/T02/Conf/IAM/policies/POLICY-244832646717280951/policy.json"
//		bytes, err := txn.Get(context.Background(), []byte(name))
//		fmt.Println(string(bytes))
//		result = string(bytes)
//		return err
//	})
//	return result
//}
//
//func set(name, value string) {
//	_ = ExecuteTxn(func(txn *transaction.KVTxn) error {
//		err := txn.Set([]byte(name), []byte(value))
//		return err
//	})
//}
//
//func del(name string) {
//	_ = ExecuteTxn(func(txn *transaction.KVTxn) error {
//		err := txn.Delete([]byte(name))
//		return err
//	})
//}

// tikv时间戳转换
func timeTest() {
	now := time.Now()
	physical := uint64(now.UnixMilli()) // 当前时间的毫秒级时间戳
	startTS := uint64(physical) << 18   // 左移 18 位（逻辑计数默认为 0）
	fmt.Printf("Current StartTS: %d\n", startTS)

	// 反向解析验证
	physical = startTS >> 18
	t := time.UnixMilli(int64(physical))
	fmt.Printf("Parsed Time: %s\n", t.Format("2006-01-02 15:04:05"))
}
