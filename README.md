# finnio
 CLI based application that uses the Finnhub API to perform cool actions 💎



# Table of Contents

* [Project Structure](#project-structure)
* [Use The cli](#use-the-cli)





# Project Structure


    ├── Finio                    
    │   ├── cmd             # cli commands
    │   ├── internal        # app functionality
    │   ├── tests           # Testing endpoints
    │   └── main            # Run the application


## Use The cli 

* Create a .env file based on the .env.example containing your Finnhub's API key
* Run the following command in order to retrieve the data

``
go run main.go shares
``
* You will get the current price, the previous closing of AAPL&MSFT in your console, your portfolio value and profit-loss accordingly.  
<img width="736" alt="Screenshot 2023-01-30 at 02 50 24" src="https://user-images.githubusercontent.com/19203770/215366506-1d51ea4c-3953-4db2-aa81-3f2f99cea899.png">




## Todos

1. Integration testing the whole flow of the app - from httm requests to calculations.
3. Concurrent execution of the finnhub API coms calls for Apple and Microsoft.
4. Dockerize the app.
5. Set up github actions for CI/CD.
