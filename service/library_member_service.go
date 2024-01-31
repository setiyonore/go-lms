package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type LibraryMember interface {
	GetLibraryMember() ([]entities.LibrarryMembers, error)
	GetLibraryMemberById(id int) (entities.LibrarryMembers, error)
	GetLibraryMemberByName(name string) (entities.LibrarryMembers, error)
	AddLibrarryMember(input entities.AddLibraryMemberInput) error
	UpdateLibrarryMember(id int, input entities.AddLibraryMemberInput) error
	DeleteLibrarryMember(id int) error
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

func (s *librarymember) GetLibraryMemberByName(name string) (entities.LibrarryMembers, error) {
	libraryMember, err := s.libraryMemberRepository.FindByName(name)
	if err != nil {
		return libraryMember, err
	}
	if libraryMember.Id == 0 {
		return libraryMember, errors.New("data not found")
	}
	return libraryMember, nil

}

func (s *librarymember) AddLibrarryMember(input entities.AddLibraryMemberInput) error {
	librarryMember := entities.LibrarryMembers{}
	librarryMember.Name = input.Name
	librarryMember.PhoneNumber = input.PhoneNumber
	err := s.libraryMemberRepository.Save(librarryMember)
	if err != nil {
		return err
	}
	return nil
}

func (s *librarymember) UpdateLibrarryMember(id int, input entities.AddLibraryMemberInput) error {
	librarryMember := entities.LibrarryMembers{}

	librarryMember, err := s.libraryMemberRepository.FindById(id)
	if err != nil {
		return err
	}
	if librarryMember.Id == 0 {
		return errors.New("data not found")
	}
	librarryMember.Name = input.Name
	librarryMember.PhoneNumber = input.PhoneNumber
	err = s.libraryMemberRepository.Update(librarryMember)
	if err != nil {
		return err
	}
	return nil
}

func (s *librarymember) DeleteLibrarryMember(id int) error {
	librarryMember, err := s.libraryMemberRepository.FindById(id)
	if err != nil {
		return err
	}
	if librarryMember.Id == 0 {
		return errors.New("data not found")
	}
	err = s.libraryMemberRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
