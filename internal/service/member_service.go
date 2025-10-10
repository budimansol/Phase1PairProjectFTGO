package service

import (
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/repository"
)

type MemberService struct {
	repo *repository.MemberRepository
}

func NewMemberService(repo *repository.MemberRepository) *MemberService {
	return &MemberService{repo: repo}
}

func (m *MemberService) CreateMember(member *model.Member) error {
	return m.repo.CreateMember(member)
}

func (m *MemberService) UpdateMember(member *model.Member, id int) error {
	return m.repo.UpdateMember(member, id)
}

func (m *MemberService) DeleteMember(id int) error {
	return m.repo.DeleteMember(id)
}

func (m *MemberService) GetMemberByID(id int) (*model.Member, error) {
	return m.repo.GetMemberByID(id)
}

func (m *MemberService) GetAllMembers() ([]model.Member, error) {
	return m.repo.GetAllMembers()
}
