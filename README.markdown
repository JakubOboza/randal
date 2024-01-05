# Randal

Simple tool that lets You establish endpoints serving random urls from a prepared list.

# Usage

start randal:

```
./bin/randal serve -f example_config.yml 
```

example config will establish three enpoints `/one`, `/two` and `/three` which serves redirect to random locations. Each have a configured destination list from which a url is drawn and served. `/four` will serve 404 error as it has invalid configuration.

```
root_url:
endpoints:
  one:
    destinations:
      - http://youtube.com/wow
      - http://google.com/wow2
      - http://bing.com/wow3
  two:
    destinations:
      - http://lambdacu.be 
  three: 
    destinations:
      - https://twitter.com/1/2/3/4
      - https://facebook.com/1/2/3/4
  four: 
    destinations:
      - "invalid%$"
```

# Build

```
make
```

to build cross platform packaged release 

```
make release
```

# Test

```
make test
```