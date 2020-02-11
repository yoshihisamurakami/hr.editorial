module crawler

go 1.13

replace model => ../model

replace db => ../db

replace service => ../service

require (
	github.com/PuerkitoBio/goquery v1.5.1
	golang.org/x/text v0.3.0
	model v0.0.0-00010101000000-000000000000
	service v0.0.0-00010101000000-000000000000
)
