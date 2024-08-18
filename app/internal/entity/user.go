package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId         uuid.UUID `gorm:"type:uuid;primaryKey;column:userid;default:uuid_generate_v4()"`
	UserRegisterId uuid.UUID `gorm:"type:uuid;column:userregisterid;default:uuid_generate_v4()"`
	Email          *string   `gorm:"column:email;default:''"`
	Role           *string   `gorm:"column:role;default:''"`
	UserType       *string   `gorm:"column:usertype;default:''"`
	ProfilePicture *string   `gorm:"column:profilepicture;default:''"`
	IsActive       *bool     `gorm:"column:isactive;default:true"`
	IsDeleted      *bool     `gorm:"column:isdeleted;default:false"`
	CreatedOn      time.Time `gorm:"column:createdon;default:CURRENT_TIMESTAMP"`
	ModifiedOn     time.Time `gorm:"column:modifiedon;default:CURRENT_TIMESTAMP"`
}
