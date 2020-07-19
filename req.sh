curl -si \
    --header "Content-Type: application/json" \
    --data '[{"name":"bottle of perfume","price":47.50,"type":"other","imported":true},{"name":"box of chocolates","price":10.00,"type":"food","imported":true}]' \
    https://pepper-sorleone.cloud.okteto.net/receipt
