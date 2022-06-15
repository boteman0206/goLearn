package main

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/panjf2000/ants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type DownLoad struct {
	ProductId  string `json:"product_id"`
	ThirdSkuId string `json:"third_sku_id"`
	ImageUrl   string `json:"image_url"`
}

var (

	// 配置环境的url
	Url      = ""
	AntsP, _ = ants.NewPool(1000)
)

func main() {

	data, err := getData()
	if err != nil {
		return
	}

	wait := sync.WaitGroup{}

	for i := range data {
		wait.Add(1)

		value := data[i]
		fmt.Println(value)
		AntsP.Submit(func() {
			defer func() {
				wait.Done()
			}()
			split := strings.Split(value.ImageUrl, ",")
			fmt.Println("split pic : ", split)

			err := os.MkdirAll("./data/img/"+value.ThirdSkuId, os.ModePerm)
			if err != nil {
				fmt.Println("os.MkdirAll error: ", err.Error())
				return
			}

			for k := range split {
				if len(strings.TrimSpace(split[k])) > 0 {
					//保存图片
					fmt.Println("正在保存的图片:", split[k])
					newUUID, err := uuid.NewUUID()
					if err != nil {
						fmt.Println("uuid.NewUUID() error: ", err.Error())
						continue
					}
					f, err := os.Create("./data/img/" + value.ThirdSkuId + "/" + newUUID.String() + ".jpg")
					if err != nil {
						fmt.Println("os.Create error: ", err.Error())
						continue
					}

					r, err := http.Get(split[k])
					if err != nil {
						fmt.Println(" http.Get error: ", err.Error())
						continue
					}
					body, _ := ioutil.ReadAll(r.Body)
					io.Copy(f, bytes.NewReader(body))

					//释放资源
					r.Body.Close()
					f.Close()
				}
			}

		})
	}

	wait.Wait()

	fmt.Println("全部下载成功")

}

func getData() ([]DownLoad, error) {

	gormConf := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(Url), gormConf)

	fmt.Println("获取数据库链接：", db, err)

	data := []DownLoad{}
	err = db.Raw("select p.id product_id,st.third_sku_id third_sku_id, p.pic image_url  from dc_product.product p join dc_product.sku_third st on p.id =st.product_id " +
		"where st.erp_id = 2 and length(pic) > 0 and length(st.third_sku_id) > 0  limit 10 ").Find(&data).Error
	if err != nil {
		fmt.Println("err : ", err.Error())
	}
	//
	fmt.Println("数据量大小： ", len(data))
	time.Sleep(5 * time.Second)
	return data, err

}
