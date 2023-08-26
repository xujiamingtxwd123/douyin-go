package conf

import "strings"

const (
	//特殊权限-白名单
	TRIAL_WHITELIST = "trial.whitelist"

	//特殊权限-增量授权
	INCREMENTAL_AUTHORIZATION = "incremental_authorization"

	//视频权限-查询授权账号的视频列表
	VIDEO_LIST_BIND = "video.list.bind"

	//用户权限-授权登录能力
	USER_INFO = "user_info"

	//用户权限 - 粉丝判断
	FANS_CHECK = "fans.check"

	//用户权限 - 授权动态续期
	RENEW_REFRESH_TOKEN = "renew_refresh_token"

	//数据权限-视频数据
	DATA_EXTERNAL_ITEM = "data.external.item"

	//数据权限-抖音影视综榜单数据
	DISCOVERY_ENT = "discovery.ent"

	//数据权限-星图数据
	STAR_TOP_SCORE_DISPLAY = "star_top_score_display"

	STAR_TOPS = "star_tops"

	STAR_AUTHOR_SCORE_DISPLAY = "star_author_score_display"

	//SDK分享视频数据
	DATA_EXTERNAL_SDK_SHARE = "data.external.sdk_share"

	//抖音热点
	HOTSEARCH = "hotsearch"

	//达人榜单-游戏
	DATA_EXTERNAL_BILLBOARD_GAME = "data.external.billboard_game"

	//达人榜单-剧情
	DATA_EXTERNAL_BILLBOARD_DRAMA = "data.external.billboard_drama"

	//达人榜单-汽车
	DATA_EXTERNAL_BILLBOARD_CAR = "data.external.billboard_car"

	//达人榜单-搞笑
	DATA_EXTERNAL_BILLBOARD_AMUSEMENT = "data.external.billboard_amusement"

	//达人榜单-二次元
	DATA_EXTERNAL_BILLBOARD_COSPA = "data.external.billboard_cospa"

	//达人榜单-美食
	DATA_EXTERNAL_BILLBOARD_FOOD = "data.external.billboard_food"

	//达人榜单-旅游
	DATA_EXTERNAL_BILLBOARD_TRAVEL = "data.external.billboard_travel"

	//达人榜单-娱乐明星
	DATA_EXTERNAL_BILLBOARD_STARS = "data.external.billboard_stars"

	//达人榜单-体育
	DATA_EXTERNAL_BILLBOARD_SPORT = "data.external.billboard_sport"

	//话题榜单数据
	DATA_EXTERNAL_BILLBOARD_TOPIC = "data.external.billboard_topic"

	//道具榜单数据
	DATA_EXTERNAL_BILLBOARD_PROP = "data.external.billboard_prop"

	//音乐榜单数据
	DATA_EXTERNAL_BILLBOARD_MUSIC = "data.external.billboard_music"

	//直播榜单数据
	DATA_EXTERNAL_BILLBOARD_LIVE = "data.external.billboard_live"

	//热门视频数据
	DATA_EXTERNAL_BILLBOARD_HOT_VIDEO = "data.external.billboard_hot_video"

	//查询用户的粉丝数据能力
	FANS_DATA_BIND     = "fans.data.bind"
	DATA_EXTERNAL_USER = "data.external.user"
	LIVE_ROOM_BASE     = "live.room.base"
	LIVE_ROOM_AUDIENCE = "live.room.audience"
)

func AllAuthScope() string {
	scopes := []string{
		TRIAL_WHITELIST,
		INCREMENTAL_AUTHORIZATION,

		VIDEO_LIST_BIND,
		//用户权限-授权登录能力
		USER_INFO,

		//用户权限 - 粉丝判断
		FANS_CHECK,

		//用户权限 - 授权动态续期
		RENEW_REFRESH_TOKEN,

		//数据权限-视频数据
		DATA_EXTERNAL_ITEM,

		//数据权限-抖音影视综榜单数据
		DISCOVERY_ENT,

		//数据权限-星图数据
		STAR_TOP_SCORE_DISPLAY,

		STAR_TOPS,

		STAR_AUTHOR_SCORE_DISPLAY,

		//SDK分享视频数据
		DATA_EXTERNAL_SDK_SHARE,

		//抖音热点
		HOTSEARCH,

		//达人榜单-游戏
		DATA_EXTERNAL_BILLBOARD_GAME,

		//达人榜单-剧情
		DATA_EXTERNAL_BILLBOARD_DRAMA,

		//达人榜单-汽车
		DATA_EXTERNAL_BILLBOARD_CAR,

		//达人榜单-搞笑
		DATA_EXTERNAL_BILLBOARD_AMUSEMENT,

		//达人榜单-二次元
		DATA_EXTERNAL_BILLBOARD_COSPA,

		//达人榜单-美食
		DATA_EXTERNAL_BILLBOARD_FOOD,

		//达人榜单-旅游
		DATA_EXTERNAL_BILLBOARD_TRAVEL,

		//达人榜单-娱乐明星
		DATA_EXTERNAL_BILLBOARD_STARS,

		//达人榜单-体育
		DATA_EXTERNAL_BILLBOARD_SPORT,

		//话题榜单数据
		DATA_EXTERNAL_BILLBOARD_TOPIC,

		//道具榜单数据
		DATA_EXTERNAL_BILLBOARD_PROP,

		//音乐榜单数据
		DATA_EXTERNAL_BILLBOARD_MUSIC,

		//直播榜单数据
		DATA_EXTERNAL_BILLBOARD_LIVE,

		//热门视频数据
		DATA_EXTERNAL_BILLBOARD_HOT_VIDEO,

		//查询用户的粉丝数据能力
		FANS_DATA_BIND,
	}
	return strings.Join(scopes, ",")
}
