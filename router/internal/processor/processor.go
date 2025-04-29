package processor

import "github.com/finnmo/mqtt-router/internal/model"

type PublishMsg struct {
    Topic   string
    Payload []byte
}

// Processor is implemented once per device “type”.
type Processor interface {
    // Process takes the raw message and returns zero, one, or many
    // messages to publish. Returning nil means “drop”.
    Process(d model.Device, topic string, payload []byte) ([]PublishMsg, error)
}
