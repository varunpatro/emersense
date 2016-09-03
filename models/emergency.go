package models

import "time"

type Emergency struct {
	Id          int          `json:"id"`
	CreatedAt   time.Time    `json:"createdAt"`
	PendingList []UserStatus `json:"pendingList"`
	SafeList    []UserStatus `json:"safeList"`
	UnsafeList  []UserStatus `json:"unsafeList"`
	SheetTitle  string       `json:"sheet_title"`
	SheetURL    string       `json:"sheet_url"`
}
