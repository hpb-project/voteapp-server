package db

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
)

type BoeNode struct {
	gorm.Model
	Coinbase   string `gorm:"column:coinbase" json:"boeAddress"`
	LockNumber uint64 `gorm:"column:locknumber" json:"lockNumber"`
	LockAddr   string `gorm:"column:lockaddr" json:"lockAddr"`
	HolderAddr string `gorm:"column:holderaddr" json:"holderAddr"`
	HID        string `gorm:"column:hid" json:"hid"`
	CID        string `gorm:"column:cid" json:"cid"`
	VoteNum    uint64 `gorm:"column:votes" json:"votes"`
}

func (dt *BoeNode) TableName() string {
	return "tb_boenode"
}

func (dt *BoeNode) NewTable() {
	orm := GetORM()
	orm.AutoMigrate(&BoeNode{})
}

func (dt *BoeNode) RefreshAll(info []*BoeNode) error {
	orm := GetORM()
	log.Debug("batch create info length ", len(info))
	if len(info) == 0 {
		return nil
	}

	t := reflect.TypeOf(info[0])
	insertList := reflect.New(reflect.SliceOf(t)).Elem()

	for _, elem := range info {
		insertList = reflect.Append(insertList, reflect.ValueOf(elem))
	}

	e := orm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("1 = 1").Delete(&BoeNode{}).Error; err != nil {
			log.Errorf("delete all failed, err:%s\n", err)
			return err
		}
		if err := tx.CreateInBatches(insertList.Interface(), 10).Error; err != nil {
			log.Errorf("insert all failed, err:%s\n", err)
			return err
		}
		return nil
	})

	if e != nil {
		log.Errorf("update all failed, err:%s\n", e)
		return e
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

func (dt *BoeNode) GetAll() (all []*BoeNode, err error) {
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

func (dt *BoeNode) Clear() error {
	orm := GetORM()
	orm.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&BoeNode{})

	return nil
}
