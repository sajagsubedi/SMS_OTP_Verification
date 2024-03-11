package controllers

import(
  "time"
  "net/http"
  "context"
 "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "github.com/sajagsubedi/SMS_OTP_Verification/models"
  "github.com/sajagsubedi/SMS_OTP_Verification/helpers"
 )
 
var validate=validator.New()

func SendOTP()gin.HandlerFunc{
  return func(c *gin.Context){
    _,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var payload models.OTPdata
    if err:=c.BindJSON(&payload);err!=nil{
      c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
      return
    }
    validationErr:=validate.Struct(payload)
    if validationErr!=nil{
      c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":validationErr.Error()})
      return
    }
    _,err:=helpers.SendTwilioOTP(payload.PhoneNumber)
    if err!=nil{
      c.JSON(http.StatusInternalServerError,gin.H{"success":false,"message":"Internal Server Error!","err":err.Error()})
      return
    }
    c.JSON(http.StatusAccepted,gin.H{"success":true,"message":"OTP sent successfully!"})
  }
}

func VerifyOTP()gin.HandlerFunc{
  return func(c *gin.Context){
    _,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var payload models.VerifyData
    if err:=c.BindJSON(&payload);err!=nil{
      c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":err.Error()})
      return
    }
    validationErr:=validate.Struct(payload)
    if validationErr!=nil{
      c.JSON(http.StatusBadRequest,gin.H{"success":false,"message":validationErr.Error()})
      return
    }
    err:=helpers.VerifyTwilioOTP(payload.User,payload.Code)
    if err!=nil{
      c.JSON(http.StatusInternalServerError,gin.H{"success":false,"message":"Internal Server Error!"})
      return
    }
    c.JSON(http.StatusAccepted,gin.H{"success":true,"message":"OTP verified successfully!"})
  }
}