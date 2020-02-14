module web

go 1.13

require (
	controller v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.5.0
	github.com/jinzhu/gorm v1.9.12
	model v0.0.0-00010101000000-000000000000 // indirect
	service v0.0.0-00010101000000-000000000000 // indirect
)

replace controller => ../controller

replace service => ../service

replace db => ../db

replace model => ../model
