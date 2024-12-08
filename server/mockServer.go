package server

type MockServer struct{}

func NewMockServer() Server {
	return &MockServer{}
}

func (m *MockServer) Start() {}

func (m *MockServer) Stop() {}
