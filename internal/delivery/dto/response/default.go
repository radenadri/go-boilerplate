package response

type Pagination struct {
	TotalItems      int  `json:"total_items"`
	TotalPages      int  `json:"total_pages"`
	CurrentPage     int  `json:"current_page"`
	ItemsPerPage    int  `json:"items_per_page"`
	HasNextPage     bool `json:"has_next_page"`
	HasPreviousPage bool `json:"has_previous_page"`
	NextPage        int  `json:"next_page"`
	PreviousPage    *int `json:"previous_page"`
}

type Links struct {
	Self     string  `json:"self"`
	First    string  `json:"first"`
	Last     string  `json:"last"`
	Next     string  `json:"next"`
	Previous *string `json:"prev"`
}

type Response struct {
	Success bool              `json:"success"`
	Data    interface{}       `json:"data,inline,omitempty"`
	Error   string            `json:"error,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

type PaginatedResponse[T any] struct {
	Success    bool       `json:"success"`
	Data       []T        `json:"data,inline,omitempty"`
	Pagination Pagination `json:"pagination"`
	Links      Links      `json:"links"`
}
