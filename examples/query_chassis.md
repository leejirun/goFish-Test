package main

import (
	"fmt"

	"github.com/stmcginnis/gofish"
)

func main() {
	// 접속 정보 정의
	config := gofish.ClientConfig{
		Endpoint: "https://10.10.1.42:443",
		Username: "hrgomp",
		Password: "hpinvent",
		Insecure: true,
	}

	// 접속, 성공시 c는 Connection정보, err은 실패시 에러 정보
	c, err := gofish.Connect(config)

	// c, err := gofish.ConnectDefault("https://10.10.1.42:443")
	// err이 nil이면 실패했다는 뜻, 에러 출력
	if err != nil {
		panic(err)
	}

	// Connection에서 Service 정보
	service := c.Service
	//Chassis는 뭔지 모르겠으나 어떤 정보들을 조회하는 뜻인듯.
	chassis, err := service.Chassis()
	if err != nil {
		panic(err)
	}
	// 조회된 정보들 for문으로 돌면서 print
	for _, chass := range chassis {
		fmt.Printf("Chassis: %#v\n\n", chass)
	}
}
