package rxerrs

import (
	log "github.com/sirupsen/logrus"
	"reactive-kafka-boilerplate/app/events"
)

type App interface {
	Shutdown()
	SendErrorEvent(eventName string, value []byte) error
}

func HandleError(err error, c *events.Ctx, app App) {
	userErr, ok := err.(UserError)
	if ok {
		errProduce := app.SendErrorEvent(c.EventName+"-response", userErr.GetJson())
		if errProduce != nil {
			app.Shutdown()
			log.Error(errProduce)
		}
		return
	}

	serverErr, serverOk := err.(ServerError)
	if serverOk {
		errProduce := app.SendErrorEvent(c.EventName+"-response", serverErr.GetJson())
		if errProduce != nil {
			app.Shutdown()
			log.Error(errProduce)
		}
		return
	}

	errProduce := app.SendErrorEvent("error-in-boilerplate", []byte(err.Error()))
	if errProduce != nil {
		app.Shutdown()
		log.Error(errProduce)
		return
	}
	app.Shutdown()
	log.Error(err)
}
