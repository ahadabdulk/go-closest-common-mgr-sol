package main

import (
	"container/list"
	"reflect"
	"sort"
	"testing"
)

func TestEmployeeStruct(t *testing.T) {
	t.Parallel()
	e := "Claire"
	ceo := &Employee{Name: e, Reportees: list.New()}

	if ceo.Name != e {
		t.Errorf("Expected: %v, Actual: %v", e, ceo.Name)
	}
}

func TestOrganisationStruct(t *testing.T) {
	t.Parallel()

	organisationName := "Bureaucr.at"
	o := Organisation{Name: organisationName, EmployeeTable: make(map[int]Employee)}

	if o.Name != organisationName {
		t.Errorf("Expected: %v, Got %v", organisationName, o.Name)
	}
}

func TestEmployeeReportee(t *testing.T) {
	t.Parallel()

	claire := Employee{Name: "Claire", Reportees: list.New()}
	alice := Employee{Name: "Alice", Reportees: list.New()}
	bob := Employee{Name: "Bob", Reportees: list.New()}

	claire.Reportees.PushBack(alice)
	claire.Reportees.PushBack(bob)

	if claire.Reportees.Len() != 2 {
		t.Errorf("Expected: 2, Actual %v", claire.Reportees.Len())
	}
}

func TestAddReporteeFunc(t *testing.T) {
	t.Parallel()

	claire := Employee{Name: "Claire", Reportees: list.New()}
	alice := Employee{Name: "Alice", Reportees: list.New()}
	bob := Employee{Name: "Bob", Reportees: list.New()}

	claire.addReportee(&alice)
	claire.addReportee(&bob)

	if claire.Reportees.Len() != 2 {
		t.Errorf("Expected: 2, Actual %v", claire.Reportees.Len())
	}

	if bob.Manager.Name != claire.Name {
		t.Errorf("Expected: %v, Got: %v", claire.Name, bob.Manager.Name)
	}

	if alice.Manager.Name != claire.Name {
		t.Errorf("Expected: %v, Got: %v", claire.Name, alice.Manager.Name)
	}

}

func TestAddReporteeError(t *testing.T) {
	t.Parallel()

	claire := Employee{Name: "Claire", Reportees: list.New()}
	alice := Employee{Name: "Alice", Reportees: list.New()}

	if ok := claire.addReportee(&alice); ok != nil {
		t.Errorf("Expected: nil, Got: %v", ok)
	}

	bob := Employee{Name: "Bob", Reportees: list.New()}

	if ok2 := bob.addReportee(&alice); ok2.Error() != "Manager already exists" {
		t.Errorf("Expected: nil, Got: %v", ok2)
	}
}

func TestManagerExistsForEmp(t *testing.T) {
	t.Parallel()

	claire := Employee{Name: "Claire", Reportees: list.New()}
	alice := Employee{Name: "Alice", Reportees: list.New()}

	claire.addReportee(&alice)

	bob := Employee{Name: "Bob", Reportees: list.New()}

	if ok := bob.addReportee(&alice); ok.Error() != "Manager already exists" {
		t.Errorf("Expected: nil, Got: %v", ok)
	}
}

func TestAddEmployeeFunc(t *testing.T) {
	t.Parallel()

	organisationName := "Bureaucr.at"

	claire := Employee{Name: "Claire", Reportees: list.New()}
	o := Organisation{Name: organisationName, EmployeeTable: make(map[int]Employee)}

	o.addEmployee(&claire) //function under test

	if len(o.EmployeeTable) != 1 {
		t.Errorf("Expected: 1, Got %v", len(o.EmployeeTable))
	}

	for _, val := range o.EmployeeTable {
		if val.Name != claire.Name {
			t.Errorf("Expected: %v, Got %v", claire.Name, val.Name)
		}
	}
}

func TestEmployeeIdGeneration(t *testing.T) {
	t.Parallel()

	var expected []int
	var observed []int

	claire := Employee{Name: "Claire", Reportees: list.New()}
	alice := Employee{Name: "Alice", Reportees: list.New()}
	bob := Employee{Name: "Bob", Reportees: list.New()}

	o := Organisation{Name: "Bureaucr.at", EmployeeTable: make(map[int]Employee)}

	claire.addReportee(&alice)
	claire.addReportee(&bob)

	o.addEmployee(&claire)
	o.addEmployee(&alice)
	o.addEmployee(&bob)

	expected = append(expected, claire.Id)
	expected = append(expected, alice.Id)
	expected = append(expected, bob.Id)

	for i := range o.EmployeeTable {
		observed = append(observed, i)
	}

	sort.Ints(observed) // Sorting so that order of []expected and []observed matches
	sort.Ints(expected)

	if !(reflect.DeepEqual(observed, expected)) {
		t.Errorf("Expected: %v , Got %v", expected, observed)
	}

}

func TestIsUnderFunc(t *testing.T) {
	t.Parallel()

	claire := Employee{Name: "Claire", Reportees: list.New()}
	alice := Employee{Name: "Alice", Reportees: list.New()}
	bob := Employee{Name: "Bob", Reportees: list.New()}

	claire.addReportee(&alice)
	claire.addReportee(&bob)

	res := bob.isUnder(&claire)

	if !res {
		t.Errorf("Expected: true, Actual %v", res)
	}
}

func TestClosestCommonManagerFunc(t *testing.T) {
	t.Parallel()

	organisationName := "Bureaucr.at"
	o := Organisation{Name: organisationName, EmployeeTable: make(map[int]Employee)}

	jonh := Employee{Name: "Jonh", Reportees: list.New()}
	steve := Employee{Name: "Steve", Reportees: list.New()}
	rima := Employee{Name: "Rima", Reportees: list.New()}

	julia := Employee{Name: "Julia", Reportees: list.New()}
	ahmet := Employee{Name: "Ahmet", Reportees: list.New()}
	ajit := Employee{Name: "Ajit", Reportees: list.New()}

	steve.addReportee(&ajit)
	steve.addReportee(&julia)
	steve.addReportee(&ahmet)

	rico := Employee{Name: "Rico", Reportees: list.New()}
	harry := Employee{Name: "Harry", Reportees: list.New()}

	julia.addReportee(&rico)
	julia.addReportee(&harry)

	anja := Employee{Name: "Anja", Reportees: list.New()}
	rima.addReportee(&anja)

	claire := Employee{Name: "Claire", Reportees: list.New()}

	claire.addReportee(&steve)
	claire.addReportee(&jonh)
	claire.addReportee(&rima)

	o.addEmployee(&claire)
	o.addEmployee(&jonh)
	o.addEmployee(&steve)
	o.addEmployee(&rima)
	o.addEmployee(&julia)
	o.addEmployee(&ahmet)
	o.addEmployee(&ajit)
	o.addEmployee(&rico)
	o.addEmployee(&harry)
	o.addEmployee(&anja)

	res := o.closestCommonManager(&claire, &harry, &jonh)

	if res.Name != claire.Name {
		t.Errorf("Expected %v, Got %v", claire.Name, res.Name)
	}

	res = o.closestCommonManager(&claire, &harry, &ahmet)

	if res.Name != steve.Name {
		t.Errorf("Expected %v, Got %v", steve.Name, res.Name)
	}

	res = o.closestCommonManager(&claire, &anja, &ajit)

	if res.Id != claire.Id {
		t.Errorf("Expected %v, Got %v", claire.Name, res.Name)
	}
}

func TestDetachedCeo(t *testing.T) {
	t.Parallel()

	organisationName := "Bureaucr.at"
	o := Organisation{Name: organisationName, EmployeeTable: make(map[int]Employee)}

	steve := Employee{Name: "Steve", Reportees: list.New()}

	rima := Employee{Name: "Rima", Reportees: list.New()}

	claire := Employee{Name: "Claire", Reportees: list.New()}

	julia := Employee{Name: "Julia", Reportees: list.New()}
	ahmet := Employee{Name: "Ahmet", Reportees: list.New()}
	ajit := Employee{Name: "Ajit", Reportees: list.New()}

	steve.addReportee(&ajit)
	steve.addReportee(&julia)
	steve.addReportee(&ahmet)

	rico := Employee{Name: "Rico", Reportees: list.New()}
	harry := Employee{Name: "Harry", Reportees: list.New()}

	julia.addReportee(&rico)
	julia.addReportee(&harry)

	anja := Employee{Name: "Anja", Reportees: list.New()}

	rima.addReportee(&anja)

	o.addEmployee(&steve)
	o.addEmployee(&rima)
	o.addEmployee(&julia)
	o.addEmployee(&ahmet)
	o.addEmployee(&ajit)
	o.addEmployee(&rico)
	o.addEmployee(&harry)
	o.addEmployee(&anja)

	res := o.closestCommonManager(&claire, &anja, &harry)

	if res != nil {
		t.Errorf("Expected nil, Got %v", res)
	}

}

func TestNonExistingEmployee(t *testing.T) {

	t.Parallel()
	organisationName := "Bureaucr.at"
	o := Organisation{Name: organisationName, EmployeeTable: make(map[int]Employee)}

	jonh := Employee{Name: "Jonh", Reportees: list.New()}
	steve := Employee{Name: "Steve", Reportees: list.New()}
	rima := Employee{Name: "Rima", Reportees: list.New()}

	julia := Employee{Name: "Julia", Reportees: list.New()}
	ahmet := Employee{Name: "Ahmet", Reportees: list.New()}
	ajit := Employee{Name: "Ajit", Reportees: list.New()}

	steve.addReportee(&ajit)
	steve.addReportee(&julia)
	steve.addReportee(&ahmet)

	rico := Employee{Name: "Rico", Reportees: list.New()}
	harry := Employee{Name: "Harry", Reportees: list.New()}

	julia.addReportee(&rico)
	julia.addReportee(&harry)

	anja := Employee{Name: "Anja", Reportees: list.New()}
	rima.addReportee(&anja)

	claire := Employee{Name: "Claire", Reportees: list.New()}

	claire.addReportee(&steve)
	claire.addReportee(&jonh)
	claire.addReportee(&rima)

	o.addEmployee(&claire)
	o.addEmployee(&jonh)
	o.addEmployee(&steve)
	o.addEmployee(&rima)
	o.addEmployee(&julia)
	o.addEmployee(&ahmet)
	o.addEmployee(&ajit)
	o.addEmployee(&rico)
	o.addEmployee(&harry)
	o.addEmployee(&anja)

	//Creating empty Employee
	noem01 := Employee{}

	res := o.closestCommonManager(&claire, &noem01, &jonh)

	if res != nil {
		t.Errorf("Excpected: nil, Got: %v", res)
	}

	// res2 := o.closestCommonManager(&noem01, &noem01, &jonh)

	// if res2 != nil {
	// 	t.Errorf("Excpected: nil, Got: %v", res2)
	// }

}

func TestNilEmployees(t *testing.T) {
	t.Parallel()

	organisationName := "Bureaucr.at"
	o := Organisation{Name: organisationName, EmployeeTable: make(map[int]Employee)}
	res := o.closestCommonManager(nil, nil, nil)

	if res != nil {
		t.Errorf("Expected: nil, Got:%v", res)
	}
}

func TestOnlyEmployeeCeo(t *testing.T) {
	t.Parallel()

	organisationName := "Bureaucr.at"
	o := Organisation{Name: organisationName, EmployeeTable: make(map[int]Employee)}

	claire := Employee{Name: "Claire", Reportees: list.New()}

	res := o.closestCommonManager(&claire, nil, nil)

	if res != nil {
		t.Errorf("Expected: nil, Got:%v", res)
	}
}
