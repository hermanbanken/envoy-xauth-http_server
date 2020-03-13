.PHONY: validate
validate:
	docker run -v $$PWD:/data envoyproxy/envoy envoy -c "/data/envoy.yaml" --mode "validate"
