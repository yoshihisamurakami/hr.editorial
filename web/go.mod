module web

go 1.13

require (
	controller v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.5.0
	model v0.0.0-00010101000000-000000000000 // indirect
	service v0.0.0-00010101000000-000000000000 // indirect
)

replace controller => ../controller

replace service => ../service

replace db => ../db

replace model => ../model
