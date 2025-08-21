package main

import (
	"github.com/Mreza2020/Image_Processing_Service/Build"
	"github.com/Mreza2020/Image_Processing_Service/Handler"
	Login "github.com/Mreza2020/Image_Processing_Service/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(r *gin.Context) {
	var loginSt Handler.LoginST
	if err := r.ShouldBindJSON(&loginSt); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	} else {
		output := Login.Login(loginSt.Username, loginSt.Password)
		switch output {
		case "err":
			r.JSON(http.StatusBadRequest, gin.H{"error": "The information is wrong"})
		case "ok":
			token, err1 := Login.GenerateJWT(loginSt.Username)
			if err1 != nil {
				r.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}
			r.JSON(http.StatusOK, gin.H{"response": "The information is correct", "token": token})
		}
	}

}

func sign(r *gin.Context) {
	var signSt Handler.SignST
	if err := r.ShouldBindJSON(&signSt); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	} else {
		output := Login.Sign(signSt.Username, signSt.Password)
		if output != "ok" {
			r.JSON(http.StatusOK, gin.H{"response": "The information is correct"})
			return
		} else {
			r.JSON(http.StatusBadRequest, gin.H{"error": "The information is wrong"})
			return
		}
	}

}

func main() {
	Api := gin.Default()
	Api.POST("/IPS/login", login)
	Api.POST("/IPS/sign", sign)
	Api.POST("/IPS/Upload", Login.AuthMiddleware(), Build.Upload)
	Api.POST("/IPS/Compress", Login.AuthMiddleware(), Build.Compress)
	Api.POST("/IPS/Crop", Login.AuthMiddleware(), Build.Crop)
	Api.POST("/IPS/ApplyFilter", Login.AuthMiddleware(), Build.ApplyFilter)
	Api.POST("/IPS/Flip", Login.AuthMiddleware(), Build.Flip)
	Api.POST("/IPS/ChangeFormat", Login.AuthMiddleware(), Build.ChangeFormat)
	Api.POST("/IPS/Resize", Login.AuthMiddleware(), Build.Resize)
	Api.POST("/IPS/Rotate", Login.AuthMiddleware(), Build.Rotate)
	Api.POST("/IPS/Watermark", Login.AuthMiddleware(), Build.Watermark)

	err := Api.Run(":8080")
	if err != nil {
		return
	}

}
