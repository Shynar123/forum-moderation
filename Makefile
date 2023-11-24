run:
	docker build -t forum-image .
	docker run -d -p 8000:8000 --name forum forum-image
	docker logs --follow forum
stop:
	docker stop forum 
	docker rm forum
	docker rmi forum-image
	# docker rmi golang:1.20-alpine
	# docker rmi alpine:latest
	docker image prune -f