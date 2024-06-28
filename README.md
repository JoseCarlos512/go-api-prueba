# go-api-prueba

curl -X POST   http://localhost:3100/qr-factorization   -H 'Content-Type: application/json'   -d '{
    "matrix": [
      [1, 2, 3],
      [4, 5, 6],
      [7, 8, 9]
    ]
  }'