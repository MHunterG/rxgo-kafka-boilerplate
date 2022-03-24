package rxerrs_test

import (
	"errors"
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/events"
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/rxerrs"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testApp struct {
	isShutdown     bool
	errorEventName string
}

func (app *testApp) Shutdown() {
	app.isShutdown = true
}
func (app *testApp) SendErrorEvent(event string, _ []byte) error {
	app.errorEventName = event
	return nil
}

func TestHandleError(t *testing.T) {
	ctx := events.Ctx{
		EventName: "randomerrortest",
	}

	tApp := testApp{isShutdown: false}

	rxerrs.HandleError(errors.New("random"), &ctx, &tApp)

	assert.True(t, tApp.isShutdown)
	assert.EqualValues(t, "error-in-boilerplate", tApp.errorEventName)
}

func TestHandleUserError(t *testing.T) {
	tApp := testApp{isShutdown: false}

	ctx := events.Ctx{
		EventName: "usererrortest",
	}

	userErrorMsg := "someusererror"
	userError := rxerrs.NewUserError(userErrorMsg, 400)
	assert.EqualValues(t, userError.Error(), userErrorMsg)

	rxerrs.HandleError(userError, &ctx, &tApp)
	assert.EqualValues(t, ctx.EventName+"-response", tApp.errorEventName)
	assert.True(t, !tApp.isShutdown)
}

func TestHandleServerError(t *testing.T) {
	tApp := testApp{isShutdown: false}

	ctx := events.Ctx{
		EventName: "servererrortest",
	}

	serverErrorMsg := "someservererror"
	serverError := rxerrs.NewServerError(serverErrorMsg, errors.New("rooterror"))
	assert.EqualValues(t, serverError.Error(), serverErrorMsg)
	rxerrs.HandleError(serverError, &ctx, &tApp)

	assert.EqualValues(t, ctx.EventName+"-response", tApp.errorEventName)
	assert.True(t, !tApp.isShutdown)
}

type testBrokenApp struct {
	isShutdown     bool
	errorEventName string
}

func (app *testBrokenApp) Shutdown() {
	app.isShutdown = true
}
func (app *testBrokenApp) SendErrorEvent(event string, _ []byte) error {
	app.errorEventName = event
	return errors.New("error send event")
}

func TestHandleServerErrorWithBrokenApp(t *testing.T) {
	tApp := testBrokenApp{isShutdown: false}

	ctx := events.Ctx{
		EventName: "servererrortest",
	}

	serverError := rxerrs.NewServerError("someservererror", errors.New("rooterror"))
	rxerrs.HandleError(serverError, &ctx, &tApp)

	assert.EqualValues(t, ctx.EventName+"-response", tApp.errorEventName)
	assert.True(t, tApp.isShutdown)
}

func TestHandleUserErrorWithBrokenApp(t *testing.T) {
	tApp := testBrokenApp{isShutdown: false}

	ctx := events.Ctx{
		EventName: "servererrortest",
	}

	serverError := rxerrs.NewUserError("someservererror", 400)
	rxerrs.HandleError(serverError, &ctx, &tApp)

	assert.EqualValues(t, ctx.EventName+"-response", tApp.errorEventName)
	assert.True(t, tApp.isShutdown)
}
