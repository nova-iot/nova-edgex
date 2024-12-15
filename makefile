.PHONY: api
api:
	goctl api go -api api/rest.api -dir .
	goctl api plugin --plugin goctl-swagger="swagger -filename gateway.json" --api ./api/rest.api --dir ./api