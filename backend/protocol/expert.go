package protocol

/**专家管理接口**/
type struct AddExpertReq{
    UserId int `json:"userId"`
    DomainId int `json:"domainId"`
}
type struct AddExpertRes{
   CommonRes
   ExpertId int `json:"expertId"`
}

type struct DelExpertReq{
    UserId int `json:"userId"`
}
type DelExpertRes CommonRes

type struct SearchExpertReq{
    ExpertIdList []int `json:"exportIdList"`
    UserId int `json:"userId"`
    DomainId int `json:"domainId"`
    MinFortune int `json:"minFortune"`
    MaxFortune int `json:"maxFortune"`
    MinExpertLevel int `json:"minExpertLevel"`
    MaxExpertLevel int `json:"maxExpertLevel"`
}
type struct SearchExpertRes{
   CommonRes
   expertList []ExpertInfo `json:"data"`
}

