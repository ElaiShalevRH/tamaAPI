package main

import "github.com/ElaiShalevRH/tamaAPI/cmd"

func main() {

	if err := cmd.Execute(); err != nil {
		panic(err)
	}

}
