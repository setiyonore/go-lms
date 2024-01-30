package entities

import (
	"time"

	"gorm.io/gorm"
)

type LibrarryMembers struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	PhoneNumber string         `json:"phone_number"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func FormatLibraryMember(libraryMember LibrarryMembers) LibrarryMembers {
	libraryMemberFormatter := LibrarryMembers{}
	libraryMemberFormatter.Id = libraryMember.Id
	libraryMemberFormatter.Name = libraryMember.Name
	libraryMemberFormatter.PhoneNumber = libraryMember.PhoneNumber
	return libraryMemberFormatter
}

func FormatLibraryMembers(libraryMembers []LibrarryMembers) []LibrarryMembers {
	libraryMemberFormtters := []LibrarryMembers{}
	for _, libraryMember := range libraryMembers {
		libraryMemberFormatter := FormatLibraryMember(libraryMember)
		libraryMemberFormtters = append(libraryMemberFormtters, libraryMemberFormatter)
	}
	return libraryMemberFormtters
}
