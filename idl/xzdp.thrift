// idl/xzdp.thrift
namespace go xzdp

// import "github.com/hertz-contrib/sessions"

struct Empty {}

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

struct User {
    1: string Phone 
    2: string code
    3: string password 
    4: i64 id
    5: string NickName
    6: string icon
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

struct ShopType {
    1: i64 id
    2: string name
    3: string icon
    4: i32 sort
}

service UserService {
    UserResp UserMethod(1: string Id) (api.get="/me");
    UserResp UserCode(1: UserLoginFrom request) (api.post="/user/code");
    Result UserLogin(1: UserLoginFrom request) (api.post="/user/login");
    UserResp UserInfo(1: string Id) (api.get="/user/{id}");
}


service ShopService {
    list<ShopType> ShopList(1: Empty request) (api.get="/shop-type/list");
}

/*
            "id": 4,
            "shopId": 4,
            "userId": 2,
            "icon": "/imgs/icons/kkjtbcr.jpg",
            "name": "可可今天不吃肉",
            "title": "无尽浪漫的夜晚丨在万花丛中摇晃着红酒杯\uD83C\uDF77品战斧牛排\uD83E\uDD69",
            "images": "/imgs/blogs/7/14/4771fefb-1a87-4252-816c-9f7ec41ffa4a.jpg,/imgs/blogs/4/10/2f07e3c9-ddce-482d-9ea7-c21450f8d7cd.jpg,/imgs/blogs/2/6/b0756279-65da-4f2d-b62a-33f74b06454a.jpg,/imgs/blogs/10/7/7e97f47d-eb49-4dc9-a583-95faa7aed287.jpg,/imgs/blogs/1/2/4a7b496b-2a08-4af7-aa95-df2c3bd0ef97.jpg,/imgs/blogs/14/3/52b290eb-8b5d-403b-8373-ba0bb856d18e.jpg",
            "content": "生活就是一半烟火·一半诗意<br/>手执烟火谋生活·心怀诗意以谋爱·<br/>当然<br/>\r\n男朋友给不了的浪漫要学会自己给\uD83C\uDF52<br/>\n无法重来的一生·尽量快乐.<br/><br/>\uD83C\uDFF0「小筑里·神秘浪漫花园餐厅」\uD83C\uDFF0<br/><br/>\n\uD83D\uDCAF这是一家最最最美花园的西餐厅·到处都是花餐桌上是花前台是花  美好无处不在\n品一口葡萄酒，维亚红酒马瑟兰·微醺上头工作的疲惫消失无际·生如此多娇\uD83C\uDF43<br/><br/>\uD83D\uDCCD地址:延安路200号(家乐福面)<br/><br/>\uD83D\uDE8C交通:地铁①号线定安路B口出右转过下通道右转就到啦～<br/><br/>--------------\uD83E\uDD63菜品详情\uD83E\uDD63---------------<br/><br/>「战斧牛排]<br/>\n超大一块战斧牛排经过火焰的炙烤发出阵阵香，外焦里嫩让人垂涎欲滴，切开牛排的那一刻，牛排的汁水顺势流了出来，分熟的牛排肉质软，简直细嫩到犯规，一刻都等不了要放入嘴里咀嚼～<br/><br/>「奶油培根意面」<br/>太太太好吃了\uD83D\uDCAF<br/>我真的无法形容它的美妙，意面混合奶油香菇的香味真的太太太香了，我真的舔盘了，一丁点美味都不想浪费‼️<br/><br/><br/>「香菜汁烤鲈鱼」<br/>这个酱是辣的 真的绝好吃‼️<br/>鲈鱼本身就很嫩没什么刺，烤过之后外皮酥酥的，鱼肉蘸上酱料根本停不下来啊啊啊啊<br/>能吃辣椒的小伙伴一定要尝尝<br/><br/>非常可 好吃子\uD83C\uDF7D\n<br/>--------------\uD83C\uDF43个人感受\uD83C\uDF43---------------<br/><br/>【\uD83D\uDC69\uD83C\uDFFB‍\uD83C\uDF73服务】<br/>小姐姐特别耐心的给我们介绍彩票 <br/>推荐特色菜品，拍照需要帮忙也是尽心尽力配合，太爱他们了<br/><br/>【\uD83C\uDF43环境】<br/>比较有格调的西餐厅 整个餐厅的布局可称得上的万花丛生 有种在人间仙境的感觉\uD83C\uDF38<br/>集美食美酒与鲜花为一体的风格店铺 令人向往<br/>烟火皆是生活 人间皆是浪漫<br/>",
            "liked": 1,
            "comments": 104,
            "createTime": "2021-12-28T19:50:01",
            "updateTime": "2022-03-10T14:26:34"
        },
*/


struct Blog {
    1: i64 id
    2: i64 shopId
    3: i64 userId
    4: string icon
    5: string name
    6: string title
    7: string images
    8: string content
    9: i64 liked
    10: i64 comments
    11: string createTime
    12: string updateTime
}

struct BlogReq {
    1: i64 current (api.query="current");
}
service BlogSerivice {
    list<Blog> GetHotBlog(1: BlogReq request) (api.get="/blog/hot");
}