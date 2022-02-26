package main

import "fmt"

type NotFoundErrors struct {

}

type MethodNotAllow struct {
	method string
}


type HookTimeOutErr struct {
}

func (h HookTimeOutErr) Error() string  {
	return fmt.Sprint("Hook Time out \n")
}

func NewMethodNotAllow(method string) MethodNotAllow  {
	return MethodNotAllow{method: method}
}

func NewNotFound() NotFoundErrors  {
	return NotFoundErrors{}
}

func (m MethodNotAllow) Error() string  {
	return fmt.Sprint("Method allow: %s",m.method)
}


func (m NotFoundErrors) Error() string  {
	return "Not found"
}