package lib

import "fmt"

var (
	ErrAppAlreadyLoaded     = fmt.Errorf("application is already loaded")
	ErrAppAlreadyStarted    = fmt.Errorf("application is already started")
	ErrAppUnknown           = fmt.Errorf("unknown application name")
	ErrAppIsNotRunning      = fmt.Errorf("application is not running")
	ErrNameUnknown          = fmt.Errorf("unknown name")
	ErrNameOwner            = fmt.Errorf("not an owner")
	ErrProcessBusy          = fmt.Errorf("process is busy")
	ErrProcessMailboxFull   = fmt.Errorf("process mailbox is full")
	ErrProcessUnknown       = fmt.Errorf("unknown process")
	ErrProcessContext       = fmt.Errorf("not a Process context")
	ErrProcessIncarnation   = fmt.Errorf("process ID belongs to the previous incarnation")
	ErrProcessTerminated    = fmt.Errorf("process terminated")
	ErrMonitorUnknown       = fmt.Errorf("unknown monitor reference")
	ErrSenderUnknown        = fmt.Errorf("unknown sender")
	ErrBehaviorUnknown      = fmt.Errorf("unknown behavior")
	ErrBehaviorGroupUnknown = fmt.Errorf("unknown behavior group")
	ErrAliasUnknown         = fmt.Errorf("unknown alias")
	ErrAliasOwner           = fmt.Errorf("not an owner")
	ErrEventMismatch        = fmt.Errorf("message type mismatch")
	ErrEventUnknown         = fmt.Errorf("unknown event type")
	ErrEventOwner           = fmt.Errorf("not an owner")
	ErrEventSelf            = fmt.Errorf("monitor events from itself")
	ErrNoRoute              = fmt.Errorf("no route to node")
	ErrTaken                = fmt.Errorf("resource is taken")
	ErrFragmented           = fmt.Errorf("fragmented data")
	ErrReferenceUnknown     = fmt.Errorf("unknown reference")

	ErrRouteName = fmt.Errorf("incorrect route name")

	ErrTimeout         = fmt.Errorf("timed out")
	ErrUnsupported     = fmt.Errorf("not supported")
	ErrUnknown         = fmt.Errorf("unknown")
	ErrPeerUnsupported = fmt.Errorf("peer does not support this feature")

	ErrUnsupportedRequest = fmt.Errorf("unsupported request")
	ErrServerTerminated   = fmt.Errorf("server terminated")

	ErrProxyUnknownRequest   = fmt.Errorf("unknown proxy request")
	ErrProxyTransitDisabled  = fmt.Errorf("proxy feature disabled")
	ErrProxyNoRoute          = fmt.Errorf("no proxy route to node")
	ErrProxyConnect          = fmt.Errorf("can't establish proxy connection")
	ErrProxyHopExceeded      = fmt.Errorf("proxy hop is exceeded")
	ErrProxyLoopDetected     = fmt.Errorf("proxy loop detected")
	ErrProxyPathTooLong      = fmt.Errorf("proxy path too long")
	ErrProxySessionUnknown   = fmt.Errorf("unknown session id")
	ErrProxySessionDuplicate = fmt.Errorf("session is already exist")
)
