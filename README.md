# Go Task API

Eine REST-API zur Aufgabenverwaltung (Task Management), gebaut mit **Go (Gin Framework)** und **PostgreSQL**, vollständig containerisiert mit **Docker Compose**.

## Features

- **CRUD-Operationen** für Tasks (Erstellen, Lesen, Aktualisieren, Löschen)
- **Automatische Datenbankschema-Erstellung** beim ersten Start
- **Dockerized Setup** mit PostgreSQL Healthcheck und automatischer Abhängigkeitsverwaltung

## Voraussetzungen

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

## Schnellstart

1. **Repository klonen / Ordner öffnen:**
   ```bash
   cd go-task-api