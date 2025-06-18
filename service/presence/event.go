package presence

import (
	"encoding/json"
	"time"
)

const (
	EventUserBanishment    = "event.banishment"
	EventUserStateChanged  = "event.userstatechanged"
	EventUserReasonChanged = "event.userreasonchanged"
	EventUserAssignChanged = "event.userassignchanged"
	EventPhoneBusy         = "event.busy"
	EventPhoneIdle         = "event.idle"
)

type Banishment struct {
	Event       string    `json:"event,omitempty"`
	OldUserID   string    `json:"oldUserId,omitempty"`
	OldAppID    string    `json:"oldAppId,omitempty"`
	OldDn       string    `json:"oldDn,omitempty"`
	OldMediaSet []string  `json:"oldMediaset,omitempty"`
	NewUserID   string    `json:"newUserId,omitempty"`
	NewAppID    string    `json:"newAppId,omitempty"`
	NewDn       string    `json:"newDn,omitempty"`
	NewMediaSet []string  `json:"newMediaset,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	NowTime     time.Time `json:"nowTime,omitempty"`
}

type UserStateChanged struct {
	Event             string      `json:"event,omitempty"`
	UserID            string      `json:"userId,omitempty"`
	OldState          string      `json:"oldState,omitempty"`
	OldReason         string      `json:"oldReason,omitempty"`
	OldCallID         string      `json:"oldCallId,omitempty"`
	NewState          string      `json:"newState,omitempty"`
	NewReason         string      `json:"newReason,omitempty"`
	CallID            string      `json:"callId,omitempty"`
	AcdID             string      `json:"acdId,omitempty"`
	ANI               string      `json:"ani,omitempty"`
	DNIS              string      `json:"dnis,omitempty"`
	OldStateTime      time.Time   `json:"oldStateTime,omitempty"`
	NewStateTime      time.Time   `json:"newStateTime,omitempty"`
	DayFirstLogin     bool        `json:"dayFirstLogin,omitempty"`
	DayFirstLoginTime time.Time   `json:"dayFirstLoginTime,omitempty"`
	DayLastLoginTime  time.Time   `json:"dayLastLoginTime,omitempty"`
	DayWorkDuration   json.Number `json:"dayWorkDuration,omitempty"`
	NowTime           time.Time   `json:"nowTime,omitempty"`
	MediaType         string      `json:"mediaType,omitempty"`
}

type UserReasonChanged struct {
	Event        string    `json:"event,omitempty"`
	UserID       string    `json:"userId,omitempty"`
	State        string    `json:"state,omitempty"`
	OldReason    string    `json:"oldReason,omitempty"`
	NewReason    string    `json:"newReason,omitempty"`
	OldStateTime time.Time `json:"oldStateTime,omitempty"`
	NewStateTime time.Time `json:"newStateTime,omitempty"`
	NowTime      time.Time `json:"nowTime,omitempty"`
	MediaType    string    `json:"mediaType,omitempty"`
}

type UserAssignChanged struct {
	Event        string    `json:"event,omitempty"`
	UserID       string    `json:"userId,omitempty"`
	OldState     string    `json:"oldState,omitempty"`
	NewState     string    `json:"newState,omitempty"`
	CallID       string    `json:"callId,omitempty"`
	AcdID        string    `json:"acdId,omitempty"`
	MediaType    string    `json:"mediaType,omitempty"`
	ANI          string    `json:"ani,omitempty"`
	DNIS         string    `json:"dnis,omitempty"`
	AssignedTime time.Time `json:"assignedTime,omitempty"`
	NowTime      time.Time `json:"nowTime,omitempty"`
}

type UserDeleted struct {
	Event   string    `json:"event,omitempty"`
	UserID  string    `json:"userId,omitempty"`
	NowTime time.Time `json:"nowTime,omitempty"`
}

// 임시 정의.
type PhoneRegisted struct {
	Event     string      `json:"event,omitempty"`
	PhoneID   string      `json:"phoneId,omitempty"`
	UserID    string      `json:"userId,omitempty"`
	Expire    json.Number `json:"expire,omitempty"`
	Extension string      `json:"extension,omitempty"`
}

type PhoneUnregisted struct {
	Event       string    `json:"event,omitempty"`
	PhoneID     string    `json:"phoneId,omitempty"`
	UserID      string    `json:"userId,omitempty"`
	LastRegDate time.Time `json:"lastRegDate,omitempty"`
	Extension   string    `json:"extension,omitempty"`
}

type PhoneIdle struct {
	Event   string `json:"event,omitempty"`
	PhoneID string `json:"phoneId,omitempty"`
}

type PhoneBusy struct {
	Event   string `json:"event,omitempty"`
	PhoneID string `json:"phoneId,omitempty"`
}

type PhoneAlert struct {
	Event   string `json:"event,omitempty"`
	PhoneID string `json:"phoneId,omitempty"`
}

type PhoneDeleted struct {
	Event   string `json:"event,omitempty"`
	PhoneID string `json:"phoneId,omitempty"`
}

type MediaChanged struct {
	Event     string    `json:"event,omitempty"`
	UserID    string    `json:"userId,omitempty"`
	Enable    bool      `json:"enable,omitempty"`
	MediaType string    `json:"mediaType,omitempty"`
	NowTime   time.Time `json:"nowTime,omitempty"`
}
