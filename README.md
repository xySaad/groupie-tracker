# Groupie Tracker

### Description

Groupie Trackers is a program written in go consists on receiving artist details from an API, manipulate the data contained in it, and display the informations (artists, locations and dates) on a user friendly [website](http://groupie.us.to/).

### Features

- The program uses go routines and channels to enable concurrent flow and fast response time.
- The program fetches banners from YouTube music and use them as a background in artist page.

## Usage

```bash
git clone https://github.com/xySaad/groupie-tracker
cd groupie-tracker
go run ./app
```

```
Server running on http://localhost:8080
```

### Tests

run the server first

```bash
go run ./app
```

then run the tests (in another terminal)

```bash
go test ./tests
```
