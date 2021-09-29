package ws

var gHub *Hub

type ClientType string

const (
	ClientTypeKLineHour           ClientType = "60"
	ClientTypeKLineFourHour       ClientType = "240"
	ClientTypeKLineDay            ClientType = "1D"
	ClientTypeKLineWeek           ClientType = "1W"
	ClientTypeKLineMouth          ClientType = "1M"
	ClientTypeKLineMinute         ClientType = "1"
	ClientTypeKLineFiveMinute     ClientType = "5"
	ClientTypeKLineFifteenMinute  ClientType = "15"
	ClientTypeKLineHalfHourMinute ClientType = "30"
	ClientTypeKLineLastPrice      ClientType = "KLineLastPrice"
	// 左侧栏LastPrice监听事件
	ClientTypeLeftLastPrice ClientType = "LeftLastPrice"
)

type BroadcastData struct {
	Data []byte
	Type ClientType
	// ClientTypeKLineHour 与 ClientTypeKLineDay才有效
	PairAddress string
}

type Hub struct {
	// Registered clients.
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client

	Broadcast chan *BroadcastData
}

func InitWsHub() *Hub {
	gHub = &Hub{
		Broadcast:  make(chan *BroadcastData),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
	return gHub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case bd := <-h.Broadcast:
			// 如果广播的类型是ClientTypeKLineLastPrice，
			// 则需要将消息发送给所有的 ClientTypeKLineHour和ClientTypeKLineDay 客户端
			if bd.Type == ClientTypeKLineLastPrice {
				for client := range h.clients {
					if client.pairAddress == bd.PairAddress &&
						(client.ct != ClientTypeLeftLastPrice) {
						select {
						case client.send <- bd.Data:
						default:
							close(client.send)
							delete(h.clients, client)
						}
					}
				}
				continue
			}

			if bd.Type == ClientTypeLeftLastPrice {
				for client := range h.clients {
					if client.ct == ClientTypeLeftLastPrice {
						select {
						case client.send <- bd.Data:
						default:
							close(client.send)
							delete(h.clients, client)
						}
					}
				}
				continue
			}

			// 发送到指定类型的客户端
			for client := range h.clients {
				if client.ct == bd.Type && client.pairAddress == bd.PairAddress {
					select {
					case client.send <- bd.Data:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

func GetWSHub() *Hub {
	return gHub
}
