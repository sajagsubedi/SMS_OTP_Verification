package main

import(
  "os"
  "github.com/gin-gonic/gin"
  "github.com/sajagsubedi/SMS_OTP_Verification/routes"
  "github.com/joho/godotenv"
  "github.com/gin-contrib/cors"
  "log"
)

func main() {
  err:= godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error on loading .env file")
  }
  port:= os.Getenv("PORT")
  if port == "" {
    port = "8000"
  }
  router:= gin.New()

  config:= cors.DefaultConfig()
  config.AllowOrigins = []string {
    "*",
  }
  config.AllowMethods = []string {
    "POST",
  }
  config.AllowHeaders = []string {
    "Origin",
    "Content-Type",
  }

  //middlewares
  router.Use(cors.New(config))
  router.Use(gin.Logger())

  //routes
  routes.OTPRoutes(router)

  router.Run(":"+port)
}