package db

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type NodeName struct {
	gorm.Model
	Name     string `gorm:"column:name" json:"nodeName"`
	Coinbase string `gorm:"column:coinbase" json:"boeAddress"`
}

func (dt *NodeName) TableName() string {
	return "tb_nodename"
}

func (dt *NodeName) Create() error {
	orm := GetORM()

	if err := orm.Model(&NodeName{}).Create(dt).Error; err != nil {
		log.WithError(err).Errorln("db create NodeName error", log.Fields{
			"record": dt,
		})
		return errors.Wrap(err, "db create NodeName error")
	}

	return nil
}

func (dt *NodeName) GetNameByCoinbase(coinbase string) (string, error) {
	orm := GetORM()
	var result NodeName

	if err := orm.Model(&NodeName{}).Where("coinbase = ?", coinbase).First(&result).Error; err != nil {
		log.Errorf("GetNameByCoinbase error:%s, %s\n",
			err.Error(), log.Fields{"coinbase": coinbase})

		return "", errors.Wrap(err, "GetNameByCoinbase error")
	}

	return result.Name, nil
}

func (dt *NodeName) GetAllInfo() (all []*NodeName, err error) {
	orm := GetORM()

	if err = orm.Model(&NodeName{}).Find(&all).Error; err != nil {
		log.Errorf("GetAllInfo error:%s\n", err.Error())
		return all, errors.Wrap(err, "GetAllInfo error")
	}

	return all, nil
}
