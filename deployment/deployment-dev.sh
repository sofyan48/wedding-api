
#!/bin/bash

awk '{print "echo " $0 " >> .env"}' deployment/.env-dev > env.sh && sh env.sh && rm env.sh