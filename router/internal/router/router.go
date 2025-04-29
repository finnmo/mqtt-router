package router

import (
    "context"
    mqtt "github.com/eclipse/paho.golang/paho"
    "github.com/finnmo/mqtt-router/router/internal/model"
    "github.com/finnmo/mqtt-router/router/internal/processor"
    "log/slog"
    "strings"
)

type Router struct {
    processors map[string]processor.Processor // keyed by Device.Type
    devices    map[string]model.Device        // keyed by first input topic segment
}

func New(cfg model.Config) *Router {
    r := &Router{
        processors: map[string]processor.Processor{
            "vs135":      processor.NewPassthrough(),
            "bosch5100i": processor.NewBosch5100i(),
        },
        devices: make(map[string]model.Device),
    }
    // map first topic token (“cameraId”) to device
    for _, d := range cfg.Devices {
        tok := strings.SplitN(d.InputTopic, "/", 2)[0]
        r.devices[tok] = d
    }
    return r
}

func (r *Router) Handle(ctx context.Context, pub *mqtt.Publish, pubFn func(processor.PublishMsg) error) {
    tok := strings.SplitN(pub.Topic, "/", 2)[0]
    d, ok := r.devices[tok]
    if !ok {
        slog.Warn("unmapped topic, dropping", "topic", pub.Topic)
        return
    }

    proc, exists := r.processors[d.Type]
    if !exists {
        slog.Error("no processor for type", "type", d.Type)
        return
    }

    outs, err := proc.Process(d, pub.Topic, pub.Payload)
    if err != nil {
        slog.Error("process", "err", err)
        return
    }
    for _, o := range outs {
        if err := pubFn(o); err != nil {
            slog.Error("publish", "topic", o.Topic, "err", err)
        }
    }
}
