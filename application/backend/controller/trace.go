package controller

import (
    "backend/pkg"
    "fmt"

    "github.com/gin-gonic/gin"
)

// 保存了新闻图片上链与查询的函数

func Uplink(c *gin.Context) {
    // 与userID不一样，取ID从第二位作为溯源码，即18位，雪花ID的结构如下（转化为10进制共19位）：
    // +--------------------------------------------------------------------------+
    // | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
    // +--------------------------------------------------------------------------+
    newsImage_traceability_code := pkg.GenerateID()[1:]
    args := buildArgs(c, newsImage_traceability_code)
    if args == nil {
        return
    }
    res, err := pkg.ChaincodeInvoke("Uplink", args)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "uplink failed: " + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":              200,
        "message":           "uplink success",
        "txid":              res,
        "traceability_code": args[1],
    })
}

// 获取新闻图片的上链信息
func GetNewsImageInfo(c *gin.Context) {
    res, err := pkg.ChaincodeQuery("GetNewsImageInfo", c.PostForm("traceability_code"))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "查询失败：" + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "query success",
        "data":    res,
    })
}

// 获取用户的新闻图片ID列表
func GetNewsImageList(c *gin.Context) {
    userID, _ := c.Get("userID")
    res, err := pkg.ChaincodeQuery("GetNewsImageList", userID.(string))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "query failed: " + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "query success",
        "data":    res,
    })
}

// 获取所有的新闻图片信息
func GetAllNewsImageInfo(c *gin.Context) {
    res, err := pkg.ChaincodeQuery("GetAllNewsImageInfo", "")
    fmt.Print("res", res)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "query failed: " + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "query success",
        "data":    res,
    })
}

// 获取新闻图片上链历史
func GetNewsImageHistory(c *gin.Context) {
    res, err := pkg.ChaincodeQuery("GetNewsImageHistory", c.PostForm("traceability_code"))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "query failed: " + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "query success",
        "data":    res,
    })
}

func buildArgs(c *gin.Context, newsImage_traceability_code string) []string {
    var args []string
    userID, _ := c.Get("userID")
    userType, _ := pkg.ChaincodeQuery("GetUserType", userID.(string))
    args = append(args, userID.(string))
    fmt.Print(userID)
    // 摄影记者不需要输入溯源码，其他用户需要，通过雪花算法生成ID
    if userType == "摄影记者" {
        args = append(args, newsImage_traceability_code)
    } else {
        // 检查溯源码是否正确
        res, err := pkg.ChaincodeQuery("GetNewsImageInfo", c.PostForm("traceability_code"))
        if res == "" || err != nil || len(c.PostForm("traceability_code")) != 18 {
            c.JSON(200, gin.H{
                "message": "请检查溯源码是否正确!!",
            })
            return nil
        } else {
            args = append(args, c.PostForm("traceability_code"))
        }
    }
    args = append(args, c.PostForm("arg1"))
    args = append(args, c.PostForm("arg2"))
    args = append(args, c.PostForm("arg3"))
    args = append(args, c.PostForm("arg4"))
    args = append(args, c.PostForm("arg5"))
    return args
}
