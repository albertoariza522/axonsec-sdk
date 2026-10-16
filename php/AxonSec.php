<?php
class AxonSecClient {
    private $apiKey;
    private $apiUrl = 'https://axonsec.net/api/ids/';

    public function __construct($apiKey = 'ax_beta_2026_public') {
        $this->apiKey = $apiKey;
    }
    public function classify($texto, $metricas) {
        $data = json_encode(['texto' => $texto, 'metricas' => $metricas]);
        $ch = curl_init($this->apiUrl);
        curl_setopt_array($ch, [
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_POST => true,
            CURLOPT_POSTFIELDS => $data,
            CURLOPT_HTTPHEADER => [
                'Content-Type: application/json',
                'X-API-Key: ' . $this->apiKey
            ]
        ]);
        $response = curl_exec($ch);
        curl_close($ch);
        return json_decode($response, true);
    }
}
