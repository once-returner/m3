// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/storage/bootstrap/types.go

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package bootstrap is a generated GoMock package.
package bootstrap

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/dbnode/storage/bootstrap/result"
	"github.com/m3db/m3/src/dbnode/storage/namespace"
	"github.com/m3db/m3/src/dbnode/topology"

	"github.com/golang/mock/gomock"
)

// MockProcessProvider is a mock of ProcessProvider interface
type MockProcessProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProcessProviderMockRecorder
}

// MockProcessProviderMockRecorder is the mock recorder for MockProcessProvider
type MockProcessProviderMockRecorder struct {
	mock *MockProcessProvider
}

// NewMockProcessProvider creates a new mock instance
func NewMockProcessProvider(ctrl *gomock.Controller) *MockProcessProvider {
	mock := &MockProcessProvider{ctrl: ctrl}
	mock.recorder = &MockProcessProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessProvider) EXPECT() *MockProcessProviderMockRecorder {
	return m.recorder
}

// SetBootstrapperProvider mocks base method
func (m *MockProcessProvider) SetBootstrapperProvider(bootstrapper BootstrapperProvider) {
	m.ctrl.Call(m, "SetBootstrapperProvider", bootstrapper)
}

// SetBootstrapperProvider indicates an expected call of SetBootstrapperProvider
func (mr *MockProcessProviderMockRecorder) SetBootstrapperProvider(bootstrapper interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBootstrapperProvider", reflect.TypeOf((*MockProcessProvider)(nil).SetBootstrapperProvider), bootstrapper)
}

// BootstrapperProvider mocks base method
func (m *MockProcessProvider) BootstrapperProvider() BootstrapperProvider {
	ret := m.ctrl.Call(m, "BootstrapperProvider")
	ret0, _ := ret[0].(BootstrapperProvider)
	return ret0
}

// BootstrapperProvider indicates an expected call of BootstrapperProvider
func (mr *MockProcessProviderMockRecorder) BootstrapperProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapperProvider", reflect.TypeOf((*MockProcessProvider)(nil).BootstrapperProvider))
}

// Provide mocks base method
func (m *MockProcessProvider) Provide() (Process, error) {
	ret := m.ctrl.Call(m, "Provide")
	ret0, _ := ret[0].(Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provide indicates an expected call of Provide
func (mr *MockProcessProviderMockRecorder) Provide() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockProcessProvider)(nil).Provide))
}

// MockProcess is a mock of Process interface
type MockProcess struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMockRecorder
}

// MockProcessMockRecorder is the mock recorder for MockProcess
type MockProcessMockRecorder struct {
	mock *MockProcess
}

// NewMockProcess creates a new mock instance
func NewMockProcess(ctrl *gomock.Controller) *MockProcess {
	mock := &MockProcess{ctrl: ctrl}
	mock.recorder = &MockProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcess) EXPECT() *MockProcessMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockProcess) Run(start time.Time, ns namespace.Metadata, shards []uint32) (ProcessResult, error) {
	ret := m.ctrl.Call(m, "Run", start, ns, shards)
	ret0, _ := ret[0].(ProcessResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockProcessMockRecorder) Run(start, ns, shards interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockProcess)(nil).Run), start, ns, shards)
}

// MockProcessOptions is a mock of ProcessOptions interface
type MockProcessOptions struct {
	ctrl     *gomock.Controller
	recorder *MockProcessOptionsMockRecorder
}

// MockProcessOptionsMockRecorder is the mock recorder for MockProcessOptions
type MockProcessOptionsMockRecorder struct {
	mock *MockProcessOptions
}

// NewMockProcessOptions creates a new mock instance
func NewMockProcessOptions(ctrl *gomock.Controller) *MockProcessOptions {
	mock := &MockProcessOptions{ctrl: ctrl}
	mock.recorder = &MockProcessOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessOptions) EXPECT() *MockProcessOptionsMockRecorder {
	return m.recorder
}

// SetCacheSeriesMetadata mocks base method
func (m *MockProcessOptions) SetCacheSeriesMetadata(value bool) ProcessOptions {
	ret := m.ctrl.Call(m, "SetCacheSeriesMetadata", value)
	ret0, _ := ret[0].(ProcessOptions)
	return ret0
}

// SetCacheSeriesMetadata indicates an expected call of SetCacheSeriesMetadata
func (mr *MockProcessOptionsMockRecorder) SetCacheSeriesMetadata(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCacheSeriesMetadata", reflect.TypeOf((*MockProcessOptions)(nil).SetCacheSeriesMetadata), value)
}

// CacheSeriesMetadata mocks base method
func (m *MockProcessOptions) CacheSeriesMetadata() bool {
	ret := m.ctrl.Call(m, "CacheSeriesMetadata")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CacheSeriesMetadata indicates an expected call of CacheSeriesMetadata
func (mr *MockProcessOptionsMockRecorder) CacheSeriesMetadata() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheSeriesMetadata", reflect.TypeOf((*MockProcessOptions)(nil).CacheSeriesMetadata))
}

// SetTopologyMapProvider mocks base method
func (m *MockProcessOptions) SetTopologyMapProvider(value topology.MapProvider) ProcessOptions {
	ret := m.ctrl.Call(m, "SetTopologyMapProvider", value)
	ret0, _ := ret[0].(ProcessOptions)
	return ret0
}

// SetTopologyMapProvider indicates an expected call of SetTopologyMapProvider
func (mr *MockProcessOptionsMockRecorder) SetTopologyMapProvider(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTopologyMapProvider", reflect.TypeOf((*MockProcessOptions)(nil).SetTopologyMapProvider), value)
}

// TopologyMapProvider mocks base method
func (m *MockProcessOptions) TopologyMapProvider() topology.MapProvider {
	ret := m.ctrl.Call(m, "TopologyMapProvider")
	ret0, _ := ret[0].(topology.MapProvider)
	return ret0
}

// TopologyMapProvider indicates an expected call of TopologyMapProvider
func (mr *MockProcessOptionsMockRecorder) TopologyMapProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopologyMapProvider", reflect.TypeOf((*MockProcessOptions)(nil).TopologyMapProvider))
}

// SetOrigin mocks base method
func (m *MockProcessOptions) SetOrigin(value topology.Host) ProcessOptions {
	ret := m.ctrl.Call(m, "SetOrigin", value)
	ret0, _ := ret[0].(ProcessOptions)
	return ret0
}

// SetOrigin indicates an expected call of SetOrigin
func (mr *MockProcessOptionsMockRecorder) SetOrigin(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrigin", reflect.TypeOf((*MockProcessOptions)(nil).SetOrigin), value)
}

// Origin mocks base method
func (m *MockProcessOptions) Origin() topology.Host {
	ret := m.ctrl.Call(m, "Origin")
	ret0, _ := ret[0].(topology.Host)
	return ret0
}

// Origin indicates an expected call of Origin
func (mr *MockProcessOptionsMockRecorder) Origin() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Origin", reflect.TypeOf((*MockProcessOptions)(nil).Origin))
}

// Validate mocks base method
func (m *MockProcessOptions) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockProcessOptionsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockProcessOptions)(nil).Validate))
}

// MockRunOptions is a mock of RunOptions interface
type MockRunOptions struct {
	ctrl     *gomock.Controller
	recorder *MockRunOptionsMockRecorder
}

// MockRunOptionsMockRecorder is the mock recorder for MockRunOptions
type MockRunOptionsMockRecorder struct {
	mock *MockRunOptions
}

// NewMockRunOptions creates a new mock instance
func NewMockRunOptions(ctrl *gomock.Controller) *MockRunOptions {
	mock := &MockRunOptions{ctrl: ctrl}
	mock.recorder = &MockRunOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunOptions) EXPECT() *MockRunOptionsMockRecorder {
	return m.recorder
}

// SetPersistConfig mocks base method
func (m *MockRunOptions) SetPersistConfig(value PersistConfig) RunOptions {
	ret := m.ctrl.Call(m, "SetPersistConfig", value)
	ret0, _ := ret[0].(RunOptions)
	return ret0
}

// SetPersistConfig indicates an expected call of SetPersistConfig
func (mr *MockRunOptionsMockRecorder) SetPersistConfig(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPersistConfig", reflect.TypeOf((*MockRunOptions)(nil).SetPersistConfig), value)
}

// PersistConfig mocks base method
func (m *MockRunOptions) PersistConfig() PersistConfig {
	ret := m.ctrl.Call(m, "PersistConfig")
	ret0, _ := ret[0].(PersistConfig)
	return ret0
}

// PersistConfig indicates an expected call of PersistConfig
func (mr *MockRunOptionsMockRecorder) PersistConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistConfig", reflect.TypeOf((*MockRunOptions)(nil).PersistConfig))
}

// SetCacheSeriesMetadata mocks base method
func (m *MockRunOptions) SetCacheSeriesMetadata(value bool) RunOptions {
	ret := m.ctrl.Call(m, "SetCacheSeriesMetadata", value)
	ret0, _ := ret[0].(RunOptions)
	return ret0
}

// SetCacheSeriesMetadata indicates an expected call of SetCacheSeriesMetadata
func (mr *MockRunOptionsMockRecorder) SetCacheSeriesMetadata(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCacheSeriesMetadata", reflect.TypeOf((*MockRunOptions)(nil).SetCacheSeriesMetadata), value)
}

// CacheSeriesMetadata mocks base method
func (m *MockRunOptions) CacheSeriesMetadata() bool {
	ret := m.ctrl.Call(m, "CacheSeriesMetadata")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CacheSeriesMetadata indicates an expected call of CacheSeriesMetadata
func (mr *MockRunOptionsMockRecorder) CacheSeriesMetadata() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheSeriesMetadata", reflect.TypeOf((*MockRunOptions)(nil).CacheSeriesMetadata))
}

// SetInitialTopologyState mocks base method
func (m *MockRunOptions) SetInitialTopologyState(value *topology.StateSnapshot) RunOptions {
	ret := m.ctrl.Call(m, "SetInitialTopologyState", value)
	ret0, _ := ret[0].(RunOptions)
	return ret0
}

// SetInitialTopologyState indicates an expected call of SetInitialTopologyState
func (mr *MockRunOptionsMockRecorder) SetInitialTopologyState(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInitialTopologyState", reflect.TypeOf((*MockRunOptions)(nil).SetInitialTopologyState), value)
}

// InitialTopologyState mocks base method
func (m *MockRunOptions) InitialTopologyState() *topology.StateSnapshot {
	ret := m.ctrl.Call(m, "InitialTopologyState")
	ret0, _ := ret[0].(*topology.StateSnapshot)
	return ret0
}

// InitialTopologyState indicates an expected call of InitialTopologyState
func (mr *MockRunOptionsMockRecorder) InitialTopologyState() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitialTopologyState", reflect.TypeOf((*MockRunOptions)(nil).InitialTopologyState))
}

// MockBootstrapperProvider is a mock of BootstrapperProvider interface
type MockBootstrapperProvider struct {
	ctrl     *gomock.Controller
	recorder *MockBootstrapperProviderMockRecorder
}

// MockBootstrapperProviderMockRecorder is the mock recorder for MockBootstrapperProvider
type MockBootstrapperProviderMockRecorder struct {
	mock *MockBootstrapperProvider
}

// NewMockBootstrapperProvider creates a new mock instance
func NewMockBootstrapperProvider(ctrl *gomock.Controller) *MockBootstrapperProvider {
	mock := &MockBootstrapperProvider{ctrl: ctrl}
	mock.recorder = &MockBootstrapperProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBootstrapperProvider) EXPECT() *MockBootstrapperProviderMockRecorder {
	return m.recorder
}

// String mocks base method
func (m *MockBootstrapperProvider) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockBootstrapperProviderMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockBootstrapperProvider)(nil).String))
}

// Provide mocks base method
func (m *MockBootstrapperProvider) Provide() (Bootstrapper, error) {
	ret := m.ctrl.Call(m, "Provide")
	ret0, _ := ret[0].(Bootstrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provide indicates an expected call of Provide
func (mr *MockBootstrapperProviderMockRecorder) Provide() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockBootstrapperProvider)(nil).Provide))
}

// MockBootstrapper is a mock of Bootstrapper interface
type MockBootstrapper struct {
	ctrl     *gomock.Controller
	recorder *MockBootstrapperMockRecorder
}

// MockBootstrapperMockRecorder is the mock recorder for MockBootstrapper
type MockBootstrapperMockRecorder struct {
	mock *MockBootstrapper
}

// NewMockBootstrapper creates a new mock instance
func NewMockBootstrapper(ctrl *gomock.Controller) *MockBootstrapper {
	mock := &MockBootstrapper{ctrl: ctrl}
	mock.recorder = &MockBootstrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBootstrapper) EXPECT() *MockBootstrapperMockRecorder {
	return m.recorder
}

// String mocks base method
func (m *MockBootstrapper) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockBootstrapperMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockBootstrapper)(nil).String))
}

// Can mocks base method
func (m *MockBootstrapper) Can(strategy Strategy) bool {
	ret := m.ctrl.Call(m, "Can", strategy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Can indicates an expected call of Can
func (mr *MockBootstrapperMockRecorder) Can(strategy interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Can", reflect.TypeOf((*MockBootstrapper)(nil).Can), strategy)
}

// BootstrapData mocks base method
func (m *MockBootstrapper) BootstrapData(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, opts RunOptions) (result.DataBootstrapResult, error) {
	ret := m.ctrl.Call(m, "BootstrapData", ns, shardsTimeRanges, opts)
	ret0, _ := ret[0].(result.DataBootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BootstrapData indicates an expected call of BootstrapData
func (mr *MockBootstrapperMockRecorder) BootstrapData(ns, shardsTimeRanges, opts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapData", reflect.TypeOf((*MockBootstrapper)(nil).BootstrapData), ns, shardsTimeRanges, opts)
}

// BootstrapIndex mocks base method
func (m *MockBootstrapper) BootstrapIndex(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, opts RunOptions) (result.IndexBootstrapResult, error) {
	ret := m.ctrl.Call(m, "BootstrapIndex", ns, shardsTimeRanges, opts)
	ret0, _ := ret[0].(result.IndexBootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BootstrapIndex indicates an expected call of BootstrapIndex
func (mr *MockBootstrapperMockRecorder) BootstrapIndex(ns, shardsTimeRanges, opts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapIndex", reflect.TypeOf((*MockBootstrapper)(nil).BootstrapIndex), ns, shardsTimeRanges, opts)
}

// MockSource is a mock of Source interface
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSource) EXPECT() *MockSourceMockRecorder {
	return m.recorder
}

// Can mocks base method
func (m *MockSource) Can(strategy Strategy) bool {
	ret := m.ctrl.Call(m, "Can", strategy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Can indicates an expected call of Can
func (mr *MockSourceMockRecorder) Can(strategy interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Can", reflect.TypeOf((*MockSource)(nil).Can), strategy)
}

// AvailableData mocks base method
func (m *MockSource) AvailableData(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, runOpts RunOptions) (result.ShardTimeRanges, error) {
	ret := m.ctrl.Call(m, "AvailableData", ns, shardsTimeRanges, runOpts)
	ret0, _ := ret[0].(result.ShardTimeRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvailableData indicates an expected call of AvailableData
func (mr *MockSourceMockRecorder) AvailableData(ns, shardsTimeRanges, runOpts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableData", reflect.TypeOf((*MockSource)(nil).AvailableData), ns, shardsTimeRanges, runOpts)
}

// ReadData mocks base method
func (m *MockSource) ReadData(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, runOpts RunOptions) (result.DataBootstrapResult, error) {
	ret := m.ctrl.Call(m, "ReadData", ns, shardsTimeRanges, runOpts)
	ret0, _ := ret[0].(result.DataBootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadData indicates an expected call of ReadData
func (mr *MockSourceMockRecorder) ReadData(ns, shardsTimeRanges, runOpts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadData", reflect.TypeOf((*MockSource)(nil).ReadData), ns, shardsTimeRanges, runOpts)
}

// AvailableIndex mocks base method
func (m *MockSource) AvailableIndex(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, opts RunOptions) (result.ShardTimeRanges, error) {
	ret := m.ctrl.Call(m, "AvailableIndex", ns, shardsTimeRanges, opts)
	ret0, _ := ret[0].(result.ShardTimeRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvailableIndex indicates an expected call of AvailableIndex
func (mr *MockSourceMockRecorder) AvailableIndex(ns, shardsTimeRanges, opts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableIndex", reflect.TypeOf((*MockSource)(nil).AvailableIndex), ns, shardsTimeRanges, opts)
}

// ReadIndex mocks base method
func (m *MockSource) ReadIndex(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, opts RunOptions) (result.IndexBootstrapResult, error) {
	ret := m.ctrl.Call(m, "ReadIndex", ns, shardsTimeRanges, opts)
	ret0, _ := ret[0].(result.IndexBootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadIndex indicates an expected call of ReadIndex
func (mr *MockSourceMockRecorder) ReadIndex(ns, shardsTimeRanges, opts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadIndex", reflect.TypeOf((*MockSource)(nil).ReadIndex), ns, shardsTimeRanges, opts)
}
