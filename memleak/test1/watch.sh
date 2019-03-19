echo "analyzer start........."

echo "0"
curl http://localhost:6060/debug/pprof/heap > heap.0.pprof
sleep 5

echo "1"
curl http://localhost:6060/debug/pprof/heap > heap.1.pprof
sleep 10

echo "2"
curl http://localhost:6060/debug/pprof/heap > heap.2.pprof
sleep 15

echo "3"
curl http://localhost:6060/debug/pprof/heap > heap.3.pprof
sleep 20

echo "4"
curl http://localhost:6060/debug/pprof/heap > heap.4.pprof
sleep 25

echo "5"
curl http://localhost:6060/debug/pprof/heap > heap.5.pprof
sleep 30

echo "6"
curl http://localhost:6060/debug/pprof/heap > heap.6.pprof