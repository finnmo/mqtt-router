package main

import (
    "context"
    "encoding/json"
    "log"
    "os"
    "time"

    paho "github.com/eclipse/paho.golang/paho"
    "github.com/your-org/mqtt-router/config"
    "github.com/your-org/mqtt-router/internal/broker"
    "github.com/your-org/mqtt-router/internal/router"
)

func must[T any](v T, err error) T { if err != nil { log.Fatal(err) }; return v }

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    cfg := must(loadConfig("config/devices.json"))
    rtr := router.New(cfg)

    // Brokers -----------------------------------------------------------
    local := must(broker.NewLocal())
    aws   := must(broker.NewAWS())

    // subscribe once per device
    for _, d := range cfg.Devices {
        must(nil, local.Subscribe(ctx, d.InputTopic))
    }

    // local → router → aws
    go func() {
        for msg := range local.Incoming() {
            rtr.Handle(ctx, msg, func(pm router.PublishMsg) error {
                return aws.Publish(ctx, pm.Topic, pm.Payload)
            })
        }
    }()

    // block forever
    <-ctx.Done()
}

// helper ---------------------------------------------------------------
func loadConfig(path string) (model.Config, error) {
    raw, err := os.ReadFile(path)
    if err != nil { return model.Config{}, err }
    var c model.Config
    err = json.Unmarshal(raw, &c)
    // expand {site} token:
    for i := range c.Devices {
        c.Devices[i].InputTopic  = strings.ReplaceAll(c.Devices[i].InputTopic,  "{site}", c.Site)
        c.Devices[i].OutputTopic = strings.ReplaceAll(c.Devices[i].OutputTopic, "{site}", c.Site)
    }
    return c, err
}
