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

func GetEmployeeById(c *fiber.Ctx) error {
	employees := []Employee{
		{
			EmployeeID: "E001",
			Name:       "Alice Johnson",
			Phone:      "555-1234",
			Address:    "123 Elm St, Springfield",
		},
		{
			EmployeeID: "E002",
			Name:       "Bob Smith",
			Phone:      "555-5678",
			Address:    "456 Oak St, Springfield",
		},
		{
			EmployeeID: "E003",
			Name:       "Charlie Brown",
			Phone:      "555-8765",
			Address:    "789 Pine St, Springfield",
		},
		{
			EmployeeID: "E004",
			Name:       "Diana Prince",
			Phone:      "555-4321",
			Address:    "321 Maple St, Springfield",
		},
	}

	empId := c.Query("empId")
	if empId != "" {
		for _, employee := range employees {
			if employee.EmployeeID == empId {
				return c.JSON(employee)
			}
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
