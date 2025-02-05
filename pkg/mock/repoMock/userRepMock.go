// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/userRepo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/msegeya56/eCommerce-smart_gads/pkg/domain"
	request "github.com/msegeya56/eCommerce-smart_gads/pkg/utils/request"
	response "github.com/msegeya56/eCommerce-smart_gads/pkg/utils/response"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddToWishlist mocks base method.
func (m *MockUserRepository) AddToWishlist(ctx context.Context, wishlistData request.AddToWishlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToWishlist", ctx, wishlistData)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToWishlist indicates an expected call of AddToWishlist.
func (mr *MockUserRepositoryMockRecorder) AddToWishlist(ctx, wishlistData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToWishlist", reflect.TypeOf((*MockUserRepository)(nil).AddToWishlist), ctx, wishlistData)
}

// CreditUserWallet mocks base method.
func (m *MockUserRepository) CreditUserWallet(ctx context.Context, data domain.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditUserWallet", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreditUserWallet indicates an expected call of CreditUserWallet.
func (mr *MockUserRepositoryMockRecorder) CreditUserWallet(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditUserWallet", reflect.TypeOf((*MockUserRepository)(nil).CreditUserWallet), ctx, data)
}

// DeleteAddress mocks base method.
func (m *MockUserRepository) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", ctx, userID, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserRepositoryMockRecorder) DeleteAddress(ctx, userID, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserRepository)(nil).DeleteAddress), ctx, userID, addressID)
}

// DeleteFromWishlist mocks base method.
func (m *MockUserRepository) DeleteFromWishlist(ctx context.Context, productId, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFromWishlist", ctx, productId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFromWishlist indicates an expected call of DeleteFromWishlist.
func (mr *MockUserRepositoryMockRecorder) DeleteFromWishlist(ctx, productId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromWishlist", reflect.TypeOf((*MockUserRepository)(nil).DeleteFromWishlist), ctx, productId, userId)
}

// FindUser mocks base method.
func (m *MockUserRepository) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, user)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockUserRepositoryMockRecorder) FindUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUserRepository)(nil).FindUser), ctx, user)
}

// GetAllAddress mocks base method.
func (m *MockUserRepository) GetAllAddress(ctx context.Context, userId uint) ([]response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAddress", ctx, userId)
	ret0, _ := ret[0].([]response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAddress indicates an expected call of GetAllAddress.
func (mr *MockUserRepositoryMockRecorder) GetAllAddress(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAddress", reflect.TypeOf((*MockUserRepository)(nil).GetAllAddress), ctx, userId)
}

// GetCartIdByUserId mocks base method.
func (m *MockUserRepository) GetCartIdByUserId(ctx context.Context, userId uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartIdByUserId", ctx, userId)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartIdByUserId indicates an expected call of GetCartIdByUserId.
func (mr *MockUserRepositoryMockRecorder) GetCartIdByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartIdByUserId", reflect.TypeOf((*MockUserRepository)(nil).GetCartIdByUserId), ctx, userId)
}

// GetCartItemsbyUserId mocks base method.
func (m *MockUserRepository) GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) ([]response.CartItemResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartItemsbyUserId", ctx, page, userID)
	ret0, _ := ret[0].([]response.CartItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartItemsbyUserId indicates an expected call of GetCartItemsbyUserId.
func (mr *MockUserRepositoryMockRecorder) GetCartItemsbyUserId(ctx, page, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartItemsbyUserId", reflect.TypeOf((*MockUserRepository)(nil).GetCartItemsbyUserId), ctx, page, userID)
}

// GetDefaultAddress mocks base method.
func (m *MockUserRepository) GetDefaultAddress(ctx context.Context, userId uint) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultAddress", ctx, userId)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultAddress indicates an expected call of GetDefaultAddress.
func (mr *MockUserRepositoryMockRecorder) GetDefaultAddress(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultAddress", reflect.TypeOf((*MockUserRepository)(nil).GetDefaultAddress), ctx, userId)
}

// GetEmailPhoneByUserId mocks base method.
func (m *MockUserRepository) GetEmailPhoneByUserId(ctx context.Context, userID uint) (response.UserContact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmailPhoneByUserId", ctx, userID)
	ret0, _ := ret[0].(response.UserContact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailPhoneByUserId indicates an expected call of GetEmailPhoneByUserId.
func (mr *MockUserRepositoryMockRecorder) GetEmailPhoneByUserId(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailPhoneByUserId", reflect.TypeOf((*MockUserRepository)(nil).GetEmailPhoneByUserId), ctx, userID)
}

// GetUserbyID mocks base method.
func (m *MockUserRepository) GetUserbyID(ctx context.Context, userId uint) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserbyID", ctx, userId)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserbyID indicates an expected call of GetUserbyID.
func (mr *MockUserRepositoryMockRecorder) GetUserbyID(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserbyID", reflect.TypeOf((*MockUserRepository)(nil).GetUserbyID), ctx, userId)
}

// GetWalletHistory mocks base method.
func (m *MockUserRepository) GetWalletHistory(ctx context.Context, userId uint) ([]domain.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletHistory", ctx, userId)
	ret0, _ := ret[0].([]domain.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletHistory indicates an expected call of GetWalletHistory.
func (mr *MockUserRepositoryMockRecorder) GetWalletHistory(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletHistory", reflect.TypeOf((*MockUserRepository)(nil).GetWalletHistory), ctx, userId)
}

// GetWishlist mocks base method.
func (m *MockUserRepository) GetWishlist(ctx context.Context, userId uint) ([]response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishlist", ctx, userId)
	ret0, _ := ret[0].([]response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishlist indicates an expected call of GetWishlist.
func (mr *MockUserRepositoryMockRecorder) GetWishlist(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishlist", reflect.TypeOf((*MockUserRepository)(nil).GetWishlist), ctx, userId)
}

// RemoveCartItem mocks base method.
func (m *MockUserRepository) RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCartItem", ctx, DelCartItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCartItem indicates an expected call of RemoveCartItem.
func (mr *MockUserRepositoryMockRecorder) RemoveCartItem(ctx, DelCartItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCartItem", reflect.TypeOf((*MockUserRepository)(nil).RemoveCartItem), ctx, DelCartItem)
}

// SaveAddress mocks base method.
func (m *MockUserRepository) SaveAddress(ctx context.Context, userAddress request.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAddress", ctx, userAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAddress indicates an expected call of SaveAddress.
func (mr *MockUserRepositoryMockRecorder) SaveAddress(ctx, userAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAddress", reflect.TypeOf((*MockUserRepository)(nil).SaveAddress), ctx, userAddress)
}

// SavetoCart mocks base method.
func (m *MockUserRepository) SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavetoCart", ctx, addToCart)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavetoCart indicates an expected call of SavetoCart.
func (mr *MockUserRepositoryMockRecorder) SavetoCart(ctx, addToCart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavetoCart", reflect.TypeOf((*MockUserRepository)(nil).SavetoCart), ctx, addToCart)
}

// UpdateAddress mocks base method.
func (m *MockUserRepository) UpdateAddress(ctx context.Context, userAddress request.AddressPatchReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", ctx, userAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserRepositoryMockRecorder) UpdateAddress(ctx, userAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserRepository)(nil).UpdateAddress), ctx, userAddress)
}

// UpdateCart mocks base method.
func (m *MockUserRepository) UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, cartUpadates)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockUserRepositoryMockRecorder) UpdateCart(ctx, cartUpadates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockUserRepository)(nil).UpdateCart), ctx, cartUpadates)
}
