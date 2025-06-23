package reflect

import "testing"

func TestName(t *testing.T) {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
}

func TestName2(t *testing.T) {

	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery2(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery2(e)
}
