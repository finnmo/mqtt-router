package processor

import (
    "encoding/json"
    "fmt"
    "github.com/your-org/mqtt-router/internal/model"
)

type bosch5100i struct{}

func NewBosch5100i() Processor { return &bosch5100i{} }

func (b *bosch5100i) Process(d model.Device, inTopic string, payload []byte) ([]PublishMsg, error) {
    var src struct {
        Source struct{ Rule string } `json:"Source"`
        Data   struct{ Count int }   `json:"Data"`
        UtcTime string               `json:"UtcTime"`
    }
    if err := json.Unmarshal(payload, &src); err != nil {
        return nil, err
    }
    dir := inferDirection(inTopic) // inbound / outbound from topic suffix

    transformed := map[string]any{
        "provider":     "MQTT-Bosch",
        "iotThingName": d.ThingName,
        "PayloadData": map[string]any{
            "timestamp": src.UtcTime,
            "count":     src.Data.Count,
            "direction": dir,
            "cameraId":  d.DeviceName,
        },
    }
    out, _ := json.Marshal(transformed)

    // two topics: …/telemetry and …/status
    return []PublishMsg{
        {Topic: d.OutputTopic, Payload: out},
        {Topic: fmt.Sprintf("%s/%s", d.OutputTopic, dir), Payload: out},
    }, nil
}
