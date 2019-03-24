package Socket

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var LinkTable = map[string]map[int]*websocket.Conn{}
var Pools = Pool{}

func LinkSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	// ===================================================
	var idName string
	if getID, ok := r.URL.Query()["ID"]; !ok {
		w.Write([]byte("ERROR"))
		return
	} else {
		idName = getID[0]
	}
	// ===================================================
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte("ERROR"))
		return
	}
	// ===================================================
	var id int
	if _, ok := LinkTable[idName]; !ok || len(LinkTable[idName]) == 0 {
		LinkTable[idName] = map[int]*websocket.Conn{
			1: connection,
		}
		id = 1
	} else {
		count := 0
		for {
			count++
			_, ok := LinkTable[idName][count]
			if !ok {
				LinkTable[idName][count] = connection
				id = count
				break
			}
		}
	}
	go Pools.Schedule(idName, id)
}

// ========================================================

func NumberSocket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if getID, ok := r.URL.Query()["ID"]; !ok {
		w.Write([]byte("ERROR"))
		return
	} else if getID[0] == "__all__" {
		allmap := map[string]int{}
		for key, val := range LinkTable {
			allmap[key] = len(val)
		}
		json.NewEncoder(w).Encode(allmap)
	} else if _, ok := LinkTable[getID[0]]; !ok {
		json.NewEncoder(w).Encode(
			map[string]int{
				getID[0]: 0,
			})
		return
	} else {
		json.NewEncoder(w).Encode(
			map[string]int{
				getID[0]: len(LinkTable[getID[0]]),
			})
		return
	}
}

// ========================================================

type Pool struct {
	Works map[string]Message
}

func (p *Pool) Schedule(name string, id int) {
	if p.Works == nil {
		p.Works = make(map[string]Message)
	}
	if _, ok := p.Works[name]; !ok || len(LinkTable[name]) == 1 {
		p.Works[name] = Message{
			name:        name,
			linkMessage: make(chan []byte, 1),
			linkType:    make(chan int, 1),
		}
		p.Works[name].MakeListen(id)
	} else {
		go p.Works[name].GetMessage(id)
	}
}

// ========================================================

type Message struct {
	linkMessage chan []byte
	linkType    chan int
	name        string
}

func (core Message) MakeListen(id int) {
	go core.GetMessage(id)
	go core.RequestMessage()
}

func (core Message) GetMessage(key int) {
	childConnection := LinkTable[core.name][key]
	for {
		messageType, message, err := childConnection.ReadMessage()
		if err != nil {
			delete(LinkTable[core.name], key)
			return
		} else {
			core.linkType <- messageType
			core.linkMessage <- message
		}
	}
}

func (core *Message) RequestMessage() {
	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case requsetMessage := <-core.linkMessage:
			requsetType := <-core.linkType
			for key, connection := range LinkTable[core.name] {
				err := connection.WriteMessage(requsetType, requsetMessage)
				if err != nil {
					delete(LinkTable[core.name], key)
				}
			}
		case <-tick.C:
			if len(LinkTable[core.name]) == 0 {
				return
			}
		}
	}
}
