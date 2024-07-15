package handler

import (
	"context"
	"errors"
	"sync"

	"Go_course/accounts/http/models"
	"Go_course/proto"
)

type AccountsHandler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
	proto.AccountsServer
}

func New() *AccountsHandler {
	return &AccountsHandler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

func (s *AccountsHandler) CreateAccount(ctx context.Context, in *proto.CreateAccountRequest) (*proto.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if len(in.Name) == 0 {
		return nil, errors.New("empty name")
	}

	if _, ok := s.accounts[in.Name]; ok {
		return nil, errors.New("account already exists")
	}

	s.accounts[in.Name] = &models.Account{
		Name:   in.Name,
		Amount: int(in.Amount),
	}

	return &proto.Empty{}, nil
}

func (s *AccountsHandler) GetAccount(ctx context.Context, in *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	s.guard.RLock()
	defer s.guard.RUnlock()

	if len(in.Name) == 0 {
		return nil, errors.New("empty name")
	}

	account, ok := s.accounts[in.Name]

	if !ok {
		return nil, errors.New("account does not exist")
	}

	return &proto.GetAccountResponse{
		Name:   account.Name,
		Amount: int32(account.Amount),
	}, nil
}

func (s *AccountsHandler) PatchAccount(ctx context.Context, in *proto.PatchAccountRequest) (*proto.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if len(in.Name) == 0 {
		return nil, errors.New("empty name")
	}

	account, ok := s.accounts[in.Name]

	if !ok {
		return nil, errors.New("account does not exist")
	}

	account.Amount = int(in.Amount)

	return &proto.Empty{}, nil
}

func (s *AccountsHandler) ChangeAccount(ctx context.Context, in *proto.ChangeAccountRequest) (*proto.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if len(in.Name) == 0 {
		return nil, errors.New("empty name")
	}

	if len(in.Newname) == 0 {
		return nil, errors.New("empty new name")
	}

	account, ok := s.accounts[in.Name]

	if !ok {
		return nil, errors.New("account does not exist")
	}

	if _, ok := s.accounts[in.Newname]; ok {
		return nil, errors.New("new account already exists")
	}

	account.Name = in.Newname
	s.accounts[in.Newname] = account
	delete(s.accounts, in.Name)

	return &proto.Empty{}, nil
}

func (s *AccountsHandler) DeleteAccount(ctx context.Context, in *proto.DeleteAccountRequest) (*proto.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()

	if len(in.Name) == 0 {
		return nil, errors.New("empty name")
	}

	if _, ok := s.accounts[in.Name]; !ok {
		return nil, errors.New("account does not exist")
	}

	delete(s.accounts, in.Name)

	return &proto.Empty{}, nil
}
