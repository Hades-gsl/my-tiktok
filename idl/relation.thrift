include 'user.thrift'

namespace go relation

struct ActionRequest {
    1: string token, // 用户鉴权token
    2: i64 to_user_id, // 对方用户id
    3: i32 action_type // 1-关注，2-取消关注
}

struct ActionResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

struct FollowListRequest {
    1: i64 user_id, // 用户id
    2: string token // 用户鉴权token
}

struct FollowListResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: list<user.User> user_list // 用户信息列表
}

struct FollowerListRequest {
    1: i64 user_id, // 用户id
    2: string token // 用户鉴权token
}

struct FollowerListResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: list<user.User> user_list // 用户列表
}

struct FriendListRequest {
    1: i64 user_id, // 用户id
    2: string token // 用户鉴权token
}

struct FriendListResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: list<FriendUser> user_list // 用户列表
}

struct FriendUser {
    1: optional string message, // 和该好友的最新聊天消息
    2: required i64 msgType // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
    3: user.User user_info // 嵌套的User结构
}

service RelationService{
    ActionResponse Action(1: ActionRequest req),
    FollowListResponse FollowList(1: FollowListRequest req),
    FollowerListResponse FollowerList(1: FollowerListRequest req),
    FriendListResponse FriendList(1: FriendListRequest req),
}