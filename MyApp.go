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
var employeeToManagerMappings map[string]string

func main() {

	// construct a mapping from employee to manager
	employeeToManagerMappings = map[string]string{"A": "A", "B": "A", "C": "B", "D": "B", "E": "D", "F": "E"}
	findEmployees()

	// // print contents of the resulting dictionary
	// for key, value := range result
	//     print(key, "—>", value)
}

func findAllReportingEmployees(manager string) string {
	reportees := ""

	if manager == "" {
		reportees = ""
	} else {
		for emp, mgr := range employeeToManagerMappings {
			if manager == mgr {
				if emp != mgr {
					reportees = reportees + findAllReportingEmployees(emp)
				} else {
					continue
				}

			}
		}

	}

	// for _, val := range reportees {
	// 	reportees = append(reportees, findAllReportingEmployees(val))
	// }
	return reportees
}

func findEmployees() {
	var managers []string
	for employee, manager := range employeeToManagerMappings {

		managers = append(managers, manager)
		managers = append(managers, employee)
	}
	sort.Strings(managers)
	for _, val := range unique(managers) {
		fmt.Println(val, "->", findAllReportingEmployees(val))

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
