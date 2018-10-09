#!/bin/bash

go test -bench=. -benchmem -cpuprofile profile.out
go-torch pprof fuzzy.test profile.out

