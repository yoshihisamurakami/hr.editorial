module crawler

go 1.13

replace model => ../model

replace db => ../db

require (
	github.com/PuerkitoBio/goquery v1.5.1
	model v0.0.0-00010101000000-000000000000
)
