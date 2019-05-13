package main

import (
	"flag"
	"fmt"
	"os"
)

func flagFunc() {
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please input config path")
	flag.IntVar(&logLevel, "d", 0, "please input log level")
	// 启用参数配置
	flag.Parse()

	fmt.Println("conPath: ", confPath)
	fmt.Println("log level: ", logLevel)
}

func main() {
	fmt.Printf("Args len = %d\n", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("Args[%d] == %s\n", i, v)
	}

	flagFunc()
}