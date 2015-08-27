// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/cloudfoundry-incubator/diego-ssh/cf-plugin/cmd"
	"golang.org/x/crypto/ssh"
)

type FakeSecureClient struct {
	NewSessionStub        func() (cmd.SecureSession, error)
	newSessionMutex       sync.RWMutex
	newSessionArgsForCall []struct{}
	newSessionReturns     struct {
		result1 cmd.SecureSession
		result2 error
	}
	ConnStub        func() ssh.Conn
	connMutex       sync.RWMutex
	connArgsForCall []struct{}
	connReturns     struct {
		result1 ssh.Conn
	}
	DialStub        func(network, address string) (net.Conn, error)
	dialMutex       sync.RWMutex
	dialArgsForCall []struct {
		network string
		address string
	}
	dialReturns struct {
		result1 net.Conn
		result2 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
}

func (fake *FakeSecureClient) NewSession() (cmd.SecureSession, error) {
	fake.newSessionMutex.Lock()
	fake.newSessionArgsForCall = append(fake.newSessionArgsForCall, struct{}{})
	fake.newSessionMutex.Unlock()
	if fake.NewSessionStub != nil {
		return fake.NewSessionStub()
	} else {
		return fake.newSessionReturns.result1, fake.newSessionReturns.result2
	}
}

func (fake *FakeSecureClient) NewSessionCallCount() int {
	fake.newSessionMutex.RLock()
	defer fake.newSessionMutex.RUnlock()
	return len(fake.newSessionArgsForCall)
}

func (fake *FakeSecureClient) NewSessionReturns(result1 cmd.SecureSession, result2 error) {
	fake.NewSessionStub = nil
	fake.newSessionReturns = struct {
		result1 cmd.SecureSession
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureClient) Conn() ssh.Conn {
	fake.connMutex.Lock()
	fake.connArgsForCall = append(fake.connArgsForCall, struct{}{})
	fake.connMutex.Unlock()
	if fake.ConnStub != nil {
		return fake.ConnStub()
	} else {
		return fake.connReturns.result1
	}
}

func (fake *FakeSecureClient) ConnCallCount() int {
	fake.connMutex.RLock()
	defer fake.connMutex.RUnlock()
	return len(fake.connArgsForCall)
}

func (fake *FakeSecureClient) ConnReturns(result1 ssh.Conn) {
	fake.ConnStub = nil
	fake.connReturns = struct {
		result1 ssh.Conn
	}{result1}
}

func (fake *FakeSecureClient) Dial(network string, address string) (net.Conn, error) {
	fake.dialMutex.Lock()
	fake.dialArgsForCall = append(fake.dialArgsForCall, struct {
		network string
		address string
	}{network, address})
	fake.dialMutex.Unlock()
	if fake.DialStub != nil {
		return fake.DialStub(network, address)
	} else {
		return fake.dialReturns.result1, fake.dialReturns.result2
	}
}

func (fake *FakeSecureClient) DialCallCount() int {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return len(fake.dialArgsForCall)
}

func (fake *FakeSecureClient) DialArgsForCall(i int) (string, string) {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return fake.dialArgsForCall[i].network, fake.dialArgsForCall[i].address
}

func (fake *FakeSecureClient) DialReturns(result1 net.Conn, result2 error) {
	fake.DialStub = nil
	fake.dialReturns = struct {
		result1 net.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureClient) Wait() error {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	} else {
		return fake.waitReturns.result1
	}
}

func (fake *FakeSecureClient) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSecureClient) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureClient) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeSecureClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSecureClient) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

var _ cmd.SecureClient = new(FakeSecureClient)
