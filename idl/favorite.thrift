include 'feed.thrift'

namespace go favorite

struct ActionRequest {
    1: i64 user_id, // 用户id
    2: i64 video_id, // 视频id
    3: i32 action_type // 1-点赞，2-取消点赞
}

struct ActionResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

struct ListRequest {
    1: i64 actor_id, // 观看者id
    2: i64 user_id // 用户id
}

struct ListResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: list<feed.Video> video_list // 用户点赞视频列表
}

service FavoriteService{
    ActionResponse Action(1: ActionRequest req),
    ListResponse List(1: ListRequest req),
}