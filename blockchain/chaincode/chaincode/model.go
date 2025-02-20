package chaincode

/*
定义用户结构体
用户ID
用户类型
实名认证信息哈希,包括用户注册的姓名、身份证号、手机号、注册平台同意协议签名的哈希
新闻图片列表
*/
type User struct {
    UserID       string      `json:"userID"`
    UserType     string      `json:"userType"`
    RealInfoHash string      `json:"realInfoHash"`
    NewsImageList []*NewsImage `json:"newsImageList"`
}

/*
定义新闻图片结构体
溯源码
摄影记者输入
媒体机构输入
转发者输入
发布平台输入
*/
type NewsImage struct {
    Traceability_code string       `json:"traceability_code"`
    Reporter_input    Reporter_input  `json:"reporter_input"`
    Media_input       Media_input     `json:"media_input"`
    Spreader_input    Spreader_input  `json:"spreader_input"`
    Platform_input    Platform_input  `json:"platform_input"`
}

// HistoryQueryResult structure used for handling result of history query
type HistoryQueryResult struct {
    Record    *NewsImage `json:"record"`
    TxId      string     `json:"txId"`
    Timestamp string     `json:"timestamp"`
    IsDelete  bool       `json:"isDelete"`
}

/*
摄影记者
新闻图片的溯源码，一物一码（自动生成）
新闻事件主题
拍摄地点
拍摄时间
摄影记者姓名
*/
type Reporter_input struct {
    Re_newsTheme      string `json:"re_newsTheme"`
    Re_shootingLocation string `json:"re_shootingLocation"`
    Re_shootingTime   string `json:"re_shootingTime"`
    Re_photographerName string `json:"re_photographerName"`
    Re_Txid           string `json:"re_txid"`
    Re_Timestamp      string `json:"re_timestamp"`
}

/*
媒体机构
报道标题
发布批次
编辑处理时间
媒体机构名称与地址
联系电话
*/
type Media_input struct {
    Me_reportTitle      string `json:"me_reportTitle"`
    Me_releaseBatch     string `json:"me_releaseBatch"`
    Me_editingTime      string `json:"me_editingTime"`
    Me_mediaName        string `json:"me_mediaName"`
    Me_contactNumber    string `json:"me_contactNumber"`
    Me_Txid             string `json:"me_txid"`
    Me_Timestamp        string `json:"me_timestamp"`
}

/*
转发者
姓名
联系电话
转发记录
*/
type Spreader_input struct {
    Sp_name       string `json:"sp_name"`
    Sp_phone      string `json:"sp_phone"`
    Sp_spreadRecord string `json:"sp_spreadRecord"`
    Sp_Txid       string `json:"sp_txid"`
    Sp_Timestamp  string `json:"sp_timestamp"`
}

/*
发布平台
图片上传时间
热门传播时间
发布平台名称
发布平台位置
发布平台电话
*/
type Platform_input struct {
    Pl_uploadTime    string `json:"pl_uploadTime"`
    Pl_popularTime   string `json:"pl_popularTime"`
    Pl_platformName  string `json:"pl_platformName"`
    Pl_platformLocation string `json:"pl_platformLocation"`
    Pl_platformPhone string `json:"pl_platformPhone"`
    Pl_Txid          string `json:"pl_txid"`
    Pl_Timestamp     string `json:"pl_timestamp"`
}