"""
AxonSec SDK — Python client
https://developers.axonsec.net
"""
import requests

API_URL = "https://axonsec.net/api/ids/"

class AxonSecClient:
    def __init__(self, api_key: str = "ax_beta_2026_public"):
        self.api_key = api_key
        self.session = requests.Session()
        self.session.headers.update({"X-API-Key": api_key})

    def classify(self, texto: str, metricas: dict) -> dict:
        resp = self.session.post(API_URL, json={"texto": texto, "metricas": metricas})
        resp.raise_for_status()
        return resp.json()
