package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type LibraryMember interface {
	GetLibraryMember() ([]entities.LibrarryMembers, error)
	GetLibraryMemberById(id int) (entities.LibrarryMembers, error)
}

type librarymember struct {
	libraryMemberRepository repository.LibraryMember
}

func NewLibraryMember(libraryMemberRepository repository.LibraryMember) *librarymember {
	return &librarymember{libraryMemberRepository: libraryMemberRepository}
}

func (s *librarymember) GetLibraryMember() ([]entities.LibrarryMembers, error) {
	libraryMembers, err := s.libraryMemberRepository.FindAll()
	if err != nil {
		return libraryMembers, err
	}
	return libraryMembers, nil
}

func (s *librarymember) GetLibraryMemberById(id int) (entities.LibrarryMembers, error) {
	libraryMember, err := s.libraryMemberRepository.FindById(id)
	if err != nil {
		return libraryMember, err
	}
	if libraryMember.Id == 0 {
		return libraryMember, errors.New("data not found")
	}
	return libraryMember, nil
}
