# Hydrator

This is a sample hydrator as in [http://openmymind.net/Practical-SOA-Hydration-Part-1/](Practical SOA - Hydration - Part 1) and [http://openmymind.net/Practical-SOA-Hydration-Part-2/](Practical SOA - Hydration - Part 2).

It's meant as a guide, not a finished solution.

## Runing

`upstream.coffee` is the mock upstream server that returns a response needing hydration. It doesn't require any packages (aside coffee-script) and can be started with:

    coffee upstream.coffee

You can see its response at: http://localhost:4005/

You can run the hydration service by running:

    go run app/main.go

You can see the hydrated response at: http://localhost:4006
