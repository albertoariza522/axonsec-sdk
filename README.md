# AxonSec SDK

**Morphogenetic AI · NDR L3/L4/L7 · F1=97.1% validado**

Multi-language client for the [AxonSec API](https://developers.axonsec.net).

## Beta Key

```
ax_beta_2026_public
```
No registration required. 1000 req/day per IP.

---

## Python

```python
from axonsec import AxonSecClient
client = AxonSecClient('ax_beta_2026_public')
result = client.classify('SSH brute force', {'freq': 0.9, 'urgencia': 0.85})
print(result)
```
→ [axonsec/](axonsec/)

## Node.js

```javascript
const { AxonSecClient } = require('./node/index');
const client = new AxonSecClient('ax_beta_2026_public');
client.classify('SSH brute force', { freq: 0.9, urgencia: 0.85 }).then(console.log);
```
→ [node/](node/)

## PHP

```php
require_once 'AxonSec.php';
$client = new AxonSecClient('ax_beta_2026_public');
$result = $client->classify('SSH brute force', ['freq' => 0.9, 'urgencia' => 0.85]);
```
→ [php/](php/)

## Go

```go
client := axonsec.NewClient('ax_beta_2026_public')
result, _ := client.Classify('SSH brute force', map[string]interface{}{'freq': 0.9, 'urgencia': 0.85})
```
→ [go/](go/)

---

## API Response

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
