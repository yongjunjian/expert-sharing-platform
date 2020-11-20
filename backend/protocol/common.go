package protocol

/**响应数据的公共接口**/
type struct CommonRes{
    Message string `json:"message"`
    Code int `json:"message"`
}

/**session结构**/
type struct Session{
    UserId int `json:"userId"`
}

/**专家信息**/
type struct ExpertInfo{
    ExpertId int `json:"exportId"`
    UserId int `json:"userId"`
    DomainId int `json:"domainId"`
    Fortune int `json:"fortune"`
    ExpertLevel int `json:"expertLevel"`
}

/**任务信息**/
type struct TaskInfo{
    DomainId int `json:"domainId"`
    Bonus int `json:"bonus"` 
    Fine int `json:"fine"` 
    Deadline string `json:"deadline"` //最大截止日期,yyyy-MM-DD
    ExportId int `json:"exportId"`
    Status int `json:"status"`
}

/**方案信息**/
type struct SolutionInfo{
    Solution int `json:"solutionId"`
    DomainId int `json:"domainId"`
    TaskId int `json:"taskId"`
    ProblemDesc string `json:"problemDesc"` 
    Abstract string `json:"abstract"` 
    Detail string `json:"detail"`
    CommunicateInfo string `json:"communicateInfo"` 
    PublisherId int `publisherId` 
    PublisherPermitted bool `json:"publisherPermitted"` 
    ExpertId int `json:"expertId"` 
    SolverPermitted bool `json:"solverPermitted"`
    Price int `json:"price"` 
    ShareType int `json:"shareType"` 
}
