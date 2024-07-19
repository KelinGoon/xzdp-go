// idl/xzdp.thrift
namespace go xzdp

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}

struct UserReq {
    1: string Name (api.query="name");
}

struct UserResp {
    1: string RespBody;
}


service UserService {
    UserResp UserMethod(1: UserReq request) (api.get="/user");
    UserResp UserLogin(1: UserReq request) (api.post="/user/login");
}