package protocol

/**添加任务**/
type struct AddTaskReq{
    DomainId int `json:"domainId"`
    ProblemDesc string `json:"problemDesc"`
    StandardForSolution string `json:"standardForSolution"`
    Bonus int `json:"bonus"` //奖金
    Fine int `json:"fine"` //罚金
    Deadline string `json:"deadline"` //截止日期,yyyy-MM-DD
    PublisherId int `json:"publisherId"`
}
type struct AddTaskRes{
    CommonRes
    taskId int `json:"TaskId"`
}

/**搜索任务**/
type struct SearchTaskReq{
    DomainId int `json:"domainId"`
    Query string `json:"query"`
    minBonus int `json:"minBonus"` //最少奖金
    maxBonus int `json:"maxBonus"` //最大奖金
    minFine int `json:"minFine"` //最少罚金
    maxFine int `json:"maxFine"` //最大罚金
    MinDeadline string `json:"minDeadline"` //最小截止日期,yyyy-MM-DD
    MaxDeadline string `json:"maxDeadline"` //最大截止日期,yyyy-MM-DD
    ExportId int `json:"exportId"`
    PublisherId int `json:"publisherId"`
    SolverId int `json:"solverId"`
    Status int `json:"status"`
}
type struct SearchTaskRes{
    CommonRes
    TaskList []TaskInfo `json:"data"`
}

/**查询任务**/
type struct QueryTaskReq{
    TaskId int `json:"taskId"`
}

type struct QueryTaskRes{
    CommonRes
    Task TaskInfo `json:"data"`
}

/**更新任务**/
type struct UpdateTaskReq{
    TaskId int `json:"taskId"`
    TaskInfo
}

type struct UpdateTaskRes{
    CommonRes
}
