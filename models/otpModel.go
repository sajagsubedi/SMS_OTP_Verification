package models

type OTPdata struct{
  PhoneNumber string `json:"phonenumber,omitempty" validate:"required"`
}

type VerifyData struct{
  User    string `json:"user,omitempty" validate:"required"`
  Code    string `json:"user,omitempty" validate:"required"`
}