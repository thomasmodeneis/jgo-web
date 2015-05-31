Description
===========

This is the web project for (JGO)[https://github.com/thomasmodeneis/jgo].


Deploy
======

```
$ heroku login
$ heroku git:remote -a jgo
$ heroku buildpacks:set https://github.com/kr/heroku-buildpack-go.git
$ git add .
$ git commit -am "#2"
$ git push heroku master
$ https://jgo.herokuapp.com/ deployed to Heroku
```

License
=======

GNU 3
-----

GNU General Public License, version 3.
