setup:
	if [ -e ./dist ]; then rm -rf ./dist; fi; mkdir ./dist; mkdir -p ./dist/darwin/bin; mkdir -p ./dist/linux/bin

build:
	@echo "******************"
	@echo "**** Building ****"
	@echo "******************"

	GO111MODULE=on GOOS=linux go build -mod=vendor ../../cmd/activity_manager_bootstrap.go;cp dist/linux/bin/activity_manager_bootstrap ../docker;mv ./activity_manager_bootstrap dist/linux/bin/activity_manager_bootstrap


docker-build-prep:
	cd docker;cp ../dist/linux/bin/activity_manager_bootstrap .

docker-build-cleanup:
	cd ../docker;rm activity_manager_bootstrap

release-tag: build

	cd ../; git tag `cat .version`
	git push --tags

release: release-tag
	cd ../docker; docker build -t "singaravelan21/todo_list_manager" . --no-cache;
	cd ../docker; docker tag "singaravelan21/todo_list_manager" "singaravelan21/todo_list_manager:`cat ../.version`";
	cd ../docker; docker push "singaravelan21/todo_list_manager:`cat ../.version`"
	make docker-build-cleanup

