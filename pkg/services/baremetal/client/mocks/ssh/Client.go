// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	sshclient "github.com/syself/cluster-api-provider-hetzner/pkg/services/baremetal/client/ssh"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CheckCloudInitLogsForSigTerm provides a mock function with given fields:
func (_m *Client) CheckCloudInitLogsForSigTerm() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// CleanCloudInitInstances provides a mock function with given fields:
func (_m *Client) CleanCloudInitInstances() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// CleanCloudInitLogs provides a mock function with given fields:
func (_m *Client) CleanCloudInitLogs() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// CloudInitStatus provides a mock function with given fields:
func (_m *Client) CloudInitStatus() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// CreateAutoSetup provides a mock function with given fields: data
func (_m *Client) CreateAutoSetup(data string) sshclient.Output {
	ret := _m.Called(data)

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func(string) sshclient.Output); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// CreateMetaData provides a mock function with given fields: hostName
func (_m *Client) CreateMetaData(hostName string) sshclient.Output {
	ret := _m.Called(hostName)

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func(string) sshclient.Output); ok {
		r0 = rf(hostName)
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// CreateNoCloudDirectory provides a mock function with given fields:
func (_m *Client) CreateNoCloudDirectory() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// CreatePostInstallScript provides a mock function with given fields: data
func (_m *Client) CreatePostInstallScript(data string) sshclient.Output {
	ret := _m.Called(data)

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func(string) sshclient.Output); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// CreateUserData provides a mock function with given fields: userData
func (_m *Client) CreateUserData(userData string) sshclient.Output {
	ret := _m.Called(userData)

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func(string) sshclient.Output); ok {
		r0 = rf(userData)
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// DownloadImage provides a mock function with given fields: path, url
func (_m *Client) DownloadImage(path string, url string) sshclient.Output {
	ret := _m.Called(path, url)

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func(string, string) sshclient.Output); ok {
		r0 = rf(path, url)
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// EnsureCloudInit provides a mock function with given fields:
func (_m *Client) EnsureCloudInit() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// ExecuteInstallImage provides a mock function with given fields: hasPostInstallScript
func (_m *Client) ExecuteInstallImage(hasPostInstallScript bool) sshclient.Output {
	ret := _m.Called(hasPostInstallScript)

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func(bool) sshclient.Output); ok {
		r0 = rf(hasPostInstallScript)
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsCPUArch provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsCPUArch() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsCPUClockGigahertz provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsCPUClockGigahertz() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsCPUCores provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsCPUCores() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsCPUFlags provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsCPUFlags() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsCPUModel provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsCPUModel() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsCPUThreads provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsCPUThreads() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsNics provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsNics() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsRAM provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsRAM() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHardwareDetailsStorage provides a mock function with given fields:
func (_m *Client) GetHardwareDetailsStorage() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// GetHostName provides a mock function with given fields:
func (_m *Client) GetHostName() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// Reboot provides a mock function with given fields:
func (_m *Client) Reboot() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

// ResetKubeadm provides a mock function with given fields:
func (_m *Client) ResetKubeadm() sshclient.Output {
	ret := _m.Called()

	var r0 sshclient.Output
	if rf, ok := ret.Get(0).(func() sshclient.Output); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sshclient.Output)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
