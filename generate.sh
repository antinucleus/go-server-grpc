#!/usr/bin/bash
protoc login/loginpb/login.proto --go_out=. --go-grpc_out=.
