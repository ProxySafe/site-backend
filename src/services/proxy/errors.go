package proxy

import "fmt"

type ErrNoAvailableProxies struct {
	Amount int
}

func (e *ErrNoAvailableProxies) Error() string {
	return fmt.Sprintf("can not find %d proxies", e.Amount)
}

func IsErrNoAvailableProxies(err error) bool {
	_, ok := err.(*ErrNoAvailableProxies)
	return ok
}
