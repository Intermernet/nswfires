#!/bin/bash

cd ~mike/nswfires
git pull
curl http://www.rfs.nsw.gov.au/feeds/majorIncidents.json -o ~mike/nswfires/docs/majorIncidents.json
curl 'https://hotspots.dea.ga.gov.au/geoserver/wfs?service=wfs&version=2.0.0&request=GetFeature&outputFormat=application/json&typeNames=public:hotspots_three_days' -o ~/mike/nswfires/input.json
~mike/nswfires/scripts/clean_hotspots/clean_hotspots -input ~mike/nswfires/input.json -output ~mike/nswfires/docs/features.json
rm -f ~mike/nswfires/input.json
git commit -am "Update Incident and HotSpot DBs"
git push gh_nswfires master

