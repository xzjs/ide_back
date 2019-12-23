package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/astaxie/beego"
)

//MainController is a struct
type MainController struct {
	beego.Controller
}

type Code struct {
	Code string
}

//Get is function
func (c *MainController) Get() {
	c.TplName = "index.html"
}

//Post 方法
func (c *MainController) Post() {
	var code Code
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &code)
	if err != nil {
		panic(err)
	}

	t, err := ioutil.ReadFile("template.py")
	if err != nil {
		panic(err)
	}

	handle, err := os.Create("code.py")
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	handle.WriteString(string(t) + "\n\n" + code.Code + "\n\n" + "print(arr)")

	cmd := exec.Command("python3", "code.py")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println(out.String())
	}

	c.Ctx.WriteString(out.String())
}
