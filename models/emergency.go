package models

import "time"

type Emergency struct {
	CreatedAt   time.Time
	PendingList []UserStatus
	SafeList    []UserStatus
	UnsafeList  []UserStatus
}
