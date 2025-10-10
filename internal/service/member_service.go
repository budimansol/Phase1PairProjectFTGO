package service

import (
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/repository"
)

type MemberService interface {
	CreateMember(member *model.Member) error
	UpdateMember(member *model.Member, id int) error
	DeleteMember(id int) error
	GetMemberByID(id int) (*model.Member, error)
	GetAllMembers() ([]model.Member, error)
}

type memberService struct {
	repo repository.MemberRepository
}

func NewMemberService(repo repository.MemberRepository) MemberService {
	return &memberService{repo: repo}
}

func (m *memberService) CreateMember(member *model.Member) error {
	return m.repo.CreateMember(member)
}

func (m *memberService) UpdateMember(member *model.Member, id int) error {
	return m.repo.UpdateMember(member, id)
}

func (m *memberService) DeleteMember(id int) error {
	return m.repo.DeleteMember(id)
}

func (m *memberService) GetMemberByID(id int) (*model.Member, error) {
	return m.repo.GetMemberByID(id)
}

func (m *memberService) GetAllMembers() ([]model.Member, error) {
	return m.repo.GetAllMembers()
}
