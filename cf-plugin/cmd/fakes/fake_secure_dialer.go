// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/diego-ssh/cf-plugin/cmd"
	"golang.org/x/crypto/ssh"
)

type FakeSecureDialer struct {
	DialStub        func(network string, address string, config *ssh.ClientConfig) (ssh.Conn, cmd.SecureClient, error)
	dialMutex       sync.RWMutex
	dialArgsForCall []struct {
		network string
		address string
		config  *ssh.ClientConfig
	}
	dialReturns struct {
		result1 ssh.Conn
		result2 cmd.SecureClient
		result3 error
	}
}

func (fake *FakeSecureDialer) Dial(network string, address string, config *ssh.ClientConfig) (ssh.Conn, cmd.SecureClient, error) {
	fake.dialMutex.Lock()
	fake.dialArgsForCall = append(fake.dialArgsForCall, struct {
		network string
		address string
		config  *ssh.ClientConfig
	}{network, address, config})
	fake.dialMutex.Unlock()
	if fake.DialStub != nil {
		return fake.DialStub(network, address, config)
	} else {
		return fake.dialReturns.result1, fake.dialReturns.result2, fake.dialReturns.result3
	}
}

func (fake *FakeSecureDialer) DialCallCount() int {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return len(fake.dialArgsForCall)
}

func (fake *FakeSecureDialer) DialArgsForCall(i int) (string, string, *ssh.ClientConfig) {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return fake.dialArgsForCall[i].network, fake.dialArgsForCall[i].address, fake.dialArgsForCall[i].config
}

func (fake *FakeSecureDialer) DialReturns(result1 ssh.Conn, result2 cmd.SecureClient, result3 error) {
	fake.DialStub = nil
	fake.dialReturns = struct {
		result1 ssh.Conn
		result2 cmd.SecureClient
		result3 error
	}{result1, result2, result3}
}

var _ cmd.SecureDialer = new(FakeSecureDialer)
