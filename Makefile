build:
	go build -o terraform-provider-twilio

serve:
	python3 -m http.server --directory ./public 19090

format:
	go fmt ./...
