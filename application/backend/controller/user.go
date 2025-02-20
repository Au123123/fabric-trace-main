package controller

import (
    "backend/model"
    "backend/pkg"

    "github.com/gin-gonic/gin"
)

// Register 处理用户注册逻辑
// 将用户信息存入 MySQL 数据库和区块链
func Register(c *gin.Context) {
    // 将用户信息存入 MySQL 数据库
    var user model.MysqlUser
    user.UserID = pkg.GenerateID()
    user.Username = c.PostForm("username")
    user.Password = c.PostForm("password")
    // 对用户名进行 MD5 加密作为实名信息哈希
    user.RealInfo = pkg.EncryptByMD5(c.PostForm("username"))
    err := pkg.InsertUser(&user)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "register failed：" + err.Error(),
        })
        return
    }
    // 将用户信息存入区块链
    // userID string, userType string, realInfoHash string
    // 将 post 请求的参数封装成一个数组 args
    var args []string
    args = append(args, user.UserID)
    args = append(args, c.PostForm("userType"))
    args = append(args, user.RealInfo)
    res, err := pkg.ChaincodeInvoke("RegisterUser", args)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "register failed：" + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "register success",
        "txid":    res,
    })
}

// Login 处理用户登录逻辑
// 验证用户信息，获取用户类型并生成 JWT
func Login(c *gin.Context) {
    var user model.MysqlUser
    user.Username = c.PostForm("username")
    user.Password = c.PostForm("password")
    // 获取用户 ID
    var err error
    user.UserID, err = pkg.GetUserID(user.Username)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "没有找到该用户",
        })
        return
    }
    userType, err := GetUserType(user.UserID)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "login failed:" + err.Error(),
        })
        return
    }
    err = pkg.Login(&user)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "login failed:" + err.Error(),
        })
        return
    }

    // 生成 JWT
    jwt, err := pkg.GenToken(user.UserID, userType)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "login failed:" + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "login success",
        "jwt":     jwt,
    })
}

// Logout 处理用户登出逻辑
func Logout(c *gin.Context) {
    c.JSON(200, gin.H{
        "code":    200,
        "message": "logout success",
    })
}

// GetUserType 获取用户类型
func GetUserType(userID string) (string, error) {
    userType, err := pkg.ChaincodeQuery("GetUserType", userID)
    if err != nil {
        return "", err
    }
    return userType, nil
}

// GetInfo 获取用户信息
func GetInfo(c *gin.Context) {
    userID, exist := c.Get("userID")
    if !exist {
        c.JSON(200, gin.H{
            "message": "get user type failed",
        })
        return
    }

    userType, err := GetUserType(userID.(string))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "get user type failed: " + err.Error(),
        })
        return
    }

    username, err := pkg.GetUsername(userID.(string))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "get user name failed: " + err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{
        "code":     200,
        "message":  "get user type success",
        "userType": userType,
        "username": username,
    })
}
