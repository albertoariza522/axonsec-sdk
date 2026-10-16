package axonsec

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

const apiURL = "https://axonsec.net/api/ids/"

type Client struct { APIKey string }
type Request struct {
    Texto    string                 `json:"texto"`
    Metricas map[string]interface{} `json:"metricas"`
}
type Response struct {
    Clasificacion string  `json:"clasificacion"`
    BiasFinal     float64 `json:"bias_final"`
    Confianza     float64 `json:"confianza"`
}

func NewClient(apiKey string) *Client { return &Client{APIKey: apiKey} }

func (c *Client) Classify(texto string, metricas map[string]interface{}) (*Response, error) {
    body, _ := json.Marshal(Request{Texto: texto, Metricas: metricas})
    req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", c.APIKey)
    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, fmt.Errorf("axonsec: %w", err) }
    defer resp.Body.Close()
    var result Response
    json.NewDecoder(resp.Body).Decode(&result)
    return &result, nil
}
