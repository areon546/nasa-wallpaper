
help:
	@echo "Usage:"
	@echo "		dl	: copies config to relevant directory, and runs cmd/download-photo"
	@echo "		copy	: copies nasa-photos.toml to the relevant config folder I made for this project"
	@echo "		setup	: compiles the set-background program, and moves relevant files around, primarily used in setting up the service"
	@echo "			systemd (alias for setup)"
	@echo "		test	: (testing purposes) sets my homescreen to example.com jpeg, sourced from wikipedia"
	@echo "		reset	: (testing purposes) resets homescreen to blank gray"
	@

copy:
	cp ./draig-nasa-photos.toml ~/.config/dragon/nasa/wallpaper.toml

dl: copy 
	go run cmd/set-background/set-background.go

setup: copy
	# compile and copy set-background command 
	go build cmd/set-background/set-background.go 
	cp ./set-background ~/.config/dragon/nasa/
	# .config systemd user is apparently a recommendation for 
	# where to put services files that run within userspace,
	# i placed the service and timer files there
	# so that they run regularly
	cp nasa.service ~/.config/systemd/user/nasa.service
	cp nasa.timer ~/.config/systemd/user/nasa.timer


run:
	go run cmd/set-background/set-background.go

test:
	feh --bg-max https://upload.wikimedia.org/wikipedia/commons/5/52/Exampledotcom_2025.png

reset: 
	xsetroot -solid "#333333"

systemd: setup 
	systemctl --user daemon-reload
	systemctl --user restart nasa.timer
	systemctl --user restart nasa.service
	systemctl --user enable nasa.timer

journal:
	journalctl --user -xeu nasa
