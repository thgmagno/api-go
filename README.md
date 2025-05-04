# Api-Go

API desenvolvida em Go para centralizar funcionalidades utilizadas por diferentes aplicações.

O projeto é pensado para ser modular, servindo como base para lidar com tarefas e serviços úteis.

[![Deploy na Railway](https://img.shields.io/badge/railway-online-success?style=flat&logo=railway)](https://railway.com/)

## Tecnologias

- **Go**
- **Gin** – Web framework leve e rápido
- **Redis** – Armazenamento rápido em memória

## Funcionalidades

- [x] **Encurtar URL**  
  Recebe uma URL longa e retorna um link curto.

  **Endpoint:**  
  `POST https://api-go-thgmagno.up.railway.app/shorten-url`

  **Body JSON:**
  ```json
  {
    "url": "https://exemplo.com"
  }
