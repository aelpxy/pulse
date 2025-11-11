package presence

import (
	"encoding/json"
	"sync"
)

type Member struct {
	UserID   string                 `json:"user_id"`
	UserInfo map[string]interface{} `json:"user_info,omitempty"`
}

type ChannelMembers struct {
	members map[string]*Member
	mu      sync.RWMutex
}

func NewChannelMembers() *ChannelMembers {
	return &ChannelMembers{
		members: make(map[string]*Member),
	}
}

func (c *ChannelMembers) Add(connectionID string, member *Member) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.members[connectionID] = member
}

func (c *ChannelMembers) Remove(connectionID string) *Member {
	c.mu.Lock()
	defer c.mu.Unlock()
	member := c.members[connectionID]
	delete(c.members, connectionID)
	return member
}

func (c *ChannelMembers) Get(connectionID string) (*Member, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	member, exists := c.members[connectionID]
	return member, exists
}

func (c *ChannelMembers) GetAll() []*Member {
	c.mu.RLock()
	defer c.mu.RUnlock()

	members := make([]*Member, 0, len(c.members))
	for _, member := range c.members {
		members = append(members, member)
	}
	return members
}

func (c *ChannelMembers) GetPresenceData() *PresenceData {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ids := make([]string, 0, len(c.members))
	hash := make(map[string]map[string]interface{})

	seenUsers := make(map[string]bool)

	for _, member := range c.members {
		if !seenUsers[member.UserID] {
			ids = append(ids, member.UserID)
			seenUsers[member.UserID] = true
		}
		if member.UserInfo != nil {
			hash[member.UserID] = member.UserInfo
		}
	}

	return &PresenceData{
		IDs:   ids,
		Hash:  hash,
		Count: len(ids),
	}
}

func (c *ChannelMembers) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.members)
}

type PresenceData struct {
	IDs   []string                  `json:"ids"`
	Hash  map[string]map[string]any `json:"hash"`
	Count int                       `json:"count"`
}

type Manager struct {
	channels map[string]*ChannelMembers
	mu       sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		channels: make(map[string]*ChannelMembers),
	}
}

func (m *Manager) AddMember(channelName, connectionID string, member *Member) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.channels[channelName]; !exists {
		m.channels[channelName] = NewChannelMembers()
	}
	m.channels[channelName].Add(connectionID, member)
}

func (m *Manager) RemoveMember(channelName, connectionID string) *Member {
	m.mu.Lock()
	defer m.mu.Unlock()

	if ch, exists := m.channels[channelName]; exists {
		member := ch.Remove(connectionID)
		if ch.Count() == 0 {
			delete(m.channels, channelName)
		}
		return member
	}
	return nil
}

func (m *Manager) GetPresenceData(channelName string) *PresenceData {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if ch, exists := m.channels[channelName]; exists {
		return ch.GetPresenceData()
	}
	return &PresenceData{
		IDs:   []string{},
		Hash:  make(map[string]map[string]interface{}),
		Count: 0,
	}
}

func (m *Manager) GetMember(channelName, connectionID string) (*Member, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if ch, exists := m.channels[channelName]; exists {
		return ch.Get(connectionID)
	}
	return nil, false
}

func ParseChannelData(data string) (*Member, error) {
	var member Member
	if err := json.Unmarshal([]byte(data), &member); err != nil {
		return nil, err
	}
	return &member, nil
}
