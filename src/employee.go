package main

import (
	"container/list"
	"errors"
)

//Employee is a struct
type Employee struct {
	Id        int
	Name      string
	Manager   *Employee
	Reportees *list.List //List of employees
}

// addReportee() is to add reportee 'r' to manager 'e'
func (e *Employee) addReportee(r *Employee) error {

	if r.Manager == nil {
		r.Manager = e
		e.Reportees.PushBack(*r)
		return nil
	} else {
		return errors.New("Manager already exists")
	}

}

//e.isUnders(m *Employee), returns true if Employee 'e' comes under Manager 'm'
func (e *Employee) isUnder(m *Employee) bool {
	if m == nil {
		return false
	}
	if m.Name == e.Name {
		return true
	}
	if m.Reportees.Len() == 0 {
		return false
	}

	cover := false

	for em := m.Reportees.Front(); em != nil; em = em.Next() {
		nm := em.Value.(Employee)
		cover = cover || e.isUnder(&nm)
	}
	return cover
}
