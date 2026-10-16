<?php
require_once 'AxonSec.php';

$client = new AxonSecClient('ax_beta_2026_public');
$result = $client->classify('SSH brute force detected', ['freq' => 0.9, 'urgencia' => 0.85]);
print_r($result);
// Array ( [clasificacion] => ATAQUE [bias_final] => -0.2769 [confianza] => 0.2769 )
