const { AxonSecClient } = require('./index');

const client = new AxonSecClient('ax_beta_2026_public');
client.classify('SSH brute force detected', { freq: 0.9, urgencia: 0.85 })
  .then(result => console.log(result));
// { clasificacion: 'ATAQUE', bias_final: -0.2769, confianza: 0.2769 }
