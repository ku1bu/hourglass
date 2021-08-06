package model

import (
	"sync"
)

type Group struct {
	Services      map[string]*Service
	ToTalCapacity int32

	ClientInfo   map[string]*Client
	MinClients   int32
	UnitQuota    int32
	MinUnitQuota int32

	hungryClients int32
	averageRemind int32

	lock sync.Mutex
}

func (g *Group) AddClient(client *Client) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.ClientInfo[client.Id] = client
	g.hungryClients++
	client.Hunger = true

	g.rebalance()
}

func (g *Group) RemoveClient(id string) {
	g.lock.Lock()
	defer g.lock.Unlock()

	c, ok := g.ClientInfo[id]
	if !ok {
		return
	}

	if c.Hunger {
		g.hungryClients--
	}

	delete(g.ClientInfo, id)
	g.rebalance()
}

func (g *Group) AddService(svc *Service) {
	if svc.Capacity < 0 {
		return
	}

	g.lock.Lock()
	defer g.lock.Unlock()

	g.Services[svc.Id] = svc
	g.ToTalCapacity += svc.Capacity
	g.MinClients = svc.MinClients

	g.rebalance()
}

func (g *Group) RemoveService(id string) {
	g.lock.Lock()
	defer g.lock.Unlock()

	svc, ok := g.Services[id]
	if !ok {
		return
	}
	g.ToTalCapacity -= svc.Capacity

	if g.ToTalCapacity < 0 {
		g.ToTalCapacity = 0
	}

	delete(g.Services, id)
	g.rebalance()
}

func (g *Group) LowActive(id string, p float32) {
	if p >= 1 {
		return
	}

	g.lock.Lock()
	defer g.lock.Unlock()

	c, ok := g.ClientInfo[id]
	if !ok {
		return
	}
	c.Hunger = false
	c.Weights = p

	g.hungryClients--
	g.averageRemind += int32((float32(c.Quota) * (1 - p))) / g.hungryClients

	g.rebalance()
	g.averageRemind = 0
}

func (g *Group) rebalance() {
	if len(g.ClientInfo) == 0 {
		return
	}

	clientCount := int32(len(g.ClientInfo))
	if clientCount < g.MinClients {
		clientCount = g.MinClients
	}

	unitQuota := g.ToTalCapacity / clientCount
	if unitQuota < g.MinUnitQuota {
		unitQuota = g.MinUnitQuota
	}
	g.UnitQuota = unitQuota

	for _, c := range g.ClientInfo {
		quota := g.UnitQuota
		if c.Hunger && g.averageRemind > 0 {
			quota = g.UnitQuota + g.averageRemind
		} else if !c.Hunger && c.Weights < 1 {
			quota = int32(float32(c.Quota) * c.Weights)
		}

		if quota == c.Quota {
			continue
		}

		c.Quota = quota

		if c.WatcBack != nil {
			c.WatcBack(c)
		}
	}
}
