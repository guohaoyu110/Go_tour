package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
	//contact contactInfo
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // use a type to describe another customized type
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"} //person is the type; use firstname and lastname as strings
	// 也可以不写firstname lastname不过那个就按照原来定义的格式来
	fmt.Println("alex: ", alex)

	// another way of declare a struct

	var alex1 person
	alex1.firstName = "Ben"
	alex1.lastName = "Guo"
	fmt.Println("alex1: ", alex1)
	fmt.Printf("%+v", alex1) // print the firstName and lastName

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, //每一行最后都要有逗号
		}, //这个最后也要加上逗号
	}

	fmt.Println("jim: ", jim)
	fmt.Printf("%+v: ", jim)

	fmt.Println("///////////////: ")
	jimPointer := &jim //give me the memory address of the value this variable is pointing at
	jimPointer.updateName("jimmy")
	jimPointer.print() //update does not take effect, so i need to take into consideration the pointers
	jim.print()
	//dive deep and understand why it does not take effect
	// go的指针问题！！

	//fmt.Println("jim: ", jim)
}

// 用指针的方法传递，普通值传递不可以
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println("%+v: ", p)
}

//如何实现指针呢？
// when you pass an integer, float, string, or struct into a function, it creats a copy of each argument, and these copies are used inside of the function

// a slice is a reference type, because a slice contains a reference to the actual underlying list of records
