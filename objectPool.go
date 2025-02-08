// thread
// database pool etc'

package main

import "fmt"

type IDBConnection interface{
	getID()string
}
type connection struct{
	id string
}
func (c* connection) getID()string{
	return c.id
}

type poolConnection struct {
	idle []IDBConnection
	active []IDBConnection
	capacity int 
}

func (p *poolConnection) getConnection()IDBConnection{ // understand what is happening heere
	if len(p.idle)==0{
		panic("no connection present in the pool")
	}
	obj:=p.idle[0]
	p.idle=p.idle[1:]
	p.active=append(p.active,obj)
	fmt.Printf("get obj pool with id:%s\n",obj.getID())
	return obj
}

func (p *poolConnection)returnConnection(target IDBConnection)error{
	err:=p.remove(target)
	if err!=
}