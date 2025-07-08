package main

import (
	"bufio"
	"flag"
	"fmt"
	"main/grpc"
	"os"
)

var (
	API_URL     = os.Getenv("IPRON_NE_API_URL")
	GRPCAPI_URL = "localhost:8000" // os.Getenv("IPRON_NE_GRPCAPI_URL")
	AppKey      = os.Getenv("IPRON_NE_APPKEY")
)

func main() {
	TenantID := flag.String("tenantid", "", "Tenant ID")
	callID := flag.String("callid", "", "Call ID")
	ani := flag.String("ani", "", "ANI Number")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	config := grpc.Config{}
	config.SetToken(AppKey)
	config.SetGrpcURI(GRPCAPI_URL)
	iwebgw := grpc.NewFromClient(config)

	{
		resp, err := iwebgw.Regist(*TenantID, *callID, *ani)
		if err != nil {
			panic(err)
		}
		fmt.Printf("TOKEN=%s\n", resp.GetToken())
	}

	ReadLine()

	{
		resp, err := iwebgw.RequestPage(*TenantID, *callID, "111", "/html/main.html", "phone1=01073220804", false)
		if err != nil {
			panic(err)
		}
		fmt.Printf("RESULT=%+v\n", resp.GetResult())
	}

	ReadLine()

	{
		resp, err := iwebgw.RequestPage(*TenantID, *callID, "111", "/html/sub1.html", "phone1=0707322", false)
		if err != nil {
			panic(err)
		}
		fmt.Printf("RESULT=%+v\n", resp.GetResult())
	}
}

func ReadLine() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
