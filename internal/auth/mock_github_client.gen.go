// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package auth

import (
	"context"
	"github.com/google/go-github/v48/github"
	"sync"
)

// Ensure, that GitHubClientMock does implement GitHubClient.
// If this is not the case, regenerate this file with moq.
var _ GitHubClient = &GitHubClientMock{}

// GitHubClientMock is a mock implementation of GitHubClient.
//
// 	func TestSomethingThatUsesGitHubClient(t *testing.T) {
//
// 		// make and configure a mocked GitHubClient
// 		mockedGitHubClient := &GitHubClientMock{
// 			ExchangeCodeToAccessKeyFunc: func(ctx context.Context, clientID string, clientSecret string, code string) (string, error) {
// 				panic("mock out the ExchangeCodeToAccessKey method")
// 			},
// 			GetUserFunc: func(ctx context.Context, accessKey string, user string) (*github.User, error) {
// 				panic("mock out the GetUser method")
// 			},
// 			IsMemberFunc: func(ctx context.Context, accessKey string, org string, user string) (bool, error) {
// 				panic("mock out the IsMember method")
// 			},
// 		}
//
// 		// use mockedGitHubClient in code that requires GitHubClient
// 		// and then make assertions.
//
// 	}
type GitHubClientMock struct {
	// ExchangeCodeToAccessKeyFunc mocks the ExchangeCodeToAccessKey method.
	ExchangeCodeToAccessKeyFunc func(ctx context.Context, clientID string, clientSecret string, code string) (string, error)

	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(ctx context.Context, accessKey string, user string) (*github.User, error)

	// IsMemberFunc mocks the IsMember method.
	IsMemberFunc func(ctx context.Context, accessKey string, org string, user string) (bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// ExchangeCodeToAccessKey holds details about calls to the ExchangeCodeToAccessKey method.
		ExchangeCodeToAccessKey []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ClientID is the clientID argument value.
			ClientID string
			// ClientSecret is the clientSecret argument value.
			ClientSecret string
			// Code is the code argument value.
			Code string
		}
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AccessKey is the accessKey argument value.
			AccessKey string
			// User is the user argument value.
			User string
		}
		// IsMember holds details about calls to the IsMember method.
		IsMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AccessKey is the accessKey argument value.
			AccessKey string
			// Org is the org argument value.
			Org string
			// User is the user argument value.
			User string
		}
	}
	lockExchangeCodeToAccessKey sync.RWMutex
	lockGetUser                 sync.RWMutex
	lockIsMember                sync.RWMutex
}

// ExchangeCodeToAccessKey calls ExchangeCodeToAccessKeyFunc.
func (mock *GitHubClientMock) ExchangeCodeToAccessKey(ctx context.Context, clientID string, clientSecret string, code string) (string, error) {
	if mock.ExchangeCodeToAccessKeyFunc == nil {
		panic("GitHubClientMock.ExchangeCodeToAccessKeyFunc: method is nil but GitHubClient.ExchangeCodeToAccessKey was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		ClientID     string
		ClientSecret string
		Code         string
	}{
		Ctx:          ctx,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
	}
	mock.lockExchangeCodeToAccessKey.Lock()
	mock.calls.ExchangeCodeToAccessKey = append(mock.calls.ExchangeCodeToAccessKey, callInfo)
	mock.lockExchangeCodeToAccessKey.Unlock()
	return mock.ExchangeCodeToAccessKeyFunc(ctx, clientID, clientSecret, code)
}

// ExchangeCodeToAccessKeyCalls gets all the calls that were made to ExchangeCodeToAccessKey.
// Check the length with:
//     len(mockedGitHubClient.ExchangeCodeToAccessKeyCalls())
func (mock *GitHubClientMock) ExchangeCodeToAccessKeyCalls() []struct {
	Ctx          context.Context
	ClientID     string
	ClientSecret string
	Code         string
} {
	var calls []struct {
		Ctx          context.Context
		ClientID     string
		ClientSecret string
		Code         string
	}
	mock.lockExchangeCodeToAccessKey.RLock()
	calls = mock.calls.ExchangeCodeToAccessKey
	mock.lockExchangeCodeToAccessKey.RUnlock()
	return calls
}

// GetUser calls GetUserFunc.
func (mock *GitHubClientMock) GetUser(ctx context.Context, accessKey string, user string) (*github.User, error) {
	if mock.GetUserFunc == nil {
		panic("GitHubClientMock.GetUserFunc: method is nil but GitHubClient.GetUser was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		AccessKey string
		User      string
	}{
		Ctx:       ctx,
		AccessKey: accessKey,
		User:      user,
	}
	mock.lockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	mock.lockGetUser.Unlock()
	return mock.GetUserFunc(ctx, accessKey, user)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//     len(mockedGitHubClient.GetUserCalls())
func (mock *GitHubClientMock) GetUserCalls() []struct {
	Ctx       context.Context
	AccessKey string
	User      string
} {
	var calls []struct {
		Ctx       context.Context
		AccessKey string
		User      string
	}
	mock.lockGetUser.RLock()
	calls = mock.calls.GetUser
	mock.lockGetUser.RUnlock()
	return calls
}

// IsMember calls IsMemberFunc.
func (mock *GitHubClientMock) IsMember(ctx context.Context, accessKey string, org string, user string) (bool, error) {
	if mock.IsMemberFunc == nil {
		panic("GitHubClientMock.IsMemberFunc: method is nil but GitHubClient.IsMember was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		AccessKey string
		Org       string
		User      string
	}{
		Ctx:       ctx,
		AccessKey: accessKey,
		Org:       org,
		User:      user,
	}
	mock.lockIsMember.Lock()
	mock.calls.IsMember = append(mock.calls.IsMember, callInfo)
	mock.lockIsMember.Unlock()
	return mock.IsMemberFunc(ctx, accessKey, org, user)
}

// IsMemberCalls gets all the calls that were made to IsMember.
// Check the length with:
//     len(mockedGitHubClient.IsMemberCalls())
func (mock *GitHubClientMock) IsMemberCalls() []struct {
	Ctx       context.Context
	AccessKey string
	Org       string
	User      string
} {
	var calls []struct {
		Ctx       context.Context
		AccessKey string
		Org       string
		User      string
	}
	mock.lockIsMember.RLock()
	calls = mock.calls.IsMember
	mock.lockIsMember.RUnlock()
	return calls
}