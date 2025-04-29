package model

type Device struct {
    DeviceName  string `json:"deviceName"`
    ThingName   string `json:"thingName"`
    Type        string `json:"type"`
    InputTopic  string `json:"inputTopic"`
    OutputTopic string `json:"outputTopic"`
}

type Config struct {
    Site    string   `json:"site"`
    Devices []Device `json:"devices"`
}
