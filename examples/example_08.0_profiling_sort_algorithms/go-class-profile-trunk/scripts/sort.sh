#! /usr/bin/env bash

for i in {1..3}
do
	curl -s http://localhost:8080/insert?loop=0 -o image/img-in-$i.gif
	sleep 1
	curl -s http://localhost:8080/qsortm?loop=0 -o image/img-qm-$i.gif
	sleep 1
	curl -s http://localhost:8080/qsortf?loop=0 -o image/img-qf-$i.gif
	sleep 1
	curl -s http://localhost:8080/qsort3?loop=0 -o image/img-q3-$i.gif
	sleep 1
done
echo "IMAGES DONE"
