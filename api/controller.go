package api

import (
	"net/http"
	"strconv"

	"github.com/tebrizetayi/optiopay/internal/corporate"
)

const (
	MIMEApplicationJSON = "application/json"
)

type Controller struct {
	DirectoryService corparate.Directory
}

func NewController(directoryService corparate.Directory) Controller {
	return Controller{DirectoryService: directoryService}
}

type Employee struct {
	Id   int
	Name string
}
type Response struct {
	Error    string
	Employee Employee
}

func (c *Controller) FindCommonManagers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id1 := r.URL.Query().Get("id1")
	if id1 == "" {
		http.Error(w, "Missing employee ID", http.StatusBadRequest)
		return
	}

	id2 := r.URL.Query().Get("id2")
	if id2 == "" {
		http.Error(w, "Missing employee ID", http.StatusBadRequest)
		return
	}

	empId1, err := strconv.Atoi(id1)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	empId2, err := strconv.Atoi(id2)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	emp, err := c.DirectoryService.FindClosestCommonManager(empId1, empId2)

	if handleIfError(w, err, http.StatusInternalServerError) {
		return
	}

	respondJson(ctx, w, Employee{emp.Id, emp.Name}, http.StatusOK)
}
