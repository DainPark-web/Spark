package models

import "time"

// User Model
type User struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	ProfileImage    string    `json:"profile_image"`
	IsProfileLocked bool      `json:"is_profile_locked"` // Profile image lock status
	CreatedAt       time.Time `json:"created_at"`
}

// ChatRoom Model
type ChatRoom struct {
	ID        string    `json:"id"`
	User1ID   string    `json:"user1_id"`
	User2ID   string    `json:"user2_id"`
	Status    string    `json:"status"`     // "waiting", "active", "expired"
	ExpiresAt time.Time `json:"expires_at"` // Expires after 1 week
	CreatedAt time.Time `json:"created_at"`
}

// Message Model (Chat messages)
type Message struct {
	ID         string    `json:"id"`
	ChatRoomID string    `json:"chat_room_id"`
	SenderID   string    `json:"sender_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

// Like Model
type Like struct {
	ID         string    `json:"id"`
	FromUserID string    `json:"from_user_id"`
	ToUserID   string    `json:"to_user_id"`
	Status     string    `json:"status"`     // "pending", "accepted", "rejected"
	ExpiresAt  time.Time `json:"expires_at"` // Expires after 1 week
	CreatedAt  time.Time `json:"created_at"`
}

// PartnerMatch Model
type PartnerMatch struct {
	ID        string    `json:"id"`
	User1ID   string    `json:"user1_id"`
	User2ID   string    `json:"user2_id"`
	Status    string    `json:"status"`     // "pending", "accepted", "rejected"
	ExpiresAt time.Time `json:"expires_at"` // Match expiration time
	CreatedAt time.Time `json:"created_at"`
}

// FeedbackQuestion Model
type FeedbackQuestion struct {
	ID         string    `json:"id"`
	ChatRoomID string    `json:"chat_room_id"`
	UserID     string    `json:"user_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
}
