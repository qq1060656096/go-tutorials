package main

import "fmt"

type PinterReceiver struct {
	name string
	age uint8
}
// 注意接收者是指针，接收变量是副本指针（相当于定义了一个新的变量，指向了老的指针），本质是值传递
func (p *PinterReceiver) setName(name string)  {
	// p 相当于 &(*p)
	p.name = name
	fmt.Printf("PinterReceiver.setName: %p, %#v\n", p, name)
	fmt.Printf("PinterReceiver.setName: %p, %#v\n", &p, name)
	fmt.Printf("PinterReceiver.setName: %p, %#v\n", &(*p), name)
	fmt.Println()
}


// 值接收器
type ValueReceiver struct {
	name string
	age uint8
}

func (v ValueReceiver) setName(name string)  {
	v.name = name
	fmt.Printf("ValueReceiver.setName: %p, %#v\n", v, name)
	fmt.Printf("ValueReceiver.setName: %p, %#v\n", &v, name)
	fmt.Println()
}

func main()  {
	pr := &PinterReceiver{}

	fmt.Printf("PinterReceiver.setName:call:before %p, %#v\n", pr, pr.name)
	pr.setName("PinterReceiver.setName:call:1")
	fmt.Printf("PinterReceiver.setName:call:1:after %p, %#v\n", pr, pr.name)
	pr.setName("PinterReceiver.setName:call:2")
	pr.setName("PinterReceiver.setName:call:3")

	fmt.Println()
	fmt.Println("demo2: ")
	vr := &ValueReceiver{}
	fmt.Printf("ValueReceiver.setName:call:before %p, %#v\n", vr, vr.name)
	fmt.Printf("ValueReceiver.setName:call:before %p, %#v\n", &vr, vr.name)
	// 因为是值传递，所以修改值在外部无效
	vr.setName("ValueReceiver.setName:call:1")
	fmt.Printf("ValueReceiver.setName:call:1:after %p, %#v\n", vr, vr.name)

	vr.setName("ValueReceiver.setName:call:2")
	vr.setName("ValueReceiver.setName:call:3")
}
