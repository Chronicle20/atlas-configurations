package services

import (
	"atlas-configurations/services/service"
	"atlas-configurations/services/task"
	"context"
	"encoding/json"
	"errors"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServiceType string

const (
	ServiceTypeLogin   = ServiceType("login-service")
	ServiceTypeChannel = ServiceType("channel-service")
	ServiceTypeDrops   = ServiceType("drops-service")
)

type Processor struct {
	l   logrus.FieldLogger
	ctx context.Context
	db  *gorm.DB
}

func NewProcessor(l logrus.FieldLogger, ctx context.Context, db *gorm.DB) *Processor {
	p := &Processor{
		l:   l,
		ctx: ctx,
		db:  db,
	}
	return p
}

func (p *Processor) ByIdProvider(id uuid.UUID) model.Provider[interface{}] {
	return model.Map(Make)(byIdEntityProvider(p.ctx)(id)(p.db))
}

func (p *Processor) AllProvider() model.Provider[[]interface{}] {
	return model.SliceMap(Make)(allEntityProvider(p.ctx)(p.db))()
}

func Make(e Entity) (interface{}, error) {
	var rm service.RestModel

	// Create a base RestModel with common fields
	if e.Type == ServiceTypeLogin {
		// For login service
		var loginData struct {
			Tasks   []task.RestModel               `json:"tasks"`
			Tenants []service.LoginTenantRestModel `json:"tenants"`
		}

		err := json.Unmarshal(e.Data, &loginData)
		if err != nil {
			return nil, err
		}

		// Create login-specific subdata
		loginSubData := service.LoginRestModel{
			Tenants: loginData.Tenants,
		}

		// Marshal the login subdata to JSON
		subdataBytes, err := json.Marshal(loginSubData)
		if err != nil {
			return nil, err
		}

		rm.Tasks = loginData.Tasks
		rm.Subtype = string(ServiceTypeLogin)
		rm.SubData = subdataBytes

	} else if e.Type == ServiceTypeChannel {
		// For channel service
		var channelData struct {
			Tasks   []task.RestModel                 `json:"tasks"`
			Tenants []service.ChannelTenantRestModel `json:"tenants"`
		}

		err := json.Unmarshal(e.Data, &channelData)
		if err != nil {
			return nil, err
		}

		// Create channel-specific subdata
		channelSubData := service.ChannelRestModel{
			Tenants: channelData.Tenants,
		}

		// Marshal the channel subdata to JSON
		subdataBytes, err := json.Marshal(channelSubData)
		if err != nil {
			return nil, err
		}

		rm.Tasks = channelData.Tasks
		rm.Subtype = string(ServiceTypeChannel)
		rm.SubData = subdataBytes

	} else if e.Type == ServiceTypeDrops {
		// For generic/drops service
		var genericData struct {
			Tasks []task.RestModel `json:"tasks"`
		}

		err := json.Unmarshal(e.Data, &genericData)
		if err != nil {
			return nil, err
		}

		rm.Tasks = genericData.Tasks
		rm.Subtype = string(ServiceTypeDrops)
		// No subdata for generic service

	} else {
		return nil, errors.New("invalid service type")
	}

	rm.Id = e.Id.String()
	return rm, nil
}

func (p *Processor) GetAll() ([]interface{}, error) {
	return p.AllProvider()()
}

func (p *Processor) GetById(id uuid.UUID) (interface{}, error) {
	return p.ByIdProvider(id)()
}
