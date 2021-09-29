package db

import (
	"encoding/json"
	"fmt"
	"github.com/hpb-project/votedapp-server/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"testing"
)

func init() {
	config.Init()
	conf := config.GetConfig()
	conf.DBUserName = "root"
	conf.DBUserPassword = "bXlzcWw="
	conf.DBIP = "1.15.85.191"
	conf.DBPort = 3306
	conf.DBName = "hpbvote"
}

func initTable(t *testing.T) error {
	err := Init()
	if err != nil {
		t.Fatal("open database failed, err ", err)
		return err
	}
	gorm := GetORM()
	if gorm.HasTable(&NodeName{}) {
		gorm.AutoMigrate(&NodeName{})
	} else {
		gorm.CreateTable(&NodeName{})
	}
	return nil
}


func TestNodeName_Create(t *testing.T) {
	initTable(t)

	datas := make([]*NodeName, 0)
	list,err := ioutil.ReadFile("nodelist.json")
	if err != nil {
		fmt.Println("open nodelist failed.")
		return
	}
	err = json.Unmarshal(list, &datas)
	if err != nil {
		fmt.Printf("unmarshal nodelist failed, err:%s.\n", err)
		return
	}

	for _, d := range datas {
		if name,err := d.GetNameByCoinbase(d.Coinbase); err == nil {
			fmt.Printf("coinbase %s existed, name is %s.\n", d.Coinbase, name)
			continue
		}
		if err := d.Create(); err != nil {
			t.Fatal("create failed, err ", err)
		} else {
			fmt.Println("create new passed ")
		}
	}
	//gorm.DropTable(&NodeName{})
}

func TestNodeName_GetAllInfo(t *testing.T) {
	var d = &NodeName{}
	all, err := d.GetAllInfo()
	if err != nil {
		fmt.Printf("get all failed.\n")
	}
	for _, info := range all {
		fmt.Printf("get node name(%s), coinbase(%s)\n", info.Name, info.Coinbase)
	}
}

func TestNodeName_CloseDB(t *testing.T) {
	GetORM().Close()
}
