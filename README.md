Fuzzy Plantation
================

A Go experiment by Plato

Design notes
------------

Start with 100 Fuzzies  
Buy animals to increase your Fuzzy score  
Each grants 0.01 Fuzzy per second  
Each costs 1 Fuzzies  

Planned Animals
---------------

Kittens  
Puppies  
Parakeets  
Elephants  

Planned Events
--------------

Battle Royale: Spooked by a [helicopter|meteorite|gust of wind|can opener|balloon], the animals go beserk. `nAnimals *0.9`  
Field Trip: A busload of happy [kids|convicts|judges|billionaires|kings] visit your Plantation. `Fuzzies *1.1`  
Thunderstorm: Zap. `nParakeets * 0.9`  
Adoption Day: Everyone's happy that some [random] got a new home! `[nRandom] * 0.9, Fuzzies *= (1 + (nRandom * .01/nTotal))`  
Population Explosion: `nRandom *= 1.5`  
Picky Eaters: Starvation hits some [randoms] due to their insistence on eating only [caviar|meatballs|spaghetti|snails|mangos]  


Planned Milestones:
-------------------

nAll == 1: Animal Lover  
nAll == 10: Noisy Backyard  
nAll == 20: Smelly Backyard  
nAll == 50: Move to the Country  
nAll == 100: Time for some fences  
nAll == 150: Gonna need a bigger barn  
nAll == 200: Help Needed  
nAll == 250: RFID Tagging  