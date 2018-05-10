package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	s := `One of the hot topics at Google I/O 2018 has been the evolution of the Material Design.
		The website has been overhauled and lots of new guidelines have been added, with some of 
		the changes being discussed in sessions at the conference. Material Theming is also being
		introduced to help developers implement the design language in their apps while retaining
		their own style. Another tool, Material Gallery, has just been launched to help designers
		organize their projects.`

	encodStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890abcdefghuijklmopqrtuvwxyz+/"

	s64 := base64.NewEncoding(encodStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))

	fmt.Println(s)
	fmt.Println(s64)
}
