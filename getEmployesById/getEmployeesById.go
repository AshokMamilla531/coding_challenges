package main

import (
	"github.com/gofiber/fiber/v2"
)

type Employee struct {
	EmployeeID string `json:"employeeId"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

var employeeMap = map[string]Employee{
	"E001": {EmployeeID: "E001", Name: "Alice Johnson", Phone: "555-1234", Address: "123 Elm St, Springfield"},
	"E002": {EmployeeID: "E002", Name: "Bob Smith", Phone: "555-5678", Address: "456 Oak St, Springfield"},
	"E003": {EmployeeID: "E003", Name: "Charlie Brown", Phone: "555-8765", Address: "789 Pine St, Springfield"},
	"E004": {EmployeeID: "E004", Name: "Diana Prince", Phone: "555-4321", Address: "321 Maple St, Springfield"},
}

func GetEmployeeById(c *fiber.Ctx) error {
	empId := c.Query("empId")
	if empId != "" {
		if employee, exists := employeeMap[empId]; exists {
			return c.JSON(employee)
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Employee not found",
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Employee ID is required",
	})
}

func main() {
	app := fiber.New()

	app.Get("/employee", GetEmployeeById)

	app.Listen(":3000")
}
