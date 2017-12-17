# The Life of a Bacterium

This is a probabilistic evolutionary simulation of the microbiome. The simulation is written in Go(golang) https://golang.org/, which is nested into an Electron desktop GUI https://electronjs.org/.



![Alt text](/simulation/data/Original.png)



## Install instructions:

#### For Standard Use Case 

1.) Clone repository into a local directory

2.) Download and install Node.js (8.9.2 LTS) from: https://nodejs.org/en/

3.) Navigate to: The-Life-of-a-Bacterium/simulation via command line

4.) Make sure your system has permission to execute simulation.exe in the The-Life-of-a-Bacterium/simulation/ directory. This is an issue mainly for linux users.

5.) To install dependencies, run on command line (in directory from step 3): npm install    

6.) To start app, run: npm start

#### To Be Able to Modify and Build simulation.exe (for developing on the go code)

1.) Clone repository into your go/src/ directory

2.) Move the graphing package directory (gonum.org located in: The-Life-of-a-Bacterium/) directly under go/src/  
  - You should now have a directory go/src/gonum.org

3.) To build the .exe, navigate to: The-Life-of-a-Bacterium/simulation/. Run: go build

4.) Follow --Standard Use Case-- above (starting at step 2) to run app as needed.

