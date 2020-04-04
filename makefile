build:
	@echo "=============building============="
	docker-compose up -d --build

up:
	@echo "=============starting============="
	docker-compose up -d


logs:
	docker-compose logs -f

down:
	docker-compose down

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f
