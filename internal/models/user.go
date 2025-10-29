package models

import (
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
type User struct {
	ID                  uint           `gorm:"primaryKey"`
	UUID                string         `gorm:"unique;not null"`
	Username            string         `gorm:"unique;not null"`
	Email               string         `gorm:"unique;not null"`
	FirstName           string
	LastName            string
	PasswordHash        string         `gorm:"not null"`
	Language            string         `gorm:"default:'en'"`
	RootAdmin           bool           `gorm:"default:false"`
	UseTOTP             bool           `gorm:"default:false"`
	TOTPSecret          string
	TOTPAuthenticatedAt *time.Time
	Gravatar            bool           `gorm:"default:true"`
	CreatedAt           time.Time
	UpdatedAt           time.Time

	// Relations
	// TODO: Create those relations, they can only be created
	// when we create the proper files, for example the Server 
	// model.
	Servers   []Server   `gorm:"foreignKey:OwnerID"`
	SSHKeys   []UserSSHKey `gorm:"foreignKey:UserID"`
	APIKeys   []APIKey     `gorm:"foreignKey:UserID"`
}

// TODO: Create the HelperFunctions, which do the following:
// - BeforeCreate hook to generate UUID
// - SetPassword hashes and sets the user's password
// - CheckPassword validates a password against the stored hash
// - FullName returns the user's full name
// - IsAdmin returns true if the user is an admin
// - AccessibleServers can preload all servers the user owns

// NOTE: This is taken as reference from Pterodactyl.