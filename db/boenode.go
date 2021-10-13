package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type BoeNode struct {
	gorm.Model
	Coinbase   string   `gorm:"column:coinbase" json:"boeAddress"`
	LockNumber *big.Int `gorm:"column:locknumber" json:"lockNumber"`
	LockAddr   string   `gorm:"column:lockaddr" json:"lockAddr"`
	HolderAddr string   `gorm:"column:holderaddr" json:"holderAddr"`
	HID        string   `gorm:"column:hid" json:"hid"`
	CID        string   `gorm:"column:cid" json:"cid"`
	VoteNum    *big.Int `gorm:"column:votes" json:"votes"`
}

func (dt *BoeNode) NewTable() {
	orm := GetORM()
	orm.AutoMigrate(&BoeNode{})
}

func (dt *BoeNode) BatchCreate(info []*BoeNode) error {
	orm := GetORM()
	log.Debug("batch create info length ", len(info))
	if err := orm.Model(&BoeNode{}).Create(&info).Error; err != nil {
		log.WithError(err).Errorln("db batch create BoeNode error", log.Fields{
			"len(info)": len(info),
		})
		return errors.Wrap(err, "db batch create BoeNode error")
	}

	return nil
}

func (dt *BoeNode) Create() error {
	orm := GetORM()

	if err := orm.Model(&BoeNode{}).Create(dt).Error; err != nil {
		log.WithError(err).Errorln("db create BoeNode error", log.Fields{
			"record": dt,
		})
		return errors.Wrap(err, "db create BoeNode error")
	}

	return nil
}

func (dt *BoeNode) GetByCoinbase(coinbase string) error {
	orm := GetORM()

	if err := orm.Model(&BoeNode{}).Where("coinbase = ?", coinbase).First(dt).Error; err != nil {
		log.Errorf("GetByCoinbase error:%s, %s\n",
			err.Error(), log.Fields{"coinbase": coinbase})

		return errors.Wrap(err, "GetNameByCoinbase error")
	}

	return nil
}

func (dt *BoeNode) GetAllNode() (all []*BoeNode, err error) {
	orm := GetORM()

	if err = orm.Model(&BoeNode{}).Find(&all).Error; err != nil {
		log.Errorf("GetAllInfo error:%s\n", err.Error())
		return all, errors.Wrap(err, "GetAllInfo error")
	}

	return all, nil
}

func (dt *BoeNode) Update() error {

	orm := GetORM()
	if err := orm.Model(dt).Save(dt).Error; err != nil {
		log.Errorf("Update boenode error:%s\n", err.Error())
		return errors.Wrap(err, "db boenode update error")
	}
	return nil
}

func (dt *BoeNode) Drop() error {
	orm := GetORM()
	orm.DropTable(&BoeNode{})
	return nil
}
