/*
 * @Author: Au123123 2502894164@qq.com
 * @Date: 2024-12-25 17:35:30
 * @LastEditors: Au123123 2502894164@qq.com
 * @LastEditTime: 2025-02-20 10:33:58
 * @FilePath: \fabric-trace-main\blockchain\chaincode\chaincode\smartcontract.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package chaincode

 import (
	 "encoding/json"
	 "fmt"
	 "log"
	 "time"
 
	 "github.com/golang/protobuf/ptypes"
	 "github.com/hyperledger/fabric-contract-api-go/contractapi"
 )
 
 // 定义合约结构体
 type SmartContract struct {
	 contractapi.Contract
 }
 
 // 注册用户
 func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, realInfoHash string) error {
	 user := User{
		 UserID:       userID,
		 UserType:     userType,
		 RealInfoHash: realInfoHash,
		 NewsImageList: []*NewsImage{},
	 }
	 userAsBytes, err := json.Marshal(user)
	 if err != nil {
		 return err
	 }
	 err = ctx.GetStub().PutState(userID, userAsBytes)
	 if err != nil {
		 return err
	 }
	 return nil
 }
 
 // 新闻图片上链，传入用户ID、新闻图片上链信息
 func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, traceability_code string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) (string, error) {
	 // 获取用户类型
	 userType, err := s.GetUserType(ctx, userID)
	 if err != nil {
		 return "", fmt.Errorf("failed to get user type: %v", err)
	 }
 
	 // 通过溯源码获取新闻图片的上链信息
	 newsImageAsBytes, err := ctx.GetStub().GetState(traceability_code)
	 if err != nil {
		 return "", fmt.Errorf("failed to read from world state: %v", err)
	 }
	 // 将新闻图片的信息转换为NewsImage结构体
	 var newsImage NewsImage
	 if newsImageAsBytes != nil {
		 err = json.Unmarshal(newsImageAsBytes, &newsImage)
		 if err != nil {
			 return "", fmt.Errorf("failed to unmarshal news image: %v", err)
		 }
	 }
 
	 //获取时间戳,修正时区
	 txtime, err := ctx.GetStub().GetTxTimestamp()
	 if err != nil {
		 return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	 }
	 timeLocation, _ := time.LoadLocation("Asia/Shanghai") // 选择你所在的时区
	 time := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")
 
	 // 获取交易ID
	 txid := ctx.GetStub().GetTxID()
	 // 给新闻图片信息中加上溯源码
	 newsImage.Traceability_code = traceability_code
	 // 不同用户类型的上链的参数不一致
	 switch userType {
	 // 摄影记者
	 case "摄影记者":
		 // 将传入的新闻图片上链信息转换为Reporter_input结构体
		 newsImage.Reporter_input.Re_newsTheme = arg1
		 newsImage.Reporter_input.Re_shootingLocation = arg2
		 newsImage.Reporter_input.Re_shootingTime = arg3
		 newsImage.Reporter_input.Re_photographerName = arg4
		 newsImage.Reporter_input.Re_Txid = txid
		 newsImage.Reporter_input.Re_Timestamp = time
	 // 媒体机构
	 case "媒体机构":
		 // 将传入的新闻图片上链信息转换为Media_input结构体
		 newsImage.Media_input.Me_reportTitle = arg1
		 newsImage.Media_input.Me_releaseBatch = arg2
		 newsImage.Media_input.Me_editingTime = arg3
		 newsImage.Media_input.Me_mediaName = arg4
		 newsImage.Media_input.Me_contactNumber = arg5
		 newsImage.Media_input.Me_Txid = txid
		 newsImage.Media_input.Me_Timestamp = time
	 // 转发者
	 case "转发者":
		 // 将传入的新闻图片上链信息转换为Spreader_input结构体
		 newsImage.Spreader_input.Sp_name = arg1
		 newsImage.Spreader_input.Sp_phone = arg2
		 newsImage.Spreader_input.Sp_spreadRecord = arg3
		 newsImage.Spreader_input.Sp_Txid = txid
		 newsImage.Spreader_input.Sp_Timestamp = time
	 // 发布平台
	 case "发布平台":
		 // 将传入的新闻图片上链信息转换为Platform_input结构体
		 newsImage.Platform_input.Pl_uploadTime = arg1
		 newsImage.Platform_input.Pl_popularTime = arg2
		 newsImage.Platform_input.Pl_platformName = arg3
		 newsImage.Platform_input.Pl_platformLocation = arg4
		 newsImage.Platform_input.Pl_platformPhone = arg5
		 newsImage.Platform_input.Pl_Txid = txid
		 newsImage.Platform_input.Pl_Timestamp = time
	 }
 
	 //将新闻图片的信息转换为json格式
	 newsImageAsBytes, err = json.Marshal(newsImage)
	 if err != nil {
		 return "", fmt.Errorf("failed to marshal news image: %v", err)
	 }
	 //将新闻图片的信息存入区块链
	 err = ctx.GetStub().PutState(traceability_code, newsImageAsBytes)
	 if err != nil {
		 return "", fmt.Errorf("failed to put news image: %v", err)
	 }
	 //将新闻图片存入用户的新闻图片列表
	 err = s.AddNewsImage(ctx, userID, &newsImage)
	 if err != nil {
		 return "", fmt.Errorf("failed to add news image to user: %v", err)
	 }
 
	 return txid, nil
 }
 
 // 添加新闻图片到用户的新闻图片列表
 func (s *SmartContract) AddNewsImage(ctx contractapi.TransactionContextInterface, userID string, newsImage *NewsImage) error {
	 userBytes, err := ctx.GetStub().GetState(userID)
	 if err != nil {
		 return fmt.Errorf("failed to read from world state: %v", err)
	 }
	 if userBytes == nil {
		 return fmt.Errorf("the user %s does not exist", userID)
	 }
	 // 将返回的结果转换为User结构体
	 var user User
	 err = json.Unmarshal(userBytes, &user)
	 if err != nil {
		 return err
	 }
	 user.NewsImageList = append(user.NewsImageList, newsImage)
	 userAsBytes, err := json.Marshal(user)
	 if err != nil {
		 return err
	 }
	 err = ctx.GetStub().PutState(userID, userAsBytes)
	 if err != nil {
		 return err
	 }
	 return nil
 }
 
 // 获取用户类型
 func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	 userBytes, err := ctx.GetStub().GetState(userID)
	 if err != nil {
		 return "", fmt.Errorf("failed to read from world state: %v", err)
	 }
	 if userBytes == nil {
		 return "", fmt.Errorf("the user %s does not exist", userID)
	 }
	 // 将返回的结果转换为User结构体
	 var user User
	 err = json.Unmarshal(userBytes, &user)
	 if err != nil {
		 return "", err
	 }
	 return user.UserType, nil
 }
 
 // 获取用户信息
 func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	 userBytes, err := ctx.GetStub().GetState(userID)
	 if err != nil {
		 return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	 }
	 if userBytes == nil {
		 return &User{}, fmt.Errorf("the user %s does not exist", userID)
	 }
	 // 将返回的结果转换为User结构体
	 var user User
	 err = json.Unmarshal(userBytes, &user)
	 if err != nil {
		 return &User{}, err
	 }
	 return &user, nil
 }
 
 // 获取新闻图片的上链信息
 func (s *SmartContract) GetNewsImageInfo(ctx contractapi.TransactionContextInterface, traceability_code string) (*NewsImage, error) {
	 newsImageAsBytes, err := ctx.GetStub().GetState(traceability_code)
	 if err != nil {
		 return &NewsImage{}, fmt.Errorf("failed to read from world state: %v", err)
	 }
 
	 // 将返回的结果转换为NewsImage结构体
	 var newsImage NewsImage
	 if newsImageAsBytes != nil {
		 err = json.Unmarshal(newsImageAsBytes, &newsImage)
		 if err != nil {
			 return &NewsImage{}, fmt.Errorf("failed to unmarshal news image: %v", err)
		 }
		 if newsImage.Traceability_code != "" {
			 return &newsImage, nil
		 }
	 }
	 return &NewsImage{}, fmt.Errorf("the news image %s does not exist", traceability_code)
 }
 
 // 获取用户的新闻图片ID列表
 func (s *SmartContract) GetNewsImageList(ctx contractapi.TransactionContextInterface, userID string) ([]*NewsImage, error) {
	 userBytes, err := ctx.GetStub().GetState(userID)
	 if err != nil {
		 return nil, fmt.Errorf("failed to read from world state: %v", err)
	 }
	 if userBytes == nil {
		 return nil, fmt.Errorf("the user %s does not exist", userID)
	 }
	 // 将返回的结果转换为User结构体
	 var user User
	 err = json.Unmarshal(userBytes, &user)
	 if err != nil {
		 return nil, err
	 }
	 return user.NewsImageList, nil
 }
 
 // 获取所有的新闻图片信息
 func (s *SmartContract) GetAllNewsImageInfo(ctx contractapi.TransactionContextInterface) ([]NewsImage, error) {
	 newsImageListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	 if err != nil {
		 return nil, fmt.Errorf("failed to read from world state: %v", err)
	 }
	 defer newsImageListAsBytes.Close()
	 var newsImages []NewsImage
	 for newsImageListAsBytes.HasNext() {
		 queryResponse, err := newsImageListAsBytes.Next()
		 if err != nil {
			 return nil, err
		 }
		 var newsImage NewsImage
		 err = json.Unmarshal(queryResponse.Value, &newsImage)
		 if err != nil {
			 return nil, err
		 }
		 // 过滤非新闻图片的信息
		 if newsImage.Traceability_code != "" {
			 newsImages = append(newsImages, newsImage)
		 }
	 }
	 return newsImages, nil
 }
 
 // 获取新闻图片上链历史
 func (s *SmartContract) GetNewsImageHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
	 log.Printf("GetAssetHistory: ID %v", traceability_code)
 
	 resultsIterator, err := ctx.GetStub().GetHistoryForKey(traceability_code)
	 if err != nil {
		 return nil, err
	 }
	 defer resultsIterator.Close()
 
	 var records []HistoryQueryResult
	 for resultsIterator.HasNext() {
		 response, err := resultsIterator.Next()
		 if err != nil {
			 return nil, err
		 }
 
		 var newsImage NewsImage
		 if len(response.Value) > 0 {
			 err = json.Unmarshal(response.Value, &newsImage)
			 if err != nil {
				 return nil, err
			 }
		 } else {
			 newsImage = NewsImage{
				 Traceability_code: traceability_code,
			 }
		 }
 
		 timestamp, err := ptypes.Timestamp(response.Timestamp)
		 if err != nil {
			 return nil, err
		 }
		 // 指定目标时区
		 targetLocation, err := time.LoadLocation("Asia/Shanghai")
		 if err != nil {
			 return nil, err
		 }
 
		 // 将时间戳转换到目标时区
		 timestamp = timestamp.In(targetLocation)
		 // 格式化时间戳为指定格式的字符串
		 formattedTime := timestamp.Format("2006-01-02 15:04:05")
 
		 record := HistoryQueryResult{
			 TxId:      response.TxId,
			 Timestamp: formattedTime,
			 Record:    &newsImage,
			 IsDelete:  response.IsDelete,
		 }
		 records = append(records, record)
	 }
 
	 return records, nil
 }