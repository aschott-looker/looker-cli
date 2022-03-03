package api

import (
	"fmt"
	"os"

	"github.com/looker-open-source/sdk-codegen/go/rtl"
	v4 "github.com/looker-open-source/sdk-codegen/go/sdk/v4"
)

func check(err error) {
	if err != nil {
		panic(any(err))
	}
}

func InitSdk() *v4.LookerSDK {
	lookerIniPath := "looker.ini"

	// Read settings from ini file
	fmt.Println("Looking for ini file at", lookerIniPath)
	pwd, _ := os.Getwd()
	fmt.Println("Currently at", pwd)
	cfg, err := rtl.NewSettingsFromFile(lookerIniPath, nil)
	check(err)

	// New instance of LookerSDK
	return v4.NewLookerSDK(rtl.NewAuthSession(cfg))
}