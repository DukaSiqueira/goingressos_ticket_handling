package tickets

type Ticket struct {
	ID        int    `json:"id"`
	EventID   int    `json:"event_id"`
	UserID    int    `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RefundRequest struct {
	ID        int    `json:"id"`
	TicketID  int    `json:"ticket_id"`
	UserID    int    `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
