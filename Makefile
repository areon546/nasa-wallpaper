
help:
	@echo "Usage:"
	@echo "		dl		:	copies config to relevant directory, and runs cmd/download-photo"
	@
	@

copy:
	cp ./draig-nasa-photos.toml ~/.config/dragon/nasa/wallpaper.toml

dl: copy 
	go run cmd/set-background/set-background.go

build:
	go build cmd/set-background/set-background.go

setup: build 
	# copy relevant files to setup the daemon
	cp ./set-background ~/.config/dragon/nasa/
	cp nasa.service ~/.config/systemd/user/nasa.service
	cp nasa.timer ~/.config/systemd/user/nasa.timer
	sudo systemctl daemon-reload
