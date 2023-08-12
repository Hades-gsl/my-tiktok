namespace go user

struct User {
    1: i64 id, // 用户id
    2: string name, // 用户名称
    3: optional i64 follow_count, // 关注总数
    4: optional i64 follower_count, // 粉丝总数
    5: required bool is_follow, // true-已关注，false-未关注
    6: optional string avatar, // 用户头像
    7: optional string background_image, // 用户个人页顶部大图
    8: optional string signature, // 个人简介
    9: optional i64 total_favorited, // 获赞数量
    10: optional i64 work_count, // 作品数量
    11: optional i64 favorite_count // 点赞数量
}

struct RegisterRequest {
    1: string username, // 注册用户名，最长32个字符
    2: string password // 密码，最长32个字符
}

struct RegisterResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: i64 user_id, // 用户id
    4: string token // 用户鉴权token
}

struct LoginRequest {
    1: string username, // 登录用户名
    2: string password // 登录密码
}

struct LoginResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: i64 user_id, // 用户id
    4: string token // 用户鉴权token
}

struct InfoRequest {
    1: i64 user_id, // 用户id
    2: string token // 用户鉴权token
}

struct InforResponse {
    1: i32 status_code, // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: User user // 用户信息
}

service UserService{
    RegisterResponse Register(1: RegisterRequest req),
    LoginResponse Login(1: LoginRequest req),
    InforResponse Info(1: InfoRequest req),
}