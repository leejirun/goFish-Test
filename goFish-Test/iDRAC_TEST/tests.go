package main

import (
	"fmt"

	"github.com/stmcginnis/gofish"
)

func main() {
	// 접속 정보 정의
	config := gofish.ClientConfig{
		Endpoint: "https://10.10.1.102:443",
		Username: "root ",
		Password: "calvin",
		Insecure: true,
	}

	c, err := gofish.Connect(config)

	if err != nil {
		panic(err)
	}
	defer c.Logout()
	fmt.Printf("connect information: %#v\n\n", c)

	service := c.Service

	/*셰시*/
	chassis, err := service.Chassis()
	if err != nil {
		panic(err)
	}

	for _, chass := range chassis {
		fmt.Printf("Chassis: %#v\n\n", chass)
	}

	/*컴퓨터 시스템*/
	comsys, err := chassis[0].ComputerSystems()
	if err != nil {
		panic(err)
	}

	for _, cs := range comsys {
		fmt.Printf("comsys: %#v\n\n", cs)
	}

	/*프로세스*/
	processors, err := comsys[0].Processors()
	if err != nil {
		panic(err)
	}

	for _, processor := range processors {
		fmt.Printf("proceessor 전체 정보 : %#v\n\n", processor)

		fmt.Printf("모델 명 : %#v\n\n", processor.Model)
		fmt.Printf("코어 수 : %#v\n\n", processor.TotalCores)
		fmt.Printf("사용 가능한 코어 수 : %#v\n\n", processor.TotalEnabledCores)
		fmt.Printf("스레드 수 : %#v\n\n", processor.TotalThreads)

	}

	/*전원 상태*/
	power, err := chassis[0].Power()

	if err != nil {
		panic(err)
	}

	fmt.Println(power.PowerControl)
	fmt.Printf("power status(전원 상태) : %#v\n\n", comsys[0].PowerState)
	fmt.Printf("processor 요약 : %#v\n\n", comsys[0].ProcessorSummary)

	/*Thermal*/
	thermals, err := chassis[0].Thermal()
	if err != nil {
		panic(err)
	}

	fmt.Println(thermals)
	fmt.Println(thermals.Status.Health)
	fmt.Pirntln(thermals.Entity)
	fmt.Pirntln(thermals.ID)

	/*dep*/
	dep, err := chassis[0].Drives()
	if err != nil {
		panic(err)
	}

	for _, deps := range dep {
		fmt.Printf(deps.PartNumber)
		fmt.Printf(deps.ID)
		fmt.Printf(deps.SerialNumber)
		fmt.Printf(deps.Name)

	}
	fmt.Println(dep)

}
