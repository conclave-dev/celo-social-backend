package kvstore

// Profile is mutable user data
type Profile struct {
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Photo   string   `json:"photo"`
	Details string   `json:"details"`
	Website string   `json:"website"`
	Contact Contact  `json:"contact"`
	Members []Member `json:"members"`
}

type Contact struct {
	Info string `json:"info"`
	Type string `json:"type"`
}

// Member is a member that the user has added
type Member struct {
	Name    string `json:"name"`
	Role    string `json:"role"`
	Email   string `json:"email"`
	Website string `json:"website"`
}
