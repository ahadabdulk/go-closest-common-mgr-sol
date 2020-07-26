package main

import (
	"container/list"
	"fmt"
)

func main() {

	organisationName := "Bureaucr.at"
	o := Organisation{Name: organisationName, EmployeeTable: make(map[int]Employee)}

	//Creating Employees
	jonh := Employee{Name: "Jonh", Reportees: list.New()}
	steve := Employee{Name: "Steve", Reportees: list.New()}
	rima := Employee{Name: "Rima", Reportees: list.New()}

	julia := Employee{Name: "Julia", Reportees: list.New()}
	ahmet := Employee{Name: "Ahmet", Reportees: list.New()}
	ajit := Employee{Name: "Ajit", Reportees: list.New()}

	//Add employees under Steve, who is a manager
	steve.addReportee(&ajit)
	steve.addReportee(&julia)
	steve.addReportee(&ahmet)

	rico := Employee{Name: "Rico", Reportees: list.New()}
	harry := Employee{Name: "Harry", Reportees: list.New()}

	//Add employees under Julia, who is a manager
	julia.addReportee(&rico)
	julia.addReportee(&harry)

	anja := Employee{Name: "Anja", Reportees: list.New()}

	//Add Anja under Rima
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

	for _, val := range o.EmployeeTable {
		fmt.Println()
		fmt.Printf("%v [Manager] <--[Reportee:]    ", val.Name)
		for em := val.Reportees.Front(); em != nil; em = em.Next() {
			fmt.Printf("%v,", em.Value.(Employee).Name)
		}
	}

	res := o.closestCommonManager(&claire, &harry, &jonh)

	res2 := o.closestCommonManager(&claire, &harry, &ahmet)

	res3 := o.closestCommonManager(&claire, &anja, &ajit)

	fmt.Println("\n\nClosest Common Manager for Harry and John is : ", res.Name)
	fmt.Println("\n\nClosest Common Manager for Harry and Ahmet is : ", res2.Name)
	fmt.Println("\n\nClosest Common Manager for Anja and Ajit is : ", res3.Name)

}
