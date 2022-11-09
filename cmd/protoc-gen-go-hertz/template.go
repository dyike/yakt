package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

//go:embed template.go.tpl
var tpl string

type serviceDesc struct {
	Name      string // Greeter
	FullName  string // helloworld.Greeter
	FilePath  string // api/helloword/helloworld.proto
	Methods   []*methodDesc
	MethodSet map[string]*methodDesc
}

func (s *serviceDesc) execute() string {
	s.MethodSet = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSet[m.Name] = m
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

// InterfaceName service interface name
func (s *serviceDesc) InterfaceName() string {
	return s.Name + "HTTPServer"
}

type methodDesc struct {
	// method
	Name    string // SayHello
	Num     int    // 一个rpc方法可以对应多个http请求
	Request string // SayHelloReq
	Reply   string // SayHelloResp
	// http_rule
	Path         string // 路由
	Method       string // HTTP Method
	Body         string
	ResponseBody string
}

// HandlerName for hertz handler name
func (m *methodDesc) HandlerName() string {
	return fmt.Sprintf("%s_%d", m.Name, m.Num)
}
