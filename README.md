## Gin + Zerolog quick demo

How to replace the default Gin logger with Zerolog.

Also how to time duration of a call out to a third party service and record it in a log. (Though arguably, you should be using a metrics library for this instead of logging.)

```console
❯ go run cmd/log/main.go 
{"level":"info","App":"logtest","Call":"main","time":"2023-10-17T15:44:35-07:00","message":"Starting server on port 8080"}
{"level":"info","foo":"bar","time":"2023-10-17T15:44:40-07:00","message":"Hello Thing"}
{"level":"info","thing":"thing","App":"logtest","Call":"main","time":"2023-10-17T15:44:40-07:00","message":"Got thing"}
{"level":"info","App":"logtest","Call":"TimeSomething","Duration":1.0001718,"time":"2023-10-17T15:44:41-07:00"}
{"level":"error","error":"this is an error","App":"logtest","Call":"TimeSomething","time":"2023-10-17T15:44:41-07:00","message":"error"}
{"level":"info","client_ip":"127.0.0.1","user_agent":"curl/7.81.0","method":"GET","path":"/ping","latency":1.0003496,"status":200,"time":"2023-10-17T15:44:41-07:00"}
```