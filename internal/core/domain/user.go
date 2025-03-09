package domain

type (
	UserID   = uint64
	UserVkID = uint64
)

type UserModel struct {
	ID   UserID   `json:"id" db:"id"`
	VkID UserVkID `json:"vk_id" db:"vk_id"`
}
