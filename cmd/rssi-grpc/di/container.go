package di

import wiregrpc "git.cie.com/ips/wire-provider/grpc"

type Container struct {
	server wiregrpc.Server
}
