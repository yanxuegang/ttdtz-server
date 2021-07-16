package models

import (
	"log"
	"time"
	"ttdtz-server/global"

	"github.com/jinzhu/gorm"
)

type Player struct {
	Id                  uint64    `gorm:"column:id" db:"id" json:"id" form:"id"`
	OpenId              string    `gorm:"column:open_id" db:"open_id" json:"open_id" form:"open_id"`                                                             //微信openId(账号)
	Type                string    `gorm:"column:type" db:"type" json:"type" form:"type"`                                                                         //来源
	Name                string    `gorm:"column:name" db:"name" json:"name" form:"name"`                                                                         //自定义昵称
	Head                string    `gorm:"column:head" db:"head" json:"head" form:"head"`                                                                         //自定义头像
	NickName            string    `gorm:"column:nick_name" db:"nick_name" json:"nick_name" form:"nick_name"`                                                     //微信昵称
	AvatarUrl           string    `gorm:"column:avatar_url" db:"avatar_url" json:"avatar_url" form:"avatar_url"`                                                 //微信头像
	Gender              int       `gorm:"column:gender" db:"gender" json:"gender" form:"gender"`                                                                 //性别：1男 2女(0未知)
	City                string    `gorm:"column:city" db:"city" json:"city" form:"city"`                                                                         //城市
	Province            string    `gorm:"column:province" db:"province" json:"province" form:"province"`                                                         //省份
	Country             string    `gorm:"column:country" db:"country" json:"country" form:"country"`                                                             //国家
	CreatedAt           time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                                                 //注册时间
	LoginAt             time.Time `gorm:"column:login_at" db:"login_at" json:"login_at" form:"login_at"`                                                         //登录时间
	IsFirst             int       `gorm:"column:is_first" db:"is_first" json:"is_first" form:"is_first"`                                                         //当天登陆(0:未登陆,1:登陆)
	LoginDays           int       `gorm:"column:login_days" db:"login_days" json:"login_days" form:"login_days"`                                                 //登陆天数
	Level               int       `gorm:"column:level" db:"level" json:"level" form:"level"`                                                                     //玩家等级
	Exp                 int       `gorm:"column:exp" db:"exp" json:"exp" form:"exp"`                                                                             //经验值
	Gold                int       `gorm:"column:gold" db:"gold" json:"gold" form:"gold"`                                                                         //玩家金币
	Money               int       `gorm:"column:money" db:"money" json:"money" form:"money"`                                                                     //玩家元宝
	ChapterId           int       `gorm:"column:chapter_id" db:"chapter_id" json:"chapter_id" form:"chapter_id"`                                                 //当前章节id
	ChapterLevel        int       `gorm:"column:chapter_level" db:"chapter_level" json:"chapter_level" form:"chapter_level"`                                     //当前章节难度等级
	IsModifyNickName    int       `gorm:"column:is_modify_nick_name" db:"is_modify_nick_name" json:"is_modify_nick_name" form:"is_modify_nick_name"`             //是否已经修改过名字(0:未修改,1:修改)
	IsNewUserPass       int       `gorm:"column:is_new_user_pass" db:"is_new_user_pass" json:"is_new_user_pass" form:"is_new_user_pass"`                         //是否通过新手引导(0:未通过,1:通过)
	NewUserGuide        string    `gorm:"column:new_user_guide" db:"new_user_guide" json:"new_user_guide" form:"new_user_guide"`                                 //新手引导步骤
	NewUserChapter      string    `gorm:"column:new_user_chapter" db:"new_user_chapter" json:"new_user_chapter" form:"new_user_chapter"`                         //新手通章
	SystemOpen          string    `gorm:"column:system_open" db:"system_open" json:"system_open" form:"system_open"`                                             //已开放系统(模块)
	SystemOpenNewuserAt time.Time `gorm:"column:system_open_newuser_at" db:"system_open_newuser_at" json:"system_open_newuser_at" form:"system_open_newuser_at"` //新手活动开启时间
	TodayOnlineAt       time.Time `gorm:"column:today_online_at" db:"today_online_at" json:"today_online_at" form:"today_online_at"`                             //今天在线时间
	TodayOnlineSecond   int       `gorm:"column:today_online_second" db:"today_online_second" json:"today_online_second" form:"today_online_second"`             //今天在线时间秒数
	RechargeStatus      int       `gorm:"column:recharge_status" db:"recharge_status" json:"recharge_status" form:"recharge_status"`                             //首充奖励(0:未充值,1:已充值,2:已领取)
	AdStatus            int       `gorm:"column:ad_status" db:"ad_status" json:"ad_status" form:"ad_status"`                                                     //广告奖励(0:未满足,1:未(可)领取,2:已领取)
	AdStatusAt          time.Time `gorm:"column:ad_status_at" db:"ad_status_at" json:"ad_status_at" form:"ad_status_at"`                                         //广告状态领取时间(关羽)
	NormalChapterNum    int       `gorm:"column:normal_chapter_num" db:"normal_chapter_num" json:"normal_chapter_num" form:"normal_chapter_num"`                 //普通关进副本次数
	EliteChapterNum     int       `gorm:"column:elite_chapter_num" db:"elite_chapter_num" json:"elite_chapter_num" form:"elite_chapter_num"`                     //精英关进副本次数
	KillMonstersNum     int       `gorm:"column:kill_monsters_num" db:"kill_monsters_num" json:"kill_monsters_num" form:"kill_monsters_num"`                     //总杀怪数量
	MonstersPacketNum   int       `gorm:"column:monsters_packet_num" db:"monsters_packet_num" json:"monsters_packet_num" form:"monsters_packet_num"`             //怪物红包数量
	VideoNum            int       `gorm:"column:video_num" db:"video_num" json:"video_num" form:"video_num"`                                                     //看视频次数
	RedPacketNum        int       `gorm:"column:red_packet_num" db:"red_packet_num" json:"red_packet_num" form:"red_packet_num"`                                 //红包领取次数
	SignMoneyNum        int       `gorm:"column:sign_money_num" db:"sign_money_num" json:"sign_money_num" form:"sign_money_num"`                                 //签到红包领取次数
	SuperRedPacketNum   int       `gorm:"column:super_red_packet_num" db:"super_red_packet_num" json:"super_red_packet_num" form:"super_red_packet_num"`         //超级红包(3000)(6次)
	SuperRedPacketAt    time.Time `gorm:"column:super_red_packet_at" db:"super_red_packet_at" json:"super_red_packet_at" form:"super_red_packet_at"`             //超级红包领取日期
}

func (m *Player) GetDB() *gorm.DB {
	db := global.GetDB("app_line")
	if db == nil {
		log.Println("[Error] Cannot connect DB")
		return nil
	}
	return db
}

func (m *Player) Create() error {
	//todo redis save
	return m.GetDB().Create(m).Error
}

func (m *Player) Update(attrs ...interface{}) error {
	//todo redis save
	rowsAffected := m.GetDB().Model(m).Update(attrs...).RowsAffected
	if rowsAffected == 0 {
		log.Println("[PLAYER][WARNING] No Content To Update.")
		return nil
	}
	return nil
}
