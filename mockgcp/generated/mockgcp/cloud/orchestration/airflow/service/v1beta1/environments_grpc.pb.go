// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/orchestration/airflow/service/v1beta1/environments.proto

package servicepb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EnvironmentsClient is the client API for Environments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentsClient interface {
	// Create a new environment.
	CreateEnvironment(ctx context.Context, in *CreateEnvironmentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Get an existing environment.
	GetEnvironment(ctx context.Context, in *GetEnvironmentRequest, opts ...grpc.CallOption) (*Environment, error)
	// List environments.
	ListEnvironments(ctx context.Context, in *ListEnvironmentsRequest, opts ...grpc.CallOption) (*ListEnvironmentsResponse, error)
	// Update an environment.
	UpdateEnvironment(ctx context.Context, in *UpdateEnvironmentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Delete an environment.
	DeleteEnvironment(ctx context.Context, in *DeleteEnvironmentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Restart Airflow web server.
	RestartWebServer(ctx context.Context, in *RestartWebServerRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Check if an upgrade operation on the environment will succeed.
	//
	// In case of problems detailed info can be found in the returned Operation.
	CheckUpgrade(ctx context.Context, in *CheckUpgradeRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Executes Airflow CLI command.
	ExecuteAirflowCommand(ctx context.Context, in *ExecuteAirflowCommandRequest, opts ...grpc.CallOption) (*ExecuteAirflowCommandResponse, error)
	// Stops Airflow CLI command execution.
	StopAirflowCommand(ctx context.Context, in *StopAirflowCommandRequest, opts ...grpc.CallOption) (*StopAirflowCommandResponse, error)
	// Polls Airflow CLI command execution and fetches logs.
	PollAirflowCommand(ctx context.Context, in *PollAirflowCommandRequest, opts ...grpc.CallOption) (*PollAirflowCommandResponse, error)
	// Creates a snapshots of a Cloud Composer environment.
	//
	// As a result of this operation, snapshot of environment's state is stored
	// in a location specified in the SaveSnapshotRequest.
	SaveSnapshot(ctx context.Context, in *SaveSnapshotRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Loads a snapshot of a Cloud Composer environment.
	//
	// As a result of this operation, a snapshot of environment's specified in
	// LoadSnapshotRequest is loaded into the environment.
	LoadSnapshot(ctx context.Context, in *LoadSnapshotRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Triggers database failover (only for highly resilient environments).
	DatabaseFailover(ctx context.Context, in *DatabaseFailoverRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Fetches database properties.
	FetchDatabaseProperties(ctx context.Context, in *FetchDatabasePropertiesRequest, opts ...grpc.CallOption) (*FetchDatabasePropertiesResponse, error)
}

type environmentsClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentsClient(cc grpc.ClientConnInterface) EnvironmentsClient {
	return &environmentsClient{cc}
}

func (c *environmentsClient) CreateEnvironment(ctx context.Context, in *CreateEnvironmentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/CreateEnvironment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) GetEnvironment(ctx context.Context, in *GetEnvironmentRequest, opts ...grpc.CallOption) (*Environment, error) {
	out := new(Environment)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/GetEnvironment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) ListEnvironments(ctx context.Context, in *ListEnvironmentsRequest, opts ...grpc.CallOption) (*ListEnvironmentsResponse, error) {
	out := new(ListEnvironmentsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/ListEnvironments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) UpdateEnvironment(ctx context.Context, in *UpdateEnvironmentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/UpdateEnvironment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) DeleteEnvironment(ctx context.Context, in *DeleteEnvironmentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/DeleteEnvironment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) RestartWebServer(ctx context.Context, in *RestartWebServerRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/RestartWebServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) CheckUpgrade(ctx context.Context, in *CheckUpgradeRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/CheckUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) ExecuteAirflowCommand(ctx context.Context, in *ExecuteAirflowCommandRequest, opts ...grpc.CallOption) (*ExecuteAirflowCommandResponse, error) {
	out := new(ExecuteAirflowCommandResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/ExecuteAirflowCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) StopAirflowCommand(ctx context.Context, in *StopAirflowCommandRequest, opts ...grpc.CallOption) (*StopAirflowCommandResponse, error) {
	out := new(StopAirflowCommandResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/StopAirflowCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) PollAirflowCommand(ctx context.Context, in *PollAirflowCommandRequest, opts ...grpc.CallOption) (*PollAirflowCommandResponse, error) {
	out := new(PollAirflowCommandResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/PollAirflowCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) SaveSnapshot(ctx context.Context, in *SaveSnapshotRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/SaveSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) LoadSnapshot(ctx context.Context, in *LoadSnapshotRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/LoadSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) DatabaseFailover(ctx context.Context, in *DatabaseFailoverRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/DatabaseFailover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) FetchDatabaseProperties(ctx context.Context, in *FetchDatabasePropertiesRequest, opts ...grpc.CallOption) (*FetchDatabasePropertiesResponse, error) {
	out := new(FetchDatabasePropertiesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/FetchDatabaseProperties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentsServer is the server API for Environments service.
// All implementations must embed UnimplementedEnvironmentsServer
// for forward compatibility
type EnvironmentsServer interface {
	// Create a new environment.
	CreateEnvironment(context.Context, *CreateEnvironmentRequest) (*longrunningpb.Operation, error)
	// Get an existing environment.
	GetEnvironment(context.Context, *GetEnvironmentRequest) (*Environment, error)
	// List environments.
	ListEnvironments(context.Context, *ListEnvironmentsRequest) (*ListEnvironmentsResponse, error)
	// Update an environment.
	UpdateEnvironment(context.Context, *UpdateEnvironmentRequest) (*longrunningpb.Operation, error)
	// Delete an environment.
	DeleteEnvironment(context.Context, *DeleteEnvironmentRequest) (*longrunningpb.Operation, error)
	// Restart Airflow web server.
	RestartWebServer(context.Context, *RestartWebServerRequest) (*longrunningpb.Operation, error)
	// Check if an upgrade operation on the environment will succeed.
	//
	// In case of problems detailed info can be found in the returned Operation.
	CheckUpgrade(context.Context, *CheckUpgradeRequest) (*longrunningpb.Operation, error)
	// Executes Airflow CLI command.
	ExecuteAirflowCommand(context.Context, *ExecuteAirflowCommandRequest) (*ExecuteAirflowCommandResponse, error)
	// Stops Airflow CLI command execution.
	StopAirflowCommand(context.Context, *StopAirflowCommandRequest) (*StopAirflowCommandResponse, error)
	// Polls Airflow CLI command execution and fetches logs.
	PollAirflowCommand(context.Context, *PollAirflowCommandRequest) (*PollAirflowCommandResponse, error)
	// Creates a snapshots of a Cloud Composer environment.
	//
	// As a result of this operation, snapshot of environment's state is stored
	// in a location specified in the SaveSnapshotRequest.
	SaveSnapshot(context.Context, *SaveSnapshotRequest) (*longrunningpb.Operation, error)
	// Loads a snapshot of a Cloud Composer environment.
	//
	// As a result of this operation, a snapshot of environment's specified in
	// LoadSnapshotRequest is loaded into the environment.
	LoadSnapshot(context.Context, *LoadSnapshotRequest) (*longrunningpb.Operation, error)
	// Triggers database failover (only for highly resilient environments).
	DatabaseFailover(context.Context, *DatabaseFailoverRequest) (*longrunningpb.Operation, error)
	// Fetches database properties.
	FetchDatabaseProperties(context.Context, *FetchDatabasePropertiesRequest) (*FetchDatabasePropertiesResponse, error)
	mustEmbedUnimplementedEnvironmentsServer()
}

// UnimplementedEnvironmentsServer must be embedded to have forward compatible implementations.
type UnimplementedEnvironmentsServer struct {
}

func (UnimplementedEnvironmentsServer) CreateEnvironment(context.Context, *CreateEnvironmentRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEnvironment not implemented")
}
func (UnimplementedEnvironmentsServer) GetEnvironment(context.Context, *GetEnvironmentRequest) (*Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvironment not implemented")
}
func (UnimplementedEnvironmentsServer) ListEnvironments(context.Context, *ListEnvironmentsRequest) (*ListEnvironmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnvironments not implemented")
}
func (UnimplementedEnvironmentsServer) UpdateEnvironment(context.Context, *UpdateEnvironmentRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEnvironment not implemented")
}
func (UnimplementedEnvironmentsServer) DeleteEnvironment(context.Context, *DeleteEnvironmentRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEnvironment not implemented")
}
func (UnimplementedEnvironmentsServer) RestartWebServer(context.Context, *RestartWebServerRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartWebServer not implemented")
}
func (UnimplementedEnvironmentsServer) CheckUpgrade(context.Context, *CheckUpgradeRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUpgrade not implemented")
}
func (UnimplementedEnvironmentsServer) ExecuteAirflowCommand(context.Context, *ExecuteAirflowCommandRequest) (*ExecuteAirflowCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteAirflowCommand not implemented")
}
func (UnimplementedEnvironmentsServer) StopAirflowCommand(context.Context, *StopAirflowCommandRequest) (*StopAirflowCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopAirflowCommand not implemented")
}
func (UnimplementedEnvironmentsServer) PollAirflowCommand(context.Context, *PollAirflowCommandRequest) (*PollAirflowCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PollAirflowCommand not implemented")
}
func (UnimplementedEnvironmentsServer) SaveSnapshot(context.Context, *SaveSnapshotRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSnapshot not implemented")
}
func (UnimplementedEnvironmentsServer) LoadSnapshot(context.Context, *LoadSnapshotRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadSnapshot not implemented")
}
func (UnimplementedEnvironmentsServer) DatabaseFailover(context.Context, *DatabaseFailoverRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabaseFailover not implemented")
}
func (UnimplementedEnvironmentsServer) FetchDatabaseProperties(context.Context, *FetchDatabasePropertiesRequest) (*FetchDatabasePropertiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchDatabaseProperties not implemented")
}
func (UnimplementedEnvironmentsServer) mustEmbedUnimplementedEnvironmentsServer() {}

// UnsafeEnvironmentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentsServer will
// result in compilation errors.
type UnsafeEnvironmentsServer interface {
	mustEmbedUnimplementedEnvironmentsServer()
}

func RegisterEnvironmentsServer(s grpc.ServiceRegistrar, srv EnvironmentsServer) {
	s.RegisterService(&Environments_ServiceDesc, srv)
}

func _Environments_CreateEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).CreateEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/CreateEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).CreateEnvironment(ctx, req.(*CreateEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_GetEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).GetEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/GetEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).GetEnvironment(ctx, req.(*GetEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_ListEnvironments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnvironmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).ListEnvironments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/ListEnvironments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).ListEnvironments(ctx, req.(*ListEnvironmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_UpdateEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).UpdateEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/UpdateEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).UpdateEnvironment(ctx, req.(*UpdateEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_DeleteEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).DeleteEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/DeleteEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).DeleteEnvironment(ctx, req.(*DeleteEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_RestartWebServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartWebServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).RestartWebServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/RestartWebServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).RestartWebServer(ctx, req.(*RestartWebServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_CheckUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).CheckUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/CheckUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).CheckUpgrade(ctx, req.(*CheckUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_ExecuteAirflowCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteAirflowCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).ExecuteAirflowCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/ExecuteAirflowCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).ExecuteAirflowCommand(ctx, req.(*ExecuteAirflowCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_StopAirflowCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAirflowCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).StopAirflowCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/StopAirflowCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).StopAirflowCommand(ctx, req.(*StopAirflowCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_PollAirflowCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollAirflowCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).PollAirflowCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/PollAirflowCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).PollAirflowCommand(ctx, req.(*PollAirflowCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_SaveSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).SaveSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/SaveSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).SaveSnapshot(ctx, req.(*SaveSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_LoadSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).LoadSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/LoadSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).LoadSnapshot(ctx, req.(*LoadSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_DatabaseFailover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseFailoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).DatabaseFailover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/DatabaseFailover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).DatabaseFailover(ctx, req.(*DatabaseFailoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_FetchDatabaseProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchDatabasePropertiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).FetchDatabaseProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments/FetchDatabaseProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).FetchDatabaseProperties(ctx, req.(*FetchDatabasePropertiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Environments_ServiceDesc is the grpc.ServiceDesc for Environments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Environments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.orchestration.airflow.service.v1beta1.Environments",
	HandlerType: (*EnvironmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEnvironment",
			Handler:    _Environments_CreateEnvironment_Handler,
		},
		{
			MethodName: "GetEnvironment",
			Handler:    _Environments_GetEnvironment_Handler,
		},
		{
			MethodName: "ListEnvironments",
			Handler:    _Environments_ListEnvironments_Handler,
		},
		{
			MethodName: "UpdateEnvironment",
			Handler:    _Environments_UpdateEnvironment_Handler,
		},
		{
			MethodName: "DeleteEnvironment",
			Handler:    _Environments_DeleteEnvironment_Handler,
		},
		{
			MethodName: "RestartWebServer",
			Handler:    _Environments_RestartWebServer_Handler,
		},
		{
			MethodName: "CheckUpgrade",
			Handler:    _Environments_CheckUpgrade_Handler,
		},
		{
			MethodName: "ExecuteAirflowCommand",
			Handler:    _Environments_ExecuteAirflowCommand_Handler,
		},
		{
			MethodName: "StopAirflowCommand",
			Handler:    _Environments_StopAirflowCommand_Handler,
		},
		{
			MethodName: "PollAirflowCommand",
			Handler:    _Environments_PollAirflowCommand_Handler,
		},
		{
			MethodName: "SaveSnapshot",
			Handler:    _Environments_SaveSnapshot_Handler,
		},
		{
			MethodName: "LoadSnapshot",
			Handler:    _Environments_LoadSnapshot_Handler,
		},
		{
			MethodName: "DatabaseFailover",
			Handler:    _Environments_DatabaseFailover_Handler,
		},
		{
			MethodName: "FetchDatabaseProperties",
			Handler:    _Environments_FetchDatabaseProperties_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/orchestration/airflow/service/v1beta1/environments.proto",
}