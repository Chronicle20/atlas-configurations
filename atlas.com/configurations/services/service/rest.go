package service

import (
	"atlas-configurations/services/task"
	"encoding/json"
)

// RestModel is a unified model for all service types
type RestModel struct {
	Id      string           `json:"-"`
	Tasks   []task.RestModel `json:"tasks"`
	Subtype string           `json:"subtype"`
	SubData json.RawMessage  `json:"subData,omitempty"`
}

func (r RestModel) GetName() string {
	return "services"
}

func (r RestModel) GetID() string {
	return r.Id
}

func (r *RestModel) SetID(id string) error {
	r.Id = id
	return nil
}

// LoginRestModel contains the login-specific data
type LoginRestModel struct {
	Tenants []LoginTenantRestModel `json:"tenants"`
}

type LoginTenantRestModel struct {
	Id   string `json:"id"`
	Port int    `json:"port"`
}

// ChannelRestModel contains the channel-specific data
type ChannelRestModel struct {
	Tenants []ChannelTenantRestModel `json:"tenants"`
}

type ChannelTenantRestModel struct {
	Id        string                  `json:"id"`
	IPAddress string                  `json:"ipAddress"`
	Worlds    []ChannelWorldRestModel `json:"worlds"`
}

type ChannelWorldRestModel struct {
	Id       byte                      `json:"id"`
	Channels []ChannelChannelRestModel `json:"channels"`
}

type ChannelChannelRestModel struct {
	Id   byte `json:"id"`
	Port int  `json:"port"`
}
