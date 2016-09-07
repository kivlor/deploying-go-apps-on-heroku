### The Caveats of Deploying Go Apps on Heroku

A lightening talk on deploying Go apps on [![Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

##### Running the Slides

Install the present tool:

```
$ go get https://golang.org/x/tools/present
```

From the directory call `present`:

```
$ present
```

Open your web browser and visit [http://127.0.0.1:3999](http://127.0.0.1:3999)

##### Deploying the App:

Make sure you have the [Heroku toolbelt](https://toolbelt.heroku.com/) installed and you've authenticated then
run:

```
$ heroku create
```

Deploy the app using git:

```
$ git push heroku master
```

Open the app:

```
$ heroku open
```
