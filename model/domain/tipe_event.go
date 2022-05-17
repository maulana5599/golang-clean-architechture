package domain

import "time"

type TipeEvent struct {
	IdEvent   uint8     `json:"id_event"`
	NamaEvent string    `json:"nama_event"`
	CreatedAt time.Time `json:"create_at"`
	Status    uint8     `json:"status"`
}
