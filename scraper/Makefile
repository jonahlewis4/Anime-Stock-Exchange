run-prod:
	echo "Running with production environment"
	docker compose -f buildfiles/docker-compose.prod.yaml up --build

#this one is likely unnecessary because the prod compose doesn't run detached, meaning you can exit it in the terminal
stop-prod:
	echo "Stopping production environment"
	docker compose -f buildfiles/docker-compose.prod.yaml down



########################################################################################################################
#FOR DEBUGGING, YOU NEED TO HAVE GO AND DELVE INSTALLED
#GOLAND COMES WITH BOTH OF THESE
#Then you need to make a go remote on port 2345
run-debug:
	echo "Running with remote dubber port"
	docker compose -f buildfiles/docker-compose.debug.yaml up --build -d
	echo "If build succeeded, you can connect and debug remotely over port. Don't forget to stop the container with make stop-debug"
run-debug-no-detach:
	echo "Running with remote dubber port"
	docker compose -f buildfiles/docker-compose.debug.yaml up --build
	echo "If build succeeded, you can connect and debug remotely over port. Don't forget to stop the container with make stop-debug"
stop-debug:
	echo "Stopping debug environment"
	docker compose -f buildfiles/docker-compose.debug.yaml down
