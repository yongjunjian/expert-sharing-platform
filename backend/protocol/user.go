package protocol

/**注册接口协议**/
type struct RegisterReq{
    UserAccount string `json:"userAccount"`
    UserName string `json:"userName"`
    PhoneNumber string `json:"phoneNumber"`
    UserPassword string `json:"userPassword"`
    ExpertDomainId int `json:"exportDomainId"`
}

type struct RegisterRes{
    CommonRes 
    UserId int `json:"userId"`
}

/**登录接口协议**/
type struct LoginReq{
    UserAccount string `json:"userAccount"`
    PhoneNumber string `json:"phoneNumber"`
    LoginMode string `json:"loginMode"` //登录方式
    UserPassword string `json:"userPassword"`
}

type struct LoginRes{
    CommonRes 
    UserId int `json:"userId"`
}

/**登出接口协议**/
type struct LogoutReq{
}

type struct LogoutRes{
    CommonRes 
}

/**用户查询协议**/
type struct QueryUserReq{
    UserId int `json:"userId"`
}

type struct QueryUserRes{
    CommonRes 
}
