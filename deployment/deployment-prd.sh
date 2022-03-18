
#!/bin/bash

awk '{print "echo " $0 " >> .env"}' deployment/.env-prd > env.sh && sh env.sh && rm env.sh