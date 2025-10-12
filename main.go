package main

import (
	"fmt"
)

func main() {
	myBill := newBill("jaime's bill")
	myBill.updateTip(10)
	myBill.addItem("onion soup", 4.50)
	myBill.addItem("butt soup", 2.55)
	myBill.addItem("caca soup", 1.00)
	myBill.addItem("chicken soup", 5.67)

	fmt.Println(myBill.format())
}
