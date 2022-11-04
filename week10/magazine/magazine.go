package magazine

type Subscriber struct { // 대문자로 선언해 주어야 함. 왜냐면 다른 파일에서 사용할 거니까.. 소문자로 선언 : private, 대문자로 선언 : public이라고 생각하자.
	Name   string
	Rate   float64
	Active bool
	// HomeAddress Address // 이렇게 하면 구독자(Susbcriber)든 직원(Employee)이든 다 주소(자기집)를 가질 수 있게 됨.
	Address // 줄여서 쓰기
}

type Employee struct {
	Name   string
	Salary float64
	// HomeAddress Address
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
