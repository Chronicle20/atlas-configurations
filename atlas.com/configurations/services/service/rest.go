package service

import (
	"atlas-configurations/services/task"
	"encoding/json"
	"errors"
	"fmt"
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

// GetLoginData returns the login-specific data if the RestModel is of login type
func (r RestModel) GetLoginData() (*LoginRestModel, error) {
	if r.Subtype != "login-service" {
		return nil, errors.New(fmt.Sprintf("RestModel is not of login type, actual type: %s", r.Subtype))
	}

	var loginData LoginRestModel
	err := json.Unmarshal(r.SubData, &loginData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal login data: %w", err)
	}

	return &loginData, nil
}

// GetChannelData returns the channel-specific data if the RestModel is of channel type
func (r RestModel) GetChannelData() (*ChannelRestModel, error) {
	if r.Subtype != "channel-service" {
		return nil, errors.New(fmt.Sprintf("RestModel is not of channel type, actual type: %s", r.Subtype))
	}

	var channelData ChannelRestModel
	err := json.Unmarshal(r.SubData, &channelData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal channel data: %w", err)
	}

	return &channelData, nil
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
