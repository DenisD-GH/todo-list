package main

import (
	"log"
	"strconv"
	"todo-list/db"
	"todo-list/endpoints"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	conn, ctx := db.DBConnect()
	defer conn.Close(ctx)

	db.CreateTaskTable(conn, ctx)

	app.Get("/tasks", func(c *fiber.Ctx) error {
		tasks, err := endpoints.GetTasks(conn, ctx)

		if err != nil {
			log.Printf("Error getting tasks: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error getting tasks")
		}
		return c.JSON(tasks)
	})

	app.Post("/tasks", func(c *fiber.Ctx) error {
		var task endpoints.Task
		err := c.BodyParser(&task)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
		}

		err = endpoints.CreateTask(conn, ctx, task)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error creating task")
		}
		return c.Status(fiber.StatusCreated).SendString("Task created")
	})

	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		var task endpoints.Task
		err = c.BodyParser(&task)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
		}
		task.ID = id

		if err := endpoints.UpdateTask(conn, ctx, task); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating task")
		}
		return c.SendString("Task updated")
	})

	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		err = endpoints.DeleteTask(conn, ctx, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error deleting task")
		}
		return c.SendString("Task deleted")
	})

	log.Println("Starting server on :3000")
	log.Fatal(app.Listen(":3000"))
}
