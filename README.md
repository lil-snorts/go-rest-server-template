# MaxTrackr

## Aim

* Be able to record workouts
* Gain experience in working with weird tech
* Have a deployed app to look good on a CV
* (Optional) sell the app or somehow make money from it

## docker implementation

Using docker - to create a new image, run the

```bash
docker build -t bigsnorts/maxtrackr:<version.number> .
```

to then run it use

```bash
# docker run -p address:hostMachinePort:containerPort bigsnorts/maxtrackr:<version.number>
# this mapps localhost:80 to the docker port 5000
# IE
docker run -p 80:5000 bigsnorts/maxtrackr:0.1
```

## trouble shooting

if the pogram wont start after the banner art

```sh
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.4
High performance, minimalist Go web framework
https://echo.labstack.com
```

Check the port isnt in use

### Can't connect to the docker contianer?

Make sure the ports have been mapped corrently. IE `-p 70:5000` or equiv

### Can't download the Go dependencies?

try changing the download proxy

```sh
export GOPROXY=https://goproxy.io,direct`
```

## Elastic Beanstalk commands **deprecated**

```bash
# for when you mess up 
eb init -p docker maxtrackr-dev
eb list
eb deploy
eb open
```
