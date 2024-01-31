rebuild:
	docker rmi motionserver-app --force && docker compose up -d
	docker network connect motionapi motionserver-app-1
