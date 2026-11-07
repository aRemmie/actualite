React 19 + TypeScript + Zustand + Webpack

## Запуск

```bash
npm install
API_URL=http://127.0.0.1:8080 npm run start
npm run typecheck
npm run build:prod
```

`API_URL` задаётся переменной окружения, по умолчанию — `http://127.0.0.1:8080` (host из swagger).

Пересоздать типы АПИ после обновления сваггера: `npm run generate-api`.
