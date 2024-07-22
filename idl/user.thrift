// idl/user.thrift
namespace go user

struct Session {
    1: string Token;
}
struct UserLoginFrom {
    1: string Phone (api.query="phone");
    2: string code (api.query="code");
    3: string password (api.query="password");
    4: Session session (api.query="session");
}

struct User {
    1: string Phone (api.query="phone");
    2: string code
    3: string password 
    4: i64 id
    5: string NickName
    6: string icon
}

struct Result {
    1: bool success,
    2: optional string errorMsg,
    3: optional string data, // 使用 string 类型来表示泛型对象。你可以根据需要选择合适的数据类型。
    4: optional i64 total
}

struct UserResp {
    1: string RespBody;
}

struct ShopType {
    1: i64 id
    2: string name
    3: string icon
    4: i32 sort
}

service UserService {
    UserResp UserMethod(1: string request) (api.get="/user/me");
    UserResp SendCode(1: UserLoginFrom request) (api.post="/user/code");
    UserResp UserLogin(1: UserLoginFrom request) (api.post="/user/login");
    UserResp UserInfo(1: UserLoginFrom request) (api.get="/user/:id");
}

