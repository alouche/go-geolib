go-geolib
=========

[![Build Status](https://travis-ci.org/alouche/go-geolib.png?branch=master)](https://travis-ci.org/alouche/go-geolib)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/alouche/go-geolib)

## About

Geo functions for Golang apps!

## Scope
Libraries/Functions/Utilities I often require in other projects!

## installation

	go get github.com/alouche/go-geolib

## API Documentation

[API](https://gowalker.org/github.com/alouche/go-geolib)

## Example Usage

	package main																																												

	import (
		"fmt"
		"github.com/alouche/go-geolib"
	)

	func main() {
		haversine_distance := fmt.Sprintf("%2.f", geolib.Haversine_Distance(50.116667, 8.683333, 52.516667, 13.3833))
		println("Havesine Distance:", haversine_distance, "km")
	}
