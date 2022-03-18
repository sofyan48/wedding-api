
#!/bin/bash

awk '{print "echo " $0 " >> .env"}' deployment/.env-stg > env.sh && sh env.sh && rm env.sh