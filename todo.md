# TODOs 

## Small 

- add function to get API map keys 

## What I want to do next with the project 

I want to setup a command, function, (or background server, but that is expensive), 
that I can call that caches the result, 
and checks it first. 
- The cache result doesn't consider the API call specifically, but the image selected, 
as I want random calls to actually be random. 
 - May have to ignore 'random' results anyhow

I need to setup a way to handle using potentially other API types for simplicity. 
I think the way to deal with this is to create different datatypes for it as an array, and not, 
and then just cycle through the handlers based on whether the schema of the JSON matches 
based on a certain rule

I need a way to create a backup chain for API calls to make,
in case one fails.
- Not too complicated, just while main loop over APIs in config
in the same order
