package main

import "fmt"

var (
	openingState = &OpenningState{}
	closingState = &ClosingState{}
	runningState = &RunningState{}
	stoppingState = &StoppingState{}
)

/////////////////////////////////
type ILift interface {
	Open()
	Close()
	Run()
	Stop()
}

type ILiftState interface {
	ILift
	GetContext() IContext
	SetContext(context IContext)
}

type IContext interface {
	ILift
	GetLiftState() ILiftState
	SetLiftState(state ILiftState)
}

///////////////////////////////////////////////
type LiftState struct {
	context IContext
}

func (l *LiftState) GetContext() IContext {
	return l.context
}

func (l *LiftState) SetContext(context IContext) {
	l.context = context
}

///////////////////////////////////////////

type OpenningState struct {
	LiftState
}

func (l *OpenningState) Close() {
	l.context.SetLiftState(closingState)
	l.context.GetLiftState().Close()
}

func (l *OpenningState) Open() {
	fmt.Println("lift open...")
}

func (l *OpenningState) Run() {
}

func (l *OpenningState) Stop() {
}
///////////////////////////////////////////

type ClosingState struct {
	LiftState
}

func (l *ClosingState) Close() {
	fmt.Println("lift close...")
}

func (l *ClosingState) Open() {
	l.context.SetLiftState(openingState)
	l.context.GetLiftState().Open()
}

func (l *ClosingState) Run() {
	l.context.SetLiftState(runningState)
	l.context.GetLiftState().Run()
}

func (l *ClosingState) Stop() {
	l.context.SetLiftState(stoppingState)
	l.context.GetLiftState().Stop()
}
///////////////////////////////////////////

type RunningState struct {
	LiftState
}

func (l *RunningState) Close() {
}

func (l *RunningState) Open() {
}

func (l *RunningState) Run() {
	fmt.Println("lift run up to down")
}

func (l *RunningState) Stop() {
	l.context.SetLiftState(stoppingState)
	l.context.GetLiftState().Stop()
}

///////////////////////////////////////////

type StoppingState struct {
	LiftState
}

func (l *StoppingState) Close() {
}

func (l *StoppingState) Open() {
	l.context.SetLiftState(openingState)
	l.context.GetLiftState().Open()
}

func (l *StoppingState) Run() {
	l.context.SetLiftState(runningState)
	l.context.GetLiftState().Run()
}

func (l *StoppingState) Stop() {
	fmt.Println("lift stop...")
}

////////////////////////////////////////////

type Context struct {
	liftState ILiftState
}

func (c *Context) GetLiftState() ILiftState {
	return c.liftState
}

func (c *Context)SetLiftState(state ILiftState) {
	c.liftState = state
	c.liftState.SetContext(c)
}

func (c *Context)Open() {
	c.liftState.Open()
}

func (c *Context)Close() {
	c.liftState.Close()
}

func (c *Context)Run() {
	c.liftState.Run()
}

func (c *Context)Stop() {
	c.liftState.Stop()
}

//////////////////////////////////////////////

func main() {
	context := new(Context)
	context.SetLiftState(closingState)
	context.Open()
	context.Close()
	context.Run()
	context.Stop()
}

//lift open...
//lift close...
//lift run up to down
//lift stop...