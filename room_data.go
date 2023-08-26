package douyingo

import (
	"context"

	"github.com/xujiamingtxwd123/douyin-go/conf"
)

type BaseErr struct {
	ErrMsg string `json:"err_msg"` //错误码描述,正常为0
	ErrNo  int64  `json:"err_no"`  //错误码
	LogID  string `json:"log_id"`
}

// DataExternalRoomIDsReq 获取房间号列表请求
type DataExternalRoomIDsReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Startime    int64  // 查询开始时间（秒级时间戳）
	EndTime     int64  // 查询结束时间（秒级时间戳）
}

type DataExternalRoomIDsRes struct {
	RoomIDs []int64 `json:"room_ids"`
	BaseErr
}

// DataExternalRoomBaseReq 获取直播间基础数据请求
type DataExternalRoomBaseReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	LiveID      int64  // 业务线id： 1-抖火 3-西瓜头条
	RoomID      int64  // 房间id
}

// DataExternalRoomBaseRes 获取直播间基础数据响应
type DataExternalRoomBaseRes struct {
	Data DataExternalRoomBaseData `json:"data"`
	BaseErr
}

type DataExternalRoomBaseData struct {
	Stats []*DataExternalRoomBaseStat `json:"stats"`
}

type DataExternalRoomBaseStat struct {
	Item     string `json:"item"`      //房间id
	ItemType int64  `json:"item_type"` //3-房间，2-主播，1-用户
	LiveID   int64  `json:"live_id"`   //1-抖火 3-西瓜头条
	Name     string `json:"name"`      //指标名
	Value    int64  `json:"value"`     //数据
	Timeslot string `json:"timeslot"`  //时间槽
}

// DataExternalRoomAudienceReq 获取直播间看播数据请求
type DataExternalRoomAudienceReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	LiveID      int64  // 业务线id： 1-抖火 3-西瓜头条
	RoomID      int64  // 房间id
}

// DataExternalRoomBaseRes 获取直播间基础数据响应
type DataExternalRoomAudienceRes struct {
	Data DataExternalRoomBaseData `json:"data"`
	BaseErr
}

// DataExternalRoomIDS 获取主播房间号
func (m *Manager) DataExternalRoomIDs(req DataExternalRoomIDsReq) (res DataExternalRoomIDsRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?access_token=%s&open_id=%s&start_time=%d&end_time=%d",
			conf.API_ROOM_DATA_ROOM_ID_GET, req.AccessToken, req.OpenId, req.Startime, req.EndTime), nil, nil)
	return res, err
}

// DataExternalRoomBase 获取直播间基础数据
func (m *Manager) DataExternalRoomBase(req DataExternalRoomBaseReq) (res DataExternalRoomAudienceRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?access_token=%s&open_id=%s&live_id=%d&room_id=%d",
			conf.API_ROOM_DATA_BASE_GET, req.AccessToken, req.OpenId, req.LiveID, req.RoomID), nil, nil)
	return res, err
}

// DataExternalRoomAudience 获取直播间看播数据
func (m *Manager) DataExternalRoomAudience(req DataExternalRoomAudienceReq) (res DataExternalRoomBaseRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?access_token=%s&open_id=%s&live_id=%d&room_id=%d",
			conf.API_ROOM_DATA_AUDIENCE_GET, req.AccessToken, req.OpenId, req.LiveID, req.RoomID), nil, nil)
	return res, err
}
