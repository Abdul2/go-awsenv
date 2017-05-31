### short golang programme to set aws env values


#### why


	using aws cli and credentials file i can make a cli mfa authonticated api call. this creates json file
	in ~/.aws/cli/cache with session credentials. to make session credentials availble to other tools that 
	expect env values such as terraforma, this short programme helps me to avoid a manual cut and paste operation

### how

	run aws cli call with MFA to get a session (in ~/.aws/cli/cache)
	clone and run 
	go to http://localhost:8081
	check your profile file 
	

