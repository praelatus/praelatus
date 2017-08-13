package models

import "gopkg.in/mgo.v2/bson"

// Workflow is the container for issues and keeps track of available transitions
type Workflow struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Transitions []Transition  `json:"transitions"`
}

func (w *Workflow) String() string {
	return jsonString(w)
}

// Transition contains information about what hooks to perform when performing
// a transition
type Transition struct {
	Name       string `json:"name"`
	FromStatus string `json:"from_status"`
	ToStatus   string `json:"to_status"`
	Hooks      []Hook `json:"hooks"`
}

func (t *Transition) String() string {
	return jsonString(t)
}

// Hook contains information about what webhooks to fire when a given
// transition is run.
type Hook struct {
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
	Body     string `json:"body"`
}

func (h *Hook) String() string {
	return jsonString(h)
}
