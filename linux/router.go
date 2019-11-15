package linux

import (
	"errors"
	"github.com/vishvananda/netlink"
)

var (
	// ErrNotFound Link not found
	ErrNotFound = errors.New("Link not found")
)

// Router is a struct for managing Linux routing table
type Router struct {
	Link netlink.Link
	Table int
	Priority int
}


// GetRouter builds router struct
func GetRouter(name string, table int, priority int) (Router, error) {
	link, err := netlink.LinkByName(name)
	if err != nil {
		return Router{}, err
	}
	out := Router {
		Link: 		link,
		Table:		table,
		Priority: 	priority,
	}
	return out, nil
}
