package routes

import (
	"encoding/json"
	"fst/project/database"
	"fst/project/models"
	"log"
	"os"
	"sync"

	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Step struct {
	Time string `json:"time"`
	Step int    `json:"step"`
}

var (
	stepsclients   = make(map[*websocket.Conn]bool)
	stepsclientsMu sync.Mutex
)

func CreateResponseStep(step models.Step) Step {
	return Step{Time: get_time(step.CreatedAt), Step: step.Value}
}

func CreateStep(c *fiber.Ctx) error {
	var step models.Step
	if err := c.BodyParser(&step); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	res := database.Database.Db.Create(&step)

	if res.Error != nil {
		return c.Status(500).JSON(res.Error.Error())

	} else {
		responseStep := CreateResponseStep(step)
		return c.Status(200).JSON(responseStep)
	}
}

func GetSteps(clientId int) []models.Step {
	steps := []models.Step{}
	database.Database.Db.Find(&steps, "user_id=?", clientId)
	return steps
}

func Step_socket(c *websocket.Conn) {

	stepsclientsMu.Lock()
	stepsclients[c] = true
	stepsclientsMu.Unlock()

	// Send the current messages array to the client
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}
	for _, msg := range GetSteps(id) {
		jsonmessage, err := json.Marshal(msg)
		log.Println(err)
		if err := c.WriteMessage(websocket.TextMessage, jsonmessage); err != nil {
			log.Println("write:", err)
			break
		}
	}

	for {
		// Read message from the client
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// Convert []byte to Step struct
		var step models.Step
		step.CreatedAt = time.Now()
		err = json.Unmarshal(msg, &step)
		// Add to database
		res := database.Database.Db.Create(&step)
		if res.Error != nil {
			log.Println(res.Error.Error())

		}
		fmt.Println(step)
		msg_to_send, err := json.Marshal(step)
		log.Println(err)

		fmt.Fprintln(os.Stdout, msg, msg_to_send)
		log.Println(msg_to_send)
		// Broadcast the new message to all clients
		stepsclientsMu.Lock()
		for client := range stepsclients {
			// Check if the client has the same id
			clientID, err := strconv.Atoi(client.Params("id"))
			if err != nil {
				log.Println(err)
				continue
			}
			if clientID == id {
				if err := client.WriteMessage(mt, msg_to_send); err != nil {
					log.Println("write:", err)
					delete(stepsclients, client)
					break
				}
			}
		}
		stepsclientsMu.Unlock()
	}

	stepsclientsMu.Lock()
	delete(stepsclients, c)
	stepsclientsMu.Unlock()
}
