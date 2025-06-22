FROM postgres:17.5

# Define o nome do banco e o usuário padrão
ENV POSTGRES_DB=dev
ENV POSTGRES_USER=postgres

# Copia o dump SQL para o diretório de inicialização automática
COPY dump.sql /docker-entrypoint-initdb.d/dump.sql

