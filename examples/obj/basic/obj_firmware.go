package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/core/response"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
)

type FirmwareObject struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *FirmwareObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *FirmwareObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *FirmwareObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *FirmwareObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *FirmwareObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func NewExampleFirmwareUpdateObject(reg Registry) *FirmwareObject {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	return &FirmwareObject{
		Model: reg.GetModel(oma.OBJECT_LWM2M_FIRMWARE_UPDATE),
		Data:  data,
	}
}
