package main

import (
	"fmt"
	"time"
)

var c = make(chan RoleReq,1024)
var batchRoleReq = 20
var batchFunc =20

type RoleReq struct {
	zone string
	server string
	rid string
}

func crc32 (s string) int{
	return 1
}
func getRoleInfoByGid(roleReq RoleReq) string{
	key := crc32(roleReq.zone+roleReq.server)
	hkey := fmt.Sprintf("%x",key)
	//get from cache redis hash cr:{hkey}:{rid} ={roleinfo}

	exist:=false
	if exist{
		return "" //{roleinfo}
	}

	produce(roleReq)
	return ""
}

func produce(roleReq RoleReq){
	//check is is exist wait cache key ck:{hkey}:{rid}=1
	exist := false
	if exist{
		return
	}

	select {
	case c<-roleReq:
		//set wait cache key ck:{hkey}:{rid}=1 TTL 1min
	default:
		//chan is full
	}
}

func consumerWorker(){
	for{
		fs := make([]func(),0,batchFunc)
		for i := 0; i < batchFunc; i++ {
			roleReqs := make([]RoleReq,0,batchRoleReq)
			for j := 0; j < batchRoleReq; j++ {
				roleReq,isok := <-c

				if !isok{
					return
				}

				roleReqs = append(roleReqs,roleReq)
			}

			fs = append(fs, func() {
				reqRoleInfos(roleReqs)
			})
		}

		result := AsyncInvokesWithTimeout(time.Second*10,fs)
		fmt.Println(result)
	}
}

func reqRoleInfos(roleReqs []RoleReq){
	//req ok
	//cache to redis hash  cr:{hkey}:{rid}={roleinfo}
	//clear wait cache key wgid:gid =1
}
