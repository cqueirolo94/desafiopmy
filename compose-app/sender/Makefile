.PHONY: down
down:
	docker-compose down -v

.PHONY: level0
level0: #1 Orders per second
	ENV_FILE=level0.env docker-compose up

.PHONY: level1
level1: #100 Orders per second
	ENV_FILE=level1.env docker-compose up

.PHONY: level2
level2: #10000 Orders per second
	ENV_FILE=level2.env docker-compose up

.PHONY: level3
level3: #100000 Orders per second
	ENV_FILE=level3.env docker-compose up

.PHONY: 
level4: #No time restriction
	ENV_FILE=level4.env docker-compose up