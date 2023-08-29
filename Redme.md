1. create file .env with content like in .env.example

2. (optional) change ports in docker-compose.yml if they are already in use

3. docker-compose build

4. docker-compose up -d

5. try to start container with app when container with db will be ready
(1. docker ps -a 2. docker start {insert here container id which have image backend-trainee-assignment-2023-go})