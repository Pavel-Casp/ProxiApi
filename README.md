## Конфигурация

Сервис поддерживает несколько способов конфигурации (приоритет по порядку):

1. Переменные окружения
2. `.env` файл
3. `configs/config.yaml` файл
4. Значения по умолчанию

### YAML-конфигурация

Пример `configs/config.yaml`:

```yaml
api:
  base_url: "https://jsonplaceholder.typicode.com"
  request_timeout: 10s

retry:
  max_retries: 3
  wait_time: 1s
  max_wait: 5s

debug:
  enabled: false
  log_level: "info"