package helpers

import (
  "errors"
  "github.com/sajagsubedi/SMS_OTP_Verification/config"
  "github.com/twilio/twilio-go"
  twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var ACCOUNT_ID = config.GetAccountId()
var SERVICE_ID = config.GetServiceId()
var AUTH_TOKEN = config.GetAuthToken()

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams {
  Username: ACCOUNT_ID,
  Password: AUTH_TOKEN,
})
func SendTwilioOTP(phoneNumber string)(string, error) {
  params:= twilioApi.CreateVerificationParams {}
  params.SetTo(phoneNumber)
  params.SetChannel("sms")
  resp,
  err:= client.VerifyV2.CreateVerification(SERVICE_ID, &params)
  if err != nil {
    return "",
    err
  }

  return *resp.Sid,
  nil
}
func VerifyTwilioOTP(phoneNumber string, code string) error {
  params:= &twilioApi.CreateVerificationCheckParams {}
  params.SetTo(phoneNumber)
  params.SetCode(code)

  resp,
  err:= client.VerifyV2.CreateVerificationCheck(SERVICE_ID, params)
  if err != nil {
    return err
  }
  if *resp.Status != "approved" {
    return errors.New("not a valid code")
  }

  return nil
}