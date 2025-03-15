package models

type Client struct {
	Id      int
	Message string
}

func (client *Client) New(id int, message string) (*Client, error) {
	return &Client{
		Id:      id,
		Message: message,
	}, nil
}
