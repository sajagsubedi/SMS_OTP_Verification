package config

import(
  "os"
  "log"
  "github.com/joho/godotenv"
)
func GetAccountId()string {
  err:= godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error on loading .env file")
  }
  return os.Getenv("ACCOUNT_ID")
}
func GetServiceId()string {
  err:= godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error on loading .env file")
  }
  return os.Getenv("SERVICE_ID")

}
func GetAuthToken()string {
  err:= godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error on loading .env file")
  }
  return os.Getenv("AUTH_TOKEN")

}