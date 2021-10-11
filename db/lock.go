package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type LockNode struct {
	gorm.Model
	Name     string `gorm:"column:name" json:"LockNode"`
	Coinbase string `gorm:"column:coinbase" json:"boeAddress"`
}

func (dt *LockNode) Create() error {
	orm := GetORM()

	if err := orm.Model(&LockNode{}).Create(dt).Error; err != nil {
		log.WithError(err).Errorln("db create LockNode error", log.Fields{
			"record": dt,
		})
		return errors.Wrap(err, "db create LockNode error")
	}

	return nil
}

func (dt *LockNode) GetNameByCoinbase(coinbase string) (string, error) {
	orm := GetORM()
	var result LockNode

	if err := orm.Model(&LockNode{}).Where("coinbase = ?", coinbase).First(&result).Error; err != nil {
		log.Errorf("GetNameByCoinbase error:%s, %s\n",
			err.Error(), log.Fields{"coinbase": coinbase})

		return "", errors.Wrap(err, "GetNameByCoinbase error")
	}

	return result.Name, nil
}

func (dt *LockNode) GetAllInfo() (all []*LockNode, err error) {
	orm := GetORM()

	if err = orm.Model(&LockNode{}).Find(&all).Error; err != nil {
		log.Errorf("GetAllInfo error:%s\n", err.Error())
		return all, errors.Wrap(err, "GetAllInfo error")
	}

	return all, nil
}
