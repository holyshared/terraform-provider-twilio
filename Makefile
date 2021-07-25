version = 0.1.5

build:
	go build -o terraform-provider-twilio

install:
	go build -o terraform-provider-twilio
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/holyshared/twilio/$(version)/darwin_amd64
	mv terraform-provider-twilio ~/.terraform.d/plugins/registry.terraform.io/holyshared/twilio/$(version)/darwin_amd64/terraform-provider-twilio_v$(version)
	rm examples/.terraform.lock.hcl

serve:
	python3 -m http.server --directory ./public 19090

format:
	go fmt ./...
