package routes
import (
  "github.com/sajagsubedi/SMS_OTP_Verification/controllers"
  "github.com/gin-gonic/gin"
)

func OTPRoutes(incomingRoutes *gin.Engine){
  incomingRoutes.POST("/getotp",controllers.SendOTP())
  incomingRoutes.POST("/verify",controllers.VerifyOTP())
}