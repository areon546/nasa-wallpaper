
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
	# .config systemd user is apparently a recommendation for where to put services files that run within userspace
	cp nasa.service ~/.config/systemd/user/nasa.service
	cp nasa.timer ~/.config/systemd/user/nasa.timer


test:
	feh --bg-max https://upload.wikimedia.org/wikipedia/commons/5/52/Exampledotcom_2025.png

reset: 
	xsetroot -solid "#333333"

systemd: setup 
	systemctl --user daemon-reload
	systemctl --user restart nasa
