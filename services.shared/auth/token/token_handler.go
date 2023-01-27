package token

import "clean-boilerplate/shared/events"

type Handler interface {
	ListenAll()
}

type handler struct {
	srv    Service
	engine events.Engine
	topic  string
}

type HandlerConfig struct {
	Srv    Service
	Engine events.Engine
	Topic  string
}

func NewHandler(cnf HandlerConfig) Handler {
	if cnf.Topic == "" {
		cnf.Topic = "Token"
	}
	return &handler{
		srv:    cnf.Srv,
		engine: cnf.Engine,
		topic:  cnf.Topic,
	}
}

func (h *handler) ListenAll() {
	h.engine.Subscribe(h.createEventName("Created"), h.onCreate)
	h.engine.Subscribe(h.createEventName("Deleted"), h.onDelete)
	h.engine.Subscribe(h.createEventName("Extended"), h.onExtend)
}

func (h *handler) createEventName(event string) string {
	return h.topic + "." + event
}

func (h *handler) onCreate(data []byte) {
	d := &dto{}
	err := h.engine.Unmarshal(data, d)
	if err != nil {
		return
	}
	h.srv.Extend(d.Token)
}

func (h *handler) onDelete(data []byte) {
	d := &dto{}
	err := h.engine.Unmarshal(data, d)
	if err != nil {
		return
	}
	h.srv.Expire(d.Token)
}

func (h *handler) onExtend(data []byte) {
	d := &dto{}
	err := h.engine.Unmarshal(data, d)
	if err != nil {
		return
	}
	h.srv.Extend(d.Token)
}
