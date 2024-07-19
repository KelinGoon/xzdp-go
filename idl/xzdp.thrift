// idl/xzdp.thrift
namespace go xzdp

// import "github.com/hertz-contrib/sessions"

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}

struct Session {
    1: string Token;
}
struct UserLoginFrom {
    1: string Phone (api.query="phone");
    2: string code (api.query="code");
    3: string password (api.query="password");
    4: Session session (api.query="session");
}

struct Result {
    1: bool success,
    2: string errorMsg,
    3: optional string data, // 使用 string 类型来表示泛型对象。你可以根据需要选择合适的数据类型。
    4: optional i64 total
}

struct UserResp {
    1: string RespBody;
}


service UserService {
    UserResp UserMethod(1: string Id) (api.get="/me");
    Result UserLogin(1: UserLoginFrom request) (api.post="/user/login");
    UserResp UserInfo(1: string Id) (api.get="/user/{id}");
}