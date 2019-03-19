curl http://localhost:6060/debug/pprof/heap > heap.0.pprof
sleep 30
curl http://localhost:6060/debug/pprof/heap > heap.1.pprof
sleep 30
curl http://localhost:6060/debug/pprof/heap > heap.2.pprof
sleep 30
curl http://localhost:6060/debug/pprof/heap > heap.3.pprof
sleep 10
curl http://localhost:6060/debug/pprof/heap > heap.4.pprof
