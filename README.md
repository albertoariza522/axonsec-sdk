# AxonSec SDK

**Morphogenetic AI · NDR L3/L4/L7 · F1=97.1% validado**

Python client for the [AxonSec API](https://developers.axonsec.net).

## Quick Start

```python
from axonsec import AxonSecClient

client = AxonSecClient("ax_beta_2026_public")

result = client.classify(
    texto="SSH brute force detected",
    metricas={"freq": 0.9, "urgencia": 0.85}
)
print(result)
# {"clasificacion": "ATAQUE", "bias_final": -0.2769, "confianza": 0.2769}
```

## Beta Key

```
ax_beta_2026_public
```

No registration required. 1000 req/day per IP.

## API Reference

**POST** `https://axonsec.net/api/ids/`

| Field | Type | Description |
|-------|------|-------------|
| `texto` | string | Event description |
| `metricas.freq` | float 0-1 | Request frequency |
| `metricas.urgencia` | float 0-1 | Urgency score |

**Response:**

```json
{
  "clasificacion": "ATAQUE | ANOMALIA | NORMAL",
  "bias_final": -0.2769,
  "confianza": 0.2769
}
```

## Links

- Portal: [developers.axonsec.net](https://developers.axonsec.net)
- Web: [axonsec.net](https://axonsec.net)
