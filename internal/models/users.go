package models

// Users - Model for the uses table
type Users struct {
    Id       int    `json:"id" orm:"auto"`
    UserName string `json:"user_name" orm:"size(32)"`
    Password string `json:"password" orm:"size(64)"`
}