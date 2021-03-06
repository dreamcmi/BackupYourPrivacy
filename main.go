package main

import (
	"BackupYourPrivacy/config"
	"BackupYourPrivacy/sql"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func perr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	cg := config.Config{}
	cg, err = config.InitConfig()
	perr(err)

	err = sql.InitSql()
	perr(err)

	client, err := oss.New(cg.Oss.EndPoint, cg.Oss.AccessKeyID, cg.Oss.AccessKeySecret)
	perr(err)

	bucket, err := client.Bucket(cg.Oss.Bucket)
	perr(err)

	lsRes, err := bucket.ListObjects()
	perr(err)

	for _, object := range lsRes.Objects {
		fmt.Println("Objects:", object.Key)
	}
}
