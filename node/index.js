const fetch = require('node-fetch');

const API_URL = 'https://axonsec.net/api/ids/';

class AxonSecClient {
  constructor(apiKey = 'ax_beta_2026_public') {
    this.apiKey = apiKey;
  }
  async classify(texto, metricas) {
    const res = await fetch(API_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'X-API-Key': this.apiKey },
      body: JSON.stringify({ texto, metricas })
    });
    if (!res.ok) throw new Error(`AxonSec API error: ${res.status}`);
    return res.json();
  }
}
module.exports = { AxonSecClient };
