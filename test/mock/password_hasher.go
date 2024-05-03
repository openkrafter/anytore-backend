package mock

import "github.com/openkrafter/anytore-backend/auth"

type MockHasher struct{}

func NewMockHasher() *MockHasher {
	mockHasher := &MockHasher{}
	auth.PassHasher = mockHasher
	return mockHasher
}

func (h *MockHasher) HashPassword(password string) (string, error) {
	return "hash_value", nil
}
