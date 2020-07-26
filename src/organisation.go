package main

import (
	"container/list"
	"sync"
)

// Employee Id will start from 100
var id int = 1000

//Organisation here is Bureaucr.at
type Organisation struct {
	Name          string
	EmployeeTable map[int]Employee
	mux           sync.Mutex
}

//addEmployee adds a new employee in the organisation corresponding to an 'id' in employee table, which is a map
func (o *Organisation) addEmployee(e *Employee) {
	o.mux.Lock()
	defer o.mux.Unlock()
	e.Id = id
	o.EmployeeTable[id] = *e
	id++

}

// func closestCommonManager() finds common manager farthest from CEO, where:
// c - root of organisation (i.e. CEO)
// e1 & e2 are employees whose common manager need to be found
func (o *Organisation) closestCommonManager(c *Employee, e1 *Employee, e2 *Employee) *Employee {
	queue := list.New()
	var closestCommonMgr *Employee = nil

	if c == nil || e1 == nil || e2 == nil {
		return nil
	}

	if !(e1.isUnder(c)) && !(e2.isUnder(c)) {
		return nil
	}
	//queue := list.New()
	queue.PushBack(*c)
	for queue.Len() > 0 {
		e := queue.Front().Value.(Employee)
		queue.Remove(queue.Front())
		if e1.isUnder(&e) && e2.isUnder(&e) {
			closestCommonMgr = &e
			for i := e.Reportees.Front(); i != nil; i = i.Next() {
				queue.PushBack(i.Value)
			}
		}
	}
	return closestCommonMgr
}
