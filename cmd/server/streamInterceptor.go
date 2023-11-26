package main

import (
	"log"

	"google.golang.org/grpc"
)

func streamLogging() grpc.StreamServerInterceptor {
	return func (srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// ストリームがopenされたときに行われる前処理
		log.Println("[pre stream] my stream server interceptor 1: ", info.FullMethod)

		err := handler(srv, &myServerStreamWrapper1{ss}) // 本来のストリーム処理

		// ストリームがcloseされるときに行われる後処理
		log.Println("[post stream] my stream server interceptor 1: ")
		return err
	}
}
