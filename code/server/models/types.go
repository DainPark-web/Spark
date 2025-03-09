package models

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Man   Gender = "man"
	Woman Gender = "woman"
)

// User Model
type User struct {
	gorm.Model             // ID, CreatedAt, UpdatedAt, DeletedAt 포함
	Email           string `json:"email" gorm:"not null"`
	Name            string `json:"name" gorm:"not null"`
	Password        string `json:"-" gorm:"not null"`
	ProfileImage    string `json:"profile_image" gorm:"null"`
	IsProfileLocked bool   `json:"is_profile_locked" gorm:"default:false"`
	Description     string `json:"description" gorm:"null"`
	Height          int    `json:"height" gorm:"null"`
	Weight          int    `json:"weight" gorm:"null"`
	Job             string `json:"job" gorm:"null"`
	Location        string `json:"location" gorm:"null"`
	Gender          Gender `json:"gender" gorm:"null"`
	// ChatRooms       []ChatRoom `json:"chat_rooms" gorm:"many2many:user_chat_rooms"`
}

type TinyUser struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
	Description  string `json:"description"`
}

// // PartnerMatch Model
type PartnerMatch struct {
	gorm.Model
	User1ID       uint      `json:"user1_id" gorm:"uniqueIndex:idx_user1_user2"`
	User2ID       uint      `json:"user2_id" gorm:"uniqueIndex:idx_user1_user2"`
	Status        string    `json:"status" gorm:"default:'pending'"` // "pending", "accepted", "rejected"
	ExpiresAt     time.Time `json:"expires_at"`                      // Match expiration time
	User1         User      `json:"-" gorm:"foreignKey:User1ID"`
	User2         User      `json:"-" gorm:"foreignKey:User2ID"`
	User1Accepted bool      `json:"user1_accepted" gorm:"default:false"`
	User2Accepted bool      `json:"user2_accepted" gorm:"default:false"`
}

// // ChatRoom Model
// type ChatRoom struct {
// 	gorm.Model
// 	User1ID   uint      `json:"user1_id" gorm:"not null"`
// 	User2ID   uint      `json:"user2_id" gorm:"not null"`
// 	Status    string    `json:"status" gorm:"default:'waiting'"` // "waiting", "active", "expired"
// 	ExpiresAt time.Time `json:"expires_at"`                      // Expires after 1 week
// 	Messages  []Message `json:"messages" gorm:"foreignKey:ChatRoomID"`
// 	Users     []User    `json:"users" gorm:"many2many:user_chat_rooms"`
// }

// // Message Model (Chat messages)
// type Message struct {
// 	gorm.Model
// 	ChatRoomID uint     `json:"chat_room_id" gorm:"not null"`
// 	SenderID   uint     `json:"sender_id" gorm:"not null"`
// 	Content    string   `json:"content" gorm:"not null"`
// 	ChatRoom   ChatRoom `json:"-" gorm:"foreignKey:ChatRoomID"`
// 	Sender     User     `json:"-" gorm:"foreignKey:SenderID"`
// }

// // Like Model
// type Like struct {
// 	gorm.Model
// 	FromUserID uint      `json:"from_user_id" gorm:"not null"`
// 	ToUserID   uint      `json:"to_user_id" gorm:"not null"`
// 	Status     string    `json:"status" gorm:"default:'pending'"` // "pending", "accepted", "rejected"
// 	ExpiresAt  time.Time `json:"expires_at"`                      // Expires after 1 week
// 	FromUser   User      `json:"-" gorm:"foreignKey:FromUserID"`
// 	ToUser     User      `json:"-" gorm:"foreignKey:ToUserID"`
// }

// // FeedbackQuestion Model
// type FeedbackQuestion struct {
// 	gorm.Model
// 	ChatRoomID uint     `json:"chat_room_id" gorm:"not null"`
// 	UserID     uint     `json:"user_id" gorm:"not null"`
// 	Answer     string   `json:"answer"`
// 	ChatRoom   ChatRoom `json:"-" gorm:"foreignKey:ChatRoomID"`
// 	User       User     `json:"-" gorm:"foreignKey:UserID"`
// }
