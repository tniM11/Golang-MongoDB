.PHONY: install-mongo init-db

install-mongo:
	@echo "Installing MongoDB (Debian/Ubuntu example)"
	sudo apt-get update && sudo apt-get install -y mongodb

init-db:
        @echo "Starting MongoDB using docker-compose"
        docker-compose up -d
