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

type Heartbeat struct {
	Time      string `json:"time"`
	Heartbeat int    `json:"heartbeat"`
}

func CreateResponseHeartbeat(heartbeat models.Heartbeat) Heartbeat {
	return Heartbeat{Time: get_time(heartbeat.CreatedAt), Heartbeat: heartbeat.Value}
}

func CreateHeartbeat(c *fiber.Ctx) error {
	var heartbeat models.Heartbeat
	if err := c.BodyParser(&heartbeat); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	res := database.Database.Db.Create(&heartbeat)

	if res.Error != nil {
		return c.Status(500).JSON(res.Error.Error())

	} else {
		responseHeartbeat := CreateResponseHeartbeat(heartbeat)
		return c.Status(200).JSON(responseHeartbeat)
	}
}

func GetHearbeats(clientId int) []models.Heartbeat {
	heartbeats := []models.Heartbeat{}
	database.Database.Db.Find(&heartbeats, "user_id=?", clientId)
	return heartbeats
}
func get_time(currentTime time.Time) string {

	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second())
}

var (
	heartbeatclients   = make(map[*websocket.Conn]bool)
	heartbeatclientsMu sync.Mutex
)

func Heartbeat_socket(c *websocket.Conn) {

	heartbeatclientsMu.Lock()
	heartbeatclients[c] = true
	heartbeatclientsMu.Unlock()

	// Send the current messages array to the client
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}
	for _, msg := range GetHearbeats(id) {
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
		// Convert []byte to Heartbeat struct
		var heartbeat models.Heartbeat
		heartbeat.CreatedAt = time.Now()
		err = json.Unmarshal(msg, &heartbeat)
		// Add to database
		res := database.Database.Db.Create(&heartbeat)
		if res.Error != nil {
			log.Println(res.Error.Error())

		}
		fmt.Println(heartbeat)
		msg_to_send, err := json.Marshal(heartbeat)
		log.Println(err)

		fmt.Fprintln(os.Stdout, msg, msg_to_send)
		log.Println(msg_to_send)
		// Broadcast the new message to all clients
		heartbeatclientsMu.Lock()
		for client := range heartbeatclients {
			// Check if the client has the same id
			clientID, err := strconv.Atoi(client.Params("id"))
			if err != nil {
				log.Println(err)
				continue
			}
			if clientID == id {
				if err := client.WriteMessage(mt, msg_to_send); err != nil {
					log.Println("write:", err)
					delete(heartbeatclients, client)
					break
				}
			}
		}
		heartbeatclientsMu.Unlock()
	}

	heartbeatclientsMu.Lock()
	delete(heartbeatclients, c)
	heartbeatclientsMu.Unlock()
}
