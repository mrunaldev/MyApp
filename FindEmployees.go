package main

import (
	"fmt"
	"sort"
)

// A —> [B, D, C, E, F]
// B —> [D, C, E, F]
// C —> []
// D —> [E, F]
// E —> [F]
// F —> []
// tests
var employeeToManagerMappings map[string]string

func FindAllReportingEmployees(manager string) string {
	reportees := ""
	if manager == "" {
		reportees = ""
	} else {
		for emp, mgr := range employeeToManagerMappings {
			if manager == mgr {
				if emp != mgr {
					reportees = reportees + emp + FindAllReportingEmployees(emp)
				}
			}
		}
	}
	return reportees
}

func FindEmployees() {
	employeeToManagerMappings = map[string]string{"A": "A", "B": "A", "C": "B", "D": "B", "E": "D", "F": "E"}

	var managers []string
	for employee, manager := range employeeToManagerMappings {

		managers = append(managers, manager)
		managers = append(managers, employee)
	}
	sort.Strings(managers)
	for _, val := range unique(managers) {
		fmt.Println(val, "->", FindAllReportingEmployees(val))

	}
}
func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
