package processor

import "github.com/your-org/mqtt-router/internal/model"

type passthrough struct{}

func NewPassthrough() Processor { return &passthrough{} }

func (p *passthrough) Process(d model.Device, _ string, payload []byte) ([]PublishMsg, error) {
    return []PublishMsg{{Topic: d.OutputTopic, Payload: payload}}, nil
}
