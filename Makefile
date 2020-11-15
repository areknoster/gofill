install-tools: tools.go
	cat tools.go | grep _ | grep \".*\" -o | xargs -tI % go install %

bundle: resources
	 fyne bundle resources > bundled.go
