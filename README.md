# Estudos com Fiber

> **Execute a criação das extensões abaixo em seus postgres antes de rodar as migrations**

```sql
CREATE USER estudos;

ALTER USER estudos WITH ENCRYPTED password 'teste123';

GRANT ALL privileges on database fiberestudos to estudos;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT * FROM pg_available_extensions;

SELECT * FROM pg_extension;
```

Desenvolvimento:

`make dev`

-   https://www.youtube.com/watch?v=MfFi4Gt-tos

## Utilize o AIR para live reload

Instale no sistema operacional o binario do air. É mais fácil

fonte: https://github.com/cosmtrek/air
