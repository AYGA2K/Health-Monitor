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

type Sleep struct {
	Time  string `json:"time"`
	Sleep int    `json:"sleep"`
}

var (
	sleepclients   = make(map[*websocket.Conn]bool)
	sleepclientsMu sync.Mutex
)

func CreateResponseSleep(sleep models.Sleep) Sleep {
	return Sleep{Time: get_time(sleep.CreatedAt), Sleep: sleep.Value}
}

func CreateSleep(c *fiber.Ctx) error {
	var sleep models.Sleep
	if err := c.BodyParser(&sleep); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	res := database.Database.Db.Create(&sleep)

	if res.Error != nil {
		return c.Status(500).JSON(res.Error.Error())

	} else {
		responseSleep := CreateResponseSleep(sleep)
		return c.Status(200).JSON(responseSleep)
	}
}

func GetSleeps(clientId int) []models.Sleep {
	sleeps := []models.Sleep{}
	database.Database.Db.Find(&sleeps, "user_id=?", clientId)
	return sleeps
}

func Sleep_socket(c *websocket.Conn) {

	sleepclientsMu.Lock()
	sleepclients[c] = true
	sleepclientsMu.Unlock()

	// Send the current messages array to the client
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}
	for _, msg := range GetSleeps(id) {
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
		// Convert []byte to Sleep struct
		var sleep models.Sleep
		sleep.CreatedAt = time.Now()
		err = json.Unmarshal(msg, &sleep)
		// Add to database
		res := database.Database.Db.Create(&sleep)
		if res.Error != nil {
			log.Println(res.Error.Error())

		}
		fmt.Println(sleep)
		msg_to_send, err := json.Marshal(sleep)
		log.Println(err)

		fmt.Fprintln(os.Stdout, msg, msg_to_send)
		log.Println(msg_to_send)
		// Broadcast the new message to all clients
		sleepclientsMu.Lock()
		for client := range sleepclients {
			// Check if the client has the same id
			clientID, err := strconv.Atoi(client.Params("id"))
			if err != nil {
				log.Println(err)
				continue
			}
			if clientID == id {
				if err := client.WriteMessage(mt, msg_to_send); err != nil {
					log.Println("write:", err)
					delete(sleepclients, client)
					break
				}
			}
		}
		sleepclientsMu.Unlock()
	}

	sleepclientsMu.Lock()
	delete(sleepclients, c)
	sleepclientsMu.Unlock()
}
