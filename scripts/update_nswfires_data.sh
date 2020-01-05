#!/bin/bash

cd ~mike/nswfires
git pull
curl http://www.rfs.nsw.gov.au/feeds/majorIncidents.json -o ~mike/nswfires/docs/majorIncidents.json
git commit -am "Update Incident DB"
git push gh_nswfires master

