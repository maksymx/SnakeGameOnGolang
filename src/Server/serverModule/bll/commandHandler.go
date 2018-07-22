package bll

import (
	"fmt"
	"sync"

	"../../../Shared/messageTypeEnum"
	. "../../../Shared/models"
	"../../../Shared/serializer"
	"../dal"
	"math"
)

type commandHandler struct {
	lastId     uint64
	inCount    float64
	validCount float64
	mtx        *sync.Mutex
}

func (this *commandHandler) Type() HandlerType {
	return HandlerType(messageTypeEnum.CommandType)
}

func (this *commandHandler) Handle(data Message, session dal.ISession) ([]byte, error) {
	command := serializer.DecodeCommand(data.Data)
	if !this.checkAndChangeId(command.Id) {
		return make([]byte, 0), nil
	}
	session.HandleCommand(int(command.Code))
	return make([]byte, 0), nil
}

func (this *commandHandler) checkAndChangeId(id uint64) bool {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	this.inCount++
	if this.inCount > 1000 {
		fmt.Printf("Packet loss = %v%%\n", math.Max(0,100 -((this.validCount*4)/this.inCount)*100))
		this.inCount = 0
		this.validCount = 0
	}
	if this.lastId < id {
		this.lastId = id
		this.validCount++
		return true
	}
	return false
}
