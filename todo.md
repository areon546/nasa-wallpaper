# TODOs 

## Small 

- add function to get API map keys 

## What I want to do next with the project 

I want to setup a command, function, (or background server, but that is expensive), 
that I can call that caches the result, 
and checks it first. 
- The cache result doesn't consider the API call specifically, but the image selected, 
as I want random calls to actually be random. 

I need to setup a way to handle using potentially other API types for simplicity. 
eg API, API-array, etc. Idea:
```go 
var api any 
switch api.type {
    case APOD:
      api = newApod()
    case APOD-array:
      api = newApod()
    default:
      break;
  }

// need to reliably store which datatype is stored in which location, 
// which is a bit more awkawrd
```

I need a way to create a backup chain for API calls to make,
in case one fails.
- Not too complicated, just while main loop over APIs in config
if it fails
