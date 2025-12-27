package main

import "fmt"

var pl = fmt.Println

type customer struct {
	name    string
	address string
	bal     float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

type contract struct {
	fname string
	lname string
	phone string
}

type business struct {
	name    string
	address string
	contract
}

func (b business) getBussinessInfo() {
	fmt.Printf("contact at %s is %s %s", b.name, b.contract.fname, b.contract.lname)
}

func main() {
	var c1 customer
	c1.name = "Tom Smith"
	c1.address = "Kissakatu 1"
	c1.bal = 2343.23
	getCustInfo(c1)
	newCustAdd(&c1, "Emannantie 1")
	pl(c1)

	c2 := customer{
		name:    "Kissa",
		address: "Katon katolla",
	}
	pl(c2)

	c3 := customer{"Koira", "Talossa", 123}
	pl(c3)

	con1 := contract{
		"James",
		"Wang",
		"837827384",
	}
	bus1 := business{
		"ABC putket",
		"Halonkatu 1",
		con1,
	}

	bus1.getBussinessInfo()
}
