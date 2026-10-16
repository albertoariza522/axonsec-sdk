from axonsec import AxonSecClient

client = AxonSecClient("ax_beta_2026_public")

result = client.classify(
    texto="SSH brute force detected",
    metricas={"freq": 0.9, "urgencia": 0.85}
)
print(result)
# {"clasificacion": "ATAQUE", "bias_final": -0.2769, "confianza": 0.2769}
