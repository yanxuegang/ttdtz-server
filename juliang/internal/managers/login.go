package managers

import (
	"context"
	"juliang/global"
	. "juliang/internal/models"
	key "juliang/internal/rmodels/keys"
	"juliang/pkg/app"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
)

type wxAuthResponseData struct {
	Errcode    int    `json:"errcode"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

type params struct {
	OpenId   string `json:"open_id" form:"open_id" binding:"required"`
	Type     string `json:"type" form:"type" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Channel  string `json:"channel" form:"channel" binding:"required"`
}

type LoginRequest struct {
	Cmd    int    `json:"cmd" form:"cmd" binding:"required"`
	Params params `json:"params" form:"params" binding:"required"`
}

type LoginResponseInfo struct {
	OpenId string `json:"open_id"`
	Sign   int    `json:"sign"`
	Money  int    `json:"money"`
}

func Login(ctx context.Context, req *LoginRequest) (*LoginResponseInfo, error) {
	respdata := &LoginResponseInfo{}
	playerId, err := GetPlayerIdByOpenId(req.Params.OpenId, 0)
	if err != nil {
		return respdata, err
	}
	if playerId == 0 {
		playerId, err = newPlayerId(req.Params.OpenId, 0)
		if err != nil {
			return respdata, err
		}
		player, err := createPlayer(playerId, req.Params.OpenId)
		if err != nil {
			return respdata, err
		}
		respdata.OpenId = player.OpenId
		respdata.Sign = 0
		respdata.Money = player.Money
		return respdata, nil

	}
	player, err := GetPlayerById(playerId)
	if err != nil {
		player, err := createPlayer(playerId, req.Params.OpenId)
		if err != nil {
			return respdata, err
		}
		respdata.OpenId = player.OpenId
		respdata.Sign = 0
		respdata.Money = player.Money
		return respdata, nil
	}
	respdata.OpenId = player.OpenId
	respdata.Sign = 1
	respdata.Money = player.Money
	return respdata, nil
}

func WxLogin(ctx context.Context, req *LoginRequest) (*LoginResponseInfo, error) {
	respdata := &LoginResponseInfo{}
	//wx登录验证
	resp, err := http.Get(app.GetWxAuthUrl(req.Params.OpenId))
	if err != nil {
		global.Logger.Infof("wx login err: %+v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var (
		wxrespData wxAuthResponseData
		json       = jsoniter.ConfigCompatibleWithStandardLibrary
	)

	if err2 := json.NewDecoder(resp.Body).Decode(&wxrespData); err2 != nil {
		global.Logger.Infof("wx login err: %+v", err)
		return nil, err2
	}
	global.Logger.Infof("wx login data: %+v", wxrespData)
	playerId, err := GetPlayerIdByOpenId(wxrespData.Openid, 0)
	if err != nil {
		return respdata, err
	}
	if playerId == 0 {
		playerId, err = newPlayerId(req.Params.OpenId, 0)
		if err != nil {
			return respdata, err
		}
		player, err := createPlayer(playerId, req.Params.OpenId)
		if err != nil {
			return respdata, err
		}
		respdata.OpenId = player.OpenId
		respdata.Sign = 0
		respdata.Money = player.Money
		return respdata, nil

	}
	player, err := GetPlayerById(playerId)
	if err != nil {
		player, err := createPlayer(playerId, req.Params.OpenId)
		if err != nil {
			return respdata, err
		}
		respdata.OpenId = player.OpenId
		respdata.Sign = 0
		respdata.Money = player.Money
		return respdata, nil
	}
	respdata.OpenId = player.OpenId
	respdata.Sign = 1
	respdata.Money = player.Money
	return respdata, nil
}

func newPlayerId(openid string, system uint8) (uint64, error) {
	if openid == "" {
		log.Println("WxLogin Auth ERROR openid")
		//todo errcode
		return 0, nil
	}

	// 角色数量限制
	// max, err := dbs.CacheGet(PlayerRestrictMaxFormat)
	// if err != dbs.CacheNil && err != nil {
	// 	return 0, dbs.NewCacheException(PlayerService, "Get Restrict Max Error")
	// }

	// cur, err := dbs.CacheIncr(PlayerRestrictCurFormat) // 递增记录累计创角数量
	// if err != dbs.CacheNil && err != nil {
	// 	return 0, dbs.NewCacheException(PlayerService, "Inc Restrict Cur Error")
	// }

	// if maxnum, _ := strconv.ParseInt(max, 10, 64); maxnum != 0 && cur > maxnum {
	// 	log.Printf("[Cache]Login Restrict Cur %d, Max %d, OpenID %s", cur, maxnum, openid)
	// 	dbs.CacheSAdd(PlayerRestrictBlkFormat, openid)
	// 	return 0, errors.ErrPlayerRestrict // 禁止新角色创建
	// }

	id, err := global.CacheIncr(key.PlayerGlobalIdKey)
	if err != nil {
		return 0, err
	}
	playerId := uint64(id)

	if err := saveOpenIdMap(openid, system, playerId); err != nil {
		return 0, err
	}

	// if _, err2 := dbs.CacheSetWithExpiration(fmt.Sprintf(OpenIdKeyFormat, openid, system), playerId, 0); err2 != nil {
	// 	return 0, dbs.NewCacheException(OpenIdService, "Set Openid Error")
	// }
	return playerId, nil
}

func saveOpenIdMap(openId string, system uint8, playerId uint64) error {
	db := global.GetDB("app_line")
	return db.Exec("INSERT INTO `openid_map` VALUES (?, ?, ?)", openId, system, playerId).Error
}

func createPlayer(playerId uint64, openid string) (*Player, error) {
	if openid == "" {
		return nil, nil
	}
	player := new(Player)
	player.Id = playerId
	player.OpenId = openid

	if err := player.Create(); err != nil {
		return player, err
	}
	return player, nil
}

func GetPlayerIdByOpenId(openId string, system uint8) (uint64, error) {
	//todo errcode
	if openId == "" {
		return 0, nil
	}

	var (
		openidMap = new(OpenidMap)
		db        = global.GetDB("app_line")
	)
	openidMap.Openid = openId
	openidMap.System = system

	err := db.First(&openidMap).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	} else if err == gorm.ErrRecordNotFound {
		// MySQL未查询到
		return 0, nil
	}
	return openidMap.PlayerId, nil
}

func GetPlayerById(playerId uint64) (*Player, error) {
	if playerId == 0 {
		return nil, nil
	}
	var (
		player = new(Player)
		err    error
	)
	player.Id = playerId
	err = getPlayerFromDB(player)
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return player, nil
}

func getPlayerFromDB(player *Player) error {
	return global.GetDB("app_line").First(&player).Error
}
